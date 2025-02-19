# 1.21.4.0
FROM registry..com/infra-auto/go@sha256:e786e9e9376e01500116bdcc024b40e13294aee256fb64b23b6969fba46be18b AS builder

ARG BUILD_VERSION=0.1.0
ARG BUILD_COMMIT=nightly
ARG BUILD_DATE=nightly

COPY . /src
WORKDIR /src
RUN go build -ldflags="-s -w -X github.com/IaC/go-kcloutie/pkg/params/version.BuildTime=${BUILD_DATE} -X github.com/IaC/go-kcloutie/pkg/params/version.BuildVersion=${BUILD_VERSION} -X github.com/IaC/go-kcloutie/pkg/params/version.Commit=${BUILD_COMMIT}" -o /tmp ./cmd/gokcloutie/gokcloutie.go

FROM registry.access.redhat.com/ubi9/ubi-minimal@sha256:b7a3642d6245446da03d14482740be5f2fe58f30b9dfe001e89a39071a50edfc

ARG BINARY_NAME=go-kcloutie

COPY --from=builder /tmp/go-kcloutie /opt/app-root/go-kcloutie

# CMD /opt/app-root/go-kcloutie
ENTRYPOINT ["/opt/app-root/go-kcloutie", "run", "server"]