package main

// mirrorAsset mirrors them from the source URL to the local destination path
func mirrorAsset(source string, destination string, overwrite bool) error {
	// Check to make sure the destination directory exists
	parentDirectory := GetParentDirectory(destination)

	// Check if the directory exists
	directoryCheck, err := DirectoryExists(parentDirectory)
	// Check for errors
	if err != nil {
		return err
	}

	// If the directory doesn't exist, create it
	if !directoryCheck {
		CreateDirectory(parentDirectory)
	}

	// Check if the directory is writable
	if IsWritable(parentDirectory) {
		// Directory exists and is writable - continue

		// Check if the file exists
		fileCheck, err := FileExists(destination)
		// Check for errors
		if err != nil {
			logStdErr("Error checking path existence: " + destination)
			return err
		}

		// If the file doesn't exist or overwrite is true, download it from the source
		if !fileCheck || overwrite {
			if overwrite {
				logStdOut("OVERWRITING FILE: " + destination)
			}
			// Download the file
			err := DownloadFile(destination, source)
			// Check for errors
			if err != nil {
				logStdErr("Error downloading file: " + source)
				return err
			}
		} else {
			// File exists
			logStdOut("The file already exists: " + destination)
		}
	} else {
		// Directory exists but is not writable
		return Stoerr("Directory is not writable!")
	}

	// Finished
	return nil
}
