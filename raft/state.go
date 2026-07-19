package raft

// State contains the Raft state maintained by each node
type State struct {
	// Persistant state on all servers
	// (Updated on stable storage before responding to RPCs)
	//
	// Terms for recovering from a restart
	// (every node should know these entries)
	// latest term server has seen(initialized to 0 on first boot, increases monotonically)
	currentTerm Term
	// candidate Id that received vote in current term(or null if none)
	votedFor NodeID

	// Volatile state on all servers
	//
	// Index of highest log entry known to be commited
	// (Initialized to 0, increases monotonically)
	commitIndex LogIndex
	// Index of highest log entry applied to state machine
	// (Initialized to 0, increases monotonically)
	lastApplied LogIndex

	// // Role state for LeaderState reference
	// isLeader bool
	// Node type as role representing the current role of node
	NodeState NodeState
}

type LeaderState struct {
	// Volatile State on leaders
	// (Reinitialized after election)
	Progress map[NodeID]*NodeProgress
}

type NodeProgress struct {
	// Stores the progress of log entries on the node
	NextIndex  LogIndex
	MatchIndex LogIndex
}
