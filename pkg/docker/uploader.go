package docker

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/iyehuda/bring/pkg/utils/commands"
)

// ImageRetagger is an interface that is able to get an image name and return a modified version of it.
type ImageRetagger interface {
	Retag(string) string
}

// ImageUploader is able to load an image bundle, retag its contents and upload it to a specified registry target.
type ImageUploader struct {
	retagger ImageRetagger
	runner   commands.Runner
}

// NewImageUploader creates a new ImageUploader with a retagger and command runner.
func NewImageUploader(retagger ImageRetagger, runner commands.Runner) *ImageUploader {
	return &ImageUploader{retagger: retagger, runner: runner}
}

func (iu *ImageUploader) loadImageBundle(bundlePath string) ([]string, error) {
	cmd := exec.Command("docker", "load", "--input", bundlePath)
	output := &bytes.Buffer{}
	cmd.Stderr = io.Discard
	cmd.Stdout = output

	if err := iu.runner.Run(cmd); err != nil {
		return nil, fmt.Errorf("failed to run `docker load --input %s`: %w", bundlePath, err)
	}

	rawOutput := strings.TrimSpace(output.String())
	lines := strings.Split(rawOutput, "\n")
	loadedImages := make([]string, len(lines))

	for i, line := range lines {
		loadedImages[i] = strings.TrimPrefix(line, "Loaded image: ")
	}

	return loadedImages, nil
}

func (iu *ImageUploader) retagImage(oldImage string) (string, error) {
	newName := iu.retagger.Retag(oldImage)

	cmd := exec.Command("docker", "tag", oldImage, newName)
	if err := iu.runner.Run(cmd); err != nil {
		return "", fmt.Errorf("failed to run `docker tag %s %s`: %w", oldImage, newName, err)
	}

	return newName, nil
}

func (iu *ImageUploader) pushImage(image string) error {
	cmd := exec.Command("docker", "push", image)
	if err := iu.runner.Run(cmd); err != nil {
		return fmt.Errorf("failed to run `docker push %s`: %w", image, err)
	}

	return nil
}

// UploadImages takes an image bundle path, loads its images, retag them using the uploader retagger
// and uploads them to the uploader destination.
func (iu *ImageUploader) UploadImages(bundlePath string) ([]string, error) {
	sourceImages, err := iu.loadImageBundle(bundlePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load images from %s: %w", bundlePath, err)
	}

	targetImages := make([]string, len(sourceImages))

	for i, image := range sourceImages {
		newImage, err := iu.retagImage(image)
		if err != nil {
			return nil, fmt.Errorf("failed to retag image %s: %w", image, err)
		}

		targetImages[i] = newImage
	}

	for _, target := range targetImages {
		if err := iu.pushImage(target); err != nil {
			return nil, fmt.Errorf("failed to push image %s: %w", target, err)
		}
	}

	return targetImages, nil
}
