package main

import (
	"log"

	"github.com/cpg1111/GoAlgos/binaryTree"
	"github.com/cpg1111/GoAlgos/bubbleSort"
)

func main() {
	testCase := []int{0, 4, 1, 2, 3, 8, 6}
	newTree := binaryTree.New(testCase)
	log.Println(newTree.Search(2))
	newTree.Print()
	log.Println("Before bubble sort:", testCase)
	bubbleSortedArr := bubbleSort.Run(testCase, 0)
	log.Println("After bubble sort:", bubbleSortedArr)
}
