package main

import (
	"log"

	"github.com/cpg1111/GoAlgos/binaryTree"
)

func main() {
	newTree := binaryTree.New([]int{0, 4, 1, 2, 3, 8, 6})
	log.Println(newTree.Search(2))
	newTree.Print()
}
