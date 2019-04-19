FROM golang:latest as base

RUN go get -u github.com/gocolly/colly/...
COPY ./ /go/src/github.com/dev-chat/baseball
RUN go build /go/src/github.com/dev-chat/baseball/

FROM base as out
COPY --from=base \
    /go/src/github.com/dev-chat/baseball/baseball\
    /go/src/github.com/dev-chat/baseball/template.html\
    /
COPY template.html ./
CMD "./baseball"