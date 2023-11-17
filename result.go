package benchmank

import (
	"fmt"
	"math"
	"testing"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Result struct {
	testing.BenchmarkResult
	TimePerOp             float64
	TimePerOpBase         string
	AllocedBytesPerOp     float64
	AllocedBytesPerOpBase string
}

func String(br testing.BenchmarkResult) string {
	return NewBenchmankResult(br).String()
}

func NewBenchmankResult(bResult testing.BenchmarkResult) Result {
	timePerOp, timePerOpBase := humanReadTime(uint64(bResult.NsPerOp()))
	allocedBytesPerOp, allocedBytesPerOpBase := humanReadBytes(uint64(bResult.AllocedBytesPerOp()))
	hbr := Result{
		BenchmarkResult:       bResult,
		TimePerOp:             timePerOp,
		TimePerOpBase:         timePerOpBase,
		AllocedBytesPerOp:     allocedBytesPerOp,
		AllocedBytesPerOpBase: allocedBytesPerOpBase,
	}

	return hbr
}

func (hbr Result) String() string {
	msg := message.NewPrinter(language.English)
	return fmt.Sprintf("n:%8s \t %12s/op\t %12s/op\t %12s allocs/op",
		msg.Sprintf("%d", hbr.N),                                     // The number of iterations.
		baseString(hbr.TimePerOp, hbr.TimePerOpBase),                 // TimePerOp
		baseString(hbr.AllocedBytesPerOp, hbr.AllocedBytesPerOpBase), // AllocedBytesPerOp
		msg.Sprintf("%d", hbr.AllocsPerOp()),                         // AllocsPerOp
	)
}

func baseString(val float64, base string) string {
	return fmt.Sprintf("%.3f %s", val, base)
}
func humanReadBytes(s uint64) (float64, string) {
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	bases := []float64{1000, 1000, 1000, 1000, 1000, 1000}

	return humanRead(s, bases, sizes)
}
func humanReadTime(s uint64) (float64, string) {
	sizes := []string{"ns", "µs", "ms", " s", " m", " h", " day"}
	bases := []float64{1000, 1000, 1000, 60, 60, 24}

	return humanRead(s, bases, sizes)
}

func humanRead(s uint64, bases []float64, sizes []string) (float64, string) {
	if float64(s) < bases[0] {
		return float64(s), sizes[0]
	}

	var i = 0
	var base float64 = 1
	for e := float64(s); e > bases[i]; i++ {
		e = math.Floor(e / bases[i])
		base *= bases[i]
	}

	return float64(s) / base, sizes[i]
}
