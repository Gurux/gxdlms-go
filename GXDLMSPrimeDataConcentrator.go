package dlms

//
// --------------------------------------------------------------------------
//  Gurux Ltd
//
//
//
// Filename:        $HeadURL$
//
// Version:         $Revision$,
//                  $Date$
//                  $Author$
//
// Copyright (c) Gurux Ltd
//
//---------------------------------------------------------------------------
//
//  DESCRIPTION
//
// This file is a part of Gurux Device Framework.
//
// Gurux Device Framework is Open Source software; you can redistribute it
// and/or modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; version 2 of the License.
// Gurux Device Framework is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU General Public License for more details.
//
// More information of Gurux products: https://www.gurux.org
//
// This code is licensed under the GNU General Public License v2.
// Full text may be retrieved at http://www.gnu.org/licenses/gpl-2.0.txt
//---------------------------------------------------------------------------

import (
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/types"
)

// GXDLMSPrimeDataConcentrator contains information that is needed
// if PRIME data concentrator is used between the client and the meter.
type GXDLMSPrimeDataConcentrator struct {
	// Notification message type.
	Type enums.PrimeDcMsgType

	// Device identifier that will be used to address the device this
	// notification refers to, both in custom messages and in the
	// standard DLMS/TCP communications.
	DeviceID uint16

	// Flags with the capabilities of the device.
	Capabilities uint16

	// DLMS Identifier of the reported device.
	DlmsId []byte

	// EUI48 of the reported device.
	Eui48 []byte
}

func (g *GXDLMSPrimeDataConcentrator) String() string {
	var err error
	var ret []byte
	bb := types.GXByteBuffer{}
	switch g.Type {
	case enums.PrimeDcMsgTypeNewDeviceNotification:
		ret, err = g.GenerateNewDeviceNotification(nil)
		if err != nil {
			return err.Error()
		}
		err = bb.Set(ret)
		if err != nil {
			return err.Error()
		}
	case enums.PrimeDcMsgTypeRemoveDeviceNotification:
		ret, err = g.GenerateRemoveDeviceNotification(nil)
		if err != nil {
			return err.Error()
		}
		err = bb.Set(ret)
		if err != nil {
			return err.Error()
		}
	case enums.PrimeDcMsgTypeStartReportingMeters:
		ret, err = g.GenerateStartReportingMeters(nil)
		if err != nil {
			return err.Error()
		}
		err = bb.Set(ret)
		if err != nil {
			return err.Error()
		}
	case enums.PrimeDcMsgTypeDeleteMeters:
		ret, err = g.GenerateDeleteMetersNotification(nil)
		if err != nil {
			return err.Error()
		}
		err = bb.Set(ret)
		if err != nil {
			return err.Error()
		}
	case enums.PrimeDcMsgTypeEnableAutoClose:
		ret, err = g.GenerateEnableAutoCloseNotification(nil)
		if err != nil {
			return err.Error()
		}
		err = bb.Set(ret)
		if err != nil {
			return err.Error()
		}
	case enums.PrimeDcMsgTypeDisableAutoClose:
		ret, err = g.GenerateDisableAutoCloseNotification(nil)
		if err != nil {
			return err.Error()
		}
		err = bb.Set(ret)
		if err != nil {
			return err.Error()
		}
	default:
		return "GXDLMSPrimeDataConcentrator"
	}
	t := GXDLMSTranslator{}
	return t.PduToXml(&bb, enums.InterfaceTypePrimeDcWrapper)
}

// GenerateNewDeviceNotification returns the this method generates new device notification message.
//
// Parameters:
//
//	client: DLMS client settings.
//
// Returns:
//
//	Generated message.
func (g *GXDLMSPrimeDataConcentrator) GenerateNewDeviceNotification(client *GXDLMSClient) ([]byte, error) {
	bb := types.GXByteBuffer{}
	err := bb.SetUint8(uint8(enums.PrimeDcMsgTypeNewDeviceNotification))
	if err != nil {
		return nil, err
	}
	err = bb.SetUint16(g.DeviceID)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint16(g.Capabilities)
	if err != nil {
		return nil, err
	}
	err = bb.SetUint8(uint8(len(g.DlmsId)))
	if err != nil {
		return nil, err
	}
	err = bb.Set(g.DlmsId)
	if err != nil {
		return nil, err
	}
	err = bb.Set(g.Eui48)
	if err != nil {
		return nil, err
	}
	if client == nil {
		return bb.Array(), nil
	}
	return getWrapperFrame(client.Settings(), enums.CommandDataNotification, &bb)
}

// GenerateRemoveDeviceNotification returns the this method generates remove device notification message.
//
// Parameters:
//
//	client: DLMS client settings.
//
// Returns:
//
//	Generated message.
func (g *GXDLMSPrimeDataConcentrator) GenerateRemoveDeviceNotification(client *GXDLMSClient) ([]byte, error) {
	bb := types.GXByteBuffer{}
	err := bb.SetUint8(uint8(enums.PrimeDcMsgTypeRemoveDeviceNotification))
	if err != nil {
		return nil, err
	}
	err = bb.SetUint16(g.DeviceID)
	if err != nil {
		return nil, err
	}
	if client == nil {
		return bb.Array(), nil
	}
	return getWrapperFrame(client.Settings(), enums.CommandDataNotification, &bb)
}

// GenerateStartReportingMeters returns the this method generates start reporting meters message.
//
// Parameters:
//
//	client: DLMS client settings.
//
// Returns:
//
//	Generated message.
func (g *GXDLMSPrimeDataConcentrator) GenerateStartReportingMeters(client *GXDLMSClient) ([]byte, error) {
	bb := types.GXByteBuffer{}
	err := bb.SetUint8(uint8(enums.PrimeDcMsgTypeStartReportingMeters))
	if err != nil {
		return nil, err
	}
	if client == nil {
		return bb.Array(), nil
	}
	return getWrapperFrame(client.Settings(), enums.CommandSetRequest, &bb)
}

// GenerateDeleteMetersNotification returns the this method generates delete meters notification message.
//
// Parameters:
//
//	client: DLMS client settings.
//
// Returns:
//
//	Generated message.
func (g *GXDLMSPrimeDataConcentrator) GenerateDeleteMetersNotification(client *GXDLMSClient) ([]byte, error) {
	bb := types.GXByteBuffer{}
	err := bb.SetUint8(uint8(enums.PrimeDcMsgTypeDeleteMeters))
	if err != nil {
		return nil, err
	}
	err = bb.SetUint16(g.DeviceID)
	if err != nil {
		return nil, err
	}
	if client == nil {
		return bb.Array(), nil
	}
	return getWrapperFrame(client.Settings(), enums.CommandSetRequest, &bb)
}

// GenerateEnableAutoCloseNotification returns the this method generates enable auto close notification message.
//
// Parameters:
//
//	client: DLMS client settings.
//
// Returns:
//
//	Generated message.
func (g *GXDLMSPrimeDataConcentrator) GenerateEnableAutoCloseNotification(client *GXDLMSClient) ([]byte, error) {
	bb := types.GXByteBuffer{}
	err := bb.SetUint8(uint8(enums.PrimeDcMsgTypeEnableAutoClose))
	if err != nil {
		return nil, err
	}
	err = bb.SetUint16(g.DeviceID)
	if err != nil {
		return nil, err
	}
	if client == nil {
		return bb.Array(), nil
	}
	return getWrapperFrame(client.Settings(), enums.CommandSetRequest, &bb)
}

// GenerateDisableAutoCloseNotification returns the this method generates disable auto close notification message.
//
// Parameters:
//
//	client: DLMS client settings.
//
// Returns:
//
//	Generated message.
func (g *GXDLMSPrimeDataConcentrator) GenerateDisableAutoCloseNotification(client *GXDLMSClient) ([]byte, error) {
	bb := types.GXByteBuffer{}
	err := bb.SetUint8(uint8(enums.PrimeDcMsgTypeDisableAutoClose))
	if err != nil {
		return nil, err
	}
	err = bb.SetUint16(g.DeviceID)
	if err != nil {
		return nil, err
	}
	if client == nil {
		return bb.Array(), nil
	}
	return getWrapperFrame(client.Settings(), enums.CommandSetRequest, &bb)
}
