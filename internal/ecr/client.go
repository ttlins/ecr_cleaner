package ecr

import (
	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/spf13/viper"
)

// Client holds all ECR logic
type Client struct {
	*ecr.ECR
}

// NewClient initializes and returns a Client
func NewClient(s *session.Session) *Client {
	return &Client{ecr.New(s)}
}

func (c *Client) listAll() (fi []*ecr.ImageIdentifier, err error) {
	out := &ecr.ListImagesOutput{NextToken: aws.String("go")}
	in := &ecr.ListImagesInput{
		Filter: &ecr.ListImagesFilter{
			TagStatus: aws.String("TAGGED"),
		},
		RepositoryName: aws.String(viper.GetString("repository")),
		MaxResults:     aws.Int64(1000),
	}

	for out.NextToken != nil {
		if out, err = c.ListImages(in); err != nil {
			return
		}

		fi = append(fi, out.ImageIds...)
		in.NextToken = out.NextToken
	}

	return
}

// ListImagesMatchingPattern gets all images for the given repository matching the passed parameter
func (c *Client) ListImagesMatchingPattern(re *regexp.Regexp) ([]*ecr.ImageIdentifier, error) {
	imgs, err := c.listAll()
	if err != nil {
		return nil, err
	}

	f := newImageFilter(matchFactory(re))

	return f.filterImagesIds(imgs), nil
}

// DeleteImages deletes the passed images
func (c *Client) DeleteImages(imgs []*ecr.ImageIdentifier) error {
	in := &ecr.BatchDeleteImageInput{
		ImageIds:       imgs,
		RepositoryName: aws.String(viper.GetString("repository")),
	}
	_, err := c.BatchDeleteImage(in)
	return err
}
