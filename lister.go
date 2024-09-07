package main

import "time"

// file types
const (
	fileRegular int = iota
	fileDirectory
	fileExecutable
	fileCompressed
	fileImage
	fileLink
)

// file extensions
const (
	exe = ".exe"
	deb = ".deb"
	zip = ".zip"
	tar = ".tar"
	rar = ".rar"
	png = ".png"
	jpg = ".jpg"
	gif = ".gif"
)

type file struct {
	name             string
	fileType         int
	isDir            bool
	isHidden         bool
	userName         string
	groupName        string
	size             int64
	modificationTime time.Time
	mode             string
}

type styleFileType struct {
	icon   string
	color  string
	symbol string
}

var mapStyleFileType = map[int]styleFileType{
	fileRegular: styleFileType{
		icon: "📄",
	},
	fileDirectory: styleFileType{
		icon:  "📁",
		color: "blue",
	},
	fileExecutable: styleFileType{
		icon:   "🚀",
		color:  "green",
		symbol: "*",
	},
	fileCompressed: styleFileType{
		icon:  "🗜️",
		color: "red",
	},
	fileImage: styleFileType{
		icon:  "🖼️",
		color: "yellow",
	},
	fileLink: styleFileType{
		icon:  "🔗",
		color: "magenta",
	},
}
