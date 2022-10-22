package docker

import (
	"fmt"
	"strings"
)

// PreservePathImageRetagger enables to re-tag a docker image under a different registry target.
// The re-tagging preserves the source relative path under the source registry and appends it to the target path.
// Examples:
//
//	r := NewPreservePathImageRetagger("quay.io/iyehuda")
//	r.Retag("alpine:3.16.2") // would give "quay.io/iyehuda/alpine:3.16.2"
//	r.Retag("grafana/grafana:latest") // would give "quay.io/iyehuda/grafana/grafana:latest"
//
//	r := NewPreservePathImageRetagger("yehudac.jfrog.io/my/images")
//	r.Retag("quay.io/argoproj/argocd:2.3.4") // would give "yehudac.jfrog.io/my/images/argoproj/argocd:2.3.4"
type PreservePathImageRetagger struct {
	prefix string
}

// NewPreservePathImageRetagger creates a new PreservePathImageRetagger with the given registry prefix.
func NewPreservePathImageRetagger(prefix string) *PreservePathImageRetagger {
	return &PreservePathImageRetagger{
		prefix: prefix,
	}
}

func stripRegistryFromImage(image string) string {
	registry, relativeName, ok := strings.Cut(image, "/")

	if !ok || !strings.ContainsAny(registry, ".:") {
		return image
	}

	return relativeName
}

// Retag takes an image and returns its new name under the retagger prefix.
// The image parameter could be any legal docker image name.
// The output will preserve the original image relative path under its registry.
func (ir *PreservePathImageRetagger) Retag(image string) string {
	relativeName := stripRegistryFromImage(image)

	return fmt.Sprintf("%s/%s", ir.prefix, relativeName)
}
