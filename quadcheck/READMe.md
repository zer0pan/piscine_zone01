 ```
  __   _  _   __   ____     ___  _  _  ____  ___  __ _  ____  ____ 
 /  \ / )( \ / _\ (    \   / __)/ )( \(  __)/ __)(  / )(  __)(  _ \
(  O )) \/ (/    \ ) D (  ( (__ ) __ ( ) _)( (__  )  (  ) _)  )   /
 \__\)\____/\_/\_/(____/   \___)\_)(_/(____)\___)(__\_)(____)(__\_)
```

## Overview

This repository contains a Main.go program that compiles the executable program `quadchecker` and its subsidiary quad functions.

Quadchecker matches a given string to quad functions (QuadA, QuadB, QuadC, QuadD, QuadE). The program will display the name of the matching quad function(s) along with the dimensions of the input. If no match is found, it will print "Not a quad function".


## Task Allocation

- **Christos**: Core Logic Implementation / Code Finalization
- **Petros**: Executable Creation / Merging / Compatibility Tweaks
- **Taha**: Research / Documentation


## How it Works

1. **Executable Generation**: The program compiles individual executables for QuadA to QuadE, as well as a main executable called `quadchecker`.
2. **Input Analysis**: The `quadchecker` executable reads the input string and determines its dimensions and structure.
3. **Pattern Matching**: It compares the input pattern against outputs from QuadA to QuadE to identify matching quad functions.
4. **Result Display**: 
   - If one or more matches are found, the matching quad names and their dimensions are displayed.
   - Multiple matches are shown in alphabetical order, separated by `||`.
   - If no matches are found, the program outputs "Not a quad function".


## How to Run the Program

1. **Compile**: Run the `main.go` file to generate all executables (`quadA`, `quadB`, `quadC`, `quadD`, `quadE`, and `quadchecker`).
2. **Execute**: Use the terminal to test by piping the output from a quad function into `quadchecker`.

### Usage Example

**Checking a 3x3 QuadA:**

```bash
$ ./quadA 3 3 | ./quadchecker
[quadA] [3] [3]
