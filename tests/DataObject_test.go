package dlms_tests

import (
	"bufio"
	"bytes"
	"slices"
	"testing"

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/objects"
)

// Test_DataAttributeCount tests the GetAttributeCount method of GXDLMSData.
func Test_DataAttributeCount(t *testing.T) {
	d := objects.GXDLMSData{}
	expected := 2
	actual := d.GetAttributeCount()
	if actual != expected {
		t.Fatalf("got %v, want %v", actual, expected)
	}
}

// Test_DataMethodCount tests the GetMethodCount method of GXDLMSData.
func Test_DataMethodCount(t *testing.T) {
	d := objects.GXDLMSData{}
	expected := 0
	actual := d.GetMethodCount()
	if actual != expected {
		t.Fatalf("got %v, want %v", actual, expected)
	}
}

// Test_DataNames tests the GetNames method of GXDLMSData.
func Test_DataNames(t *testing.T) {
	d := objects.GXDLMSData{}
	expected := []string{"Logical Name", "Value"}
	actual := d.GetNames()
	if !slices.Equal(actual, expected) {
		t.Fatalf("got %v, want %v", actual, expected)
	}
}

// Test_DataDataTypes tests the GetDataType method of GXDLMSData.
func Test_DataTypes(t *testing.T) {
	d := objects.GXDLMSData{}
	expected := enums.DataType(enums.DataTypeOctetString)
	actual, err := d.GetDataType(1)
	if err != nil || actual != expected {
		t.Fatalf("got %v, want %v", actual, expected)
	}
}

func Test_DataObjectSerialize(t *testing.T) {
	var _ objects.IGXDLMSBase = (*objects.GXDLMSData)(nil)

	d, err := objects.NewGXDLMSData("0.0.42.0.0.255", 0)
	if err != nil {
		t.Fatalf("Failed to create GXDLMSData: %v", err)
	}
	d.Value = int32(123456)
	objs := objects.GXDLMSObjectCollection{}
	objs = append(objs, d)
	var ms bytes.Buffer
	bw := bufio.NewWriter(&ms)
	err = objs.SaveToStream(bw, nil)
	if err != nil {
		t.Fatalf("Failed to save GXDLMSData: %v", err)
	}
	bw.Flush()
	objs = objects.GXDLMSObjectCollection{}
	br := bufio.NewReader(&ms)
	err = objs.LoadFromStream(br)
	if err != nil {
		t.Fatalf("Failed to load GXDLMSData: %v", err)
	}
}
