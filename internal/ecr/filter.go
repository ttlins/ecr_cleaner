package ecr

import (
	"log"
	"regexp"

	"github.com/aws/aws-sdk-go/service/ecr"
)

var normalSemVerMatch = matchFactory(
	regexp.MustCompile(`^[0-9]+\.[0-9]+\.[0-9]+$`),
)

type matchFunc func(s string) bool

func matchFactory(re *regexp.Regexp) matchFunc {
	return func(s string) bool {
		return re.MatchString(s)
	}
}

func filterImagesIds(imageIds []*ecr.ImageIdentifier, match matchFunc) (fi []*ecr.ImageIdentifier) {
	for _, img := range imageIds {
		if img == nil || img.ImageTag == nil {
			log.Printf("ImageIdentifier or ImageTag nil for %q. Skipping..\n", img.String())
			continue
		}

		if match(*img.ImageTag) && !normalSemVerMatch(*img.ImageTag) {
			fi = append(fi, img)
		}
	}

	return
}
