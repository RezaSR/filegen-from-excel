package prog

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
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
	ShowVersion  bool
	Mode         modeType
	TemplateFile templateFileType
	OutDir       outDirType
	OutFileName  outFileNameType
	DataFile     dataFileType
)

// Methods of modeType
func (t *modeType) Set(isCli bool) {
	*t.isCli = isCli
}
func (t *modeType) SetP(isCli *bool) {
	t.isCli = isCli
}
func (t *modeType) IsCli() bool {
	return *t.isCli
}
func (t *modeType) Init() error {
	return nil
}

// Methods of dataFileType
func (t *dataFileType) Set(name string) {
	*t.name = name
}
func (t *dataFileType) SetP(name *string) {
	t.name = name
}
func (t *dataFileType) Name() string {
	return *t.name
}
func (t *dataFileType) Validate() error {
	if len(*t.name) == 0 {
		return errors.New("Excel data file is not specified.")
	}
	*t.name = NormalizePath(*t.name)
	return FileExists(t.Name())
}

// Methods of templateFileType
func (t *templateFileType) Set(name string) {
	*t.name = name
}
func (t *templateFileType) SetP(name *string) {
	t.name = name
}
func (t *templateFileType) Name() string {
	return *t.name
}
func (t *templateFileType) Validate() error {
	if len(*t.name) == 0 {
		return errors.New("Template file is not specified.")
	}
	*t.name = NormalizePath(*t.name)
	return FileExists(t.Name())
}

// Methods of outDirType
func (t *outDirType) Set(name string) {
	*t.name = name
}
func (t *outDirType) SetP(name *string) {
	t.name = name
}
func (t *outDirType) Name() string {
	return *t.name
}
func (t *outDirType) Validate() error {
	*t.name = NormalizePath(*t.name)
	return DirExists(t.Name(), true)
}

// Methods of outFileNameType
func (t *outFileNameType) Set(name string) {
	*t.name = name
}
func (t *outFileNameType) SetP(name *string) {
	t.name = name
}
func (t *outFileNameType) Name() string {
	return *t.name
}
func (t *outFileNameType) Validate() error {
	if len(*t.name) == 0 {
		return errors.New("Output file name is not specified.")
	}
	*t.name = NormalizePath(*t.name)
	return nil
}

var (
	UsageTemplateFile = `Template file that contains patterns to be replaced by excel data:
[COLUMN]:
    Replaces with the content of corresponding column from excel data
    For example:
    [A] replaces with the data of cell "A" of current row
Patterns can be escaped by adding ":" after "["
    For example:
    [:A] generates [A]`
	UsageOutFileName = `Output file name that contains special patterns:
[0000]:
    Generates auto increment number padded to the specified zeros
    For example:
    [00].txt generates: 00.txt, 01.txt, 02.txt, 03.txt, ...
[COLUMN]:
    Replaces with the content of corresponding column from excel data
    For example:
    [A].txt replaces [A] with the data of cell "A" of current row
Patterns can be escaped by adding ":" after "["
    For example:
    [:00].txt generates [00].txt`
	DefaultOutFileName = "[0000].txt"
	DefaultOutDir      = "filegen_out"
)

func InitUsage() {
	homeDir, err := os.UserHomeDir()
	if err == nil {
		DefaultOutDir = filepath.Clean(homeDir + "/Documents/" + DefaultOutDir)
	}

	Mode.SetP(flag.Bool("c", false, "Run in CLI mode and do not open GUI"))
	DataFile.SetP(flag.String("d", "", "Excel data file"))
	TemplateFile.SetP(flag.String("t", "", UsageTemplateFile))
	OutDir.SetP(flag.String("o", DefaultOutDir, "Output directory"))
	OutFileName.SetP(flag.String("f", DefaultOutFileName, UsageOutFileName+"\n"))

	v := flag.Bool("v", false, "Version number")

	flag.Parse()

	if *v {
		fmt.Printf("Version: %v\n", VERSION)
		os.Exit(0)
	}
}
