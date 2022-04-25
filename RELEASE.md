# Teardown Lua API v1.0.0 Stubs

## Changes
- Extend lua stubs with overloaded functions to prevent false positive missing arguments

## Usage
- Download the file you need, then copy it into your mod folder.
    - Lua
        - The IDE should pick up the functions by themselves.
        - If you install an extension with support for the annotations, types will be inferred as well.
        - Don't `#include` it!
    - Teal
        - Define `global_env_def = 'teardown'` in your `tlconfig.lua` file.
