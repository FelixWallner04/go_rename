package pkg

func TidySpecialCharacters(folder string) {
	RenameFiles(folder, `^[^a-zA-Z0-9.]+`, "")
	RenameFiles(folder, `^[.]`, "Empty.")
	RenameFiles(folder, `_\.`, ".")
}
