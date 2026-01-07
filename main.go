package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/shopspring/decimal"

	"github.com/robojandro/loggenerator"
)

func main() {

	var verbose bool
	flag.BoolVar(&verbose, "verbose", false,
		"whether to output execution statistics to STDOUT after log statements have been output",
	)

	var delay int64
	flag.Int64Var(&delay, "delay", 10, "delay between log messages in milliseconds")

	var outputLimit int
	flag.IntVar(&outputLimit, "output", 1000,
		"default amount of log lines to output before exiting")

	// leverage default value to
	var fatalRatio, errorRatio, warnRatio, infoRatio, debugRatio, traceRatio int64
	flag.Int64Var(&fatalRatio, "fatal_ratio", -1, "expected rate of fatal level log messages")
	flag.Int64Var(&errorRatio, "error_ratio", -1, "expected rate of error level log messages")
	flag.Int64Var(&warnRatio, "warn_ratio", -1, "expected rate of warn level log messages")
	flag.Int64Var(&infoRatio, "info_ratio", -1, "expected rate of info level log messages")
	flag.Int64Var(&debugRatio, "debug_ratio", -1, "expected rate of debug level log messages")
	flag.Int64Var(&traceRatio, "trace_ratio", -1, "expected rate of trace level log messages")

	flag.Parse()

	specified := make(map[int64]bool, 6)

	// set up some sensible defaults
	levelRatios := []decimal.Decimal{
		decimal.NewFromInt(0),  // Fatal
		decimal.NewFromInt(10), // Error
		decimal.NewFromInt(20), // Warn
		decimal.NewFromInt(50), // Info
		decimal.NewFromInt(20), // Debug
		decimal.NewFromInt(0),  // Trace
	}

	// override defaults if any are specified to allow re-distribution to all
	// but fatal and trace
	if fatalRatio != -1 || errorRatio != -1 || warnRatio != -1 || infoRatio != -1 ||
		debugRatio != -1 || traceRatio != -1 {
		levelRatios[loggenerator.LvlFatal] = decimal.NewFromInt(0)
		levelRatios[loggenerator.LvlError] = decimal.NewFromInt(0)
		levelRatios[loggenerator.LvlWarn] = decimal.NewFromInt(0)
		levelRatios[loggenerator.LvlInfo] = decimal.NewFromInt(0)
		levelRatios[loggenerator.LvlDebug] = decimal.NewFromInt(0)
		levelRatios[loggenerator.LvlTrace] = decimal.NewFromInt(0)
	}

	if fatalRatio != -1 {
		specified[loggenerator.LvlFatal] = true
		levelRatios[loggenerator.LvlFatal] = decimal.NewFromInt(fatalRatio)
	}

	if errorRatio != -1 {
		specified[loggenerator.LvlError] = true
		levelRatios[loggenerator.LvlError] = decimal.NewFromInt(errorRatio)
	}

	if warnRatio != -1 {
		specified[loggenerator.LvlWarn] = true
		levelRatios[loggenerator.LvlWarn] = decimal.NewFromInt(warnRatio)
	}

	if infoRatio != -1 {
		specified[loggenerator.LvlInfo] = true
		levelRatios[loggenerator.LvlInfo] = decimal.NewFromInt(infoRatio)
	}

	if debugRatio != -1 {
		specified[loggenerator.LvlDebug] = true
		levelRatios[loggenerator.LvlDebug] = decimal.NewFromInt(debugRatio)
	}

	if traceRatio != -1 {
		specified[loggenerator.LvlTrace] = true
		levelRatios[loggenerator.LvlTrace] = decimal.NewFromInt(traceRatio)
	}

	generator, errs := loggenerator.New(specified, levelRatios)
	if len(errs) != 0 {
		fmt.Printf("Invalid ratios set: %+v\nErrors: %+v\n", levelRatios, errs)
		os.Exit(1)
	}

	ranges := generator.DeriveDistributionRanges()

	outputCounts := generator.Output(ranges, outputLimit, delay)

	if verbose {
		fmt.Printf(
			"level ratios\n Fatal  %s\n Error  %s\n Warn   %s\n Info   %s\n Debug  %s\n Trace  %s\n",
			levelRatios[loggenerator.LvlFatal].String(),
			levelRatios[loggenerator.LvlError].String(),
			levelRatios[loggenerator.LvlWarn].String(),
			levelRatios[loggenerator.LvlInfo].String(),
			levelRatios[loggenerator.LvlDebug].String(),
			levelRatios[loggenerator.LvlTrace].String(),
		)
		fmt.Printf(
			"random ranges\n Fatal  %d\n Error  %d\n Warn   %d\n Info   %d\n Debug  %d\n Trace  %d\n",
			ranges[loggenerator.LvlFatal],
			ranges[loggenerator.LvlError],
			ranges[loggenerator.LvlWarn],
			ranges[loggenerator.LvlInfo],
			ranges[loggenerator.LvlDebug],
			ranges[loggenerator.LvlTrace],
		)
		fmt.Printf("sum: %d\n", ranges[0]+ranges[1]+ranges[2]+ranges[3]+ranges[4]+ranges[5])

		fmt.Printf("output counts: %+v\n", outputCounts)
	}
}
