// Code generated by jzero. DO NOT EDIT.
// type: fake_{{.Scope}}_client

package fake

import (
	"github.com/jzero-io/restc"
	"{{$.Module}}/typed/{{.Scope}}"
)g

type Fake{{.Scope | FirstUpper}} struct {}

func (f *Fake{{.Scope | FirstUpper}}) RESTClient() restc.Interface {
	var ret *restc.RESTClient
	return ret
}

{{range $v := .Resources}}func (f *Fake{{$.Scope | FirstUpper}}) {{$v | FirstUpper}}() {{$.Scope}}.{{$v | FirstUpper}}Interface {
	return &Fake{{$v | FirstUpper}}{Fake: f}
}

{{end}}