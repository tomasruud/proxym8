package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

const layout = `<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="go-import" content="{{.Alias}} git {{.URL}}">
    <title>proxym8 module</title>
</head>
<body>proxym8 module</body>
</html>`

type config struct {
	Root  string `yaml:"root"`
	Repos []struct {
		Path string `yaml:"path"`
		Repo string `yaml:"repo"`
	} `yaml:"repos"`
}

func main() {
	in := flag.String("in", "./m8.yaml", "input file name")
	out := flag.String("out", "./out", "output folder")
	flag.Parse()

	file, err := ioutil.ReadFile(*in)
	if err != nil {
		log.Fatalf("unable to read config file: %v", err)
	}

	var cfg config
	if err = yaml.Unmarshal(file, &cfg); err != nil {
		log.Fatalf("unable to parse yaml config: %v", err)
	}

	tmpl, err := template.New("template").Parse(layout)
	if err != nil {
		log.Fatalf("unable to create template: %v", err)
	}

	for _, repo := range cfg.Repos {
		repo.Path = strings.TrimPrefix(repo.Path, "/")
		repo.Path = strings.TrimSuffix(repo.Path, "/")

		alias := fmt.Sprintf("%v/%v", cfg.Root, repo.Path)
		alias = strings.TrimSuffix(alias, "/")

		loc := fmt.Sprintf("%v/%v/%v", *out, repo.Path, "index.html")
		if err := os.MkdirAll(filepath.Dir(loc), 0770); err != nil {
			log.Fatalf("unable to create output directories: %v", err)
		}

		outFile, err := os.Create(loc)
		if err != nil {
			log.Fatalf("unable to create output file: %v", err)
		}

		err = tmpl.Execute(outFile, map[string]string{
			"Alias": alias,
			"URL":   repo.Repo,
		})
		if err != nil {
			log.Fatalf("unable to execute template: %v", err)
		}

		if err := outFile.Close(); err != nil {
			log.Fatalf("unable to close file: %v", err)
		}
	}
}
