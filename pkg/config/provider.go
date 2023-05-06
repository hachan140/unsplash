package config

type provider struct {
	mySQL          MySQL
	unsplashConfig UnsplashConfig
}

type Provider interface {
	MySQL() MySQL
	UnsplashConfig() UnsplashConfig
}

func NewConfigProvider() Provider {
	return &provider{
		mySQL:          NewMySQLConfig(),
		unsplashConfig: NewUnsplashConfig(),
	}
}

func (p *provider) MySQL() MySQL {
	return p.mySQL
}
func (p *provider) UnsplashConfig() UnsplashConfig {
	return p.unsplashConfig
}
