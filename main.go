package main

import (
	"container/list"
	"fmt"
	"os"
	"strconv"
)

////////////////////////////////////////

var targetAmount int = 6
var buckets []int = []int{9, 4, 3}
var maxMoves int = 10
var numResults int = -1
var firstResultOnly = false

////////////////////////////////////////

func ParseFlags(args []string) {

	var mode int = 0
	var localBuckets []int

	for _, v := range args {

		// mode switcher
		switch v {

		case "-target":
			mode = 1
			continue
		case "-moves":
			mode = 2
			continue
		case "-buckets":
			mode = 3
			continue
		case "-results":
			mode = 4
			continue
		case "-quick":
			mode = 5
			firstResultOnly = true
			continue
		}

		switch mode {
		case 1:
			ta, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				fmt.Println("Invalid usage")
				os.Exit(-1)
			} else {
				targetAmount = int(ta)
			}
			break
		case 2:
			mm, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				fmt.Println("Invalid usage")
				os.Exit(-1)
			} else {
				maxMoves = int(mm)
			}
			break
		case 3:
			bucket, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				fmt.Println("Invalid usage")
				os.Exit(-1)
			} else {
				localBuckets = append(localBuckets, int(bucket))
			}
		case 4:
			nr, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				fmt.Println("Invalid usage")
				os.Exit(-1)
			} else {
				numResults = int(nr)
			}

			break
		default:
			fmt.Println("Invalid usage")
			os.Exit(-1)
		}

		if len(localBuckets) > 0 {
			buckets = localBuckets
		}
	}
}

func PrintState(node *Node) {

	fmt.Printf("=> %02d [", node.depth)
	for i, v := range node.state {
		fmt.Printf("%vl: %v", v.capacity, v.currentAmount)
		if i < len(node.state)-1 {
			fmt.Printf(", ")
		}
	}
	fmt.Printf("] (%v)\n", node.desc)
}

func main() {

	ParseFlags(os.Args[1:])

	fmt.Printf("\nWater Puzzle Solver\n")
	fmt.Printf("----------\n")
	fmt.Printf("Puzzle: %v in %v\n", targetAmount, buckets[0])
	fmt.Printf("Buckets: %v\n", buckets)
	fmt.Printf("Moves Allowed: %v\n\n", maxMoves)

	workingBuckets := make([]Bucket, len(buckets))
	for bucketIndex := 0; bucketIndex < len(buckets); bucketIndex++ {
		workingBuckets[bucketIndex].capacity = buckets[bucketIndex]
		workingBuckets[bucketIndex].currentAmount = 0
	}

	nodeList, moves := DoLogic(workingBuckets, targetAmount, maxMoves, firstResultOnly)

	//// results
	if nodeList.Len() == 0 {

		fmt.Println("No results! Try increasing the amount of moves...")

	} else {

		for e := nodeList.Front(); e != nil; e = e.Next() {

			node := e.Value.(*Node)

			if node.depth > moves {
				continue
			}

			if numResults == 0 {
				break
			} else {
				numResults--
			}

			var pathList list.List
			for node.parent != nil {
				pathList.PushFront(node)
				node = node.parent
			}

			fmt.Println("Route:")

			for n := pathList.Front(); n != nil; n = n.Next() {
				path := n.Value.(*Node)
				PrintState(path)
			}
		}
	}
}
