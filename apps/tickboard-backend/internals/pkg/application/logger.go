package application

import (
	"strings"

	"github.com/halimdotnet/tickboard-backend/internals/pkg/logger"
	"go.uber.org/fx/fxevent"
)

type fxLogger struct {
	log logger.Logger
}

func (l *fxLogger) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.OnStartExecuting:
		l.log.Debugw("OnStart hook executing", logger.String("callee", e.FunctionName), logger.String("caller", e.CallerName))
	case *fxevent.OnStartExecuted:
		if e.Err != nil {
			l.log.Errorw("OnStart hook failed", logger.String("callee", e.FunctionName), logger.String("caller", e.CallerName), logger.Error(e.Err))
		} else {
			l.log.Debugw("OnStart hook executed", logger.String("callee", e.FunctionName), logger.String("caller", e.CallerName), logger.String("runtime", e.Runtime.String()))
		}
	case *fxevent.OnStopExecuting:
		l.log.Debugw("OnStop hook executing", logger.String("callee", e.FunctionName), logger.String("caller", e.CallerName))
	case *fxevent.OnStopExecuted:
		if e.Err != nil {
			l.log.Errorw("OnStop hook failed", logger.String("callee", e.FunctionName), logger.String("caller", e.CallerName), logger.Error(e.Err))
		} else {
			l.log.Debugw("OnStop hook executed", logger.String("callee", e.FunctionName), logger.String("caller", e.CallerName), logger.String("runtime", e.Runtime.String()))
		}
	case *fxevent.Supplied:
		if e.Err != nil {
			l.log.Errorw("supplied failed", logger.String("type", e.TypeName), logger.Error(e.Err))
		} else {
			l.log.Debugw("supplied", logger.String("type", e.TypeName))
		}
	case *fxevent.Provided:
		for _, rtype := range e.OutputTypeNames {
			l.log.Debugw("provided", logger.String("constructor", e.ConstructorName), logger.String("type", rtype))
		}
		if e.Err != nil {
			l.log.Errorw("error encountered while applying options", logger.Error(e.Err))
		}
	case *fxevent.Decorated:
		for _, rtype := range e.OutputTypeNames {
			l.log.Debugw("decorated", logger.String("decorator", e.DecoratorName), logger.String("type", rtype))
		}
		if e.Err != nil {
			l.log.Errorw("error encountered while applying options", logger.Error(e.Err))
		}
	case *fxevent.Invoking:
		l.log.Debugw("invoking", logger.String("function", e.FunctionName))
	case *fxevent.Invoked:
		if e.Err != nil {
			l.log.Errorw("invoke failed", logger.Error(e.Err), logger.String("stack", e.Trace), logger.String("function", e.FunctionName))
		}
	case *fxevent.Stopping:
		l.log.Infow("received signal", logger.String("signal", strings.ToUpper(e.Signal.String())))
	case *fxevent.Stopped:
		if e.Err != nil {
			l.log.Errorw("stop failed", logger.Error(e.Err))
		}
	case *fxevent.RollingBack:
		l.log.Errorw("starting rollback", logger.Error(e.StartErr))
	case *fxevent.RolledBack:
		if e.Err != nil {
			l.log.Errorw("rollback failed", logger.Error(e.Err))
		}
	case *fxevent.Started:
		if e.Err != nil {
			l.log.Errorw("start failed", logger.Error(e.Err))
		} else {
			l.log.Infow("started")
		}
	case *fxevent.LoggerInitialized:
		if e.Err != nil {
			l.log.Errorw("custom logger initialization failed", logger.Error(e.Err))
		} else {
			l.log.Debugw("initialized custom fxevent.Logger", logger.String("function", e.ConstructorName))
		}
	}
}
