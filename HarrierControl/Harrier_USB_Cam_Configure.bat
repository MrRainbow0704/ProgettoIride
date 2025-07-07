@echo off
CLS
REM **********************************************************************
REM AS-CIB-USBHDMI-002-A USB Configuration Batch File
REM **********************************************************************
REM Version: 08
SET SCRIPT_VER=08
REM Date:    10-September-2024
REM **********************************************************************
REM Required equipment:
REM     USB SS cable Type-C
REM     High current/power SS USB socket
REM **********************************************************************

SET HC_VER=v1_4_2
SET WHT=[0m
SET RED=[91m
SET GRN=[92m
SET YEL=[93m
SET BLU=[94m
SET MAG=[95m
SET BOLD=[1m
SET UNDR=[4m
SET RST=[0m

:start
cls
echo ------------------------------------------------
echo Harrier USB Camera Configuration Utility V.1.%SCRIPT_VER%
echo ------------------------------------------------
echo:
echo Select the make/model of camera you wish to configure for: 
echo:
echo   (1) = Harrier 10LHD camera
echo   (2) = Harrier 36LGHD camera
echo   (3) = Harrier 40LHD camera
echo   (4) = Harrier 55LHD camera
echo   (5) = Tamron MP1010M-VC/WP camera
echo   (6) = Tamron MP1110M-VC/WP camera
echo   (7) = Tamron MP3010M-EV camera
echo   (8) = Tamron MP2030M-GS camera
echo   (9) = Sony EV7520A camera
echo   (A) = Sony EV7520 camera
echo   (B) = Sony EV9500L camera
echo   (C) = Sony EV9520L camera
echo   (q) = Quit
echo:

choice /C:123456789ABCQ /M "Select option"
REM echo ERRORLEVEL=%ERRORLEVEL%
REM SERIAL_CAMERA=8 chars max
IF %ERRORLEVEL% EQU 1 SET SERIAL_CAMERA=AS10LHD
IF %ERRORLEVEL% EQU 2 SET SERIAL_CAMERA=AS36LGHD
IF %ERRORLEVEL% EQU 3 SET SERIAL_CAMERA=AS40LHD
IF %ERRORLEVEL% EQU 4 SET SERIAL_CAMERA=AS55LHD
IF %ERRORLEVEL% EQU 5 SET SERIAL_CAMERA=TNMP1010
IF %ERRORLEVEL% EQU 6 SET SERIAL_CAMERA=TNMP1110
IF %ERRORLEVEL% EQU 7 SET SERIAL_CAMERA=TNMP3010
IF %ERRORLEVEL% EQU 8 SET SERIAL_CAMERA=TN2030GS
IF %ERRORLEVEL% EQU 9 SET SERIAL_CAMERA=SY7520A
IF %ERRORLEVEL% EQU 10 SET SERIAL_CAMERA=SY7520
IF %ERRORLEVEL% EQU 11 SET SERIAL_CAMERA=SY9500L
IF %ERRORLEVEL% EQU 12 SET SERIAL_CAMERA=SY9520L
IF %ERRORLEVEL% EQU 13 GOTO :END
echo: 

REM echo Select the operating system you wish to configure for: 
REM echo:
REM echo   (W) = Windows 10
REM echo   (L) = Linux
REM echo:

REM choice /C:wl /M "Select option (W/L): "
REM echo ERRORLEVEL=%ERRORLEVEL%
REM echo:
REM IF %ERRORLEVEL% EQU 1 SET HOST_OS=WIN10
REM IF %ERRORLEVEL% EQU 2 SET HOST_OS=LINUX
SET HOST_OS=LINUX

IF %HOST_OS%==WIN10 SET SERIAL_SUFFIX=1
IF %HOST_OS%==LINUX SET SERIAL_SUFFIX=2

REM muliptle IFs due to delayed expansion

REM Harrier 10LHD camera
IF %SERIAL_CAMERA%==AS10LHD SET CAMERA=Harrier 10LHD camera
IF %SERIAL_CAMERA%==AS10LHD SET FTDI_CONF_FILE=AP23C06_%HC_VER%_10LHD.flash

REM Harrier 36LGHD camera
IF %SERIAL_CAMERA%==AS36LGHD SET CAMERA=Harrier 36LGHD camera
IF %SERIAL_CAMERA%==AS36LGHD SET FTDI_CONF_FILE=AP23C06_%HC_VER%_36LGHD.flash

REM Harrier 40LHD camera
IF %SERIAL_CAMERA%==AS40LHD SET CAMERA=Harrier 40LHD camera
IF %SERIAL_CAMERA%==AS40LHD SET FTDI_CONF_FILE=AP23C06_%HC_VER%_40LHD.flash

REM Harrier 55LHD camera
IF %SERIAL_CAMERA%==AS55LHD SET CAMERA=Harrier 55LHD camera
IF %SERIAL_CAMERA%==AS55LHD SET FTDI_CONF_FILE=AP23C06_%HC_VER%_55LHD.flash

REM Tamron MP1110M-VC/WP camera
IF %SERIAL_CAMERA%==TNMP1110 SET CAMERA=Tamron MP1110M-VC/WP camera
IF %SERIAL_CAMERA%==TNMP1110 SET FTDI_CONF_FILE=AP23C06_%HC_VER%_MP1110.flash

REM Tamron MP1010 camera
IF %SERIAL_CAMERA%==TNMP1010 SET CAMERA=Tamron MP1010 camera
IF %SERIAL_CAMERA%==TNMP1010 SET FTDI_CONF_FILE=AP23C06_%HC_VER%_MP1010.flash

REM Tamron MP3010 camera
IF %SERIAL_CAMERA%==TNMP3010 SET CAMERA=Tamron MP3010 camera
IF %SERIAL_CAMERA%==TNMP3010 SET FTDI_CONF_FILE=AP23C06_%HC_VER%_MP3010.flash

REM Tamron MP2030GS camera
IF %SERIAL_CAMERA%==TN2030GS SET CAMERA=TN2030GS
IF %SERIAL_CAMERA%==TN2030GS SET FTDI_CONF_FILE=AP23C06_%HC_VER%_MP2030.flash
IF %SERIAL_CAMERA%==TN2030GS SET SERIAL_SUFFIX=2

REM Sony EV-FCB7520A camera
IF %SERIAL_CAMERA%==SY7520A SET CAMERA=Sony FCB-EV7520A camera
IF %SERIAL_CAMERA%==SY7520A SET FTDI_CONF_FILE=AP23C06_%HC_VER%_EV7520A.flash

REM Sony EV-FCB7520 camera
IF %SERIAL_CAMERA%==SY7520 SET CAMERA=Sony FCB-EV7520 camera
IF %SERIAL_CAMERA%==SY7520 SET FTDI_CONF_FILE=AP23C06_%HC_VER%_EV7520.flash

REM Sony EV-FCB9500L camera
IF %SERIAL_CAMERA%==SY9500L SET CAMERA=Sony FCB-EV9500L camera
IF %SERIAL_CAMERA%==SY9500L SET FTDI_CONF_FILE=AP23C06_%HC_VER%_EV9500L.flash
IF %SERIAL_CAMERA%==SY9500L SET SERIAL_SUFFIX=2

REM Sony EV-FCB9520L camera
IF %SERIAL_CAMERA%==SY9520L SET CAMERA=Sony FCB-EV9520L camera
IF %SERIAL_CAMERA%==SY9520L SET FTDI_CONF_FILE=AP23C06_%HC_VER%_EV9520L.flash
IF %SERIAL_CAMERA%==SY9520L SET SERIAL_SUFFIX=2

REM IF %HOST_OS%==WIN10 GOTO :WIN10
REM Harrier 10LHD camera
REM IF %SERIAL_CAMERA%==AS10LHD SET FTDI_CONF_FILE=AP23C06_%HC_VER%_10LHD.flash
REM Harrier 36LGHD camera already configured as self-powered

REM Harrier 40LHD camera
REM IF %SERIAL_CAMERA%==AS40LHD SET FTDI_CONF_FILE=AP23C06_%HC_VER%_40LHD_Linux.flash

REM Harrier 55LHD camera
REM IF %SERIAL_CAMERA%==AS55LHD SET FTDI_CONF_FILE=AP23C06_%HC_VER%_55LHD.flash

REM Tamron MP1110M-VC/WP camera
REM IF %SERIAL_CAMERA%==TNMP1110 SET FTDI_CONF_FILE=AP23C06_%HC_VER%_MP1110_Linux.flash
REM Tamron MP1010 camera
REM IF %SERIAL_CAMERA%==TNMP1010 SET FTDI_CONF_FILE=AP23C06_%HC_VER%_MP1010_Linux.flash
REM Tamron MP1010 camera
REM IF %SERIAL_CAMERA%==TNMP3010 SET FTDI_CONF_FILE=AP23C06_%HC_VER%_MP3010_Linux.flash
REM Tamron MP2030GS camera already configured as self-powered
REM Sony EV9500L camera already configured as self-powered
REM Sony EV9520L camera already configured as self-powered
REM Sony FCB-EV7520A camera
REM IF %SERIAL_CAMERA%==SY7520A SET FTDI_CONF_FILE=AP23C06_%HC_VER%_EV7520A_Linux.flash
REM Sony FCB-EV7520 camera
REM IF %SERIAL_CAMERA%==SY7520 SET FTDI_CONF_FILE=AP23C06_%HC_VER%_EV7520_Linux.flash

:WIN10

REM version of config 1 char only
SET SERIAL_MACHINE=C

REM machine used 1 char only
SET FTDI_CONF_EXE=Harriercontrol.exe
SET FTDI_CONF_CHANGE=%FTDI_CONF_EXE% USB3 LoadConfig %FTDI_CONF_FILE% SetSerial
SET FTDI_CONF_SER_NUM_CHECK=%FTDI_CONF_EXE% USB3 GetSerial

REM for /F "delims= " %%G IN ('"powershell get-date -format yyyy-MM-dd"') do SET TODAY_DATE=%%G

SET CONF_VER=%SERIAL_CAMERA%-%SERIAL_SUFFIX%

REM Check filenames exist
SET F_ERROR=0

IF NOT EXIST "%FTDI_CONF_FILE%" (
echo Missing %FTDI_CONF_FILE%
SET F_ERROR=1
)
IF NOT EXIST "%FTDI_CONF_EXE%" (
echo Missing %FTDI_CONF_EXE%
SET F_ERROR=2
)
IF %F_ERROR% NEQ 0 GOTO FILE_ERROR
REM pause
cls

echo -----------------------------------------------
echo Harrier USB Camera Configuration Utility V.%SCRIPT_VER%
echo -----------------------------------------------
echo:
echo Opening Windows Device Manager....
start devmgmt.msc
echo:
echo You will have to click on this window to restore the focus 
echo:
echo:
pause

REM **********************************************************************
:PROGRAM_BOARD
REM **********************************************************************

cls
echo ----------------------------------------
echo Harrier USB Configuration Utility V.%SCRIPT_VER%
echo ----------------------------------------
echo:
echo Please do the following:
echo:
echo 1) Check that DIP Switch 7 is set to ON (down)
echo 2) Connect Superspeed (5Gbps) USB Type-C cable to SuperSpeed socket on PC.
echo 3) Connect Superspeed (5Gbps) USB Type-C cable to J6 (USB Video) on Harrier USB/HDMI Board to be configured. 
echo 4) Check that the LED on the board illuminates.
echo:
echo:
echo Note: when you plug in the USB Type-C cable the USB board will power-up and the LED should illuminate. 
echo:
echo Please wait for the Harrier device to appear in Windows Device Manager under Cameras 
echo:
pause
cls
echo ----------------------------------------------
echo Harrier USB Camera Configuration Utility V.%SCRIPT_VER%
echo ----------------------------------------------
echo:
echo Current serial number/configuration:
%FTDI_CONF_SER_NUM_CHECK%
IF %ERRORLEVEL% NEQ 0 GOTO HC_ERROR
echo:
echo:
echo Serial number codes:
echo   TNMP1010 = Tamron MP1010M-VC/WP camera
echo   TNMP1110 = Tamron MP1110M-VC/WP camera
echo   TNMP3010 = Tamron MP3010M-EV/WP camera
echo   TN2030GS = Tamron MP2030M-GS camera
echo   SY7520A  = Sony EV7520A camera
echo   SY7520   = Sony EV7520 camera
echo   SY9500L  = Sony EV9500L camera
echo   SY9520L  = Sony EV9520L camera
echo   AS10LHD  = Harrier 10LHD camera
echo   AS36LGHD = Harrier 36LGHD camera
echo   AS40LHD  = Harrier 40LHD camera
echo   AS55LHD  = Harrier 55LHD camera
echo:
echo:
REM echo %CAMERA%   %SERIAL_CAMERA%   %HOST_OS%   %SERIAL_SUFFIX%   %FTDI_CONF_FILE% 
echo %CAMERA%   %SERIAL_CAMERA%   %SERIAL_SUFFIX%   %FTDI_CONF_FILE% 
echo:
echo:
REM echo New configuration to be set: %GRN%%CAMERA% - %HOST_OS%%WHT%
echo New configuration to be set: %GRN%%CAMERA% %WHT%
echo:
choice /C:YN /M " Do you really want to change the USB board configuration"
IF ERRORLEVEL 2 GOTO END

REM ********** program device **************

cls
echo ----------------------------------------------
echo Harrier USB Camera Configuration Utility V.%SCRIPT_VER%
echo ----------------------------------------------
echo:
echo Configuring USB device...
REM - get universal epoch time and write to user serial number field with FT602 release configuration.
for /F "tokens=1,2" %%G IN ('"powershell [long]((date).touniversaltime()-[datetime]'1970-01-01').totalseconds"') do SET UNI_TIME=%%G

SET SERIAL_NUM=%SERIAL_MACHINE%-%UNI_TIME%-%SERIAL_CAMERA%-%SERIAL_SUFFIX%
echo:
echo Generated Serial Num : %SERIAL_NUM%

%FTDI_CONF_CHANGE% %SERIAL_NUM%
IF %ERRORLEVEL% NEQ 0 GOTO HC_ERROR

echo:
echo:
echo Wait for Windows to recognize and setup the new configuration
echo:
echo Please do NOT quit
TIMEOUT 25 /NOBREAK
echo:
echo:
echo Reading back serial number...
echo:

%FTDI_CONF_SER_NUM_CHECK%
IF %ERRORLEVEL% NEQ 0 GOTO HC_ERROR

echo:
echo:
echo ...USB configuration done.
echo:
echo Generated and read back serial number should be identical 
echo:
echo Serial number codes:
echo   TNMP1010 = Tamron MP1010M-VC/WP camera
echo   TNMP1110 = Tamron MP1110M-VC/WP camera
echo   TNMP3010 = Tamron MP3010M-EV/WP camera
echo   TN2030GS = Tamron MP2030M-GS camera
echo   SY7520A  = Sony EV7520A camera
echo   SY7520   = Sony EV7520 camera
echo   SY9500L  = Sony EV9500L camera
echo   SY9520L  = Sony EV9520L camera
echo   AS10LHD  = Harrier 10LHD camera
echo   AS36LGHD = Harrier 36LGHD camera
echo   AS40LHD  = Harrier 40LHD camera
echo   AS55LHD  = Harrier 55LHD camera
echo:
echo:
echo The name of the Harrier device in Windows Device Manager should change to indicate the %CAMERA% camera.
echo:
echo ----
echo:
choice /C:YN /M "Q: Configure another camera (same camera make/model)"
IF ERRORLEVEL 2 GOTO END
IF ERRORLEVEL 1 GOTO PROGRAM_BOARD
GOTO END
REM **********************************************************************

:HC_ERROR
echo:
echo HarrierControl error - install VC_redist.exe and check FTD2XX.dll
echo:
goto :end


:FILE_ERROR
echo:
echo File is missing or incorrectly set
echo:
pause

:END
echo:
echo:
echo Press any key to exit.
echo:
pause > nul
REM --------------------------------------------------------
REM End of File.
REM --------------------------------------------------------
