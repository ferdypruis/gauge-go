package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/getgauge/common"
	"github.com/manuviswam/gauge-go/constants"
)

var pluginDir = ""
var projectRoot = ""
var start = flag.Bool("start", false, "Start go runner")
var initialize = flag.Bool("init", false, "Initialize Go project structure")

func main() {
	flag.Parse()

	setPluginAndProjectRoots()
	if *start {
		startGo()
	} else if *initialize {
		initGo()
	} else {
		printUsage()
	}
}

func startGo() {
	os.Chdir(filepath.Join(projectRoot, constants.DefaultStepImplDir))
	cmd := exec.Command(constants.CommandGo, constants.ArgTest)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error occured while executing 'go test' :", err)
	}
}

func initGo() {
	stepImplDir := filepath.Join(projectRoot, constants.DefaultStepImplDir)
	showMessage("create", stepImplDir)
	createDirectory(stepImplDir)
	stepImplFile := filepath.Join(stepImplDir, constants.DefaultStepImplFileName)
	showMessage("create", stepImplFile)
	common.CopyFile(filepath.Join(constants.SkelDir, constants.DefaultStepImplFileName), stepImplFile)
}

func printUsage() {
	flag.PrintDefaults()
}

func showMessage(action, filename string) {
	fmt.Printf(" %s  %s\n", action, filename)
}

func setPluginAndProjectRoots() {
	var err error
	pluginDir, err = os.Getwd()
	if err != nil {
		fmt.Printf("Failed to find current working directory: %s \n", err)
		os.Exit(1)
	}
	projectRoot = os.Getenv(common.GaugeProjectRootEnv)
	if projectRoot == "" {
		fmt.Printf("Could not find %s env. Go Runner exiting...", common.GaugeProjectRootEnv)
		os.Exit(1)
	}
}

func createDirectory(dirPath string) {
	showMessage("create", dirPath)
	if !common.DirExists(dirPath) {
		err := os.MkdirAll(dirPath, common.NewDirectoryPermissions)
		if err != nil {
			fmt.Printf("Failed to make directory. %s\n", err.Error())
		}
	} else {
		fmt.Println("skip ", dirPath)
	}
}
