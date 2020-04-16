package main

func f1() int {
	var result int
	defer func() {
		result++

	}()
	return result
}
func f2() (result int) {
	defer func() {
		result++
	}()
	return
}
// GOOS=linux GOARCH=amd64 go tool compile -N -l -S return.go >return.s 2>&1
func main() {
	f1()
	f2()
}
