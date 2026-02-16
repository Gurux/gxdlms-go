package dlms_tests

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/Gurux/gxdlms-go/objects"
)

func Test_Serialize(t *testing.T) {
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

func Test_FileSerialize(t *testing.T) {
	var _ objects.IGXDLMSBase = (*objects.GXDLMSData)(nil)

	d, err := objects.NewGXDLMSData("0.0.42.0.0.255", 0)
	if err != nil {
		t.Fatalf("Failed to create GXDLMSData: %v", err)
	}
	d.Value = int32(123456)
	objs := objects.GXDLMSObjectCollection{}
	objs = append(objs, d)
	err = objs.SaveToFile("testfile.xml", nil)
	if err != nil {
		t.Fatalf("Failed to save GXDLMSData: %v", err)
	}
	objs = objects.GXDLMSObjectCollection{}
	err = objs.LoadFromFile("testfile.xml")
	if err != nil {
		t.Fatalf("Failed to load GXDLMSData: %v", err)
	}
}
