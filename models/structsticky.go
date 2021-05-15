package models

import "fmt"

type StructSticky struct {
	Files         []string
	Directories   []string
	Suffix        string
	SearchDir     string
	CountingFiles int
	CountingTotal int
	CountingDir   int
	StickyMode    int
}

func (st StructSticky) Print() {
	println("search directory: ", st.SearchDir)
	println("suffix of file / append dir to search: ", st.Suffix)
	println("counting of suffix file: ", st.CountingFiles)
	println("counting of directory: ", st.CountingDir)
	println("counting of total files: ", st.CountingTotal)
}

func (st StructSticky) PrintDetail() {
	if len(st.Files) > 0 {
		for _, values := range st.Files {
			fmt.Printf("file %s\n", values)
		}
	} else {
		fmt.Println("no file")
	}

	if len(st.Directories) > 0 {
		for _, values := range st.Directories {
			fmt.Printf("dir %s\n", values)
		}
	} else {
		fmt.Println("no directory")
	}
}
