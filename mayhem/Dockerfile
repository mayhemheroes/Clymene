FROM golang:1.25 AS builder

COPY . /Clymene
WORKDIR /Clymene

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN GOFLAGS=-mod=mod go get github.com/dvyukov/go-fuzz/go-fuzz-dep golang.org/x/net@latest
RUN apt update && apt install -y clang

RUN go-fuzz-build -libfuzzer -o fuzzClymene.a github.com/Clymene-project/Clymene/mayhem && \
    clang -fsanitize=fuzzer fuzzClymene.a -o mayhem/fuzzClymene.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /Clymene/mayhem/fuzzClymene.libfuzzer /
