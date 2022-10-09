//go:build integration

package integration

import (
	"bytes"
	"testing"
)

func buildDownloadArgs(images []string, path string) []string {
	args := append([]string{"docker", "download"}, images...)

	if path != "" {
		args = append(args, "--to", path)
	}

	return args
}

func TestDockerDownload(t *testing.T) {
	type fields struct {
		images []string
		path   string
	}

	tests := []struct {
		name            string
		fields          fields
		wantErr         bool
		wantHelpMessage bool
	}{
		{
			name:            "Should succeed - single image",
			fields:          fields{images: []string{"alpine:3.16.2"}, path: "/tmp/single.tar"},
			wantErr:         false,
			wantHelpMessage: false,
		},
		{
			name:            "Should succeed - multiple images",
			fields:          fields{images: []string{"alpine:3.16.2", "redis:7.0.5-alpine"}, path: "/tmp/multiple.tar"},
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
			fields:          fields{images: []string{"alpine:3.16.2", "redis:7.0.5-alpine"}},
			wantErr:         true,
			wantHelpMessage: true,
		},
		{
			name:            "Should fail - image not found",
			fields:          fields{images: []string{"iyehuda/not:exists"}, path: "/tmp/a"},
			wantErr:         true,
			wantHelpMessage: false,
		},
		{
			name:            "Should fail - target path not found",
			fields:          fields{images: []string{"alpine:3.16.2"}, path: "/tmp/not/exists"},
			wantErr:         true,
			wantHelpMessage: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			output := &bytes.Buffer{}

			err := execute(
				buildDownloadArgs(tt.fields.images, tt.fields.path),
				withOutput(output),
				withError(output),
			)

			assertErrorIf(t, err, tt.wantErr)
			assertErrorPrintedIf(t, output, tt.wantHelpMessage)
		})
	}
}
