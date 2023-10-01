/*
Package img_get provides functions to download
	images from URLs or copy images from local paths.

Usage:
1. Import the package:
   import "github.com/jstevens8185/img_get"

2. To download an image from a URL:
   sourceURL := "https://example.com/image.jpg"
   outputPath := "downloaded_image.jpg"
   err := img_get.GetImage(sourceURL, "", outputPath)

3. To copy an image from a local path:
   localPath := "/path/to/local/image.jpg"
   outputPath := "copied_image.jpg"
   err := img_get.GetImage("", localPath, outputPath)

4. Handling errors:
   The GetImage function returns an error if any
	  issues occur during the download or copying process.
*/

package img_get

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// GetImage downloads an image from a URL or copies it from a local path
// and saves it to a specified file path.
func GetImage(sourceURL, localPath, outputPath string) error {
	var reader io.Reader

	// Check if the source is a URL or a local file
	if sourceURL != "" {
		// Source is a URL
		response, err := http.Get(sourceURL)
		if err != nil {
			return fmt.Errorf("error making the request: %v", err)
		}
		defer response.Body.Close()

		// Check if the response status code is OK (200)
		if response.StatusCode != http.StatusOK {
			return fmt.Errorf("error: status code %d", response.StatusCode)
		}

		reader = response.Body
	} else if localPath != "" {
		// Source is a local file
		file, err := os.Open(localPath)
		if err != nil {
			return fmt.Errorf("error opening the local file: %v", err)
		}
		defer file.Close()

		reader = file
	} else {
		return fmt.Errorf("both source URL and local path are empty")
	}

	// Create the directory structure if it doesn't exist
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("error creating directory: %v", err)
	}

	// Create a new file to save the image
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating the file: %v", err)
	}
	defer outputFile.Close()

	// Copy the image data to the file
	_, err = io.Copy(outputFile, reader)
	if err != nil {
		return fmt.Errorf("error saving the image: %v", err)
	}

	log.Printf("Image downloaded/saved as '%s'\n", outputPath)
	return nil
}
