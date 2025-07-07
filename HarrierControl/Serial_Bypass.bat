REM **********************************************************************
REM Sets interface board into serial bypass mode
REM **********************************************************************
@echo off

SET PROG_EXE=HarrierControl.exe
SET ARG_BAUD=9600

cls
echo ---------------------------
echo Set Serial Bypass mode
echo ---------------------------
echo:
echo:

choice /C:12345 /M "Select Comms mode: 1(TTL), 2(RS-232), 3(RS-485), 4(USB), 5(COM port)?"
REM echo ERRORLEVEL=%ERRORLEVEL%
IF %ERRORLEVEL% EQU 1 SET ARG_COMMS=TTL 
IF %ERRORLEVEL% EQU 2 SET ARG_COMMS=RS232
IF %ERRORLEVEL% EQU 3 SET ARG_COMMS=RS485
IF %ERRORLEVEL% EQU 4 SET ARG_COMMS=USB3
IF %ERRORLEVEL% EQU 5 SET ARG_COMMS=COM
echo:

REM muliptle IFs due to delayed expansion
IF %ARG_COMMS%==COM SET /P PORT_NUM=Enter COM port number:
IF %ARG_COMMS%==COM SET ARG_COMMS=COM%PORT_NUM%

IF %ARG_COMMS%==USB3 (
REM  SET PROG_CMD=%PROG_EXE% %ARG_COMMS%
  echo:
  echo Note: serial bypass does not work in USB mode...
  pause
  goto :END
) else (
  SET PROG_CMD=%PROG_EXE% %ARG_COMMS% %ARG_BAUD%
)

echo: 
echo ARG_COMMS=%ARG_COMMS%
echo PROG_CMD=%PROG_CMD%
echo:
echo Setting serial bypass...
REM serial bypass command
%PROG_CMD% 82,01,0A,06,00,FF 
echo:
%PROG_CMD% FF
echo:
echo:
echo Testing connection to camera...
%PROG_CMD% 81,09,00,02,FF /P
echo:
echo:
pause


:END