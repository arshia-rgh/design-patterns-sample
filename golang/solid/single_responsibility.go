package solid

import "fmt"

type Journal struct {
	Entries []string
}

var entryCount = 0

func (j *Journal) AddEntry(entry string) int {
	entryCount++
	j.Entries = append(j.Entries, fmt.Sprintf("%v: %v", entryCount, entry))
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	entryCount--
	j.Entries = append(j.Entries[:index], j.Entries[index+1:]...)
}
