package engine

import (
	"context"
	"log"
	"regexp"
	"sync"

	"github.com/dlclark/regexp2"
	"github.com/insidersec/insider/report"
)

// Regex interface that abstracts both regexp and regexp2
type Regex interface {
	MatchString(s string) bool
}

// StdRegexp wraps Go's builtin regexp.Regexp
type StdRegexp struct {
	*regexp.Regexp
}

func (r *StdRegexp) MatchString(s string) bool {
	return r.Regexp.MatchString(s)
}

// Regex2 wraps regexp2.Regexp
type Regex2 struct {
	*regexp2.Regexp
}

func (r *Regex2) MatchString(s string) bool {
	ok, _ := r.Regexp.MatchString(s) // regexp2 returns (bool, error)
	return ok
}

// Engine uses the abstract Regex interface
type Engine struct {
	logger      *log.Logger
	exclude     []Regex
	ruleBuilder RuleBuilder
	jobs        int
}

// New creates a new Engine with any Regex implementation
func New(ruleBuilder RuleBuilder, exclude []Regex, jobs int, logger *log.Logger) *Engine {
	return &Engine{
		logger:      logger,
		exclude:     exclude,
		ruleBuilder: ruleBuilder,
		jobs:        jobs,
	}
}

func (e *Engine) Scan(ctx context.Context, dir string) (report.Result, error) {
	e.logger.Printf("Analysing files on directory %s\n", dir)
	scanner := &scanner{
		logger:      e.logger,
		mutext:      new(sync.Mutex),
		wg:          new(sync.WaitGroup),
		ch:          make(chan bool, e.jobs),
		errors:      make([]error, 0),
		ctx:         ctx,
		result:      new(Result),
		ruleBuilder: e.ruleBuilder,
		ruleSet:     NewRuleSet(),
		dir:         dir,
		exclude:     extractRegexp2(e.exclude), // TODO: Check for this so that regexp and regexp2 can coexist
	}
	return scanner.Process()
}

// extractRegexp2 converts []Regex to []*regexp2.Regexp by extracting from Regex2
func extractRegexp2(regexes []Regex) []*regexp2.Regexp {
	var result []*regexp2.Regexp
	for _, r := range regexes {
		if rx2, ok := r.(*Regex2); ok && rx2.Regexp != nil {
			result = append(result, rx2.Regexp)
		}
	}
	return result
}
