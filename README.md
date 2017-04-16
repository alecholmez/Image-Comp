# Image Comp
[![CircleCI](https://circleci.com/gh/alecholmez/Image-Comp/tree/develop.svg?style=shield)](https://circleci.com/gh/alecholmez/Image-Comp/tree/develop)
[![Coverage Status](https://coveralls.io/repos/github/alecholmez/Image-Comp/badge.svg?branch=develop)](https://coveralls.io/github/alecholmez/Image-Comp?branch=develop)

__V 0.0.1__
#### Current Image Formats Supported
-   PNG

## Getting Started
This program will compress a file with the specified amount of compression.

To run the program, follow these steps:
```bash
git clone git@github.com:alecholmez/Image-Comp.git
cd Image-Comp
go build
./Image-Comp -h
```
This will yield the following:
```bash
Usage of ./Image-Comp:
  -imgPath string
        Path to a directory (default "./imgs")
```
The program currently accepts one flag which is a path to the file to be compressed.

## Roadmap
-   [X] Directory tree walking
-   [X] Speed optimization
-   [X] Async compression
-   [X] Tests
-   [ ] Better memory management (Currently using high amounts of memory)
