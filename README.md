# ECS with Raylib Renderer

[![Go Report Card](https://goreportcard.com/badge/github.com/andygeiss/ecs-example)](https://goreportcard.com/report/github.com/andygeiss/ecs-example)
[![BCH compliance](https://bettercodehub.com/edge/badge/andygeiss/ecs-example?branch=master)](https://bettercodehub.com/)

**A basic ECS benchmark using Raylib for rendering.**

[![Raylib](https://www.raylib.com/images/raylib_architecture_v4.2.png)](https://www.raylib.com/index.html)

## Start benchmark 

    go run main.go

Settings in [entrypoint.go](https://github.com/andygeiss/ecs-example/tree/main/internal/app/entrypoint.go):
- 200.000 Entities
- Resolution: 1366 x 768

Result on my `i7-10510U with Intel GPU` ;-)

![Example](example.png)
