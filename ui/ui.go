package ui

import (
	"embed"
	"io/fs"
)

//go:embed dist
var static embed.FS

func FS() (fs.FS, error) {
	return fs.Sub(static, "dist")
}
