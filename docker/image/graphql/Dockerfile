FROM golang:1.21.0

WORKDIR /go/src/goal-minder

# go extentions 
RUN go install github.com/cweill/gotests/gotests@v1.6.0 && \
    go install github.com/fatih/gomodifytags@v1.16.0 && \
    go install github.com/josharian/impl@v1.1.0 && \
    go install github.com/haya14busa/goplay/cmd/goplay@v1.0.0 && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install honnef.co/go/tools/cmd/staticcheck@latest && \
    go install golang.org/x/tools/gopls@latest && \
    go install golang.org/x/tools/cmd/goimports@latest

# develop tools
RUN go install github.com/cosmtrek/air@latest && \
    go install github.com/go-task/task/v3/cmd/task@latest && \
    go install github.com/fe3dback/go-arch-lint@latest && \
    go install github.com/google/wire/cmd/wire@latest && \
    go install github.com/99designs/gqlgen@latest && \
    go install github.com/volatiletech/sqlboiler/v4@latest && \
    go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest && \
    curl -OL https://github.com/sqldef/sqldef/releases/download/v0.16.13/mysqldef_linux_amd64.tar.gz && \
    tar xf mysqldef_linux_amd64.tar.gz -C /usr/local/bin/
