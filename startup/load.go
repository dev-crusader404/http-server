package startup

import (
	"sync"

	"github.com/magiconair/properties"
)

var (
	appProperties = make(map[string]any)
	onceInit      sync.Once
)

const (
	propertyKey = "propKeys"
)

func Load(path string) {
	onceInit.Do(func() {
		loadproperties(path)
	})
}

func loadproperties(path string) {
	p := properties.MustLoadFile(path+"app.properties", properties.UTF8)
	appProperties[propertyKey] = p
}

func GetAll() *properties.Properties {
	p, _ := appProperties[propertyKey].(*properties.Properties)
	return p
}
