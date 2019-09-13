package localservice

import (
	"testing"

	"github.com/dynamicgo/go-config"
	"github.com/gomeshnetwork/gomesh"
	"github.com/stretchr/testify/require"
)

type testA struct {
	B *testB `inject:"test.B"`
}

func (a *testA) Start(config config.Config) error {
	return nil
}

type testB struct {
	A *testA `inject:"test.A"`
}

func (b *testB) Start(config config.Config) error {
	return nil
}

func TestInject(t *testing.T) {
	mesh := gomesh.New()

	localService := New(mesh)

	localService.Register("test.A", func() (gomesh.Service, error) {
		return &testA{}, nil
	})

	localService.Register("test.B", func() (gomesh.Service, error) {
		return &testB{}, nil
	})

	err := mesh.Start()

	require.NoError(t, err)

	var a *testA

	mesh.ServiceByName("test.A", &a)

	require.NotNil(t, a)

	var b *testB

	mesh.ServiceByName("test.B", &b)

	require.NotNil(t, b)

	require.Equal(t, b.A, a)

	require.Equal(t, a.B, b)

}
