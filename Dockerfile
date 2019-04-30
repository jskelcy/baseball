FROM node:latest as frontendCompiler
COPY ./ /go/src/github.com/dev-chat/baseball
RUN cd /go/src/github.com/dev-chat/baseball/baseball_frontend && \
    npm i && \
    npm run build && \
    rm -rf node_modules

FROM golang:latest as backendCompiler
COPY ./ /go/src/github.com/dev-chat/baseball
COPY --from=frontendCompiler \
    go/src/github.com/dev-chat/baseball/baseball_frontend/build\
    go/src/github.com/dev-chat/baseball/baseball_frontend/build\
    /
RUN cd /go/src/github.com/dev-chat/baseball && go build .

FROM alpine:latest as out
RUN apk add --no-cache \
    libc6-compat
COPY --from=backendCompiler /go/src/github.com/dev-chat/baseball/baseball .
COPY --from=backendCompiler \
    go/src/github.com/dev-chat/baseball/baseball_frontend/build \ 
    ./baseball_frontend/build
CMD "./baseball"