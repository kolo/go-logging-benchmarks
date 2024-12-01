package bench

import (
	"io"

	"github.com/remerge/cue"
	"github.com/remerge/cue/collector"
)

type cueBench struct {
	l cue.Logger

}

func newCue(w io.Writer) cue.Logger {
	collector := collector.Terminal{
		Writer: w,
	}
	cue.Collect(cue.INFO, collector.New())
	return cue.NewLogger("cue")
}

func (b *cueBench) new(w io.Writer) logBenchmark {
	return &cueBench{
		l: newCue(w),
	}
}

func (b *cueBench) newWithCtx(w io.Writer) logBenchmark {
	l := newCue(w)
	return &cueBench{
		l: l.WithFields(mapFields()),
	}
}

func (b *cueBench) name() string {
	return "cue"
}

func (b *cueBench) logEvent(msg string) {
	b.l.Info(msg)
}

func (b *cueBench) logEventFmt(msg string, args ...any) {
	b.l.Infof(msg, args)
}

func (b *cueBench) logEventCtx(msg string) {
	b.l.WithFields(mapFields()).Info(msg)
}

func (b *cueBench) logEventCtxWeak(msg string) {
	b.logEventCtx(msg)
}

func (b *cueBench) logDisabled(msg string) {
	b.l.Debug(msg)
}

func (b *cueBench) logDisabledFmt(msg string, args ...any) {
	b.l.Debugf(msg, args)
}

func (b *cueBench) logDisabledCtx(msg string) {
	b.l.WithFields(mapFields()).Debug(msg)
}

func (b *cueBench) logDisabledCtxWeak(msg string) {
	b.logDisabledCtx(msg)
}
