/*
Copyright © 2022 Yehuda Chikvashvili <yehudaac1@gmail.com>
*/
package docker

import (
	"fmt"
	"os"
	"os/exec"
)

type Fetcher struct {
	images      []string
	destination string
}

func NewFetcher(images []string, destination string) *Fetcher {
	return &Fetcher{
		images:      images,
		destination: destination,
	}
}

func (df *Fetcher) Fetch() error {
	if err := df.pull(); err != nil {
		return fmt.Errorf("failed to fetch images: %w", err)
	}

	if err := df.save(); err != nil {
		return fmt.Errorf("failed to save images: %w", err)
	}

	return nil
}

func buildPullCommand(image string) *exec.Cmd {
	cmd := exec.Command("docker", "pull", image)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd
}

func (df *Fetcher) pull() error {
	for _, image := range df.images {
		cmd := buildPullCommand(image)
		err := cmd.Run()
		if err != nil {
			return &ImagePullError{
				Image: image,
				Err:   err,
			}
		}
	}

	return nil
}

func buildSaveCommand(destination string, images []string) *exec.Cmd {
	saveArgs := append([]string{"save", "--output", destination}, images...)
	cmd := exec.Command("docker", saveArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd
}

func (df *Fetcher) save() error {
	cmd := buildSaveCommand(df.destination, df.images)

	if err := cmd.Run(); err != nil {
		return &ImageSaveError{
			Destination: df.destination,
			Images:      df.images,
			Err:         err,
		}
	}

	return nil
}
