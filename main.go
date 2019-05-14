package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/Kretech/xgo/word"
)

const (
	Version = ``
)

var (
	app             = kingpin.New("goption", "Generator class fields as constructor options")
	verbose         = app.Flag("verbose", "show log").Short('v').Bool()
	pkgName         = app.Flag("package", "package full path").Short('p').Required().String()
	classesName     = app.Flag(`classes`, `match classes name like "person,car".`).Short('c').String()
	enableWriteFile = app.Flag(`writefile`, `write generated code`).Short('w').Default("false").Bool()
	strictCase      = app.Flag(`strictcase`, `strict case sensitive`).Short('s').Default("false").Bool()
	tplPath         = app.Flag(`tpl`, `template file path`).Short('t').Default("").String()
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func init() {
	app.HelpFlag.Short('h')
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		app.Usage(args)
		return
	}

	_, err := app.Parse(args)
	if err != nil {
		log.Printf("parse os.Args: %v\n\n", err)
		app.Usage(args)
		return
	}

	vlog(*verbose, `package:`, *pkgName)
	classes, err := getPackageClasses(*pkgName)
	if err != nil {
		log.Println(err)
		return
	}

	if tplFileName, ok := tplAlias[*tplPath]; ok {
		*tplPath = tplFileName
	}

	if !*strictCase {
		*classesName = strings.ToUpper(*classesName)
	}
	onlyClasses := strings.Split(*classesName, `,`)
	classesMatch := ``
	for _, value := range onlyClasses {
		if value != `` {
			classesMatch += `|` + value + `|`
		}
	}

	g := &Generator{
		TplPath: *tplPath,
		toolOption: toolOption{
			verbose: *verbose,
		},
	}
	for _, class := range classes {
		if classesMatch != `` {
			name := class.Name
			if !*strictCase {
				name = strings.ToUpper(name)
			}

			if !strings.Contains(classesMatch, `|`+name+`|`) {
				vlog(*verbose, `SKIP struct:`, class.Name)
				continue
			}
		}

		src, err := g.generateCode(class)
		if err != nil {
			log.Println(err)
			return
		}

		if !*enableWriteFile {
			fmt.Println(src)
		} else {
			filename := fmt.Sprintf("%s_option_gen.go", word.LowerFirst(word.CamelCase(class.Name)))
			fullpath := path.Join(*pkgName, filename)

			err := ioutil.WriteFile(fullpath, []byte(src), 0644)
			if err != nil {
				log.Println(err)
				continue
			}
			log.Println(`write to`, fullpath)
		}
	}
}
