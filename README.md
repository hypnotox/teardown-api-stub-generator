# Teardown API Stub Generator

This project aims to provide a parser for the XML files at https://www.teardowngame.com/modding/api.xml

These files describe the available Lua API for mods.

The file is currently being modified beforehand, since there are some types that don't exist in Lua and `Vector`, `Quaternion` as well as `Transform` are just described as `table`.

The parser generates a stub file for [Lua](https://www.lua.org/).

## Usage

#### If you just want the stub file

- Go to the [releases section](https://github.com/hypnotox/teardown-api-stub-generator/releases).
- There you can download the file you need, then copy it into your mod folder.
  - Lua
    - The IDE should pick up the functions by themselves.
    - If you install an extension with support for the annotations, types will be inferred as well.
    - Don't `#include` it!

#### If you want to, for some reason, run the generator yourself

You have to have [Go](https://go.dev/) installed and in your PATH.

Download the source and run `go run` in the directory to build the stubs from the current `api.xml` file.

## Useful IDE Extensions

### Lua
These extensions extend type inference for Lua.
- Visual Studio Code: https://marketplace.visualstudio.com/items?itemName=sumneko.lua
- JetBrains IDEs: https://plugins.jetbrains.com/plugin/9768-emmylua
