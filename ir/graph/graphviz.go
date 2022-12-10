package graph

import (
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
	"github.com/slince/php-plus/ir"
	"log"
)

type Canvas struct {
	program  *ir.Program
	graphviz *graphviz.Graphviz
	graph    *cgraph.Graph
	nodes    map[*ir.BasicBlock]*cgraph.Node
}

func (c *Canvas) Draw() error {
	var err = c.create()
	if err != nil {
		return err
	}
	for _, fun := range c.program.Modules[0].Functions {
		err = c.drawFunction(fun)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Canvas) create() error {
	var err error
	c.graphviz = graphviz.New()
	c.graph, err = c.graphviz.Graph()
	return err
}

func (c *Canvas) SaveTo(file string) error {
	defer func() {
		if err := c.graph.Close(); err != nil {
			log.Fatal(err)
		}
		_ = c.graphviz.Close()
	}()
	// 3. write to file directly
	return c.graphviz.RenderFilename(c.graph, graphviz.PNG, file)
}

func (c *Canvas) drawFunction(fun *ir.Function) error {
	for _, block := range fun.Blocks {
		var err = c.drawBlock(block)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Canvas) drawBlock(block *ir.BasicBlock) error {
	var node, err = c.createNode(block)
	if err != nil {
		return err
	}
	if block.Terminator != nil {
		if inst, ok := block.Terminator.(*ir.Jmp); ok {
			var target *cgraph.Node
			target, err = c.createNode(inst.Target.(*ir.BasicBlock))
			if err == nil {
				_, err = c.graph.CreateEdge("e", node, target)
			}
		} else if inst, ok := block.Terminator.(*ir.CondJmp); ok {
			var trueTarget *cgraph.Node
			trueTarget, err = c.createNode(inst.TrueTarget.(*ir.BasicBlock))
			if err == nil {
				_, err = c.graph.CreateEdge("e", node, trueTarget)
			}
			var falseTarget *cgraph.Node
			falseTarget, err = c.createNode(inst.FalseTarget.(*ir.BasicBlock))
			if err == nil {
				_, err = c.graph.CreateEdge("e", node, falseTarget)
			}
		}
	}
	return err
}

func (c *Canvas) createNode(block *ir.BasicBlock) (*cgraph.Node, error) {
	var node, ok = c.nodes[block]
	if ok {
		return node, nil
	}
	node, err := c.graph.CreateNode(block.Name)
	if err == nil {
		node.SetShape("box")
		node.SetLabel(block.Name)
		c.nodes[block] = node
	}
	return node, err
}

func NewCanvas(program *ir.Program) *Canvas {
	return &Canvas{
		program: program,
		nodes:   map[*ir.BasicBlock]*cgraph.Node{},
	}
}
