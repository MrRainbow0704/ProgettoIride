@echo off

REM ****** Camera Demo Batch file for Hariercontrol v1.4.2 ******
REM ****** 10 Sept 2024 ******

SETLOCAL ENABLEDELAYEDEXPANSION
SETLOCAL ENABLEEXTENSIONS

SET ARG_BAUD=9600
SET NO_1080P60=FALSE
SET XH_SET=ON
SET LVDS_SDI=LVDS
SET LVDS_NEXT=SDI
SET FLIP=ON
SET MIRROR=ON
SET ZVALUE=0000
SET COMVALUE=0

SET   CMD_DUAL_LVDS=81,01,04,24,74,00,01,FF
SET CMD_SINGLE_LVDS=81,01,04,24,74,00,00,FF
SET   CMD_CIB_RESET=82,01,0A,00,FF
SET      CAM_REBOOT=81,01,04,00,03,FF

SET ZOOMx10=81,01,04,47,04,00,00,00,FF
SET ZOOMx4=81,01,04,47,02,0D,06,00,FF
SET ZOOMx0=81,01,04,47,00,00,00,00,FF

SET TEXT_ON=81,01,04,74,2F,FF
SET TEXT_OFF=81,01,04,74,3F,FF
SET FLIP_ON=81,01,04,66,02,FF
SET FLIP_OFF=81,01,04,66,03,FF
SET COLOUR_HUE=81,01,04,4F,00,00,00,00,FF
SET CAM_LR_Reverse_ON=81,01,04,61,02,FF
SET CAM_LR_Reverse_OFF=81,01,04,61,03,FF
SET CMD_CIB_X_HAIR_OFF=82,01,0A,03,00,FF
SET CMD_CIB_X_HAIR_ON=82,01,0A,03,01,FF

rem set title position, size and string for Tamron 1110 camera
rem %PROG_EXE% 81,01,04,73,10,00,05,03,03,00,00,00,00,00,00,FF 
rem %PROG_EXE% 81,01,04,73,21,43,2A,44,1B,00,2A,3B,30,3D,2C,FF 
rem %PROG_EXE% 81,01,04,73,31,1B,12,30,33,30,2A,36,35,1B,1B,FF 

rem set title position, size and string for Tamron 2030 camera
REM %PROG_EXE% 81,01,04,73,11,00,00,03,07,00,00,00,00,00,00,FF >NUL 
REM %PROG_EXE% 81,01,04,73,21,1B,02,1B,1B,00,02,13,08,15,04,FF >NUL
REM %PROG_EXE% 81,01,04,73,31,1B,12,08,0B,08,02,0E,0D,1B,1B,FF >NUL

cls
echo ------------------------------------
echo  [7m    HarrierControl Camera Demo    [0m
echo ------------------------------------
echo.
echo.
echo Select Camera Communications mode.....
echo.
echo  1. TTL (9600 baud)
echo  2. RS-232 (9600 baud)
echo  3. USB 
echo  4. COM port
echo.
echo.
choice /C:1234 /M "Select Comms mode " 
rem echo ERRORLEVEL=%ERRORLEVEL%
echo.
IF %ERRORLEVEL% EQU 1 SET ARG_COMMS=TTL
IF %ERRORLEVEL% EQU 2 SET ARG_COMMS=RS232
IF %ERRORLEVEL% EQU 3 SET ARG_COMMS=USB3
IF %ERRORLEVEL% EQU 4 SET /P COMVALUE=Enter COM port number... 

IF %COMVALUE% NEQ 0 SET ARG_COMMS=COM%COMVALUE%
echo.
echo %ARG_COMMS%
echo.

SET PROG_EXE=HarrierControl.exe %ARG_COMMS% %ARG_BAUD%
IF %ARG_COMMS%==USB3 SET PROG_EXE=HarrierControl.exe %ARG_COMMS%

echo.
echo %PROG_EXE%

cls
echo ------------------------------------
echo  [7m    HarrierControl Camera Demo    [0m
echo ------------------------------------
echo.
echo.
echo Select Camera.....
echo.
echo.
echo 1. Harrier 10x AFZ camera
echo 2. Harrier 36x AFZ camera
echo 3. Harrier 40x AFZ camera
echo 4. Harrier 55x AFZ camera
echo 5. Tamron MP3010 camera
echo 6. Sony FCB-EV7520/7520A/9500l/9520L
echo.
echo.


choice /C:123456 /M "Select Camera  "
rem echo ERRORLEVEL=%ERRORLEVEL%
IF %ERRORLEVEL% EQU 1 GOTO H10x
IF %ERRORLEVEL% EQU 2 GOTO H36x
IF %ERRORLEVEL% EQU 3 GOTO H40x
IF %ERRORLEVEL% EQU 4 GOTO H55x
IF %ERRORLEVEL% EQU 5 GOTO T3010
IF %ERRORLEVEL% EQU 6 GOTO SONY

:H36x
:H55x
SET NO_1080P60=TRUE
:H10x
:H40x
SET CMD_MODE_1080P60=81,01,04,24,72,01,03,FF
SET CMD_MODE_1080P30=81,01,04,24,72,00,06,FF
SET  CMD_MODE_720P60=81,01,04,24,72,00,09,FF
GOTO :START
:T3010
SET CMD_MODE_1080P60=81,01,04,24,72,00,07,FF
SET CMD_MODE_1080P30=81,01,04,24,72,00,01,FF
SET  CMD_MODE_720P60=81,01,04,24,72,00,05,FF
GOTO :START
:SONY
SET CMD_MODE_1080P60=81,01,04,24,72,01,05,FF
SET CMD_MODE_1080P30=81,01,04,24,72,00,07,FF
SET  CMD_MODE_720P60=81,01,04,24,72,00,0A,FF
GOTO :START

:START
cls
echo ------------------------------------
echo  [7m    HarrierControl Camera Demo    [0m
echo ------------------------------------
echo.
echo 1. Set Max Optical Zoom 
echo 2. Set Mid range Optical Zoom
echo 3. Set 0x Optical Zoom
echo 4. Optical Zoom manual setting
echo 5. Set Crosshairs %XH_SET% 
echo 6. Set Camera Image 180 Flip %FLIP% 
echo 7. Set Camera Image Mirror %MIRROR% 
echo 8. Set 1080p60
echo 9. Set 1080p30
echo A. Set 720p60
echo Q. Quit
echo.
choice /C 123456789AQX /N /M "Please Select Option: "
if !errorlevel! EQU 12 goto end
if !errorlevel! EQU 11 goto end
if !errorlevel! EQU 10 goto 720p60
if !errorlevel! EQU 9  goto 1080p30
if !errorlevel! EQU 8  goto 1080p60
if !errorlevel! EQU 7  goto MIRROR
if !errorlevel! EQU 6  goto FLIP
if !errorlevel! EQU 5  goto XHAIR
if !errorlevel! EQU 4  goto MZOOM
if !errorlevel! EQU 3  %PROG_EXE% %ZOOMx0%
if !errorlevel! EQU 2  %PROG_EXE% %ZOOMx4%
if !errorlevel! EQU 1  %PROG_EXE% %ZOOMx10%
REM pause
goto START

:720p60
%PROG_EXE% %CMD_SINGLE_LVDS%  /P
echo.
%PROG_EXE% %CMD_MODE_720P60% /P
echo.
%PROG_EXE% %CMD_CIB_RESET% /P
goto START


:1080p30
%PROG_EXE% %CMD_SINGLE_LVDS% /P
echo.
%PROG_EXE% %CMD_MODE_1080P30% /P
echo.
%PROG_EXE% %CMD_CIB_RESET% /P
goto START

:1080p60
if %NO_1080P60%==TRUE (
 echo.
 echo 1080p60 not supported on this camera. 
 echo.
 pause
) else (
%PROG_EXE% %CMD_DUAL_LVDS% /P
echo.
%PROG_EXE% %CMD_MODE_1080P60% /P
echo.
%PROG_EXE% %CMD_CIB_RESET% /P
)
goto START

:FLIP
if %FLIP%==ON (
    echo switch on FLIP
    echo.
    %PROG_EXE% %FLIP_ON%
    SET FLIP=OFF
) else (
    echo switch off FLIP
    echo.
    %PROG_EXE% %FLIP_OFF%
    SET FLIP=ON
)
goto START

:MIRROR
if %MIRROR%==ON (
    echo switch on MIRROR
    echo.
    %PROG_EXE% %CAM_LR_Reverse_ON%
    SET MIRROR=OFF
) else (
    echo switch off MIRROR
    echo.
    %PROG_EXE% %CAM_LR_Reverse_OFF%
    SET MIRROR=ON
)
goto START

:XHAIR
if %XH_SET%==ON (
    echo switch on Crosshair
    echo.
	%PROG_EXE% %CMD_CIB_X_HAIR_ON%
    SET XH_SET=OFF
) else (
    echo switch off Crosshair
    echo.
	%PROG_EXE% %CMD_CIB_X_HAIR_OFF%
    SET XH_SET=ON
)
goto START

:MZOOM
echo.
set ZVALUE=0
set /P ZVALUE=Enter 'Q' to quit Manual Zoom or Enter Zoom Value [0000-4000(hex)]:
echo.

if /I "%ZVALUE%"=="Q" goto START

SET /A VALUE=0x%ZVALUE% 2>NUL

IF %ERRORLEVEL% GTR 100000 (
   echo Error: Invalid number, all characters must be hexadecimal numbers [0-F]
   rem reset errorlevel to 0
   cd .\
   goto MZOOM
)

if %VALUE% GTR 0x4000 (
   echo Error: value too large - %ZVALUE%[hex] - 4000[hex] maximum
   goto MZOOM
) else (
   if "!ZVALUE:~1,1!"=="" set ZVALUE=000%ZVALUE%
   if "!ZVALUE:~2,1!"=="" set ZVALUE=00%ZVALUE%
   if "!ZVALUE:~3,1!"=="" set ZVALUE=0%ZVALUE%
   echo Send !ZVALUE! zoom value to camera
   %PROG_EXE% 81,01,04,47,0!ZVALUE:~0,1!,0!ZVALUE:~1,1!,0!ZVALUE:~2,1!,0!ZVALUE:~3,1!,FF
)
goto MZOOM
goto START

:end
echo.
echo bye...
echo.

