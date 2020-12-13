package plugin

var (
	registry map[string]*Plugin
)

func init() {
	registry = make(map[string]*Plugin, 0)
}

func Register(name string, setup func(ctx *Context) error) error {
	registry[name] = &Plugin{
		Name: name,
		Setup: setup,
	}

	return nil
}

func RegisterWithTearDown(name string, setup func(ctx *Context) error, teardown func(ctx *Context) error) error {
	registry[name] = &Plugin{
		Name: name,
		Setup: setup,
		TearDown: teardown,
	}

	return nil
}

func Get(name string) *Plugin {
	return registry[name]
}

