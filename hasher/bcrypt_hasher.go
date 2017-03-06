//Simple wrapper around golang's bcrypt library
package hasher

import (
	"golang.org/x/crypto/bcrypt"
)

type BcryptHasher struct {
}

//Some cryptic hash value would be the first return value of the expression
//While an error would be the second. The hash was successful only if the error equals nil
func (b *BcryptHasher) Hash(p string) (string, error) {

	hashedInBytes, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedInBytes), nil
}

func (b *BcryptHasher) Verify(hashed, plain string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain)) == nil
}

func NewBcryptHasher() *BcryptHasher {
	return &BcryptHasher{}
}
