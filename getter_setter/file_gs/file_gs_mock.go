package file_gs

type fileGsMockTest struct {
	FileReader
}

var fileGsEntityMock = NewFileGs(
	"",
	"",
	"",
	[]Attribute{NewAttribute("", "")},
)

var fileGsEntityMockLocal = NewFileGs(
	"../../testFile.php",
	"php",
	"private",
	[]Attribute{NewAttribute("", "")},
)
