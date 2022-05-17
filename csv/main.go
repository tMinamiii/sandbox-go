package main

import (
	"bytes"
	"encoding/csv"
	"fmt"

	"github.com/jszwec/csvutil"
)

type migrationUser struct {
	FromUserID int64 `csv:"from"`
	ToUserID   int64 `csv:"to"`
}
type migrationUsers []migrationUser

func main() {
	emptyCSV := `from, to`
	csvReader := csv.NewReader(bytes.NewReader([]byte(emptyCSV)))
	decoder, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		fmt.Println(err)
	}

	var users migrationUsers
	if err = decoder.Decode(&users); err != nil {
		if err.Error() == "EOF" {
			users = migrationUsers{}
		} else {
			fmt.Println(err)
		}
	}
	fmt.Println(users)
}
