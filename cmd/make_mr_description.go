package main

import (
	"flag"
	"fmt"

	"github.com/mski-iksm/mrgpt/pkg/ai"
	"github.com/mski-iksm/mrgpt/pkg/description"
	"github.com/mski-iksm/mrgpt/pkg/git"
)

func main() {
	var openaiApiKey *string = flag.String("openai-api-key", "DEFAULT", "openai api key")

	flag.Parse()

	// masterとのdiffを取得
	gitDiff := git.GitDiff()

	// chatGPTでdescriptionを生成
	mrDescription := ai.MakeDescription(gitDiff, *openaiApiKey)

	// descriptionをprint
	fmt.Println(mrDescription)

	// clipboardにコピー
	description.CopyToClipboard(mrDescription)
}
