#Account Numbers

A response to [this](http://codingdojo.org/cgi-bin/index.pl?KataBankOCR) in Go.

## Getting started

1. Make sure that you have Go installed and are working within your `$GOPATH`.
1. `cd` into the project `src` root.
1. `go test`

I've left a failing test to get started.

## Next steps

1. Fix the failing test case.
1. Make it more performant.
1. Unit test the helper functions? This seems to be not very Go-y.
1. `parseChars` should be able to tell how many digits are in the string it accepts instead of being told explicitly.
1. Make the binary accept piped `stdin`.
1. Get better at type-casting in Go. `Account.number` probably shouldn't be a string.
