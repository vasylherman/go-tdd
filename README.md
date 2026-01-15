### installing go
https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/install-go
Using Modules is pretty straightforward. Select any directory outside GOPATH as the root of your project, and create a new module with the go mod init command.
```shell
go mod init go-tdd
```
#### The built-in documentation provides an overview of all available go mod commands.
```shell
go help mod
go help mod init
```
#### An improvement over the default linter can be configured using GolangCI-Lint.
```shell
brew install golangci-lint
```
#### Go's tool for viewing documentation
```shell
go install golang.org/x/pkgsite/cmd/pkgsite@latest
pkgsite -open .
```