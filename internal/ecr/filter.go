package ecr

import (
	"log"
	"regexp"

	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/spf13/viper"
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

type imageFilter struct {
	matchers []matchFunc
}

func newImageFilter(matchers ...matchFunc) *imageFilter {
	if viper.GetBool("ignoreNormal") {
		matchers = append(matchers, notMatcher(normalSemVerMatch))
	}
	return &imageFilter{
		matchers: matchers,
	}
}

func (f *imageFilter) filterImagesIds(imageIds []*ecr.ImageIdentifier) (fi []*ecr.ImageIdentifier) {
	for _, img := range imageIds {
		if img == nil || img.ImageTag == nil {
			log.Printf("ImageIdentifier or ImageTag nil for %q. Skipping..\n", img.String())
			continue
		}

		if f.match(*img.ImageTag) {
			fi = append(fi, img)
		}
	}

	return
}

func (f *imageFilter) match(s string) bool {
	for _, m := range f.matchers {
		if !m(s) {
			return false
		}
	}

	return true
}

func notMatcher(matcher matchFunc) matchFunc {
	return func(s string) bool {
		return !matcher(s)
	}
}
