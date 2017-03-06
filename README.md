### Gotils 

[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE)
[![Build Status](https://img.shields.io/travis/adelowo/gbowo/master.svg?style=flat-square)](https://travis-ci.org/adelowo/gotils.svg?branch=master)

Some useful set of packages (utilities) i think i'd be needing over and over while writing web services in go.

Current Packages include :

> TODO => Update doc

- `Bag` 

Well, this is essentially a bag. It is used to hold values in a key value format. 

There is currently just one implementation of this right now and it is a `ValidatorErrorBag`. This can be attached to a view (json response) in other to provide feedback as per the user's entry.

- `Hasher`

Provides a simple wrapper for `crypto/bcrypt`.