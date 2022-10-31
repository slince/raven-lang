package main

class FooClass extends ParentClass implements InterfaceA, InterfaceB{
	public const a: string = "123"
	public static b: bool = false

	abstract public function hello(): string{
		return "hello" + "world";
    }

	final public static function world(): void{
        return 1 + 2
    }
}


let cls = class FooClass extends ParentClass implements InterfaceA, InterfaceB{
	public const a: string = "123"
	public static b: bool = false

	public function hello(): string{
		return "hello" + "world";
    }

	final public static function world(): void{
        return 1 + 2
    }
}