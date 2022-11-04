package main

import (
	sp "github.com/scipipe/scipipe"
	spc "github.com/scipipe/scipipe/components"
)

func main() {
	wf := sp.NewWorkflow("workflow", 4)

	src_ref := spc.NewFileSource(wf, "in_ref", "input/ref_e_coli.fna.gz")
	src_source := spc.NewFileSource(wf, "in_source", "input/e_coli.fastq.gz")
	src_parser := spc.NewFileSource(wf, "in_parser", "input/parser.py")

	index := wf.NewProc("index", "minimap2 -d {o:ref} {i:in_ref}")
	index.SetOut("ref", "ref.mmi")
	index.In("in_ref").From(src_ref.Out())

	//fastqc := wf.NewProc("fastqc", "fastqc {i:in_source} > {o:fres}")
	//fastqc.SetOut("fres", "./output/e_coli_fastq_fastqc.html")
	//fastqc.In("in_source").From(src_source.Out())

	alignment := wf.NewProc("alignment", "minimap2 -a {i:ref} {i:in_source} > {o:ares}")
	alignment.SetOut("ares", "output/alignment.sam")
	alignment.In("in_source").From(src_source.Out())
	alignment.In("ref").From(index.Out("ref"))

	samtools := wf.NewProc("samtools", "samtools flagstat {i:ares} > {o:sres}")
	samtools.SetOut("sres", "output/result.txt")
	samtools.In("ares").From(alignment.Out("ares"))

	parse := wf.NewProc("parse", "python3 {i:in_parser} {i:sres} > {o:parsed}")
	parse.SetOut("parsed", "output/parsed.txt")
	parse.In("in_parser").From(src_parser.Out())
	parse.In("sres").From(samtools.Out("sres"))

	// Run workflow
	wf.Run()
}
