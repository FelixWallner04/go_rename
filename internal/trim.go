package internal

import (
	"fmt"

	"github.com/FelixWallner04/rename/pkg"
)

func TrimExecute(regex, folder string, tidy bool) {
	oldFiles := pkg.GetFilesNameInFolder(folder)
	pkg.RenameFiles(folder, regex, "")
	if tidy {
		pkg.TidySpecialCharacters(folder)
	}
	newFiles := pkg.GetFilesNameInFolder(folder)

	for i := range oldFiles {
		fmt.Printf("%s -> %s\n", oldFiles[i], newFiles[i])
	}
}
