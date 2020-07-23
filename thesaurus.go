package thesaurus

// Thesaurus is a thesaurus representation.
type Thesaurus interface {
	Synonyms(term string) ([]string, error)
}
