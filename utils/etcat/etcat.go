// Copyright (c) 2021, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/emer/etable/v2/agg"
	"github.com/emer/etable/v2/etable"
	"goki.dev/gi"
)

var (
	Output    string
	OutFile   *os.File
	OutWriter *bufio.Writer
	LF        = []byte("\n")
	Delete    bool
	LogPrec   = 4
)

func main() {
	var help bool
	var avg bool
	flag.BoolVar(&help, "help", false, "if true, report usage info")
	flag.BoolVar(&avg, "avg", false, "if true, files must have same cols (ideally rows too, though not necessary), outputs average of any float-type columns across files")
	flag.StringVar(&Output, "output", "", "name of output file -- stdout if not specified")
	flag.StringVar(&Output, "o", "", "name of output file -- stdout if not specified")
	flag.BoolVar(&Delete, "delete", false, "if true, delete the source files after cat -- careful!")
	flag.BoolVar(&Delete, "d", false, "if true, delete the source files after cat -- careful!")
	flag.IntVar(&LogPrec, "prec", 4, "precision for number output -- defaults to 4")
	flag.Parse()

	files := flag.Args()

	sort.StringSlice(files).Sort()

	if Output != "" {
		OutFile, err := os.Create(Output)
		if err != nil {
			fmt.Println("Error creating output file:", err)
			os.Exit(1)
		}
		defer OutFile.Close()
		OutWriter = bufio.NewWriter(OutFile)
	} else {
		OutWriter = bufio.NewWriter(os.Stdout)
	}

	switch {
	case help || len(files) == 0:
		fmt.Printf("\netcat is a data table concatenation utility\n\tassumes all files have header lines, and only retains the header for the first file\n\t(otherwise just use regular cat)\n")
		flag.PrintDefaults()
	case avg:
		AvgCat(files)
	default:
		RawCat(files)
	}
	OutWriter.Flush()
}

// RawCat concatenates all data in one big file
func RawCat(files []string) {
	for fi, fn := range files {
		fp, err := os.Open(fn)
		if err != nil {
			fmt.Println("Error opening file: ", err)
			continue
		}
		scan := bufio.NewScanner(fp)
		li := 0
		for {
			if !scan.Scan() {
				break
			}
			ln := scan.Bytes()
			if li == 0 {
				if fi == 0 {
					OutWriter.Write(ln)
					OutWriter.Write(LF)
				}
			} else {
				OutWriter.Write(ln)
				OutWriter.Write(LF)
			}
			li++
		}
		fp.Close()
		if Delete {
			os.Remove(fn)
		}
	}
}

// AvgCat computes average across all runs
func AvgCat(files []string) {
	dts := make([]*etable.Table, 0, len(files))
	for _, fn := range files {
		dt := &etable.Table{}
		err := dt.OpenCSV(gi.FileName(fn), etable.Tab)
		if err != nil {
			fmt.Println("Error opening file: ", err)
			continue
		}
		if dt.Rows == 0 {
			fmt.Printf("File %v empty\n", fn)
			continue
		}
		dts = append(dts, dt)
	}
	if len(dts) == 0 {
		fmt.Println("No files or files are empty, exiting")
		return
	}
	avgdt := agg.MeanTables(dts)
	avgdt.SetMetaData("precision", strconv.Itoa(LogPrec))
	avgdt.SaveCSV(gi.FileName(Output), etable.Tab, etable.Headers)
}
