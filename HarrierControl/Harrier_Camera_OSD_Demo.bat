@echo off

REM ****** Harrier Camera OSD Demo Batch file for HarierControl v1.2.1 ******
REM ****** 27 September 2021 ******

SETLOCAL ENABLEDELAYEDEXPANSION
SETLOCAL ENABLEEXTENSIONS

SET ARG_BAUD=9600
SET COMVALUE=0 
SET WHT=[0m
SET BOLD=[1m
SET RST=[0m

cls
echo ---------------------------------------------
echo Harrier Camera OSD Demo
echo ---------------------------------------------
echo:
HarrierCOMPortDiscovery.exe
echo:
echo Select Camera Communications mode.....
echo:
echo %BOLD%
echo  1. TTL (9600 baud)
echo  2. RS-232 (9600 baud)
echo  3. USB 
echo  4. COM port
echo %RST%
echo:
choice /C:1234 /M "Select Comms mode: 1(TTL), 2(RS-232), 3(USB), 4(COM port)?"
echo:
rem echo ERRORLEVEL=%ERRORLEVEL%
IF %ERRORLEVEL% EQU 1 SET ARG_COMMS=TTL
IF %ERRORLEVEL% EQU 2 SET ARG_COMMS=RS232
IF %ERRORLEVEL% EQU 3 SET ARG_COMMS=USB3
IF %ERRORLEVEL% EQU 4 SET /P COMVALUE=Enter COM port number... 

IF %COMVALUE% NEQ 0 SET ARG_COMMS=COM%COMVALUE%
echo:
echo %ARG_COMMS%
echo:

SET PROG_EXE=HarrierControl %ARG_COMMS% %ARG_BAUD%
IF %ARG_COMMS%==USB3 SET PROG_EXE=HARRIERCONTROL.exe %ARG_COMMS%

echo:
echo %PROG_EXE%
rem pause
SET UP=81,01,04,16,01,FF
SET DOWN=81,01,04,16,02,FF
SET LEFT=81,01,04,16,04,FF
SET RIGHT=81,01,04,16,08,FF
SET MENU=81,01,04,16,10,FF
SET ESC=81,01,04,16,20,FF
SET BK_ON=81,01,05,20,11,FF
rem 25% transparent bluescreen
SET BK_OFF=81,01,05,20,10,FF

:START
cls
echo:
echo ================================
echo Harrier Camera OSD Demo
echo ================================

echo %BOLD%
echo  UP    ( 8 / U )
echo  DOWN  ( 2 / D )
echo  LEFT  ( 4 / L )
echo  RIGHT ( 6 / R )
echo  MENU  ( 5 / M )
echo:
echo  Quit OSD Menu ( 7 / Q )
echo %RST%
echo:
echo Select MENU to open menu system or select option highlighted on the on-screen menu
echo:
echo (B) for backgound on, (N) for no background
echo:
echo (X) to Exit batch file
echo:
echo Set NumLock ON to use number keypad.
echo Note: responses to keypresses are slow and are not buffered.
echo:
choice /C 824657udlrmQbnX /N /M "Please Select Option: "
rem choice /C 824657udlrmQbnX /N 
rem echo sending....
if !errorlevel! EQU 15 goto end
if !errorlevel! EQU 14 %PROG_EXE% %BK_OFF% >NUL
if !errorlevel! EQU 13 %PROG_EXE% %BK_ON% >NUL
if !errorlevel! EQU 12 %PROG_EXE% %ESC% >NUL
if !errorlevel! EQU 11 %PROG_EXE% %MENU% >NUL
if !errorlevel! EQU 10 %PROG_EXE% %RIGHT% >NUL
if !errorlevel! EQU 9  %PROG_EXE% %LEFT% >NUL
if !errorlevel! EQU 8  %PROG_EXE% %DOWN% >NUL
if !errorlevel! EQU 7  %PROG_EXE% %UP% >NUL
if !errorlevel! EQU 6  %PROG_EXE% %ESC% >NUL
if !errorlevel! EQU 5  %PROG_EXE% %MENU% >NUL
if !errorlevel! EQU 4  %PROG_EXE% %RIGHT% >NUL
if !errorlevel! EQU 3  %PROG_EXE% %LEFT% >NUL
if !errorlevel! EQU 2  %PROG_EXE% %DOWN% >NUL
if !errorlevel! EQU 1  %PROG_EXE% %UP% >NUL
REM pause
goto START

:end
echo:
echo Bye...
echo:

