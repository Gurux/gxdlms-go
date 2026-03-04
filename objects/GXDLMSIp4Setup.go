package objects

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
	"encoding/binary"
	"fmt"
	"net"

	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// Online help:
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSIp4Setup
type GXDLMSIp4Setup struct {
	GXDLMSObject

	DataLinkLayerReference string
	IPAddress              net.IP
	MulticastIPAddress     []net.IP
	IPOptions              []GXDLMSIp4SetupIpOption
	SubnetMask             net.IP
	GatewayIPAddress       net.IP
	UseDHCP                bool
	PrimaryDNSAddress      net.IP
	SecondaryDNSAddress    net.IP
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSIp4Setup) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

// Invoke invokes method.
func (g *GXDLMSIp4Setup) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}

// GetAttributeIndexToRead returns the collection of attributes to read.
func (g *GXDLMSIp4Setup) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	if all || !g.IsRead(2) {
		attributes = append(attributes, 2)
	}
	if all || g.CanRead(3) {
		attributes = append(attributes, 3)
	}
	if all || g.CanRead(4) {
		attributes = append(attributes, 4)
	}
	if all || g.CanRead(5) {
		attributes = append(attributes, 5)
	}
	if all || g.CanRead(6) {
		attributes = append(attributes, 6)
	}
	if all || g.CanRead(7) {
		attributes = append(attributes, 7)
	}
	if all || !g.IsRead(8) {
		attributes = append(attributes, 8)
	}
	if all || g.CanRead(9) {
		attributes = append(attributes, 9)
	}
	if all || g.CanRead(10) {
		attributes = append(attributes, 10)
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSIp4Setup) GetNames() []string {
	return []string{
		"Logical Name",
		"Data LinkLayer Reference",
		"IP Address",
		"Multicast IP Address",
		"IP Options",
		"Subnet Mask",
		"Gateway IP Address",
		"Use DHCP",
		"Primary DNS Address",
		"Secondary DNS Address",
	}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSIp4Setup) GetMethodNames() []string {
	return []string{"Add mc IP address", "Delete mc IP address", "Get nbof mc IP addresses"}
}

// GetAttributeCount returns the amount of attributes.
func (g *GXDLMSIp4Setup) GetAttributeCount() int {
	return 10
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSIp4Setup) GetMethodCount() int {
	return 3
}

// GetValues returns an array containing the COSEM object's attribute values.
func (g *GXDLMSIp4Setup) GetValues() []any {
	return []any{
		g.LogicalName(),
		g.DataLinkLayerReference,
		ip4StringOrEmpty(g.IPAddress),
		g.MulticastIPAddress,
		g.IPOptions,
		ip4StringOrEmpty(g.SubnetMask),
		ip4StringOrEmpty(g.GatewayIPAddress),
		g.UseDHCP,
		ip4StringOrEmpty(g.PrimaryDNSAddress),
		ip4StringOrEmpty(g.SecondaryDNSAddress),
	}
}

// GetDataType returns the data type of selected attribute index.
func (g *GXDLMSIp4Setup) GetDataType(index int) (enums.DataType, error) {
	switch index {
	case 1, 2:
		return enums.DataTypeOctetString, nil
	case 3, 6, 7, 9, 10:
		return enums.DataTypeUint32, nil
	case 4, 5:
		return enums.DataTypeArray, nil
	case 8:
		return enums.DataTypeBoolean, nil
	default:
		return 0, dlmserrors.ErrInvalidAttributeIndex
	}
}

// GetValue returns the value of given attribute.
func (g *GXDLMSIp4Setup) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	switch e.Index {
	case 1:
		return helpers.LogicalNameToBytes(g.LogicalName())
	case 2:
		return helpers.LogicalNameToBytes(g.DataLinkLayerReference)
	case 3:
		return ip4ToUint32(g.IPAddress), nil
	case 4:
		data := types.NewGXByteBuffer()
		if err := data.SetUint8(uint8(enums.DataTypeArray)); err != nil {
			return nil, err
		}
		if err := types.SetObjectCount(len(g.MulticastIPAddress), data); err != nil {
			return nil, err
		}
		for _, it := range g.MulticastIPAddress {
			if err := internal.SetData(settings, data, enums.DataTypeUint32, ip4ToUint32(it)); err != nil {
				return nil, err
			}
		}
		return data.Array(), nil
	case 5:
		data := types.NewGXByteBuffer()
		if err := data.SetUint8(uint8(enums.DataTypeArray)); err != nil {
			return nil, err
		}
		if err := types.SetObjectCount(len(g.IPOptions), data); err != nil {
			return nil, err
		}
		for _, it := range g.IPOptions {
			if err := data.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
				return nil, err
			}
			if err := data.SetUint8(3); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, data, enums.DataTypeUint8, uint8(it.Type())); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, data, enums.DataTypeUint8, it.Length); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, data, enums.DataTypeOctetString, it.Data); err != nil {
				return nil, err
			}
		}
		return data.Array(), nil
	case 6:
		return ip4ToUint32(g.SubnetMask), nil
	case 7:
		return ip4ToUint32(g.GatewayIPAddress), nil
	case 8:
		return g.UseDHCP, nil
	case 9:
		return ip4ToUint32(g.PrimaryDNSAddress), nil
	case 10:
		return ip4ToUint32(g.SecondaryDNSAddress), nil
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		return nil, nil
	}
}

// SetValue sets value of given attribute.
func (g *GXDLMSIp4Setup) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	switch e.Index {
	case 1:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		return g.SetLogicalName(ln)
	case 2:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		g.DataLinkLayerReference = ln
	case 3:
		ip, err := toIPv4(e.Value)
		if err != nil {
			return err
		}
		g.IPAddress = ip
	case 4:
		ips := make([]net.IP, 0)
		for _, it := range e.Value.(types.GXArray) {
			ip, err := toIPv4(it)
			if err != nil {
				continue
			}
			if ip != nil {
				ips = append(ips, ip)
			}
		}
		g.MulticastIPAddress = ips
	case 5:
		opts := make([]GXDLMSIp4SetupIpOption, 0)
		for _, it := range e.Value.(types.GXArray) {
			parts := it.(types.GXStructure)
			if len(parts) < 3 {
				continue
			}
			item := GXDLMSIp4SetupIpOption{}
			optionType, err := toUint8(parts[0])
			if err != nil {
				continue
			}
			if err := item.SetType(enums.IP4SetupIpOptionType(optionType)); err != nil {
				continue
			}
			length, err := toUint8(parts[1])
			if err != nil {
				continue
			}
			item.Length = length
			switch v := parts[2].(type) {
			case []byte:
				item.Data = v
			case string:
				item.Data = types.HexToBytes(v)
			default:
				item.Data = nil
			}
			opts = append(opts, item)
		}
		g.IPOptions = opts
	case 6:
		ip, err := toIPv4(e.Value)
		if err != nil {
			return err
		}
		g.SubnetMask = ip
	case 7:
		ip, err := toIPv4(e.Value)
		if err != nil {
			return err
		}
		g.GatewayIPAddress = ip
	case 8:
		value, err := toBool(e.Value)
		if err != nil {
			return err
		}
		g.UseDHCP = value
	case 9:
		ip, err := toIPv4(e.Value)
		if err != nil {
			return err
		}
		g.PrimaryDNSAddress = ip
	case 10:
		ip, err := toIPv4(e.Value)
		if err != nil {
			return err
		}
		g.SecondaryDNSAddress = ip
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}

// Load loads object content from XML.
func (g *GXDLMSIp4Setup) Load(reader *GXXmlReader) error {
	var err error
	g.DataLinkLayerReference, err = reader.ReadElementContentAsString("DataLinkLayerReference", "")
	if err != nil {
		return err
	}

	if str, err := reader.ReadElementContentAsString("IPAddress", ""); err != nil {
		return err
	} else {
		g.IPAddress = parseIPv4String(str)
	}

	mc := make([]net.IP, 0)
	if ok, err := reader.IsStartElementNamed("MulticastIPAddress", true); err != nil {
		return err
	} else if ok {
		for {
			ok, err = reader.IsStartElementNamed("Value", false)
			if err != nil {
				return err
			}
			if !ok {
				break
			}
			v, err := reader.ReadElementContentAsString("Value", "")
			if err != nil {
				return err
			}
			mc = append(mc, parseIPv4String(v))
		}
		if err := reader.ReadEndElement("MulticastIPAddress"); err != nil {
			return err
		}
	}
	g.MulticastIPAddress = mc

	opts := make([]GXDLMSIp4SetupIpOption, 0)
	if ok, err := reader.IsStartElementNamed("IPOptions", true); err != nil {
		return err
	} else if ok {
		for {
			itemFound := false
			if itemFound, err = reader.IsStartElementNamed("IPOption", true); err != nil {
				return err
			}
			if !itemFound {
				if itemFound, err = reader.IsStartElementNamed("IPOptions", true); err != nil {
					return err
				}
			}
			if !itemFound {
				break
			}
			item := GXDLMSIp4SetupIpOption{}
			typeValue, err := reader.ReadElementContentAsInt("Type", 0)
			if err != nil {
				return err
			}
			_ = item.SetType(enums.IP4SetupIpOptionType(typeValue))
			length, err := reader.ReadElementContentAsInt("Length", 0)
			if err != nil {
				return err
			}
			item.Length = uint8(length)
			hex, err := reader.ReadElementContentAsString("Data", "")
			if err != nil {
				return err
			}
			item.Data = types.HexToBytes(hex)
			opts = append(opts, item)
		}
		if err := reader.ReadEndElement("IPOptions"); err != nil {
			return err
		}
	}
	g.IPOptions = opts

	if str, err := reader.ReadElementContentAsString("SubnetMask", ""); err != nil {
		return err
	} else {
		g.SubnetMask = parseIPv4String(str)
	}
	if str, err := reader.ReadElementContentAsString("GatewayIPAddress", ""); err != nil {
		return err
	} else {
		g.GatewayIPAddress = parseIPv4String(str)
	}
	if v, err := reader.ReadElementContentAsInt("UseDHCP", 0); err != nil {
		return err
	} else {
		g.UseDHCP = v != 0
	}
	if str, err := reader.ReadElementContentAsString("PrimaryDNSAddress", ""); err != nil {
		return err
	} else {
		g.PrimaryDNSAddress = parseIPv4String(str)
	}
	if str, err := reader.ReadElementContentAsString("SecondaryDNSAddress", ""); err != nil {
		return err
	} else {
		g.SecondaryDNSAddress = parseIPv4String(str)
	}
	return nil
}

// Save saves object content to XML.
func (g *GXDLMSIp4Setup) Save(writer *GXXmlWriter) error {
	if err := writer.WriteElementString("DataLinkLayerReference", g.DataLinkLayerReference); err != nil {
		return err
	}
	if err := writer.WriteElementString("IPAddress", ip4StringOrZero(g.IPAddress)); err != nil {
		return err
	}
	writer.WriteStartElement("MulticastIPAddress")
	for _, it := range g.MulticastIPAddress {
		if err := writer.WriteElementString("Value", ip4StringOrZero(it)); err != nil {
			return err
		}
	}
	writer.WriteEndElement()

	writer.WriteStartElement("IPOptions")
	for _, it := range g.IPOptions {
		writer.WriteStartElement("IPOption")
		if err := writer.WriteElementString("Type", int(it.Type())); err != nil {
			return err
		}
		if err := writer.WriteElementString("Length", it.Length); err != nil {
			return err
		}
		if err := writer.WriteElementString("Data", types.ToHex(it.Data, false)); err != nil {
			return err
		}
		writer.WriteEndElement()
	}
	writer.WriteEndElement()

	if err := writer.WriteElementString("SubnetMask", ip4StringOrZero(g.SubnetMask)); err != nil {
		return err
	}
	if err := writer.WriteElementString("GatewayIPAddress", ip4StringOrZero(g.GatewayIPAddress)); err != nil {
		return err
	}
	if err := writer.WriteElementString("UseDHCP", g.UseDHCP); err != nil {
		return err
	}
	if err := writer.WriteElementString("PrimaryDNSAddress", ip4StringOrZero(g.PrimaryDNSAddress)); err != nil {
		return err
	}
	if err := writer.WriteElementString("SecondaryDNSAddress", ip4StringOrZero(g.SecondaryDNSAddress)); err != nil {
		return err
	}
	return nil
}

// PostLoad handles actions after Load.
func (g *GXDLMSIp4Setup) PostLoad(reader *GXXmlReader) error {
	return nil
}

func ip4ToUint32(value net.IP) uint32 {
	if value == nil {
		return 0
	}
	ip := value.To4()
	if ip == nil {
		return 0
	}
	return binary.BigEndian.Uint32(ip)
}

func uint32ToIP4(value uint32) net.IP {
	if value == 0 {
		return net.IPv4zero
	}
	buff := make([]byte, 4)
	binary.BigEndian.PutUint32(buff, value)
	return net.IP(buff)
}

func parseIPv4String(value string) net.IP {
	if value == "" {
		return net.IPv4zero
	}
	if ip := net.ParseIP(value); ip != nil {
		if ip4 := ip.To4(); ip4 != nil {
			return ip4
		}
	}
	return net.IPv4zero
}

func ip4StringOrEmpty(value net.IP) string {
	if value == nil {
		return ""
	}
	if ip := value.To4(); ip != nil {
		return ip.String()
	}
	return ""
}

func ip4StringOrZero(value net.IP) string {
	if value == nil {
		return net.IPv4zero.String()
	}
	if ip := value.To4(); ip != nil {
		return ip.String()
	}
	return net.IPv4zero.String()
}

func toIPv4(value any) (net.IP, error) {
	switch v := value.(type) {
	case nil:
		return nil, nil
	case net.IP:
		if ip := v.To4(); ip != nil {
			return ip, nil
		}
		return nil, fmt.Errorf("invalid IPv4 address")
	case string:
		if v == "" {
			return nil, nil
		}
		ip := net.ParseIP(v)
		if ip == nil || ip.To4() == nil {
			return nil, fmt.Errorf("invalid IPv4 address: %s", v)
		}
		return ip.To4(), nil
	case []byte:
		if len(v) == 0 {
			return nil, nil
		}
		if len(v) == 4 {
			return net.IP(v), nil
		}
		if ip := net.ParseIP(string(v)); ip != nil && ip.To4() != nil {
			return ip.To4(), nil
		}
		return nil, fmt.Errorf("invalid IPv4 byte length: %d", len(v))
	case *types.GXByteBuffer:
		return toIPv4(v.Array())
	default:
		u32, err := toUint32(v)
		if err != nil {
			return nil, err
		}
		return uint32ToIP4(u32), nil
	}
}

// NewGXDLMSIp4Setup creates a new Ip4 setup object instance.
func NewGXDLMSIp4Setup(ln string, sn int16) (*GXDLMSIp4Setup, error) {
	if err := ValidateLogicalName(ln); err != nil {
		return nil, err
	}
	return &GXDLMSIp4Setup{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeIP4Setup,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
