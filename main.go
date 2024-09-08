package main

import (
	"flag"
	"fmt"
	"io/fs"
	"lister/utils"
	"os"
	"os/user"
	"regexp"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	//Flag declaration
	flagPattern := flag.String("p", "", "Filter by file path")
	//flagAll := flag.Bool("a", false, "List all files")
	flagNumberRecords := flag.Int("n", 10, "Number of records to display")
	//flagOrderByTime := flag.Bool("t", false, "Order by time")
	//flagOrderBySize := flag.Bool("s", false, "Order by size")
	//flagOrderByName := flag.Bool("N", false, "Order by name")

	//Parse the flags
	flag.Parse()

	path := flag.Arg(0)
	if path == "" {
		path = "."
	}

	memDirs, err := os.ReadDir(path)
	utils.ErrHandler(err, true)

	fs := []file{}

	for _, dir := range memDirs {
		file, err := getFile(dir, false)
		utils.ErrHandler(err, true)

		math, err := regexp.MatchString("(?i)"+*flagPattern, file.name)
		utils.ErrHandler(err, true)

		if !math {
			continue
		}

		fs = append(fs, file)
	}

	if *flagNumberRecords > len(fs) {
		*flagNumberRecords = len(fs)
	}
	formatList(fs, *flagNumberRecords)
}

func formatList(fs []file, nRecords int) {
	for _, f := range fs[:nRecords] {
		style := mapStyleFileType[f.fileType]
		fmt.Printf("%s %s %8d %s %s %s\n", f.mode, f.userName, f.size, f.modificationTime.Format("2006-01-02 15:04:05"), style.icon, f.name)
	}
}

func getFile(dir fs.DirEntry, isHidden bool) (file, error) {
	inf, err := dir.Info()
	utils.ErrHandler(err, true)

	stat := inf.Sys().(*syscall.Stat_t)

	// Obtener el nombre de usuario y grupo a partir de Uid y Gid
	usr, err := user.LookupId(strconv.Itoa(int(stat.Uid)))
	utils.ErrHandler(err, true)

	group, err := user.LookupGroupId(strconv.Itoa(int(stat.Gid)))
	utils.ErrHandler(err, true)

	f := file{
		name:             dir.Name(),
		isDir:            dir.IsDir(),
		isHidden:         isHidden,
		userName:         usr.Username,
		groupName:        group.Name,
		size:             inf.Size(),
		modificationTime: inf.ModTime(),
		mode:             inf.Mode().String(),
	}

	designFile(&f)

	return f, nil
}

func designFile(f *file) {
	isLink := strings.HasPrefix(strings.ToUpper(f.mode), "L")
	isExecutable := strings.Contains(f.mode, "x")
	isCompress := strings.HasSuffix(f.mode, zip) || strings.HasSuffix(f.mode, tar) || strings.HasSuffix(f.mode, rar) || strings.HasSuffix(f.mode, deb)
	isImage := strings.HasSuffix(f.mode, png) || strings.HasSuffix(f.mode, jpg) || strings.HasSuffix(f.mode, gif)

	switch {
	case isLink:
		f.fileType = fileLink
	case f.isDir:
		f.fileType = fileDirectory
	case isExecutable:
		f.fileType = fileExecutable
	case isCompress:
		f.fileType = fileCompressed
	case isImage:
		f.fileType = fileImage
	default:
		f.fileType = fileRegular
	}
}
