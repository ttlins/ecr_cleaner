package cmd

import (
	"fmt"
	"log"
	"regexp"

	awsecr "github.com/aws/aws-sdk-go/service/ecr"
	"github.com/c-bata/go-prompt"
	"github.com/spf13/viper"
)

func displayImages(imgs []*awsecr.ImageIdentifier) {
	fmt.Println()
	fmt.Printf(
		"listing images for %q matching %q\n",
		viper.GetString("repository"),
		viper.GetString("pattern"),
	)
	fmt.Println("=========================")
	for i, img := range imgs {
		fmt.Printf("%d:\n", i)
		fmt.Printf("\tTag: %s\n", *img.ImageTag)
		fmt.Printf("\tDigest: %s\n", *img.ImageDigest)
	}
	fmt.Println()
}

func compileUserPattern() *regexp.Regexp {
	pattern := viper.GetString("pattern")
	if viper.GetBool("escape") {
		pattern = regexp.QuoteMeta(pattern)
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatalf(
			"failed to compile user provided regex %q: %v\n",
			pattern,
			err,
		)
	}
	return re
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "yes", Description: "Deletes all the images printed above"},
		{Text: "no", Description: "Cancels and exits ecr_cleaner"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func deletePrompt() string {
	return prompt.Input(
		"Confirm deletion? (yes / no) > ",
		completer,
		prompt.OptionPrefixBackgroundColor(prompt.Red),
		prompt.OptionPrefixTextColor(prompt.White),
	)
}
