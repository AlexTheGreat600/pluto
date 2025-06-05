//  ___ _   _   _ _____ ___
// | _ \ | | | | |_   _/ _ \
// |  _/ |_| |_| | | || (_) |
// |_| |____\___/  |_| \___/
//

package main

import (
	"os"
	"fmt"
	"strings"
	"path/filepath"
	"github.com/fatih/color"
	"github.com/iancoleman/strcase"
)

/* Variable to confirm rename file or not */
var confirmRename string

/* Function to show the help of program */
func plutoHelp() {

	/* Show the pluto logo at the start */
	color.HiBlue	("___  _    _  _ ___ ____")
	color.HiBlue	("|__] |    |  |  |  |  |")
	color.Blue		("|    |___ |__|  |  |__|")
	color.Blue		("                       ")

	/* Show pluto usage at start */
	fmt.Printf("%s [name-type] [your-files]\n\n", color.HiWhiteString("PLUTO :"))

	/* Show the case types in colors */
	color.HiWhite("NAME-TYPES : \n\n")
	fmt.Printf("%s %s | %s\n", color.HiRedString("*"), color.HiWhiteString("-k "), color.HiWhiteString("--kebab"))
	fmt.Printf("%s %s | %s\n", color.HiGreenString("*"), color.HiWhiteString("-c "), color.HiWhiteString("--camel"))
	fmt.Printf("%s %s | %s\n", color.HiBlueString("*"), color.HiWhiteString("-s "), color.HiWhiteString("--snake"))
	fmt.Printf("%s %s | %s\n", color.HiYellowString("*"), color.HiWhiteString("-lc"), color.HiWhiteString("--lowerCamel"))
	fmt.Printf("%s %s | %s\n", color.HiMagentaString("*"), color.HiWhiteString("-ss"), color.HiWhiteString("--screamingSnake"))
	fmt.Printf("%s %s | %s\n", color.HiCyanString("*"), color.HiWhiteString("-sk"), color.HiWhiteString("--screamingKebab"))
	fmt.Println("")

	/* Show the message at the end */
	os.Exit(1)
}

/* Function to check if file exists or not */
func fileExists(path string) bool {
    _, err := os.Stat(path)
    return !os.IsNotExist(err)
}

/* Function to fix file names */
func fixFileName(filePath string) {

	/* Some basic file info */
	fileDir := filepath.Dir(filePath)
	fileName := filepath.Base(filePath)
	fileExt := filepath.Ext(fileName)
	fileBaseName := strings.Replace(fileName, fileExt, "", 1)

	/* Convert the file name to specific case */
	var newFileName string

	switch (os.Args[1]) {
		case "-k","--kebab": newFileName = strcase.ToKebab(fileBaseName)
		case "-c","--camel": newFileName = strcase.ToCamel(fileBaseName)
		case "-s","--snake": newFileName = strcase.ToSnake(fileBaseName)
		case "-lc","--lowerCamel": newFileName = strcase.ToLowerCamel(fileBaseName)
		case "-ss","--screamingSnake": newFileName = strcase.ToScreamingSnake(fileBaseName)
		case "-sk","--screamingKebab": newFileName = strcase.ToScreamingKebab(fileBaseName)

		default: newFileName = strcase.ToKebab(fileBaseName)
	}

	/* New info about files */
	newFileName += fileExt
	newFilePath := fmt.Sprintf("%s/%s", fileDir, newFileName)

	/* Confirmation to rename or not */
	if (confirmRename != "a") {
		fmt.Printf("%s -> %s -> %s: ", color.HiRedString("[" + fileName + "]"), color.HiGreenString("[" + newFileName + "]"), color.HiWhiteString("[y|a|n]"))
		fmt.Scan(&confirmRename)

	} else {
		fmt.Printf("%s -> %s", color.HiRedString("[" + fileName + "]"), color.HiGreenString("[" + newFileName + "]"))
	}

	/* Rename the user provided files */
	if (confirmRename == "n") {
		return

	} else if (confirmRename == "y" || confirmRename == "a") {
		os.Rename(filePath, newFilePath)
	}
}

/* Function to fix file names */
func main() {

	/* Iterate through the arguments if there are enough arguments */
	if len(os.Args) > 2 {
		for _, f := range os.Args[2:] {

			/* Fix the file name if file exists */
			if fileExists(f) {
				fixFileName(f)

			/* Else show the help and exit */
			} else {
				plutoHelp()
			}
		}

	/* If invalid amount of arguments show help */
	} else {
		plutoHelp()
	}
}