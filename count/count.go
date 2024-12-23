package count

import (
	"bufio"
	"io"
	"os"
	"fmt"
	"errors"
	"flag"
)

type counter struct {
	input io.Reader
	output io.Writer
}

type option func(*counter) error

func WithInput(input io.Reader) option {
	return func(c *counter) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		c.input = input
		return nil
	}
}

func WithInputFromArgs(args []string) option {
	return func (c *counter) error {
		if len(args) < 1 {
			return nil
		}
		f, err := os.Open(args[0])
		if err != nil {
			return err
		}
		c.input = f
		return nil
	}
}

func WithOutput(output io.Writer) option {
	return func(c *counter) error {
		if output == nil {
			return errors.New("nil output writer")
		}
		c.output = output
		return nil
	}
}

func NewCounter(opts ...option) (*counter, error) {
	c := &counter{
		input: os.Stdin,
		output: os.Stdout,
	}
	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c *counter) Lines() int {
	lines := 0
	input := bufio.NewScanner(c.input)
	for input.Scan() {
		lines++
	}
	return lines
}


func (c *counter) Words() int {
	words := 0
	input := bufio.NewScanner(c.input)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		words++
	}
	return words
}

func Main() int {
	lineMode := flag.Bool("lines", false, "Count lines, not words")
	flag.Parse()
	c, err := NewCounter(
		WithInputFromArgs(flag.Args()),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	if *lineMode {
		fmt.Println(c.Lines())
	} else {
		fmt.Println(c.Words())
	}
	return 0
}

func MainLines() int {
	c, err := NewCounter(
		WithInputFromArgs(os.Args[1:]),	
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	fmt.Println(c.Lines())
	return 0
}


func MainWords() int {
	c, err := NewCounter(
		WithInputFromArgs(os.Args[1:]),	
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	fmt.Println(c.Words())
	return 0
}
