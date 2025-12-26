package ascii_art

import (
	"strings"
)

func Sep_Fonts(datafonts string) [][]string {
	var chars [][]string
	var char []string
	datafonts = strings.ReplaceAll(datafonts, "\r\n", "\n")

	lines := strings.Split(datafonts, "\n")
	for _, line := range lines {
		if len(char) < 8 && line != "" {
			char = append(char, line)
		} else if len(char) == 8 {
			chars = append(chars, char)
			char = nil
		}
	}
	return chars
}
