package docker

import "testing"

func TestPreservePathImageRetagger_Retag(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		source string
		target string
		want   string
	}{
		{
			name:   "from Docker Hub (official, no tag) to Docker Hub (community)",
			source: "alpine",
			target: "example",
			want:   "example/alpine",
		},
		{
			name:   "from Docker Hub (official, with tag) to Docker Hub (community)",
			source: "alpine:3.16.2",
			target: "example",
			want:   "example/alpine:3.16.2",
		},
		{
			name:   "from Docker Hub (community, no tag) to Docker Hub (community)",
			source: "nested/alpine",
			target: "example",
			want:   "example/nested/alpine",
		},
		{
			name:   "from Docker Hub (community, with tag) to Docker Hub (community)",
			source: "nested/alpine:3.16.2",
			target: "example",
			want:   "example/nested/alpine:3.16.2",
		},
		{
			name:   "from private registry (flat, no tag) to Docker Hub (community)",
			source: "quay.io/alpine",
			target: "example",
			want:   "example/alpine",
		},
		{
			name:   "from private registry (flat, with tag) to Docker Hub (community)",
			source: "quay.io/alpine:3.16.2",
			target: "example",
			want:   "example/alpine:3.16.2",
		},
		{
			name:   "from private registry (nested, no tag) to Docker Hub (community)",
			source: "quay.io/nested/alpine",
			target: "example",
			want:   "example/nested/alpine",
		},
		{
			name:   "from private registry (nested, with tag) to Docker Hub (community)",
			source: "quay.io/nested/alpine:3.16.2",
			target: "example",
			want:   "example/nested/alpine:3.16.2",
		},
		{
			name:   "from Docker Hub (official, no tag) to private registry (naked)",
			source: "alpine",
			target: "gcr.io",
			want:   "gcr.io/alpine",
		},
		{
			name:   "from Docker Hub (official, with tag) to private registry (naked)",
			source: "alpine:3.16.2",
			target: "gcr.io",
			want:   "gcr.io/alpine:3.16.2",
		},
		{
			name:   "from Docker Hub (community, no tag) to private registry (naked)",
			source: "nested/alpine",
			target: "gcr.io",
			want:   "gcr.io/nested/alpine",
		},
		{
			name:   "from Docker Hub (community, with tag) to private registry (naked)",
			source: "nested/alpine:3.16.2",
			target: "gcr.io",
			want:   "gcr.io/nested/alpine:3.16.2",
		},
		{
			name:   "from private registry (flat, no tag) to private registry (naked)",
			source: "quay.io/alpine",
			target: "gcr.io",
			want:   "gcr.io/alpine",
		},
		{
			name:   "from private registry (flat, with tag) to private registry (naked)",
			source: "quay.io/alpine:3.16.2",
			target: "gcr.io",
			want:   "gcr.io/alpine:3.16.2",
		},
		{
			name:   "from private registry (nested, no tag) to private registry (naked)",
			source: "quay.io/nested/alpine",
			target: "gcr.io",
			want:   "gcr.io/nested/alpine",
		},
		{
			name:   "from private registry (nested, with tag) to private registry (naked)",
			source: "quay.io/nested/alpine:3.16.2",
			target: "gcr.io",
			want:   "gcr.io/nested/alpine:3.16.2",
		},
		{
			name:   "from Docker Hub (official, no tag) to private registry (flat)",
			source: "alpine",
			target: "gcr.io/example",
			want:   "gcr.io/example/alpine",
		},
		{
			name:   "from Docker Hub (official, with tag) to private registry (flat)",
			source: "alpine:3.16.2",
			target: "gcr.io/example",
			want:   "gcr.io/example/alpine:3.16.2",
		},
		{
			name:   "from Docker Hub (community, no tag) to private registry (flat)",
			source: "nested/alpine",
			target: "gcr.io/example",
			want:   "gcr.io/example/nested/alpine",
		},
		{
			name:   "from Docker Hub (community, with tag) to private registry (flat)",
			source: "nested/alpine:3.16.2",
			target: "gcr.io/example",
			want:   "gcr.io/example/nested/alpine:3.16.2",
		},
		{
			name:   "from private registry (flat, no tag) to private registry (flat)",
			source: "quay.io/alpine",
			target: "gcr.io/example",
			want:   "gcr.io/example/alpine",
		},
		{
			name:   "from private registry (flat, with tag) to private registry (flat)",
			source: "quay.io/alpine:3.16.2",
			target: "gcr.io/example",
			want:   "gcr.io/example/alpine:3.16.2",
		},
		{
			name:   "from private registry (nested, no tag) to private registry (flat)",
			source: "quay.io/nested/alpine",
			target: "gcr.io/example",
			want:   "gcr.io/example/nested/alpine",
		},
		{
			name:   "from private registry (nested, with tag) to private registry (flat)",
			source: "quay.io/nested/alpine:3.16.2",
			target: "gcr.io/example",
			want:   "gcr.io/example/nested/alpine:3.16.2",
		},
		{
			name:   "from Docker Hub (official, no tag) to private registry (nested)",
			source: "alpine",
			target: "gcr.io/example/images",
			want:   "gcr.io/example/images/alpine",
		},
		{
			name:   "from Docker Hub (official, with tag) to private registry (nested)",
			source: "alpine:3.16.2",
			target: "gcr.io/example/images",
			want:   "gcr.io/example/images/alpine:3.16.2",
		},
		{
			name:   "from Docker Hub (community, no tag) to private registry (nested)",
			source: "nested/alpine",
			target: "gcr.io/example/images",
			want:   "gcr.io/example/images/nested/alpine",
		},
		{
			name:   "from Docker Hub (community, with tag) to private registry (nested)",
			source: "nested/alpine:3.16.2",
			target: "gcr.io/example/images",
			want:   "gcr.io/example/images/nested/alpine:3.16.2",
		},
		{
			name:   "from private registry (flat, no tag) to private registry (nested)",
			source: "quay.io/alpine",
			target: "gcr.io/example/images",
			want:   "gcr.io/example/images/alpine",
		},
		{
			name:   "from private registry (flat, with tag) to private registry (nested)",
			source: "quay.io/alpine:3.16.2",
			target: "gcr.io/example/images",
			want:   "gcr.io/example/images/alpine:3.16.2",
		},
		{
			name:   "from private registry (nested, no tag) to private registry (nested)",
			source: "quay.io/nested/alpine",
			target: "gcr.io/example/images",
			want:   "gcr.io/example/images/nested/alpine",
		},
		{
			name:   "from private registry (nested, with tag) to private registry (nested)",
			source: "quay.io/nested/alpine:3.16.2",
			target: "gcr.io/example/images",
			want:   "gcr.io/example/images/nested/alpine:3.16.2",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ir := &PreservePathImageRetagger{
				prefix: tt.target,
			}
			if got := ir.Retag(tt.source); got != tt.want {
				t.Errorf("ImageRetagger.Retag() = %v, want %v", got, tt.want)
			}
		})
	}
}
