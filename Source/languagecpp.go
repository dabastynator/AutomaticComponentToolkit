/*++

Copyright (C) 2018 Autodesk Inc. (Original Author)

All rights reserved.

Redistribution and use in source and binary forms, with or without modification,
are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
list of conditions and the following disclaimer.
2. Redistributions in binary form must reproduce the above copyright notice,
this list of conditions and the following disclaimer in the documentation
and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR
ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

--*/

//////////////////////////////////////////////////////////////////////////////////////////////////////
// languagecpp.go
// functions to generate a CPP-layer of a library's API (can be used in bindings or implementation)
//////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

// CreateCPPTypesHeader creates a CPP header file for the types in component's API
func CreateCPPTypesHeader(component ComponentDefinition, CTypesHeaderName string) (error) {
	hTypesFile, err := CreateLanguageFile(CTypesHeaderName, "  ");
	if (err != nil) {
		return err;
	}
	hTypesFile.WriteCLicenseHeader (component,
		fmt.Sprintf ("This is an autogenerated C++-Header file with basic types in\norder to allow an easy use of %s", component.LibraryName),
		true);

	err = buildCCPPTypesHeader(component, hTypesFile, component.NameSpace, true);
	return err;
}

func getCPPMemberLine(member ComponentDefinitionMember, NameSpace string, arraysuffix string, structName string) (string, error) {
	switch (member.Type) {
		case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "single", "double", "bool", "pointer":
			typeName, err := getCPPParameterTypeName(member.Type, NameSpace, "")
			if (err != nil) {
				return "", err
			}
			return fmt.Sprintf("%s m_%s%s;", typeName, member.Name, arraysuffix), nil
		case "enum":
			return fmt.Sprintf("e%s m_%s%s;", member.Class, member.Name, arraysuffix), nil
		default:
			return "", fmt.Errorf ("it is not possible for struct %s to contain a %s member", structName, member.Type);
		
	}
}

// CreateCPPAbiHeader creates a CPP header file for the component's API
func CreateCPPAbiHeader(component ComponentDefinition, CHeaderName string) (error) {
	hfile, err := CreateLanguageFile(CHeaderName, "  ");
	if (err != nil) {
		return err;
	}
	hfile.WriteCLicenseHeader (component,
		fmt.Sprintf ("This is an autogenerated C++-Header file in order to allow an easy\n use of %s", component.LibraryName),
		true);
	err = buildCAbiHeader(component, hfile, component.NameSpace, component.BaseName, true);
	return err;
}

func getCPPParameterTypeName(ParamTypeName string, NameSpace string, ParamClass string)(string, error) {
	paramNameSpace, paramClassName, err := decomposeParamClassName(ParamClass)
	if (err != nil) {
		return "", err
	}
	if len(paramNameSpace) == 0 {
		paramNameSpace = NameSpace
	}

	cppParamTypeName := "";
	switch (ParamTypeName) {
		case "enum":
			cppParamTypeName = fmt.Sprintf ("%s::e%s", paramNameSpace, paramClassName);
		case "struct":
			cppParamTypeName = fmt.Sprintf ("%s::s%s *", paramNameSpace, paramClassName);
		case "structarray":
			cppParamTypeName = fmt.Sprintf ("%s::s%s *", paramNameSpace, paramClassName)
		case "class", "optionalclass":
			cppParamTypeName = fmt.Sprintf ("%s_%s", paramNameSpace, paramClassName)
		case "functiontype":
			cppParamTypeName = fmt.Sprintf ("%s::%s", paramNameSpace, paramClassName)
		default:
			cParamTypeName, err := getCParameterTypeName(ParamTypeName, paramNameSpace, paramClassName)
			if (err != nil) {
				return "", err
			}
			cppParamTypeName = cParamTypeName
	}
	return cppParamTypeName, nil;
}
