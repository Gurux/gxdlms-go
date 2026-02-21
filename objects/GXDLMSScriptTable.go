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
	"reflect"

	"github.com/Gurux/gxdlms-go/dlmserrors"
	"github.com/Gurux/gxdlms-go/enums"
	"github.com/Gurux/gxdlms-go/internal"
	"github.com/Gurux/gxdlms-go/internal/helpers"
	"github.com/Gurux/gxdlms-go/settings"
	"github.com/Gurux/gxdlms-go/types"
)

// Script table objects contain a table of script entries. Each entry consists of a script identifier
// and a series of action specifications.
// Online help:
// https://www.gurux.fi/Gurux.DLMS.Objects.GXDLMSScriptTable
type GXDLMSScriptTable struct {
	GXDLMSObject
	Scripts []GXDLMSScript
}

// Base returns the base GXDLMSObject of the object.
func (g *GXDLMSScriptTable) Base() *GXDLMSObject {
	return &g.GXDLMSObject
}

// Invoke returns the invokes method.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Invoke parameters.
func (g *GXDLMSScriptTable) Invoke(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) ([]byte, error) {
	if e.Index != 1 {
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return nil, nil
}

// GetAttributeIndexToRead returns the collection of attributes to read.
// If attribute is static and already read or device is returned HW error it is not returned.
//
// Parameters:
//
//	all: All items are returned even if they are read already.
//
// Returns:
//
//	Collection of attributes to read.
func (g *GXDLMSScriptTable) GetAttributeIndexToRead(all bool) []int {
	var attributes []int
	// LN is static and read only once.
	if all || g.LogicalName() == "" {
		attributes = append(attributes, 1)
	}
	// Scripts
	if all || !g.CanRead(2) {
		attributes = append(attributes, 2)
	}
	return attributes
}

// GetNames returns the names of attribute indexes.
func (g *GXDLMSScriptTable) GetNames() []string {
	return []string{"Logical Name", "Scripts"}
}

// GetMethodNames returns the names of method indexes.
func (g *GXDLMSScriptTable) GetMethodNames() []string {
	return []string{"Execute"}
}

// GetAttributeCount returns the amount of attributes.
//
// Returns:
//
//	Count of attributes.
func (g *GXDLMSScriptTable) GetAttributeCount() int {
	return 2
}

// GetMethodCount returns the amount of methods.
func (g *GXDLMSScriptTable) GetMethodCount() int {
	return 1
}

// GetValue returns the value of given attribute.
// When raw parameter us not used example register multiplies value by scalar.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Get parameters.
//
// Returns:
//
//	Value of the attribute index.
func (g *GXDLMSScriptTable) GetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) (any, error) {
	var err error
	if e.Index == 1 {
		v, err := helpers.LogicalNameToBytes(g.LogicalName())
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return v, err
	}
	if e.Index == 2 {
		cnt := len(g.Scripts)
		data := types.NewGXByteBuffer()
		err = data.SetUint8(uint8(enums.DataTypeArray))
		if err != nil {
			return nil, err
		}
		err = types.SetObjectCount(cnt, data)
		for _, it := range g.Scripts {
			err = data.SetUint8(uint8(enums.DataTypeStructure))
			if err != nil {
				return nil, err
			}
			err = data.SetUint8(2)
			if err != nil {
				return nil, err
			}
			err = internal.SetData(settings, data, enums.DataTypeUint16, it.Id)
			if err != nil {
				return nil, err
			}
			err = data.SetUint8(uint8(enums.DataTypeArray))
			if err != nil {
				return nil, err
			}
			err = data.SetUint8(uint8(len(it.Actions)))
			if err != nil {
				return nil, err
			}
			for _, a := range it.Actions {
				err = data.SetUint8(uint8(enums.DataTypeStructure))
				if err != nil {
					return nil, err
				}
				err = data.SetUint8(5)
				if err != nil {
					return nil, err
				}
				err = internal.SetData(settings, data, enums.DataTypeEnum, a.Type)
				if err != nil {
					return nil, err
				}
				err = internal.SetData(settings, data, enums.DataTypeUint16, a.Target.Base().ObjectType())
				if err != nil {
					return nil, err
				}
				ln, err := helpers.LogicalNameToBytes(a.Target.Base().LogicalName())
				if err != nil {
					return nil, err
				}
				err = internal.SetData(settings, data, enums.DataTypeOctetString, ln)
				if err != nil {
					return nil, err
				}
				err = internal.SetData(settings, data, enums.DataTypeInt8, a.Index)
				if err != nil {
					return nil, err
				}
				// parameter
				tp := a.ParameterDataType
				if tp == enums.DataTypeNone {
					tp, err = internal.GetDLMSDataType(reflect.TypeOf(a.Parameter))
					if err != nil {
						return nil, err
					}
				}
				err = internal.SetData(settings, data, tp, a.Parameter)
				if err != nil {
					return nil, err
				}
			}
		}
		return data.Array(), nil
	}
	e.Error = enums.ErrorCodeReadWriteDenied
	return nil, nil
}

// SetValue returns the set value of given attribute.
// When raw parameter us not used example register multiplies value by scalar.
//
// Parameters:
//
//	settings: DLMS settings.
//	e: Set parameters.
func (g *GXDLMSScriptTable) SetValue(settings *settings.GXDLMSSettings, e *internal.ValueEventArgs) error {
	var err error
	if e.Index == 1 {
		ln, err := helpers.ToLogicalName(e.Value)
		if err != nil {
			e.Error = enums.ErrorCodeReadWriteDenied
		}
		return g.SetLogicalName(ln)
	} else if e.Index == 2 {
		g.Scripts = g.Scripts[:0]
		for _, tmp := range e.Value.(types.GXArray) {
			item := tmp.(types.GXStructure)
			script := GXDLMSScript{}
			script.Id = item[0].(uint16)
			g.Scripts = append(g.Scripts, script)
			for _, tmp2 := range item[1].(types.GXArray) {
				arr := tmp2.(types.GXStructure)
				it := GXDLMSScriptAction{}
				it.Type = enums.ScriptActionType(arr[0].(types.GXEnum).Value)
				ot := enums.ObjectType(arr[1].(uint16))
				ln, err := helpers.ToLogicalName(arr[2])
				if err != nil {
					return err
				}
				it.Target = getObjectCollection(settings.Objects).FindByLN(ot, ln)
				if it.Target == nil {
					it.Target, err = CreateObject(ot, ln, 0)
					if err != nil {
						return err
					}
				}
				it.Index = arr[3].(int8)
				it.Parameter = arr[4]
				if it.Parameter != nil {
					it.ParameterDataType, err = internal.GetDLMSDataType(reflect.TypeOf(it.Parameter))
					if err != nil {
						break
					}
				}
				script.Actions = append(script.Actions, it)
			}
		}
	} else {
		e.Error = enums.ErrorCodeReadWriteDenied
	}
	return err
}

// Load returns the load object content from XML.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSScriptTable) Load(reader *GXXmlReader) error {
	g.Scripts = g.Scripts[:0]
	var err error
	if ret, err := reader.IsStartElementNamed("Scripts", true); ret && err == nil {
		for {
			ret, err = reader.IsStartElementNamed("Script", true)
			if err != nil {
				return err
			}
			if !ret {
				break
			}
			it := GXDLMSScript{}
			g.Scripts = append(g.Scripts, it)
			ret, err := reader.ReadElementContentAsInt("ID", 0)
			if err != nil {
				return err
			}
			it.Id = uint16(ret)
			if ret, err := reader.IsStartElementNamed("Actions", true); ret && err == nil {
				for {
					ret, err = reader.IsStartElementNamed("Action", true)
					if err != nil {
						return err
					}
					if !ret {
						break
					}
					a := GXDLMSScriptAction{}
					ret, err := reader.ReadElementContentAsInt("Type", 0)
					if err != nil {
						return err
					}
					a.Type = enums.ScriptActionType(ret)
					ret, err = reader.ReadElementContentAsInt("ObjectType", 0)
					if err != nil {
						return err
					}
					ot := enums.ObjectType(ret)
					ln, err := reader.ReadElementContentAsString("LN", "")
					if err != nil {
						return err
					}
					ret, err = reader.ReadElementContentAsInt("Index", 0)
					if err != nil {
						return err
					}
					a.Index = int8(ret)
					a.Target = reader.Objects.FindByLN(ot, ln)
					if a.Target == nil {
						a.Target, err = CreateObject(ot, ln, 0)
						if err != nil {
							return err
						}
					}
					ret, err = reader.ReadElementContentAsInt("ParameterDataType", 0)
					if err != nil {
						return err
					}
					a.ParameterDataType = enums.DataType(ret)
					a.Parameter, err = reader.ReadElementContentAsString("Parameter", "")
					if err != nil {
						return err
					}
					if a.ParameterDataType != enums.DataTypeNone {
						//TODO: a.Parameter = ChangeType(a.Parameter, a.ParameterDataType)
					}
					it.Actions = append(it.Actions, a)
				}
				reader.ReadEndElement("Actions")
			}
		}
		reader.ReadEndElement("Scripts")
	}
	return err
}

// Save returns the save object content to XML.
//
// Parameters:
//
//	writer: XML writer.
func (g *GXDLMSScriptTable) Save(writer *GXXmlWriter) error {
	var err error
	if g.Scripts != nil {
		writer.WriteStartElement("Scripts")
		for _, it := range g.Scripts {
			writer.WriteStartElement("Script")
			err = writer.WriteElementString("ID", fmt.Sprint())
			if err != nil {
				return err
			}
			writer.WriteStartElement("Actions")
			for _, a := range it.Actions {
				writer.WriteStartElement("Action")
				err = writer.WriteElementString("Type", int(a.Type))
				if err != nil {
					return err
				}
				if a.Target == nil {
					err = writer.WriteElementString("ObjectType", int(enums.ObjectTypeNone))
					if err != nil {
						return err
					}
					err = writer.WriteElementString("LN", "0.0.0.0.0.0")
					if err != nil {
						return err
					}
					err = writer.WriteElementString("Index", "0")
					if err != nil {
						return err
					}
					err = writer.WriteElementString("ParameterDataType", int(enums.DataTypeNone))
					if err != nil {
						return err
					}
					err = writer.WriteElementString("Parameter", "")
					if err != nil {
						return err
					}
				} else {
					err = writer.WriteElementString("ObjectType", int(a.Target.Base().ObjectType()))
					if err != nil {
						return err
					}
					err = writer.WriteElementString("LN", a.Target.Base().LogicalName())
					if err != nil {
						return err
					}
					err = writer.WriteElementString("Index", a.Index)
					if err != nil {
						return err
					}
					err = writer.WriteElementString("ParameterDataType", int(a.ParameterDataType))
					if err != nil {
						return err
					}
					if v, ok := a.Parameter.([]byte); ok {
						err = writer.WriteElementString("Parameter", types.ToHex(v, false))
						if err != nil {
							return err
						}
					} else {
						err = writer.WriteElementString("Parameter", fmt.Sprint(a.Parameter))
						if err != nil {
							return err
						}
					}
				}
				writer.WriteEndElement()
			}
			writer.WriteEndElement()
			writer.WriteEndElement()
		}
		writer.WriteEndElement()
	}
	return err
}

// PostLoad returns the handle actions after Load.
//
// Parameters:
//
//	reader: XML reader.
func (g *GXDLMSScriptTable) PostLoad(reader *GXXmlReader) error {
	return nil
}

// GetValues returns the an array containing the COSEM object's attribute values.
func (g *GXDLMSScriptTable) GetValues() []any {
	return []any{g.LogicalName(), g.Scripts}
}

// Execute executes selected script.
//
// Parameters:
//
//	client: DLMS client.
//	script: Executed script.
//
// Returns:
//
//	Action bytes.
func (g *GXDLMSScriptTable) Execute(client IGXDLMSClient, script *GXDLMSScript) ([][]byte, error) {
	return client.Method(g, 1, uint16(script.Id), enums.DataTypeUint16)
}

// ExecuteById  executes selected script by id.
//
// Parameters:
//
//	client: DLMS client.
//	scriptId: Executed script id.
//
// Returns:
//
//	Action bytes.
func (g *GXDLMSScriptTable) ExecuteById(client IGXDLMSClient, scriptId uint16) ([][]byte, error) {
	return client.Method(g, 1, scriptId, enums.DataTypeUint16)
}

// GetDataType returns the device data type of selected attribute index.
//
// Parameters:
//
//	index: Attribute index of the object.
//
// Returns:
//
//	Device data type of the object.
func (g *GXDLMSScriptTable) GetDataType(index int) (enums.DataType, error) {
	if index == 1 {
		return enums.DataTypeOctetString, nil
	}
	if index == 2 {
		return enums.DataTypeArray, nil
	}
	return enums.DataTypeNone, dlmserrors.ErrInvalidAttributeIndex
}

// NewGXDLMSScriptTable creates a new script table object instance.
//
// The function validates `ln` before creating the object.
//`ln` is the Logical Name and `sn` is the Short Name of the object.
func NewGXDLMSScriptTable(ln string, sn int16) (*GXDLMSScriptTable, error) {
	err := ValidateLogicalName(ln)
	if err != nil {
		return nil, err
	}
	return &GXDLMSScriptTable{
		GXDLMSObject: GXDLMSObject{
			objectType:  enums.ObjectTypeScriptTable,
			logicalName: ln,
			ShortName:   sn,
		},
	}, nil
}
