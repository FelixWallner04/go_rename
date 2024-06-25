package internal

import (
	"rename/pkg"
)

func TrimExecute(regex, folder string, tidy bool) {
	pkg.RenameFiles(folder, regex, "")
	if tidy {
		pkg.TidySpecialCharacters(folder)
	}
}
