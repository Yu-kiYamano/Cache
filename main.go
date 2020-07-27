package main

type database interface {
	connect()
	set()
	get()
}

type mysql struct {
}

type cache struct {
}

//インターフェースを返すメソッド
func providerhandler() interface {
}

func main() {

}
