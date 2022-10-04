/*
Copyright © 2022 Yehuda Chikvashvili <yehudaac1@gmail.com>
*/
package docker

import "fmt"

// ImagePullError represents an error encountered when pulling a docker image
type ImagePullError struct {
	Image string
	Err   error
}

// Error implements the error interface for formatting an error message
func (e *ImagePullError) Error() string {
	return fmt.Sprintf("failed to pull image %v: %v", e.Image, e.Err)
}

// Unwrap implements the Unwrap interface for retrieving inner error
func (e *ImagePullError) Unwrap() error {
	return e.Err
}

// ImageSaveError represents an error encountered when saving docker images to file
type ImageSaveError struct {
	Destination string
	Images      []string
	Err         error
}

// Error implements the error interface for formatting an error message
func (e *ImageSaveError) Error() string {
	return fmt.Sprintf("failed to save images %v to %v: %v", e.Images, e.Destination, e.Err)
}

// Unwrap implements the Unwrap interface for retrieving inner error
func (e *ImageSaveError) Unwrap() error {
	return e.Err
}
