package main

import (
	"fmt"
	"io"
	"os"

	"github.com/jessevdk/go-flags"

	"bom"
)

func addBom(r io.Reader, w io.Writer, target bom.BOM) {
	buf := make([]byte, 1024)

	c, _ := r.Read(buf)
	if c == 0 {
		return
	}

	b := bom.GetBom(buf)

	if b == bom.NOTBOM {
		w.Write([]byte(target))
	}
	w.Write(buf[:c])

	for {
		c, _ := r.Read(buf)
		if c == 0 {
			break
		}
		w.Write(buf[:c])
	}
}

func removeBom(r io.Reader, w io.Writer, target bom.BOM) {
	buf := make([]byte, 1024)

	c, _ := r.Read(buf)
	if c == 0 {
		return
	}

	b := bom.GetBom(buf)
	if b == target {
		w.Write(buf[len(b):c])
	} else {
		w.Write(buf[:c])
	}

	for {
		c, _ := r.Read(buf)
		if c == 0 {
			break
		}
		w.Write(buf[:c])
	}
}

func main() {
	var opts struct {
		Remove bool   `short:"r" long:"remove" description:"remove bom."`
		Target string `short:"t" long:"target" default:"8" description:"target format (8|16le|16be|32le|32be)."`
	}
	parser := flags.NewParser(&opts, flags.Default)
	parser.Usage = "[OPTION]...[FILE]"

	args, err := parser.Parse()
	if err != nil {
		return
	}

	var target bom.BOM
	switch opts.Target {
	case "8":
		target = bom.UTF8
	case "16le":
		target = bom.UTF16LE
	case "16be":
		target = bom.UTF16BE
	case "32le":
		target = bom.UTF32LE
	case "32be":
		target = bom.UTF32BE
	default:
		fmt.Fprintf(os.Stderr, "unknown target: %s\n", opts.Target)
		return
	}

	var r io.Reader = os.Stdin
	if len(args) > 0 && args[0] != "-" {
		f, err := os.Open(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot open file: %s\n", args[0])
			return
		}
		defer f.Close()
		r = f
	}

	if opts.Remove {
		removeBom(r, os.Stdout, target)
	} else {
		addBom(r, os.Stdout, target)
	}
}
