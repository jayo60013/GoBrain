# gobrain

gobrain is a Brainfuck interpreter written in Go. Brainfuck is an esoteric programming language known for its minimalism and simplicity.

## Installation

To install gobrain, you need to have Go installed on your system. Then, you can simply clone this repository:

```bash
git clone https://github.com/jayo60013/GoBrain.git
```
Navigate to the cloned directory
```bash
cd gobrain
```
Then build the executable using the `go build` command
```bash
go build -o gobrain main.go
```

After building the executable, you can run the gobrain interpreter by executing this command. The following example runs hello world.
```bash
./gobrain examples/hello.bf
```
Replace `examples/hello.bf` to the path of your brainfuck program file.

### Examples
In the `examples` directory you will find some example Brainfuck programs taken from the [Brainfuck Wiki](https://en.wikipedia.org/wiki/Brainfuck) you can use to test GoBrain.

## License
This project is licensed under the MIT License.

