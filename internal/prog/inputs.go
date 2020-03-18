package prog

import (
	"errors"
	"flag"
)

type modeType struct {
	isCli *bool
}

type templateFileType struct {
	name *string
}

type outDirType struct {
	name *string
}

type outFileNameType struct {
	name *string
}

type dataFileType struct {
	name *string
}

var (
	Mode         modeType
	TemplateFile templateFileType
	OutDir       outDirType
	OutFileName  outFileNameType
	DataFile     dataFileType
)

// Methods of modeType
func (t *modeType) Set(isCli *bool) {
	t.isCli = isCli
}
func (t *modeType) IsCli() bool {
	return *t.isCli
}
func (t *modeType) Init() error {
	return nil
}

// Methods of templateFileType
func (t *templateFileType) Set(name *string) {
	t.name = name
}
func (t *templateFileType) Name() string {
	return *t.name
}
func (t *templateFileType) Init() error {
	if len(*t.name) == 0 {
		return errors.New("Template file is not specified.")
	}
	*t.name = normalizePath(*t.name)
	return fileExists(t.Name())
}

// Methods of outDirType
func (t *outDirType) Set(name *string) {
	t.name = name
}
func (t *outDirType) Name() string {
	return *t.name
}
func (t *outDirType) Init() error {
	*t.name = normalizePath(*t.name)
	return dirExists(t.Name(), true)
}

// Methods of outFileNameType
func (t *outFileNameType) Set(name *string) {
	t.name = name
}
func (t *outFileNameType) Name() string {
	return *t.name
}
func (t *outFileNameType) Init() error {
	if len(*t.name) == 0 {
		return errors.New("Output file name is not specified.")
	}
	*t.name = normalizePath(*t.name)
	return nil
}

// Methods of dataFileType
func (t *dataFileType) Set(name *string) {
	t.name = name
}
func (t *dataFileType) Name() string {
	return *t.name
}
func (t *dataFileType) Init() error {
	if len(*t.name) == 0 {
		return errors.New("Excel data file is not specified.")
	}
	*t.name = normalizePath(*t.name)
	return fileExists(t.Name())
}

func InitUsage() {
	usageTemplateFile := `Template file that contains patterns to be replaced by excel data:
    [COLUMN]: Replaces with the content of the corresponding column from excel data
        For example: [A] replaces with the data of cell "A" of current row
    Patterns can be escaped by adding ":" after "["
        For example: [:A] generates [A]
`

	usageOutFileName := `Output file name that contains special patterns:
    [0000]: Generates auto increment number padded to the specified zeros
        For example: [00].txt generates: 00.txt, 01.txt, 02.txt, 03.txt, ...
    [COLUMN]: Replaces with the content of the corresponding column from excel data
        For example: [A].txt replaces [A] with the data of cell "A" of current row
    Patterns can be escaped by adding ":" after "["
        For example: [:00].txt generates [00].txt
`

	Mode.Set(flag.Bool("c", false, "Run in CLI mode and do not open GUI"))
	TemplateFile.Set(flag.String("t", "", usageTemplateFile))
	OutDir.Set(flag.String("o", "out", "Output directory"))
	OutFileName.Set(flag.String("f", "[0000].txt", usageOutFileName))
	DataFile.Set(flag.String("d", "", "Excel data file"))
	flag.Parse()
}
