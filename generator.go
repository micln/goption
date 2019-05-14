package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var tplAlias map[string]string

func init() {
	tplAlias = map[string]string{
		"": "tpl/class_option.gohtml",
	}
}

type Generator struct {
	TplPath string

	toolOption
}

type toolOption struct {
	verbose bool
}

func (this *Generator) generateCode(clz *Class) (src string, err error) {
	return this.render(clz, this.TplPath)
}

func (this *Generator) render(clz *Class, tplPath string) (src string, err error) {
	t := this.getTemplate(tplPath)

	buf := bytes.NewBuffer(make([]byte, 0, 1024))
	err = t.Execute(buf, clz)

	src = buf.String()

	return
}

func (this *Generator) getTemplate(tplPath string) *template.Template {
	vlog(this.verbose, `template.path:`, tplPath)

	paths := strings.Split(tplPath, "/")
	tplName := paths[len(paths)-1]

	bs, err := ioutil.ReadFile(tplPath)
	if os.IsNotExist(err) {
		bs, err = Asset(tplPath)
	}
	if err != nil {
		panic(err)
	}

	return template.Must(template.New(tplName).
		Funcs(funcs).Parse(string(bs)))
}

func vlog(show bool, args ...interface{}) {
	if show {
		log.Output(1, fmt.Sprintln(args...))
	}
}
