# dockmaster

## Build

Prerequisites:

- GO (1.22)
- C++ compiler

### Windows:

#### Build for Windows using WSL

Pre-requisites:

- package `gcc-mingw-w64-x86-64-win32` (`sudo apt install gcc-mingw-w64-x86-64`)


```
make build_wsl_to_windows
```

#### Build for Windows using minGW

- Tested with: https://github.com/skeeto/w64devkit

Unzip devkit to desired location, and add to PATH.

```
make build_mingw_to_windows
```