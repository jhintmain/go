package main

func get(p *int) *int {
	example := *p
	return &example
}

func main() {
	var a int
	var p *int = &a
	a = 20
	_ = get(p)
}


