package docker

import (
	"testing"

	"github.com/iyehuda/bring/pkg/executils"
	"github.com/iyehuda/bring/pkg/testutils"
	"github.com/stretchr/testify/mock"
)

func TestFetcher_Fetch(t *testing.T) {
	t.Parallel()

	type fields struct {
		images      []string
		destination string
		runner      executils.Runner
	}

	failOnPull := new(testutils.MockCommandRunner)
	failOnPull.On("Run", mock.MatchedBy(testutils.CommandIncludes("docker pull"))).
		Return(&ImagePullError{Image: "alpine:not-exists", Err: nil})

	failOnSave := new(testutils.MockCommandRunner)
	failOnSave.On("Run", mock.MatchedBy(testutils.CommandIncludes("docker pull"))).
		Return(nil)
	failOnSave.On("Run", mock.MatchedBy(testutils.CommandIncludes("docker save"))).
		Return(&ImageSaveError{Images: []string{"alpine:3.16"}, Destination: "/foo/bar", Err: nil})

	noOpCommandRunner := &testutils.FakeCommandRunner{}

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
				runner:      noOpCommandRunner,
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
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
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
