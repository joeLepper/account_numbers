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

  f, err := os.Open("numbers.txt")
  check(err)

  g, err := os.Open(name)
  check(err)

  var accountNumber []byte
  digits := parseChars(bufio.NewScanner(f), 10)
  input := parseChars(bufio.NewScanner(g), 9)
  for i := 0; i < 9; i++ {
    num := getVal(input[i], digits)
    accountNumber = strconv.AppendInt(accountNumber, int64(num), 10)
  }
  account.num = string(accountNumber)
  return account
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func checkCharEq(a, b []string) bool {
  if a == nil && b == nil {
    return true;
  }

  if a == nil || b == nil {
    return false;
  }

  if len(a) != len(b) {
    return false
  }

  for i := range a {
    if a[i] != b[i] {
      return false
    }
  }

  return true
}

func getVal(char []string, digits [][]string) int {
  for i := 0; i < 10; i++ {
    isEq := checkCharEq(char, digits[i])
    if isEq {
      return i
    }
  }
  return -1
}

func parseChars(s *bufio.Scanner, length int) [][]string {
  var chars [][]string

  for j:= 0; s.Scan(); j++ {
    text := s.Text()
    for i:= 0; i < length; i++ {
      var start int = i * 3
      var end int = i * 3 + 3

      arr := strings.Split(text, "")
      if j > 0 {
        chars[i] = append(chars[i], arr[start:end]...)
      } else {
        var char []string
        chars = append(chars, char)
      }
    }
  }
  check(s.Err())
  return chars
}
