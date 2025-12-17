package logger

type Config struct {
	Level    string `mapstructure:"level"`
	Format   string `mapstructure:"format"`
	Dir      string `mapstructure:"director"`
	Filename string `mapstructure:"filename"`
	MaxSize  int    `mapstructure:"max_size"`
	MaxAge   int    `mapstructure:"max_age"`
	Compress bool   `mapstructure:"compress"`
}
