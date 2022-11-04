# Blitz: ECS-based 2D game engine

<div style="align-content: center">
    <img src="https://img.shields.io/badge/Made%20with-Go-1f425f.svg" alt="made with Go">
</div>

## Concept

Library based on sdl2 graphics library (and also go wrapper for sdl2 & sdl2 libs)

The goal of ECS pattern is data-logic separation. It improves code modularity and convenience

### Basis

Engine separated by categories:

- Entities
- Components
- Systems

#### Entities

Objects which contains id and a list of properties (components)

#### Components

Properties with data inside

#### Systems

Logics implementation that affect components each engine Update() event

Systems using global entities list to process each component in every single entity

Systems can interact with several components, for example:
local2world system using position, rotation and size to calculate child entity global position

## Build examples - Ubuntu

### Prebuild

```bash
# install dependencies
sudo apt install libglu1-mesa-dev freeglut3-dev mesa-common-dev libsdl2-dev libsdl2-*-dev
```

### Build and run example

```bash
cd examples/{example name}
go build -o {example name} main.go
./{example name}
```