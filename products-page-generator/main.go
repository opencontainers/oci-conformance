package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/Masterminds/sprig"
	"github.com/joshdk/go-junit"
	"gopkg.in/yaml.v3"
)

type (
	submission struct {
		IsOSS          bool
		LatestVersion  string
		AllVersions    []string
		Meta           submissionMeta
		BadgesMarkdown string
		ReadmeMarkdown string
		Workflows      map[string][]workflow
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

	workflow struct {
		Name      string
		Supported bool
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
	junitFilename    = "junit.xml"
	junitTestPrefix  = "OCI Distribution Conformance Tests"
	metaTypeOSS      = "distribution"
	templateFilename = "index.md.tpl"
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
`

	workflowPull              = "Pull"
	workflowPush              = "Push"
	workflowContentDiscovery  = "Content Discovery"
	workflowContentManagement = "Content Management"
)

func main() {
	os.RemoveAll(outputDir)
	os.Mkdir(outputDir, 0755)
	os.Mkdir(filepath.Join(outputDir, staticDir), 0755)
	submissions, err := getSubmissions()
	if err != nil {
		log.Fatal(err)
	}
	tpl, err := template.New(templateFilename).
		Funcs(sprig.FuncMap()).ParseFiles(templateFilename)
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
	for i := len(minorVersions) - 1; i >= 0; i-- {
		minorVersion := minorVersions[i]
		dirEntries, err := os.ReadDir(filepath.Join(rootDir, minorVersion))
		if err != nil {
			return nil, err
		}
		for _, dirEntry := range dirEntries {
			k := dirEntry.Name()
			readmeRaw, err := ioutil.ReadFile(filepath.Join(rootDir, minorVersion, k, readmeFilename))
			if err != nil {
				return nil, err
			}
			readmeParentDirPermalink := filepath.Join(staticDir, minorVersion, instructionsDir, k)
			readmeRaw = append([]byte(fmt.Sprintf(indexPermalinkTemplate, readmeParentDirPermalink)), readmeRaw...)
			readmeParentDir := filepath.Join(outputDir, readmeParentDirPermalink, readmeFilename)
			os.MkdirAll(readmeParentDir, 0755)
			err = ioutil.WriteFile(
				filepath.Join(readmeParentDir, readmeFilename),
				readmeRaw, 0644)
			if err != nil {
				return nil, err
			}
			reportRaw, err := ioutil.ReadFile(filepath.Join(rootDir, minorVersion, k, reportFilename))
			if err != nil {
				return nil, err
			}
			os.MkdirAll(filepath.Join(outputDir, staticDir, minorVersion, reportsDir, k), 0755)
			err = ioutil.WriteFile(
				filepath.Join(outputDir, staticDir, minorVersion, reportsDir, k, staticIndex),
				reportRaw, 0644)
			if err != nil {
				return nil, err
			}
			if s, ok := submissions[k]; ok {
				allVersions := append(s.AllVersions, minorVersion)
				s.AllVersions = allVersions
				junitPath := filepath.Join(rootDir, minorVersion, k, junitFilename)
				b, err := ioutil.ReadFile(junitPath)
				if err != nil {
					return nil, err
				}
				workflows, err := workflowsFromJunitBytes(b)
				if err != nil {
					return nil, err
				}
				s.Workflows[minorVersion] = workflows
				submissions[k] = s
			} else {
				var meta submissionMeta
				b, err := os.ReadFile(filepath.Join(rootDir, minorVersion, k, metaFilename))
				if err != nil {
					return nil, err
				}
				err = yaml.Unmarshal(b, &meta)
				if err != nil {
					return nil, err
				}
				readmePath := filepath.Join(rootDir, minorVersion, k, readmeFilename)
				b, err = ioutil.ReadFile(readmePath)
				if err != nil {
					return nil, err
				}
				readmeMarkdown := string(b)
				var badgesMarkdown string
				badgesPath := filepath.Join(rootDir, liveSubdir, k, badgesFilename)
				if _, err := os.Stat(badgesPath); err == nil {
					b, err := ioutil.ReadFile(badgesPath)
					if err != nil {
						return nil, err
					}
					badgesMarkdown = string(b)
				}
				junitPath := filepath.Join(rootDir, minorVersion, k, junitFilename)
				b, err = ioutil.ReadFile(junitPath)
				if err != nil {
					return nil, err
				}
				workflows, err := workflowsFromJunitBytes(b)
				if err != nil {
					return nil, err
				}
				submissions[k] = submission{
					IsOSS:          meta.Type == metaTypeOSS,
					LatestVersion:  minorVersion,
					AllVersions:    []string{minorVersion},
					Meta:           meta,
					BadgesMarkdown: badgesMarkdown,
					ReadmeMarkdown: readmeMarkdown,
					Workflows:      map[string][]workflow{minorVersion: workflows},
				}
			}
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
	sort.Sort(semver.Collection(vs))
	minorVersions := make([]string, numVersions)
	for i, v := range vs {
		minorVersions[i] = fmt.Sprintf("v%d.%d", v.Major(), v.Minor())
	}
	return minorVersions, nil
}

func workflowsFromJunitBytes(b []byte) ([]workflow, error) {
	suites, err := junit.Ingest(b)
	if err != nil {
		return nil, err
	}
	skippedTests := map[string]bool{
		workflowPull:              true,
		workflowPush:              true,
		workflowContentDiscovery:  true,
		workflowContentManagement: true,
	}
	for _, suite := range suites {
		for _, test := range suite.Tests {
			if test.Status != junit.StatusSkipped {
				for k, v := range skippedTests {
					if v {
						if strings.HasPrefix(test.Name, fmt.Sprintf("%s %s ", junitTestPrefix, k)) {
							skippedTests[k] = false
						}
					}
				}
			}
		}
	}
	workflows := []workflow{
		{workflowPull, !skippedTests[workflowPull], true},
		{workflowPush, !skippedTests[workflowPush], false},
		{workflowContentDiscovery, !skippedTests[workflowContentDiscovery], false},
		{workflowContentManagement, !skippedTests[workflowContentManagement], false},
	}
	return workflows, nil
}
