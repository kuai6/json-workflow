package jsonworkflow

import (
	"fmt"
	"sync"
)

// ContextInterface describes Context bucket
type ContextInterface interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{})
}

// Context general context implementation
type Context struct {
	*sync.RWMutex
	bucket map[string]interface{}
}

// Get function implements getter value from context
func (c *Context) Get(key string) (interface{}, error) {
	if val, ok := c.bucket[key]; ok {
		return val, nil
	}
	return nil, fmt.Errorf(fmt.Sprintf("key %s not found in context", key))
}

// Set function implements setter value with key to context
func (c *Context) Set(key string, value interface{}) {
	c.RWMutex.Lock()
	defer c.RWMutex.Unlock()
	c.bucket[key] = value
}
