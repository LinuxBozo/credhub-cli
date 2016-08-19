package models

import (
	"fmt"
	"strings"
)

type Ca struct {
	Name   string
	CaBody CaBody
}

func NewCa(name string, caBody CaBody) Item {
	return Ca{
		Name:   name,
		CaBody: caBody,
	}
}

func (ca Ca) String() string {
	lines := []string{}

	caBody := ca.CaBody
	lines = append(lines,
		fmt.Sprintf("Type:		%s", caBody.ContentType),
		fmt.Sprintf("Name:		%s", ca.Name),
	)

	if caBody.Value.Certificate != "" {
		lines = append(lines, fmt.Sprintf("Certificate:		%s", caBody.Value.Certificate))
	}

	if caBody.Value.Private != "" {
		lines = append(lines, fmt.Sprintf("Private:	%s", caBody.Value.Private))
	}

	if caBody.UpdatedAt != "" {
		lines = append(lines, fmt.Sprintf("Updated:	%s", caBody.UpdatedAt))
	}

	return strings.Join(lines, "\n")
}
