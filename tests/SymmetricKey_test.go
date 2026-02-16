package dlms_tests

import (
	"testing"

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

func Test_Authentication(t *testing.T) {
	expected := "06 72 5D 91 0F 92 21 D2 63 87 75 16"
	data := "C0 01 00 00 08 00 00 01 00 00 FF 02 00"
	p := settings.NewAesGcmParameter(0x10, nil, enums.SecurityAuthentication,
		enums.SecuritySuite0, 0x01234567,
		types.HexToBytes("4D4D4D0000BC614E"), types.HexToBytes("000102030405060708090A0B0C0D0E0F"),
		types.HexToBytes("D0D1D2D3D4D5D6D7D8D9DADBDCDDDEDF"),
	)
	p.Type = settings.CountTypeTag
	d := types.HexToBytes(data)
	ret, err := settings.EncryptAesGcm(p, d)
	if err != nil {
		t.Errorf("DateTimeFromString failed. Expected no error, but got: %s", err.Error())
	}
	actual := types.ToHex(ret, true)
	if actual != expected {
		t.Errorf("DateTimeFromString failed. Expected: %s, Actual: %s", expected, actual)
	}
	bb := types.GXByteBuffer{}
	bb.SetHexString(data)
	bb.SetHexString(expected)
	ret, err = settings.DecryptAesGcm(p, &bb)
	if err != nil {
		t.Errorf("DateTimeFromString failed. Expected no error, but got: %s", err.Error())
	}
	actual = types.ToHex(ret, true)
	if actual != data {
		t.Errorf("DateTimeFromString failed. Expected: %s, Actual: %s", data, actual)
	}
}
