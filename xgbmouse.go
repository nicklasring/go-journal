package main

import (
	"image"
	"log"

	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/mousebind"
	"github.com/BurntSushi/xgbutil/xevent"
)

func getMouseCoords() (p1 image.Point, p2 image.Point) {
	X, err := xgbutil.NewConn()
	if err != nil {
		log.Fatal(err)
	}

	mousebind.Initialize(X)

	cb1 := mousebind.ButtonPressFun(
		func(X *xgbutil.XUtil, e xevent.ButtonPressEvent) {
			p1 = image.Point{int(e.EventX), int(e.EventY)}
		})

	cb2 := mousebind.ButtonReleaseFun(
		func(X *xgbutil.XUtil, e xevent.ButtonReleaseEvent) {
			p2 = image.Point{int(e.EventX), int(e.EventY)}
			xevent.Quit(X)
		})

	err = cb1.Connect(X, X.RootWin(), "1", false, true)
	if err != nil {
		log.Fatal(err)
	}

	err = cb2.Connect(X, X.RootWin(), "1", false, true)
	if err != nil {
		log.Fatal(err)
	}

	xevent.Main(X)
	return p1, p2
}
