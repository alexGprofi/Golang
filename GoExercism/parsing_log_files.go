package main

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re, _ := regexp.Compile(`(^\[TRC\]|^\[DBG\]|^\[INF\]|^\[WRN\]|^\[ERR\]|^\[FTL\])`)
	if re.MatchString(text) {
		return true
	}
	return false
}

func SplitLogLine(text string) []string {
	re, _ := regexp.Compile(`\<\W*\>`)
	sl := re.Split(text, -1)
	fmt.Println("sl: ", sl)
	return sl
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	for _, line := range lines {
		match, _ := regexp.MatchString(`(?i)\".*(password).*\"`, line)
		fmt.Println(match)
		if match == true {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line+\d*`)
	s := re.ReplaceAllString(text, "")
	return s
}

func TagWithUserName(lines []string) []string {
	s := make([]string, len(lines))
	var myString, myString2 string
	for index, line := range lines {
		re, _ := regexp.Compile(`User+\s+(\w+)`)
		if re.MatchString(line) {
			myString = re.FindString(line)
			re2, _ := regexp.Compile(`User+\s+`)
			myString2 = re2.ReplaceAllString(myString, "[USR] ")
			s[index] = myString2 + " " + line
			fmt.Println(s[index])
		} else {
			s[index] = line
		}

	}
	return s
}

func main() {
	lines := []string{
		`[INF] passWord`, // contains 'password' but not surrounded by quotation marks
		`"passWord"`,     // count this one
		`[INF] User saw error message "Unexpected Error" on page load.`,          // does not contain 'password'
		`[INF] The message "Please reset your password" was ignored by the user`, // count this one
	}
	numLines := CountQuotedPasswords(lines)
	fmt.Println("numLines: ", numLines)

	sEOL := RemoveEndOfLineText("[INF] end-of-line23033 Network Failure end-of-line27")
	fmt.Println("sEOL: ", sEOL)
	// => "[INF]  Network Failure "

	result := TagWithUserName([]string{
		"[WRN] User James123 has exceeded storage space.",
		"[WRN] Host down. User   Michelle4 lost connection.",
		"[INF] Users can login again after 23:00.",
		"[DBG] We need to check that user names are at least 6 chars long.",
	})
	fmt.Println("result: ", result)
}
