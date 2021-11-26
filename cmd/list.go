package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/titolins/ecr_cleaner/internal/ecr"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists images by tag pattern",
	Long: `List lists all images found for the given configuration matching the tag pattern
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
	},
}
