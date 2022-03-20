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
	"../../testFiles/php/php7/testPhpFile.php",
	"php7",
	"private",
	"all",
	[]Attribute{NewAttribute("", "")},
	logger.NewLogging(),
)

var fileGsEntityMockLocalEight = NewFileGs(
	"../../testFiles/php/php8/myPhpEight.php",
	"php8",
	"private",
	"all",
	[]Attribute{NewAttribute("", "")},
	logger.NewLogging(),
)
