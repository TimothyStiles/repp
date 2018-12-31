package defrag

// Type is the Fragment building type to be used in the assembly
type Type int

const (
	// circular is a circular sequence of DNA, e.g.: many of Addgene's plasmids
	circular Type = 0

	// pcr fragments are those prepared by pcr, often a subselection of their parent vector
	pcr Type = 1

	// synthetic fragments are those that will be fully synthesized (ex: gBlocks)
	synthetic Type = 2

	// linear fragment, ie the type of a fragment as it was uploaded submitted and without PCR/synthesis
	existing Type = 3
)

// Fragment is a single building block stretch of DNA for assembly
type Fragment struct {
	// ID is a unique identifier for this fragment
	ID string `json:"-"`

	// URL, eg link to a vector's addgene page
	URL string `json:"url,omitempty"`

	// Cost to make the fragment
	Cost float64 `json:"costDollars"`

	// fragment's sequence (linear)
	Seq string `json:"seq,omitempty"`

	// primers necessary to create this (if pcr fragment)
	Primers []Primer `json:"primers,omitempty"`

	// Entry of this fragment In the DB that it came from
	// Used to look for off-targets
	Entry string `json:"-"`

	// Type of this fragment
	Type Type `json:"-"`
}
