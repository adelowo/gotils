package bag

import (
	"flag"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var validatorErrorBag *ValidatorErrorBag

func TestMain(m *testing.M) {

	validatorErrorBag = NewValidatorErrorBag()

	flag.Parse()

	os.Exit(m.Run())
}

func ExampleValidatorErrorBag_Add() {

	message := "Some message you want to tell the user based on his/her username"

	//Add a message with the key "username"
	validatorErrorBag.Add("username", message)

}

func ExampleValidatorErrorBag_Get() {

	message := "Some message you want to tell the user based on his/her username"

	//Add a message with the key "username"
	validatorErrorBag.Add("username", message)

	//Fetch a non existent key
	_, err := validatorErrorBag.Get("password")

	if err != nil {
		fmt.Println(err)
	}

	//Output: Key, password does not exist on this bag

}

func ExampleValidatorErrorBag_Count() {

	validatorErrorBag.Reset()

	fmt.Println(validatorErrorBag.Count())
	//Output: 0
}

func ExampleValidatorErrorBag_Has() {

	exists := validatorErrorBag.Has("password")
	fmt.Println(exists)
	//Output: false
}

func ExampleValidatorErrorBag_Reset() {
	validatorErrorBag.Reset()

	fmt.Println(validatorErrorBag.Count())
	//Output: 0

}

func ExampleValidatorErrorBag_Delete() {
	validatorErrorBag.Add("foo", "bar")
	validatorErrorBag.Delete("foo")
}

func TestValidatorErrorBag_Delete(t *testing.T) {
	validatorErrorBag.Add("foo", "bar")
	validatorErrorBag.Delete("foo")

	_, err := validatorErrorBag.Get("foo")

	assert.EqualError(t, err, "Key, foo does not exist on this bag", "Foo should not be found on this bag")

}

func TestValidatorErrorBagHasWorksAsExpected(t *testing.T) {

	validatorErrorBag.Reset()

	assert.Exactly(t, 0, validatorErrorBag.Count(), "ValidatorErrorBag should contain no elements")

	validatorErrorBag.Add("lanre", "Used to be a PHPer. Go has sold me")

	assert.False(t, validatorErrorBag.Has("password"), "ValidatorErrorBag should not have the password key")

	val, err := validatorErrorBag.Get("lanre")

	if err != nil {
		assert.FailNow(t, "Error should be nil")
	}

	assert.Equal(t, "Used to be a PHPer. Go has sold me", val, "Go didn't sell me ?")

	assert.Equal(t, 1, validatorErrorBag.Count(), "The validatorErrorBag should contain only one key")
}
