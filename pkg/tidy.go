package pkg

func TidySpecialCharacters(folder string) []string {
	var files []string

	files = RenameFiles(folder, `^[^a-zA-Z0-9.]+`, "")
	files = RenameFiles(folder, `^[.]`, "Empty.")
	files = RenameFiles(folder, `_\.`, ".")

	return files
}
