// This program generates contributors.go. It can be invoked by running
// go generate
package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"time"
)

func main() {
	os.Remove("flag_gen.go")
	f, err := os.Create("flag_gen.go")
	die(err)
	defer f.Close()

	packageTemplate.Execute(f, struct {
		Timestamp  time.Time
		FloatFlags []string
		IntFlags   []string
	}{
		Timestamp:  time.Now(),
		FloatFlags: []string{"float32", "float64"},
		IntFlags:   []string{"int", "int8", "int16", "int32", "int64"},
	})
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var removeNonDigits = regexp.MustCompile("[^0-9]+")
var funcMap = template.FuncMap{
	"Ucfirst": func(s string) string {
		return strings.ToUpper(string(s[0])) + s[1:]
	},
	"GetBitLen": func(s string) int {
		i, err := strconv.ParseInt(removeNonDigits.ReplaceAllString(s, ""), 10, 0)
		if err != nil {
			return 0
		}
		return int(i)
	},
}
var packageTemplate = template.Must(template.New("").Funcs(funcMap).Parse(`// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// {{ .Timestamp }}
package flag

import (
	"fmt"
	"strconv"
)

{{- range .FloatFlags }}
type {{ printf "%s" . | Ucfirst }}Flag struct {
	ValueFlag[{{ printf "%s" . }}]

	Names       []string
	Required    bool
	Description string
	value       {{ printf "%s" . }}
	wasSet      bool
}

func (f *{{ printf "%s" . | Ucfirst }}Flag) NameList() []string {
	return f.Names
}

func (f *{{ printf "%s" . | Ucfirst }}Flag) Value() {{ printf "%s" . }} {
	return f.value
}

func (f *{{ printf "%s" . | Ucfirst }}Flag) IsSet() bool {
	return f.wasSet
}

func (f *{{ printf "%s" . | Ucfirst }}Flag) IsRequired() bool {
	return f.Required
}

func (f *{{ printf "%s" . | Ucfirst }}Flag) Apply(ctx *ParseCtx) error {
	a, ok := ctx.Args().Peek()
	if !ok || a.IsFlag() {
		return fmt.Errorf("flag '%s' requires a parameter", f.Names[0])
	}

	ctx.Args().Next() // take next argument
	v, e := strconv.ParseFloat(a.String(), {{ printf "%s" . | GetBitLen }})
	if e != nil {
		return fmt.Errorf("flag '%s' contains invalid float '%v'", f.Names[0], a.String())
	}
	f.value = {{ printf "%s" . }}(v)
	f.wasSet = true
	return nil
}

func (f *{{  printf "%s" . | Ucfirst }}Flag) HelpDescription() string {
	return f.Description
}
{{- end }}

{{- range .IntFlags }}
type {{ printf "%s" . | Ucfirst }}Flag struct {
	ValueFlag[{{ printf "%s" . }}]

	Names       []string
	Required    bool
	Description string
	value       {{ printf "%s" . }}
	wasSet      bool
}

func (f *{{ printf "%s" . | Ucfirst }}Flag) NameList() []string {
	return f.Names
}

func (f *{{ printf "%s" . | Ucfirst }}Flag) Value() {{ printf "%s" . }} {
	return f.value
}

func (f *{{ printf "%s" . | Ucfirst }}Flag) IsSet() bool {
	return f.wasSet
}

func (f *{{ printf "%s" . | Ucfirst }}Flag) IsRequired() bool {
	return f.Required
}

func (f *{{ printf "%s" . | Ucfirst }}Flag) Apply(ctx *ParseCtx) error {
	a, ok := ctx.Args().Peek()
	if !ok || a.IsFlag() {
		return fmt.Errorf("flag '%s' requires a parameter", f.Names[0])
	}

	ctx.Args().Next() // take next argument
	v, e := strconv.ParseInt(a.String(), 10, {{ printf "%s" . | GetBitLen }})
	if e != nil {
		return fmt.Errorf("flag '%s' contains invalid int '%v'", f.Names[0], a.String())
	}
	f.value = {{ printf "%s" . }}(v)
	f.wasSet = true
	return nil
}

func (f *{{  printf "%s" . | Ucfirst }}Flag) HelpDescription() string {
	return f.Description
}

type {{ printf "u%s" . | Ucfirst }}Flag struct {
	ValueFlag[{{ printf "u%s" . }}]

	Names       []string
	Required    bool
	Description string
	value       {{ printf "u%s" . }}
	wasSet      bool
}

func (f *{{ printf "u%s" . | Ucfirst }}Flag) NameList() []string {
	return f.Names
}

func (f *{{ printf "u%s" . | Ucfirst }}Flag) Value() {{ printf "u%s" . }} {
	return f.value
}

func (f *{{ printf "u%s" . | Ucfirst }}Flag) IsSet() bool {
	return f.wasSet
}

func (f *{{ printf "u%s" . | Ucfirst }}Flag) IsRequired() bool {
	return f.Required
}

func (f *{{ printf "u%s" . | Ucfirst }}Flag) Apply(ctx *ParseCtx) error {
	a, ok := ctx.Args().Peek()
	if !ok || a.IsFlag() {
		return fmt.Errorf("flag '%s' requires a parameter", f.Names[0])
	}

	ctx.Args().Next() // take next argument
	v, e := strconv.ParseUint(a.String(), 10, {{ printf "u%s" . | GetBitLen }})
	if e != nil {
		return fmt.Errorf("flag '%s' contains invalid int '%v'", f.Names[0], a.String())
	}
	f.value = {{ printf "u%s" . }}(v)
	f.wasSet = true
	return nil
}

func (f *{{  printf "u%s" . | Ucfirst }}Flag) HelpDescription() string {
	return f.Description
}
{{- end }}
`))
