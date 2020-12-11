package data

import (
	"fmt"
	"strings"

	gjson "github.com/tidwall/gjson"
)

// Circuit represents a circuit, duh.
type Circuit struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

const circuitsJSON = `
[
	{"id": "1", "name": "IP Transit"},
	{"id": "2", "name": "Private Extension"},
	{"id": "3", "name": "SD-WAN"},
	{"id": "4", "name": "Cross Connect"}
]
`

// GetCircuitType gets a circuit type, duh.
func GetCircuitType(search string) (r Circuit, e error) {
	gjson.AddModifier("case", func(json, arg string) string {
		if arg == "upper" {
			return strings.ToUpper(json)
		}
		if arg == "lower" {
			return strings.ToLower(json)
		}
		return json
	})
	circuits := gjson.Parse(circuitsJSON).Array()
	sl := strings.ToLower(search)
	e = nil
	for _, s := range circuits {
		if sl == s.Get("id").Str || strings.Contains(s.Get("name|@case:lower").Str, sl) {
			r = Circuit{
				ID:   s.Get("id").Str,
				Name: s.Get("name").Str,
			}
			break
		}
	}
	if (Circuit{} == r) {
		e = fmt.Errorf("No Circuit Type Matches '%s'", search)
	}
	return
}
