package goword

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

func ParseText(filename string) (string, error) {
	doc, err := openWordFile(filename)
	if err != nil {
		return "", fmt.Errorf("error opening file %s - %s", filename, err)
	}

	docx, err := Parse(doc)
	if err != nil {
		return "", fmt.Errorf("error parsing %s - %s", filename, err)
	}

	return docx.AsText(), nil
}

func Parse(doc string) (WordDocument, error) {
	docx := WordDocument{}
	r := strings.NewReader(doc)
	decoder := xml.NewDecoder(r)

	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "p" {
				var p Paragraph
				err := decoder.DecodeElement(&p, &se)
				if err != nil {
					return docx, err
				}
				docx.Paragraphs = append(docx.Paragraphs, p)
			}
		}
	}

	return docx, nil
}

func openWordFile(filename string) (string, error) {
	// Open a zip archive for reading. word files are zip archives
	r, err := zip.OpenReader(filename)
	if err != nil {
		return "", err
	}
	defer r.Close()

	// Iterate through the files in the archive,
	// find document.xml
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return "", err
		}
		if f.Name == "word/document.xml" {
			doc, err := io.ReadAll(rc)
			if err != nil {
				return "", err
			}

			return string(doc), nil
		}
		rc.Close()
	}

	return "", nil
}
