## `github.com/myitcv/neogo` - a Neovim plugin for Go development

[![Build Status](https://travis-ci.org/myitcv/neogo.svg?branch=master)](https://travis-ci.org/myitcv/neogo)

[![YouTube demo of plugin in action](http://img.youtube.com/vi/r7Aw5znRfNA/0.jpg)](https://youtu.be/r7Aw5znRfNA)

A proof of concept Neovim plugin written against the [`neovim` Go package](http://godoc.org/github.com/myitcv/neovim)
to support Go development in Neovim. In particular `neogo` highlights the Go code being edited based on a `go/parser` AST
generated from the current buffer as opposed to Neovim's regex-based highlighting.

Very very alpha.

## Running the plugin

```bash
mkdir -p $HOME/.nvim/plugins/go
go get github.com/juju/errgo
go get gopkg.in/check.v1
go get github.com/myitcv/neovim
go get github.com/myitcv/neogo
go get github.com/myitcv/neovim/cmd/neovim-go-plugin-manager
$GOPATH/bin/neovim-go-plugin-manager github.com/myitcv/neogo
```

This should give some output along the following lines:

```
2015/08/25 21:39:56 Working directory is /tmp/neovim-go-plugin-manager_1440535196810988589
2015/08/25 21:39:56 Just about to go get -d -t github.com/myitcv/neogo
2015/08/25 21:39:56 Now running go test github.com/myitcv/neogo
2015/08/25 21:39:57 Generating plugin host
2015/08/25 21:39:57 Building plugin host
2015/08/25 21:39:58 Putting plugin host source into place: /home/myitcv/.nvim/plugins/go/plugin_host.go
2015/08/25 21:39:58 Putting plugin host into place: /home/myitcv/.nvim/plugins/go/plugin_host
2015/08/25 21:39:58 Done!
```

Now:

```
nvim -u $GOPATH/src/github.com/myitcv/neogo/special.vimrc test.go
```

Now write some go code and watch it highlight as you type!

e.g. try entering:

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello, playground")
}
```

## Features implemented

* syntax highlighting via [`go/parser`](http://godoc.org/go/parser) (partial)

## Features TODO list

See the [wiki](https://github.com/myitcv/neogo/wiki/TODO)

