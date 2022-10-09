package docker

import (
	"testing"

	"github.com/iyehuda/bring/pkg/utils/commands"
	"github.com/iyehuda/bring/pkg/utils/tests"
	"github.com/stretchr/testify/mock"
)

func TestFetcher_Fetch(t *testing.T) {
	type fields struct {
		images      []string
		destination string
		runner      commands.Runner
	}

	failOnPull := new(tests.MockCommandRunner)
	failOnPull.On("Run", mock.MatchedBy(tests.CommandIncludes("docker pull"))).
		Return(&ImagePullError{Image: "alpine:not-exists", Err: nil})

	failOnSave := new(tests.MockCommandRunner)
	failOnSave.On("Run", mock.MatchedBy(tests.CommandIncludes("docker pull"))).
		Return(nil)
	failOnSave.On("Run", mock.MatchedBy(tests.CommandIncludes("docker save"))).
		Return(&ImageSaveError{Images: []string{"alpine:3.16"}, Destination: "/foo/bar", Err: nil})

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				images:      []string{"alpine:3.16"},
				destination: "/tmp/test",
				runner:      tests.NoOpCommandRunner,
			},
			wantErr: false,
		},
		{
			name: "Should fail - image not found",
			fields: fields{
				images:      []string{"alpine:not-exists"},
				destination: "/tmp/test",
				runner:      failOnPull,
			},
			wantErr: true,
		},
		{
			name: "Should fail - path not found",
			fields: fields{
				images:      []string{"alpine:3.16"},
				destination: "/foo/bar",
				runner:      failOnSave,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			df := &Fetcher{
				images:      tt.fields.images,
				destination: tt.fields.destination,
				runner:      tt.fields.runner,
			}
			if err := df.Fetch(); (err != nil) != tt.wantErr {
				t.Errorf("Fetcher.Fetch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
