package main

func main() {
	app := App{}

	app.initialize()
	app.run("localhost:10000")
}
