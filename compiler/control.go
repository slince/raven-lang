package compiler

import (
	"github.com/samber/lo"
	"github.com/slince/php-plus/compiler/ast"
	"github.com/slince/php-plus/ir"
	"strconv"
)

func (c *Compiler) compileIfStmt(node *ast.IfStmt) *ir.BasicBlock {
	c.ctx.LeaveBlock = c.Function.NewBlock("if.done")
	var consequent = c.compileBlockStmt(node.Consequent, "if.then")
	if consequent.Terminator == nil {
		consequent.NewJmp(c.ctx.LeaveBlock)
	}
	var ifElse ir.Block = c.ctx.LeaveBlock
	if node.Alternate != nil {
		if alternate, ok := node.Alternate.(*ast.BlockStmt); ok {
			ifElse = c.compileBlockStmt(alternate, "if.else")
		} else if alternate, ok := node.Alternate.(*ast.IfStmt); ok {
			ifElse = c.compileIfStmt(alternate)
		}
	}
	var test = c.Function.NewBlock("if.test")
	c.compileBlock(test, func() {
		c.ctx.NewCondJmp(c.compileExpr(node.Test), consequent, ifElse)
	})
	return test
}

func (c *Compiler) compileSwitchStmt(node *ast.SwitchStmt) {
	// compile switch cases
	var disc = c.compileExpr(node.Discriminant)
	c.ctx.LeaveBlock = c.Function.NewBlock("switch.done")
	c.enterScope()
	var caseNum = len(node.Cases)
	var _, defaultIdx, _ = lo.FindIndexOf(node.Cases, func(clause *ast.SwitchCase) bool {
		return clause.Default
	})
	for idx, clause := range node.Cases {
		var caseBody = c.compileSwitchCaseBody(clause, idx, idx == caseNum-1)
		c.compileSwitchCaseDisc(disc, caseBody, clause, idx, idx == caseNum-1, defaultIdx)
	}
	c.leaveScope()
	// jmp first case discriminant
	c.ctx.NewJmp(ir.NewReference("switch.case.disc.0"))
}

func (c *Compiler) compileSwitchCaseDisc(disc ir.Operand, caseBody *ir.BasicBlock, node *ast.SwitchCase, idx int, last bool, defaultIdx int) *ir.BasicBlock {
	var discBlock = c.Function.NewBlock("switch.case.disc." + strconv.Itoa(idx))
	c.compileBlock(discBlock, func() {
		if node.Default {
			c.ctx.NewJmp(caseBody)
			return
		}
		var cond = ir.NewTemporary(nil)
		c.ctx.NewEq(cond, disc, c.compileExpr(node.Test))
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
		c.ctx.NewCondJmp(cond, caseBody, leaveTarget)
	})
	return discBlock
}

func (c *Compiler) compileSwitchCaseBody(node *ast.SwitchCase, idx int, last bool) *ir.BasicBlock {
	var caseBlock = c.Function.NewBlock("switch.case." + strconv.Itoa(idx))
	c.compileBlock(caseBlock, func() {
		c.compileSwitchCaseConsequent(node)
		if c.ctx.Terminator == nil {
			var leaveTarget ir.Block
			if last {
				leaveTarget = c.ctx.LeaveBlock
			} else {
				leaveTarget = ir.NewReference("switch.case." + strconv.Itoa(idx+1))
			}
			c.ctx.NewJmp(leaveTarget)
		}
	})
	return caseBlock
}

func (c *Compiler) compileSwitchCaseConsequent(node *ast.SwitchCase) {
	for _, consequent := range node.Consequent {
		c.compileStmt(consequent)
	}
}

func (c *Compiler) compileDoWhileStmt(node *ast.DoWhileStmt) {
	c.ctx.LeaveBlock = c.Function.NewBlock("do.while.done")
	var body = c.compileBlockStmt(node.Body, "do.while.body")
	var test = c.Function.NewBlock("do.while.test")
	c.compileBlock(test, func() {
		c.ctx.NewCondJmp(c.compileExpr(node.Test), body, c.ctx.LeaveBlock)
	})
	if body.Terminator == nil {
		body.NewJmp(test)
	}
	c.ctx.NewJmp(body)
}

func (c *Compiler) compileWhileStmt(node *ast.WhileStmt) {
	c.ctx.LeaveBlock = c.Function.NewBlock("while.done")
	var body = c.compileBlockStmt(node.Body, "while.body")
	var test = c.Function.NewBlock("while.test")
	c.compileBlock(test, func() {
		c.ctx.NewCondJmp(c.compileExpr(node.Test), body, c.ctx.LeaveBlock)
	})
	if body.Terminator == nil {
		body.NewJmp(test)
	}
	c.ctx.NewJmp(test)
}

func (c *Compiler) compileBreakStmt(node *ast.BreakStmt) {
	c.ctx.NewJmp(c.ctx.LeaveBlock)
}
