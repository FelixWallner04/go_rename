package internal

import (
	"github.com/FelixWallner04/rename/pkg"
)

func TrimExecute(regex, folder string, tidy bool) {
	pkg.RenameFiles(folder, regex, "")
	if tidy {
		pkg.TidySpecialCharacters(folder)
	}
}
