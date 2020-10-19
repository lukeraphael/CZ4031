package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ticklepoke/CZ4031/bptree"
	"github.com/ticklepoke/CZ4031/tsvparser"
)

func experiment1And2(n int) *bptree.Tree {
	fmt.Println("================= Experiment 1 =================")
	t := bptree.NewTree(n)
	rows := tsvparser.ParseTSV("../../data.tsv")

	i := 0
	for _, s := range rows {
		if i == 1000 {
			break
		}
		tconsts, rating, votes := s[0], s[1], s[2]
		key, _ := strconv.ParseFloat(rating, 64)
		t.Insert(key, tconsts, rating, votes)
		i++
	}

	t.BlckMngr.DisplayStatus(false)
	fmt.Println("================= Experiment 2 =================")
	fmt.Println("B+ tree has parameter n of", n)
	fmt.Println("B+ tree has height of", t.Height())
	fmt.Println()
	fmt.Println("Printing B+ tree structure")
	fmt.Println()
	t.PrintTree()
	return t
}

func experiment3(t *bptree.Tree) {
	fmt.Println("================= Experiment 3 =================")
	t.BlckMngr.ResetBlocksAccessed()
	t.FindAndPrint(8.0, true)

	t.BlckMngr.GetBlocksAccessed()
}

func experiment4(t *bptree.Tree) {
	fmt.Println("================= Experiment 4 =================")
	t.BlckMngr.ResetBlocksAccessed()
	t.FindAndPrintRange(7.0, 9.0, true)

	t.BlckMngr.GetBlocksAccessed()
}

func experiment5(t *bptree.Tree) {
	fmt.Println("================= Experiment 5 =================")
	start := time.Now()
	recPtrs, _ := t.Find(7.0, false)
	t.PrintTree()
	t.Delete(7.0)
	t.FindNumDeletions()
	t.PrintHeight()
	t.PrintTree()
	// t.PrintLeaves()

	// TODO: add this into the DeleteRecord func
	for _, recPtr := range recPtrs {
		t.BlckMngr.DeleteRecord(recPtr.Value)
	}
	t.BlckMngr.DisplayStatus(false)
	elapse := time.Since(start)
	fmt.Println("Experiment 5 elapsed time: ", elapse)
}

func main() {
	n := 5
	t := experiment1And2(n)
	experiment3(t)
	// experiment4(t)
	// experiment5(t)
}
