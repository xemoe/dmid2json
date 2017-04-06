package dmid

import (
	"reflect"
	"testing"
)

var expectedBiosString string = `# dmidecode 2.12
Reading SMBIOS/DMI data from file dmitable.bin.
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
		BIOS ROM is socketed
		EDD is supported
		5.25"/1.2 MB floppy services are supported (int 13h)
		3.5"/720 kB floppy services are supported (int 13h)
		3.5"/2.88 MB floppy services are supported (int 13h)
		Print screen service is supported (int 5h)
		8042 keyboard services are supported (int 9h)
		Serial services are supported (int 14h)
		Printer services are supported (int 17h)
		ACPI is supported
		USB legacy is supported
		BIOS boot specification is supported
		Targeted content distribution is supported
		UEFI is supported
	BIOS Revision: 4.6

Handle 0x004B, DMI type 13, 22 bytes
BIOS Language Information
	Language Description Format: Long
	Installable Languages: 1
		en|US|iso8859-1
	Currently Installed Language: en|US|iso8859-1`

func TestRawDecode_ExpectedStdout_ShoudlNotBeNull(t *testing.T) {
	stdout, _, _ := RawDecode([]string{"-t", "0"})
	if len(stdout) == 0 {
		t.Errorf("RawDecode result shouldn't be empty string or return error")
	}
}

func TestRawDecode_ExpectedStderr_ShouldNotBeNull(t *testing.T) {
	_, _, stderr := RawDecode([]string{"-t"})
	if len(stderr) == 0 {
		t.Errorf("RawDecode result shouldn't be empty string or return error")
	}
}

func TestStringDecode_ExpectedStdout_ShouldNotBeNull(t *testing.T) {
	stdout := StringDecode("bios-vendor")
	if len(stdout) == 0 {
		t.Errorf("StringDecode result shouldn't be empty string or return error")
	}
}

func TestTypeDecode_ExpectedStdout_ShouldNotBeNull(t *testing.T) {
	stdout := TypeDecode("0")
	if len(stdout) == 0 {
		t.Errorf("TypeDecode result shouldn't be empty string or return error")
	}
}

func TestAvailableStringKeywords_ExpectedKeywords_ShouldBeExpectedArray(t *testing.T) {

	expected := []string{
		"bios-vendor",
		"bios-version",
		"bios-release-date",
		"system-manufacturer",
		"system-product-name",
		"system-version",
		"system-serial-number",
		"system-uuid",
		"baseboard-manufacturer",
		"baseboard-product-name",
		"baseboard-version",
		"baseboard-serial-number",
		"baseboard-asset-tag",
		"chassis-manufacturer",
		"chassis-type",
		"chassis-version",
		"chassis-serial-number",
		"chassis-asset-tag",
		"processor-family",
		"processor-manufacturer",
		"processor-version",
		"processor-frequency",
	}

	keywords := AvailableStringKeywords()

	if len(keywords) == 0 {
		t.Errorf("Keywords shouldn't be empty")
	}
	if !reflect.DeepEqual(expected, keywords) {
		t.Errorf("Keywords should match expected array")
	}
}

func TestAvailableTypeKeywords_ExpectedKeywords_ShouldBeExpectedArray(t *testing.T) {

	expected := []string{
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

	keywords := AvailableTypeKeywords()

	if len(keywords) == 0 {
		t.Errorf("Keywords shouldn't be empty")
	}
	if !reflect.DeepEqual(expected, keywords) {
		t.Errorf("Keywords should match expected array")
	}
}

/**
func TestReadFromBin_ShouldReturnExpectedString(t *testing.T) {

	expected := expectedBiosString

	tablefile := "dmitable.bin"

	out := ReadFromBin(tablefile, "bios")

	if expected != out {
		t.Errorf("Result should match expected string")
	}
}
**/
