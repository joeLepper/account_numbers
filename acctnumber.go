package acctnumber

import  (
   "os"
   "strings"
   "bufio"
   "strconv"
)

type Account struct {
  num string
  filename string
}

func NewAccount(name string) *Account {
  account := new(Account)
  account.filename = name
  account.num = ""
  return account
}
