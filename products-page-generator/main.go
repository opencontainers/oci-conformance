package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/Masterminds/sprig/v3"
	"github.com/goccy/go-yaml"
	"github.com/joshdk/go-junit"
)

type (
	results struct {
		APIs map[string]status `yaml:"apis"`
		Data map[string]status `yaml:"data"`
	}

	submission struct {
		IsOSS          bool
		LatestVersion  string
		AllVersions    []string
		Meta           submissionMeta
		BadgesMarkdown string
		Workflows      map[string][]workflow
		Summary        map[string]summary
	}

	submissionMeta struct {
		Vendor           string `yaml:"vendor"`
		Name             string `yaml:"name"`
		Version          string `yaml:"version"`
		WebsiteURL       string `yaml:"website_url"`
		RepoURL          string `yaml:"repo_url"`
		DocumentationURL string `yaml:"documentation_url"`
		ProductLogoURL   string `yaml:"product_logo_url"`
		Type             string `yaml:"type"`
		Description      string `yaml:"description"`
	}

	summary struct {
		Counts map[string]int
	}

	workflow struct {
		Name      string
		Supported status
		Required  bool
	}
)

const (
	rootDir          = "../distribution-spec"
	liveSubdir       = "live"
	metaFilename     = "PRODUCT.yaml"
	badgesFilename   = "badges.md"
	readmeFilename   = "README.md"
	reportFilename   = "report.html"
	resultsFilename  = "results.yaml"
	junitFilename    = "junit.xml"
	junitTestPrefix  = "OCI Distribution Conformance Tests"
	metaTypeOSS      = "distribution"
	outputDir        = "output"
	outputFilename   = "README.md"
	staticDir        = "static"
	staticIndex      = "index.html"
	gitCommitEnvVar  = "GIT_COMMIT"
	gitCommitDefault = "unknown"
	instructionsDir  = "instructions"
	reportsDir       = "reports"

	indexPermalinkTemplate = `---
permalink: %s/index.html
---
%s
`

	workflowPull              = "Pull"
	workflowPush              = "Push"
	workflowContentDiscovery  = "Content Discovery"
	workflowContentManagement = "Content Management"
)

//go:embed index.md.tpl
var templateEmbed string

type status int

const (
	statusUnknown  status = iota // status is undefined
	statusDisabled               // test was disabled by configuration
	statusSkip                   // test was skipped
	statusPass                   // test passed
	statusFail                   // test detected a conformance failure
	statusError                  // failure of the test engine itself
)

func (s status) String() string {
	switch s {
	case statusPass:
		return "Pass"
	case statusSkip:
		return "Skip"
	case statusDisabled:
		return "Disabled"
	case statusFail:
		return "FAIL"
	case statusError:
		return "Error"
	default:
		return "Unknown"
	}
}

func (s status) MarshalText() ([]byte, error) {
	ret := s.String()
	if ret == "Unknown" {
		return []byte(ret), fmt.Errorf("unknown status %d", s)
	}
	return []byte(ret), nil
}

func (s *status) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {
	case "pass":
		*s = statusPass
	case "skip":
		*s = statusSkip
	case "disabled":
		*s = statusDisabled
	case "fail":
		*s = statusFail
	case "error":
		*s = statusError
	default:
		return fmt.Errorf("unknown status %s", string(text))
	}
	return nil
}

func main() {
	os.RemoveAll(outputDir)
	os.Mkdir(outputDir, 0755)
	os.Mkdir(filepath.Join(outputDir, staticDir), 0755)
	submissions, err := getSubmissions()
	if err != nil {
		log.Fatal(err)
	}
	// TODO: update template to compress workflows tables to only exceptions/highlights with an expansion option
	// TODO: is sprig needed?
	tpl, err := template.New("index").
		Funcs(sprig.FuncMap()).Parse(templateEmbed)
	if err != nil {
		log.Fatal(err)
	}
	gitCommit := os.Getenv(gitCommitEnvVar)
	if gitCommit == "" {
		gitCommit = gitCommitDefault
	}
	var b bytes.Buffer
	err = tpl.Execute(&b, struct {
		Submissions map[string]submission
		GitCommit   string
	}{
		Submissions: submissions,
		GitCommit:   gitCommit,
	})
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(filepath.Join(outputDir, outputFilename), b.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("New Markdown site generated to %s/", outputDir)
}

func getSubmissions() (map[string]submission, error) {
	submissions := map[string]submission{}
	minorVersions, err := getMinorVersionsSorted()
	if err != nil {
		return nil, err
	}
	slices.Reverse(minorVersions)
	for _, minorVersion := range minorVersions {
		dirProjects, err := os.ReadDir(filepath.Join(rootDir, minorVersion))
		if err != nil {
			return nil, err
		}
		for _, dirProject := range dirProjects {
			projectName := dirProject.Name()
			// import the readme with a template
			readmeOrig, err := os.ReadFile(filepath.Join(rootDir, minorVersion, projectName, readmeFilename))
			if err != nil {
				return nil, err
			}
			projectInstructionsDir := filepath.Join(staticDir, minorVersion, instructionsDir, projectName)
			projectReadmeRaw := fmt.Appendf(nil, indexPermalinkTemplate, projectInstructionsDir, readmeOrig)
			projectReadmeFilename := filepath.Join(outputDir, projectInstructionsDir, readmeFilename)
			err = os.MkdirAll(filepath.Dir(projectReadmeFilename), 0755)
			if err != nil {
				return nil, err
			}
			err = os.WriteFile(projectReadmeFilename, projectReadmeRaw, 0644)
			if err != nil {
				return nil, err
			}
			// copy the report
			projectReportRaw, err := os.ReadFile(filepath.Join(rootDir, minorVersion, projectName, reportFilename))
			if err != nil {
				return nil, err
			}
			projectReportFilename := filepath.Join(outputDir, staticDir, minorVersion, reportsDir, projectName, staticIndex)
			err = os.MkdirAll(filepath.Dir(projectReportFilename), 0755)
			if err != nil {
				return nil, err
			}
			err = os.WriteFile(projectReportFilename, projectReportRaw, 0644)
			if err != nil {
				return nil, err
			}
			// load metadata yaml if this is the first submission for the project
			if _, ok := submissions[projectName]; !ok {
				var meta submissionMeta
				projectMetaFilename := filepath.Join(rootDir, minorVersion, projectName, metaFilename)
				metaRaw, err := os.ReadFile(projectMetaFilename)
				if err != nil {
					return nil, err
				}
				err = yaml.Unmarshal(metaRaw, &meta)
				if err != nil {
					return nil, fmt.Errorf("failed to parse %s: %w", projectMetaFilename, err)
				}
				var badgesMarkdown string
				badgesPath := filepath.Join(rootDir, liveSubdir, projectName, badgesFilename)
				if _, err := os.Stat(badgesPath); err == nil {
					b, err := os.ReadFile(badgesPath)
					if err != nil {
						return nil, err
					}
					badgesMarkdown = string(b)
				}
				submissions[projectName] = submission{
					IsOSS:          meta.Type == metaTypeOSS,
					AllVersions:    []string{},
					LatestVersion:  minorVersion,
					Meta:           meta,
					BadgesMarkdown: badgesMarkdown,
					Workflows:      map[string][]workflow{},
					Summary:        map[string]summary{},
				}
			}
			s := submissions[projectName]
			s.AllVersions = append(s.AllVersions, minorVersion)
			// attempt to read results.yaml
			projectResultsFilename := filepath.Join(rootDir, minorVersion, projectName, resultsFilename)
			resultsRaw, err := os.ReadFile(projectResultsFilename)
			if err == nil {
				projectResults := results{}
				err = yaml.Unmarshal(resultsRaw, &projectResults)
				if err != nil {
					return nil, fmt.Errorf("failed to parse %s: %w", projectResultsFilename, err)
				}
				workflows := []workflow{}
				for api, s := range projectResults.APIs {
					name := fmt.Sprintf("API: %s", api)
					w := workflow{
						Name:      name,
						Supported: s,
					}
					// required API workflows
					switch api {
					case "Blob get", "Blob head", "Manifest get by digest", "Manifest get by tag", "Manifest head by digest", "Manifest head by tag":
						w.Required = true
					}
					workflows = append(workflows, w)
				}
				for data, s := range projectResults.Data {
					name := fmt.Sprintf("Data: %s", data)
					w := workflow{
						Name:      name,
						Supported: s,
					}
					// required data workflows
					switch data {
					case "Blobs sha256", "Image", "Index":
						w.Required = true
					}
					workflows = append(workflows, w)
				}
				// sort the workflows to undo random range over maps
				slices.SortFunc(workflows, func(a, b workflow) int {
					return strings.Compare(a.Name, b.Name)
				})
				s.Workflows[minorVersion] = workflows
				s.Summary[minorVersion] = genSummary(workflows)
			} else {
				// fallback to reading original conformance junit and workflows
				junitPath := filepath.Join(rootDir, minorVersion, projectName, junitFilename)
				b, err := os.ReadFile(junitPath)
				if err != nil {
					return nil, err
				}
				workflows, err := workflowsFromJunitBytes(b)
				if err != nil {
					return nil, err
				}
				s.Workflows[minorVersion] = workflows
				s.Summary[minorVersion] = genSummary(workflows)
			}
			submissions[projectName] = s
		}
	}
	return submissions, nil
}

func getMinorVersionsSorted() ([]string, error) {
	dirEntries, err := os.ReadDir(rootDir)
	if err != nil {
		return nil, err
	}
	var raw []string
	for _, dirEntry := range dirEntries {
		if name := dirEntry.Name(); name != liveSubdir {
			raw = append(raw, name)
		}
	}
	numVersions := len(raw)
	vs := make([]*semver.Version, numVersions)
	for i, r := range raw {
		v, err := semver.NewVersion(r)
		if err != nil {
			return nil, err
		}
		vs[i] = v
	}
	// TODO: make a simple semver sort to remove MasterMinds dependency
	sort.Sort(semver.Collection(vs))
	minorVersions := make([]string, numVersions)
	for i, v := range vs {
		minorVersions[i] = fmt.Sprintf("v%d.%d", v.Major(), v.Minor())
	}
	return minorVersions, nil
}

func genSummary(workflows []workflow) summary {
	ret := summary{
		Counts: map[string]int{},
	}
	for _, w := range workflows {
		ret.Counts[w.Supported.String()]++
	}
	return ret
}

func workflowsFromJunitBytes(b []byte) ([]workflow, error) {
	suites, err := junit.Ingest(b)
	if err != nil {
		return nil, err
	}
	workflows := []workflow{
		{workflowPull, statusSkip, true},
		{workflowPush, statusSkip, false},
		{workflowContentDiscovery, statusSkip, false},
		{workflowContentManagement, statusSkip, false},
	}
	for _, suite := range suites {
		for _, test := range suite.Tests {
			if test.Status != junit.StatusSkipped {
				for i, w := range workflows {
					if w.Supported == statusSkip {
						if strings.HasPrefix(test.Name, fmt.Sprintf("%s %s ", junitTestPrefix, w.Name)) {
							workflows[i].Supported = statusPass
						}
					}
				}
			}
		}
	}
	return workflows, nil
}
