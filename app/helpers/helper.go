package helpers

import (
	"fmt"
	"strings"
	"time"
)

func GetMonth(m time.Month) string {
	months := []string{"jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec"}
	s := m.String()
	first3 := s[0:3]
	date := strings.ToLower(first3)

	var month []string
	for _, v := range months {
		if date == v {
			month = append(month, v)
		}
	}
	strDate := strings.Join(month, ", ")
	return strDate
}

func GetTablename(providerCode string, m time.Month) string {
	providers := []string{"jok", "hbn", "afb"}

	var prov []string
	toLowerProviderCode := strings.ToLower(providerCode)
	for _, v := range providers {
		if toLowerProviderCode == v {
			prov = append(prov, v)
		}
	}

	strProviderCode := strings.Join(prov, ", ")
	month := GetMonth(m)
	allTableName := fmt.Sprintf("%s.%s_%s", strProviderCode, "summary_members", month)

	return allTableName
}

func GetProviderCode() []string {
	providers := []string{"jok", "hbn", "afb"}

	return providers
}
