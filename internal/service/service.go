package service

import (
	"log"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// TextDifinition function determines which string is passed and, depending on what is passed,
// converts the string to text or Morse code
func TextDifinition(s string) string {
	if strings.Contains(s, "--") || strings.Contains(s, ".-") || strings.Contains(s, "-.") {
		return morse.ToText(s)
	}
	if len(s) == 0 {
		log.Fatal("an empty string was passed")
	}
	return morse.ToMorse(s)
}
