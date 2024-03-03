package parser

import (
	"fmt"
	"strings"
)

var tracingLevel int = 0

const traceIdentPlaceholder string = "\t"

func identLevel() string {
	return strings.Repeat(traceIdentPlaceholder, tracingLevel-1)
}

func tracePrint(fs string) {
	fmt.Printf("%s%s\n", identLevel(), fs)
}
func incIdentLevel() { tracingLevel = tracingLevel + 1 }
func decIdentLevel() { tracingLevel = tracingLevel - 1 }

func trace(msg string) string {
	incIdentLevel()
	tracePrint("BEGIN" + " " + msg)
	return msg
}

func untrace(msg string) {
	tracePrint("END" + " " + msg)
	decIdentLevel()
}
