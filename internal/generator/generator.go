package generator

import (
	"bytes"
	"go/format"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Generator struct {
	templateDir string
	outputDir   string
}

type DAOTemplate struct {
	Package    string
	EntityName string
	Methods    []MethodTemplate
}

type MethodTemplate struct {
	Name       string
	Parameters []ParameterTemplate
	ReturnType string
	SQLFile    string
}

type ParameterTemplate struct {
	Name string
	Type string
}

func NewGenerator(templateDir, outputDir string) *Generator {
	return &Generator{
		templateDir: templateDir,
		outputDir:   outputDir,
	}
}

func (g *Generator) GenerateDAO(data DAOTemplate) error {
	tmpl, err := template.ParseFiles(filepath.Join(g.templateDir, "dao.go.tmpl"))
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return err
	}

	// Goのコードをフォーマット
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	// 出力ファイルパスの作成
	outputFile := filepath.Join(g.outputDir, strings.ToLower(data.EntityName)+"_dao.go")

	return os.WriteFile(outputFile, formatted, 0644)
}
