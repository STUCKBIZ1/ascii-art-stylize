package ascii_art_web

func IsSupportedText(text string) bool {
	for _, v := range text {
		if v >= 32 && v <= 126 {
			return true
		}
	}
	return false
}
