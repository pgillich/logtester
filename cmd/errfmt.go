package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/pgillich/errfmt"
	"github.com/pgillich/logtester/errfmt_tester"
)

const (
	// OptFormatter is the formatter name
	OptFormatter = "formatter"

	// OptTestCase is the test case
	OptTestCase = "testCase"

	// OptFlagExtractDetails extracts errors.Details to logrus.Fields
	OptFlagExtractDetails = "extractDetails"
	// OptFlagCallStackInFields extracts errors.StackTrace() to logrus.Fields
	OptFlagCallStackInFields = "callStackInFields"
	// OptFlagCallStackOnConsole extracts errors.StackTrace() to logrus.Field "callstack"
	OptFlagCallStackOnConsole = "callStackOnConsole"
	// OptFlagCallStackInHTTPProblem extracts errors.StackTrace() to HTTPProblem
	OptFlagCallStackInHTTPProblem = "callStackInHTTPProblem"
	// OptFlagPrintStructFieldNames renders non-scalar Details values are by "%+v", instead of "%v"
	OptFlagPrintStructFieldNames = "printStructFieldNames"
	// OptFlagTrimJSONDquote trims the leading and trailing '"' of JSON-formatted values
	OptFlagTrimJSONDquote = "trimJSONDquote"

	// OptCallStackSkipLast skips the last n items of call stack
	OptCallStackSkipLast = "callStackSkipLast"
)

var errorformatterCmd = &cobra.Command{
	Use:   "errfmt",
	Short: "github.com/pgillich/errfmt",
	Long: `Start service.
Example commands:

./logtester errfmt --formatter syslog --extractDetails --callStackOnConsole true --callStackInHTTPProblem
`,
	Run: func(cmd *cobra.Command, args []string) {
		testErrorformatter()
	},
}

func init() { // nolint:gochecknoinits
	RootCmd.AddCommand(errorformatterCmd)

	errorformatterCmd.PersistentFlags().String(OptFormatter, "text", "Formatter: text, syslog, json")
	// nolint:gosec,errcheck
	viper.BindPFlag(OptFormatter, errorformatterCmd.PersistentFlags().Lookup(OptFormatter))

	// nolint:lll
	errorformatterCmd.PersistentFlags().String(OptTestCase, "errorhttp", "Test Case: sampletext, samplehttp, info, error, errorhttp")
	// nolint:gosec,errcheck
	viper.BindPFlag(OptTestCase, errorformatterCmd.PersistentFlags().Lookup(OptTestCase))

	errorformatterCmd.PersistentFlags().Bool(OptFlagExtractDetails, false, "Extracts errors.Details to logrus.Fields")
	// nolint:gosec,errcheck
	viper.BindPFlag(OptFlagExtractDetails, errorformatterCmd.PersistentFlags().Lookup(OptFlagExtractDetails))

	// nolint:lll
	errorformatterCmd.PersistentFlags().Bool(OptFlagCallStackInFields, false, "Extracts errors.StackTrace() to logrus.Fields")
	// nolint:gosec,errcheck
	viper.BindPFlag(OptFlagCallStackInFields, errorformatterCmd.PersistentFlags().Lookup(OptFlagCallStackInFields))

	// nolint:lll
	errorformatterCmd.PersistentFlags().Bool(OptFlagCallStackOnConsole, false, `Extracts errors.StackTrace() to logrus.Field "callstack"`)
	// nolint:gosec,errcheck
	viper.BindPFlag(OptFlagCallStackOnConsole, errorformatterCmd.PersistentFlags().Lookup(OptFlagCallStackOnConsole))

	// nolint:lll
	errorformatterCmd.PersistentFlags().Bool(OptFlagCallStackInHTTPProblem, false, "Extracts errors.StackTrace() to HTTPProblem")
	// nolint:gosec,errcheck,lll
	viper.BindPFlag(OptFlagCallStackInHTTPProblem, errorformatterCmd.PersistentFlags().Lookup(OptFlagCallStackInHTTPProblem))

	// nolint:lll
	errorformatterCmd.PersistentFlags().Bool(OptFlagPrintStructFieldNames, false, `Renders non-scalar Details values are by "%+v", instead of "%v"`)
	// nolint:gosec,errcheck
	viper.BindPFlag(OptFlagPrintStructFieldNames, errorformatterCmd.PersistentFlags().Lookup(OptFlagPrintStructFieldNames))

	// nolint:lll
	errorformatterCmd.PersistentFlags().Bool(OptFlagTrimJSONDquote, false, `Trims the leading and trailing '"' of JSON-formatted values`)
	// nolint:gosec,errcheck
	viper.BindPFlag(OptFlagTrimJSONDquote, errorformatterCmd.PersistentFlags().Lookup(OptFlagTrimJSONDquote))

	// nolint:lll
	errorformatterCmd.PersistentFlags().Int(OptCallStackSkipLast, 0, "Skips the last n items of call stack")
	// nolint:gosec,errcheck
	viper.BindPFlag(OptCallStackSkipLast, errorformatterCmd.PersistentFlags().Lookup(OptCallStackSkipLast))
}

func testErrorformatter() {
	flags := errfmt.FlagNone
	if viper.GetBool(OptFlagExtractDetails) {
		flags += errfmt.FlagExtractDetails
	}
	if viper.GetBool(OptFlagCallStackInFields) {
		flags += errfmt.FlagCallStackInFields
	}
	if viper.GetBool(OptFlagCallStackOnConsole) {
		flags += errfmt.FlagCallStackOnConsole
	}
	if viper.GetBool(OptFlagCallStackInHTTPProblem) {
		flags += errfmt.FlagCallStackInHTTPProblem
	}
	if viper.GetBool(OptFlagPrintStructFieldNames) {
		flags += errfmt.FlagPrintStructFieldNames
	}
	if viper.GetBool(OptFlagTrimJSONDquote) {
		flags += errfmt.FlagTrimJSONDquote
	}

	errfmt_tester.TryErrorformatter(
		viper.GetString(OptFormatter),
		viper.GetString(OptTestCase),
		flags,
		viper.GetInt(OptCallStackSkipLast),
	)
}
