package main

import (
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/Kretech/xgo/word"
)

func TestDoc(t *testing.T) {

	pkgs, err := parser.ParseDir(token.NewFileSet(), "./testdata", func(fileInfo os.FileInfo) bool {
		return strings.HasSuffix(fileInfo.Name(), ".go")
	}, parser.ParseComments)
	log.Println(err)

	p := doc.New(pkgs["testdata"], "./testdata", doc.AllDecls)
	log.Println(p.Name)
	log.Println(p.Types)

}

func Test_Ast(t *testing.T) {
	//dir, err := ioutil.TempDir("", "goption_test")
	//if err != nil {
	//	t.Error(err)
	//}
	//defer os.RemoveAll(dir)

	g := &Generator{}

	classes, err := getPackageClasses("./testdata")
	log.Println(err)
	for _, class := range classes {
		src, err := g.generateCode(class)

		writeFileName := "./testdata/" + word.LowerFirst(word.UnderlineCase(class.Name)) + "_option_gen.go"

		ioutil.WriteFile(writeFileName, []byte(src), 0644)

		//log.Println(src)
		log.Println(err)
	}

	return

	type args struct {
		file     *ast.File
		filename string
	}
	tests := []struct {
		name        string
		args        args
		wantClasses []*Class
		wantErr     bool
	}{
		// TODO: Add test cases.
		{`person`, args{filename: "Person.go"}, []*Class{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotClasses, err := parseClass(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseClass() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotClasses, tt.wantClasses) {
				t.Errorf("parseClass() = %v, want %v", gotClasses, tt.wantClasses)
			}
		})
	}
}
