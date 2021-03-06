package main

import (
	"io/ioutil"

	"github.com/AndreKR/multiface"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/ewmh"
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/zachomedia/go-bdf"
)

func initX() error {
	// Set up a connection to the X server.
	var err error
	X, err = xgbutil.NewConn()
	if err != nil {
		return err
	}

	// Run the main X event loop, this is used to catch events.
	go xevent.Main(X)

	return nil
}

func initEWMH(w xproto.Window) error {
	// TODO: `WmStateSet` and `WmDesktopSet` are basically here to keep OpenBox
	// happy, can I somehow remove them and just use `_NET_WM_WINDOW_TYPE_DOCK`
	// like I can with WindowChef?
	if err := ewmh.WmWindowTypeSet(X, w, []string{
		"_NET_WM_WINDOW_TYPE_DOCK"}); err != nil {
		return err
	}
	if err := ewmh.WmStateSet(X, w, []string{
		"_NET_WM_STATE_STICKY"}); err != nil {
		return err
	}
	return ewmh.WmNameSet(X, w, "melonnotify")
}

func initFace() error {
	face = new(multiface.Face)

	fpl := []string{
		"/fonts/cure.punpun.bdf",
		"/fonts/kochi.small.bdf",
		"/fonts/baekmuk.small.bdf",
	}

	for _, fp := range fpl {
		f, err := runtime.Open(fp)
		if err != nil {
			return err
		}
		fb, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}
		ff, err := bdf.Parse(fb)
		if err != nil {
			return err
		}

		face.AddFace(ff.NewFace())

		f.Close()
	}

	return nil
}
