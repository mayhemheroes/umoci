FROM golang:1.21 AS builder

ENV GOTOOLCHAIN=auto

ARG project=umoci
ARG fuzzTargetPath=oci/casext
ARG fuzzTarget=FuzzCasext
ARG fuzzBin=oci-casext-fuzz

WORKDIR /src
COPY . .

RUN go test -c -fuzz="${fuzzTarget}" \
    -o "/${fuzzBin}.libfuzzer" \
    ./${fuzzTargetPath}

FROM ubuntu:22.04
COPY --from=builder /oci-casext-fuzz.libfuzzer /
