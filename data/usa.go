package data

import (
	"fmt"
	"strings"

	gjson "github.com/tidwall/gjson"
)

// USState represents a US State, duh.
type USState struct {
	FullName   string `json:"fullName"`
	ShortName  string `json:"shortName"`
	RegionID   string `json:"regionID"`
	DivisionID string `json:"divisionID"`
	FIPSCode   string `json:"fipsCode"`
}

// USStateJSON is a JSON representation of all US States & metadata.
const USStateJSON = `
[
	{"fullName": "Alabama", "shortName": "AL", "regionID": "6", "divisionID": "3", "fipsCode": "01"},
	{"fullName": "Alaska", "shortName": "AK", "regionID": "8", "divisionID": "4", "fipsCode": "02"},
	{"fullName": "Arizona", "shortName": "AZ", "regionID": "7", "divisionID": "4", "fipsCode": "04"},
	{"fullName": "Arkansas", "shortName": "AR", "regionID": "6", "divisionID": "3", "fipsCode": "05"},
	{"fullName": "California", "shortName": "CA", "regionID": "8", "divisionID": "4", "fipsCode": "06"},
	{"fullName": "Colorado", "shortName": "CO", "regionID": "7", "divisionID": "4", "fipsCode": "08"},
	{"fullName": "Connecticut", "shortName": "CT", "regionID": "1", "divisionID": "1", "fipsCode": "09"},
	{"fullName": "Delaware", "shortName": "DE", "regionID": "5", "divisionID": "3", "fipsCode": "10"},
	{"fullName": "District of Columbia", "shortName": "DC", "regionID": "5", "divisionID": "3", "fipsCode": "11"},
	{"fullName": "Florida", "shortName": "FL", "regionID": "5", "divisionID": "3", "fipsCode": "12"},
	{"fullName": "Georgia", "shortName": "GA", "regionID": "5", "divisionID": "3", "fipsCode": "13"},
	{"fullName": "Hawaii", "shortName": "HI", "regionID": "8", "divisionID": "4", "fipsCode": "15"},
	{"fullName": "Idaho", "shortName": "ID", "regionID": "7", "divisionID": "4", "fipsCode": "16"},
	{"fullName": "Illinois", "shortName": "IL", "regionID": "3", "divisionID": "2", "fipsCode": "17"},
	{"fullName": "Indiana", "shortName": "IN", "regionID": "3", "divisionID": "2", "fipsCode": "18"},
	{"fullName": "Iowa", "shortName": "IA", "regionID": "4", "divisionID": "2", "fipsCode": "19"},
	{"fullName": "Kansas", "shortName": "KS", "regionID": "4", "divisionID": "2", "fipsCode": "20"},
	{"fullName": "Kentucky", "shortName": "KY", "regionID": "6", "divisionID": "3", "fipsCode": "21"},
	{"fullName": "Louisiana", "shortName": "LA", "regionID": "6", "divisionID": "3", "fipsCode": "22"},
	{"fullName": "Maine", "shortName": "ME", "regionID": "1", "divisionID": "1", "fipsCode": "23"},
	{"fullName": "Maryland", "shortName": "MD", "regionID": "5", "divisionID": "3", "fipsCode": "24"},
	{"fullName": "Massachusetts", "shortName": "MA", "regionID": "1", "divisionID": "1", "fipsCode": "25"},
	{"fullName": "Michigan", "shortName": "MI", "regionID": "3", "divisionID": "2", "fipsCode": "26"},
	{"fullName": "Minnesota", "shortName": "MN", "regionID": "4", "divisionID": "2", "fipsCode": "27"},
	{"fullName": "Mississippi", "shortName": "MS", "regionID": "6", "divisionID": "3", "fipsCode": "28"},
	{"fullName": "Missouri", "shortName": "MO", "regionID": "4", "divisionID": "2", "fipsCode": "29"},
	{"fullName": "Montana", "shortName": "MT", "regionID": "7", "divisionID": "4", "fipsCode": "30"},
	{"fullName": "Nebraska", "shortName": "NE", "regionID": "4", "divisionID": "2", "fipsCode": "31"},
	{"fullName": "Nevada", "shortName": "NV", "regionID": "7", "divisionID": "4", "fipsCode": "32"},
	{"fullName": "New Hampshire", "shortName": "NH", "regionID": "1", "divisionID": "1", "fipsCode": "33"},
	{"fullName": "New Jersey", "shortName": "NJ", "regionID": "2", "divisionID": "1", "fipsCode": "34"},
	{"fullName": "New Mexico", "shortName": "NM", "regionID": "7", "divisionID": "4", "fipsCode": "35"},
	{"fullName": "New York", "shortName": "NY", "regionID": "2", "divisionID": "1", "fipsCode": "36"},
	{"fullName": "North Carolina", "shortName": "NC", "regionID": "5", "divisionID": "3", "fipsCode": "37"},
	{"fullName": "North Dakota", "shortName": "ND", "regionID": "4", "divisionID": "2", "fipsCode": "38"},
	{"fullName": "Ohio", "shortName": "OH", "regionID": "3", "divisionID": "2", "fipsCode": "39"},
	{"fullName": "Oklahoma", "shortName": "OK", "regionID": "6", "divisionID": "3", "fipsCode": "40"},
	{"fullName": "Oregon", "shortName": "OR", "regionID": "8", "divisionID": "4", "fipsCode": "41"},
	{"fullName": "Pennsylvania", "shortName": "PA", "regionID": "2", "divisionID": "1", "fipsCode": "42"},
	{"fullName": "Rhode Island", "shortName": "RI", "regionID": "1", "divisionID": "1", "fipsCode": "44"},
	{"fullName": "South Carolina", "shortName": "SC", "regionID": "5", "divisionID": "3", "fipsCode": "45"},
	{"fullName": "South Dakota", "shortName": "SD", "regionID": "4", "divisionID": "2", "fipsCode": "46"},
	{"fullName": "Tennessee", "shortName": "TN", "regionID": "6", "divisionID": "3", "fipsCode": "47"},
	{"fullName": "Texas", "shortName": "TX", "regionID": "6", "divisionID": "3", "fipsCode": "48"},
	{"fullName": "Utah", "shortName": "UT", "regionID": "7", "divisionID": "4", "fipsCode": "49"},
	{"fullName": "Vermont", "shortName": "VT", "regionID": "1", "divisionID": "1", "fipsCode": "50"},
	{"fullName": "Virginia", "shortName": "VA", "regionID": "5", "divisionID": "3", "fipsCode": "51"},
	{"fullName": "Washington", "shortName": "WA", "regionID": "8", "divisionID": "4", "fipsCode": "53"},
	{"fullName": "West Virginia", "shortName": "WV", "regionID": "5", "divisionID": "3", "fipsCode": "54"},
	{"fullName": "Wisconsin", "shortName": "WI", "regionID": "3", "divisionID": "2", "fipsCode": "55"},
	{"fullName": "Wyoming", "shortName": "WY", "regionID": "7", "divisionID": "4", "fipsCode": "56"}
]
`

// GetState gets a state, duh.
func GetState(search string) (r USState, e error) {
	gjson.AddModifier("case", func(json, arg string) string {
		if arg == "upper" {
			return strings.ToUpper(json)
		}
		if arg == "lower" {
			return strings.ToLower(json)
		}
		return json
	})
	states := gjson.Parse(USStateJSON).Array()
	sl := strings.ToLower(search)
	e = nil
	for _, s := range states {
		if sl == s.Get("fullName|@case:lower").Str || sl == s.Get("shortName|@case:lower").Str {
			r = USState{
				FullName:   s.Get("fullName").Str,
				ShortName:  s.Get("shortName").Str,
				RegionID:   s.Get("regionID").Str,
				DivisionID: s.Get("divisionID").Str,
				FIPSCode:   s.Get("fipsCode").Str,
			}
			break
		}
	}
	if (USState{} == r) {
		e = fmt.Errorf("No State Matches '%s'", search)
	}
	return
}
