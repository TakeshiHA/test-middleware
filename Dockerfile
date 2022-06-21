# Builder
FROM golang:latest as builder

RUN mkdir /build

WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/TakeshiHA/test-middleware/main
RUN cd build && git clone https://github.com/TakeshiHA/test-middleware.git

# Distribution
RUN cd build/test-middleware/main && go build 

EXPOSE 8080

ENTRYPOINT ["build/test-middleware/main/main"]
