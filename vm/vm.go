package vm

import (
	"encoding/binary"
	"fmt"
	"github.com/slince/php-plus/vm/object"
)

type VM struct {
	instructions Instructions
	constants    []object.Object
	globals      []object.Object
	functions    map[string]*Function
	slot         *CallSlot
	stack        *Stack
	// 栈顶地址寄存器
	sp uint64
	// 站基地址寄存器
	bp uint64
	// 指令寄存器
	pc uint64
	// 参数寄存器
	argv [255]object.Object
	// 运算结果寄存器
	r    object.Object
	flag uint8
}

type CallSlot struct {
	prev     *CallSlot
	ip       uint64
	function *Function
}

func New() *VM {
	return &VM{}
}

// 下一个立即数
func (v *VM) nextImm() uint32 {
	var val = binary.BigEndian.Uint32(v.instructions[v.pc : v.pc+4])
	v.pc += 4
	return val
}

// 下一个小立即数
func (v *VM) nextSmaImm() uint8 {
	var val = v.instructions[v.pc]
	v.pc++
	return val
}

// 下一个操作数
func (v *VM) nextDynOperand() Operand {
	var val = v.nextImm()
	var kind = OperandType(val >> 31 & 0x1)
	var num = val & 0x7FFFFFFF
	return Operand{Type: kind, Value: num}
}

// 下一个操作数对应的值
func (v *VM) nextDynOperandValue() object.Object {
	var operand = v.nextDynOperand()
	var val object.Object
	switch operand.Type {
	case OPE_IMM, OPE_SMA_IMM:
		val = &object.Int{Value: int64(operand.Value)}
	case OPE_CONST:
		val = v.constants[operand.Value]
	case OPE_GLOBAL:
		val = v.globals[operand.Value]
	case OPE_NONE:
		val = &object.Null{}
	}
	return val
}

func (v *VM) setGlobal(idx uint32, val object.Object) {
	v.globals[idx] = val
}

func (v *VM) getRegister(idx uint8) object.Object {
	return v.stack.Get(v.bp + uint64(idx)).(object.Object)
}

func (v *VM) setRegister(idx uint8, val object.Object) {
	v.stack.Set(v.bp+uint64(idx), val)
}

// 执行二元运算
func (v *VM) execBinaryOp(opcode Opcode, lhs, rhs object.Object) object.Object {
	var result object.Object
	switch lhs.Type() {
	case object.INT:
		var left = lhs.(*object.Int).Value
		var right = rhs.(*object.Int).Value
		switch opcode {
		// 算术运算
		case OP_ADD:
			result = &object.Int{Value: left + right}
		case OP_SUB:
			result = &object.Int{Value: left - right}
		case OP_MUL:
			result = &object.Int{Value: left * right}
		case OP_DIV:
			result = &object.Int{Value: left / right}
		case OP_MOD:
			result = &object.Int{Value: left % right}
		// 位运算
		case OP_AND:
			result = &object.Int{Value: left & right}
		case OP_OR:
			result = &object.Int{Value: left | right}
		case OP_XOR:
			result = &object.Int{Value: left ^ right}
		case OP_SHL:
			result = &object.Int{Value: left << right}
		case OP_SHR:
			result = &object.Int{Value: left >> right}
		// 逻辑运算
		case OP_EQ:
			if left == right {
				result = object.TrueValue
			} else {
				result = object.FalseValue
			}
		case OP_LT:
			if left < right {
				result = object.TrueValue
			} else {
				result = object.FalseValue
			}
		case OP_GT:
			if left > right {
				result = object.TrueValue
			} else {
				result = object.FalseValue
			}
		case OP_NEQ:
			if left != right {
				result = object.TrueValue
			} else {
				result = object.FalseValue
			}
		case OP_LEQ:
			if left <= right {
				result = object.TrueValue
			} else {
				result = object.FalseValue
			}
		case OP_GEQ:
			if left >= right {
				result = object.TrueValue
			} else {
				result = object.FalseValue
			}
		}
	case object.FLOAT:
		var left = lhs.(*object.Float).Value
		var right = rhs.(*object.Float).Value
		switch opcode {
		// 算术运算
		case OP_ADD:
			result = &object.Float{Value: left + right}
		case OP_SUB:
			result = &object.Float{Value: left - right}
		case OP_MUL:
			result = &object.Float{Value: left * right}
		case OP_DIV:
			result = &object.Float{Value: left / right}
		// 逻辑运算
		case OP_EQ:
			if left == right {
				result = object.TrueValue
			} else {
				result = object.FalseValue
			}
		case OP_LT:
			if left < right {
				result = object.TrueValue
			} else {
				result = object.FalseValue
			}
		case OP_GT:
			if left > right {
				result = object.TrueValue
			} else {
				result = object.FalseValue
			}
		case OP_NEQ:
			if left != right {
				result = object.TrueValue
			} else {
				result = object.FalseValue
			}
		case OP_LEQ:
			if left <= right {
				result = object.TrueValue
			} else {
				result = object.FalseValue
			}
		case OP_GEQ:
			if left >= right {
				result = object.TrueValue
			} else {
				result = object.FalseValue
			}
		}
	case object.BOOL:
		var left = lhs.(*object.Bool).Value
		var right = rhs.(*object.Bool).Value
		switch opcode {
		// 逻辑运算
		case OP_LOGIC_AND:
			if left && right {
				result = object.TrueValue
			} else {
				result = object.FalseValue
			}
		case OP_LOGIC_OR:
			if left || right {
				result = object.TrueValue
			} else {
				result = object.FalseValue
			}
		}
	}
	return result
}

// 执行一元运算
func (v *VM) execUnaryOp(opcode Opcode, operand object.Object) object.Object {
	var result object.Object
	switch operand.Type() {
	case object.INT:
		var intValue = operand.(*object.Int).Value
		switch opcode {
		case OP_INC:
			result = &object.Int{Value: intValue + 1}
		case OP_DEC:
			result = &object.Int{Value: intValue - 1}
		}
	case object.FLOAT:
		var floatValue = operand.(*object.Float).Value
		switch opcode {
		case OP_INC:
			result = &object.Float{Value: floatValue + 1}
		case OP_DEC:
			result = &object.Float{Value: floatValue - 1}
		}
	case object.BOOL:
		var boolValue = operand.(*object.Bool).Value
		switch opcode {
		case OP_LOGIC_NOT:
			if boolValue {
				result = object.FalseValue
			} else {
				result = object.TrueValue
			}
		}
	}
	return result
}

func (v *VM) execute() error {
	var run = true
	for run {
		var err error
		var opcode = Opcode(v.instructions[v.pc])
		switch opcode {
		case OP_NOP:
		case OP_HALT:
			run = false
		// 寄存器设置
		case OP_MOVE:
			var reg1 = v.nextSmaImm()
			var reg2 = v.nextSmaImm()
			v.setRegister(reg1, v.getRegister(reg2))
		case OP_LOAD:
			var reg = v.nextSmaImm()
			// 立即数支持有符号扩展
			var num = int64(v.nextImm())
			v.setRegister(reg, &object.Int{Value: num})
		case OP_LOAD_NULL:
			var start = v.nextSmaImm()
			// 立即数支持有符号扩展
			var num = v.nextSmaImm()
			var null = &object.Null{}
			var i uint8
			for i = 0; i < num; i++ {
				v.setRegister(start+i, null)
			}
		case OP_LOAD_BOOL:
			var start = v.nextSmaImm()
			// 立即数支持有符号扩展
			var num = v.nextSmaImm()
			var flag = v.nextSmaImm()
			var val object.Object
			if flag == 1 {
				val = &object.Bool{Value: true}
			} else {
				val = &object.Bool{Value: false}
			}
			var i uint8
			for i = 0; i < num; i++ {
				v.setRegister(start+i, val)
			}
		case OP_LOAD_CONST:
			var reg = v.nextSmaImm()
			var idx = v.nextImm()
			v.setRegister(reg, v.constants[idx])
		case OP_LOAD_GLOBAL:
			var reg = v.nextSmaImm()
			var idx = v.nextImm()
			v.setRegister(reg, v.globals[idx])
		case OP_STORE_GLOBAL:
			var reg = v.nextSmaImm()
			var idx = v.nextImm()
			v.globals[idx] = v.getRegister(reg)
		// 二元运算
		case OP_ADD, OP_SUB, OP_MUL, OP_DIV, OP_MOD, OP_AND, OP_OR, OP_XOR, OP_NOT, OP_SHL, OP_SHR, OP_LOGIC_AND,
			OP_LOGIC_OR, OP_EQ, OP_LT, OP_GT, OP_NEQ, OP_LEQ, OP_GEQ:
			var reg = v.nextSmaImm()
			var left = v.nextDynOperandValue()
			var right = v.nextDynOperandValue()
			v.setRegister(reg, v.execBinaryOp(opcode, left, right))
		// 一元运算
		case OP_LOGIC_NOT, OP_INC, OP_DEC:
			var reg = v.nextSmaImm()
			var operand = v.nextDynOperandValue()
			v.setRegister(reg, v.execUnaryOp(opcode, operand))
		// 函数调用
		case OP_CALL:
			var reg = v.nextSmaImm()
			var fn = v.getRegister(reg).(*object.String).Value
			err = v.call(fn)
		case OP_ENTER:
			v.stack.Push(v.bp)
			v.bp = v.sp
		case OP_LEAVE:
			v.sp = v.bp
			v.bp = v.stack.Pop().(uint64)
		case OP_RET:
			err = v.ret()
		default:
			err = fmt.Errorf("unknown opcode: %d", opcode)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (v *VM) call(fn string) error {
	var function = v.functions[fn]
	v.slot = &CallSlot{
		prev:     v.slot,
		function: function,
	}
	v.instructions = function.Instructions
	return v.execute()
}

func (v *VM) ret() error {
	v.slot = v.slot.prev
	v.instructions = v.slot.function.Instructions
	v.pc = v.slot.ip
	return v.execute()
}

func (v *VM) Run(bytecode Bytecode) error {
	v.constants = bytecode.Constants
	v.globals = bytecode.Globals
	v.functions = bytecode.Functions
	return v.call(bytecode.Entry)
}
