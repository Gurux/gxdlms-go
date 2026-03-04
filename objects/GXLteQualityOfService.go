package objects

import "github.com/Gurux/gxdlms-go/enums"

// GXLteQualityOfService holds LTE quality-of-service values.
type GXLteQualityOfService struct {
	SignalQuality       int8
	SignalLevel         int8
	SignalToNoiseRatio  int8
	CoverageEnhancement enums.LteCoverageEnhancement
}
