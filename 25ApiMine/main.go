package main

type Games struct {
	Game_id        string     `json:"id"`
	Game_name      string     `json:"name"`
	Game_price     float64    `json:"price"`
	Game_meta_data *Meta_data `json:"meta"`
	//can also remove * and it'll work but it copies it, using * will use reference of it rather that copyinf it
}

type Meta_data struct {
	Studio_name  string `json:"studio"`
	Release_data string `json:"release-date"`
}

func main() {

}
