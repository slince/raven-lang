package ir

type Context struct {
	*BasicBlock
	LeaveBlock Block
	Outer      *Context
}

//func (c *Context) EnterContext(b *BasicBlock) *Context {
//	var ctx = NewContext(b)
//	ctx.Outer = c
//	return ctx
//}
//
//func (c *Context) LeaveContext() *Context {
//	return c.Outer
//}

func NewContext(block *BasicBlock, outer *Context) *Context {
	return &Context{
		BasicBlock: block,
		Outer:      outer,
	}
}
