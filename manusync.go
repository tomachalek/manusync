package main

import (
	"flag"
	"fmt"
	"sort"

	"manusync/fsops"
)

func main() {
	flag.Parse()
	files, err := fsops.ListFilesInDirRecur(flag.Arg(0))
	fmt.Println("ERR: ", err)
	sort.Sort(files)
	files.ForEach(func(finfo *fsops.FileItem, i int) bool {
		fmt.Println(">> ", finfo.Path)
		return true
	})
}
