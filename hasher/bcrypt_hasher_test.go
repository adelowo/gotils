package hasher

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var _ Hasher = NewBcryptHasher()

func ExampleBcryptHasher_Hash() {

	hasher := BcryptHasher{}

	hasher.Hash("*72t723c(#fji3)@")
}

func ExampleBcryptHasher_Verify() {

	badPassword := "my-password"

	hasher := BcryptHasher{}

	hashedPassword, _ := hasher.Hash(badPassword)

	fmt.Println(hasher.Verify(hashedPassword, badPassword))
	// Output: true
}

func TestBcryptHasherVerificationFails(t *testing.T) {
	hasher := BcryptHasher{}

	hashedPassword, err := hasher.Hash("bad-password")

	if err != nil {
		assert.FailNow(t, "An error was encountered while trying to hash the ")
	}

	assert.False(t, hasher.Verify(hashedPassword, "yet-another-bad-password"), "Hasher verification is supposed to fail")

}
