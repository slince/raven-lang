package graphviz

import (
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
	"github.com/slince/php-plus/ir"
	"strconv"
)

type Canvas struct {
}

func (c *Canvas) Draw(program *ir.Program) error {
	var g = graphviz.New()
	graph, err := g.Graph()

	if err != nil {
		return err
	}
}

func (c *Canvas) drawFunction(graph *cgraph.Graph, fun *ir.Function) error {
	for idx, block := range fun.Blocks {
		var name = block.Name
		if len(name) == 0 {
			name = "L" + strconv.Itoa(idx)
		}
		node, err := graph.CreateNode(name)
		if err != nil {
			return err
		}
	}
}
