/*++

Copyright (C) 2019 ACT Developers


This file has been generated by the Automatic Component Toolkit (ACT) version 1.6.0.

Abstract: This is an autogenerated C++-Header file with basic types in
order to allow an easy use of Optional Class Library

Interface version: 1.0.0

*/

#ifndef __OPTCLASS_TYPES_HEADER_CPP
#define __OPTCLASS_TYPES_HEADER_CPP


/*************************************************************************************************************************
 Scalar types definition
**************************************************************************************************************************/

#ifdef OPTCLASS_USELEGACYINTEGERTYPES

typedef unsigned char OptClass_uint8;
typedef unsigned short OptClass_uint16 ;
typedef unsigned int OptClass_uint32;
typedef unsigned long long OptClass_uint64;
typedef char OptClass_int8;
typedef short OptClass_int16;
typedef int OptClass_int32;
typedef long long OptClass_int64;

#else // OPTCLASS_USELEGACYINTEGERTYPES

#include <stdint.h>

typedef uint8_t OptClass_uint8;
typedef uint16_t OptClass_uint16;
typedef uint32_t OptClass_uint32;
typedef uint64_t OptClass_uint64;
typedef int8_t OptClass_int8;
typedef int16_t OptClass_int16;
typedef int32_t OptClass_int32;
typedef int64_t OptClass_int64 ;

#endif // OPTCLASS_USELEGACYINTEGERTYPES

typedef float OptClass_single;
typedef double OptClass_double;

/*************************************************************************************************************************
 General type definitions
**************************************************************************************************************************/

typedef OptClass_int32 OptClassResult;
typedef void * OptClassHandle;
typedef void * OptClass_pvoid;

/*************************************************************************************************************************
 Version for OptClass
**************************************************************************************************************************/

#define OPTCLASS_VERSION_MAJOR 1
#define OPTCLASS_VERSION_MINOR 0
#define OPTCLASS_VERSION_MICRO 0
#define OPTCLASS_VERSION_PRERELEASEINFO ""
#define OPTCLASS_VERSION_BUILDINFO ""

/*************************************************************************************************************************
 Error constants for OptClass
**************************************************************************************************************************/

#define OPTCLASS_SUCCESS 0
#define OPTCLASS_ERROR_NOTIMPLEMENTED 1
#define OPTCLASS_ERROR_INVALIDPARAM 2
#define OPTCLASS_ERROR_INVALIDCAST 3
#define OPTCLASS_ERROR_BUFFERTOOSMALL 4
#define OPTCLASS_ERROR_GENERICEXCEPTION 5
#define OPTCLASS_ERROR_COULDNOTLOADLIBRARY 6
#define OPTCLASS_ERROR_COULDNOTFINDLIBRARYEXPORT 7
#define OPTCLASS_ERROR_INCOMPATIBLEBINARYVERSION 8

/*************************************************************************************************************************
 Declaration of handle classes 
**************************************************************************************************************************/

typedef OptClassHandle OptClass_Base;

namespace OptClass {

} // namespace OptClass;

// define legacy C-names for enums, structs and function types

#endif // __OPTCLASS_TYPES_HEADER_CPP