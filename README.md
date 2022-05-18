# go-cmd-proxy
Proxy replacement for go cmd.

So you can use `go get`,`go install` or some go command else on a net proxy but without system proxy or command poroxy settings.

# usage
Rename the original go file to `original-go` (Windows needs to keep the suffix).

Place the replacement in the same folder.

# q&a
## How can i use a custom proxy host?
Just edit the source code and recompile.
