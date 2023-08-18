//go:build integration

package integration

import (
	"bytes"
	"path"
	"testing"
)

const (
	singleImageBundle      = "testdata/images/single.tar"
	multipleImageBundle    = "testdata/images/multiple.tar"
	invalidImageBundle     = "testdata/images/invalid.tar"
	nonexistentImageBundle = "testdata/images/nonexistent.tar"
)

func buildDownloadArgs(images []string, path string) []string {
	args := append([]string{"docker", "download"}, images...)

	if path != "" {
		args = append(args, "--to", path)
	}

	return args
}

func buildUploadArgs(path string, destination string) []string {
	args := []string{"docker", "upload"}

	if path != "" {
		args = append(args, path)
	}

	if destination != "" {
		args = append(args, "--to", destination)
	}

	return args
}

func TestDockerDownload(t *testing.T) {
	t.Parallel()

	type fields struct {
		images []string
		path   string
	}

	tempDir := t.TempDir()

	tests := []struct {
		name            string
		fields          fields
		wantErr         bool
		wantHelpMessage bool
	}{
		{
			name:            "Should succeed - single image",
			fields:          fields{images: []string{"busybox:1.35.0"}, path: path.Join(tempDir, "single.tar")},
			wantErr:         false,
			wantHelpMessage: false,
		},
		{
			name:            "Should succeed - multiple images",
			fields:          fields{images: []string{"alpine:3.16.2", "busybox:1.35.0"}, path: path.Join(tempDir, "multiple.tar")},
			wantErr:         false,
			wantHelpMessage: false,
		},
		{
			name:            "Should fail - no images given",
			fields:          fields{},
			wantErr:         true,
			wantHelpMessage: true,
		},
		{
			name:            "Should fail - no target given",
			fields:          fields{images: []string{"alpine:3.16.2", "busybox:1.35.0"}},
			wantErr:         true,
			wantHelpMessage: true,
		},
		{
			name:            "Should fail - image not found",
			fields:          fields{images: []string{"iyehuda/not:exists"}, path: path.Join(tempDir, "output.tar")},
			wantErr:         true,
			wantHelpMessage: false,
		},
		{
			name:            "Should fail - target path not found",
			fields:          fields{images: []string{"busybox:1.35.0"}, path: path.Join(tempDir, "not/exists.tar")},
			wantErr:         true,
			wantHelpMessage: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			output := &bytes.Buffer{}

			err := execute(
				buildDownloadArgs(tt.fields.images, tt.fields.path),
				withOutput(output),
				withError(output),
			)

			assertErrorIf(t, err, tt.wantErr)
			assertUsagePrintedIf(t, output, tt.wantHelpMessage)
		})
	}
}

func TestDockerUpload(t *testing.T) {
	t.Parallel()

	type fields struct {
		path        string
		destination string
	}

	tests := []struct {
		name            string
		fields          fields
		wantErr         bool
		wantHelpMessage bool
	}{
		{
			name:            "Should succeed - single image bundle",
			fields:          fields{path: singleImageBundle, destination: "registry.gitlab.com/iyehuda/bring"},
			wantErr:         false,
			wantHelpMessage: false,
		},
		{
			name:            "Should succeed - multiple images",
			fields:          fields{path: multipleImageBundle, destination: "registry.gitlab.com/iyehuda/bring"},
			wantErr:         false,
			wantHelpMessage: false,
		},
		{
			name:            "Should succeed - single image, target subpath",
			fields:          fields{path: singleImageBundle, destination: "registry.gitlab.com/iyehuda/bring/subpath"},
			wantErr:         false,
			wantHelpMessage: false,
		},
		{
			name:            "Should succeed - multiple images, target subpath",
			fields:          fields{path: multipleImageBundle, destination: "registry.gitlab.com/iyehuda/bring/subpath"},
			wantErr:         false,
			wantHelpMessage: false,
		},
		{
			name:            "Should fail - no input file path given",
			fields:          fields{destination: "registry.gitlab.com/iyehuda/bring"},
			wantErr:         true,
			wantHelpMessage: true,
		},
		{
			name:            "Should fail - no target given",
			fields:          fields{path: singleImageBundle},
			wantErr:         true,
			wantHelpMessage: true,
		},
		{
			name:            "Should fail - invalid input file",
			fields:          fields{path: invalidImageBundle, destination: "registry.gitlab.com/iyehuda/bring"},
			wantErr:         true,
			wantHelpMessage: false,
		},
		{
			name:            "Should fail - input file not found",
			fields:          fields{path: nonexistentImageBundle, destination: "registry.gitlab.com/iyehuda/bring"},
			wantErr:         true,
			wantHelpMessage: false,
		},
		{
			name:            "Should fail - registry not found",
			fields:          fields{path: singleImageBundle, destination: "somenonexistentregistry.io/iyehuda"},
			wantErr:         true,
			wantHelpMessage: false,
		},
		{
			name:            "Should fail - registry path not found",
			fields:          fields{path: singleImageBundle, destination: "registry.gitlab.com/iyehuda/not-exists"},
			wantErr:         true,
			wantHelpMessage: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			output := &bytes.Buffer{}

			err := execute(
				buildUploadArgs(tt.fields.path, tt.fields.destination),
				withOutput(output),
				withError(output),
			)

			assertErrorIf(t, err, tt.wantErr)
			assertUsagePrintedIf(t, output, tt.wantHelpMessage)
		})
	}
}
