package utils

import (
	"flag"
	"os"
)

func Flags() {

	// Set New Custom Flag
	creditcardFlag := flag.NewFlagSet("command", flag.ExitOnError)

	// Init custom flag arguments
	creditcardFlag.BoolVar(&Pick, "pick", false, "Pick")
	creditcardFlag.BoolVar(&Help, "help", false, "Help")
	creditcardFlag.BoolVar(&Stdin, "stdin", false, "Stdin")
	creditcardFlag.StringVar(&Brands, "brands", "data/brands.txt", "Brands")
	creditcardFlag.StringVar(&Issuers, "issuers", "data/issuers.txt", "Issuers")
	creditcardFlag.StringVar(&Brand, "brand", "", "Brand")
	creditcardFlag.StringVar(&Issuer, "issuer", "", "Issuer")

	OsArgs = os.Args

	// Parse command line
	creditcardFlag.Parse(os.Args[2:])
}
