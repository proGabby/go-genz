package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GeneratePasscode(passcodeLength int) (string, error) {
	max := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(passcodeLength)), nil)

	randomNumber, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}

	password := fmt.Sprintf("%0*d", passcodeLength, randomNumber)

	return password, nil

}
