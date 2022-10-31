package llvm

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/slince/php-plus/compiler/ast"
	"math"
)

type Compiler struct {
	ctx *Context
}

func NewCompiler() Compiler {
	var block = ir.NewBlock("root")
	return Compiler{
		ctx: NewContext(block),
	}
}

func (c *Compiler) compile(node ast.Node){
	var m = ir.NewModule()
}

func (c *Compiler) compileExpr(node ast.Expr) value.Value {

}

func (c *Compiler) compileIdentifier(node ast.Identifier) string{
	return node.Value
}

func (c *Compiler) compileLiteralExpr(node ast.Literal) constant.Constant{
	var con constant.Constant
	switch node.Value.(type) {
	case int64:
		con = constant.NewInt(types.I32, node.Value.(int64))
	case bool:
		con = constant.NewBool(node.Value.(bool))
	case float64:
		var num = node.Value.(float64)
		if num > math.MaxFloat32 {
			con = constant.NewFloat(types.Float, node.Value.(float64))
		} else {
			con = constant.NewFloat(types.Double, node.Value.(float64))
		}
	case string:
		con = constant.NewCharArrayFromString(node.Value.(string))
	}
	return con
}

func (c *Compiler) compileStmt(node ast.Stmt){
	switch node.(type) {
	case ast.BlockStmt:
		c.compileStmt(node)
	case ast.ExpressionStmt:
		c.compileExprStmt(node.(ast.ExpressionStmt))
	case ast.WhileStmt:
		c.compileWhileStmt(node.(ast.WhileStmt))
	case ast.DoWhileStmt:
		c.compileDoWhileStmt(node.(ast.DoWhileStmt))
	case ast.SwitchStmt:
		c.compileSwitchStmt(node.(ast.SwitchStmt))
	case ast.BreakStmt:
		c.compileReturnStmt(node.(ast.ReturnStmt))
	}
}

func (c *Compiler) compileBlockStmt(node ast.BlockStmt){
	for _, stmt := range node.Body {
		c.compileStmt(stmt)
	}
}

func (c *Compiler) compileExprStmt(node ast.ExpressionStmt){
	c.compileExpr(node.Expr)
}

func (c *Compiler) compileReturnStmt(node ast.ReturnStmt){
	c.ctx.NewRet(c.compileExpr(node.Argument))
}

func (c *Compiler) compileBreakStmt(node ast.BreakStmt){
	c.ctx.NewBr(c.ctx.leaveBlock)
}

func (c *Compiler) compileIfStmt(node ast.IfStmt) {
	var ifThen = ir.NewBlock("if.then")
	c.newContext(ifThen)
	c.compileStmt(node.Consequent)
	var ifElse = ir.NewBlock("if.else")
	if node.Alternate != nil {
		c.newContext(ifElse)
		c.compileStmt(node.Alternate)
	}
	c.ctx.NewCondBr(c.compileExpr(node.Test), ifThen, ifElse)
	if c.ctx.Term == nil {
		leaveB := ir.NewBlock("leave.if")
		c.ctx.NewBr(leaveB)
	}
	c.endContext()
}

func (c *Compiler) compileDoWhileStmt(node ast.DoWhileStmt) {
	var whileBody = ir.NewBlock("do.while.body")
	var leaveBlock = ir.NewBlock("leave.do.while")

	c.ctx.NewBr(whileBody)

	c.subCompile(whileBody, func() {
		c.compileStmt(node.Body)
		c.ctx.leaveBlock = leaveBlock
		c.ctx.NewCondBr(c.compileExpr(node.Test), whileBody, leaveBlock)
	})
}

func (c *Compiler) compileWhileStmt(node ast.WhileStmt) {
	var whileTest = ir.NewBlock("while.test")
	var whileBody = ir.NewBlock("while.body")
	var leaveBlock = ir.NewBlock("leave.do.while")

	c.ctx.NewBr(whileTest)
	c.subCompile(whileTest, func() {
		c.ctx.NewCondBr(c.compileExpr(node.Test), whileBody, leaveBlock)
		c.ctx.leaveBlock = leaveBlock
	})

	c.subCompile(whileBody, func() {
		c.ctx.leaveBlock = leaveBlock
		c.compileStmt(node.Body)
		c.ctx.NewBr(whileTest)
	})
}

func (c *Compiler) compileSwitchStmt(node ast.SwitchStmt) {
	var cases = make([]*ir.Case, 0)
	for _, ca := range node.Cases {
		var caseBlock = ir.NewBlock("switch.case")
		c.subCompile(caseBlock, func() {
			for _, consequent := range ca.Consequent {
				c.compileStmt(consequent)
			}
		})
		cases = append(cases, ir.NewCase(c.compileLiteralExpr(ca.Test), caseBlock))
	}
	var defaultBlock = ir.NewBlock("switch.default")
	c.subCompile(defaultBlock, func() {
		c.compileStmt(node)
	})
	c.ctx.NewSwitch(c.compileExpr(node.Discriminant), defaultBlock, cases...)
}

func (c *Compiler) compileVarDecl(node ast.VariableDeclarator){

}

func (c *Compiler) compileFuncDecl(node ast.FunctionDeclaration){
	var fn = ir.NewFunc(node.)
}

func (c *Compiler) compileType(node ast.Identifier) types.Type {
	var _type types.Type
	switch node.Value {
	case "int64":
		_type = types.I64
	case "int32":
		_type = types.I32
	case "float32":
		_type = types.Float
	case "float64":
		_type = types.Double
	}
	return _type
}

func (c *Compiler) compileFunc(node ast.Function) *ir.Func{
	var id = c.compileIdentifier(node.Id.(ast.Identifier))

	var kind types.Type = types.Void
	if node.Kind != nil {
		kind = c.compileType(node.Kind.(ast.Identifier))
	}

	var args []*ir.Param
	for _, arg := range node.Args {
		args = append(args, c.compileFuncArg(arg))
	}

	var fn = ir.NewFunc(id, kind, args ...)

	var block = fn.NewBlock("")

	c.subCompile(block, func() {
		c.compileBlockStmt(node.Body)
	})
	return fn
}

func (c *Compiler) compileFuncArg(node ast.FunctionArgument) *ir.Param {
	return ir.NewParam(c.compileIdentifier(node.Id), c.compileType(node.Kind))
}

func (c *Compiler) subCompile(b *ir.Block, fn func()){
	c.newContext(b)
	fn()
	c.endContext()
}

func (c *Compiler) newContext(b *ir.Block) *Context {
	 c.ctx = c.ctx.NewContext(b)
	 return c.ctx
}

func (c *Compiler) endContext(){
	c.ctx = c.ctx.parent
}
