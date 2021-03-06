(*++

Copyright (C) 2019 Calculation developers

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.6.0.

Abstract: This is an autogenerated Pascal implementation file in order to allow easy
development of Calculation library. It needs to be generated only once.

Interface version: 1.0.0

*)

{$MODE DELPHI}
Unit calculation_impl;

interface

uses
  calculation_types,
  calculation_exception,
  calculation_interfaces,
  Classes,
  Unit_Numbers,
  calculation_impl_calculator,
  sysutils;

type
  TCalculationWrapper = class(TObject)
    private
      class var GNumbersWrapper: TNumbersWrapper;
      class function GetNumbersWrapper: TNumbersWrapper; static;
      class procedure SetNumbersWrapper(ANumbersWrapper: TNumbersWrapper); static;
    public
      class property NumbersWrapper: TNumbersWrapper read GetNumbersWrapper write SetNumbersWrapper;
      class function CreateCalculator(): TObject;
      class procedure GetVersion(out AMajor: Cardinal; out AMinor: Cardinal; out AMicro: Cardinal);
      class function GetLastError(AInstance: TObject; out AErrorMessage: String): Boolean;
      class procedure ReleaseInstance(AInstance: TObject);
      class procedure AcquireInstance(AInstance: TObject);
  end;


implementation

class procedure TCalculationWrapper.SetNumbersWrapper(ANumbersWrapper: TNumbersWrapper);
  begin
  if assigned(GNumbersWrapper) then
    raise ECalculationException.Create(CALCULATION_ERROR_COULDNOTLOADLIBRARY);
  GNumbersWrapper := ANumbersWrapper;
end;

class function TCalculationWrapper.GetNumbersWrapper(): TNumbersWrapper;
begin
  if not assigned(GNumbersWrapper) then
    raise ECalculationException.Create(CALCULATION_ERROR_COULDNOTLOADLIBRARY);
  result := GNumbersWrapper;
end;

class function TCalculationWrapper.CreateCalculator(): TObject;
begin
  result := TCalculationCalculator.Create();
end;

class procedure TCalculationWrapper.GetVersion(out AMajor: Cardinal; out AMinor: Cardinal; out AMicro: Cardinal);
begin
  AMajor := CALCULATION_VERSION_MAJOR;
  AMinor := CALCULATION_VERSION_MINOR;
  AMicro := CALCULATION_VERSION_MICRO;
end;

class function TCalculationWrapper.GetLastError(AInstance: TObject; out AErrorMessage: String): Boolean;
begin
  result := false;
  if supports(AInstance, ICalculationBase) then
    result := (AInstance as ICalculationBase).GetLastErrorMessage(AErrorMessage);
end;

class procedure TCalculationWrapper.ReleaseInstance(AInstance: TObject);
begin
  (AInstance as ICalculationBase).DecRefCount(); 
end;

class procedure TCalculationWrapper.AcquireInstance(AInstance: TObject);
begin
  (AInstance as ICalculationBase).IncRefCount(); 
end;


end.
