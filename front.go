package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/mum4k/termdash/widgets/text"

	"github.com/mum4k/termdash/linestyle"

	"github.com/mum4k/termdash/terminal/terminalapi"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/terminal/termbox"
)

// We only need to redraw the screen when input changes happens right?
// i.e when we input characters in the search and so on..

// Create header (Search bar, date and time, session info)
// How to add widgets??
// Make make custom colors

// Create segment display
// call go func(ctx, obj) and then its works??

func clock(ctx context.Context, t *text.Text) {
	ticker := time.NewTicker(time.Second) // Channel to time every second
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			now := time.Now() // New tick
			nowStr := now.Format("15 04")
			parts := strings.Split(nowStr, " ")

			spacer := ":"
			if now.Second()%2 == 0 {
				spacer = " "
			}

			date := now.Format("Jan 2 Mon")

			//fmt.Println(date)

			t.Reset() // First delete text buffer

			if err := t.Write(parts[0] + spacer + parts[1] + " - " + date); err != nil {
				panic(err)
			}

		case <-ctx.Done():
			return
		}
	}
}

const reDraw = 250 * time.Millisecond

func main() {
	t, err := termbox.New()
	if err != nil {
		fmt.Errorf("termbox.New => %v", err)
	}
	defer t.Close()

	ctx, cancel := context.WithCancel(context.Background())

	clockCon, err := text.New()
	if err != nil {
		panic(err)
	}
	//clockCon.Write("hej")
	go clock(ctx, clockCon)

	// Layout
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
						container.PlaceWidget(clockCon),
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

	quitter := func(k *terminalapi.Keyboard) {
		if k.Key == 'q' || k.Key == 'Q' {
			cancel()
		}
	}

	if err := termdash.Run(ctx, t, c, termdash.RedrawInterval(reDraw), termdash.KeyboardSubscriber(quitter)); err != nil {
		fmt.Errorf("termbox.New => %v", err)
	}
}
