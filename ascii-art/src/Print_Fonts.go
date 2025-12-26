package ascii_art

func Print_Fonts(OutputFonts [][]string) string {
	var result string
	if len(OutputFonts) == 0 {
		return ""
	}

	for i := 0; i < 8; i++ {
		for _, v := range OutputFonts {
			result += v[i]
		}
		result += "\n"
	}
	return result
}
