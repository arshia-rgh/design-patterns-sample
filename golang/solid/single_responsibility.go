package solid

import "fmt"

// Journal struct with responsibility to manage entries
type Journal struct {
	Entries []string
}

var entryCount = 0

// AddEntry adds entry to the entries for the Journal (single responsibility is ok)
func (j *Journal) AddEntry(entry string) int {
	entryCount++
	j.Entries = append(j.Entries, fmt.Sprintf("%v: %v", entryCount, entry))
	return entryCount
}

// RemoveEntry removes entry from the entries of the Journal (single responsibility is ok)
func (j *Journal) RemoveEntry(index int) {
	entryCount--
	j.Entries = append(j.Entries[:index], j.Entries[index+1:]...)
}

// SaveToFile breaks the single responsibility
func (j *Journal) SaveToFile(filename string) {

}

// LoadFromFile breaks the single responsibility
func (j *Journal) LoadFromFile(filename string) {

}
