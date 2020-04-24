package prog

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/tealeg/xlsx"
)

const SHEET_INDEX = 0

const (
	TOKEN_TYPE_Cell = iota
	TOKEN_TYPE_COUNTER
)

type outFileNameTokenInfo struct {
	tokenType int
	data      int
}

var (
	ExePath    string
	WorkingDir string

	templateContent    []byte
	templateTokens     map[int][]byte
	outFileNameTokens  map[string]outFileNameTokenInfo
	outFileNameCounter int = 0
)

func Main() (int, error) {
	err := validateInputs()
	if err != nil {
		return 0, err
	}

	err = parseTemplateFile()
	if err != nil {
		return 0, err
	}

	err = parseOutFileName()
	if err != nil {
		return 0, err
	}

	n, err := processExcelFile()
	if err != nil {
		return n, err
	}

	return n, nil
}

func validateInputs() error {
	var errs []string

	err := DataFile.Validate()
	if err != nil {
		errs = append(errs, err.Error())
	}

	err = TemplateFile.Validate()
	if err != nil {
		errs = append(errs, err.Error())
	}

	err = OutDir.Validate()
	if err != nil {
		errs = append(errs, err.Error())
	}

	err = OutFileName.Validate()
	if err != nil {
		errs = append(errs, err.Error())
	}

	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, "\n"))
	} else {
		return nil
	}
}

func parseTemplateFile() error {
	var err error
	templateTokens = make(map[int][]byte)

	templateContent, err = ioutil.ReadFile(TemplateFile.Name())
	if err != nil {
		return err
	}

	r, err := regexp.Compile(`\[([A-Za-z]+)\]`)
	if err != nil {
		return err
	}
	matches := r.FindAllSubmatch(templateContent, -1)
	for i := range matches {
		templateTokens[xlsx.ColLettersToIndex(string(matches[i][1]))] = matches[i][0]
	}

	return nil
}

func parseOutFileName() error {
	outFileNameTokens = make(map[string]outFileNameTokenInfo)

	r, err := regexp.Compile(`\[([A-Za-z]+)\]`)
	if err != nil {
		return err
	}
	matches := r.FindAllStringSubmatch(OutFileName.Name(), -1)
	for i := range matches {
		outFileNameTokens[matches[i][0]] = outFileNameTokenInfo{tokenType: TOKEN_TYPE_Cell, data: xlsx.ColLettersToIndex(string(matches[i][1]))}
	}

	r, err = regexp.Compile(`\[(0+)\]`)
	if err != nil {
		return err
	}
	matches = r.FindAllStringSubmatch(OutFileName.Name(), -1)
	for i := range matches {
		outFileNameTokens[matches[i][0]] = outFileNameTokenInfo{tokenType: TOKEN_TYPE_COUNTER, data: len(matches[i][1])}
	}

	return nil
}

func processExcelFile() (int, error) {
	dataFile, err := xlsx.OpenFile(DataFile.Name())
	if err != nil {
		return 0, err
	}

	sheetLen := len(dataFile.Sheets)
	if sheetLen == 0 {
		return 0, errors.New("The excel file is empty.")
	}

	sheet := dataFile.Sheets[SHEET_INDEX]

	numOfGenerated := 0
	for _, row := range sheet.Rows {
		if row != nil {
			fileContent := templateContent
			for colIndex, token := range templateTokens {
				cellVal := "" // cell := row.GetCell(colIndex)
				if colIndex < len(row.Cells) {
					cell := row.Cells[colIndex]
					if cell != nil {
						cellVal, err = cell.FormattedValue()
						if err != nil {
							cellVal = err.Error()
						}
					}
				}

				fileContent = bytes.ReplaceAll(fileContent, token, []byte(cellVal))
			}

			fileName := OutFileName.Name()
			for token, tokenInfo := range outFileNameTokens {
				replacement := ""
				switch tokenInfo.tokenType {
				case TOKEN_TYPE_Cell:
					if tokenInfo.data < len(row.Cells) {
						cell := row.Cells[tokenInfo.data]
						if cell != nil {
							replacement, err = cell.FormattedValue()
							if err != nil {
								replacement = err.Error()
							}
						}
					}
				case TOKEN_TYPE_COUNTER:
					replacement = fmt.Sprintf("%0*d", tokenInfo.data, outFileNameCounter)
					outFileNameCounter++
				}
				fileName = strings.ReplaceAll(fileName, token, replacement)
			}

			err = ioutil.WriteFile(filepath.Join(OutDir.Name(), fileName), fileContent, 0644)
			if err == nil {
				numOfGenerated++
			} else {
				return numOfGenerated, err
			}
		}
	}

	return numOfGenerated, nil
}
