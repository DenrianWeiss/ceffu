# Ceffu Client for Go

This is a trivial client for the Ceffu API written in Go. Currently, this package is in experimental stage.

## Usage

```go
package main

import "github.com/DenrianWeiss/ceffu"
import "net/http"

const API_KEY = ""
const PRIVATE_KEY = ""

func main() {
	cl, err := ceffu.New(API_KEY, PRIVATE_KEY, http.DefaultClient, nil, ceffu.CeffuApiBaseUrl)
	if err != nil {
		panic(err)
	}
	cl.GetStatus(ceffu.BusinessTypeDeposit, "10")
}

```