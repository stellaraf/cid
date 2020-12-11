package data

import (
	"fmt"
	"strings"

	gjson "github.com/tidwall/gjson"
)

// Country represents a country, duh.
type Country struct {
	FullName  string `json:"fullName"`
	ShortName string `json:"shortName"`
	ISO       string `json:"iso"`
}

const countriesJSON = `[{"fullName": "United States of America", "shortName": "US", "iso": "840"}]`

// GetCountry gets a country, duh.
func GetCountry(search string) (r Country, e error) {
	gjson.AddModifier("case", func(json, arg string) string {
		if arg == "upper" {
			return strings.ToUpper(json)
		}
		if arg == "lower" {
			return strings.ToLower(json)
		}
		return json
	})
	countries := gjson.Parse(countriesJSON).Array()
	sl := strings.ToLower(search)
	e = nil
	for _, s := range countries {
		if sl == s.Get("fullName|@case:lower").Str || sl == s.Get("shortName|@case:lower").Str {
			r = Country{
				FullName:  s.Get("fullName").Str,
				ShortName: s.Get("shortName").Str,
				ISO:       s.Get("iso").Str,
			}
			break
		}
	}
	if (Country{} == r) {
		e = fmt.Errorf("No Country Matches '%s'", search)
	}
	return
}
