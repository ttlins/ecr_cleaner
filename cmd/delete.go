package cmd

import (
	"fmt"
	"log"

	awsecr "github.com/aws/aws-sdk-go/service/ecr"
	"github.com/spf13/cobra"
	"github.com/titolins/ecr_cleaner/internal/ecr"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes images by tag pattern",
	Long: `Delete deletes all images found for the given configuration matching the tag pattern
	NOTE: ignores normal SemVer tags (X.Y.Z)`,
	Run: func(cmd *cobra.Command, args []string) {
		re := compileUserPattern()
		c := ecr.NewClient(sess)

		imgs, err := c.ListImagesMatchingPattern(re)
		if err != nil {
			log.Fatalf("failed to list images: %v\n", err)
		}
		if len(imgs) == 0 {
			log.Printf("no images found. exiting..")
			return
		}

		displayImages(imgs)
		handleInput(c, imgs)
	},
}

func handleInput(c *ecr.Client, imgs []*awsecr.ImageIdentifier) {
	switch deletePrompt() {
	case "yes":
		if err := c.DeleteImages(imgs); err != nil {
			log.Fatalf("failed to delete images: %v\n", err)
		}
		log.Println("images deleted")
	case "no":
	default:
		fmt.Println("invalid choice")
	}
	log.Println("exiting ecr_cleaner..")
}
