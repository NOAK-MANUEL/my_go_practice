package main

type Writer interface {
	Write()
}
type File struct{}

func (f File) Write() {
	println("Printing file")
}

type Db struct{}

func (db Db) Write() {
	println("Printing file")
}
func save(w Writer) {
	w.Write()
}
func main() {
	save(File{})
}
