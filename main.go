package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
)

// This will just be $FUZZ like other tools by default.
const DEF_TOKEN = "$FUZZ"

func init() {
	// Display usage information
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: czuffr <FLAGS> -- <TARGET_BINARY> %s\n", DEF_TOKEN)
		fmt.Println("\nOptions:")

		flag.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(os.Stderr, "  --%s=%s: %s\n", f.Name, f.DefValue, f.Usage)
		})

		fmt.Println("\nExamples:")
		fmt.Println("\tAdd examples here.")
	}
}

// PrintBanner to console at start of program.
func PrintBanner() {
	fmt.Println("[ cZuffr - Command Line Fuzzer ]\n")
}

func main() {

	// Setup signal handler
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	PrintBanner()

	// Parse command line options
	var (
		flagMutFuzz  = flag.Bool("mutfuzz", true, "Generate any mutation payload to fuzz with.")
		flagCrash    = flag.Bool("crash", true, "If fuzz case causes crash, save payload.")
		flagExe      = flag.Bool("exec", true, "Execute the command with the fuzzed payload.")
		flagScript   = flag.Bool("script", false, "Execute the script with the fuzzed payload.")
		flagWordlist = flag.String("wordlist", "", "Wordlist bruteforce argument fuzzing.")
		flagNeedle   = flag.String("needle", "", "If this option is set, the string ")
		flagToken    = flag.String("token", "$FUZZ", "The placeholder token that is replaced with fuzzing/wordlist values.")
	)
	flag.Parse()

	_ = flagToken
	_ = flagNeedle
	_ = flagWordlist
	_ = flagScript
	_ = flagExe
	_ = flagCrash
	_ = flagMutFuzz

	fmt.Println("Lets FUZZZ")

}
