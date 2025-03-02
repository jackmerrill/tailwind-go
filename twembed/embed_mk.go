//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jackmerrill/tailwind-go/twfiles"
	"github.com/jackmerrill/tailwind-go/twpurge"
)

// read from upstream folder and build copy of tailwindcss embedded in this project

func main() {

	out, err := os.Create("embed_gen.go")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	twf := twfiles.New("../upstream/out")

	fmt.Fprintf(out, `package twembed

// WARNING: DO NOT MODIFY, THIS FILE IS GENERATED BY embed_mk.go

`)

	for _, name := range []string{"base", "components", "utilities"} {

		rc, err := twf.OpenDist(name)
		if err != nil {
			panic(err)
		}
		defer rc.Close()

		b, err := ioutil.ReadAll(rc)
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(out, "func tw%s() string {\n\treturn %q\n}\n\n", name, b)

	}

	// build the purge keys
	pkm, err := twpurge.PurgeKeysFromDist(twf)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(out, "var twPurgeKeyMap = %#v\n\n", pkm)

	// var b []byte

	// b, err = ioutil.ReadFile("../upstream/out/base.css")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Fprintf(out, "func twbase() string {\n\treturn %q\n}\n\n", b)

	// b, err = ioutil.ReadFile("../upstream/out/components.css")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Fprintf(out, "func twcomponents() string {\n\treturn %q\n}\n\n", b)

	// b, err = ioutil.ReadFile("../upstream/out/utilities.css")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Fprintf(out, "func twutilities() string {\n\treturn %q\n}\n\n", b)

}
