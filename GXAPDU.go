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
// Gurux Device Framework is Open Source software you can redistribute it
// and/or modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation version 2 of the License.
// Gurux Device Framework is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU General Public License for more details.
//
// More information of Gurux products: https://www.gurux.org
//
// This code is licensed under the GNU General Public License v2.
// Full text may be retrieved at http://www.gnu.org/licenses/gpl-2.0.txt
//---------------------------------------------------------------------------

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/constants"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/objects"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

func getAuthenticationString(settings *settings.GXDLMSSettings, data *types.GXByteBuffer, ignoreACSE bool) error {
	if settings.Authentication != enums.AuthenticationNone ||
		(!ignoreACSE && settings.Cipher != nil && settings.Cipher.Security() != enums.SecurityNone) {
		err := data.SetUint8(uint8(constants.BerTypeContext | constants.BerType(internal.PduTypeSenderAcseRequirements)))
		if err != nil {
			return err
		}
		err = data.SetUint8(2)
		if err != nil {
			return err
		}
		err = data.SetUint8(uint8(constants.BerTypeBitString | constants.BerTypeOctetString))
		if err != nil {
			return err
		}
		err = data.SetUint8(0x80)
		if err != nil {
			return err
		}
		err = data.SetUint8(uint8(constants.BerTypeContext | constants.BerType(internal.PduTypeMechanismName)))
		if err != nil {
			return err
		}
		err = data.SetUint8(7)
		if err != nil {
			return err
		}
		oid := []byte{0x60, 0x85, 0x74, 0x05, 0x08, 0x02, byte(settings.Authentication)}
		err = data.Set(oid)
		if err != nil {
			return err
		}
	}
	if settings.Authentication == enums.AuthenticationNone {
		return nil
	}
	var authValue []byte
	if settings.Authentication == enums.AuthenticationLow {
		authValue = settings.Password
	} else {
		authValue = settings.CtoSChallenge()
	}
	l := len(authValue)
	err := data.SetUint8(uint8(constants.BerTypeContext | constants.BerTypeConstructed | constants.BerType(internal.PduTypeCallingAuthenticationValue)))
	if err != nil {
		return err
	}
	err = data.SetUint8(uint8(2 + l))
	if err != nil {
		return err
	}
	err = data.SetUint8(uint8(constants.BerTypeContext))
	if err != nil {
		return err
	}
	err = data.SetUint8(uint8(l))
	if err != nil {
		return err
	}
	if l != 0 {
		return data.Set(authValue)
	}
	return nil
}

// SetBitString converts bit string to DLMS bytes.
func SetBitString(buff *types.GXByteBuffer, value any, addCount bool) error {
	switch v := value.(type) {
	case types.GXBitString:
		value = v.Value()
	case *types.GXBitString:
		if v == nil {
			value = nil
		} else {
			value = v.Value()
		}
	}
	switch v := value.(type) {
	case string:
		if addCount {
			types.SetObjectCount(len(v), buff)
		}
		val := byte(0)
		index := 7
		for pos := 0; pos < len(v); pos++ {
			it := v[pos]
			if it == '1' {
				val |= 1 << index
			} else if it != '0' {
				return errors.New("Not a bit string.")
			}
			index--
			if index == -1 {
				index = 7
				err := buff.SetUint8(val)
				if err != nil {
					return err
				}
				val = 0
			}
		}
		if index != 7 {
			return buff.SetUint8(val)
		}
		return nil
	case []byte:
		types.SetObjectCount(8*len(v), buff)
		return buff.Set(v)
	case nil:
		return buff.SetUint8(0)
	case byte:
		types.SetObjectCount(8, buff)
		return buff.SetUint8(v)
	default:
		return errors.New("BitString must give as string.")
	}
}

// GenerateApplicationContextName returns the code application context name.
//
// Parameters:
//
//	settings: DLMS settings.
//	data: Byte buffer where data is saved.
//	cipher: Is ciphering settings.
func GenerateApplicationContextName(settings *settings.GXDLMSSettings, data *types.GXByteBuffer, cipher settings.GXICipher) error {
	var err error
	if settings.ProtocolVersion != "" {
		err = data.SetUint8(uint8(constants.BerTypeContext) | uint8(internal.PduTypeProtocolVersion))
		if err != nil {
			return err
		}
		err = data.SetUint8(2)
		if err != nil {
			return err
		}
		err = data.SetUint8(uint8(8 - len(settings.ProtocolVersion)))
		if err != nil {
			return err
		}
		err = SetBitString(data, settings.ProtocolVersion, false)
		if err != nil {
			return err
		}
	}
	err = data.SetUint8(uint8(constants.BerTypeContext | constants.BerTypeConstructed | constants.BerType(internal.PduTypeApplicationContextName)))
	if err != nil {
		return err
	}
	err = data.SetUint8(0x09)
	if err != nil {
		return err
	}
	err = data.SetUint8(uint8(constants.BerTypeObjectIdentifier))
	if err != nil {
		return err
	}
	err = data.SetUint8(0x07)
	if err != nil {
		return err
	}
	ciphered := settings.IsCiphered(true)
	for _, b := range []byte{0x60, 0x85, 0x74, 0x05, 0x08, 0x01} {
		err = data.SetUint8(b)
		if err != nil {
			return err
		}
	}
	if settings.UseLogicalNameReferencing() {
		err = data.SetUint8(map[bool]byte{true: 3, false: 1}[ciphered])
		if err != nil {
			return err
		}
	} else {
		err = data.SetUint8(map[bool]byte{true: 4, false: 2}[ciphered])
		if err != nil {
			return err
		}
	}
	if !settings.IsServer() && cipher != nil && (ciphered || settings.Authentication == enums.AuthenticationHighGMAC || settings.Authentication == enums.AuthenticationHighSHA256 || settings.Authentication == enums.AuthenticationHighECDSA) {
		st := cipher.SystemTitle()
		if len(st) == 0 {
			return fmt.Errorf("system title is required")
		}
		err = data.SetUint8(uint8(constants.BerTypeContext | constants.BerTypeConstructed | constants.BerType(internal.PduTypeCallingApTitle)))
		if err != nil {
			return err
		}
		err = data.SetUint8(uint8(2 + len(st)))
		if err != nil {
			return err
		}
		err = data.SetUint8(uint8(constants.BerTypeOctetString))
		if err != nil {
			return err
		}
		err = data.SetUint8(uint8(len(st)))
		if err != nil {
			return err
		}
		if err = data.Set(st); err != nil {
			return err
		}
		if settings.ClientPublicKeyCertificate != nil {
			raw, err := settings.ClientPublicKeyCertificate.Encoded()
			if err != nil {
				return err
			}
			err = data.SetUint8(uint8(constants.BerTypeContext | constants.BerTypeConstructed | constants.BerType(internal.PduTypeCallingAeQualifier)))
			if err != nil {
				return err
			}
			types.SetObjectCount(1+int(helpers.GetObjectCountSizeInBytes(len(raw)))+len(raw), data)
			err = data.SetUint8(uint8(constants.BerTypeOctetString))
			if err != nil {
				return err
			}
			types.SetObjectCount(len(raw), data)
			if err = data.Set(raw); err != nil {
				return err
			}
		}
	}
	if !settings.IsServer() && settings.UserID != -1 {
		if err := data.SetUint8(uint8(constants.BerTypeContext | constants.BerTypeConstructed | constants.BerType(internal.PduTypeCallingAeInvocationId))); err != nil {
			return err
		}
		for _, b := range []byte{3, byte(constants.BerTypeInteger), 1, byte(settings.UserID)} {
			if err := data.SetUint8(b); err != nil {
				return err
			}
		}
	}
	return nil
}

// getInitiateRequest returns the generate User information initiate request.
//
// Parameters:
//
//	settings: DLMS settings.
func getInitiateRequest(settings *settings.GXDLMSSettings, data *types.GXByteBuffer) error {
	err := data.SetUint8(uint8(enums.CommandInitiateRequest))
	if err != nil {
		return err
	}
	// Usage field for dedicated-key component.
	if settings.Cipher == nil || settings.Cipher.DedicatedKey() != nil || settings.Cipher.Security() == enums.SecurityNone {
		err = data.SetUint8(0x00)
		if err != nil {
			return err
		}
	} else {
		err = data.SetUint8(0x1)
		if err != nil {
			return err
		}
		err = types.SetObjectCount(len(settings.Cipher.DedicatedKey()), data)
		err = data.Set(settings.Cipher.DedicatedKey())
		if err != nil {
			return err
		}
	}
	err = data.SetUint8(0)
	if err != nil {
		return err
	}
	// Usage field of the proposed-quality-of-service component. Not used
	if settings.QualityOfService == 0 {
		err = data.SetUint8(0x00)
		if err != nil {
			return err
		}
	} else {
		err = data.SetUint8(0x01)
		if err != nil {
			return err
		}
		err = data.SetUint8(settings.QualityOfService)
		if err != nil {
			return err
		}
	}
	err = data.SetUint8(settings.DLMSVersion)
	if err != nil {
		return err
	}
	err = data.SetUint8(0x5F)
	if err != nil {
		return err
	}
	err = data.SetUint8(0x1F)
	if err != nil {
		return err
	}
	err = data.SetUint8(0x04)
	if err != nil {
		return err
	}
	err = data.SetUint8(0x00)
	if err != nil {
		return err
	}
	bs, err := types.NewGXBitStringFromInteger(int(settings.ProposedConformance), 24)
	err = data.Set(bs.Value())
	if err != nil {
		return err
	}
	err = data.SetUint16(settings.MaxPduSize())
	if err != nil {
		return err
	}
	return err
}

func GetConformance(value int, xml *settings.GXDLMSTranslatorStructure) {
	if xml.OutputType() == enums.TranslatorOutputTypeSimpleXML {
		for _, it := range enums.AllConformance() {
			if int(it)&value != 0 {
				xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsConformanceBit), "Name", it.String())
			}
		}
	} else {
		for _, it := range enums.AllConformance() {
			if int(it)&value != 0 {
				xml.AppendString(it.String() + " ")
			}
		}
	}
}

func parse(initiateRequest bool,
	settings *settings.GXDLMSSettings,
	cipher settings.GXICipher,
	data *types.GXByteBuffer,
	xml *settings.GXDLMSTranslatorStructure, tag uint8) error {
	var err error
	response := tag == uint8(enums.CommandInitiateResponse)
	if response {
		if xml != nil {
			xml.AppendStartTag(int(enums.CommandInitiateResponse), "", "", false)
		}
		tag, err := data.Uint8()
		if err != nil {
			return err
		}
		if tag != 0 {
			ret, err := data.Uint8()
			if err != nil {
				return err
			}
			settings.QualityOfService = ret
			if xml != nil {
				xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsNegotiatedQualityOfService), "Value", strconv.Itoa(int(ret)))
			}
		}
	} else if tag == uint8(enums.CommandInitiateRequest) {
		if xml != nil {
			xml.AppendStartTag(int(enums.CommandInitiateRequest), "", "", false)
		}
		tag, err := data.Uint8()
		if err != nil {
			return err
		}
		if settings.Cipher != nil {
			settings.Cipher.SetDedicatedKey(nil)
		}
		if tag != 0 {
			len_, err := data.Uint8()
			if err != nil {
				return err
			}
			tmp2 := make([]byte, len_)
			err = data.Get(tmp2)
			if err != nil {
				return err
			}
			if settings.Cipher != nil {
				settings.Cipher.SetDedicatedKey(tmp2)
			}
			if xml != nil {
				xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsDedicatedKey), "", types.ToHex(tmp2, false))
			}
		}
		tag, err = data.Uint8()
		if err != nil {
			return err
		}
		if tag != 0 {
			ret, err := data.Uint8()
			if err != nil {
				return err
			}
			settings.QualityOfService = ret
			if xml != nil && (initiateRequest || xml.OutputType() == enums.TranslatorOutputTypeSimpleXML) {
				xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsProposedQualityOfService), "", strconv.Itoa(int(settings.QualityOfService)))
			}
		} else {
			if xml != nil && xml.OutputType() == enums.TranslatorOutputTypeStandardXML {
				xml.AppendLineFromTag(int(internal.TranslatorTagsResponseAllowed), "", "true")
			}
		}
		tag, err = data.Uint8()
		if err != nil {
			return err
		}
		if tag != 0 {
			ret, err := data.Uint8()
			if err != nil {
				return err
			}
			settings.QualityOfService = ret
			if xml != nil && xml.OutputType() == enums.TranslatorOutputTypeSimpleXML {
				xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsProposedQualityOfService), "", strconv.Itoa(int(settings.QualityOfService)))
			}
		}
	} else if tag == uint8(enums.CommandConfirmedServiceError) {
		if xml != nil {
			xml.AppendStartTag(int(enums.CommandConfirmedServiceError), "", "", false)
			if xml.OutputType() == enums.TranslatorOutputTypeStandardXML {
				data.Uint8()
				xml.AppendStartTag(int(internal.TranslatorTagsInitiateError), "", "", false)
				ret, err := data.Uint8()
				if err != nil {
					return err
				}
				type_ := enums.ServiceError(ret)
				str := standardServiceErrorToString(type_)
				ret, err = data.Uint8()
				if err != nil {
					return err
				}
				value := standardGetServiceErrorValue(type_, ret)
				xml.AppendLine("x:"+str, "", value)
				xml.AppendEndTag(int(internal.TranslatorTagsInitiateError), false)
			} else {
				ret, err := data.Uint8()
				if err != nil {
					return err
				}
				xml.AppendLineFromTag(int(internal.TranslatorTagsService), "Value", xml.IntegerToHex(ret, 2, false))
				ret, err = data.Uint8()
				if err != nil {
					return err
				}
				type_ := enums.ServiceError(ret)
				xml.AppendStartTag(int(internal.TranslatorTagsServiceError), "", "", false)
				ret, err = data.Uint8()
				if err != nil {
					return err
				}
				xml.AppendLineFromTag(int(internal.TranslatorTagsServiceError), "Value", simpleGetServiceErrorValue(type_, ret))
				xml.AppendEndTag(int(internal.TranslatorTagsServiceError), false)
			}
			xml.AppendEndTag(int(enums.CommandConfirmedServiceError), false)
			return nil
		}
		ret, err := data.Uint8()
		if err != nil {
			return err
		}
		ret1, err := data.Uint8()
		if err != nil {
			return err
		}
		ret2, err := data.Uint8()
		if err != nil {
			return err
		}
		return dlmserrors.NewGXDLMSConfirmedServiceError(enums.ConfirmedServiceError(ret), enums.ServiceError(ret1), ret2)
	} else {
		if xml != nil {
			xml.AppendComment("Error: Failed to decrypt data.")
			data.SetPosition(data.Size())
			return nil
		}
		return errors.New("Invalid tag.")
	}
	// Get DLMS version number.
	if !response {
		settings.DLMSVersion, err = data.Uint8()
		if err != nil {
			return err
		}
		// ProposedDlmsVersionNumber
		if xml != nil && (initiateRequest || xml.OutputType() == enums.TranslatorOutputTypeSimpleXML) {
			xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsProposedDlmsVersionNumber), "", xml.IntegerToHex(settings.DLMSVersion, 2, false))
		}
	} else {
		ret, err := data.Uint8()
		if err != nil {
			return err
		}
		if ret != 6 {
			return errors.New("Invalid DLMS version number.")
		}
		if xml != nil && (initiateRequest || xml.OutputType() == enums.TranslatorOutputTypeSimpleXML) {
			xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsNegotiatedDlmsVersionNumber), "", xml.IntegerToHex(settings.DLMSVersion, 2, false))
		}
	}
	tag, err = data.Uint8()
	if err != nil {
		return err
	}
	if tag != 0x5F {
		return errors.New("Invalid tag.")
	}
	// Old Way...
	ret, err := data.Uint8At(data.Position())
	if err != nil {
		return err
	}
	if ret == 0x1F {
		data.Uint8()
	}
	//len
	_, err = data.Uint8()
	if err != nil {
		return err
	}

	conformance := make([]byte, 4)
	data.Get(conformance)
	bs, err := types.NewGXBitStringFromByteArray(conformance)
	if err != nil {
		return err
	}
	v := bs.ToInteger()
	if settings.IsServer() {
		settings.NegotiatedConformance = enums.Conformance(v) & settings.ProposedConformance
		if xml != nil {
			xml.AppendStartTag(int(enums.TranslatorGeneralTagsProposedConformance), "", "", xml.OutputType() == enums.TranslatorOutputTypeStandardXML)
			GetConformance(v, xml)
		}
	} else {
		if xml != nil {
			xml.AppendStartTag(int(enums.TranslatorGeneralTagsNegotiatedConformance), "", "", xml.OutputType() == enums.TranslatorOutputTypeStandardXML)
			GetConformance(v, xml)
		}
		settings.NegotiatedConformance = enums.Conformance(v)
	}
	if !response {
		// Proposed max PDU size.
		maxPdu, err := data.Uint16()
		if err != nil {
			return err
		}
		settings.SetMaxPduSize(maxPdu)
		if xml != nil {
			// ProposedConformance closing
			if xml.OutputType() == enums.TranslatorOutputTypeSimpleXML {
				xml.AppendEndTag(int(enums.TranslatorGeneralTagsProposedConformance), false)
			} else if initiateRequest {
				xml.AppendEndTag(int(enums.TranslatorGeneralTagsProposedConformance), true)
			}
			xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsProposedMaxPduSize), "", xml.IntegerToHex(settings.MaxPduSize, 4, false))
		}
		// If client asks too high PDU.
		if settings.MaxPduSize() > settings.GetMaxServerPDUSize() {
			settings.SetMaxPduSize(settings.GetMaxServerPDUSize())
		}
	} else {
		ret, err := data.Uint16()
		if err != nil {
			return err
		}
		settings.SetMaxPduSize(ret)
		if xml != nil {
			// NegotiatedConformance closing
			if xml.OutputType() == enums.TranslatorOutputTypeSimpleXML {
				xml.AppendEndTag(int(enums.TranslatorGeneralTagsNegotiatedConformance), false)
			} else if initiateRequest {
				xml.Append(int(enums.TranslatorGeneralTagsNegotiatedConformance), false)
			}
			xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsNegotiatedMaxPduSize), "", xml.IntegerToHex(settings.MaxPduSize(), 4, false))
		}
		// If client asks too high PDU.
		if settings.MaxPduSize() > settings.GetMaxServerPDUSize() {
			settings.SetMaxPduSize(settings.GetMaxServerPDUSize())
		}
	}
	if response {
		tag, err := data.Uint16()
		if err != nil {
			return err
		}
		if xml != nil {
			if initiateRequest || xml.OutputType() == enums.TranslatorOutputTypeSimpleXML {
				xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsVaaName), "", xml.IntegerToHex(tag, 4, false))
			}
		}
		switch tag {
		case 0x0007:
			if initiateRequest {
				settings.SetUseLogicalNameReferencing(true)
			} else {
				// If LN
				if !settings.UseLogicalNameReferencing() && xml == nil {
					return errors.New("Invalid VAA.")
				}
			}
		case 0xFA00:
			// If SN
			if initiateRequest {
				settings.SetUseLogicalNameReferencing(false)
			} else {
				if settings.UseLogicalNameReferencing() {
					return errors.New("Invalid VAA.")
				}
			}
		default:
			return errors.New("Invalid VAA.")
		}
		if xml != nil {
			xml.AppendEndTag(int(enums.CommandInitiateResponse), false)
		}
	} else if xml != nil {
		xml.AppendEndTag(int(enums.CommandInitiateRequest), false)
	}
	return nil
}

// parseApplicationContextName returns the parse application context name.
//
// Parameters:
//
//	settings: DLMS settings.
//	buff: Received data.
func parseApplicationContextName(settings *settings.GXDLMSSettings,
	buff *types.GXByteBuffer,
	xml *settings.GXDLMSTranslatorStructure) (enums.ApplicationContextName, error) {
	// Get length.
	len_, err := buff.Uint8()
	if err != nil {
		return enums.ApplicationContextNameUnknown, err
	}
	if buff.Size()-buff.Position() < int(len_) {
		return enums.ApplicationContextNameUnknown, errors.New("Encoding failed. Not enough data.")
	}
	ret, err := buff.Uint8()
	if err != nil {
		return enums.ApplicationContextNameUnknown, err
	}
	if ret != 6 {
		return enums.ApplicationContextNameUnknown, errors.New("Encoding failed. Not an Object ID.")
	}
	if settings.IsServer() && settings.Cipher != nil {
		settings.Cipher.SetSecurity(enums.SecurityNone)
	}
	len_, err = buff.Uint8()
	if err != nil {
		return enums.ApplicationContextNameUnknown, err
	}
	tmp := make([]byte, len_)
	err = buff.Get(tmp)
	if err != nil {
		return enums.ApplicationContextNameUnknown, err
	}
	if tmp[0] != 0x60 || tmp[1] != 0x85 || tmp[2] != 0x74 || tmp[3] != 0x5 || tmp[4] != 0x8 || tmp[5] != 0x1 {
		if xml != nil {
			xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsApplicationContextName), "Value", "UNKNOWN")
			return enums.ApplicationContextNameUnknown, nil
		}
		return enums.ApplicationContextNameUnknown, errors.New("Encoding failed. Invalid Application context name.")
	}
	name := tmp[6]
	if xml != nil {
		if name == 1 {
			if xml.OutputType() == enums.TranslatorOutputTypeSimpleXML {
				xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsApplicationContextName), "Value", "LN")
			} else {
				xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsApplicationContextName), "", "1")
			}
			settings.SetUseLogicalNameReferencing(true)
		} else if name == 3 {
			if xml.OutputType() == enums.TranslatorOutputTypeSimpleXML {
				xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsApplicationContextName), "Value", "LN_WITH_CIPHERING")
			} else {
				xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsApplicationContextName), "", "3")
			}
			settings.SetUseLogicalNameReferencing(true)
		} else if name == 2 {
			if xml.OutputType() == enums.TranslatorOutputTypeSimpleXML {
				xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsApplicationContextName), "Value", "SN")
			} else {
				xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsApplicationContextName), "", "2")
			}
			settings.SetUseLogicalNameReferencing(false)
		} else if name == 4 {
			if xml.OutputType() == enums.TranslatorOutputTypeSimpleXML {
				xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsApplicationContextName), "Value", "SN_WITH_CIPHERING")
			} else {
				xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsApplicationContextName), "", "4")
			}
			settings.SetUseLogicalNameReferencing(false)
		} else {
			if xml.OutputType() == enums.TranslatorOutputTypeSimpleXML {
				xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsApplicationContextName), "Value", "UNKNOWN")
			} else {
				xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsApplicationContextName), "", "5")
			}
		}
		return enums.ApplicationContextNameUnknown, nil
	}
	if ln, ok := settings.AssignedAssociation().(*objects.GXDLMSAssociationLogicalName); ok {
		if byte(ln.ApplicationContextName.ContextId) == name {
			return enums.ApplicationContextNameUnknown, nil
		}
		//All connections are accepted if the There might be only one association view in some test meters.
		if settings.UseLogicalNameReferencing() &&
			len(settings.Objects.(*objects.GXDLMSObjectCollection).GetObjects(enums.ObjectTypeAssociationLogicalName)) == 1 {
			return enums.ApplicationContextNameUnknown, nil
		}
		if !settings.UseLogicalNameReferencing() &&
			len(settings.Objects.(*objects.GXDLMSObjectCollection).GetObjects(enums.ObjectTypeAssociationShortName)) == 1 {
			return enums.ApplicationContextNameUnknown, nil
		}
		return ln.ApplicationContextName.ContextId, nil
	}
	if settings.UseLogicalNameReferencing() {
		if name == uint8(enums.ApplicationContextNameLogicalName) && (settings.Cipher == nil || settings.Cipher.Security() == enums.SecurityNone) {
			return enums.ApplicationContextNameUnknown, nil
		}
		// If ciphering is used.
		if name == uint8(enums.ApplicationContextNameLogicalNameWithCiphering) && (settings.Cipher != nil && settings.Cipher.Security() != enums.SecurityNone) {
			return enums.ApplicationContextNameUnknown, nil
		}
	} else {
		if name == uint8(enums.ApplicationContextNameShortName) && (settings.Cipher == nil || settings.Cipher.Security() == enums.SecurityNone) {
			return enums.ApplicationContextNameUnknown, nil
		}
		// If ciphering is used.
		if name == uint8(enums.ApplicationContextNameShortNameWithCiphering) && (settings.Cipher != nil && settings.Cipher.Security() != enums.SecurityNone) {
			return enums.ApplicationContextNameUnknown, nil
		}
	}
	return enums.ApplicationContextName(name), nil
}

func UpdateAuthentication(settings *settings.GXDLMSSettings, buff *types.GXByteBuffer) error {
	_, err := buff.Uint8()
	if err != nil {
		return err
	}
	ret, err := buff.Uint8()
	if err != nil {
		return err
	}
	if ret != 0x60 {
		return errors.New("Invalid tag.")
	}
	ret, err = buff.Uint8()
	if err != nil {
		return errors.New("Invalid tag.")
	}
	if ret != 0x85 {
		return errors.New("Invalid tag.")
	}
	ret, err = buff.Uint8()
	if err != nil {
		return errors.New("Invalid tag.")
	}
	if ret != 0x74 {
		return errors.New("Invalid tag.")
	}
	ret, err = buff.Uint8()
	if err != nil {
		return errors.New("Invalid tag.")
	}
	if ret != 0x05 {
		return errors.New("Invalid tag.")
	}
	ret, err = buff.Uint8()
	if err != nil {
		return errors.New("Invalid tag.")
	}

	if ret != 0x08 {
		return errors.New("Invalid tag.")
	}
	ret, err = buff.Uint8()
	if err != nil {
		return errors.New("Invalid tag.")
	}
	if ret != 0x02 {
		return errors.New("Invalid tag.")
	}
	ret, err = buff.Uint8()
	if err != nil {
		return errors.New("Invalid tag.")
	}
	if ret > 7 {
		return errors.New("Invalid tag.")
	}
	settings.Authentication = enums.Authentication(ret)
	return nil
}

func AppendServerSystemTitleToXml(settings *settings.GXDLMSSettings, xml *settings.GXDLMSTranslatorStructure, tag int) {
	if xml != nil {
		// RespondingAuthentication
		if xml.OutputType() == enums.TranslatorOutputTypeSimpleXML {
			xml.AppendLineFromTag(tag, "Value", types.ToHex(settings.StoCChallenge(), false))
		} else {
			xml.Append(tag, true)
			xml.Append(int(enums.TranslatorGeneralTagsCharString), true)
			xml.AppendString(types.ToHex(settings.StoCChallenge(), false))
			xml.Append(int(enums.TranslatorGeneralTagsCharString), false)
			xml.Append(tag, false)
			xml.AppendString("\n")
		}
	}
}

func ParseProtocolVersion(settings *settings.GXDLMSSettings,
	buff *types.GXByteBuffer,
	xml *settings.GXDLMSTranslatorStructure) (enums.AcseServiceProvider, error) {
	_, err := buff.Uint8()
	if err != nil {
		return enums.AcseServiceProviderNone, err
	}
	unusedBits, err := buff.Uint8()
	if err != nil {
		return enums.AcseServiceProviderNone, err
	}
	if unusedBits > 8 {
		return enums.AcseServiceProviderNone, errors.New("unusedBits")
	}
	value, err := buff.Uint8()
	if err != nil {
		return enums.AcseServiceProviderNone, err
	}
	sb := strings.Builder{}
	types.ToBitString(&sb, value, 8-int(unusedBits))
	settings.ProtocolVersion = sb.String()
	if xml != nil {
		xml.AppendLineFromTag(int(internal.TranslatorTagsProtocolVersion), "Value", settings.ProtocolVersion)
	} else {
		if settings.ProtocolVersion != "100001" {
			return enums.AcseServiceProviderNoCommonAcseVersion, nil
		}
	}
	return enums.AcseServiceProviderNone, nil
}

func UpdatePassword(settings *settings.GXDLMSSettings,
	buff *types.GXByteBuffer,
	xml *settings.GXDLMSTranslatorStructure) error {
	len_, err := buff.Uint8()
	if err != nil {
		return err
	}
	// Get authentication information.
	ret, err := buff.Uint8()
	if err != nil {
		return err
	}
	if ret != 0x80 {
		return errors.New("Invalid tag.")
	}
	len_, err = buff.Uint8()
	if err != nil {
		return err
	}
	if settings.Authentication == enums.AuthenticationLow {
		settings.Password = make([]byte, len_)
		buff.Get(settings.Password)
	} else {
		tmp := make([]byte, len_)
		err := buff.Get(tmp)
		if err != nil {
			return err
		}
		settings.SetCtoSChallenge(tmp)
	}
	if xml != nil {
		if xml.OutputType() == enums.TranslatorOutputTypeSimpleXML {
			if settings.Authentication == enums.AuthenticationLow {
				xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsCallingAuthentication), "Value", types.ToHex(settings.Password, false))
			} else {
				xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsCallingAuthentication), "Value", types.ToHex(settings.CtoSChallenge(), false))
			}
		} else {
			xml.AppendStartTag(int(enums.TranslatorGeneralTagsCallingAuthentication), "", "", true)
			xml.AppendStartTag(int(enums.TranslatorGeneralTagsCharString), "", "", true)
			if settings.Authentication == enums.AuthenticationLow {
				xml.AppendString(types.ToHex(settings.Password, false))
			} else {
				xml.AppendString(types.ToHex(settings.CtoSChallenge(), false))
			}
			xml.AppendEndTag(int(enums.TranslatorGeneralTagsCharString), false)
			xml.AppendEndTag(int(enums.TranslatorGeneralTagsCallingAuthentication), false)
		}
	}
	return nil
}

// GenerateUserInformation returns the generate user information.
//
// Parameters:
//
//	settings: DLMS settings.
//	data: Generated user information.
func GenerateUserInformation(conf *settings.GXDLMSSettings, cipher settings.GXICipher, encryptedData *types.GXByteBuffer, data *types.GXByteBuffer) error {
	err := data.SetUint8(uint8(constants.BerTypeContext) | uint8(constants.BerTypeConstructed) | uint8(internal.PduTypeUserInformation))
	if err != nil {
		return err
	}
	if !conf.IsCiphered(true) {
		err = data.SetUint8(0x10)
		if err != nil {
			return err
		}
		err = data.SetUint8(uint8(constants.BerTypeOctetString))
		if err != nil {
			return err
		}
		tmp := types.GXByteBuffer{}
		err = getInitiateRequest(conf, &tmp)
		if err != nil {
			return err
		}
		err = types.SetObjectCount(tmp.Size(), data)
		if err != nil {
			return err
		}
		err = data.SetByteBuffer(&tmp)
		if err != nil {
			return err
		}
	} else {
		if encryptedData != nil && encryptedData.Size() != 0 {
			err = data.SetUint8(uint8((4 + encryptedData.Size())))
			if err != nil {
				return err
			}
			err = data.SetUint8(uint8(constants.BerTypeOctetString))
			if err != nil {
				return err
			}
			err = data.SetUint8(uint8((2 + encryptedData.Size())))
			if err != nil {
				return err
			}
			err = data.SetUint8(uint8(enums.CommandGloInitiateRequest))
			if err != nil {
				return err
			}
			err = data.SetUint8(uint8(encryptedData.Size()))
			if err != nil {
				return err
			}
			err = data.SetByteBuffer(encryptedData)
			if err != nil {
				return err
			}
		} else {
			value := types.GXByteBuffer{}
			getInitiateRequest(conf, &value)
			p := settings.NewAesGcmParameter(
				byte(enums.CommandGloInitiateRequest),
				conf, cipher.Security(), cipher.SecuritySuite(),
				uint64(cipher.InvocationCounter()), cipher.SystemTitle(),
				cipher.BlockCipherKey(), cipher.AuthenticationKey())
			crypted, err := settings.EncryptAesGcm(p, value.Array())
			if err != nil {
				return err
			}
			cipher.SetInvocationCounter(cipher.InvocationCounter() + 1)
			err = data.SetUint8(uint8((2 + len(crypted))))
			if err != nil {
				return err
			}
			err = data.SetUint8(uint8(constants.BerTypeOctetString))
			if err != nil {
				return err
			}
			err = data.SetUint8(uint8(len(crypted)))
			if err != nil {
				return err
			}
			err = data.Set(crypted)
			if err != nil {
				return err
			}
		}
	}
	return err
}

// GenerateAarq returns the generates Aarq.
func generateAarq(settings *settings.GXDLMSSettings,
	cipher settings.GXICipher,
	encryptedData *types.GXByteBuffer,
	data *types.GXByteBuffer) error {
	tmp := types.GXByteBuffer{}
	err := GenerateApplicationContextName(settings, &tmp, cipher)
	if err != nil {
		return err
	}
	err = getAuthenticationString(settings, &tmp, encryptedData != nil && encryptedData.Size() != 0)
	if err != nil {
		return err
	}

	err = GenerateUserInformation(settings, cipher, encryptedData, &tmp)
	if err != nil {
		return err
	}
	err = data.SetUint8(uint8(constants.BerTypeApplication | constants.BerTypeConstructed))
	if err != nil {
		return err
	}
	types.SetObjectCount(tmp.Size(), data)
	err = data.SetByteBuffer(&tmp)
	if err != nil {
		return err
	}
	return err
}

func parseInitiate(initiateRequest bool,
	conf *settings.GXDLMSSettings,
	cipher settings.GXICipher,
	data *types.GXByteBuffer,
	xml *settings.GXDLMSTranslatorStructure) (enums.ExceptionServiceError, error) {
	var err error
	var originalPos int
	// Tag for xDLMS-Initate.response
	tag, err := data.Uint8()
	if err != nil {
		return 0, err
	}
	var tmp []byte
	var encrypted []byte
	if tag == uint8(enums.CommandGloInitiateResponse) || tag == uint8(enums.CommandGloInitiateRequest) || tag == uint8(enums.CommandDedInitiateResponse) || tag == uint8(enums.CommandDedInitiateRequest) || tag == uint8(enums.CommandGeneralGloCiphering) || tag == uint8(enums.CommandGeneralDedCiphering) {
		if xml != nil {
			originalPos = data.Position()
			var st []byte
			var cnt int
			if tag == uint8(enums.CommandGeneralGloCiphering) || tag == uint8(enums.CommandGeneralDedCiphering) {
				cnt, err = types.GetObjectCount(data)
				if err != nil {
					return 0, err
				}
				st = make([]byte, cnt)
				data.Get(st)
			} else {
				if tag == uint8(enums.CommandGloInitiateRequest) || tag == uint8(enums.CommandDedInitiateRequest) {
					st = conf.Cipher.SystemTitle()
				} else {
					st = conf.SourceSystemTitle()
				}
			}
			cnt, err = types.GetObjectCount(data)
			if err != nil {
				return 0, err
			}
			encrypted = make([]byte, cnt)
			err = data.Get(encrypted)
			if err != nil {
				return 0, err
			}
			if st != nil && cipher != nil && cipher.BlockCipherKey() != nil && cipher.AuthenticationKey() != nil && xml.Comments {
				pos := xml.GetXmlLength()
				pos2 := data.Position()
				data.SetPosition(originalPos - 1)
				p := settings.NewAesGcmParameter3(conf, st, conf.Cipher.BlockCipherKey(), conf.Cipher.AuthenticationKey())
				p.Xml = xml
				tmp, err = settings.DecryptAesGcm(p, data)
				if err != nil {
					return 0, err
				}
				data.Clear()
				err = data.Set(tmp)
				if err != nil {
					return 0, err
				}
				cipher.SetSecurity(p.Security())
				tag1, err := data.Uint8()
				if err != nil {
					return 0, err
				}
				xml.StartComment("Decrypted data:")
				xml.AppendStringLine("Security: " + p.Security().String())
				xml.AppendStringLine("Invocation Counter: " + strconv.FormatUint(p.InvocationCounter, 10))
				err = parse(initiateRequest, conf, cipher, data, xml, tag1)
				if err != nil {
					// It's OK if this fails.
					xml.SetXmlLength(pos)
					data.SetPosition(pos2)
					return 0, err
				}
				xml.EndComment()
			}
			xml.AppendLineFromTag(int(tag), "Value", types.ToHex(encrypted, false))
			return enums.ExceptionServiceErrorNone, nil
		}
		data.SetPosition(data.Position() - 1)
		p := settings.NewAesGcmParameter3(conf, conf.SourceSystemTitle(), conf.Cipher.BlockCipherKey(), conf.Cipher.AuthenticationKey())
		tmp, err = settings.DecryptAesGcm(p, data)
		if err != nil {
			return 0, err
		}
		/*TODO
		if xml == nil && conf.ExpectedInvocationCounter() != 0 {
			if p.InvocationCounter < conf.ExpectedInvocationCounter() {
				return enums.ExceptionServiceErrorInvocationCounterError, nil
			}
			conf.SetExpectedInvocationCounter(1 + p.InvocationCounter)
		}
		*/
		data.Clear()
		err = data.Set(tmp)
		if err != nil {
			return 0, err
		}
		// Update used security to server.
		if conf.IsServer() {
			cipher.SetSecurity(p.Security())
		}
		tag, err = data.Uint8()
		if err != nil {
			return 0, err
		}
	}
	err = parse(initiateRequest, conf, cipher, data, xml, tag)
	if err != nil {
		return 0, err
	}
	return enums.ExceptionServiceErrorNone, nil
}

// parsePDU returns the parse APDU.
func parsePDU(settings *settings.GXDLMSSettings,
	cipher settings.GXICipher,
	buff *types.GXByteBuffer,
	xml *settings.GXDLMSTranslatorStructure) (any, error) {
	// Get AARE tag and length
	tag, err := buff.Uint8()
	if err != nil {
		return nil, err
	}
	if settings.IsServer() {
		if tag != (uint8(constants.BerTypeApplication) | uint8(constants.BerTypeConstructed)) {
			return nil, errors.New("Invalid tag.")
		}
	} else {
		if tag != (uint8(constants.BerTypeApplication) | uint8(constants.BerTypeConstructed) | uint8(internal.PduTypeApplicationContextName)) {
			return nil, errors.New("Invalid tag.")
		}
	}
	len, err := types.GetObjectCount(buff)
	if err != nil {
		return nil, err
	}
	size := buff.Size() - buff.Position()
	if len > size {
		if xml == nil {
			return nil, errors.New("Not enough data.")
		}
		xml.AppendComment("Error: Invalid data size.")
	}
	// Opening tags
	if xml != nil {
		if settings.IsServer() {
			xml.AppendStartTag(int(enums.CommandAarq), "", "", false)
		} else {
			xml.AppendStartTag(int(enums.CommandAare), "", "", false)
		}
	}
	ret, err := parsePDU2(settings, cipher, buff, xml)
	if err != nil {
		return 0, err
	}
	// Closing tags
	if xml != nil {
		if settings.IsServer() {
			xml.AppendEndTag(int(enums.CommandAarq), true)
		} else {
			xml.AppendEndTag(int(enums.CommandAare), true)
		}
	}
	return ret, nil
}

func parsePDU2(settings *settings.GXDLMSSettings, cipher settings.GXICipher,
	buff *types.GXByteBuffer, xml *settings.GXDLMSTranslatorStructure) (any, error) {
	resultComponent := enums.AssociationResultAccepted
	var ret any = 0
	for buff.Position() < buff.Size() {
		tag, err := buff.Uint8()
		if err != nil {
			return ret, err
		}
		switch tag {
		case uint8(constants.BerTypeContext) | uint8(constants.BerTypeConstructed) | uint8(internal.PduTypeApplicationContextName):
			name, err := parseApplicationContextName(settings, buff, xml)
			if err != nil {
				return ret, err
			}
			if name != enums.ApplicationContextNameUnknown {
				if settings.IsServer() {
					return name, nil
				}
				msg := ""
				switch name {
				case enums.ApplicationContextNameLogicalName:
					msg = "\nMeter expects Logical Name referencing."
				case enums.ApplicationContextNameShortName:
					msg = "\nMeter expects Short Name referencing."
				case enums.ApplicationContextNameLogicalNameWithCiphering:
					msg = "\nMeter expects Logical Name referencing with secured connection."
				case enums.ApplicationContextNameShortNameWithCiphering:
					msg = "\nMeter expects Short Name referencing with secured connection."
				}
				return ret, fmt.Errorf("association rejected: %v %v%s",
					enums.AssociationResultPermanentRejected,
					enums.SourceDiagnosticApplicationContextNameNotSupported,
					msg)
			}
		case uint8(constants.BerTypeContext) | uint8(constants.BerTypeConstructed) | uint8(internal.PduTypeCalledApTitle):
			if v, err := buff.Uint8(); err != nil || v != 3 {
				return ret, fmt.Errorf("invalid tag")
			}
			if settings.IsServer() {
				if v, err := buff.Uint8(); err != nil || v != uint8(constants.BerTypeOctetString) {
					return ret, fmt.Errorf("invalid tag")
				}
				length, err := buff.Uint8()
				if err != nil {
					return ret, err
				}
				title := make([]byte, length)
				err = buff.Get(title)
				if err != nil {
					return ret, err
				}
				settings.SetSourceSystemTitle(title)
				if xml != nil {
					xml.AppendLineFromTag(int(internal.TranslatorTagsCalledAPTitle), "Value", types.ToHex(settings.SourceSystemTitle(), false))
				}
			} else {
				if v, err := buff.Uint8(); err != nil || v != uint8(constants.BerTypeInteger) {
					return ret, fmt.Errorf("invalid tag")
				}
				if v, err := buff.Uint8(); err != nil || v != 1 {
					return ret, fmt.Errorf("invalid tag")
				}
				v, err := buff.Uint8()
				if err != nil {
					return ret, err
				}
				resultComponent = enums.AssociationResult(v)
				if xml != nil {
					if resultComponent != enums.AssociationResultAccepted {
						xml.AppendComment(resultComponent.String())
					}
					xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsAssociationResult), "Value",
						xml.IntegerToHex(int(resultComponent), 2, false))
					xml.AppendStartTag(int(enums.TranslatorGeneralTagsResultSourceDiagnostic), "", "", false)
				}
			}
		case uint8(constants.BerTypeContext) | uint8(constants.BerTypeConstructed) | uint8(internal.PduTypeCalledAeQualifier):
			if _, err := buff.Uint8(); err != nil {
				return ret, err
			}
			acseTag, err := buff.Uint8()
			if err != nil {
				return ret, err
			}
			length, err := buff.Uint8()
			if err != nil {
				return ret, err
			}
			tag2, err := buff.Uint8()
			if err != nil {
				return ret, err
			}
			if tag2 == uint8(constants.BerTypeOctetString) {
				calledAEQualifier := make([]byte, length)
				err = buff.Get(calledAEQualifier)
				if err != nil {
					return ret, err
				}
				if xml != nil {
					xml.AppendLineFromTag(int(internal.TranslatorTagsCalledAEQualifier), "Value",
						types.ToHex(calledAEQualifier, false))
				}
			} else {
				if tag2 != uint8(constants.BerTypeInteger) {
					if xml != nil {
						xml.AppendComment(fmt.Sprintf("Invalid tag. %X", tag2))
						continue
					}
					return ret, fmt.Errorf("invalid tag")
				}
				if v, err := buff.Uint8(); err != nil || v != 1 {
					return ret, fmt.Errorf("invalid tag")
				}
				v, err := buff.Uint8()
				if err != nil {
					return ret, err
				}
				if acseTag == 0xA1 {
					ret = enums.SourceDiagnostic(v)
					if xml != nil {
						if ret.(enums.SourceDiagnostic) != enums.SourceDiagnosticNone {
							xml.AppendComment(ret.(enums.SourceDiagnostic).String())
						}
						xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsACSEServiceUser), "Value",
							xml.IntegerToHex(int(ret.(enums.SourceDiagnostic)), 2, false))
					}
				} else {
					ret = enums.AcseServiceProvider(v)
					if xml != nil {
						if ret.(enums.AcseServiceProvider) != enums.AcseServiceProviderNone {
							xml.AppendComment(ret.(enums.AcseServiceProvider).String())
						}
						xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsACSEServiceProvider), "Value",
							xml.IntegerToHex(int(ret.(enums.AcseServiceProvider)), 2, false))
					}
				}
				if xml != nil {
					xml.AppendEndTag(int(enums.TranslatorGeneralTagsResultSourceDiagnostic), false)
				}
			}
		case uint8(constants.BerTypeContext) | uint8(constants.BerTypeConstructed) | uint8(internal.PduTypeCalledApInvocationId):
			if settings.IsServer() {
				if v, err := buff.Uint8(); err != nil || v != 3 {
					return ret, fmt.Errorf("invalid tag")
				}
				if v, err := buff.Uint8(); err != nil || v != uint8(constants.BerTypeInteger) {
					return ret, fmt.Errorf("invalid tag")
				}
				if v, err := buff.Uint8(); err != nil || v != 1 {
					return ret, fmt.Errorf("invalid tag length")
				}
				value, err := buff.Uint8()
				if err != nil {
					return ret, err
				}
				if xml != nil {
					xml.AppendLineFromTag(int(internal.TranslatorTagsCalledAPInvocationId), "Value",
						xml.IntegerToHex(int(value), 2, false))
				}
			} else {
				if v, err := buff.Uint8(); err != nil || v != 0xA {
					return ret, fmt.Errorf("invalid tag")
				}
				if v, err := buff.Uint8(); err != nil || v != uint8(constants.BerTypeOctetString) {
					return ret, fmt.Errorf("invalid tag")
				}
				length, err := buff.Uint8()
				if err != nil {
					return ret, err
				}
				title := make([]byte, length)
				err = buff.Get(title)
				if err != nil {
					return ret, err
				}
				settings.SetSourceSystemTitle(title)
				if xml != nil {
					if xml.Comments {
						xml.AppendComment(internal.SystemTitleToString(
							settings.Standard, settings.SourceSystemTitle(), true))
					}
					xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsRespondingAPTitle), "Value",
						types.ToHex(settings.SourceSystemTitle(), false))
				}
			}
		case uint8(constants.BerTypeContext) | uint8(constants.BerTypeConstructed) | uint8(internal.PduTypeCallingApTitle):
			if _, err := buff.Uint8(); err != nil {
				return ret, err
			}
			if _, err := buff.Uint8(); err != nil {
				return ret, err
			}
			length, err := buff.Uint8()
			if err != nil {
				return ret, err
			}
			title := make([]byte, length)
			err = buff.Get(title)
			if err != nil {
				return ret, err
			}
			settings.SetSourceSystemTitle(title)
			if xml != nil {
				if len(title) != 8 {
					xml.AppendComment("Invalid system title.")
					xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsCallingAPTitle), "Value",
						types.ToHex(title, false))
					if len(title) > 8 {
						settings.SetSourceSystemTitle(title[:8])
					}
				} else {
					xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsCallingAPTitle), "Value",
						types.ToHex(title, false))
				}
			}
		case uint8(constants.BerTypeContext) | uint8(constants.BerTypeConstructed) | uint8(internal.PduTypeSenderAcseRequirements):
			if _, err := buff.Uint8(); err != nil {
				return ret, err
			}
			acseTag, err := buff.Uint8()
			if err != nil {
				return ret, err
			}
			length, err := buff.Uint8()
			if err != nil {
				return ret, err
			}
			challenge := make([]byte, length)
			err = buff.Get(challenge)
			if err != nil {
				return ret, err
			}
			settings.SetStoCChallenge(challenge)
			AppendServerSystemTitleToXml(settings, xml, int(acseTag))
		case uint8(constants.BerTypeContext) | uint8(constants.BerTypeConstructed) | uint8(internal.PduTypeCallingAeInvocationId):
			if _, err := buff.Uint8(); err != nil {
				return ret, err
			}
			if _, err := buff.Uint8(); err != nil {
				return ret, err
			}
			if _, err := buff.Uint8(); err != nil {
				return ret, err
			}
			userID, err := buff.Uint8()
			if err != nil {
				return ret, err
			}
			settings.UserID = int(userID)
			if xml != nil {
				xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsCallingAeInvocationId), "Value",
					xml.IntegerToHex(settings.UserID, 2, false))
			}
		case uint8(constants.BerTypeContext) | uint8(constants.BerTypeConstructed) | uint8(internal.PduTypeCalledAeInvocationId):
			if _, err := types.GetObjectCount(buff); err != nil {
				return ret, err
			}
			tag2, err := buff.Uint8()
			if err != nil {
				return ret, err
			}
			length, err := types.GetObjectCount(buff)
			if err != nil {
				return ret, err
			}
			if tag2 == uint8(constants.BerTypeOctetString) {
				data := make([]byte, length)
				err = buff.Get(data)
				if err != nil {
					return ret, err
				}
				if xml != nil {
					xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsCallingAeQualifier), "Value",
						types.ToHex(data, false))
				}
				if err := handleCertificate(settings, xml, true, data); err != nil {
					if xml == nil {
						return ret, err
					}
					xml.AppendStringLine("Invalid certificate.")
				}
			} else {
				userID, err := buff.Uint8()
				if err != nil {
					return ret, err
				}
				settings.UserID = int(userID)
				if xml != nil {
					xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsCalledAeInvocationId), "Value",
						xml.IntegerToHex(settings.UserID, 2, false))
				}
			}
		case uint8(constants.BerTypeContext) | uint8(constants.BerTypeConstructed) | uint8(internal.PduTypeCallingAeQualifier):
			if _, err := types.GetObjectCount(buff); err != nil {
				return ret, err
			}
			tag2, err := buff.Uint8()
			if err != nil {
				return ret, err
			}
			length, err := types.GetObjectCount(buff)
			if err != nil {
				return ret, err
			}
			if tag2 == uint8(constants.BerTypeOctetString) {
				data := make([]byte, length)
				err = buff.Get(data)
				if err != nil {
					return ret, err
				}
				if xml != nil {
					xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsCallingAeQualifier), "Value",
						types.ToHex(data, false))
				}
				if err := handleCertificate(settings, xml, false, data); err != nil {
					if xml == nil {
						return ret, err
					}
					xml.AppendStringLine("Invalid certificate.")
				}
			} else {
				userID, err := buff.Uint8()
				if err != nil {
					return ret, err
				}
				settings.UserID = int(userID)
				if xml != nil {
					if settings.IsServer() {
						xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsCallingAeQualifier), "Value",
							xml.IntegerToHex(settings.UserID, 2, false))
					} else {
						xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsRespondingAeInvocationId), "Value",
							xml.IntegerToHex(settings.UserID, 2, false))
					}
				}
			}
		case uint8(constants.BerTypeContext) | uint8(constants.BerTypeConstructed) | uint8(internal.PduTypeCallingApInvocationId):
			if v, err := buff.Uint8(); err != nil || v != 3 {
				return ret, fmt.Errorf("invalid tag")
			}
			if v, err := buff.Uint8(); err != nil || v != 2 {
				return ret, fmt.Errorf("invalid length")
			}
			if v, err := buff.Uint8(); err != nil || v != 1 {
				return ret, fmt.Errorf("invalid tag length")
			}
			value, err := buff.Uint8()
			if err != nil {
				return ret, err
			}
			if xml != nil {
				xml.AppendLineFromTag(int(internal.TranslatorTagsCallingApInvocationId), "Value",
					xml.IntegerToHex(int(value), 2, false))
			}
		case uint8(constants.BerTypeContext) | uint8(internal.PduTypeSenderAcseRequirements),
			uint8(constants.BerTypeContext) | uint8(internal.PduTypeCallingApInvocationId):
			if v, err := buff.Uint8(); err != nil || v != 2 {
				return ret, fmt.Errorf("invalid tag")
			}
			if v, err := buff.Uint8(); err != nil || v != uint8(constants.BerTypeObjectDescriptor) {
				return ret, fmt.Errorf("invalid tag")
			}
			if v, err := buff.Uint8(); err != nil || v != 0x80 {
				return ret, fmt.Errorf("invalid tag")
			}
			if xml != nil {
				xml.AppendLineFromTag(int(tag), "Value", "1")
			}
		case uint8(constants.BerTypeContext) | uint8(internal.PduTypeMechanismName), //0x8B
			uint8(constants.BerTypeContext) | uint8(internal.PduTypeCallingAeInvocationId): //0x89
			if err := UpdateAuthentication(settings, buff); err != nil {
				return ret, err
			}
			if xml != nil {
				if xml.OutputType() == enums.TranslatorOutputTypeSimpleXML {
					xml.AppendLineFromTag(int(tag), "Value", settings.Authentication.String())
				} else {
					xml.AppendLineFromTag(int(tag), "Value", fmt.Sprintf("%d", settings.Authentication))
				}
			}
		case uint8(constants.BerTypeContext) | uint8(constants.BerTypeConstructed) | uint8(internal.PduTypeCallingAuthenticationValue):
			if err := UpdatePassword(settings, buff, xml); err != nil {
				return ret, err
			}
		case uint8(constants.BerTypeContext) | uint8(constants.BerTypeConstructed) | uint8(internal.PduTypeUserInformation):
			serviceErr, err := ParseUserInformation(settings, cipher, buff, xml)
			if err != nil {
				if xml == nil {
					return ret, err
				}
				if ex, ok := err.(*GXDLMSExceptionResponse); ok &&
					ex.ExceptionServiceError() == enums.ExceptionServiceErrorInvocationCounterError {
					return enums.ExceptionServiceErrorInvocationCounterError, nil
				}
				return enums.ExceptionServiceErrorDecipheringError, nil
			}
			if serviceErr == enums.ExceptionServiceErrorInvocationCounterError {
				return enums.ExceptionServiceErrorInvocationCounterError, nil
			}
		case uint8(constants.BerTypeContext):
			tmp, err := ParseProtocolVersion(settings, buff, xml)
			if err != nil {
				return ret, err
			}
			if tmp != enums.AcseServiceProviderNone {
				resultComponent = enums.AssociationResultPermanentRejected
			}
			ret = tmp
		default:
			if xml != nil {
				xml.AppendComment(fmt.Sprintf("Unknown tag: %d.", tag))
			}
			if buff.Position() < buff.Size() {
				length, err := buff.Uint8()
				if err != nil {
					return ret, err
				}
				if err := buff.SetPosition(buff.Position() + int(length)); err != nil {
					return ret, err
				}
			}
		}
	}
	if !settings.IsServer() && xml == nil &&
		resultComponent != enums.AssociationResultAccepted && ret != 0 {
		return ret, fmt.Errorf("association rejected: %v %v", resultComponent, ret)
	}
	return ret, nil
}

func handleCertificate(settings *settings.GXDLMSSettings,
	xml *settings.GXDLMSTranslatorStructure,
	isServer bool,
	data []byte) error {
	cert, err := types.NewGXx509Certificate(data)
	if err != nil {
		return err
	}
	if isServer {
		settings.ServerPublicKeyCertificate = cert
	} else {
		settings.ClientPublicKeyCertificate = cert
	}
	if settings.Cipher != nil {
		if cert.KeyUsage&enums.KeyUsageKeyCertSign != 0 {
			kp := settings.Cipher.KeyAgreementKeyPair()
			var priv *types.GXPrivateKey
			if kp != nil {
				priv = kp.Value
			}
			settings.Cipher.SetKeyAgreementKeyPair(
				types.NewGXKeyValuePair(cert.PublicKey, priv))
		}
		if cert.KeyUsage&enums.KeyUsageDigitalSignature != 0 {
			kp := settings.Cipher.SigningKeyPair()
			var priv *types.GXPrivateKey
			if kp != nil {
				priv = kp.Value
			}
			settings.Cipher.SetSigningKeyPair(
				types.NewGXKeyValuePair(cert.PublicKey, priv))
		}
	}
	if xml != nil && xml.Comments {
		xml.AppendComment(cert.String())
	}
	return nil
}

func getUserInformation(conf *settings.GXDLMSSettings, cipher settings.GXICipher) ([]byte, error) {
	data := types.GXByteBuffer{}
	err := data.SetUint8(enums.CommandInitiateResponse)
	if err != nil {
		return nil, err
	}
	if conf.QualityOfService == 0 {
		err = data.SetUint8(0x00)
		if err != nil {
			return nil, err
		}
	} else {
		err = data.SetUint8(1)
		if err != nil {
			return nil, err
		}
		err = data.SetUint8(conf.QualityOfService)
		if err != nil {
			return nil, err
		}
	}
	err = data.SetUint8(06)
	if err != nil {
		return nil, err
	}
	err = data.SetUint8(0x5F)
	if err != nil {
		return nil, err
	}
	err = data.SetUint8(0x1F)
	if err != nil {
		return nil, err
	}
	err = data.SetUint8(0x04)
	if err != nil {
		return nil, err
	}
	err = data.SetUint8(0x00)
	if err != nil {
		return nil, err
	}
	bs, err := types.NewGXBitStringFromInteger(int(conf.NegotiatedConformance), 24)
	if err != nil {
		return nil, err
	}
	err = data.Set(bs.Value())
	if err != nil {
		return nil, err
	}
	err = data.SetUint16(conf.MaxPduSize())
	if err != nil {
		return nil, err
	}
	// VAA Name VAA name (0x0007 for LN referencing and 0xFA00 for SN)
	if conf.UseLogicalNameReferencing() {
		err = data.SetUint16(0x0007)
		if err != nil {
			return nil, err
		}
	} else {
		err = data.SetUint16(0xFA00)
		if err != nil {
			return nil, err
		}
	}
	if conf.IsCiphered(false) {
		var cmd byte
		if (conf.NegotiatedConformance & enums.ConformanceGeneralProtection) != 0 {
			if conf.Cipher.DedicatedKey() != nil && len(conf.Cipher.DedicatedKey()) != 0 {
				cmd = byte(enums.CommandGeneralDedCiphering)
			} else {
				cmd = byte(enums.CommandGeneralGloCiphering)
			}
		} else {
			cmd = byte(enums.CommandGloInitiateResponse)
		}
		p := settings.NewAesGcmParameter(cmd, conf,
			cipher.Security(), cipher.SecuritySuite(),
			uint64(cipher.InvocationCounter()), cipher.SystemTitle(),
			cipher.BlockCipherKey(), cipher.AuthenticationKey())
		tmp, err := settings.EncryptAesGcm(p, data.Array())
		if err != nil {
			return nil, err
		}
		cipher.SetInvocationCounter(cipher.InvocationCounter() + 1)
		return tmp, nil
	}
	return data.Array(), nil
}

// GenerateAARE returns the server generates AARE message.
func GenerateAARE(conf *settings.GXDLMSSettings,
	data *types.GXByteBuffer,
	result enums.AssociationResult,
	diagnostic any,
	cipher settings.GXICipher,
	errorData *types.GXByteBuffer,
	encryptedData *types.GXByteBuffer) error {
	offset := data.Size()
	err := data.SetUint8((uint8(constants.BerTypeApplication) | uint8(constants.BerTypeConstructed) | uint8(internal.PduTypeApplicationContextName)))
	if err != nil {
		return err
	}
	err = GenerateApplicationContextName(conf, data, cipher)
	if err != nil {
		return err
	}
	err = data.SetUint8(uint8(constants.BerTypeContext) | uint8(constants.BerTypeConstructed) | uint8(constants.BerTypeInteger))
	if err != nil {
		return err
	}
	err = data.SetUint8(3)
	if err != nil {
		return err
	}
	err = data.SetUint8(uint8(constants.BerTypeInteger))
	if err != nil {
		return err
	}
	err = data.SetUint8(1)
	if err != nil {
		return err
	}
	err = data.SetUint8(uint8(result))
	if err != nil {
		return err
	}
	err = data.SetUint8(0xA3)
	if err != nil {
		return err
	}
	err = data.SetUint8(5)
	if err != nil {
		return err
	}
	// Tag
	if _, ok := diagnostic.(enums.SourceDiagnostic); ok {
		err = data.SetUint8(0xA1)
		if err != nil {
			return err
		}
	} else {
		err = data.SetUint8(0xA2)
		if err != nil {
			return err
		}
	}
	err = data.SetUint8(3)
	if err != nil {
		return err
	}
	err = data.SetUint8(2)
	if err != nil {
		return err
	}
	err = data.SetUint8(1)
	if err != nil {
		return err
	}
	err = data.SetUint8(diagnostic.(byte))
	if err != nil {
		return err
	}
	// SystemTitle
	if conf.IsCiphered(false) ||
		conf.Authentication == enums.AuthenticationHighGMAC ||
		conf.Authentication == enums.AuthenticationHighSHA256 ||
		conf.Authentication == enums.AuthenticationHighECDSA {
		err = data.SetUint8(uint8(constants.BerTypeContext) | uint8(constants.BerTypeConstructed) | uint8(internal.PduTypeCalledApInvocationId))
		if err != nil {
			return err
		}
		err = data.SetUint8(uint8((2 + len(cipher.SystemTitle()))))
		if err != nil {
			return err
		}
		err = data.SetUint8(uint8(constants.BerTypeOctetString))
		if err != nil {
			return err
		}
		err = data.SetUint8(uint8(len(cipher.SystemTitle())))
		if err != nil {
			return err
		}
		err = data.Set(cipher.SystemTitle())
		if err != nil {
			return err
		}
	}
	// Add CallingAeQualifier.
	if conf.Authentication == enums.AuthenticationHighECDSA && conf.ServerPublicKeyCertificate != nil {
		raw := conf.ServerPublicKeyCertificate.RawData()
		err = data.SetUint8(uint8(constants.BerTypeContext) | uint8(constants.BerTypeConstructed) | uint8(internal.PduTypeCallingAeQualifier))
		if err != nil {
			return err
		}
		err = types.SetObjectCount(4+len(raw), data)
		if err != nil {
			return err
		}
		err = data.SetUint8(uint8(constants.BerTypeOctetString))
		if err != nil {
			return err
		}
		err = types.SetObjectCount(len(raw), data)
		if err != nil {
			return err
		}
		err = data.Set(raw)
	} else if conf.UserID != -1 {
		err = data.SetUint8(uint8(constants.BerTypeContext) | uint8(constants.BerTypeConstructed) | uint8(internal.PduTypeCallingAeQualifier))
		if err != nil {
			return err
		}
		err = data.SetUint8(3)
		if err != nil {
			return err
		}
		err = data.SetUint8(uint8(constants.BerTypeInteger))
		if err != nil {
			return err
		}
		err = data.SetUint8(1)
		if err != nil {
			return err
		}
		err = data.SetUint8(uint8(conf.UserID))
		if err != nil {
			return err
		}
	}
	if conf.Authentication > enums.AuthenticationLow {
		err = data.SetUint8(0x88)
		if err != nil {
			return err
		}
		err = data.SetUint8(0x02)
		if err != nil {
			return err
		}
		err = data.SetUint16(0x0780)
		if err != nil {
			return err
		}
		err = data.SetUint8(0x89)
		if err != nil {
			return err
		}
		err = data.SetUint8(0x07)
		if err != nil {
			return err
		}
		err = data.SetUint8(0x60)
		if err != nil {
			return err
		}
		err = data.SetUint8(0x85)
		if err != nil {
			return err
		}
		err = data.SetUint8(0x74)
		if err != nil {
			return err
		}
		err = data.SetUint8(0x05)
		if err != nil {
			return err
		}
		err = data.SetUint8(0x08)
		if err != nil {
			return err
		}
		err = data.SetUint8(0x02)
		if err != nil {
			return err
		}
		err = data.SetUint8(uint8(conf.Authentication))
		if err != nil {
			return err
		}
		err = data.SetUint8(0xAA)
		if err != nil {
			return err
		}
		err = data.SetUint8(uint8((2 + len(conf.StoCChallenge()))))
		if err != nil {
			return err
		}
		err = data.SetUint8(uint8(constants.BerTypeContext))
		if err != nil {
			return err
		}
		err = data.SetUint8(uint8(len(conf.StoCChallenge())))
		if err != nil {
			return err
		}
		err = data.Set(conf.StoCChallenge())
		if err != nil {
			return err
		}
	}
	if result == enums.AssociationResultAccepted || cipher == nil || cipher.Security() == enums.SecurityNone {
		var tmp []byte
		err = data.SetUint8(uint8(constants.BerTypeContext) | uint8(constants.BerTypeConstructed) | uint8(internal.PduTypeUserInformation))
		if err != nil {
			return err
		}
		if encryptedData != nil && encryptedData.Size() != 0 {
			tmp2 := types.NewGXByteBufferWithCapacity(2 + encryptedData.Size())
			err = tmp2.SetUint8(uint8(enums.CommandGloInitiateResponse))
			if err != nil {
				return err
			}
			types.SetObjectCount(encryptedData.Size(), tmp2)
			err = tmp2.SetByteBuffer(encryptedData)
			if err != nil {
				return err
			}
			tmp = tmp2.Array()
		} else {
			if errorData != nil && errorData.Size() != 0 {
				tmp = errorData.Array()
			} else {
				tmp, err = getUserInformation(conf, cipher)
			}
		}
		err = data.SetUint8(uint8((2 + len(tmp))))
		if err != nil {
			return err
		}
		err = data.SetUint8(uint8(constants.BerTypeOctetString))
		if err != nil {
			return err
		}
		err = data.SetUint8(uint8(len(tmp)))
		if err != nil {
			return err
		}
		err = data.Set(tmp)
		if err != nil {
			return err
		}
	}
	types.InsertObjectCount(data.Size()-offset-1, data, offset+1)
	if conf.Gateway != nil && conf.Gateway.PhysicalDeviceAddress != nil {
		tmp := types.GXByteBuffer{}
		err = tmp.SetUint8(enums.CommandGatewayResponse)
		if err != nil {
			return err
		}
		err = tmp.SetUint8(conf.Gateway.NetworkID)
		if err != nil {
			return err
		}
		err = tmp.SetUint8(uint8(len(conf.Gateway.PhysicalDeviceAddress)))
		err = tmp.Set(conf.Gateway.PhysicalDeviceAddress)
		if err != nil {
			return err
		}
		err = data.SetAt(tmp.Array(), 0, tmp.Size())
		if err != nil {
			return err
		}
	}
	return err
}

// ParseUserInformation returns the parse User Information from PDU.
func ParseUserInformation(settings *settings.GXDLMSSettings, cipher settings.GXICipher,
	data *types.GXByteBuffer, xml *settings.GXDLMSTranslatorStructure) (enums.ExceptionServiceError, error) {
	length, err := data.Uint8()
	if err != nil {
		return enums.ExceptionServiceErrorNone, err
	}
	if data.Size()-data.Position() < int(length) {
		if xml == nil {
			return enums.ExceptionServiceErrorNone, fmt.Errorf("not enough data")
		}
		xml.AppendComment("Error: Invalid data size.")
	}
	// Excoding the choice for user information
	tag, err := data.Uint8()
	if err != nil {
		return enums.ExceptionServiceErrorNone, err
	}
	if tag != 0x4 {
		return enums.ExceptionServiceErrorNone, fmt.Errorf("invalid tag")
	}
	length, err = data.Uint8()
	if err != nil {
		return enums.ExceptionServiceErrorNone, err
	}
	if data.Size()-data.Position() < int(length) {
		if xml == nil {
			return enums.ExceptionServiceErrorNone, fmt.Errorf("not enough data")
		}
		xml.AppendComment("Error: Invalid data size.")
	}
	if xml != nil && xml.OutputType() == enums.TranslatorOutputTypeStandardXML {
		xml.AppendLineFromTag(int(enums.TranslatorGeneralTagsUserInformation), "",
			types.ToHexWithRange(data.Array(), false, data.Position(), int(length)))
		if err := data.SetPosition(data.Position() + int(length)); err != nil {
			return enums.ExceptionServiceErrorNone, err
		}
		return enums.ExceptionServiceErrorNone, nil
	}
	return parseInitiate(false, settings, cipher, data, xml)
}
