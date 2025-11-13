package rows

//import "strings"

type Row []string

// New erwartet eine Länge und einen String.
// Gibt eine neue `Row` zurück, die mit dem String gefüllt ist.
func New(length int, fill string) Row {
	//return strings.Split(strings.Repeat(fill, length), "")

	var result Row
	for i := 0; i < length; i++ {

		result = append(result, fill)
	}

	return result

}

// String gibt die `Row` als String zurück.
// Die Elemente sind durch ` | ` getrennt und von `|` umgeben.

func (r Row) String() string {

	if len(r) == 0 {
		return "| |"

	}

	result := "| " + r[0]

	for i := 1; i < len(r); i++ {

		result = result + " | " + r[i]

	}
	result = result + " |"
	return result

}

// ContainsOnly erwartet einen String.
// Gibt `true` zurück, wenn die `Row` nur aus dem String besteht.
func (r Row) ContainsOnly(s string) bool {

	for i := 0; i < len(r); i++ {
		if r[i] != s {
			return false
		}
	}
	return true

}
