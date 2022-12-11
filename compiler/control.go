package compiler

import (
	"github.com/samber/lo"
	"github.com/slince/php-plus/compiler/ast"
	"github.com/slince/php-plus/ir"
	"github.com/slince/php-plus/ir/value"
	"strconv"
)

func (c *Compiler) compileIfStmt(node *ast.IfStmt) (*ir.BasicBlock, error) {
	c.ctx.LeaveBlock = c.function.NewBlock("if.done")
	var consequent, err = c.compileBlockStmt(node.Consequent, "if.then")
	if err != nil {
		return nil, err
	}
	// Compile if else body
	var ifElse ir.Block = c.ctx.LeaveBlock
	if alternate, ok := node.Alternate.(*ast.BlockStmt); ok {
		ifElse, err = c.compileBlockStmt(alternate, "if.else")
	} else if alternate, ok := node.Alternate.(*ast.IfStmt); ok {
		ifElse, err = c.compileIfStmt(alternate)
	}
	if err != nil {
		return nil, err
	}
	// Compile if head
	var test = c.function.NewBlock("if.test")
	err = c.compileBlock(test, func() error {
		var test, err = c.compileExpr(node.Test)
		if err == nil {
			c.ctx.NewCondJmp(test, consequent, ifElse)
		}
		return err
	})
	c.ctx.NewJmp(test)
	return test, err
}

func (c *Compiler) compileSwitchStmt(node *ast.SwitchStmt) error {
	// compile switch cases
	var disc, err = c.compileExpr(node.Discriminant)
	if err != nil {
		return err
	}
	// jmp first case discriminant
	c.ctx.NewJmp(ir.NewReference("switch.case.disc.0"))
	c.ctx.LeaveBlock = c.function.NewBlock("switch.done")
	c.enterScope()
	var caseNum = len(node.Cases)
	var _, defaultIdx, _ = lo.FindIndexOf(node.Cases, func(clause *ast.SwitchCase) bool {
		return clause.Default
	})
	for idx, clause := range node.Cases {
		var caseBody, err = c.compileSwitchCaseBody(clause, idx, idx == caseNum-1)
		if err == nil {
			_, err = c.compileSwitchCaseDisc(disc, caseBody, clause, idx, idx == caseNum-1, defaultIdx)
		}
		if err != nil {
			return err
		}
	}
	c.leaveScope()
	return nil
}

func (c *Compiler) compileSwitchCaseDisc(disc value.Value, caseBody *ir.BasicBlock, node *ast.SwitchCase, idx int, last bool, defaultIdx int) (*ir.BasicBlock, error) {
	var discBlock = c.function.NewBlock("switch.case.disc." + strconv.Itoa(idx))
	var err = c.compileBlock(discBlock, func() error {
		if node.Default {
			c.ctx.NewJmp(caseBody)
			return nil
		}
		var test, err = c.compileExpr(node.Test)
		if err != nil {
			return err
		}
		var leaveTarget ir.Block
		if last {
			// jump to default case when not match the case, if the default case is present.
			if defaultIdx > -1 {
				leaveTarget = ir.NewReference("switch.case.disc." + strconv.Itoa(defaultIdx))
			} else {
				leaveTarget = c.ctx.LeaveBlock
			}
		} else {
			// Skip default branch
			if defaultIdx == idx+1 {
				leaveTarget = ir.NewReference("switch.case.disc." + strconv.Itoa(idx+1))
			} else {
				leaveTarget = ir.NewReference("switch.case.disc." + strconv.Itoa(idx+2))
			}
		}
		c.ctx.NewCondJmp(c.ctx.NewEq(disc, test), caseBody, leaveTarget)
		return nil
	})
	return discBlock, err
}

func (c *Compiler) compileSwitchCaseBody(node *ast.SwitchCase, idx int, last bool) (*ir.BasicBlock, error) {
	var caseBlock = c.function.NewBlock("switch.case." + strconv.Itoa(idx))
	var err = c.compileBlock(caseBlock, func() error {
		var err = c.compileSwitchCaseConsequent(node)
		if err != nil {
			return err
		}
		if c.ctx.Terminator == nil {
			var leaveTarget ir.Block
			if last {
				leaveTarget = c.ctx.LeaveBlock
			} else {
				leaveTarget = ir.NewReference("switch.case." + strconv.Itoa(idx+1))
			}
			c.ctx.NewJmp(leaveTarget)
		}
		return nil
	})
	return caseBlock, err
}

func (c *Compiler) compileSwitchCaseConsequent(node *ast.SwitchCase) error {
	for _, consequent := range node.Consequent {
		err := c.compileStmt(consequent)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Compiler) compileDoWhileStmt(node *ast.DoWhileStmt) error {
	c.ctx.LeaveBlock = c.function.NewBlock("do.while.done")
	var body, err = c.compileBlockStmt(node.Body, "do.while.body")
	if err != nil {
		return err
	}
	var test = c.function.NewBlock("do.while.test")
	err = c.compileBlock(test, func() error {
		var test, err = c.compileExpr(node.Test)
		if err == nil {
			c.ctx.NewCondJmp(test, body, c.ctx.LeaveBlock)
		}
		return err
	})
	if err != nil {
		return err
	}
	if body.Terminator == nil {
		body.NewJmp(test)
	}
	c.ctx.NewJmp(body)
	return nil
}

func (c *Compiler) compileWhileStmt(node *ast.WhileStmt) error {
	c.ctx.LeaveBlock = c.function.NewBlock("while.done")
	var body, err = c.compileBlockStmt(node.Body, "while.body")
	if err != nil {
		return err
	}
	var test = c.function.NewBlock("while.test")
	err = c.compileBlock(test, func() error {
		var test, err = c.compileExpr(node.Test)
		if err == nil {
			c.ctx.NewCondJmp(test, body, c.ctx.LeaveBlock)
		}
		return err
	})
	if body.Terminator == nil {
		body.NewJmp(test)
	}
	c.ctx.NewJmp(test)
	return nil
}

func (c *Compiler) compileBreakStmt(node *ast.BreakStmt) {
	c.ctx.NewJmp(c.ctx.LeaveBlock)
}
