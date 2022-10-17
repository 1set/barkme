//go:build !windows
// +build !windows

package cmd

import cl "bitbucket.org/ai69/colorlogo"

var (
	colorLogoArt = cl.RoseWaterByColumn(logoArt)
	checkMark    = "✔"
)
