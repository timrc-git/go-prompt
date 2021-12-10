package main

import (
	"fmt"
	"os"

	prompt "github.com/timrc-git/go-prompt"
)

var p *prompt.Prompt

var LivePrefixState struct {
	LivePrefix string
	IsEnable   bool
}

func executor(in string) {
	if p.WasCanceled() {
		fmt.Println("(canceled)")
	} else {
		fmt.Println("Your input: " + in)
	}
	if in == "" {
		LivePrefixState.IsEnable = false
		LivePrefixState.LivePrefix = in
		return
	}
	if in == "pw" {
		p.Obscure(true)
		in = "((password))"
	}
	if in == "exit" {
		os.Exit(0)
	}
	LivePrefixState.LivePrefix = in + "> "
	LivePrefixState.IsEnable = true
}

func completer(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "users", Description: "Store the username and age"},
		{Text: "articles", Description: "Store the article text posted by user"},
		{Text: "comments", Description: "Store the text commented to articles"},
		{Text: "groups", Description: "Combine users with specific rules"},
		{Text: "pw", Description: "Super-secret password entry"},
		{Text: "exit", Description: "Does what it says"},
	}
	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func changeLivePrefix() (string, bool) {
	return LivePrefixState.LivePrefix, LivePrefixState.IsEnable
}

func main() {
	p = prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(">>> "),
		prompt.OptionLivePrefix(changeLivePrefix),
		prompt.OptionTitle("live-prefix-example"),
	)
	p.Run()
}
