package healthz

import (
	"context"
	"fmt"
	"sync"

	"github.com/halimdotnet/tickboard-backend/internals/pkg/pgsql"
	rds "github.com/halimdotnet/tickboard-backend/internals/pkg/redis"
)

type checker struct {
	mu     sync.RWMutex
	checks map[string]CheckFunc
}

func NewChecker() Checker {
	return &checker{
		mu:     sync.RWMutex{},
		checks: make(map[string]CheckFunc),
	}
}

func (c *checker) Register(name string, fn CheckFunc) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.checks[name] = fn
}

func (c *checker) Check(ctx context.Context) map[string]CheckResult {
	c.mu.RLock()
	checks := make(map[string]CheckFunc, len(c.checks))
	for name, fn := range c.checks {
		checks[name] = fn
	}
	c.mu.RUnlock()

	results := make(map[string]CheckResult)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for name, fn := range checks {
		wg.Add(1)

		go func(name string, fn CheckFunc) {
			defer wg.Done()

			status := StatusOk
			errStr := ""

			if err := fn(ctx); err != nil {
				status = StatusError
				errStr = err.Error()
			}

			mu.Lock()
			results[name] = CheckResult{
				Status: status,
				Error:  errStr,
			}
			mu.Unlock()
		}(name, fn)
	}

	wg.Wait()
	return results
}

func RedisCheck(rd rds.Client) CheckFunc {
	return func(ctx context.Context) error {
		err := rd.Ping(ctx)
		if err != nil {
			return fmt.Errorf("health check redis: %w", err)
		}
		return nil
	}
}

func PostgresCheck(pg pgsql.Client) CheckFunc {
	return func(ctx context.Context) error {
		err := pg.Ping(ctx)
		if err != nil {
			return fmt.Errorf("health check postgres: %w", err)
		}
		return nil
	}
}
