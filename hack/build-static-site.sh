#!/usr/bin/env bash

# build-static-site.sh
# Generates a static HTML site into _site/
# Requires Go ...and Ruby

set -o errexit
set -o nounset
set -o pipefail
set -x

# Go to root of repo, get git commit hash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "${DIR}/.."
GIT_COMMIT="$(git rev-parse HEAD)"

# Run "products-page-generator" which is a custom Go program to generate
# a Jekyll-compatible Markdown site based on all the conformance
# results submitted inside the distribution-spec/ directory
cd products-page-generator/
pwd
go version
go run main.go

# Use Jekyll to generate a general purpose static HTML site
# into _site/ (git ignored) which is later hosted via Netlify
cd jekyll/
bundle config set --local path 'vendor/bundle'
bundle install --path 'vendor/bundle'
bundle exec jekyll build
cp favicon.ico _site/
echo '/** Disable unnecessary site navbar/menu */' >> _site/assets/main.css
echo '.trigger { display: none; }' >> _site/assets/main.css
rm -rf ../../_site && mv _site ../../

# Done.
set +x
echo
echo "Site has been successfully generated to _site/."
echo
echo "Testing locally? Run the following to see the site at http://localhost:8000/"
echo
# Why not add some Python to the mix eh?
echo "    (cd _site && python3 -m http.server)"
echo
