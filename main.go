package main

import (
	"embed"
	"os"
	"runtime/debug"
	"text/template"
	"time"

	flag "github.com/spf13/pflag"
	"gopkg.in/yaml.v3"
)

type id struct {
	Name   string `yaml:"name"`
	Prefix string `yaml:"prefix"`
}

type conf struct {
	Package string `yaml:"package"`
	IDs     []id   `yaml:"ids"`
	Epoch   *Epoch `yaml:"epoch,omitempty"`
}

type Epoch struct {
	Year  int        `yaml:"year"`
	Month time.Month `yaml:"month"`
	Day   int        `yaml:"day"`
}

//go:embed ids.gotmpl
var fs embed.FS

func main() {
	flag.Parse()

	input, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	c := conf{}
	if err := yaml.Unmarshal(input, &c); err != nil {
		panic(err)
	}
	if c.Epoch == nil {
		c.Epoch = &Epoch{
			Year:  1970,
			Month: time.January,
			Day:   0,
		}
	}

	tmpl, err := template.New("ids.gotmpl").ParseFS(fs, "ids.gotmpl")
	if err != nil {
		panic(err)
	}
	err = tmpl.ExecuteTemplate(os.Stdout, "ids.gotmpl", c)
	if err != nil {
		panic(err)
	}
}

var filename string
var version string

func init() {
	flag.StringVarP(&filename, "config", "c", "", "input file name")
}

func goInstallVersion() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "unknown"
	}
	return info.Main.Version
}

func getVersion() string {
	if version != "" {
		return version
	}
	return goInstallVersion()
}
