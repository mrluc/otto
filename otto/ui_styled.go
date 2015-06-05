package otto

import (
	"bufio"
	"bytes"
	"strings"
	"unicode"
)

// StyledUi is a wrapper around an existing UI that automatically
// adds formatting around the UI text.
type StyledUi struct {
	Ui
}

func (u *StyledUi) Header(msg string) {
	u.Ui.Header(u.prefix("==> ", msg))
}

func (u *StyledUi) Message(msg string) {
	u.Ui.Message(u.prefix("    ", msg))
}

func (u *StyledUi) prefix(prefix, msg string) string {
	var buf bytes.Buffer

	scan := bufio.NewScanner(strings.NewReader(msg))
	for scan.Scan() {
		buf.WriteString(prefix)
		buf.WriteString(scan.Text() + "\n")
	}

	return strings.TrimRightFunc(buf.String(), unicode.IsSpace)
}