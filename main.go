package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	c "github.com/hikachan/ex-redmine/config"
	"github.com/k0kubun/pp"
	//"github.com/hikachan/ex-redmine/config"
)

func fatal(format string, err error) {
	fmt.Errorf(format, err)
	os.Exit(1)
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fatal("Failed to get dir: %s\n", err)
	}
	path := filepath.Join(dir, ".config/settings.json")
	_, err = toml.DecodeFile(path, &c.Conf)
	pp.Println(c.Conf)
	if err != nil {
		fatal("toml could not load: %s\n", err)
	}
}
