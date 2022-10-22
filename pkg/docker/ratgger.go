package docker

import (
	"fmt"
	"strings"
)

type PreservePathImageRetagger struct {
	prefix string
}

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

func (ir *PreservePathImageRetagger) Retag(image string) string {
	relativeName := stripRegistryFromImage(image)

	return fmt.Sprintf("%s/%s", ir.prefix, relativeName)
}
