package insider

import (
	"context"

	"github.com/PLZENTERTEXT/insider/report"
)

type CsharpAnalyzer struct {
}

func NewCsharpAnalyzer() CsharpAnalyzer {
	return CsharpAnalyzer{}
}

func (a CsharpAnalyzer) Analyze(ctx context.Context, dir string) (report.Reporter, error) {
	// C# has no special analysis
	return report.Report{}, nil
}
