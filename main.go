package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/indaco/static-templ-plus/internal/finder"
	"github.com/indaco/static-templ-plus/internal/generator"
)

//go:embed .version
var versionFile embed.FS

const (
	templVersion         = "0.2.747"
	outputScriptDirPath  = "temp"
	outputScriptFileName = "templ_static_generate_script.go"
)

type flags struct {
	InputDir    string
	OutputDir   string
	Mode        string
	RunFormat   bool
	RunGenerate bool
	Debug       bool
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	switch os.Args[1] {
	case "version", "--version":
		handleVersionCmd()
		return
	default:
		// Continue with existing flag parsing
	}

	flags := parseFlags()

	if flags.OutputDir != flags.InputDir {
		if err := clearAndCreateDir(flags.OutputDir); err != nil {
			log.Fatal("Error preparing output directory:", err)
		}
	}

	modulePath, groupedFiles := prepareDirectories(flags.InputDir)

	if flags.RunFormat {
		runTemplFmt(groupedFiles)
	}

	if flags.RunGenerate {
		groupedFiles = runTemplGenerate(flags.InputDir)
	}

	funcs := findFunctions(groupedFiles.TemplGoFiles)

	if err := os.MkdirAll(outputScriptDirPath, os.ModePerm); err != nil {
		log.Fatal("Error creating temp dir:", err)
	}

	if err := copyFilesIntoOutputDir(groupedFiles.OtherFiles, flags.InputDir, flags.OutputDir); err != nil {
		log.Fatal("Error copying files:", err)
	}

	if err := generator.Generate(getOutputScriptPath(), finder.FindImports(funcs, modulePath), funcs, flags.InputDir, flags.OutputDir); err != nil {
		log.Fatal("Error generating script:", err)
	}

	runGeneratedScript(flags.Debug)
}

func handleVersionCmd() {
	versionCmd := flag.NewFlagSet("version", flag.ExitOnError)
	err := versionCmd.Parse(os.Args[2:])
	if err != nil {
		return
	}
	printVersion(getVersion(), templVersion)
}

func parseFlags() flags {
	var flags flags

	flag.StringVar(&flags.InputDir, "i", "web/pages", "Specify input directory.")
	flag.StringVar(&flags.OutputDir, "o", "dist", "Specify output directory.")
	flag.StringVar(&flags.Mode, "mode", "standard", "Set the operational mode (standard or inline).")
	flag.BoolVar(&flags.RunFormat, "f", false, "Run templ fmt.")
	flag.BoolVar(&flags.RunGenerate, "g", false, "Run templ generate.")
	flag.BoolVar(&flags.Debug, "d", false, "Keep the generation script after completion for inspection and debugging.")
	flag.Usage = usage
	flag.Parse()

	flags.InputDir = strings.TrimRight(flags.InputDir, "/")
	flags.OutputDir = strings.TrimRight(flags.OutputDir, "/")

	return flags
}

func prepareDirectories(inputDir string) (string, *finder.GroupedFiles) {
	modulePath, err := finder.FindModulePath()
	if err != nil {
		log.Fatal("Error finding module name:", err)
	}

	groupedFiles, err := finder.FindFilesInDir(inputDir)
	if err != nil {
		log.Fatal("Error finding files:", err)
	}

	return modulePath, groupedFiles
}

func runTemplFmt(groupedFiles *finder.GroupedFiles) {
	done := make(chan struct{})
	go func() {
		err := generator.RunTemplFmt(groupedFiles.TemplFiles, done)
		if err != nil {
			log.Fatalf("failed to run 'templ fmt' command: %v", err)
		}
	}()
	<-done
	log.Println("completed running 'templ fmt'")
}

func runTemplGenerate(inputDir string) *finder.GroupedFiles {
	done := make(chan struct{})
	go func() {
		err := generator.RunTemplGenerate(done)
		if err != nil {
			log.Fatalf("failed to run 'templ generate' command: %v", err)
		}
	}()
	<-done
	log.Println("completed running 'templ generate'")

	groupedFiles, err := finder.FindFilesInDir(inputDir)
	if err != nil {
		log.Fatal("Error finding _templ.go files after templ generate completion:", err)
	}
	return groupedFiles
}

func findFunctions(templGoFiles []string) []finder.FunctionToCall {
	funcs, err := finder.FindFunctionsInFiles(templGoFiles)
	if err != nil {
		log.Fatal("Error finding functions:", err)
	} else if len(funcs) < 1 {
		log.Fatalf(`No components found`)
	}
	return funcs
}

func runGeneratedScript(debug bool) {
	cmd := exec.Command("go", "run", getOutputScriptPath())
	if err := cmd.Start(); err != nil {
		log.Fatal("Error starting script:", err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal("Error running script:", err)
	}

	if !debug {
		if err := os.RemoveAll(outputScriptDirPath); err != nil {
			log.Fatal("Error removing script folder:", err)
		}
	}
}

func usage() {
	output := fmt.Sprintf(`Usage of %[1]v:
%[1]v [flags] [subcommands]

Flags:
  -i  Specify input directory (default "web/pages").
  -o  Specify output directory (default "dist").
  -f  Run templ fmt.
  -g  Run templ generate.
  -d  Keep the generation script after completion for inspection and debugging.

Subcommands:
  version  Display the version information.

Examples:
  # Specify input and output directories
  %[1]v -i web/demos -o output

  # Specify input directory, run templ generate and output to default directory
  %[1]v -i web/demos -g=true

  # Display the version information
  %[1]v version
`, os.Args[0])

	fmt.Println(output)
}

func getVersion() string {
	content, err := versionFile.ReadFile(".version")
	if err != nil {
		return "unknown"
	}
	return strings.TrimSpace(string(content))
}

func printVersion(version, templVersion string) {
	templModulePath := "github.com/a-h/templ"
	fmt.Printf("Version: %s (built with %s@v%s)\n", version, templModulePath, templVersion)
}

func clearAndCreateDir(dir string) error {
	if err := os.RemoveAll(dir); err != nil {
		return err
	}
	return os.MkdirAll(dir, os.ModePerm)
}

func copyFile(fromPath string, toPath string) error {
	if err := os.MkdirAll(path.Dir(toPath), os.ModePerm); err != nil {
		return err
	}

	src, err := os.ReadFile(fromPath)
	if err != nil {
		return err
	}

	if err = os.WriteFile(toPath, src, 0644); err != nil {
		return err
	}
	return nil
}

func copyFilesIntoOutputDir(files []string, inputDir string, outputDir string) error {
	for _, f := range files {
		if err := copyFile(f, strings.Replace(f, inputDir, outputDir, 1)); err != nil {
			return err
		}
	}
	return nil
}

func getOutputScriptPath() string {
	return fmt.Sprintf("%s/%s", outputScriptDirPath, outputScriptFileName)
}
