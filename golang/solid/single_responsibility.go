package solid

import (
	"fmt"
	"net/url"
)

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
func (j *Journal) SaveToFile(fileName string) {}

// LoadFromFile breaks the single responsibility
func (j *Journal) LoadFromFile(fileName string) {}

/*
	We should separate the persistence from the journal struct
	by separating to a separated struct or stand-alone functions
*/

type Persistence struct {
	LineSeparator string
}

func (p *Persistence) SaveToFile(fileName string) {}

func (p *Persistence) LoadFromFile(fileName string) {}

func (p *Persistence) LoadFromWeb(url *url.URL) {}

/*
	Now we moved the persistence to separated component
	and everything is ok now (we could do stand-alone functions)
*/
