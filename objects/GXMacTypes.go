package objects

// GXMacMulticastEntry represents an entry in PRIME multicast switching table.
type GXMacMulticastEntry struct {
	Id      int8
	Members int16
}

// GXMacDirectTable represents a direct routing table entry.
type GXMacDirectTable struct {
	SourceSId       int16
	SourceLnId      int16
	SourceLcId      int16
	DestinationSId  int16
	DestinationLnId int16
	DestinationLcId int16
	Did             []byte
}

// GXMacAvailableSwitch represents a switch discovered in the network.
type GXMacAvailableSwitch struct {
	Sna     []byte
	LsId    int16
	Level   int8
	RxLevel int8
	RxSnr   int8
}

// GXMacPhyCommunication represents PHY communication parameters.
type GXMacPhyCommunication struct {
	Eui              []byte
	TxPower          int8
	TxCoding         int8
	RxCoding         int8
	RxLvl            int8
	Snr              int8
	TxPowerModified  int8
	TxCodingModified int8
	RxCodingModified int8
}
