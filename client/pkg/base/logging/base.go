package logging

import (
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func LoggerSetup() {
	// Global Time Format
	zerolog.TimeFieldFormat = "2 January 2006 Monday, 15:04:05 -07"

	// Global Path Format
	// Trims the file path, writes the path after backend folder and line.
	// If you want full path, just remove it. Default attitude is full path + line number
	zerolog.CallerMarshalFunc = func(file string, line int) string {
		return path.Join(strings.Split(path.Dir(file), "/wow/")[1], path.Base(file)) + ":" + strconv.Itoa(line)
	}

	log.Logger = prettyLoggerWithColor()
}

// // Formatting and Customization for json output: Caller (address of log) and Timestamp added.
// func jsonLogger() (jsonLogger zerolog.Logger) {
// 	jsonLogger = zerolog.New(os.Stderr).With().Caller().Timestamp().Logger()
// 	return jsonLogger
// }

// // Formatting and Customization for console output.
// // --- Colorless ---
// func prettyLogger() (prettyLogger zerolog.Logger) {
// 	// Caller (address of log) and Timestamp added.
// 	prettyLogger = zerolog.New(os.Stderr).With().Caller().Timestamp().Logger()

// 	// Output Format
// 	prettyLogger = prettyLogger.Output(
// 		zerolog.ConsoleWriter{
// 			Out:        os.Stderr,
// 			NoColor:    true, // Color Configuration
// 			TimeFormat: "15:04",
// 			FormatLevel: func(i interface{}) string {
// 				if i == nil {
// 					i = " log"
// 				}
// 				return strings.ToUpper(fmt.Sprintf("| %-6s|", i))

// 			},
// 			FormatCaller: func(i interface{}) string {
// 				return fmt.Sprintf("%s |", i)
// 			},
// 			FormatMessage: func(i interface{}) string {
// 				if i == nil {
// 					return ""
// 				}
// 				return fmt.Sprintf("Message: '%s' |", i)

// 			},
// 			FormatErrFieldName: func(i interface{}) string {
// 				return strings.Title(fmt.Sprintf("%s: ", i))
// 			},
// 			FormatErrFieldValue: func(i interface{}) string {
// 				return fmt.Sprintf("%s |", i)
// 			},
// 			FormatFieldValue: func(i interface{}) string {
// 				return fmt.Sprintf("'%s' |", i)
// 			},
// 		})
// 	return prettyLogger
// }

// Formatting and Customization for console output.
// --- Colorfull ---
func prettyLoggerWithColor() (prettyLogger zerolog.Logger) {
	// Caller (address of log) and Timestamp added.
	prettyLogger = zerolog.New(os.Stderr).With().Caller().Timestamp().Logger()

	// Output Format
	prettyLogger = prettyLogger.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false, // Color Configuration
		})
	return prettyLogger
}
