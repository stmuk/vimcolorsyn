package vimcolorsyn

// use vim to generate colour syntax files

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// Vimcolour type implements a simple library for colour syntax highlighting based on vim.
type Vimcolour struct {
	Fin, Fout, Lang string
}

/*
New creates a colour syntax highlighted file

Example:

v := &Vimcolour{"vimcolorsyn.go","/tmp/foo.html","go"}
v.New()

*/
func (v *Vimcolour) New() {

	_, err := os.Open(v.Fin)

	if err != nil {
		log.Fatal(err)
	}

	vimCmdStr := fmt.Sprintf("let g:html_no_progress=1|syntax on|set noswapfile|set bg=light|set ft=%s|runtime syntax/2html.vim|wq! %s|quit", v.Lang, v.Fout)

	vimCmd := exec.Command("/opt/vim/bin/ex", "-E", "-c", vimCmdStr, v.Fin)
	_, err = vimCmd.Output()

	if err != nil {
		log.Fatal(err)
	}

}
