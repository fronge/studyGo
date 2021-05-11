package main

func a() (err error) {
	return nil
}

func main() {
	go a()
}
