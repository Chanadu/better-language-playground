package utils

import (
	"errors"
	"fmt"
	"syscall/js"
)

func ReportError(e error) {
	// _, _ = fmt.Fprintf(os.Stderr, "%s", color.RedString(fmt.Sprintf("Error: %s\n", e.Error())))
	js.Global().Call("goPrint", fmt.Sprintf("Error: %s\n", e.Error()), true)
}

func ReportDebugf(format string, args ...any) {
	// _, _ = fmt.Println(color.CyanString(format, args...))
	js.Global().Call("goPrint", fmt.Sprintf("%s\n", fmt.Sprintf(format, args...)), false)
}

func CreateErrorf(format string, args ...any) error {
	return errors.New(fmt.Sprintf("%s", fmt.Sprintf(format, args...)))
}

func CreateAndReportErrorf(format string, args ...any) {
	errorMessage := CreateErrorf(fmt.Sprintf(format, args...))
	ReportError(errorMessage)
}

func CreateScannerErrorf(line int, format string, args ...any) error {
	return errors.New(fmt.Sprintf("Line: %d, %s", line, fmt.Sprintf(format, args...)))
}

func CreateAndReportScannerErrorf(line int, format string, args ...any) {
	errorMessage := CreateScannerErrorf(line, format, args...)
	ReportError(errorMessage)
}

func CreateAndReportParsingErrorf(format string, args ...any) {
	errorMessage := fmt.Sprintf("Parsing: %s\n", fmt.Sprintf(format, args...))
	ReportError(errors.New(errorMessage))
}

// func CreateAndReportRuntimeErrorf(line int, format string, args ...any) {
// 	errorMessage := CreateRuntimeErrorf(line, fmt.Sprintf("Runtime: %s", format), args...)
// 	ReportError(errorMessage)
// }

func CreateRuntimeErrorf(line int, format string, args ...any) error {
	return errors.New(fmt.Sprintf("Runtime: Line: %d, %s", line, fmt.Sprintf(format, args...)))
}
