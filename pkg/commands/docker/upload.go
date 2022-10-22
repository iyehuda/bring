package docker

import (
	"errors"
	"fmt"
	"strings"

	"github.com/iyehuda/bring/pkg/docker"
	"github.com/iyehuda/bring/pkg/utils/commands"
	"github.com/spf13/cobra"
)

var errUploadTargetNotSet = errors.New("please specify an upload target")

type uploadOptions struct {
	uploadTarget string
}

// NewUploadCommand creates new docker upload command for extracting images from archive file,
// re-tagging, and pushing them to docker registry.
func NewUploadCommand() *cobra.Command {
	opts := &uploadOptions{}
	cmd := &cobra.Command{
		Use:     "upload <archive file> --to <target registry>",
		Short:   "Upload a set of docker images from an archive file",
		Long:    `This command depends on having docker installed and logged in to your registry (if necessary).`,
		PreRunE: validateUploadArgs(opts),
		RunE:    uploadImages(opts),
		Args:    cobra.ExactArgs(1),
	}

	cmd.Flags().StringVar(&opts.uploadTarget, "to", "", "target registry to upload (e.g. 'quay.io/iyehuda', 'quay.io/iyehuda/subpath', etc.)")

	return cmd
}

func validateUploadArgs(opts *uploadOptions) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if opts.uploadTarget == "" {
			return errUploadTargetNotSet
		}

		cmd.SilenceUsage = true

		return nil
	}
}

func uploadImages(opts *uploadOptions) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		retagger := docker.NewPreservePathImageRetagger(opts.uploadTarget)
		uploader := docker.NewImageUploader(retagger, &commands.LocalRunner{})

		images, err := uploader.UploadImages(args[0])
		if err != nil {
			return fmt.Errorf("failed to upload images: %w", err)
		}

		cmd.Printf("Uploaded images:\n%s\n", strings.Join(images, "\n"))

		return nil
	}
}
