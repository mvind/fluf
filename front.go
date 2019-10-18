package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/terminal/termbox"
)

// func main() {
// 	t, err := termbox.New()
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer t.Close()

// 	bc, err := barchart.New()
// 	if err != nil {
// 		panic(err)
// 	}

// 	values := []int{1, 2, 3, 4, 5}
// 	max := 10
// 	if err := bc.Values(values, max); err != nil {
// 		panic(err)
// 	}

// 	c, err := container.New(
// 		t,
// 		container.Border(linestyle.Light),
// 		container.BorderTitle("PRESS Q TO QUIT"),
// 		container.PlaceWidget(bc),
// 	)
// 	if err != nil {
// 		panic(err)
// 	}

// 	quitter := func(k *terminalapi.Keyboard) {
// 		if k.Key == 'q' || k.Key == 'Q' {
// 			cancel()
// 		}
// 	}

// 	if err := termdash.Run(ctx, t, c, termdash.KeyboardSubscriber(quitter)); err != nil {
// 		panic(err)
// 	}

// }

func main() {
	t, err := termbox.New()
	if err != nil {
		fmt.Errorf("termbox.New => %v", err)
	}

	defer t.Close()

	c, err := container.New(t)
	if err != nil {
		fmt.Errorf("termbox.New => %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := termdash.Run(ctx, t, c); err != nil {
		fmt.Errorf("termbox.New => %v", err)
	}
}
