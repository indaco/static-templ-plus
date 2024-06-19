package main

import (
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

const (
	outputScriptDirPath  string = "temp"
	outputScriptFileName string = "templ_static_generate_script.go"
)

func usage() {
	fmt.Fprintf(os.Stderr, `Usage of %s:
%s [options]

Options:
  -i  Specify input directory (default "web/pages").
  -o  Specify output directory (default "dist").
  -t  Run templ fmt & generate commands.
  -d  Keep the generation script after completion for inspection and debugging.

Examples:
  # Run templ fmt & templ generate commands and output to default directory
  %s -i web/pages -t=true

  # Specify input and output directories
  %s -i web/pages -o web/pages
`, os.Args[0], os.Args[0], os.Args[0], os.Args[0])
}

func main() {
	var inputDir, outputDir string
	var runTempl, debug bool

	flag.StringVar(&inputDir, "i", "web/pages", `Specify input directory.`)
	flag.StringVar(&outputDir, "o", "dist", `Specify output directory.`)
	flag.BoolVar(&runTempl, "t", false, "Run templ fmt & generate commands.")
	flag.BoolVar(&debug, "d", false, "Keep the generation script after completion for inspection and debugging.")
	flag.Usage = usage
	flag.Parse()

	inputDir = strings.TrimRight(inputDir, "/")
	outputDir = strings.TrimRight(outputDir, "/")

	if outputDir != inputDir {
		if err := clearAndCreateDir(outputDir); err != nil {
			log.Fatal("Error preparing output directory:", err)
		}
	}

	modulePath, err := finder.FindModulePath()
	if err != nil {
		log.Fatal("Error finding module name:", err)
	}

	groupedFiles, err := finder.FindFilesInDir(inputDir)
	if err != nil {
		log.Fatal("Error finding files:", err)
	}

	if runTempl {
		err := generator.RunTemplFmt(groupedFiles.TemplFiles)
		if err != nil {
			log.Fatalf("failed to run 'templ fmt' command: %v", err)
		}
		err = generator.RunTemplGenerate()
		if err != nil {
			log.Fatalf("failed to run 'templ generate' command: %v", err)
		}
	}

	funcs, err := finder.FindFunctionsInFiles(groupedFiles.TemplGoFiles)
	if err != nil {
		log.Fatal("Error finding funcs:", err)
	} else if len(funcs) < 1 {
		log.Fatalf(`No components found in "%s"`, inputDir)
	}

	if err = os.MkdirAll(outputScriptDirPath, os.ModePerm); err != nil {
		log.Fatal("err creating temp dir:", err)
	}

	if err = copyFilesIntoOutputDir(groupedFiles.OtherFiles, inputDir, outputDir); err != nil {
		log.Fatal("err copying files:", err)
	}

	if err = generator.Generate(
		getOutputScriptPath(),
		finder.FindImports(funcs, modulePath),
		funcs,
		inputDir,
		outputDir,
	); err != nil {
		log.Fatal("err generating script", err)
	}

	cmd := exec.Command("go", "run", getOutputScriptPath())
	if err = cmd.Start(); err != nil {
		log.Fatal("err starting script", err)
	}
	if err = cmd.Wait(); err != nil {
		log.Fatal("err running script", err)
	}

	if debug {
		if err = os.RemoveAll(outputScriptDirPath); err != nil {
			log.Fatal("err removing script folder", err)
		}
	}

}

func getOutputScriptPath() string {
	return fmt.Sprintf("%s/%s", outputScriptDirPath, outputScriptFileName)
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
