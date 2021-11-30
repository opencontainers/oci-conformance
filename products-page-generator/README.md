# product-page-generator

This is a simple Go-based tool to generate a Markdown webpage
used to display conformance across all specs and products.

To run:

```
go run main.go
```

This will produce a new `output/` directory.

You can then use [Jekyll](https://jekyllrb.com/)
(or something similar) to produce a static HTML website:

```
cd jekyll/
bundle config set --local path 'vendor/bundle'
bundle install
bundle exec jekyll serve
```
