package prog

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
	*t.name = normalizePath(*t.name)
	return fileExists(t.Name())
}
