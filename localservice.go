package localservice

import (
	"github.com/dynamicgo/go-config"
	"github.com/dynamicgo/xerrors"
	"github.com/gomeshnetwork/gomesh"
)

// F .
type F func() (gomesh.Service, error)

type localServiceImp struct {
	creators map[string]F
	builder  gomesh.ModuleBuilder
}

// LocalService .
type LocalService interface {
	Register(name string, f F)
}

// New .
func New(mesh gomesh.Mesh) LocalService {
	impl := &localServiceImp{
		creators: make(map[string]F),
	}

	impl.builder = mesh.Module(impl)

	return impl
}

func (module *localServiceImp) Register(name string, f F) {
	module.creators[name] = f
}

func (module *localServiceImp) Start(config config.Config) error {
	return nil
}

func (module *localServiceImp) Name() string {
	return "gomesh.module.local"
}

func (module *localServiceImp) BeginCreateService() error {
	return nil
}
func (module *localServiceImp) CreateService(name string) (gomesh.Service, error) {
	f, ok := module.creators[name]

	if !ok {
		return nil, xerrors.Errorf("module %s service %s creator not found", module.Name(), name)
	}

	return f()
}

func (module *localServiceImp) EndCreateService() error {
	return nil
}

func (module *localServiceImp) BeginSetupService() error {
	return nil
}

func (module *localServiceImp) SetupService(service gomesh.Service) error {
	return nil
}

func (module *localServiceImp) EndSetupService() error {
	return nil
}

func (module *localServiceImp) BeginStartService() error {
	return nil
}

func (module *localServiceImp) StartService(service gomesh.Service, config config.Config) error {
	return service.Start(config)
}

func (module *localServiceImp) EndStarService() error {
	return nil
}
