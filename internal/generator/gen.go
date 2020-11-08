package generator

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	text "text/template"

	"github.com/magejiCoder/gen/internal/template"
)

type GenSet struct {
	SetPackage string
	SetName    string
	SetType    string
}

func NewGenSet(pkgname, setname, settype string) GenSet {
	return GenSet{
		SetPackage: pkgname,
		SetName:    setname,
		SetType:    settype,
	}
}

var ErrOnlySupportGoGeneration error = errors.New("gen: only support generate go code")

func (g GenSet) GenerateTo(dest string) error {
	ext := filepath.Ext(dest)
	if ext != ".go" {
		return ErrOnlySupportGoGeneration
	}
	testFileDest := strings.TrimSuffix(dest, ext) + "_test" + ext
	raw, err := template.LoadSetTemplate()
	if err != nil {
		return fmt.Errorf("load template: %w", err)
	}
	rawTest, err := template.LoadSetTestTemplate()
	if err != nil {
		return fmt.Errorf("load test template: %w", err)
	}

	genmap := map[string][]byte{
		dest:         raw,
		testFileDest: rawTest,
	}
	for dest, raw := range genmap {
		if err := g.gen(raw, dest); err != nil {
			return fmt.Errorf("gen: %w", err)
		}
	}
	return nil
}

func (g GenSet) gen(raw []byte, dest string) error {

	w := bytes.NewBuffer(raw)
	t, err := text.New("SetTemplate").Parse(w.String())
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}
	newBuffer := bytes.NewBuffer(nil)
	if err := t.Execute(newBuffer, g); err != nil {
		return fmt.Errorf("execute template: %w", err)
	}
	return ioutil.WriteFile(dest, newBuffer.Bytes(), os.ModePerm)
}
