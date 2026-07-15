// Package btree implements a disk-backed B+ Tree.
//
// Responsibilities:
//   - Search
//   - Insert
//   - Delete
//
// The package does not perform file I/O directly.
// All page access goes through the pager package.
package btree

import (
	"github.com/iips-oss/distributed-kv/metadata"
	"github.com/iips-oss/distributed-kv/pager"
)

type BPlusTree struct {
    pager    *pager.Pager	
    metadata *metadata.Metadata
}

func New(p *pager.Pager, m *metadata.Metadata) *BPlusTree

// Search
// Insert
// Delete

// Open db
// Close db