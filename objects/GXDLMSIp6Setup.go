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
	"fmt"
	"net"
	"strings"

	"github.com/Gurux/gxcommon-go"
	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// Online help:
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSIp6Setup
type GXDLMSIp6Setup struct {
	GXDLMSObject

	DataLinkLayerReference string
	AddressConfigMode      enums.AddressConfigMode
	UnicastIPAddress       []net.IP
	MulticastIPAddress     []net.IP
	GatewayIPAddress       []net.IP
	PrimaryDNSAddress      net.IP
	SecondaryDNSAddress    net.IP
	TrafficClass           uint8
	NeighborDiscoverySetup []GXNeighborDiscoverySetup
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSIp6Setup) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

// Invoke invokes object method.
func (g *GXDLMSIp6Setup) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	params := e.Parameters.(types.GXStructure)
	if len(params) < 2 {
		e.Error = enums.ErrorCodeReadWriteDenied
		return nil, nil
	}
	addrType := params[0].(uint8)
	address, err := toIPv6(params[1])
	if err != nil || address == nil {
		e.Error = enums.ErrorCodeReadWriteDenied
		return nil, nil
	}

	switch e.Index {
	case 1:
		g.addAddressByType(enums.IPv6AddressType(addrType), address)
	case 2:
		if !g.removeAddressByType(enums.IPv6AddressType(addrType), address) {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil, nil
}

// GetAttributeIndexToRead returns the collection of attributes to read.
func (g *GXDLMSIp6Setup) GetAttributeIndexToRead(all bool) []int {
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
func (g *GXDLMSIp6Setup) GetNames() []string {
	return []string{
		"Logical Name",
		"Data LinkLayer Reference",
		"Address Config Mode",
		"Unicast IP Address",
		"Multicast IP Address",
		"Gateway IP Address",
		"Primary DNS Address",
		"Secondary DNS Address",
		"Traffic Class",
		"Neighbor Discovery Setup",
	}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSIp6Setup) GetMethodNames() []string {
	return []string{"Add IP v6 address", "Remove IP v6 address"}
}

// GetAttributeCount returns the amount of attributes.
func (g *GXDLMSIp6Setup) GetAttributeCount() int {
	return 10
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSIp6Setup) GetMethodCount() int {
	return 2
}

// GetUIDataType returns UI data type for selected index.
func (g *GXDLMSIp6Setup) GetUIDataType(index int) enums.DataType {
	if index == 7 || index == 8 {
		return enums.DataTypeString
	}
	return g.GXDLMSObject.GetUIDataType(index)
}

// GetValues returns an array containing the COSEM object's attribute values.
func (g *GXDLMSIp6Setup) GetValues() []any {
	return []any{
		g.LogicalName(),
		g.DataLinkLayerReference,
		g.AddressConfigMode,
		g.UnicastIPAddress,
		g.MulticastIPAddress,
		g.GatewayIPAddress,
		ip6StringOrEmpty(g.PrimaryDNSAddress),
		ip6StringOrEmpty(g.SecondaryDNSAddress),
		g.TrafficClass,
		g.NeighborDiscoverySetup,
	}
}

// GetDataType returns data type of selected attribute index.
func (g *GXDLMSIp6Setup) GetDataType(index int) (enums.DataType, error) {
	switch index {
	case 1, 2, 7, 8:
		return enums.DataTypeOctetString, nil
	case 3:
		return enums.DataTypeEnum, nil
	case 4, 5, 6, 10:
		return enums.DataTypeArray, nil
	case 9:
		return enums.DataTypeUint8, nil
	default:
		return 0, dlmserrors.ErrInvalidAttributeIndex
	}
}

// GetValue returns value of given attribute.
func (g *GXDLMSIp6Setup) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	switch e.Index {
	case 1:
		return helpers.LogicalNameToBytes(g.LogicalName())
	case 2:
		return helpers.LogicalNameToBytes(g.DataLinkLayerReference)
	case 3:
		return uint8(g.AddressConfigMode), nil
	case 4:
		return encodeIPv6Array(settings, g.UnicastIPAddress)
	case 5:
		return encodeIPv6Array(settings, g.MulticastIPAddress)
	case 6:
		return encodeIPv6Array(settings, g.GatewayIPAddress)
	case 7:
		if g.PrimaryDNSAddress == nil {
			return nil, nil
		}
		return g.PrimaryDNSAddress.To16(), nil
	case 8:
		if g.SecondaryDNSAddress == nil {
			return nil, nil
		}
		return g.SecondaryDNSAddress.To16(), nil
	case 9:
		return g.TrafficClass, nil
	case 10:
		data := types.NewGXByteBuffer()
		if err := data.SetUint8(uint8(enums.DataTypeArray)); err != nil {
			return nil, err
		}
		if err := types.SetObjectCount(len(g.NeighborDiscoverySetup), data); err != nil {
			return nil, err
		}
		for _, it := range g.NeighborDiscoverySetup {
			if err := data.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
				return nil, err
			}
			if err := data.SetUint8(3); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, data, enums.DataTypeUint8, it.MaxRetry); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, data, enums.DataTypeUint16, it.RetryWaitTime); err != nil {
				return nil, err
			}
			if err := internal.SetData(settings, data, enums.DataTypeUint32, it.SendPeriod); err != nil {
				return nil, err
			}
		}
		return data.Array(), nil
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
		return nil, nil
	}
}

// SetValue sets value of given attribute.
func (g *GXDLMSIp6Setup) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	switch e.Index {
	case 1:
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		return g.SetLogicalName(ln)
	case 2:
		if v, ok := e.Value.(string); ok {
			g.DataLinkLayerReference = v
			return nil
		}
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
			return err
		}
		g.DataLinkLayerReference = ln
	case 3:
		value, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.AddressConfigMode = enums.AddressConfigMode(value)
	case 4:
		g.UnicastIPAddress = decodeIPv6Array(e.Value.(types.GXArray), true)
	case 5:
		g.MulticastIPAddress = decodeIPv6Array(e.Value.(types.GXArray), false)
	case 6:
		g.GatewayIPAddress = decodeIPv6Array(e.Value.(types.GXArray), false)
	case 7:
		ip, err := toIPv6(e.Value)
		if err != nil {
			return err
		}
		g.PrimaryDNSAddress = ip
	case 8:
		ip, err := toIPv6(e.Value)
		if err != nil {
			return err
		}
		g.SecondaryDNSAddress = ip
	case 9:
		value, err := toUint8(e.Value)
		if err != nil {
			return err
		}
		g.TrafficClass = value
	case 10:
		items := make([]GXNeighborDiscoverySetup, 0)
		for _, tmp := range e.Value.(types.GXArray) {
			parts := tmp.(types.GXStructure)
			if len(parts) < 3 {
				return gxcommon.ErrInvalidArgument
			}
			maxRetry := parts[0].(uint8)
			retryWait := parts[1].(uint16)
			sendPeriod := parts[2].(uint32)
			items = append(items, GXNeighborDiscoverySetup{
				MaxRetry:      maxRetry,
				RetryWaitTime: retryWait,
				SendPeriod:    sendPeriod,
			})
		}
		g.NeighborDiscoverySetup = items
	default:
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil
}

// Load loads object content from XML.
func (g *GXDLMSIp6Setup) Load(reader *GXXmlReader) error {
	var err error
	g.DataLinkLayerReference, err = reader.ReadElementContentAsString("DataLinkLayerReference", "")
	if err != nil {
		return err
	}
	mode, err := reader.ReadElementContentAsInt("AddressConfigMode", 0)
	if err != nil {
		return err
	}
	g.AddressConfigMode = enums.AddressConfigMode(mode)

	if g.UnicastIPAddress, err = loadIPv6AddressList(reader, "UnicastIPAddress"); err != nil {
		return err
	}
	if g.MulticastIPAddress, err = loadIPv6AddressList(reader, "MulticastIPAddress"); err != nil {
		return err
	}
	if g.GatewayIPAddress, err = loadIPv6AddressList(reader, "GatewayIPAddress"); err != nil {
		return err
	}
	if str, err := reader.ReadElementContentAsString("PrimaryDNSAddress", ""); err != nil {
		return err
	} else {
		g.PrimaryDNSAddress, _ = toIPv6(str)
	}
	if str, err := reader.ReadElementContentAsString("SecondaryDNSAddress", ""); err != nil {
		return err
	} else {
		g.SecondaryDNSAddress, _ = toIPv6(str)
	}
	if g.TrafficClass, err = reader.ReadElementContentAsUInt8("TrafficClass", 0); err != nil {
		return err
	}
	if g.NeighborDiscoverySetup, err = loadNeighborDiscovery(reader, "NeighborDiscoverySetup"); err != nil {
		return err
	}
	return nil
}

// Save saves object content to XML.
func (g *GXDLMSIp6Setup) Save(writer *GXXmlWriter) error {
	if err := writer.WriteElementString("DataLinkLayerReference", g.DataLinkLayerReference); err != nil {
		return err
	}
	if err := writer.WriteElementString("AddressConfigMode", int(g.AddressConfigMode)); err != nil {
		return err
	}
	if err := saveIPv6AddressList(writer, g.UnicastIPAddress, "UnicastIPAddress"); err != nil {
		return err
	}
	if err := saveIPv6AddressList(writer, g.MulticastIPAddress, "MulticastIPAddress"); err != nil {
		return err
	}
	if err := saveIPv6AddressList(writer, g.GatewayIPAddress, "GatewayIPAddress"); err != nil {
		return err
	}
	if g.PrimaryDNSAddress != nil {
		if err := writer.WriteElementString("PrimaryDNSAddress", g.PrimaryDNSAddress.String()); err != nil {
			return err
		}
	}
	if g.SecondaryDNSAddress != nil {
		if err := writer.WriteElementString("SecondaryDNSAddress", g.SecondaryDNSAddress.String()); err != nil {
			return err
		}
	}
	if err := writer.WriteElementString("TrafficClass", g.TrafficClass); err != nil {
		return err
	}
	if err := saveNeighborDiscovery(writer, g.NeighborDiscoverySetup, "NeighborDiscoverySetup"); err != nil {
		return err
	}
	return nil
}

// PostLoad handles actions after Load.
func (g *GXDLMSIp6Setup) PostLoad(reader *GXXmlReader) error {
	return nil
}

// AddAddress adds IP v6 address to the meter.
func (g *GXDLMSIp6Setup) AddAddress(client IGXDLMSClient, addressType enums.IPv6AddressType, address net.IP) ([][]byte, error) {
	data, err := encodeAddressAction(addressType, address)
	if err != nil {
		return nil, err
	}
	return client.Method(g, 1, data, enums.DataTypeStructure)
}

// RemoveAddress removes IP v6 address from the meter.
func (g *GXDLMSIp6Setup) RemoveAddress(client IGXDLMSClient, addressType enums.IPv6AddressType, address net.IP) ([][]byte, error) {
	data, err := encodeAddressAction(addressType, address)
	if err != nil {
		return nil, err
	}
	return client.Method(g, 2, data, enums.DataTypeStructure)
}

func encodeAddressAction(addressType enums.IPv6AddressType, address net.IP) ([]byte, error) {
	addr := address.To16()
	if addr == nil {
		return nil, fmt.Errorf("invalid IPv6 address")
	}
	bb := types.NewGXByteBuffer()
	if err := bb.SetUint8(uint8(enums.DataTypeStructure)); err != nil {
		return nil, err
	}
	if err := bb.SetUint8(2); err != nil {
		return nil, err
	}
	if err := bb.SetUint8(uint8(enums.DataTypeEnum)); err != nil {
		return nil, err
	}
	if err := bb.SetUint8(uint8(addressType)); err != nil {
		return nil, err
	}
	if err := bb.SetUint8(uint8(enums.DataTypeOctetString)); err != nil {
		return nil, err
	}
	if err := types.SetObjectCount(len(addr), bb); err != nil {
		return nil, err
	}
	if err := bb.Set(addr); err != nil {
		return nil, err
	}
	return bb.Array(), nil
}

func encodeIPv6Array(settings *settings.GXDLMSSettings, list []net.IP) ([]byte, error) {
	data := types.NewGXByteBuffer()
	if err := data.SetUint8(uint8(enums.DataTypeArray)); err != nil {
		return nil, err
	}
	if err := types.SetObjectCount(len(list), data); err != nil {
		return nil, err
	}
	for _, ip := range list {
		addr := ip.To16()
		if addr == nil {
			continue
		}
		if err := internal.SetData(settings, data, enums.DataTypeOctetString, addr); err != nil {
			return nil, err
		}
	}
	return data.Array(), nil
}

func decodeIPv6Array(value types.GXArray, ignoreInvalid bool) []net.IP {
	list := make([]net.IP, 0)
	for _, it := range value {
		ip, err := toIPv6(it)
		if err != nil {
			if ignoreInvalid {
				if b, ok := it.([]byte); ok {
					parsed := net.ParseIP(strings.TrimSpace(string(b)))
					if parsed != nil {
						list = append(list, parsed.To16())
					}
				}
				continue
			}
			continue
		}
		if ip != nil {
			list = append(list, ip)
		}
	}
	return list
}

func (g *GXDLMSIp6Setup) addAddressByType(addressType enums.IPv6AddressType, address net.IP) {
	switch addressType {
	case enums.IPv6AddressTypeUnicast:
		g.UnicastIPAddress = append(g.UnicastIPAddress, address)
	case enums.IPv6AddressTypeMulticast:
		g.MulticastIPAddress = append(g.MulticastIPAddress, address)
	case enums.IPv6AddressTypeGateway:
		g.GatewayIPAddress = append(g.GatewayIPAddress, address)
	}
}

func (g *GXDLMSIp6Setup) removeAddressByType(addressType enums.IPv6AddressType, address net.IP) bool {
	switch addressType {
	case enums.IPv6AddressTypeUnicast:
		updated, ok := removeIPv6FromSlice(g.UnicastIPAddress, address)
		if ok {
			g.UnicastIPAddress = updated
		}
		return ok
	case enums.IPv6AddressTypeMulticast:
		updated, ok := removeIPv6FromSlice(g.MulticastIPAddress, address)
		if ok {
			g.MulticastIPAddress = updated
		}
		return ok
	case enums.IPv6AddressTypeGateway:
		updated, ok := removeIPv6FromSlice(g.GatewayIPAddress, address)
		if ok {
			g.GatewayIPAddress = updated
		}
		return ok
	default:
		return false
	}
}

func removeIPv6FromSlice(list []net.IP, address net.IP) ([]net.IP, bool) {
	for pos, it := range list {
		if it.Equal(address) {
			return append(list[:pos], list[pos+1:]...), true
		}
	}
	return list, false
}

func toIPv6(value any) (net.IP, error) {
	switch v := value.(type) {
	case nil:
		return nil, nil
	case net.IP:
		ip := v.To16()
		if ip == nil {
			return nil, fmt.Errorf("invalid IPv6 address")
		}
		return ip, nil
	case string:
		if v == "" {
			return nil, nil
		}
		ip := net.ParseIP(v)
		if ip == nil {
			return nil, fmt.Errorf("invalid IPv6 address: %s", v)
		}
		return ip.To16(), nil
	case *types.GXByteBuffer:
		return toIPv6(v.Array())
	case []byte:
		if len(v) == 0 {
			return nil, nil
		}
		if len(v) == net.IPv6len {
			return net.IP(v), nil
		}
		if ip := net.ParseIP(string(v)); ip != nil {
			return ip.To16(), nil
		}
		if len(v)%2 == 0 {
			parts := make([]string, 0, len(v)/2)
			for pos := 0; pos < len(v); pos += 2 {
				parts = append(parts, fmt.Sprintf("%02x%02x", v[pos], v[pos+1]))
			}
			if ip := net.ParseIP(strings.Join(parts, ":")); ip != nil {
				return ip.To16(), nil
			}
		}
		return nil, fmt.Errorf("invalid IPv6 byte length: %d", len(v))
	default:
		return nil, fmt.Errorf("invalid IPv6 value type: %T", value)
	}
}

func ip6StringOrEmpty(value net.IP) string {
	if value == nil {
		return ""
	}
	if ip := value.To16(); ip != nil {
		return ip.String()
	}
	return ""
}

func loadIPv6AddressList(reader *GXXmlReader, name string) ([]net.IP, error) {
	list := make([]net.IP, 0)
	ok, err := reader.IsStartElementNamed(name, true)
	if err != nil {
		return nil, err
	}
	if !ok {
		return list, nil
	}
	for {
		ok, err = reader.IsStartElementNamed("Value", false)
		if err != nil {
			return nil, err
		}
		if !ok {
			break
		}
		value, err := reader.ReadElementContentAsString("Value", "")
		if err != nil {
			return nil, err
		}
		if ip := net.ParseIP(value); ip != nil {
			list = append(list, ip.To16())
		}
	}
	if err := reader.ReadEndElement(name); err != nil {
		return nil, err
	}
	return list, nil
}

func loadNeighborDiscovery(reader *GXXmlReader, name string) ([]GXNeighborDiscoverySetup, error) {
	list := make([]GXNeighborDiscoverySetup, 0)
	ok, err := reader.IsStartElementNamed(name, true)
	if err != nil {
		return nil, err
	}
	if !ok {
		return list, nil
	}
	for {
		ok, err = reader.IsStartElementNamed("Item", true)
		if err != nil {
			return nil, err
		}
		if !ok {
			break
		}
		item := GXNeighborDiscoverySetup{}
		if item.MaxRetry, err = reader.ReadElementContentAsUInt8("MaxRetry", 0); err != nil {
			return nil, err
		}
		if item.RetryWaitTime, err = reader.ReadElementContentAsUInt16("RetryWaitTime", 0); err != nil {
			return nil, err
		}
		if item.SendPeriod, err = reader.ReadElementContentAsUInt32("SendPeriod", 0); err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	if err := reader.ReadEndElement(name); err != nil {
		return nil, err
	}
	return list, nil
}

func saveIPv6AddressList(writer *GXXmlWriter, list []net.IP, name string) error {
	if list == nil {
		return nil
	}
	writer.WriteStartElement(name)
	for _, it := range list {
		if err := writer.WriteElementString("Value", it.String()); err != nil {
			return err
		}
	}
	writer.WriteEndElement()
	return nil
}

func saveNeighborDiscovery(writer *GXXmlWriter, list []GXNeighborDiscoverySetup, name string) error {
	if list == nil {
		return nil
	}
	writer.WriteStartElement(name)
	for _, it := range list {
		writer.WriteStartElement("Item")
		if err := writer.WriteElementString("MaxRetry", it.MaxRetry); err != nil {
			return err
		}
		if err := writer.WriteElementString("RetryWaitTime", it.RetryWaitTime); err != nil {
			return err
		}
		if err := writer.WriteElementString("SendPeriod", it.SendPeriod); err != nil {
			return err
		}
		writer.WriteEndElement()
	}
	writer.WriteEndElement()
	return nil
}

// NewGXDLMSIp6Setup creates a new Ip6 setup object instance.
func NewGXDLMSIp6Setup(ln string, sn int16) (*GXDLMSIp6Setup, error) {
	if err := ValidateLogicalName(ln); err != nil {
		return nil, err
	}
	return &GXDLMSIp6Setup{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeIP6Setup,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
