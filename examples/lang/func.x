package main

import os/console

function hello(a: string, b: int64): string{
	return a + b
}

let func = function hello2(a: string, b: int64): string{
	return a + b
}

console.print(hello("hello", "world"));