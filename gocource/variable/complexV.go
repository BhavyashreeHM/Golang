package main

func complexV() {
	//complex variable declarations in Go
	var c1 complex64
	c1 = complex(2, 3)
	println("Value of complex number c1 is:", c1)

	var c2 complex128 = complex(5.5, 7.7)
	println("Value of complex number c2 is:", c2)

	var c3 = complex(1.1, 2.2)
	println("Value of complex number c3 is:", c3)

	c4 := complex(9.9, 8.8)
	println("Value of complex number c4 is:", c4)

}
