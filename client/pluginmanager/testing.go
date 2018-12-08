package pluginmanager

type MockPluginManager struct {
	RunF      func()
	ShutdownF func()
}

func (m *MockPluginManager) Run()               { m.RunF() }
func (m *MockPluginManager) Shutdown()          { m.ShutdownF() }
func (m *MockPluginManager) PluginType() string { return "mock" }
func (m *MockPluginManager) Ready() <-chan struct{} {
	ch := make(chan struct{})
	close(ch)
	return ch
}
