package grifts

import (
  "github.com/gobuffalo/buffalo"
	"wordplay/wordofplay/actions"
)

func init() {
  buffalo.Grifts(actions.App())
}
