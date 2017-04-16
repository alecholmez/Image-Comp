# Image Comp
[![CircleCI](https://circleci.com/gh/alecholmez/Image-Comp/tree/develop.svg?style=svg)](https://circleci.com/gh/alecholmez/Image-Comp/tree/develop)

__V 0.0.1__
#### Current Image Formats Supported
-   JPEG
-   PNG
-   GIF

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
Usage of ./image-analysis:
  -imgPath string
        Path to an image (default "./DarkBB8.png")
```
The program currently accepts one flag which is a path to the file to be compressed.

## Roadmap
-   [X] Directory tree walking
-   [X] Speed optimization
-   [X] Async compression
-   [ ] Tests
-   [ ] Raw image handling
-   [ ] Better memory management (Currently using high amounts of memory)
