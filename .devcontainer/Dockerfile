FROM golang:1.19.1

RUN go install github.com/cweill/gotests/gotests@latest && \
  go install github.com/fatih/gomodifytags@latest && \
  go install github.com/josharian/impl@latest && \
  go install github.com/haya14busa/goplay/cmd/goplay@latest && \
  go install honnef.co/go/tools/cmd/staticcheck@latest && \
  go install golang.org/x/tools/gopls@latest

WORKDIR /workspace

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/cache/gomod \
  GOMODCACHE=/cache/gomod GOCACHE=/cache/gobuild \
  go mod download -x && \
  mkdir -p /go/pkg/mod && \
  cp -r /cache/gomod/* /go/pkg/mod
