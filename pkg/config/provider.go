package config

type provider struct {
	mySQL          MySQL
	unsplashConfig UnsplashConfig
}

type Provider interface {
}
