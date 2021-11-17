FROM docker.io/golang:1.17 as builder

WORKDIR /workspace

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY pkg/ pkg/
COPY main.go main.go

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o server main.go

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/server .
USER root

ENTRYPOINT ["/server"]