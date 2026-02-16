package dlms_tests

import (
	"encoding/base64"
	"testing"

	"github.com/Gurux/gxdlms-go/internal/buffer"
	"github.com/Gurux/gxdlms-go/types"
)

// PKCS#8 test
func Test_Pkcs8T(t *testing.T) {
	expected :=
		"MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQg4422s5KWA4fE8ekXWpck4YcryTMm6otXVGoVleIpUxKhRANCAAQl2Jb5nKsNAfJQ8NgaRTYzdEsHoXXqUUxElXgrAjbD+Tez5wD7OvoT/rKBdQIQM5onCfe5/1X6udp34QDSnn+9"
	data := "-----BEGIN PRIVATE KEY-----\n" + expected + "\n-----END PRIVATE KEY-----\n"
	cert, err := types.Pkcs8FromPem(data)
	if err != nil {
		t.Fatalf("Error parse pkcs8: %v", err)
	}
	der := cert.ToDer()
	e, err := base64.StdEncoding.DecodeString(expected)
	if err != nil {
		t.Fatalf("Error decode base64: %v", err)
	}
	a, err := base64.StdEncoding.DecodeString(der)
	if err != nil {
		t.Fatalf("Error decode base64: %v", err)
	}
	if buffer.ToHex(e, false) != buffer.ToHex(a, false) {
		t.Fatalf("got %v, want %v", buffer.ToHex(a, true), buffer.ToHex(e, true))
	}
	pem, err := cert.ToPem()
	if err != nil || data != pem {
		t.Fatalf("got %v, want %v", pem, data)
	}
	cert, err = types.Pkcs8FromPem(pem)
	if err != nil || data != pem {
		t.Fatalf("got %v, want %v", pem, data)
	}
}
