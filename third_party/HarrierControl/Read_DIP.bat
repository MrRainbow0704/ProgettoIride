REM **********************************************************************
REM Reads DIP switches and prints value
REM **********************************************************************
@echo off

SET PROG_EXE=HarrierControl.exe
SET %ARG_BAUD%=9600
cls
echo ---------------------------
echo Read DIP Switch settings
echo ---------------------------
echo.
echo.
echo Select Comms mode:
echo.
echo.

choice /C:12345 /M "Select Comms mode: 1(TTL), 2(RS-232), 3(RS-485), 4(USB), 5(COM port)?"
REM echo ERRORLEVEL=%ERRORLEVEL%
IF %ERRORLEVEL% EQU 1 SET ARG_COMMS=TTL
IF %ERRORLEVEL% EQU 2 SET ARG_COMMS=RS232
IF %ERRORLEVEL% EQU 3 SET ARG_COMMS=RS485
IF %ERRORLEVEL% EQU 4 SET ARG_COMMS=USB3
IF %ERRORLEVEL% EQU 5 SET ARG_COMMS=COM
echo.

REM multiple IFs due to delayed expansion
IF %ARG_COMMS%==COM SET /P PORT_NUM=Enter COM port number:
IF %ARG_COMMS%==COM SET ARG_COMMS=COM%PORT_NUM%

IF %ARG_COMMS%==USB3 (
  SET PROG_CMD=%PROG_EXE% %ARG_COMMS%
) else (
  SET PROG_CMD=%PROG_EXE% %ARG_COMMS% %ARG_BAUD%
)

echo: 
echo PROG_EXE=%PROG_EXE%
echo ARG_COMMS=%ARG_COMMS%
echo PROG_CMD=%PROG_CMD%
echo:
REM pause

cls

REM %PROG_CMD% %CMD_READ_DIP% %ARG_PARSED%
REM echo Rx: [ 90 50 0%GRN%P%WHT% 0%GRN%Q%WHT% ff ]

echo ---------------------------
echo Read DIP Switch settings
echo ---------------------------
echo.

for /F "skip=1 tokens=5" %%G IN ('"%PROG_CMD% %CMD_READ_DIP%"') do SET DIP_READ=%%G
echo Current setting = %DIP_READ%
SET /A "SW1=(DIP_READ & 1)"
SET /A "SW2=(DIP_READ & 2)>>1"
SET /A "SW3=(DIP_READ & 4)>>2"
SET /A "SW4=(DIP_READ & 8)>>3"
SET /A "SW5=(DIP_READ & 16)>>4"
SET /A "SW6=(DIP_READ & 32)>>5"
SET /A "SW7=(DIP_READ & 64)>>6"
SET /A "SW8=(DIP_READ & 128)>>7"
echo SW1 = %SW1%  
echo SW2 = %SW2%  
echo SW3 = %SW3% 
echo SW4 = %SW4%
echo SW5 = %SW5%  
echo SW6 = %SW6%  
echo SW7 = %SW7% 
echo SW8 = %SW8%
echo:
pause