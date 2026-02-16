package dlms_tests

import (
	"bufio"
	"bytes"
	"slices"
	"testing"

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/objects"
)

// Test_RegisterAttributeCount tests the GetAttributeCount method of GXDLMSRegister.
func Test_RegisterAttributeCount(t *testing.T) {
	d := objects.GXDLMSRegister{}
	expected := 3
	actual := d.GetAttributeCount()
	if actual != expected {
		t.Fatalf("got %v, want %v", actual, expected)
	}
}

// Test_RegisterMethodCount tests the GetMethodCount method of GXDLMSRegister.
func Test_RegisterMethodCount(t *testing.T) {
	d := objects.GXDLMSRegister{}
	expected := 1
	actual := d.GetMethodCount()
	if actual != expected {
		t.Fatalf("got %v, want %v", actual, expected)
	}
}

// Test_RegisterNames tests the GetNames method of GXDLMSRegister.
func Test_RegisterNames(t *testing.T) {
	d := objects.GXDLMSRegister{}
	expected := []string{"Logical Name", "Value", "Scaler and Unit"}
	actual := d.GetNames()
	if !slices.Equal(actual, expected) {
		t.Fatalf("got %v, want %v", actual, expected)
	}
}

// Test_RegisterRegisterTypes tests the GetRegisterType method of GXDLMSRegister.
func Test_RegisterTypes(t *testing.T) {
	d := objects.GXDLMSRegister{}
	expected := enums.DataType(enums.DataTypeOctetString)
	actual, err := d.GetDataType(1)
	if err != nil || actual != expected {
		t.Fatalf("got %v, want %v", actual, expected)
	}
}

func Test_RegisterObjectSerialize(t *testing.T) {
	var _ objects.IGXDLMSBase = (*objects.GXDLMSRegister)(nil)

	d, err := objects.NewGXDLMSRegister("0.0.42.0.0.255", 0)
	if err != nil {
		t.Fatalf("Failed to create GXDLMSRegister: %v", err)
	}
	d.Value = int32(123456)
	objs := objects.GXDLMSObjectCollection{}
	objs = append(objs, d)
	var ms bytes.Buffer
	bw := bufio.NewWriter(&ms)
	err = objs.SaveToStream(bw, nil)
	if err != nil {
		t.Fatalf("Failed to save GXDLMSRegister: %v", err)
	}
	bw.Flush()
	objs = objects.GXDLMSObjectCollection{}
	br := bufio.NewReader(&ms)
	err = objs.LoadFromStream(br)
	if err != nil {
		t.Fatalf("Failed to load GXDLMSRegister: %v", err)
	}
}
