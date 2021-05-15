package util

import (
	"github.com/pkg/errors"
	"gostiky/models"
	"os"
	"path/filepath"
	"strings"
)

func CountFiles(args []string, modesticky int) (sticky *models.StructSticky, err error) {
	if len(args) < 2 {
		return nil, errors.New("count does not have two arguments" +
			": searchDirectory" +
			" and suffix of file to count")
	}

	val, err := countFilesParam(args[0], args[1], modesticky)
	if err != nil {
		return nil, err
	}
	return val, nil
}

func countFilesParam(searchDirectory string, suffixFile string, modesticky int) (sticky *models.StructSticky, err error) {

	err = validateCountFiles(searchDirectory, suffixFile)
	if err != nil {
		return nil, err
	}

	return startCounting(searchDirectory, suffixFile, modesticky)
}

func startCounting(searchDirectory string, suffixFile string, modesticky int) (sticky *models.StructSticky, err error) {
	println("searchDir ", searchDirectory)
	println("suffixFile ", suffixFile)
	println("mode_sticky", modesticky)
	var count int
	var countDir int
	var totalcount int
	var fileMode os.FileMode
	fileList := []string{}
	dirList := []string{}
	err = filepath.Walk(searchDirectory, func(path string, f os.FileInfo,
		err error) error {
		switch fileMode = f.Mode(); {
		case fileMode.IsRegular():
			if modesticky < MODE_STICK_IMMUTABLE_RECURSIVE {
				totalcount = totalcount + 1
				if strings.HasSuffix(path, suffixFile) {
					count = count + 1
					fileList = append(fileList, path)
				}
			}
		case fileMode.IsDir():
			if modesticky > MODE_UNSTICK_APPEND {
				if strings.HasSuffix(path, suffixFile) {
					dirList = append(dirList, path)
					countDir = countDir + 1
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	if modesticky <= MODE_UNSTICK_APPEND {
		if len(fileList) < 1 {
			return nil, errors.New("no suffix file found")
		}
	} else {
		if len(dirList) < 1 {
			return nil, errors.New("no append dir found")
		}
	}

	st := models.StructSticky{}
	st.CountingFiles = count
	st.CountingTotal = totalcount
	st.CountingDir = countDir
	st.Suffix = suffixFile
	st.SearchDir = searchDirectory
	if modesticky <= MODE_UNSTICK_APPEND {
		st.Files = append(fileList[:0:0], fileList...) // deep copy
	} else {
		st.Directories = append(dirList[:0:0], dirList...) // deep copy
	}
	return &st, nil
}

// validateCountFiles is to check input parameter to be more than 1 char
// validate searchDirectory
// validate suffixFile (ending of file or extension file)
func validateCountFiles(searchDirectory string, suffixFile string) error {
	var myError string
	if len(searchDirectory) < 1 {
		myError = "you enter less than desirable searchDirectory\n"
	}
	if len(suffixFile) < 1 {
		myError += "you have define less than desirable suffix file\n"
	}
	if len(myError) > 0 {
		return errors.New(myError)
	}
	return nil
}
