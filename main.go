package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var fromBase int
var toBase int
var pretty bool

func main(){
	flag.IntVar(&fromBase, "f", 16, "Base to convert from")
	flag.IntVar(&toBase, "t", 10, "Base to convert to")
	flag.BoolVar(&pretty, "p", false, "Add prefixes to bases where possible (e.g. 0x for hex)")

	flag.Parse()

	if toBase < 2 || toBase > 36 || fromBase < 2 || fromBase > 36 {
		fmt.Println("Base must be 2<=base<=36")
		os.Exit(1)
	}

	conv := flag.Args()
	if len(conv) < 1 {
		fmt.Println("No number provided")
		os.Exit(1)
	}

	// Trim prefixes if they were left on
	if fromBase == 16 && conv[0][:2] == "0x" {
		conv[0] = conv[0][2:]
	}
	if fromBase == 2 && conv[0][:2] == "0b" {
		conv[0] = conv[0][2:]
	}

	// Get decimal repr of input
	raw, err := strconv.ParseInt(conv[0], fromBase, 64)
	if err != nil {
		// this could either be the number being too big or
		// it not being a valid number at all
		// i dont care enough to make it look pretty
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}

	// Convert
	res := strconv.FormatInt(raw, toBase)

	// Add prefix
	if (pretty){
		switch toBase {
			case 2: res = "0b"+res
			case 16: res = "0x"+res
		}
	}

	// boop c:
	fmt.Printf("%s\n", res)
}
