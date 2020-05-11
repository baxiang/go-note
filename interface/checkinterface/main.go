package main

type Foo interface {
	bar()
}
type FooType struct {

}

//func (f FooType)bar(){
//	fmt.Println("hello world")
//}
func main() {
	var _ Foo = (*FooType)(nil)
}
