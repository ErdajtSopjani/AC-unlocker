package main

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
)

//go:embed collider.kn5
var colliderData []byte

func main() {
	fmt.Print("Assetto Corsa DLC Unlocker\nCode on: github.com/ErdajtSopjani/AC-unlocker\n\n")

	fmt.Print("Default root folder: 'C:/Program Files (x86)/Steam/steamapps/common/assettocorsa/content'\n(leave empty to use the default one)\n\n")

	fmt.Println("Enter the file path: ")
	var path string

	// get the root folder from the user
	fmt.Scanln(&path)
	if path == "" { // go with the default one if the user input is empty
		path = "C:/Program Files (x86)/Steam/steamapps/common/assettocorsa/content"
	}

	// convert slashes to appropriate system format
	path = filepath.FromSlash(path)

	// cleanup path
	path, err := filepath.Abs(path)
	if err != nil {
		fmt.Println("Error resolving the path:", err)
		return
	}

	fmt.Println("\nRoot folder set to: ", path)

	// get all the folders in the cars folder
	fmt.Println("\nReading files in the cars folder...")
	folders, err := os.ReadDir(filepath.Join(path, "cars"))
	if err != nil {
		fmt.Println(err)
		return
	}

	errMessage := "Make sure the game is closed and the file path is correct.\n\nIf the problem persists report it on github.com/ErdajtSopjani/AC-unlocker/issues\n\n"

	fmt.Print("\nDoing the magic...\n")

	// iterate through all folders while copying collider data
	for _, folder := range folders {
		if folder.IsDir() {
			// create new collider file in directory
			newColliderFile, err := os.Create(filepath.Join(path, "cars", folder.Name(), "collider.kn5"))
			if err != nil {
				fmt.Print("Something went wrong while creating the new collider file.\n", errMessage)
				return
			}
			defer newColliderFile.Close()

			// write the embedded collider data to the new file
			_, err = newColliderFile.Write(colliderData)
			if err != nil {
				fmt.Print("Something went wrong while writing the collider file.\n", errMessage)
				return
			}
		}
	}

	fmt.Println("Done, enjoy :)\n\np.s give it a star")
	fmt.Scanln()
}
