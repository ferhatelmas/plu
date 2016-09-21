package plu

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type codeMap map[string]string

var validRe = regexp.MustCompile(`^[89]?[34]\d{3}$`)

func (c codeMap) Valid(i interface{}) bool {
	return validRe.MatchString(fmt.Sprint(i))
}

func (c codeMap) Name(i interface{}) string {
	if !c.Valid(i) {
		return ""
	}
	s := fmt.Sprint(i)
	if len(s) == 5 {
		s = s[1:]
	}
	return c[s]
}

func (c codeMap) Organic(i interface{}) bool {
	return strings.HasPrefix(fmt.Sprint(i), "9")
}

func (c codeMap) GM(i interface{}) bool {
	return strings.HasPrefix(fmt.Sprint(i), "8")
}

func (c codeMap) RetailerAssigned(i interface{}) bool {
	return strings.HasPrefix(c.Name(i), "Retailer Assigned")
}

func (c codeMap) All() map[string]string {
	m := map[string]string{}
	for k, v := range m {
		m[k] = v
	}
	return m
}

var codes codeMap

func readCodes() error {
	f, err := os.Open("plu_codes.csv")
	if err != nil {
		return err
	}
	r := csv.NewReader(f)

	codes = map[string]string{}
	for i := 0; ; i++ {
		record, err := r.Read()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
		codes[record[0]] = record[1]
	}
}

// Codes provides easy access into Price Lookup Code Mapping
type Codes interface {
	Valid(interface{}) bool
	Organic(interface{}) bool
	GM(interface{}) bool
	RetailerAssigned(interface{}) bool
	Name(interface{}) string
	All() map[string]string
}

// New reads codes if first call and then creates an accessor for codes
func New() (Codes, error) {
	if codes == nil {
		if err := readCodes(); err != nil {
			return nil, err
		}
	}
	return codes, nil
}
