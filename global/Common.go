package global

import "github.com/BurntSushi/toml"

func GetConfig(fileName string) (sysConfig SysConfig, err error) {
	_, err = toml.DecodeFile(fileName, &sysConfig)
	return
}
