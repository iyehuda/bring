/*
Copyright © 2022 Yehuda Chikvashvili <yehudaac1@gmail.com>
*/
package docker

import "fmt"

type ImagePullError struct {
	Image string
	Err   error
}

func (e *ImagePullError) Error() string {
	return fmt.Sprintf("failed to pull image %v: %v", e.Image, e.Err)
}

func (e *ImagePullError) Unwrap() error {
	return e.Err
}

type ImageSaveError struct {
	Destination string
	Images      []string
	Err         error
}

func (e *ImageSaveError) Error() string {
	return fmt.Sprintf("failed to save images %v to %v: %v", e.Images, e.Destination, e.Err)
}

func (e *ImageSaveError) Unwrap() error {
	return e.Err
}
