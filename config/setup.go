package config

import (
	"isso0424/oroka-discord/loader"
	"os"
	"path"
	"path/filepath"
)

func Setup(patternFile string) {
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}

	dir := filepath.Dir(exe)
	patterns, err := loader.Load(path.Join(dir, patternFile))
	if err != nil {
		panic(err)
	}

	ReactionPatterns = patterns
}
