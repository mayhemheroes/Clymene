FROM golang:1.18 as builder

COPY . /Clymene
WORKDIR /Clymene/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN apt update && apt install -y clang

RUN go-fuzz-build -libfuzzer -o fuzzClymene.a && \
    clang -fsanitize=fuzzer fuzzClymene.a -o fuzzClymene.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /Clymene/mayhem/fuzzClymene.libfuzzer /