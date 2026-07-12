package runtime

type Context struct {
	values map[string]any
}

func New() *Context {
	return &Context{
		values: make(map[string]any),
	}
}

func (c *Context) Set(key string, value any) {
	c.values[key] = value
}

func (c *Context) Get(key string) (any, bool) {
	v, ok := c.values[key]
	return v, ok
}
