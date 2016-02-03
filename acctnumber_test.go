package acctnumber

import (
	"io/ioutil"
	"testing"
)

func TestAccount(t *testing.T) {
	cases, err := ioutil.ReadDir("./cases")
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(cases); i++ {

		name := cases[i].Name()
		if name != ".DS_Store" && name != ".git" {
			acct := NewAccount("./cases/" + name)
			if acct.num != name {
				t.Log(name)
				t.Fail()
			}
		}
	}
}
