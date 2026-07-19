package raft

type LogEntry struct {
	LogIndex LogIndex
	Term     Term
	Command  Command
}
