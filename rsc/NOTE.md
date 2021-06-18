| https://golang.org/doc/tutorial/getting-started#call

## Failed add mod

I failed launch .go file with below message.  
When use `go` first need to initialize module.

```shell
# /rsc
$ go mod tidy
go: go.mod file not found in current directory or any parent directory; see 'go help modules'
```

After I Initialize module. no error anymore.

```shell
go mod init hello
# generates .mod and .sum files.
```

## References

- Initialize module
  - https://stackoverflow.com/questions/64442213/cannot-find-package-rsc-io-quote