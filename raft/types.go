// We will put every kind of shared types that are used throughout different packages here
package raft

// Term is like the session ID to the leadership
// Each candidate node increments the Term before Requesting for votes
type Term uint64

// LogIndex provide the index of log entries
type LogIndex uint64

// Command is just a slice of byte
type Command []byte

// Node ID is basically a unique name to the node thus a string
type NodeID string

// Role provides the role of node
type NodeState uint8

const (
	Follower NodeState = iota
	Candidate
	Leader
)
