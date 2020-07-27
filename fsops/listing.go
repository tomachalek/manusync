package fsops

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

type FileItem struct {
	Path  string
	Mtime int64
}

func (f *FileItem) Basename() string {
	return filepath.Base(f.Path)
}

// FileList is an abstraction for list of files along with their
// modification time information. It supports sorting.
type FileList struct {
	files []*FileItem
}

func (f *FileList) Len() int {
	return len(f.files)
}

func (f *FileList) Less(i, j int) bool {
	return strings.Compare(f.files[i].Path, f.files[j].Path) < 0
}

func (f *FileList) Swap(i, j int) {
	f.files[i], f.files[j] = f.files[j], f.files[i]
}

func (f *FileList) Concat(other *FileList) {
	f.files = append(f.files, other.files...)
}

func (f *FileList) ForEach(fn func(info *FileItem, idx int) bool) {
	for i, v := range f.files {
		if !fn(v, i) {
			break
		}
	}
}

func ListFilesInDirRecur(path string) (*FileList, error) {
	return listFilesInDir(path)
}

func listFilesInDir(path string) (*FileList, error) {
	var ans FileList
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return &ans, err
	}
	ans.files = make([]*FileItem, 0, len(files))
	for _, v := range files {
		fullPath := filepath.Join(path, v.Name())
		if v.IsDir() {
			nested, err := listFilesInDir(fullPath)
			if err != nil {
				log.Fatal("Failed to read dir ", fullPath, err)

			} else {
				ans.Concat(nested)
			}

		} else {
			ans.files = append(ans.files, &FileItem{Path: fullPath, Mtime: v.ModTime().Unix()})
		}
	}
	return &ans, nil
}
