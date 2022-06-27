package parsinglogfiles

import (
	"fmt"
	"regexp"
)

var (
	IsValidLineRegexp          = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	SplitLogLineRegexp         = regexp.MustCompile(`<[~*=-]*>`)
	CountQuotedPasswordsRegexp = regexp.MustCompile(`"[^"]*(password)[^"]*"`)
	RemoveEndOfLineTextRegexp  = regexp.MustCompile(`end-of-line\d+`)
	TagWithUserNameRegexp      = regexp.MustCompile(`User\s+(\w+)`)
)

func IsValidLine(text string) bool {
	return IsValidLineRegexp.MatchString(text)
}

func SplitLogLine(text string) []string {
	return SplitLogLineRegexp.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	for _, v := range lines {
		count += len(CountQuotedPasswordsRegexp.FindStringSubmatch(v))
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	return RemoveEndOfLineTextRegexp.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	result := []string{}
	for _, line := range lines {
		name := TagWithUserNameRegexp.FindStringSubmatch(line)
		if name != nil {
			result = append(result, fmt.Sprintf("[USR] %s %s", name[1], line))
			continue
		}
		result = append(result, line)
	}
	return result
}
