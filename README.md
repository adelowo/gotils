### Gotils

[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE)
[![Build Status](https://img.shields.io/travis/adelowo/gbowo/master.svg?style=flat-square)](https://travis-ci.org/adelowo/gotils.svg?branch=master)

Some useful set of packages (utilities) i think i'd be needing over and over while writing web services in go.

Current Packages include :

- `Hasher`

Provides a simple wrapper for `crypto/bcrypt`.

```go

package main

import (
  "fmt"
  "github.com/adelowo/gotils/hasher"
)

func main() {
  h := hasher.NewBcryptHasher()

  hashedPassword, err := h.Hash(plainPassword)

  fmt.Println(hashedPassword, err)

  fmt.Println(h.Verify(hashed, plain))
}

```

- `Bag`

It is used to hold values in a key value format.

There is currently just one implementation of this right now and it is a `ValidatorErrorBag`. This can be attached to a view (json response) in other to provide feedback as per the user's entry.


```go

//import github.com/adelowo/gotils/bag

//Some sample handler
func postLogin(w http.ResponseWriter, r *http.Request) {

	validatorErrorBag := bag.NewValidatorErrorBag()

	r.ParseForm()

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	if email == "" {
			validatorErrorBag.Add("email", "Please provide a valid email address")
	}

	if password == "" {
		validatorErrorBag.Add("password", "Please provide your password")
	}

	if validatorErrorBag.Count() != 0 {
		sendLoginFailureResponse(w, r, validatorErrorBag)
    //sendLoginFailureResponse can make use of the `Get` and/or `Reset` method on the bag
		return
	}

	currentUser, err := model.FindUserByEmail(email)

	if err != nil {
		validatorErrorBag.Add("password", "Invalid password/email combination")
		sendLoginFailureResponse(w, r, validatorErrorBag)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)

	}
```
