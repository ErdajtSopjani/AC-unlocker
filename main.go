package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Print("Assetto Corsa DLC Unlocker\nCode on: github.com/ErdajtSopjani/AC-unlocker\n\n")

	fmt.Print("Default root folder: 'C:/Program Files (x86)/Steam/steamapps/common/assettocorsa'\n(leave empty to use the default one)\n\n")

	fmt.Println("Enter the file path: ")
	var path string

	// get the root folder from the user
	fmt.Scanln(&path)
	if path == "" { // go with the default one if the user input is empty
		// C:/Program Files (x86)/Steam/steamapps/common/assettocorsa
		path = "C:/Program Files (x86)/Steam/steamapps/common/assettocorsa"
	}
	fmt.Println("\nRoot folder set to: ", path)

	// get all the folders in the cars folder
	fmt.Println("\nReading files in the cars folder...")
	folders, err := os.ReadDir(path + "/cars")
	if err != nil {
		fmt.Println(err)
		return
	}

	errMessage := "Make sure the game is closed and the file path is correct.\n\nIf the problem presits report it on github.com/ErdajtSopjani/AC-unlocker/issues\n\n"
	fmt.Print("\nDoing the magic...\n")
	// iterate through all the folders while reading the collider.kn5 files
	for _, folder := range folders {
		if folder.IsDir() {
			colliderFile, err := os.Open("./collider.kn5")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer colliderFile.Close() // This should stay here to close the file after the loop ends

			newColliderFile, err := os.Create(path + "/cars/" + folder.Name() + "/collider.kn5")
			if err != nil {
				fmt.Print("Something went wrong while creating the new collider file.\n", errMessage)
				colliderFile.Close()
				return
			}
			defer newColliderFile.Close()

			_, err = io.Copy(newColliderFile, colliderFile)
			if err != nil {
				fmt.Print("Something went wrong while copying the collider file.\n", errMessage)
				return
			}
		}
	}

	fmt.Println("Done, enjoy :)")
	fmt.Scanln()
}
