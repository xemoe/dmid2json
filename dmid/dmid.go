package dmid

import (
	"bufio"
	"bytes"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

type DmiArg struct {
	StringKeyword []string
	TypeKeyword   []string
}

func PredefinedArg() DmiArg {
	return DmiArg{
		StringKeyword: []string{"-s"},
		TypeKeyword:   []string{"-t"},
	}
}

func RawDecode(arg []string) (string, error, string) {

	cmd := exec.Command("dmidecode", arg...)

	var sout, serr bytes.Buffer

	cmd.Stdout = &sout
	cmd.Stderr = &serr

	err := cmd.Run()

	return sout.String(), err, serr.String()
}

func StringDecode(kw string) string {

	arg := PredefinedArg().StringKeyword
	arg = append(arg, kw)

	output, err, _ := RawDecode(arg)
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(output)
}

func TypeDecode(kw string) (output string) {

	arg := PredefinedArg().TypeKeyword
	arg = append(arg, kw)

	sout, err, _ := RawDecode(arg)

	if err != nil {
		log.Fatal(err)
	}

	read := bytes.NewReader([]byte(sout))
	scan := bufio.NewScanner(read)

	for scan.Scan() {
		line := scan.Text()
		if line != "" {
			if len(output) > 0 {
				output = output + "\n" + line
			} else {
				output = line
			}
		}
	}

	return
}

func AvailableStringKeywords() (keywords []string) {

	rxp := regexp.MustCompile(`^(([a-z]+\-)+[a-z]+)$`)
	arg := PredefinedArg().StringKeyword

	_, _, serr := RawDecode(arg)

	read := bytes.NewReader([]byte(serr))
	scan := bufio.NewScanner(read)

	for scan.Scan() {
		line := strings.TrimSpace(scan.Text())
		if line != "" && rxp.MatchString(line) {
			keywords = append(keywords, line)
		}
	}

	return
}

func AvailableTypeKeywords() (keywords []string) {

	rxp := regexp.MustCompile(`^([a-z]+)$`)
	arg := PredefinedArg().TypeKeyword

	_, _, serr := RawDecode(arg)

	read := bytes.NewReader([]byte(serr))
	scan := bufio.NewScanner(read)

	for scan.Scan() {
		line := strings.TrimSpace(scan.Text())
		if line != "" && rxp.MatchString(line) {
			keywords = append(keywords, line)
		}
	}

	return
}

func ReadFromBin(file string, keyword string) string {

	options := []string{"--from-dump", file, "--type", keyword}
	output, err, _ := RawDecode(options)
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(output)
}
