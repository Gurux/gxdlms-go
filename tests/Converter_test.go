package dlms_tests

import (
	"testing"

	dlms "github.com/Gurux/gxdlms-go"
	"github.com/Gurux/gxdlms-go/enums"
)

func Test_Converter(t *testing.T) {
	c := dlms.GXDLMSConverter{}
	expected := "Ch. 0 Clock object  #1"
	actual, err := c.GetDescription("0.0.1.0.0.255", enums.ObjectTypeNone, "")
	if err != nil || actual[0] != expected {
		t.Fatalf("got %v, want %v", actual, expected)
	}
}
