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
	graphs   map[*ir.Module]*SubGraph
	current  *SubGraph
}

type SubGraph struct {
	graph *cgraph.Graph
	nodes map[*ir.BasicBlock]*cgraph.Node
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
	return c.graphviz.RenderFilename(c.graph, graphviz.PNG, file)
}

func (c *Canvas) Draw() error {
	var err = c.create()
	if err == nil {
		for _, module := range c.program.Modules {
			err = c.drawModule(module)
		}
	}
	return err
}

func (c *Canvas) drawModule(module *ir.Module) error {
	var err error
	c.current, err = c.createSubGraph(module)
	if err == nil {
		for _, fun := range module.Functions {
			err = c.drawFunction(fun)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *Canvas) drawFunction(fun *ir.Func) error {
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

func (c *Canvas) createSubGraph(module *ir.Module) (*SubGraph, error) {
	var graph = c.graph.SubGraph(module.Name, 0)
	var err error
	var subGraph *SubGraph
	subGraph = &SubGraph{
		graph: graph,
		nodes: map[*ir.BasicBlock]*cgraph.Node{},
	}
	c.graphs[module] = subGraph
	return subGraph, err
}

func (c *Canvas) createNode(block *ir.BasicBlock) (*cgraph.Node, error) {
	var node, ok = c.current.nodes[block]
	if ok {
		return node, nil
	}
	node, err := c.graph.CreateNode(block.Name)
	if err == nil {
		node.SetShape("box")
		//var builder strings.Builder
		//builder.WriteString("<B>")
		//builder.WriteString(block.Name)
		//builder.WriteString("</B><BR/>")
		//for _, _ = range block.Instructions {
		//	builder.WriteString("hehe")
		//}
		//node.SetLabel(builder.String())
		c.current.nodes[block] = node
	}
	return node, err
}

func NewCanvas(program *ir.Program) *Canvas {
	return &Canvas{
		program: program,
		graphs:  map[*ir.Module]*SubGraph{},
	}
}
