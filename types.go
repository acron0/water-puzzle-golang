package main

////////////////////////////////////////

type Bucket struct {
	capacity      int
	currentAmount int
}

type Action struct {
	capacity      int
	currentAmount int
}

type Node struct {
	identifier int
	depth      int
	parent     *Node
	children   []*Node
	state      BucketState
	desc       string
}

type BucketState []Bucket

////////////////////////////////////////

func (a BucketState) Equals(b BucketState) bool {
	if a == nil || b == nil {
		return false
	}
	for i := 0; i < len(a); i++ {
		if !(a[i].capacity == b[i].capacity && a[i].currentAmount == b[i].currentAmount) {
			return false
		}
	}

	return true
}

func (a BucketState) Clone() BucketState {
	newState := make([]Bucket, len(a))
	for i, _ := range a {
		newState[i] = a[i]
	}
	return newState
}
