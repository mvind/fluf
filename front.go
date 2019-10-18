package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mum4k/termdash/linestyle"

	"github.com/mum4k/termdash/terminal/terminalapi"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/terminal/termbox"
)

// We only need to redraw the screen when input changes happens right?
// i.e when we input characters in the search and so on..

// Create header (Search bar, date and time, session info)

const reDraw = 10 // Milliseconds until we refresh screen

func main() {
	t, err := termbox.New()
	if err != nil {
		fmt.Errorf("termbox.New => %v", err)
	}
	defer t.Close()

	c, err := container.New(
		t,
		container.SplitHorizontal(
			container.Top(
				container.SplitVertical(
					container.Left( // search
						container.Border(linestyle.Light),
					),
					container.Right( // time and date
						container.Border(linestyle.Light),
					),
					container.SplitPercent(80),
				),
			),
			container.Bottom(
				container.Border(linestyle.Light),
			),
			container.SplitPercent(10),
		),
	)
	if err != nil {
		fmt.Errorf("termbox.New => %v", err)
	}

	// ctx, cancel := context.WithTimeout(context.Background(), reDraw*time.Millisecond)
	// defer cancel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	quitter := func(k *terminalapi.Keyboard) {
		if k.Key == 'q' || k.Key == 'Q' {
			cancel()
		}
	}

	if err := termdash.Run(ctx, t, c, termdash.RedrawInterval(reDraw*time.Millisecond), termdash.KeyboardSubscriber(quitter)); err != nil {
		fmt.Errorf("termbox.New => %v", err)
	}
}
