package bn256

import (
	"math/big"
	"crypto/sha256"
)

// Note that it might be a good idea to do this in internal mont arithmetic.
func hashToCurvePoint(m []byte) (*big.Int, *big.Int) {
	three := big.NewInt(3)
	one := big.NewInt(1)

	h := sha256.Sum256(m)
	x := new(big.Int).SetBytes(h[:])
	x.Mod(x, p)

	for {
		xxx := new(big.Int).Mul(x, x)
		xxx.Mul(xxx, x)
		t := new(big.Int).Add(xxx, three)

		y := new(big.Int).ModSqrt(t, p)
		if y != nil {
			return x, y
		}

		x.Add(x, one)
	}
}


