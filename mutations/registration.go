package mutations

// A Registration represents a mutation that supports registration in the mutation register
type Registration struct {
	Label    string
	Mutation func(in []byte) []byte
}
