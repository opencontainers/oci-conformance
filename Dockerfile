FROM chainguard/wolfi-base:latest AS build
RUN apk update && apk add git go ruby-3.3 ruby-3.3-dev ruby3.3-bundler
ADD . /work
RUN /work/hack/build-static-site.sh

FROM chainguard/nginx:latest
COPY --from=build /work/_site /usr/share/nginx/html
