package getterminal

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// GetTerminal teturns the best terminal program option.
// Defaults to powershell on Windows, or runs through lists
// of known terminals on other platforms.
func GetTerminal() string {
	if runtime.GOOS == "windows" {
		return "powershell.exe"
	}

	if runtime.GOOS == "darwin" {
		return getTermForMac()
	}

	return getBestTermForUnix()
}

func commandExists(cmd string) bool {
	if runtime.GOOS == "darwin" {
		return commandExistsMac(cmd)
	}

	return commandExistsOther(cmd)
}

func commandExistsOther(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func commandExistsMac(cmd string) bool {
	s := fmt.Sprintf(`osascript -e 'id of application "%s"'`, cmd)
	c := exec.Command("bash", "-c", s)
	if err := c.Run(); err != nil {
		return false
	}

	return true
}

func getTermForMac() string {
	macTerms := []string{
		"Hyper",
		"Alacritty",
		"kitty",
		"iterm2",
		"iTerm",
		"Terminal",
	}

	for _, term := range macTerms {
		if commandExists(term) {
			return term
		}
	}

	return "Terminal"
}

func getBestTermForUnix() string {
	unixTerms := []string{
		os.Getenv("TERMINAL"),
		"x-terminal-emulator", // debian
		"terminal",            // arch, i think?

		// fancy terminals
		"Eterm",
		"alacritty",
		"aterm",
		"cool-retro-term",
		"eterm",
		"finalterm",
		"gnome-terminal",
		"guake",
		"hyper",
		"kitty",
		"konsole",
		"lilyterm",
		"lxterminal",
		"mate-terminal",
		"mrxvt",
		"qterminal",
		"roxterm",
		"sakura",
		"terminator",
		"terminix",
		"terminology",
		"termit",
		"termite",
		"termkit",
		"tilda",
		"tilix",
		"wterm",
		"xfce4-terminal",
		"yakuake",

		// less fancy terminals
		"stterm", // debian
		"st",     // not-debian -- might conflict with the st server package...

		// defaults
		"urxvt",
		"uxterm",
		"rxvt",
		"xterm",
		os.Getenv("COLORTERM"),
		os.Getenv("XTERM"),
	}

	for _, term := range unixTerms {
		if commandExists(term) {
			return term
		}
	}

	return "xterm"
}
