package main

import "github.com/alibaabaa/mutata/mutations"

type mutation func(in []byte) []byte
type mutationRegister map[string]mutation

func (r mutationRegister) add(mutation mutations.Registration) {
	r[mutation.Label] = mutation.Mutation
}

func getMutations() mutationRegister {
	register := make(mutationRegister)

	register.add(mutations.Base64Mutation())
	register.add(mutations.BCryptMutation())
	register.add(mutations.HexMutation())
	register.add(mutations.Md5Mutation())
	register.add(mutations.Sha1Mutation())

	return register
}
