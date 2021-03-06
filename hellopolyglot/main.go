// Copyright 2012 The polyglot Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
)

import (
	"github.com/lxn/polyglot"
)

var trDict *polyglot.Dict

// An implemententation of a "tr" function like this is required for every 
// package that you want to work with polyglot.
func tr(source string, context ...string) string {
	return trDict.Translation(source, context...)
}

// The example translation file was generated by executing
// cd l10n
// polyglot -dir=".." -locales="de" -name="hellopolyglot"
func main() {
	var err error
	if trDict, err = polyglot.NewDict("l10n", "de"); err != nil {
		log.Fatal(err)
	}

	// The polyglot command scans Go code for calls to a "tr" function like this 
	// one. If context args are provided, they are used for disambiguation. 
	fmt.Println(tr("Hello"))
}
