package keys

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type KeyConfig struct {
	BottomKeys   []string
	TopKeys      []string
	DownKeys     []string
	UpKeys       []string
	CollapseKeys []string
	HelpKeys     []string
	QuitKeys     []string
}

func NewConfig(data []byte) (*KeyConfig, error) {
	var c KeyConfig
	err := toml.Unmarshal(data, &c)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshall style config: %w", err)
	}
	return &c, nil
}
