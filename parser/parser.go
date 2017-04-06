package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

const KEY_HEADER = "Header"
const KEY_ATTRIBUTES = "Attributes"
const KEY_VALUES = "Values"
const KEY_HANDLE = "Handle"
const KEY_TYPE = "Type"
const KEY_TITLE = "Title"

func ParseDmidecodeTypeKeywords(out string) []string {
	return []string{
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
}

func ParseDmidecodeType(out string) map[string]map[string]map[string]map[string]string {

	result := make(map[string]map[string]map[string]map[string]string)

	codetype, header, key := "", "", ""

	rxp_header := regexp.MustCompile(`^Handle (\w+), DMI type (\d+),`)
	rxp_type := regexp.MustCompile(`^([\w\ ]+)$`)
	rxp_attr := regexp.MustCompile(`^\s+(.+): (.*)$`)
	rxp_subkey := regexp.MustCompile(`^\s+([\w ]+):$`)
	rxp_subattr := regexp.MustCompile(`^\s+(.+) ((is|are) .*)$`)
	rxp_subattr_nokey := regexp.MustCompile(`^\s+(.*)`)
	rpl_attr := strings.NewReplacer(" ", "-")

	read := bytes.NewReader([]byte(out))
	scan := bufio.NewScanner(read)

	headerKey := KEY_HEADER
	attributesKey := KEY_ATTRIBUTES
	valuesKey := KEY_VALUES
	handleKey := KEY_HANDLE
	typeKey := KEY_TYPE
	titleKey := KEY_TITLE

	for scan.Scan() {
		line := scan.Text()
		if line != "" {

			if rxp_header.MatchString(line) {
				header = line
			} else if rxp_type.MatchString(line) && rxp_header.MatchString(header) {
				//
				// Match type
				//
				match := rxp_header.FindStringSubmatch(header)
				codetype = fmt.Sprintf("%s::%s", match[1], match[2])

				result[codetype] = make(map[string]map[string]map[string]string)
				result[codetype][headerKey] = make(map[string]map[string]string)
				result[codetype][attributesKey] = make(map[string]map[string]string)

				result[codetype][headerKey][valuesKey] = make(map[string]string)
				result[codetype][headerKey][valuesKey][handleKey] = match[1]
				result[codetype][headerKey][valuesKey][typeKey] = match[2]
				result[codetype][headerKey][valuesKey][titleKey] = line

			} else if codetype != "" && rxp_attr.MatchString(line) {
				//
				// Match attribute value
				//
				match := rxp_attr.FindStringSubmatch(line)
				key = rpl_attr.Replace(match[1])

				result[codetype][attributesKey][key] = make(map[string]string)

				attributeLength := fmt.Sprint(int(len(result[codetype][attributesKey][key])))

				result[codetype][attributesKey][key][attributeLength] = strings.Trim(match[2], " ")

			} else if codetype != "" && rxp_subkey.MatchString(line) {
				//
				// Match array attribute key
				//
				match := rxp_subkey.FindStringSubmatch(line)
				key = rpl_attr.Replace(match[1])

				result[codetype][attributesKey][key] = make(map[string]string)

			} else if codetype != "" && key != "" && rxp_subattr.MatchString(line) {
				//
				// Match array attribute value
				//
				match := rxp_subattr.FindStringSubmatch(line)
				match[1] = rpl_attr.Replace(match[1])

				//
				// If sub attribute key is exist
				//
				if _, ok := result[codetype][attributesKey][key][match[1]]; ok {
					for i := 1; i < 30; i++ {
						if _, ok := result[codetype][attributesKey][key][fmt.Sprintf("%s-%d", match[1], i)]; !ok {
							match[1] = fmt.Sprintf("%s-%d", match[1], i)
							i = 30 // break for loop
						}
					}
				}

				result[codetype][attributesKey][key][match[1]] = strings.Trim(match[2], " ")

			} else if codetype != "" && key != "" && rxp_subattr_nokey.MatchString(line) {
				attributeLength := fmt.Sprint(int(len(result[codetype][attributesKey][key])))
				result[codetype][attributesKey][key][attributeLength] = strings.Trim(line, "\t")
			}
		}
	}

	return result
}
