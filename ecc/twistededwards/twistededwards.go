// Package twistededwards define unique identifier for twited edwards curves implemented in gnark-crypto
package twistededwards

// ID represent a unique ID for a twisted edwards curve
type ID uint16

const (
	UNKNOWN ID = iota
	BN254
	BLS12_377
	BLS12_381
	BLS12_381_BANDERSNATCH
	BLS24_315
	BW6_761
	BW6_633
)
