package btree

type Node struct {
	IsLeaf bool

	Keys   [][]byte
	Values [][]byte

	Children []uint64
}
