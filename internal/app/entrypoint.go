package app

// Entrypoint ...
func Entrypoint() {
	benchmark := NewApp(200000, 1366, 768, "ECS benchmark with Raylib")
	benchmark.Setup()
	defer benchmark.Teardown()
	benchmark.Run()
}
