package pkg

import (
	"log"
	"os"
	"regexp"
)

func GetFilesNameInFolder(folder string) []string {
	var files []string

	filesInfo, err := os.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}

	for _, fileInfo := range filesInfo {
		files = append(files, fileInfo.Name())
	}

	return files
}

func RenameFiles(folder, regex, substitution string) []string {
	var renamedFiles []string
	files := GetFilesNameInFolder(folder)
	reg := regexp.MustCompile(regex)

	for _, file := range files {
		if !reg.MatchString(file) {
			renamedFiles = append(renamedFiles, file)
			continue
		}
		newFileName, oldFileName := file, file
		newFileName = reg.ReplaceAllString(newFileName, substitution)

		if err := os.Rename(folder+"/"+oldFileName, folder+"/"+newFileName); err != nil {
			log.Fatal()
		}
		renamedFiles = append(renamedFiles, newFileName)
	}

	return renamedFiles
}
