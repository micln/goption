package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"os"
	"path"
)

const tplDir = `tpl`

type Generator struct {
}

func (this *Generator) generateCode(clz *Class) (src string, err error) {
	return this.render(clz, "class_option.gohtml")
}

func (this *Generator) render(clz *Class, tplName string) (src string, err error) {
	t := this.getTemplate(tplName)

	buf := bytes.NewBuffer(make([]byte, 0, 1024))
	err = t.Execute(buf, clz)

	src = buf.String()

	return
}

func (this *Generator) getTemplate(tplName string) (*template.Template) {

	tplPath := path.Join(tplDir, tplName)
	bs, err := ioutil.ReadFile(tplPath)
	if os.IsNotExist(err) {
		bs, err = Asset("tpl/class_option.gohtml")
	}
	if err != nil {
		panic(err)
	}

	return template.Must(template.New(tplName).
		Funcs(funcs).Parse(string(bs)))
}
