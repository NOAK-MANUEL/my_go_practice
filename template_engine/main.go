package main

func main() {
	template := "Hello {{name}}, you are {{age}} years old."
	data := map[string]string{"name": "Sam", "age": "30"}
	result := render(template, data)
	println(result)
}
