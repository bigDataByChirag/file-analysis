package dir

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// get the working dir
func GetWorkingDir() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	// Navigate one level up to reach the parent directory and then add .env
	dirPath := filepath.Join(cwd, "..")
	return dirPath, nil
}

func GetEnvPath() (string, error) {
	// Get the current working directory
	workDir, err := GetWorkingDir()
	if err != nil {
		return "", err
	}
	envFilePath := filepath.Join(workDir, ".env")
	return envFilePath, nil
}

// find the largest dir in the whole directory
func FindLargestFile(name string) (int64, string, error) {

	fullDirPath, exists, err := FolderExists(name)
	if err != nil || !exists {
		return 0, "", fmt.Errorf("directory %s does not exist", name)
	}

	var largestFileSize int64
	var largestFileName string

	err = filepath.WalkDir(fullDirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		fileInfo, err := d.Info()
		if err != nil {
			return err
		}

		if !d.IsDir() {
			size := fileInfo.Size()
			if size > largestFileSize {
				largestFileSize = size
				largestFileName = fileInfo.Name()
			}
			//completeDirSize += fileInfo.Size()
		}
		return nil
	})
	if err != nil {
		return 0, "", err
	}

	return largestFileSize, largestFileName, nil
}

// get the total size of the directory
func CompleteDirSize(name string) (int64, error) {

	fullDirPath, exists, err := FolderExists(name)
	if err != nil && !exists {
		return 0, err
	}

	var completeDirSize int64

	err = filepath.WalkDir(fullDirPath, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fileInfo, err := info.Info()
		if err != nil {
			return err
		}
		completeDirSize += fileInfo.Size()
		return nil
	})
	if err != nil {
		return 0, fmt.Errorf("failed to locate directory: %s", name)
	}
	return completeDirSize, nil
}

func FolderExists(folderName string) (string, bool, error) {
	// Get the root directory path
	wd, err := GetWorkingDir()
	if err != nil {
		return "", false, err
	}

	// Initialize variables to store the path and found flag
	var foundPath string
	found := false

	// Walk through the directory tree to check for existence
	err = filepath.Walk(wd, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Check if the current path is a directory and matches the folderName
		if info.IsDir() && info.Name() == folderName {
			foundPath = path
			found = true
			return filepath.SkipDir // Skip further traversal since the directory is found
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error occurred while walking directory tree:", err)
		return "", false, err
	}

	return foundPath, found, nil
}
