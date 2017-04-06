package parser

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func printArray(result interface{}) {
	p, _ := json.MarshalIndent(result, "", " ")
	fmt.Println(string(p))
}

func TestParseValidTypeKeywords_ShouldReturnExpectedArray(t *testing.T) {

	expected := T_expected_ValidTypeKeywords

	out := rawValidTypeKeywords
	result := ParseDmidecodeTypeKeywords(out)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Parse result should match expected array")
	}
}

func TestParseTypeBios_ShouldReturnExpectedArray(t *testing.T) {

	expected := make(map[string]map[string]map[string]map[string]string)
	expected[T_id_BiosInformation] = T_expected_BiosInformation
	expected[T_id_BiosLanguageInformation] = T_expected_BiosLanguageInformation

	out := rawTypeBios
	result := ParseDmidecodeType(out)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Parse result should match expected array")
	}
}

func TestParseTypeSystem_ShouldReturnExpectedArray(t *testing.T) {

	expected := make(map[string]map[string]map[string]map[string]string)
	expected[T_id_SystemInformation] = T_expected_SystemInformation
	expected[T_id_SystemConfigurationOptions] = T_expected_SystemConfigurationOptions
	expected[T_id_SystemBootInformation] = T_expected_SystemBootInformation

	out := rawTypeSystem
	result := ParseDmidecodeType(out)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Parse result should match expected array")
	}
}

func TestParseTypeBaseBoard_ShouldReturnExpectedArray(t *testing.T) {

	expected := make(map[string]map[string]map[string]map[string]string)
	expected[T_id_BaseBoardInformation] = T_expected_BaseBoardInformation
	expected[T_id_OnBoardDeviceInformation] = T_expected_OnBoardDeviceInformation
	expected[T_id_OnBoardDevice_0x003A] = T_expected_OnBoardDevice_0x003A
	expected[T_id_OnBoardDevice_0x003B] = T_expected_OnBoardDevice_0x003B
	expected[T_id_OnBoardDevice_0x003C] = T_expected_OnBoardDevice_0x003C

	out := rawTypeBaseBoard
	result := ParseDmidecodeType(out)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Parse result should match expected array")
	}
}

func TestParseTypeChassis_ShouldReturnExpectedArray(t *testing.T) {

	expected := make(map[string]map[string]map[string]map[string]string)
	expected[T_id_ChassisInformation] = T_expected_ChassisInformation

	out := rawTypeChassis
	result := ParseDmidecodeType(out)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Parse result should match expected array")
	}
}

func TestParseTypeProcessor_ShouldReturnExpectedArray(t *testing.T) {

	expected := make(map[string]map[string]map[string]map[string]string)
	expected[T_id_ProcessorInformation] = T_expected_ProcessorInformation

	out := rawTypeProcessor
	result := ParseDmidecodeType(out)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Parse result should match expected array")
	}
}

func TestParseTypeMemory_ShouldReturnExpectedArray(t *testing.T) {

	expected := make(map[string]map[string]map[string]map[string]string)
	expected[T_id_PhysicalMemoryArray] = T_expected_PhysicalMemoryArray
	expected[T_id_MemoryDevice_0x0042] = T_expected_MemoryDevice_0x0042
	expected[T_id_MemoryDevice_0x0044] = T_expected_MemoryDevice_0x0044
	expected[T_id_MemoryDevice_0x0045] = T_expected_MemoryDevice_0x0045
	expected[T_id_MemoryDevice_0x0047] = T_expected_MemoryDevice_0x0047

	out := rawTypeMemory
	result := ParseDmidecodeType(out)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Parse result should match expected array")
	}
}

func TestParseTypeConnector_ShouldReturnExpectedArray(t *testing.T) {

	expected := make(map[string]map[string]map[string]map[string]string)
	expected[T_id_PortConnectorInformation_0x0004] = T_expected_PortConnectorInformation_0x0004
	expected[T_id_PortConnectorInformation_0x0005] = T_expected_PortConnectorInformation_0x0005

	out := rawTypeConnector
	result := ParseDmidecodeType(out)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Parse result should match expected array")
	}
}

func TestParseTypeSlot_ShouldReturnExpectedArray(t *testing.T) {

	expected := make(map[string]map[string]map[string]map[string]string)
	expected[T_id_SystemSlotInformation_0x001C] = T_expected_SystemSlotInformation_0x0001C
	expected[T_id_SystemSlotInformation_0x001D] = T_expected_SystemSlotInformation_0x0001D

	out := rawTypeSlot
	result := ParseDmidecodeType(out)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Parse result should match expected array")
	}
}
