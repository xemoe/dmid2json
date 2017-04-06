package parser

var T_expected_ValidTypeKeywords = []string{
	"bios",
	"system",
	"baseboard",
	"chassis",
	"processor",
	"memory",
	"cache",
	"connector",
	"slot",
}

var T_id_BiosInformation = "0x0000::0"
var T_expected_BiosInformation = map[string]map[string]map[string]string{
	"Header": map[string]map[string]string{
		"Values": map[string]string{
			"Handle": "0x0000",
			"Type":   "0",
			"Title":  "BIOS Information",
		},
	},
	"Attributes": map[string]map[string]string{
		"Vendor": map[string]string{
			"0": "American Megatrends Inc.",
		},
		"Version": map[string]string{
			"0": "4.6.5",
		},
		"Release-Date": map[string]string{
			"0": "10/08/2012",
		},
		"Address": map[string]string{
			"0": "0xF0000",
		},
		"Runtime-Size": map[string]string{
			"0": "64 kB",
		},
		"ROM-Size": map[string]string{
			"0": "2560 kB",
		},
		"Characteristics": map[string]string{
			"PCI":                    "is supported",
			"BIOS":                   "is upgradeable",
			"BIOS-shadowing":         "is allowed",
			"Boot-from-CD":           "is supported",
			"Selectable-boot":        "is supported",
			"8042-keyboard-services": "are supported (int 9h)",
			"Serial-services":        "are supported (int 14h)",
			"ACPI":                   "is supported",
			"UEFI":                   "is supported",
		},
		"BIOS-Revision": map[string]string{
			"0": "4.6",
		},
	},
}

var T_id_BiosLanguageInformation = "0x004B::13"
var T_expected_BiosLanguageInformation = map[string]map[string]map[string]string{
	"Header": map[string]map[string]string{
		"Values": map[string]string{
			"Handle": "0x004B",
			"Type":   "13",
			"Title":  "BIOS Language Information",
		},
	},
	"Attributes": map[string]map[string]string{
		"Language-Description-Format": map[string]string{
			"0": "Long",
		},
		"Installable-Languages": map[string]string{
			"0": "1",
			"1": "en|US|iso8859-1",
		},
		"Currently-Installed-Language": map[string]string{
			"0": "en|US|iso8859-1",
		},
	},
}

var T_id_SystemInformation = "0x0001::1"
var T_expected_SystemInformation = map[string]map[string]map[string]string{
	"Header": map[string]map[string]string{
		"Values": map[string]string{
			"Handle": "0x0001",
			"Type":   "1",
			"Title":  "System Information",
		},
	},
	"Attributes": map[string]map[string]string{
		"Manufacturer": map[string]string{
			"0": "To be filled by O.E.M.",
		},
		"Product-Name": map[string]string{
			"0": "To be filled by O.E.M.",
		},
		"Version": map[string]string{
			"0": "To be filled by O.E.M.",
		},
		"Serial-Number": map[string]string{
			"0": "To be filled by O.E.M.",
		},
		"UUID": map[string]string{
			"0": "03000200-0400-0500-0006-000700080009",
		},
		"Wake-up-Type": map[string]string{
			"0": "Power Switch",
		},
		"SKU-Number": map[string]string{
			"0": "To be filled by O.E.M.",
		},
		"Family": map[string]string{
			"0": "To be filled by O.E.M.",
		},
	},
}

var T_id_SystemConfigurationOptions = "0x0023::12"
var T_expected_SystemConfigurationOptions = map[string]map[string]map[string]string{
	"Header": map[string]map[string]string{
		"Values": map[string]string{
			"Handle": "0x0023",
			"Type":   "12",
			"Title":  "System Configuration Options",
		},
	},
	"Attributes": map[string]map[string]string{
		"Option-1": map[string]string{
			"0": "To Be Filled By O.E.M.",
		},
	},
}

var T_id_SystemBootInformation = "0x0024::32"
var T_expected_SystemBootInformation = map[string]map[string]map[string]string{
	"Header": map[string]map[string]string{
		"Values": map[string]string{
			"Handle": "0x0024",
			"Type":   "32",
			"Title":  "System Boot Information",
		},
	},
	"Attributes": map[string]map[string]string{
		"Status": map[string]string{
			"0": "No errors detected",
		},
	},
}

var T_id_BaseBoardInformation = "0x0002::2"
var T_expected_BaseBoardInformation = map[string]map[string]map[string]string{
	"Header": map[string]map[string]string{
		"Values": map[string]string{
			"Handle": "0x0002",
			"Type":   "2",
			"Title":  "Base Board Information",
		},
	},
	"Attributes": map[string]map[string]string{
		"Manufacturer": map[string]string{
			"0": "INTEL Corporation",
		},
		"Product-Name": map[string]string{
			"0": "MAHOBAY",
		},
		"Version": map[string]string{
			"0": "To be filled by O.E.M.",
		},
		"Serial-Number": map[string]string{
			"0": "To be filled by O.E.M.",
		},
		"Asset-Tag": map[string]string{
			"0": "To be filled by O.E.M.",
		},
		"Features": map[string]string{
			"Board":   "is a hosting board",
			"Board-1": "is replaceable",
		},
		"Location-In-Chassis": map[string]string{
			"0": "To be filled by O.E.M.",
		},
		"Chassis-Handle": map[string]string{
			"0": "0x0003",
		},
		"Type": map[string]string{
			"0": "Motherboard",
		},
		"Contained-Object-Handles": map[string]string{
			"0": "0",
		},
	},
}

var T_id_OnBoardDeviceInformation = "0x0021::10"
var T_expected_OnBoardDeviceInformation = map[string]map[string]map[string]string{
	"Header": map[string]map[string]string{
		"Values": map[string]string{
			"Handle": "0x0021",
			"Type":   "10",
			"Title":  "On Board Device Information",
		},
	},
	"Attributes": map[string]map[string]string{
		"Type": map[string]string{
			"0": "Video",
		},
		"Status": map[string]string{
			"0": "Enabled",
		},
		"Description": map[string]string{
			"0": "To Be Filled By O.E.M.",
		},
	},
}

var T_id_OnBoardDevice_0x003A = "0x003A::41"
var T_expected_OnBoardDevice_0x003A = map[string]map[string]map[string]string{
	"Header": map[string]map[string]string{
		"Values": map[string]string{
			"Handle": "0x003A",
			"Type":   "41",
			"Title":  "Onboard Device",
		},
	},
	"Attributes": map[string]map[string]string{
		"Reference-Designation": map[string]string{
			"0": "Onboard IGD",
		},
		"Type": map[string]string{
			"0": "Video",
		},
		"Status": map[string]string{
			"0": "Enabled",
		},
		"Type-Instance": map[string]string{
			"0": "1",
		},
		"Bus-Address": map[string]string{
			"0": "0000:00:02.0",
		},
	},
}

var T_id_OnBoardDevice_0x003B = "0x003B::41"
var T_expected_OnBoardDevice_0x003B = map[string]map[string]map[string]string{
	"Header": map[string]map[string]string{
		"Values": map[string]string{
			"Handle": "0x003B",
			"Type":   "41",
			"Title":  "Onboard Device",
		},
	},
	"Attributes": map[string]map[string]string{
		"Reference-Designation": map[string]string{
			"0": "Onboard LAN",
		},
		"Type": map[string]string{
			"0": "Ethernet",
		},
		"Status": map[string]string{
			"0": "Enabled",
		},
		"Type-Instance": map[string]string{
			"0": "1",
		},
		"Bus-Address": map[string]string{
			"0": "0000:00:19.0",
		},
	},
}

var T_id_OnBoardDevice_0x003C = "0x003C::41"
var T_expected_OnBoardDevice_0x003C = map[string]map[string]map[string]string{
	"Header": map[string]map[string]string{
		"Values": map[string]string{
			"Handle": "0x003C",
			"Type":   "41",
			"Title":  "Onboard Device",
		},
	},
	"Attributes": map[string]map[string]string{
		"Reference-Designation": map[string]string{
			"0": "Onboard 1394",
		},
		"Type": map[string]string{
			"0": "Other",
		},
		"Status": map[string]string{
			"0": "Enabled",
		},
		"Type-Instance": map[string]string{
			"0": "1",
		},
		"Bus-Address": map[string]string{
			"0": "0000:03:1c.2",
		},
	},
}

var T_id_ChassisInformation = "0x0003::3"
var T_expected_ChassisInformation = map[string]map[string]map[string]string{
	"Header": map[string]map[string]string{
		"Values": map[string]string{
			"Handle": "0x0003",
			"Type":   "3",
			"Title":  "Chassis Information",
		},
	},
	"Attributes": map[string]map[string]string{
		"Manufacturer": map[string]string{
			"0": "To Be Filled By O.E.M.",
		},
		"Type": map[string]string{
			"0": "Desktop",
		},
		"Lock": map[string]string{
			"0": "Not Present",
		},
		"Version": map[string]string{
			"0": "To Be Filled By O.E.M.",
		},
		"Serial-Number": map[string]string{
			"0": "To Be Filled By O.E.M.",
		},
		"Asset-Tag": map[string]string{
			"0": "To Be Filled By O.E.M.",
		},
		"Boot-up-State": map[string]string{
			"0": "Safe",
		},
		"Power-Supply-State": map[string]string{
			"0": "Safe",
		},
		"Thermal-State": map[string]string{
			"0": "Safe",
		},
		"Security-Status": map[string]string{
			"0": "None",
		},
		"OEM-Information": map[string]string{
			"0": "0x00000000",
		},
		"Height": map[string]string{
			"0": "Unspecified",
		},
		"Number-Of-Power-Cords": map[string]string{
			"0": "1",
		},
		"Contained-Elements": map[string]string{
			"0": "0",
		},
		"SKU-Number": map[string]string{
			"0": "To be filled by O.E.M.",
		},
	},
}

var T_id_ProcessorInformation = "0x0041::4"
var T_expected_ProcessorInformation = map[string]map[string]map[string]string{
	"Header": map[string]map[string]string{
		"Values": map[string]string{
			"Handle": "0x0041",
			"Type":   "4",
			"Title":  "Processor Information",
		},
	},
	"Attributes": map[string]map[string]string{
		"Socket-Designation": map[string]string{
			"0": "SOCKET 0",
		},
		"Type": map[string]string{
			"0": "Central Processor",
		},
		"Family": map[string]string{
			"0": "Xeon",
		},
		"Manufacturer": map[string]string{
			"0": "Intel(R) Corporation",
		},
		"ID": map[string]string{
			"0": "FF FF FF FF FF FF FF FF",
		},
		"Signature": map[string]string{
			"0": "Type 0, Family 6, Model 42, Stepping 7",
		},
		"Flags": map[string]string{
			"0": "FPU (Floating-point unit on-chip)",
			"1": "VME (Virtual mode extension)",
			"2": "PSE-36 (36-bit page size extension)",
		},
		"Version": map[string]string{
			"0": "Intel(R) Xeon(R) CPU E31225 @ 3.10GHz",
		},
		"Voltage": map[string]string{
			"0": "5.0 V",
		},
		"Status": map[string]string{
			"0": "Populated, Enabled",
		},
		"Serial-Number": map[string]string{
			"0": "Not Specified",
		},
		"Asset-Tag": map[string]string{
			"0": "Fill By OEM",
		},
		"Characteristics": map[string]string{
			"0": "64-bit capable",
		},
	},
}

var T_id_PhysicalMemoryArray = "0x0040::16"
var T_expected_PhysicalMemoryArray = map[string]map[string]map[string]string{
	"Header": map[string]map[string]string{
		"Values": map[string]string{
			"Handle": "0x0040",
			"Type":   "16",
			"Title":  "Physical Memory Array",
		},
	},
	"Attributes": map[string]map[string]string{
		"Location": map[string]string{
			"0": "System Board Or Motherboard",
		},
		"Use": map[string]string{
			"0": "System Memory",
		},
		"Error-Correction-Type": map[string]string{
			"0": "None",
		},
		"Maximum-Capacity": map[string]string{
			"0": "32 GB",
		},
		"Error-Information-Handle": map[string]string{
			"0": "Not Provided",
		},
		"Number-Of-Devices": map[string]string{
			"0": "4",
		},
	},
}

var T_id_MemoryDevice_0x0042 = "0x0042::17"
var T_expected_MemoryDevice_0x0042 = map[string]map[string]map[string]string{
	"Header": map[string]map[string]string{
		"Values": map[string]string{
			"Handle": "0x0042",
			"Type":   "17",
			"Title":  "Memory Device",
		},
	},
	"Attributes": map[string]map[string]string{
		"Array-Handle": map[string]string{
			"0": "0x0040",
		},
		"Error-Information-Handle": map[string]string{
			"0": "Not Provided",
		},
		"Total-Width": map[string]string{
			"0": "64 bits",
		},
		"Data-Width": map[string]string{
			"0": "64 bits",
		},
		"Size": map[string]string{
			"0": "8192 MB",
		},
		"Form-Factor": map[string]string{
			"0": "DIMM",
		},
		"Set": map[string]string{
			"0": "None",
		},
		"Locator": map[string]string{
			"0": "ChannelA-DIMM0",
		},
		"Bank-Locator": map[string]string{
			"0": "BANK 0",
		},
		"Type": map[string]string{
			"0": "DDR3",
		},
		"Type-Detail": map[string]string{
			"0": "Synchronous",
		},
		"Speed": map[string]string{
			"0": "1333 MHz",
		},
		"Manufacturer": map[string]string{
			"0": "Kingston",
		},
		"Serial-Number": map[string]string{
			"0": "FFFFFFFF",
		},
		"Asset-Tag": map[string]string{
			"0": "FFFFFFFFFF",
		},
		"Part-Number": map[string]string{
			"0": "FFFF",
		},
		"Rank": map[string]string{
			"0": "2",
		},
		"Configured-Clock-Speed": map[string]string{
			"0": "1333 MHz",
		},
	},
}

var T_id_MemoryDevice_0x0044 = "0x0044::17"
var T_expected_MemoryDevice_0x0044 = map[string]map[string]map[string]string{
	"Header": map[string]map[string]string{
		"Values": map[string]string{
			"Handle": "0x0044",
			"Type":   "17",
			"Title":  "Memory Device",
		},
	},
	"Attributes": map[string]map[string]string{
		"Array-Handle": map[string]string{
			"0": "0x0040",
		},
		"Error-Information-Handle": map[string]string{
			"0": "Not Provided",
		},
		"Total-Width": map[string]string{
			"0": "Unknown",
		},
		"Data-Width": map[string]string{
			"0": "Unknown",
		},
		"Size": map[string]string{
			"0": "No Module Installed",
		},
		"Form-Factor": map[string]string{
			"0": "DIMM",
		},
		"Set": map[string]string{
			"0": "None",
		},
		"Locator": map[string]string{
			"0": "ChannelA-DIMM1",
		},
		"Bank-Locator": map[string]string{
			"0": "BANK 1",
		},
		"Type": map[string]string{
			"0": "Unknown",
		},
		"Type-Detail": map[string]string{
			"0": "None",
		},
		"Speed": map[string]string{
			"0": "Unknown",
		},
		"Manufacturer": map[string]string{
			"0": "[Empty]",
		},
		"Serial-Number": map[string]string{
			"0": "[Empty]",
		},
		"Asset-Tag": map[string]string{
			"0": "FFFFFFFFFF",
		},
		"Part-Number": map[string]string{
			"0": "[Empty]",
		},
		"Rank": map[string]string{
			"0": "Unknown",
		},
		"Configured-Clock-Speed": map[string]string{
			"0": "Unknown",
		},
	},
}

var T_id_MemoryDevice_0x0045 = "0x0045::17"
var T_expected_MemoryDevice_0x0045 = map[string]map[string]map[string]string{
	"Header": map[string]map[string]string{
		"Values": map[string]string{
			"Handle": "0x0045",
			"Type":   "17",
			"Title":  "Memory Device",
		},
	},
	"Attributes": map[string]map[string]string{
		"Array-Handle": map[string]string{
			"0": "0x0040",
		},
		"Error-Information-Handle": map[string]string{
			"0": "Not Provided",
		},
		"Total-Width": map[string]string{
			"0": "64 bits",
		},
		"Data-Width": map[string]string{
			"0": "64 bits",
		},
		"Size": map[string]string{
			"0": "8192 MB",
		},
		"Form-Factor": map[string]string{
			"0": "DIMM",
		},
		"Set": map[string]string{
			"0": "None",
		},
		"Locator": map[string]string{
			"0": "ChannelB-DIMM0",
		},
		"Bank-Locator": map[string]string{
			"0": "BANK 2",
		},
		"Type": map[string]string{
			"0": "DDR3",
		},
		"Type-Detail": map[string]string{
			"0": "Synchronous",
		},
		"Speed": map[string]string{
			"0": "1333 MHz",
		},
		"Manufacturer": map[string]string{
			"0": "Kingston",
		},
		"Serial-Number": map[string]string{
			"0": "FFFFFFFF",
		},
		"Asset-Tag": map[string]string{
			"0": "FFFFFFFFFF",
		},
		"Part-Number": map[string]string{
			"0": "FFFF",
		},
		"Rank": map[string]string{
			"0": "2",
		},
		"Configured-Clock-Speed": map[string]string{
			"0": "1333 MHz",
		},
	},
}

var T_id_MemoryDevice_0x0047 = "0x0047::17"
var T_expected_MemoryDevice_0x0047 = map[string]map[string]map[string]string{
	"Header": map[string]map[string]string{
		"Values": map[string]string{
			"Handle": "0x0047",
			"Type":   "17",
			"Title":  "Memory Device",
		},
	},
	"Attributes": map[string]map[string]string{
		"Array-Handle": map[string]string{
			"0": "0x0040",
		},
		"Error-Information-Handle": map[string]string{
			"0": "Not Provided",
		},
		"Total-Width": map[string]string{
			"0": "Unknown",
		},
		"Data-Width": map[string]string{
			"0": "Unknown",
		},
		"Size": map[string]string{
			"0": "No Module Installed",
		},
		"Form-Factor": map[string]string{
			"0": "DIMM",
		},
		"Set": map[string]string{
			"0": "None",
		},
		"Locator": map[string]string{
			"0": "ChannelB-DIMM1",
		},
		"Bank-Locator": map[string]string{
			"0": "BANK 3",
		},
		"Type": map[string]string{
			"0": "Unknown",
		},
		"Type-Detail": map[string]string{
			"0": "None",
		},
		"Speed": map[string]string{
			"0": "Unknown",
		},
		"Manufacturer": map[string]string{
			"0": "[Empty]",
		},
		"Serial-Number": map[string]string{
			"0": "[Empty]",
		},
		"Asset-Tag": map[string]string{
			"0": "FFFFFFFFFF",
		},
		"Part-Number": map[string]string{
			"0": "[Empty]",
		},
		"Rank": map[string]string{
			"0": "Unknown",
		},
		"Configured-Clock-Speed": map[string]string{
			"0": "Unknown",
		},
	},
}

var T_id_PortConnectorInformation_0x0004 = "0x0004::8"
var T_expected_PortConnectorInformation_0x0004 = map[string]map[string]map[string]string{
	"Header": map[string]map[string]string{
		"Values": map[string]string{
			"Handle": "0x0004",
			"Type":   "8",
			"Title":  "Port Connector Information",
		},
	},
	"Attributes": map[string]map[string]string{
		"Internal-Reference-Designator": map[string]string{
			"0": "J1A1",
		},
		"Internal-Connector-Type": map[string]string{
			"0": "None",
		},
		"External-Reference-Designator": map[string]string{
			"0": "PS2Mouse",
		},
		"External-Connector-Type": map[string]string{
			"0": "PS/2",
		},
		"Port-Type": map[string]string{
			"0": "Mouse Port",
		},
	},
}

var T_id_PortConnectorInformation_0x0005 = "0x0005::8"
var T_expected_PortConnectorInformation_0x0005 = map[string]map[string]map[string]string{
	"Header": map[string]map[string]string{
		"Values": map[string]string{
			"Handle": "0x0005",
			"Type":   "8",
			"Title":  "Port Connector Information",
		},
	},
	"Attributes": map[string]map[string]string{
		"Internal-Reference-Designator": map[string]string{
			"0": "J1A1",
		},
		"Internal-Connector-Type": map[string]string{
			"0": "None",
		},
		"External-Reference-Designator": map[string]string{
			"0": "Keyboard",
		},
		"External-Connector-Type": map[string]string{
			"0": "PS/2",
		},
		"Port-Type": map[string]string{
			"0": "Keyboard Port",
		},
	},
}

var T_id_SystemSlotInformation_0x001C = "0x001C::9"
var T_expected_SystemSlotInformation_0x0001C = map[string]map[string]map[string]string{
	"Header": map[string]map[string]string{
		"Values": map[string]string{
			"Handle": "0x001C",
			"Type":   "9",
			"Title":  "System Slot Information",
		},
	},
	"Attributes": map[string]map[string]string{
		"Designation": map[string]string{
			"0": "J6B2",
		},
		"Type": map[string]string{
			"0": "x16 PCI Express",
		},
		"Current-Usage": map[string]string{
			"0": "In Use",
		},
		"Length": map[string]string{
			"0": "Long",
		},
		"ID": map[string]string{
			"0": "0",
		},
		"Bus-Address": map[string]string{
			"0": "0000:00:01.0",
		},
		"Characteristics": map[string]string{
			"3.3-V":      "is provided",
			"Opening":    "is shared",
			"PME-signal": "is supported",
		},
	},
}

var T_id_SystemSlotInformation_0x001D = "0x001D::9"
var T_expected_SystemSlotInformation_0x0001D = map[string]map[string]map[string]string{
	"Header": map[string]map[string]string{
		"Values": map[string]string{
			"Handle": "0x001D",
			"Type":   "9",
			"Title":  "System Slot Information",
		},
	},
	"Attributes": map[string]map[string]string{
		"Designation": map[string]string{
			"0": "J6B1",
		},
		"Type": map[string]string{
			"0": "x1 PCI Express",
		},
		"Current-Usage": map[string]string{
			"0": "In Use",
		},
		"Length": map[string]string{
			"0": "Short",
		},
		"ID": map[string]string{
			"0": "1",
		},
		"Bus-Address": map[string]string{
			"0": "0000:00:1c.3",
		},
		"Characteristics": map[string]string{
			"3.3-V":      "is provided",
			"Opening":    "is shared",
			"PME-signal": "is supported",
		},
	},
}
