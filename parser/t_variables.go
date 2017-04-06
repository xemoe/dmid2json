package parser

var rawValidTypeKeywords string = `
dmidecode: option requires an argument -- 't'
Type number or keyword expected
Valid type keywords are:
  bios
  system
  baseboard
  chassis
  processor
  memory
  cache
  connector
  slot
`
var rawTypeBios string = `
# dmidecode 2.12
SMBIOS 2.7 present.

Handle 0x0000, DMI type 0, 24 bytes
BIOS Information
	Vendor: American Megatrends Inc.
	Version: 4.6.5
	Release Date: 10/08/2012
	Address: 0xF0000
	Runtime Size: 64 kB
	ROM Size: 2560 kB
	Characteristics:
		PCI is supported
		BIOS is upgradeable
		BIOS shadowing is allowed
		Boot from CD is supported
		Selectable boot is supported
		8042 keyboard services are supported (int 9h)
		Serial services are supported (int 14h)
		ACPI is supported
		UEFI is supported
	BIOS Revision: 4.6

Handle 0x004B, DMI type 13, 22 bytes
BIOS Language Information
	Language Description Format: Long
	Installable Languages: 1
		en|US|iso8859-1
	Currently Installed Language: en|US|iso8859-1
`

var rawTypeSystem string = `
# dmidecode 2.12
SMBIOS 2.7 present.

Handle 0x0001, DMI type 1, 27 bytes
System Information
	Manufacturer: To be filled by O.E.M.
	Product Name: To be filled by O.E.M.
	Version: To be filled by O.E.M.
	Serial Number: To be filled by O.E.M.
	UUID: 03000200-0400-0500-0006-000700080009
	Wake-up Type: Power Switch
	SKU Number: To be filled by O.E.M.
	Family: To be filled by O.E.M.

Handle 0x0023, DMI type 12, 5 bytes
System Configuration Options
	Option 1: To Be Filled By O.E.M.

Handle 0x0024, DMI type 32, 20 bytes
System Boot Information
	Status: No errors detected
`

var rawTypeBaseBoard string = `
# dmidecode 2.12
SMBIOS 2.7 present.

Handle 0x0002, DMI type 2, 15 bytes
Base Board Information
	Manufacturer: INTEL Corporation
	Product Name: MAHOBAY
	Version: To be filled by O.E.M.
	Serial Number: To be filled by O.E.M.
	Asset Tag: To be filled by O.E.M.
	Features:
		Board is a hosting board
		Board is replaceable
	Location In Chassis: To be filled by O.E.M.
	Chassis Handle: 0x0003
	Type: Motherboard
	Contained Object Handles: 0

Handle 0x0021, DMI type 10, 6 bytes
On Board Device Information
	Type: Video
	Status: Enabled
	Description:    To Be Filled By O.E.M.

Handle 0x003A, DMI type 41, 11 bytes
Onboard Device
	Reference Designation:  Onboard IGD
	Type: Video
	Status: Enabled
	Type Instance: 1
	Bus Address: 0000:00:02.0

Handle 0x003B, DMI type 41, 11 bytes
Onboard Device
	Reference Designation:  Onboard LAN
	Type: Ethernet
	Status: Enabled
	Type Instance: 1
	Bus Address: 0000:00:19.0

Handle 0x003C, DMI type 41, 11 bytes
Onboard Device
	Reference Designation:  Onboard 1394
	Type: Other
	Status: Enabled
	Type Instance: 1
	Bus Address: 0000:03:1c.2
`

var rawTypeChassis string = `
# dmidecode 2.12
SMBIOS 2.7 present.

Handle 0x0003, DMI type 3, 22 bytes
Chassis Information
	Manufacturer: To Be Filled By O.E.M.
	Type: Desktop
	Lock: Not Present
	Version: To Be Filled By O.E.M.
	Serial Number: To Be Filled By O.E.M.
	Asset Tag: To Be Filled By O.E.M.
	Boot-up State: Safe
	Power Supply State: Safe
	Thermal State: Safe
	Security Status: None
	OEM Information: 0x00000000
	Height: Unspecified
	Number Of Power Cords: 1
	Contained Elements: 0
	SKU Number: To be filled by O.E.M.
`

var rawTypeProcessor string = `
# dmidecode 2.12
SMBIOS 2.7 present.

Handle 0x0041, DMI type 4, 42 bytes
Processor Information
	Socket Designation: SOCKET 0
	Type: Central Processor
	Family: Xeon
	Manufacturer: Intel(R) Corporation
	ID: FF FF FF FF FF FF FF FF
	Signature: Type 0, Family 6, Model 42, Stepping 7
	Flags:
		FPU (Floating-point unit on-chip)
		VME (Virtual mode extension)
		PSE-36 (36-bit page size extension)
	Version: Intel(R) Xeon(R) CPU E31225 @ 3.10GHz
	Voltage: 5.0 V
	Status: Populated, Enabled
	Serial Number: Not Specified
	Asset Tag: Fill By OEM
	Characteristics:
		64-bit capable
`

var rawTypeMemory string = `
# dmidecode 2.12
SMBIOS 2.7 present.

Handle 0x0040, DMI type 16, 23 bytes
Physical Memory Array
	Location: System Board Or Motherboard
	Use: System Memory
	Error Correction Type: None
	Maximum Capacity: 32 GB
	Error Information Handle: Not Provided
	Number Of Devices: 4

Handle 0x0042, DMI type 17, 34 bytes
Memory Device
	Array Handle: 0x0040
	Error Information Handle: Not Provided
	Total Width: 64 bits
	Data Width: 64 bits
	Size: 8192 MB
	Form Factor: DIMM
	Set: None
	Locator: ChannelA-DIMM0
	Bank Locator: BANK 0
	Type: DDR3
	Type Detail: Synchronous
	Speed: 1333 MHz
	Manufacturer: Kingston
	Serial Number: FFFFFFFF
	Asset Tag: FFFFFFFFFF
	Part Number: FFFF
	Rank: 2
	Configured Clock Speed: 1333 MHz

Handle 0x0044, DMI type 17, 34 bytes
Memory Device
	Array Handle: 0x0040
	Error Information Handle: Not Provided
	Total Width: Unknown
	Data Width: Unknown
	Size: No Module Installed
	Form Factor: DIMM
	Set: None
	Locator: ChannelA-DIMM1
	Bank Locator: BANK 1
	Type: Unknown
	Type Detail: None
	Speed: Unknown
	Manufacturer: [Empty]
	Serial Number: [Empty]
	Asset Tag: FFFFFFFFFF
	Part Number: [Empty]
	Rank: Unknown
	Configured Clock Speed: Unknown

Handle 0x0045, DMI type 17, 34 bytes
Memory Device
	Array Handle: 0x0040
	Error Information Handle: Not Provided
	Total Width: 64 bits
	Data Width: 64 bits
	Size: 8192 MB
	Form Factor: DIMM
	Set: None
	Locator: ChannelB-DIMM0
	Bank Locator: BANK 2
	Type: DDR3
	Type Detail: Synchronous
	Speed: 1333 MHz
	Manufacturer: Kingston
	Serial Number: FFFFFFFF
	Asset Tag: FFFFFFFFFF
	Part Number: FFFF
	Rank: 2
	Configured Clock Speed: 1333 MHz

Handle 0x0047, DMI type 17, 34 bytes
Memory Device
	Array Handle: 0x0040
	Error Information Handle: Not Provided
	Total Width: Unknown
	Data Width: Unknown
	Size: No Module Installed
	Form Factor: DIMM
	Set: None
	Locator: ChannelB-DIMM1
	Bank Locator: BANK 3
	Type: Unknown
	Type Detail: None
	Speed: Unknown
	Manufacturer: [Empty]
	Serial Number: [Empty]
	Asset Tag: FFFFFFFFFF
	Part Number: [Empty]
	Rank: Unknown
	Configured Clock Speed: Unknown
`

var rawTypeCache string = `
# dmidecode 2.12
SMBIOS 2.7 present.

Handle 0x003E, DMI type 7, 19 bytes
Cache Information
	Socket Designation: CPU Internal L1
	Configuration: Enabled, Not Socketed, Level 1
	Operational Mode: Write Through
	Location: Internal
	Installed Size: 256 kB
	Maximum Size: 256 kB
	Supported SRAM Types:
		Unknown
	Installed SRAM Type: Unknown
	Speed: Unknown
	Error Correction Type: Parity
	System Type: Other
	Associativity: 16-way Set-associative
`

var rawTypeConnector string = `
# dmidecode 2.12
SMBIOS 2.7 present.

Handle 0x0004, DMI type 8, 9 bytes
Port Connector Information
	Internal Reference Designator: J1A1
	Internal Connector Type: None
	External Reference Designator: PS2Mouse
	External Connector Type: PS/2
	Port Type: Mouse Port

Handle 0x0005, DMI type 8, 9 bytes
Port Connector Information
	Internal Reference Designator: J1A1
	Internal Connector Type: None
	External Reference Designator: Keyboard
	External Connector Type: PS/2
	Port Type: Keyboard Port
`

var rawTypeSlot string = `
# dmidecode 2.12
SMBIOS 2.7 present.

Handle 0x001C, DMI type 9, 17 bytes
System Slot Information
	Designation: J6B2
	Type: x16 PCI Express
	Current Usage: In Use
	Length: Long
	ID: 0
	Characteristics:
		3.3 V is provided
		Opening is shared
		PME signal is supported
	Bus Address: 0000:00:01.0

Handle 0x001D, DMI type 9, 17 bytes
System Slot Information
	Designation: J6B1
	Type: x1 PCI Express
	Current Usage: In Use
	Length: Short
	ID: 1
	Characteristics:
		3.3 V is provided
		Opening is shared
		PME signal is supported
	Bus Address: 0000:00:1c.3
`
