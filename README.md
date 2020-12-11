<div align="center">
  <br/>
  <img src="https://res.cloudinary.com/stellaraf/image/upload/v1604277355/stellar-logo-gradient.svg" width=300 />
  <br/>
  <h3>Circuit ID Generator</h3>
  <br/>
  <a href="https://github.com/stellaraf/cid/actions?query=workflow%3Agoreleaser">
    <img alt="GitHub Workflow Status" src="https://img.shields.io/github/workflow/status/stellaraf/cid/goreleaser?color=9100fa&style=for-the-badge">
  </a>
  <br/>
  This repository contains source code for Stellar's command-line application to generate a circuit ID.
</div>

## Usage

### Download the latest [release](https://github.com/stellaraf/cid/releases/latest)

There are multiple builds of the release, for different CPU architectures/platforms:

There are multiple builds of the release, for different CPU architectures/platforms. Download and unpack the release for your platform:

```shell
wget <release url>
tar xvfz <release file> cid
```

### Run the binary

```console
$ ./cid --help

Circuit ID Generator

Options:

  -h, --help          Show this Help Menu
  -t, --type         *Circuit Type
  -c, --country      *Circuit Country
  -s, --state        *Circuit US State
  -i, --customer-id  *Customer ID
```

If no arguments are passed, you'll be prompted for them:

```console
$ ./cid
Circuit Type (IP Transit, Private Extension, SD-WAN, Cross Connect): transit
Country: us
US State: az
Customer ID: 123456

Country: United States of America
US State: Arizona
Type: IP Transit
Customer ID: 123456

Circuit ID: 1.84047.123456.8780
```

## Format

The format is specific to Stellar's standards:

Circuit IDs are composed of the following variables:
- Service Type (IP Transit, MPLS, SD-WAN)
- Location Identifier
  - Combination of country (USA, in my case)
  - US Region ID (From [list](https://www.census.gov/geo/reference/gtc/gtc_census_divreg.html) from the US Census Bureau)
  - US Division ID (From the same list)
  - US State FIPS Code (From the same list)
- Customer ID
  - Should be something reusable that maps to the customer across all systems. In our case, a customer CRM ID
- Service ID
  - Randomly generated 4 digit number

## Creating a New Release

This project uses [GoReleaser](https://goreleaser.com/) to manage releases. After completing code changes and committing them via Git, be sure to tag the release before pushing:

```
git tag <release>
```

Once a new tag is pushed, GoReleaser will automagically create a new build & release.
