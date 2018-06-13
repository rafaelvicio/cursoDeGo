package matematica

//Divisor - divisão
func Divisor(x int, b int) (resultado int) {
	resultado = x / b
	return
}

//DivisorComResto - divisão com resto
func DivisorComResto(x int, b int) (resultado int, resto int) {
	resultado = x / b
	resto = x % b
	return
}
