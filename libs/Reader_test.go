package libs

import (
	"testing"
)

func TestCSVInput_Parser(t *testing.T) {
	csv_formatter,err := CreateFormatter(map[string]string{
		"FORMAT":"CSV",
		"LOCATION":"D:\\clientsFA.csv",
		"SEPARATOR":",",
	})

	checkErr(err)
	_,err = csv_formatter.Parser(csv_formatter.Read())
	if err != nil {
		t.Error(err)
	}
}
