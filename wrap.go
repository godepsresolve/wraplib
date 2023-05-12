package wraplib

import (
	"fmt"
	"github.com/godepsresolve/corelib"
	"github.com/godepsresolve/helperlib"
)

func Wrap(input string) string {
	format := "%s@v%s -> %s"
	helperResult := fmt.Sprintf(format, pkg, version, helperlib.ProvideAssistance(input))
	corelibResult := fmt.Sprintf(format, pkg, version, corelib.Format(input))
	return fmt.Sprintf("%s\n%s", helperResult, corelibResult)
}
