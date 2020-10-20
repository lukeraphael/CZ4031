package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ticklepoke/CZ4031/logger"

	"github.com/ticklepoke/CZ4031/bptree"
	"github.com/ticklepoke/CZ4031/tsvparser"
)

func experiment1And2(n int) *bptree.Tree {
	fmt.Println("Running experiment 1")
	t := bptree.NewTree(n, 100)
	rows := tsvparser.ParseTSV("data.tsv")
	logger.InitializeLogger("experiment1")

	i := 0
	for _, s := range rows {
		tconsts, rating, votes := s[0], s[1], s[2]
		key, _ := strconv.ParseFloat(rating, 64)
		t.Insert(key, tconsts, rating, votes)
		i++
	}

	t.BlckMngr.DisplayStatus(false)
	fmt.Println("Running experiment 2")
	logger.InitializeLogger("experiment2")
	logger.Logger.Println("B+ tree has parameter n of", n)
	logger.Logger.Println("B+ tree has height of", t.Height())
	logger.Logger.Println("Printing B+ tree structure")
	t.PrintTree()
	return t
}

func experiment3(t *bptree.Tree) {
	fmt.Println("Running experiment 3")
	logger.InitializeLogger("experiment3")
	t.BlckMngr.ResetBlocksAccessed()
	t.FindAndPrint(8.0, true)

	t.BlckMngr.GetBlocksAccessed()
}

func experiment4(t *bptree.Tree) {
	fmt.Println("Running experiment 4")
	logger.InitializeLogger("experiment4")
	t.BlckMngr.ResetBlocksAccessed()
	t.FindAndPrintRange(7.0, 9.0, true)

	t.BlckMngr.GetBlocksAccessed()
}

func experiment5(t *bptree.Tree) {
	fmt.Println("Running experiment 5")
	logger.InitializeLogger("experiment5")
	start := time.Now()
	recPtrs, _ := t.Find(7.0, false)
	t.PrintTree()
	t.Delete(7.0)
	t.FindNumDeletions()
	logger.Logger.Println("B+ tree has height of", t.Height())
	logger.Logger.Println("Printing B+ tree structure")
	logger.Logger.Println()
	t.PrintTree()

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
	experiment4(t)
	experiment5(t)
}
