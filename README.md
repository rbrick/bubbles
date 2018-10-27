# bubbles
Ⓑⓤⓑⓑⓛⓔⓕⓨ ⓨⓞⓤⓡ ⓣⓔⓧⓣ

# Install
If you have Go installed you can do

```
go get -u github.com/rbrick/bubbles
```

If you do not have Go installed, you can check out the pre-built binaries under the "releases" tab.

# Usage

```sh
Usage of ./bubbles:
  -f string
        create a new file with the "bubble-fied" output
```

Note: By default this outputs to Stdout, the file is completely optional.

### Windows
Reads in an input file, and outputs it into `output.txt`
```sh
bubbles -f output.txt < input.txt
```

### *nix

Pipe your input from another command

```sh
echo "Hello, World" | ./bubbles -f output.txt
```

or just read in from a file
```sh
bubbles -f output.txt < input.txt
```

or pipe everything
```sh
echo dank memes bro | bubbles | cat
```
