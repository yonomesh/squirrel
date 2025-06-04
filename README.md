# squirrel
squirrel is log library for Golang project

zap API

```
logger := zap.NewExample()
defer logger.Sync()

sugar := logger.Sugar()
plain := sugar.Desugar()

logger, err := zap.NewProduction()
if err != nil {
  log.Fatalf("can't initialize zap logger: %v", err)
}
defer logger.Sync()
```

sq API
```
sq.NewLogEntry

logger.LazyRecord
logger.Record

log.New

logger.Trace
logger.Debug
logger.Info
logger.Warn
logger.Error
logger.Fatal
logger.Panic
```

log std

```
log.Fatal
log.Fatalf
log.Fataln

log.Panic
log.Panicf
log.Panicln

log.Print
log.Printf
log.Println

log.Prefix
log.SetPrefix

log.SetFlags
log.Flags

log.SetOutput
log.Output

log.Writer

type Logger struct {
	// contains filtered or unexported fields
}

```