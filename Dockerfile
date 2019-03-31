FROM golang:latest as base

RUN go get -u github.com/gocolly/colly/...
RUN go build .