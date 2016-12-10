package main

import clipboardMgr "github.com/atotto/clipboard"

type clipboardWriter interface {
	writeToClipboard(s string) error
}

type clipboard struct{}

func (cb *clipboard) writeToClipboard(s string) error {
	return clipboardMgr.WriteAll(s)
}
