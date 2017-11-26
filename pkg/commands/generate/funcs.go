package generate

import (
	"bytes"
	"html/template"
	"path/filepath"
)

func getFuncMap(rootDir string) template.FuncMap {
	return template.FuncMap(map[string]interface{}{
		"partial": func(tplName string, ctx interface{}) (string, error) {
			fullDir := filepath.Join(rootDir, tplName)
			buf := new(bytes.Buffer)
			tpl, err := template.ParseFiles(fullDir)
			if err != nil {
				return "", err
			}
			if err := tpl.Execute(buf, ctx); err != nil {
				return "", err
			}
			return string(buf.Bytes()), nil
		},
	})
}
