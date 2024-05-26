# Compressing Algorithms on Golang
[![License](https://img.shields.io/github/license/evrone/go-clean-template.svg)](https://github.com/Shetami/compress/blob/main/LICENSE)
## Overview

This work was carried out as part of laboratory work in the course Algorithms and Data Structures. I did everything in Go, and since I am not a professional coder in this language, some of my writing decisions may be controversial. Basic compression algorithms are implemented here. At the moment, everything is implemented only within the framework of working with a text file. In the future I plan to add image compression and refine some existing algorithms for working with text. The implementation is not stable and may crash. which will be fixed later.

## Start Up

### Linux

Clone reposytory:

```
$ git clone https://github.com/Shetami/compress.git
```
Building the code into a binary file:
```
$ go build
```
If you don't have Go installed and configured, you can do this from the official [documentation](https://go.dev/doc/install).

Select a directory for your binary: Typically this will be `/usr/local/bin`, but you can choose any other directory that is in your `PATH` variable or create your own.

```
$ sudo cp [your_file_path] /usr/local/bin/[your_file]
```

Add your command execution to your `.bashrc` or `.bash_profile` file: These files are located in your home directory. If you are using another shell such as Zsh, use the appropriate file for it.

Next we set the global variable:

```bash
export PATH=$PATH:/usr/local/bin
```
Update changes: Run source `~/.bashrc` or `source ~/.bash_profile` to make the changes take effect in the current terminal session.

```
$ source ~/.bashrc
```

## Usage

```
$ compress [param] [file_path]
```
All available flags are called by the help command.
```
$ compress -h
```
