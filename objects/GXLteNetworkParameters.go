package objects

// GXLteNetworkParameters holds LTE network parameter values.
type GXLteNetworkParameters struct {
	T3402        uint16
	T3412        uint16
	T3412ext2    uint32
	T3324        uint16
	TeDRX        uint32
	TPTW         uint16
	QRxlevMin    int8
	QRxlevMinCE  int8
	QRxLevMinCE1 int8
}
