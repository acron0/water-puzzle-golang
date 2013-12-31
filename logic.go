package main

import (
	"container/list"
	"fmt"
)

////////////////////////////////////////

func StateIsUnique(buckets BucketState, parent *Node) bool {

	intParent := parent
	for intParent != nil {
		if intParent.state.Equals(buckets) {
			return false
		}
		intParent = intParent.parent
	}

	return true
}

func PerformAction(parent *Node, buckets BucketState, currentBucketIndex int, actionID int) (BucketState, string) {

	var description string = ""
	var newBuckets BucketState = nil
	exitsPerBucket := ((len(buckets) - 1) + 2)

	if actionID == (exitsPerBucket - 1) { // empty

		if buckets[currentBucketIndex].currentAmount > 0 {
			newBuckets = buckets.Clone()
			description = fmt.Sprintf("Emptied bucket %d", currentBucketIndex)
			newBuckets[currentBucketIndex].currentAmount = 0
		}
	} else if actionID == (exitsPerBucket - 2) { // fill

		if buckets[currentBucketIndex].currentAmount != buckets[currentBucketIndex].capacity {
			newBuckets = buckets.Clone()
			description = fmt.Sprintf("Filled bucket %d", currentBucketIndex)
			newBuckets[currentBucketIndex].currentAmount = newBuckets[currentBucketIndex].capacity

		}
	} else { // transfer

		transferBucketIndex := currentBucketIndex
		for i := -1; i < actionID; i++ {
			transferBucketIndex++
			if transferBucketIndex >= len(buckets) {
				transferBucketIndex = 0
			}
		}

		didTransfer := false
		newBuckets = buckets.Clone()
		for newBuckets[transferBucketIndex].currentAmount != newBuckets[transferBucketIndex].capacity &&
			newBuckets[currentBucketIndex].currentAmount > 0 {
			newBuckets[currentBucketIndex].currentAmount--
			newBuckets[transferBucketIndex].currentAmount++
			didTransfer = true
		}
		if didTransfer {
			description = fmt.Sprintf("Transferred bucket %d to bucket %d", currentBucketIndex, transferBucketIndex)
		} else {
			newBuckets = nil
		}
	}

	if newBuckets != nil && StateIsUnique(newBuckets, parent) {
		return newBuckets, description
	} else {
		return nil, ""
	}

}

func MakeDecision(parent *Node, target int, maxMoves *int, nodeList *list.List, firstResultOnly bool) bool {

	if parent.depth >= *maxMoves {
		return false
	}

	exitsPerBucket := ((len(parent.state) - 1) + 2)
	numberOfExits := len(parent.state) * exitsPerBucket

	// create
	if parent.children == nil {
		parent.children = make([]*Node, numberOfExits)
	}

	// iterate exits
	for i := 0; i < numberOfExits; i++ {

		// create node
		if parent.children[i] == nil {
			parent.children[i] = new(Node)
			parent.children[i].depth = parent.depth + 1
			parent.children[i].parent = parent

			bucket := i / exitsPerBucket
			exit := i % exitsPerBucket
			newState, desc := PerformAction(parent, parent.state, bucket, exit)

			parent.children[i].state = newState
			parent.children[i].desc = desc
			if newState != nil {

				if parent.children[i].state[0].currentAmount == target {

					nodeList.PushBack(parent.children[i])

					// we don't want any results longer than this one, so reduce the maxMoves
					*maxMoves = parent.children[i].depth
					return true
				}
			}
		}

		if parent.children[i].state != nil {
			result := MakeDecision(parent.children[i], target, maxMoves, nodeList, firstResultOnly)
			if result && firstResultOnly {
				return true
			}
		}
	}

	return false
}

func DoLogic(buckets BucketState, target int, maxMoves int, firstResultOnly bool) (list.List, int) {

	rootNode := new(Node)
	rootNode.depth = 0
	rootNode.state = buckets

	var nodeList list.List
	var moves int = maxMoves

	MakeDecision(rootNode, target, &moves, &nodeList, firstResultOnly)

	return nodeList, moves
}
