package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/juju/errgo"
	"github.com/myitcv/neovim"
	. "gopkg.in/check.v1"
)

type NeovimGoTest struct {
	client *neovim.Client
	nvim   *exec.Cmd
}

func Test(t *testing.T) { TestingT(t) }

var _ = Suite(&NeovimGoTest{})

func (t *NeovimGoTest) SetUpTest(c *C) {
	// now start the process and wait for the socket file to be created
	t.nvim = exec.Command(os.Getenv("NEOVIM_BIN"), "-u", "/dev/null")
	t.nvim.Dir = "/tmp"

	// now we can create a new client
	client, err := neovim.NewCmdClient(t.nvim)
	if err != nil {
		log.Fatalf("Could not setup client: %v", errgo.Details(err))
	}

	// TODO need to handle nvim subprocess bombing out...

	// this is important; all tests below ignore errors...
	client.PanicOnError = true
	t.client = client
}

func (t *NeovimGoTest) TearDownTest(c *C) {
	done := make(chan struct{})
	go func() {
		state, err := t.nvim.Process.Wait()
		if err != nil {
			log.Fatalf("Process did not exit cleanly: %v, %v\n", err, state)
		}
		done <- struct{}{}
	}()
	err := t.client.Close()
	if err != nil {
		log.Fatalf("Could not close client: %v\n", err)
	}
	<-done
}

func (t *NeovimGoTest) BenchmarkBufferGetSlice(c *C) {
	// TODO this needs Neovim to be started with the right file
	// can be updated when we can use headless testing
	cb, _ := t.client.GetCurrentBuffer()
	for i := 0; i < c.N; i++ {
		bc, _ := cb.GetSlice(0, -1, true, true)
		_ = []byte(strings.Join(bc, "\n"))
	}
}

func (t *NeovimGoTest) BenchmarkParse(c *C) {
	data, err := ioutil.ReadFile("_testfiles/20k_lines.go")
	if err != nil {
		panic(err)
	}
	c.ResetTimer()
	for i := 0; i < c.N; i++ {
		if _, err := parser.ParseFile(token.NewFileSet(), "", data, parser.AllErrors|parser.ParseComments); err != nil {
			c.Fatalf("benchmark failed due to parse error: %s", err)
		}
	}
}

func (t *NeovimGoTest) BenchmarkWalk(c *C) {
	data, err := ioutil.ReadFile("_testfiles/20k_lines.go")
	if err != nil {
		panic(err)
	}
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "dummy", data, parser.AllErrors|parser.ParseComments)
	if err != nil {
		c.Fatalf("benchmark failed due to parse error: %s", err)
	}
	sg := NewSynGenerator()
	sg.fset = fset
	sg.f = f
	c.ResetTimer()
	for i := 0; i < c.N; i++ {
		sg.nodes = make(map[position]*match)
		ast.Walk(sg, f)
		for _, c := range f.Comments {
			ast.Walk(sg, c)
		}
	}
}
