package dlms_tests

import (
	"testing"

	"github.com/Gurux/gxdlms-go/types"
)

// Test_SetUint8AndUint8 tests setting and getting a uint8 value in GXByteBuffer.
func Test_SetUint8AndUint8(t *testing.T) {
	bb := types.GXByteBuffer{}
	bb.SetUint8(5)
	val, err := bb.Uint8()
	if err != nil || val != 5 {
		t.Error("Expected 5")
	}
}

func Test_Hex(t *testing.T) {
	expected := "7E A0 07 03 21 93 0F 01 7E"
	bb := types.GXByteBuffer{}
	bb.SetHexString("7E-A007:03 21 93 0F 01 7E")
	actual := bb.String()
	if actual != expected {
		t.Error("Expected " + expected + ", got " + actual)
	}
}
