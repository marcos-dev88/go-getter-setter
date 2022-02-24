package file_gs

import "github.com/marcos-dev88/go-getter-setter/getter_setter/logger"

type fileGsMockTest struct {
	FileReader
}

var fileGsEntityMock = NewFileGs(
	"",
	"",
	"",
	"",
	[]Attribute{NewAttribute("", "")},
	logger.NewLogging(),
)

var fileGsEntityMockLocal = NewFileGs(
	"../../testFiles/php/testPhpFile.php",
	"php",
	"private",
	"all",
	[]Attribute{NewAttribute("", "")},
	logger.NewLogging(),
)
