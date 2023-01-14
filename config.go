package main

type Config struct {
	volumesFrom string
	volumesTo string
}

func load(volume string) Config {
	config := Config{"1", "2"}
	return config
}
