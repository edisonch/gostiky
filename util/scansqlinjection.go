package util

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"gostiky/models"
	"os"
	"regexp"
)

func ProcessFiles(sticky *models.StructSticky) ([]string, error) {

	if len(sticky.Files) < 1 {
		return nil, errors.New("Receive empty input")
	}

	stmtsql := [3]string{"INSERT INTO", "UPDATE", "SELECT"}
	r1, _ := regexp.Compile(stmtsql[0])
	r2, _ := regexp.Compile(stmtsql[1])
	r3, _ := regexp.Compile(stmtsql[2])

	pattern := `->prepare(`
	r4, _ := regexp.Compile(regexp.QuoteMeta(pattern))

	foundfiles := []string{}

	for _, val := range sticky.Files {
		if parseLines(val, r1, r2, r3, r4) {
			foundfiles = append(foundfiles, val)
		}
	}

	return foundfiles, nil
}

// parseLines is to find file that has r1 or r2 or r3 but it does not have
// r4 anywhere in the file thus it would return true
func parseLines(filename string, r ...*regexp.Regexp) bool {

	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Errorf("NOT FOUND %q", filename)
		return false
	}
	defer file.Close()

	var line string
	found := false
	foundPrepare := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		if len(r[0].FindAllString(line, -1)) > 0 {
			//fmt.Printf("found sql in %s : %s \n",filename,line)
			found = true
			continue
		}
		if len(r[1].FindAllString(line, -1)) > 0 {
			//fmt.Printf("found sql in %s : %s \n",filename,line)
			found = true
			continue
		}
		if len(r[2].FindAllString(line, -1)) > 0 {
			//fmt.Printf("found sql in %s : %s \n",filename,line)
			found = true
			continue
		}
		if found {
			if len(r[3].FindAllString(line, -1)) > 0 {
				foundPrepare = true
				break
			}
		}
	}

	if found && !foundPrepare {
		return true
	}
	return false
}
