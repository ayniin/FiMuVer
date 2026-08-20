package middleware

import "testing"

func TestIsOriginAllowed(t *testing.T) {
	// Dev-Default (keine Whitelist): jeder localhost-Port
	if !isOriginAllowed("http://localhost:5175", nil) {
		t.Error("localhost:5175 sollte im Dev-Default erlaubt sein")
	}
	if isOriginAllowed("http://evil.com", nil) {
		t.Error("fremde Origin darf nicht erlaubt sein")
	}
	// Prod (CORS_ALLOWED_ORIGINS gesetzt): strikte Whitelist
	prod := []string{"https://fimuver.de"}
	if isOriginAllowed("http://localhost:5175", prod) {
		t.Error("Whitelist muss strikt sein wenn gesetzt")
	}
	if !isOriginAllowed("https://fimuver.de", prod) {
		t.Error("Whitelist-Origin muss durchkommen")
	}
}
