package main

type Config struct {
	YoutiseURL   string
	SocketServer string
	SerialNumber func() string
}

var appConfig map[string]Config = map[string]Config{
	"development": Config{
		YoutiseURL:   "http://localhost:6000/api/",
		SocketServer: "http://localhost:6000/player-socket",
		SerialNumber: func() string { return "4C4C4544-0035-3210-804B-B4C04F585332" },
	},
	"staging": Config{
		YoutiseURL:   "https://youtise-location-dev.herokuapp.com/api/",
		SocketServer: "https://youtise-location-dev.herokuapp.com/player-socket",
		SerialNumber: func() string { return "" },
	},
	"production": Config{
		YoutiseURL:   "https://location.youtise.com/api/",
		SocketServer: "https://location.youtise.com/player-socket",
		SerialNumber: func() string { return "" },
	},
}

func GetConfig(env string) Config {
	config, ok := appConfig[env]
	if !ok {
		return appConfig["staging"]
	}
	return config
}
