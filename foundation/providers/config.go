package providers

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/nickbryan/framework/config"
	"github.com/nickbryan/framework/di"
)

type ConfigurationProvider struct{}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func (p *ConfigurationProvider) Register(container di.Container) {
	items := map[string]interface{}{}
	loadedFromCache := false
	configPath, err := container.Make("path.config")
	checkError(err)

	// TODO: load cached config

	conf := config.NewPopulatedRepository(items)

	container.Instance("config", conf)

	if loadedFromCache == false {
		p.loadConfigurationFiles(configPath.(string), conf)
	}

	container.Instance("env", conf.GetDefault("env", "production"))
}

func (p *ConfigurationProvider) loadConfigurationFiles(configPath string, conf *config.Repository) {
	files, err := filepath.Glob(filepath.Join(configPath, "*.json"))
	checkError(err)

	for _, file := range files {
		jsn, err := ioutil.ReadFile(file)
		checkError(err)

		var parsed interface{}
		err = json.Unmarshal([]byte(jsn), &parsed)
		checkError(err)

		for key, val := range parsed.(map[string]interface{}) {
			conf.Set(key, val)
		}
	}
}
