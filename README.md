# Image Analysis
(Also known as Image-Comp)

__V 0.0.1__
## Getting Started
This program will compress a file with the specific amount of compression (Currently slow).

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
-   [ ] Directory tree walking
-   [ ] Speed optimization
-   [ ] Async compression
-   [ ] Tests
-   [ ] Raw image handling
