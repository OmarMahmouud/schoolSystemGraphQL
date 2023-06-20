package main

func main() {
	ConnectToDataBase()
	Migrate()
	schema, _ = createSchema()
	run()
}
