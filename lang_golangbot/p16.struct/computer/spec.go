package computer

type Spec struct { // exported struct
	Maker string // exported field
	model string // UNexported field
	Price int    // exported field
}
