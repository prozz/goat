package ai

import (
	"bufio"
	"github.com/mb-14/gomarkov"
	"io"
	"strings"
)

type Generator interface {
	RandomTitle() string
}

type generator struct {
	chain *gomarkov.Chain
}

func NewGenerator(reader io.Reader) Generator {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	chain := gomarkov.NewChain(1)
	for scanner.Scan() {
		chain.Add(strings.Split(scanner.Text(), " "))
	}

	return &generator{
		chain: chain,
	}
}

func (g *generator) RandomTitle() string {
	tokens := []string{gomarkov.StartToken}
	for tokens[len(tokens)-1] != gomarkov.EndToken {
		next, _ := g.chain.Generate(tokens[(len(tokens) - 1):])
		tokens = append(tokens, next)
	}
	return strings.Join(tokens[1:len(tokens)-1], " ")
}