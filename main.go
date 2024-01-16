// main.go
package main


func main() {
	g, err := NewGame()
	if err != nil {
		panic(err)
	}
	g.run()
}
