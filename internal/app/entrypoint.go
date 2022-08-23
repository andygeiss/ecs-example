package app

// Entrypoint ...
func Entrypoint() {
	benchmark := NewApp().(*App).
		WithNumEntities(200000).
		WithWidth(1366).
		WithHeight(768).
		WithTitle("ECS benchmark with Raylib")
	benchmark.Setup()
	defer benchmark.Teardown()
	benchmark.Run()
}
