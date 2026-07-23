package raft

type LogEntry struct {
	Index LogIndex
	Term     Term
	Command  Command
}

type LogManager struct {
	entries []LogEntry
}

func NewLogManager() *LogManager {
	return &LogManager{
		entries: []LogEntry{},
	}
}

// Methods on LogManager

// Appends entry to the in memory log manager
// Returns Log entry by automatically assigning a index
func (lm *LogManager) Append(term Term, command Command) (LogEntry, error) {
	// Find current index
	nextIndex := lm.LastIndex()+1
	// Create a new LogEntry
	entry := LogEntry{
		Index: nextIndex,
		Term: term,
		Command: command,
	}
	// Append to the in memory log
	lm.entries = append(lm.entries, entry)
	return entry, nil
}

func (lm *LogManager) AppendEntries(entries []LogEntry) error

func (lm *LogManager) Get(index LogIndex) (LogEntry, error)

// Returns the Index of the last Entry in the log
func (lm *LogManager) LastIndex() LogIndex {
	if len(lm.entries) == 0 {
		return 0
	}
	last := len(lm.entries)-1
	return lm.entries[last].Index
}

// Returns the Term of the last entry in the log
func (lm *LogManager) LastTerm() Term {
	// Last term is 0 if no entry in log
	if len(lm.entries) == 0 {
		return Term(0)
	}
	last := len(lm.entries)-1
	return lm.entries[last].Term

}

func (lm *LogManager) DeleteFrom(index LogIndex) error

func (lm *LogManager) Entries(start, end LogIndex) ([]LogEntry, error)

func (lm *LogManager) Contains(index LogIndex) bool

func (lm *LogManager) IsUpToDate(lastIndex LogIndex, lastTerm Term) bool