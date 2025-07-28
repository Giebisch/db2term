package config

import (
	"fmt"
	"os"
)

func EnsureDirExists(path string) error {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		// Directory does not exist, create it
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	} else if err != nil {
		// Some other error accessing the path
		return fmt.Errorf("error checking directory: %w", err)
	} else if !info.IsDir() {
		// Path exists but is not a directory
		return fmt.Errorf("path exists but is not a directory: %s", path)
	}

	// Directory exists or was successfully created
	return nil
}
