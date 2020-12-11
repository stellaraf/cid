package main

import (
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mkideal/cli"
	"github.com/stellaraf/cid/data"
)

type args struct {
	Help       bool   `cli:"!h,help" usage:"Show this Help Menu" json:"-"`
	Type       string `cli:"*t,type" usage:"Circuit Type" prompt:"Circuit Type (IP Transit, Private Extension, SD-WAN, Cross Connect)"`
	Country    string `cli:"*c,country" usage:"Circuit Country" prompt:"Country"`
	State      string `cli:"*s,state" usage:"Circuit US State" prompt:"US State"`
	CustomerID string `cli:"*i,customer-id" usage:"Customer ID" prompt:"Customer ID"`
}

func (argv *args) AutoHelp() bool {
	return argv.Help
}

func getRandom(low, hi int) int {
	rand.Seed(time.Now().UnixNano())
	return low + rand.Intn(hi-low)
}

func buildCID(c data.Country, s data.USState, t data.Circuit, i string, r int) (cID string) {
	locID := c.ISO + s.DivisionID + s.RegionID
	cID = strings.Join([]string{t.ID, locID, i, strconv.Itoa(r)}, ".")
	return
}

func main() {

	os.Exit(cli.Run(new(args), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*args)

		country, err := data.GetCountry(argv.Country)
		if err != nil {
			return err
		}
		usstate, err := data.GetState(argv.State)
		if err != nil {
			return err
		}
		circuitType, err := data.GetCircuitType(argv.Type)
		if err != nil {
			return err
		}

		custID := argv.CustomerID

		randID := getRandom(1000, 9999)
		cid := buildCID(country, usstate, circuitType, custID, randID)
		c := ctx.Color()
		ctx.String("\n")

		ctx.String(c.White("Country: ") + c.Bold((c.Red(country.FullName) + "\n")))
		ctx.String(c.White("US State: ") + c.Bold((c.Blue(usstate.FullName) + "\n")))
		ctx.String(c.White("Type: ") + c.Bold((c.Magenta(circuitType.Name) + "\n")))
		ctx.String(c.White("Customer ID: ") + c.Bold((c.Yellow(custID) + "\n")))
		ctx.String(c.Bold(c.White("\nCircuit ID: ")) + c.Bold((c.Green(cid) + "\n")))

		return nil
	}, "\nCircuit ID Generator"))
}
