package localservice

import (
	"sync"

	"github.com/dynamicgo/go-config"
	"github.com/dynamicgo/xerrors"
	"github.com/gomeshnetwork/gomesh"
)

// F .
type F func(config config.Config) (gomesh.Service, error)

type localServiceExtension struct {
	creators map[string]F
}

func newExtension() *localServiceExtension {
	return &localServiceExtension{
		creators: make(map[string]F),
	}
}

func (extension *localServiceExtension) register(name string, f F) {
	extension.creators[name] = f
}

func (extension *localServiceExtension) Name() string {
	return "gomesh.extension.local"
}

func (extension *localServiceExtension) Begin(config config.Config, builder gomesh.MeshBuilder) error {

	for name := range extension.creators {
		builder.RegisterService(extension.Name(), name)
	}

	return nil
}

func (extension *localServiceExtension) CreateSerivce(serviceName string, config config.Config) (gomesh.Service, error) {
	f, ok := extension.creators[serviceName]

	if !ok {
		return nil, xerrors.Wrapf(gomesh.ErrNotFound, "service %s not found", serviceName)
	}

	return f(config)
}

func (extension *localServiceExtension) End() error {
	return nil
}

var extension *localServiceExtension
var once sync.Once

func get() *localServiceExtension {
	once.Do(func() {
		extension = newExtension()
		gomesh.Builder().RegisterExtension(extension)
	})

	return extension
}

// Register .
func Register(name string, f F) {
	get().register(name, f)
}
