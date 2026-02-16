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
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/manufacturersettings"
)

// GXDLMSObject provides an base class for DLMS COSEM objects.
type GXDLMSObject struct {
	parent GXDLMSObjectCollection

	// Interface type of the DLMS object.
	objectType enums.ObjectType

	// Gets or sets the object that contains data about the control.
	Tag any

	// DLMS version number.
	Version uint8

	// The base name of the object, if using SN.
	ShortName int16

	// Logical Name of DLMS object.
	logicalName string

	// Description of DLMS object.
	Description string

	// Object attribute collection.
	Attributes manufacturersettings.GXAttributeCollection

	// object attribute collection.
	MethodAttributes manufacturersettings.GXAttributeCollection

	// Read times of attributes.
	readTimes map[int]time.Time
}

func getObjectCollection(objects interface{}) *GXDLMSObjectCollection {
	ret := objects.(GXDLMSObjectCollection)
	return &ret
}

// Interface type of the DLMS object.
func (g *GXDLMSObject) ObjectType() enums.ObjectType {
	return g.objectType
}

// LogicalName returns the logical name of DLMS object.
func (g *GXDLMSObject) LogicalName() string {
	return g.logicalName
}

// SetLogicalName sets the logical name of DLMS object.
func (g *GXDLMSObject) SetLogicalName(ln string) error {
	err := ValidateLogicalName(ln)
	if err != nil {
		return err
	}
	g.logicalName = ln
	return nil
}

// Name returns the logical or Short Name of DLMS object.
//
//	Returns:
func (g *GXDLMSObject) Name() any {
	if g.ShortName != 0 {
		return g.ShortName
	}
	return g.logicalName
}

func (g *GXDLMSObject) GetAttribute(index int) *manufacturersettings.GXDLMSAttributeSettings {
	att := g.Attributes.Find(index)
	if att == nil {
		att = &manufacturersettings.GXDLMSAttributeSettings{Index: index}
		g.Attributes = append(g.Attributes, att)
		// LN is read only.
		if index == 1 {
			att.Access = enums.AccessModeRead
		}
	}
	return att
}

// IsRead returns the is attribute read.
// index: Attribute index to read.
//
//	Returns:
//	    Returns true if attribute is read.
func (g *GXDLMSObject) IsRead(index int) bool {
	if !g.CanRead(index) {
		return true
	}
	return g.GetLastReadTime(index).IsZero()
}

// GetLastReadTime returns the last read time of attribute.
// attributeIndex: Attribute index.
//
//	Returns:
//	    Last read time.
func (g *GXDLMSObject) GetLastReadTime(attributeIndex int) time.Time {
	return g.readTimes[attributeIndex]
}

// SetLastReadTime sets the last read time of attribute.
// attributeIndex: Attribute index.
// tm: Read time.
func (g *GXDLMSObject) SetLastReadTime(attributeIndex int, tm time.Time) {
	g.readTimes[attributeIndex] = tm
}

// CanRead returns the is attribute of the object readable.
// index: Attribute index of the object.
//
//	Returns:
//	    True, if attribute of the object is readable.
func (g *GXDLMSObject) CanRead(index int) bool {
	// Association version number is not known and for this reason all access levels must be checked.
	var access = g.GetAccess(index)
	return access&enums.AccessModeRead != 0 || access == enums.AccessModeAuthenticatedRead ||
		access == enums.AccessModeAuthenticatedReadWrite || g.GetAccess3(index)&enums.AccessMode3Read != 0
}

// ValidateLogicalName returns the validate logical name.
// ln: Logical name to validate.
//
//	Returns:
//	    Error if logical name is invalid.
func ValidateLogicalName(ln string) error {
	var items = strings.Split(ln, ".")
	if len(items) != 6 {
		return errors.New("Invalid Logical Name.")
	}
	return nil
}

// String produces logical or Short Name of DLMS object.
//
//	Returns:
func (g *GXDLMSObject) String() string {
	if g.ShortName != 0 {
		return fmt.Sprint(g.ShortName) + " " + g.Description
	}
	return g.logicalName + " " + g.Description
}

// GetAccess returns the returns attribute access mode.index: Attribute index.
//
//	Returns:
//	    Is attribute read only.
func (g *GXDLMSObject) GetAccess(index int) enums.AccessMode {
	var att = g.Attributes.Find(index)
	if att == nil {
		return enums.AccessModeReadWrite
	}
	return att.Access
}

// SetAccess returns the set attribute access.index:
// access:
func (g *GXDLMSObject) SetAccess(index int, access enums.AccessMode) {
	var att = g.Attributes.Find(index)
	if att == nil {
		att = &manufacturersettings.GXDLMSAttributeSettings{Index: index}
		g.Attributes = append(g.Attributes, att)
	}
	att.Access = access
	att.Access3 = enums.AccessMode3NoAccess
}

// GetAccess3 returns the returns is attribute read only.index: Attribute index.
//
//	Returns:
//	    Is attribute read only.
func (g *GXDLMSObject) GetAccess3(index int) enums.AccessMode3 {
	var att = g.Attributes.Find(index)
	if att == nil {
		return enums.AccessMode3Read
	}
	return att.Access3
}

// SetAccess3 returns the set attribute access.
// index:
// access:
func (g *GXDLMSObject) SetAccess3(index int, access enums.AccessMode3) {
	var att = g.Attributes.Find(index)
	if att == nil {
		att = &manufacturersettings.GXDLMSAttributeSettings{Index: index}
		g.Attributes = append(g.Attributes, att)
	}
	att.Access = enums.AccessModeNoAccess
	att.Access3 = access
}

// GetMethodAccess returns the returns is Method attribute read only.index: Method Attribute index.
//
//	Returns:
//	    Is attribute read only.
func (g *GXDLMSObject) GetMethodAccess(index int) enums.MethodAccessMode {
	att := g.MethodAttributes.Find(index)
	if att != nil {
		return att.MethodAccess
	}
	return enums.MethodAccessModeAccess
}

// SetMethodAccess returns the set Method attribute access.
// index: Method Attribute index.
// access: Method access mode.
func (g *GXDLMSObject) SetMethodAccess(index int, access enums.MethodAccessMode) error {
	if int(access) > int(enums.MethodAccessModeAuthenticatedAccess) {
		return fmt.Errorf("access")
	}
	var att = g.MethodAttributes.Find(index)
	if att == nil {
		att = &manufacturersettings.GXDLMSAttributeSettings{Index: index}
		g.MethodAttributes = append(g.MethodAttributes, att)
	}
	att.MethodAccess = access
	att.MethodAccess3 = enums.MethodAccessMode3NoAccess
	return nil
}

// GetMethodAccess3 returns the returns is Method attribute read only.
// index: Method Attribute index.
//
//	Returns:
//	    Is attribute read only.
func (g *GXDLMSObject) GetMethodAccess3(index int) enums.MethodAccessMode3 {
	var att = g.MethodAttributes.Find(index)
	if att != nil {
		return att.MethodAccess3
	}
	return enums.MethodAccessMode3Access
}

// SetMethodAccess3 returns the set Method attribute access.index:
// access:
func (g *GXDLMSObject) SetMethodAccess3(index int, access enums.MethodAccessMode3) {
	var att = g.MethodAttributes.Find(index)
	if att == nil {
		att = &manufacturersettings.GXDLMSAttributeSettings{Index: index}
		g.MethodAttributes = append(g.MethodAttributes, att)
	}
	att.MethodAccess3 = access
	att.MethodAccess = enums.MethodAccessModeNoAccess
}

// GetUIDataType returns UI data type of selected index.
// index: Attribute index of the object.
//
//	Returns:
//	    UseUI data type of the object.
func (g *GXDLMSObject) GetUIDataType(index int) enums.DataType {
	var att = g.GetAttribute(index)
	return att.UIType
}

// GetDataType returns the returns data type of selected index.
// index: Attribute index of the object.
//
//	Returns:
//	    Device data type of the object.
func (g *GXDLMSObject) GetDataType(index int) (enums.DataType, error) {
	var att = g.GetAttribute(index)
	return att.Type(), nil
}

// GetValues returns the returns attributes as an array.
//
//	Returns:
//	    Collection of COSEM object values.
func (g *GXDLMSObject) GetValues() (any, error) {
	panic("GetValues: not implemented")
}

func (g *GXDLMSObject) SetDataType(index int, dt enums.DataType) {
	var att = g.GetAttribute(index)
	att.SetType(dt)
}

func (g *GXDLMSObject) SetUIDataType(index int, dt enums.DataType) {
	var att = g.GetAttribute(index)
	att.UIType = dt
}

func (g *GXDLMSObject) SetStatic(index int, isStatic bool) {
	var att = g.GetAttribute(index)
	att.Static = isStatic
}

func (g *GXDLMSObject) GetStatic(index int) bool {
	var att = g.GetAttribute(index)
	return att.Static
}

// GetAccessSelector returns the returns is Method attribute read only.
// index: Method Attribute index.
//
//	Returns:
//	    Access selector value.
func (g *GXDLMSObject) GetAccessSelector(index int) uint8 {
	var att = g.GetAttribute(index)
	return att.AccessSelector
}

// SetAccessSelector returns the set Method attribute access.
// index: Attribute inde
// value: Access selector value.
func (g *GXDLMSObject) SetAccessSelector(index int, value uint8) {
	var att = g.GetAttribute(index)
	att.AccessSelector = value
}
