package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"text/template"
)

func readVariablesFromEnv(prefix string) (variables map[string]string) {
	variables = make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if strings.HasPrefix(pair[0], prefix) {
			variables[pair[0]] = pair[1]
		}
	}
	return
}

func applyTemplate(tmpl string, variables map[string]string) (result string) {
	var funcMap = template.FuncMap{
		"default": func(arg interface{}, values ...interface{}) interface{} {
			if len(values) == 0 {
				return arg
			}
			value := values[0]
			v := reflect.ValueOf(value)
			switch v.Kind() {
			case reflect.String:
				if v.Len() == 0 {
					return arg
				}
			default:
				panic("default supports only strings")
			}

			return value
		},
	}

	t := template.Must(template.New("env").Funcs(funcMap).Parse(tmpl))
	var wr bytes.Buffer
	err := t.Execute(&wr, variables)
	if err != nil {
		panic(err)
	}
	result = wr.String()
	return
}

func main() {
	prefixPtr := flag.String("prefix", "", "prefix of env variables")
	flag.Parse()
	if *prefixPtr == "" {
		fmt.Fprintf(os.Stderr, "prefix option is required\n")
		os.Exit(1)
	}
	prefix := *prefixPtr + "_"
	variables := readVariablesFromEnv(prefix)
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "provide template in stdin: %s", err.Error())
		os.Exit(1)
	}
	fmt.Print(applyTemplate(string(bytes), variables))
}
