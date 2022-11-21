package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"
)

// App struct
type App struct {
	ctx context.Context
}

var exPath string

var fileNumber int
var fileNumProcessed int

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// CompFile composits two files ontop of each other and saves them to the given filepathRegionNew
func CompFile(filepathRegion, filepathOriginal, filepathNew string) {

	var err error

	// Solution for not enough memory - try again after one second and stop loop when the command succeeded. Kinda wonky but it works..
	for {
		// run ImageMagick's compos.exe without terminal window appearing
		cmd_path := "C:\\Windows\\system32\\cmd.exe"
		cmd_instance := exec.Command(cmd_path, "/c", exPath+"/lib/composite.exe", "-compose", "atop", "-geometry", "-0-0", filepathRegion, filepathOriginal, filepathNew)
		cmd_instance.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		err = cmd_instance.Run()

		if err == nil {
			break
		}

		fmt.Println("Error:", err)
		fmt.Println("Trying again for ", filepathNew)
		time.Sleep(time.Second)
	}
	fileNumProcessed++
}

func FolderExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func (a *App) CompIt(imageMagickPath, folderRegion, folderOriginal, ignorePrefixes string, overwrite bool) {

	if !FolderExists(folderOriginal) || !FolderExists(folderRegion) {
		return
	}

	fileNumProcessed = 0

	filesRegion, err := ioutil.ReadDir(folderRegion)
	filesOriginal, err := ioutil.ReadDir(folderOriginal)

	fileNumber = len(filesRegion)

	if err != nil {
		fmt.Println(err)
	}

	ex, err := os.Executable()
	exPath = filepath.Dir(ex)

	mRegionFiles := make(map[int]string)

	reFilterNumber := regexp.MustCompile("[0-9]+")

	ignorePrefixesArr := strings.Split(ignorePrefixes, ";")

	for _, fileRegion := range filesRegion {

		hasPrefix := false

		for _, prefix := range ignorePrefixesArr {
			if strings.Contains(fileRegion.Name(), prefix) {
				hasPrefix = true
			}
		}

		if hasPrefix {
			continue
		}

		numbersInName := reFilterNumber.FindAllString(fileRegion.Name(), -1)

		frameNumberRegion, err := strconv.Atoi(numbersInName[len(numbersInName)-1])

		if err != nil {
			println("cannot get region frame number for ", fileRegion.Name(), err)
			continue
		}

		mRegionFiles[frameNumberRegion] = fileRegion.Name()
	}

	for _, fileOriginal := range filesOriginal {

		if fileOriginal.IsDir() {
			continue
		}

		numbersInName := reFilterNumber.FindAllString(fileOriginal.Name(), -1)

		frameNumberOriginal, err := strconv.Atoi(numbersInName[len(numbersInName)-1])

		filepathOriginal := filepath.Join(folderOriginal, fileOriginal.Name())
		filepathRegion := filepath.Join(folderRegion, mRegionFiles[frameNumberOriginal])

		fmt.Println("check for existence: ", filepathRegion)
		if !FileExists(filepathRegion) {
			fmt.Println(filepathRegion, " does not exist - continue")
			continue
		}

		fileNew := ""

		if overwrite {
			fileNew = filepath.Join(folderOriginal, fileOriginal.Name())
		} else {
			err = os.MkdirAll(filepath.Join(path.Join(folderOriginal, "anrender")), os.ModePerm)
			if err != nil {
				println("Could not create output folder ", err)
				return
			}
			fileNew = filepath.Join(folderOriginal, "anrender", fileOriginal.Name())
		}

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(fileOriginal.Name(), frameNumberOriginal)

		// go CompFile(filepathRegion, filepathRegion, filepathRegionNew) //TODO solve memory issue...
		go CompFile(filepathRegion, filepathOriginal, fileNew)

		if err != nil {
			fmt.Println(err)
		}
	}

	if err != nil {
		panic(err)
	}

}

func (a *App) GetProgress() int {
	if fileNumProcessed >= fileNumber {
		return 100
	}
	return int(float32(fileNumProcessed) / float32(fileNumber) * 100)
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}
