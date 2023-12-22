package main

import(
    "regexp"
)

//Ansi escape code for color
const (
	Reset  = "\033[0m"
	Blue   = "\033[34;1m"
	Red    = "\033[31m"
	Green  = "\033[3221m"
	White  = "\033[37;1m"
	Yellow = "\033[33;1m"
	Purple = "\033[35m"
    Cyan   = "\033[36;1m"
	Gray   = "\033[90m"

    // 256 color pallete
    Yellow256 = "\033[;38;5;214m"
    Blue256   = "\033[;38;5;26m"
	Cyan256   = "\033[;38;5;51m"
    Orange256 = "\033[;38;5;208m"
)


func RemoveANSI(str string) string {
	// Define a regular expression to match ANSI escape codes
	ansiRegex := regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)

	// Use the regular expression to replace ANSI escape codes with an empty string
	cleanStr := ansiRegex.ReplaceAllString(str, "")

	return cleanStr
}
