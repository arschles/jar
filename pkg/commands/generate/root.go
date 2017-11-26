package generate

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	glob "github.com/mattn/go-zglob"
	"github.com/spf13/cobra"
)

// Root returns the command for the root of the tree of the generate commands
func Root() *cobra.Command {
	return &cobra.Command{
		Use:   "generate",
		Short: "Generate the site from templates",
		RunE: func(cmd *cobra.Command, args []string) error {
			var rootDir string
			if len(args) > 0 {
				abs, err := filepath.Abs(args[0])
				if err != nil {
					return fmt.Errorf("Getting absolute path for %s (%s)", args[0], err)
				}
				rootDir = abs
			} else {
				rd, err := os.Getwd()
				if err != nil {
					return fmt.Errorf("Getting the current directory (%s)", err)
				}
				rootDir = rd
			}
			globPattern := fmt.Sprintf("%s/**/*.tpl", rootDir)
			files, err := glob.Glob(globPattern)
			if err != nil {
				return fmt.Errorf("Globbing templates from %s (%s)", rootDir, err)
			}
			logger.Printf("Generating files %v", files)

			tpl := template.New("root").Funcs(getFuncMap(rootDir))
			for _, file := range files {
				t, err := tpl.ParseFiles(file)
				if err != nil {
					return fmt.Errorf("Parsing %s (%s)", file, err)
				}
				tpl = t
			}
			logger.Printf("executing %d templates", len(tpl.Templates()))
			if err := tpl.Execute(os.Stdout, nil); err != nil {
				return fmt.Errorf("Executing template (%s)", err)
			}
			return nil
		},
	}
}
