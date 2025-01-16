ARG GO_VERSION=1.23
ARG RUBY_VERSION=3.4
ARG ALPINE_VERSION=3.21

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS golang

FROM ruby:${RUBY_VERSION}-alpine${ALPINE_VERSION}

RUN apk update && \
    apk add --no-cache git alpine-sdk libffi-dev \
                       gdb strace binutils valgrind # for debug in container

COPY --from=golang /usr/local/go /usr/local/go
COPY --from=golang /go /go

ENV PATH=/go/bin:/usr/local/go/bin:$PATH \
    GOPATH=/go \
    GOLANGCI_LINT_VERSION=v1.60.3

RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s ${GOLANGCI_LINT_VERSION}

WORKDIR /app

COPY Gemfile Gemfile.lock ./
COPY _gem/go_gem.gemspec _gem/
COPY _gem/lib/go_gem/version.rb _gem/lib/go_gem/

RUN bundle config --local path vendor/bundle/ && \
    bundle config --local deployment true && \
    bundle install --jobs 4

COPY . .
