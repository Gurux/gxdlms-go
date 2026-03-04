package objects

// GXDLMSImageActivateInfo holds image metadata used by GXDLMSImageTransfer.
type GXDLMSImageActivateInfo struct {
	Size           uint32
	Identification []byte
	Signature      []byte
}
