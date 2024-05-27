package orchester

import (
	"animatic-v2/internal/argparse"
	"animatic-v2/internal/debug"
	"animatic-v2/internal/tui"
)

func Start(args argparse.Args) {
	animeName, err := tui.SearchAnime()
	if err != nil {
		debug.LogError("catch the anime from download", err)
	}

	println(animeName)
}
