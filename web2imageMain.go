////////////////////////////////////////////////////////////////////////////
// Program: web2image
// Purpose: Web to image
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

//go:generate sh -v web2imageCLIGen.sh

import (
	"fmt"
	"os"

	"github.com/mkideal/cli"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

type rootOptsT struct {
	Verbose int
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	progname  = "web2image"
	VERSION   = "0.1.0"
	buildTime = "2017-07-23"
)

var (
	rootArgv *rootT
	rootOpts rootOptsT
)

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
func main() {
	//NOTE: You can set any writer implements io.Writer
	// default writer is os.Stdout
	if err := cli.Root(root).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if rootOpts.Verbose >= 1 {
		fmt.Println("Screenshot taken.")
	}
	os.Exit(0)
}
