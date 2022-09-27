/*
Copyright © 2022 Yehuda Chikvashvili <yehudaac1@gmail.com>
*/
package docker

import (
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
		return err
	}

	return df.save()
}

func (df *Fetcher) pull() error {
	for _, image := range df.images {
		cmd := exec.Command("docker", "pull", image)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			return err
		}
	}

	return nil
}

func (df *Fetcher) save() error {
	saveArgs := append([]string{"save", "--output", df.destination}, df.images...)
	cmd := exec.Command("docker", saveArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
