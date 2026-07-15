package btree

type Node struct {
	IsLeaf bool

	Keys   [][]byte
	Values [][]byte

	Children []uint64
}

// true/ false
// key count

// key length
// key data
// key length
// key data....

// val len
// val data
// val len 
// val data

// child data
// child data

// 30 | 40 | 50


// 	SET user1223: name shyam


// btree :---: b+ tree

// b+ tree: current version

// 