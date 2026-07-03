HarrierControl Software Download
================================

This distribution contains several executables/software applications and batch scripts that are to be used with Harrier cameras and Harrier interface boards   
(e.g. HarrierControl.exe, HCGUI.exe, HarrierCOMPortDiscovery.exe, SetBaudRate.exe). 
These software applications are samples intended to support the evaluation and testing of Active Silicon products (Active Silicon camera interface boards, adapters and supported cameras). 
They should not be used in commercial/end user applications. They should not be used with products from other suppliers. 
There is no warranty or obligation for Active Silicon to support, maintain or provide fixes to reported failures in the software applications. 
Supported operating system:  Windows 10/11 - 64 bit.
For licensing terms and conditions please see the file License.txt.

Please report any issues to Tech Support techsupport@activesilicon.com


USB/HDMI Interface Board USB Firmware Update
============================================

This distribution contains tools and files for updating the USB firmware in the AS-CIB-USBHDMI-002-A.

Run the batch file: 

   Harrier_USB_Cam_Configure.bat 

and select the correct options for your camera.

The batch script requires write access to the folder it is in to store temporary data.

The USB firmware update is done using the HarrierControl.exe program.
This relies on DLLs from Visual C that may need to be installed.
If a missing file error is reported when trying to run the program, please install VC_redist.x64.exe.  



HarrierControl.exe
==================

For licensing terms and conditions please see the file License.txt.

The HarrierControl program is a only a sample application, so there may be issues with some of the features. 
Please report any issues to Tech Support techsupport@activesilicon.com

The Harrier Evaluation Kit Board supports serial communication with Tamron Cameras via an FTDI mini-module.
RS-232, RS-485 or TTL serial communication formats are supported.

Only one Harrier Evaluation Kit Board or one USB/HDMI camera interface board at a time should be connected to the PC used for testing.

The HarrierControl application automatically finds and connects to USB/HDMI camera interface boards and/or the FTDI mini-module fitted to the Harrier Evaluation Kit board.
Only one Harrier camera interface board at a time should be connected to the Harrier Evaluation Kit Board.

When run without command line arguments, the HarrierControl application will scan the USB ports to identify and communicate with USB/HDMI interface board over USB. 
If it does not find a USB connection it will scan the COM ports, automatically detect the Harrier Evaluation Kit Board and identify the port type (RS-232, RS-485 or TTL) over which the Harrier camera interface board is connected. 

It will then report the camera type, firmware version and display a simple text based menu system which allows a basic set of common setting and query VISCA commands to be sent. 

-----

The Harriercontrol application can be used with command line arguments as follows:

HarrierControl P1 P2 P3 P4

Argument P1:
Interface/Serial communication type.

/?      displays help text. 
/h      displays help text. 
/H      displays help text. 

RS232   RS-232 comms via Harrier Evaluation Kit Board selected.
RS485   RS-485 comms via Harrier Evaluation Kit Board selected.
TTL     TTL comms via Harrier Evaluation Kit Board selected.
USB3    USB comms selected.
COMX    COM port X selected.

Argument P2:
Baud Rate. Not used with USB3 option, instead omit P2 and enter values for P3.
Valid values are 9600, 19200, 38400, 57600 and 115200.

Argument P3:
Comma separated VISCA command to the camera / camera interface board. 
Each value will be treated as a hexadecimal number.
The application requires commands to be entered in a comma delimited format, 
but strips the commas before sending the commands. To work in PowerShell
the comma separated VISCA command should be enclosed in double quotation marks ""

Argument P4:
Optionally append /P to parse the command and response and print a description 
of the command on the screen.
Note: not all commands are supported so there may be no description 
for some commands.

Examples:

C:> HarrierControl TTL 9600 81,09,00,02,ff /P
Tx: [ 81 09 00 02 ff ]
    Get Camera Details
Rx: [ 90 50 00 23 f0 12 00 37 02 ff ]
    Tamron MP1110-VC
    Firmware Version: 0x0037
-----
C:> HarrierControl USB3 "81,09,00,02,ff" /P
Tx: [ 81 09 00 02 ff ]
    Get Camera Details
Rx: [ 90 50 00 23 f0 12 00 37 02 ff ]
    Tamron MP1110-VC
    Firmware Version: 0x0037
-----
C:> HarrierControl
Application Version 1.4.2
Copyright © Active Silicon 2022
Auto-detecting connected device
....
TTL port connected at 9600 Baud
Tamron MP1110-VC
Camera Firmware Version 0x0037
Harrier Firmware Version 1.2.1 
Choose Query Command (Q), Setting Command (S), or press X to exit.

-----

If the application is run with specific COM port and baud rate specified on the command line but no parameters it will look for cameras on that COM port.

e.g. HarrierControl COM6 9600

This enables use of HarrierControl with Harrier cameras that do not have a Camera Interface board. 

===============================

HarrierControl.exe can also be used to read and write USB/HDMI interface board camera firmware/configuration including the Harrier USB serial number.

Command line argument: 
'HarrierControl.exe USB3 GetSerial' reads the serial number (caps sensitive).
'HarrierControl.exe USB3 SetSerial Example123' sets the serial number to "Example123", serial numbers can contain up to 31 numerals and/or ordinary characters.
'HarrierControl.exe USB3 SaveConfig FileName.ext' saves the current USB/HDMI board configuration to "Filename.ext". A filename extension is not mandatory.
'HarrierControl.exe USB3 LoadConfig FileName.ext' loads the previously saved board configuration to the filename specified.
'HarrierControl.exe USB3 LoadConfig FileName.ext SetSerial SerialNumber" loads the configuration and sets the serial number to "SerialNumber"

Writing non-standard characters to the serial number can cause the device to become unreadable and the device may no longer be able to be re-configured (bricked).
Writing anything other than the configuration files provided and serial number generated by Active Silicon batch files will invalidate your warrantee.
The configuration files should only be written to the camera using the batch file: 
Harrier_USB_Cam_Configure.bat
otherwise invalid serial numbers will be configured into the camera.

---

HarrierCOMPortDiscovery.exe
===========================
This command line program scans the COM ports and returns the COM port values for 
the TTL, RS-232 and RS-485 interfaces on the AP23C03 Harrier Evaluation Kit Board.  

---

HCGUI.exe
===========================
This program includes a graphical user interface (GUI) that enables the user to connect to available COM ports.
Once connected, the application can then be used to identify Harrier cameras and Harrier interface boards connected to the COM port.
The GUI offers the user easy control of a range of Harrier camera and Harrier interface board features. 
The HCGUI program is a only a sample application, so there may be issues with some of the features. 
Please report any issues to Tech Support techsupport@activesilicon.com

================================
© Active Silicon Limited 2024
All rights reserved.
================================