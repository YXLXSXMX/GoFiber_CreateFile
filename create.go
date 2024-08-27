package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var createFilesCmd = &cobra.Command{
	Use:   "cf [base_filename]",
	Short: "Create Go files with the specified base filename",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		baseFilename := args[0]

		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error getting current directory: %v\n", err)
			return
		}

		folderName := filepath.Join(currentDir, fmt.Sprintf("%s", baseFilename))

		if err := os.MkdirAll(folderName, os.ModePerm); err != nil {
			fmt.Printf("Error creating folder %s: %v\n", folderName, err)
			return
		}

		files := []string{
			fmt.Sprintf("%s/%s_databases.go", folderName, baseFilename),
			fmt.Sprintf("%s/%s_handlers.go", folderName, baseFilename),
			fmt.Sprintf("%s/%s_model.go", folderName, baseFilename),
			fmt.Sprintf("%s/%s_routes.go", folderName, baseFilename),
		}

		for _, file := range files {
			if _, err := os.Stat(file); err == nil {
				fmt.Printf("File %s already exists. Skipping...\n", file)
			} else {
				content := fmt.Sprintf("package %s", baseFilename)
				if err := os.WriteFile(file, []byte(content), 0644); err != nil {
					fmt.Printf("Error creating file %s: %v\n", file, err)
				} else {
					fmt.Printf("Successfully created file: %s\n", file)
				}
			}
		}
	},
}

func main() {
	var rootCmd = &cobra.Command{Use: "myapp"}
	rootCmd.AddCommand(createFilesCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
