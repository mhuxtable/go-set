// vim: set tw=120

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/mhuxtable/go-set/cmd/genset/generator"
	"github.com/mhuxtable/go-set/cmd/genset/templates"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type options struct {
	fileName, packageName, setName string
	withGoGenerateComment          bool
	withTests, quiet               bool
}

func (opts *options) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&opts.fileName, "filename", opts.fileName, "Output filename. This file will be truncated if it exists.")
	fs.StringVar(&opts.packageName, "package-name", opts.packageName,
		"Go package name for the generated set, defaults to last element of output directory if omitted")
	fs.StringVar(&opts.setName, "set-name", opts.setName, "Name of the generated Set datatype")
	fs.BoolVar(&opts.withTests, "generate-tests", opts.withTests, "Generate tests for the Set. (You must supply an item getter function.)")
	fs.BoolVar(&opts.withGoGenerateComment, "generate-comment", opts.withGoGenerateComment, "Generate a go:generate comment to enable re-generation of the file. Assumes this tool is invoked as genset on your PATH.")
	fs.BoolVarP(&opts.quiet, "quiet", "q", opts.quiet, "Quiet mode. Don't prompt and assume yes to all inputs. May clobber files")
}

func (opts *options) String() string {
	var args []string
	args = append(args, fmt.Sprintf(`--filename "%s"`, opts.fileName))
	args = append(args, fmt.Sprintf(`--package-name "%s"`, opts.packageName))
	args = append(args, fmt.Sprintf(`--set-name "%s"`, opts.setName))
	args = append(args, fmt.Sprintf(`--generate-tests=%t`, opts.withTests))
	args = append(args, fmt.Sprintf(`--generate-comment=%t`, opts.withGoGenerateComment))
	args = append(args, fmt.Sprintf(`--quiet=%t`, opts.quiet))

	return strings.Join(args, " ")
}

func main() {
	opts := options{
		fileName:  "zz_genset.go",
		setName:   "Set",
		withTests: false,
	}

	cmd := cobra.Command{
		// TODO(MH) convert DIRECTORY to be a package path (and make module aware)
		Use:   "genset DATATYPE DIRECTORY",
		Short: "Generates a Go set data structure for the datatype DATATYPE emitted to directory DIRECTORY.",
		Args:  cobra.ExactArgs(2),
		RunE: func(_ *cobra.Command, args []string) error {
			return run(args, opts)
		},
	}
	opts.AddFlags(cmd.Flags())

	if err := cmd.Execute(); err != nil {
		log.Fatalf("Error while running command: %s", err)
	}
}

func run(args []string, opts options) error {
	dataType := args[0]
	var outputDir string

	{
		var err error
		outputDir, err = filepath.Abs(args[1])
		if err != nil {
			return fmt.Errorf("while resolving output directory: %w", err)
		}
	}

	if !strings.HasSuffix(opts.fileName, ".go") {
		return fmt.Errorf("output filename does not have suffix .go: %s", opts.fileName)
	}
	if !IsIdentifier(opts.setName) {
		return fmt.Errorf("cannot use %s as set name, as it is not a valid identifier", opts.setName)
	}
	if err := checkOutputDirectory(outputDir); err != nil {
		return err
	}

	if opts.packageName == "" {
		opts.packageName = filepath.Base(outputDir)
	}

	m := generator.Model{
		PackageName:         opts.packageName,
		DataType:            dataType,
		SetTypeName:         opts.setName,
		InternalSetTypeName: internalTypeName(opts.setName, dataType),
		GoGenerateComment:   goGenerateComment(opts, dataType),
	}

	g := generator.Generator{
		Model:               m,
		IgnoreExistingFiles: opts.quiet,
	}

	return GenerateAll(&g, opts.withTests, filepath.Join(outputDir, opts.fileName))
}

func GenerateAll(g *generator.Generator, withTests bool, path string) error {
	if err := g.Generate(templates.Set, path); err != nil {
		return err
	}

	if withTests {
		testOut := strings.TrimSuffix(path, ".go") + "_test.go"
		if err := g.Generate(templates.SetTest, testOut); err != nil {
			return err
		}
	}

	return nil
}

func checkOutputDirectory(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return os.MkdirAll(path, 0755)
		}

		return fmt.Errorf("while getting information about output directory: %w", err)
	}

	return nil
}

func goGenerateComment(opts options, dataType string) string {
	if !opts.withGoGenerateComment {
		return ""
	}

	return strings.Join([]string{
		"//go:generate genset",
		opts.String(), // flags
		dataType,
		".", // output directory
	}, " ")
}

func internalTypeName(setName, dataType string) string {
	// swap out any characters in the data type name to form a valid identifier for the internal set type name
	internalTypeName := "_set_" + setName + "_"
	if IsIdentifier(dataType) {
		internalTypeName += dataType
	} else {
		internalTypeName += strings.Map(func(r rune) rune {
			if unicode.IsLetter(r) || unicode.IsDigit(r) {
				return r
			}
			return '_'
		}, dataType)
	}

	return internalTypeName
}
