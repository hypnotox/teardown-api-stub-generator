# Teardown Lua API v0.9.0 Stubs

Every release is a build of the current version at that date.

- Download the file you need, then copy it into your mod folder.
    - Lua
        - The IDE should pick up the functions by themselves.
        - If you install an extension with support for the annotations, types will be inferred as well.
        - Don't `#include` it!
    - Teal
        - Define `global_env_def = 'teardown'` in your `tlconfig.lua` file.