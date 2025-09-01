//go:build wireinject
// +build wireinject

package simple

import (
	"io"
	"os"
	"github.com/google/wire"
)

func InitializeService(isError bool) (*SimpleService, error) {
	wire.Build(
		NewSimpleRepository,NewSimpleService,
	)
	return nil, nil
}

func InitializeDatabaseRepository() *DatabaseRepository {
	wire.Build(
		NewDatabaseMongoDB,
		NewDatabasePostgreSql,
		NewDatabaseRepository,
	)
	return nil
}

var fooSet = wire.NewSet(
	NewFooRepository,
	NewFooService,
)

var barSet = wire.NewSet(
	NewBarRepository,
	NewBarService,
)

func InitializeFooBarService() *FooBarService {
	wire.Build(
		fooSet,
		barSet,
		NewFooBarService,
	)
	return nil
}

var helloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

func InitializeHelloService() *HelloService {
	wire.Build(
		helloSet,
		NewHelloService,
	)
	return nil
}

func InitializedFooBar() *FooBar {
	wire.Build(
		NewFoo,
		NewBar,
		wire.Struct(new(FooBar), "Foo", "Bar"),
	)
	return nil
}

var fooValue = &Foo{}
var barValue = &Bar{}

func InitializedFooBarUseingValue() *FooBar{
	wire.Build(
		wire.Value(fooValue),
		wire.Value(barValue),
		wire.Struct(new(FooBar), "*"),
	)
	return nil
}

func InitializedReader() io.Reader {
	wire.Build(
		wire.InterfaceValue(new(io.Reader), os.Stdin),
	)
	return nil
}

func InitializedConfiguration() *Configuration {
	// application := NewApplication()
	// configuration := application.Configuration
	// return configuration

	wire.Build(
		NewApplication,
		wire.FieldsOf(new(*Application), "Configuration"),
	)

	return nil
}

func InitializedConnection(name string) (*Connection, func()){
	wire.Build(NewConnection, NewFile)
	return nil, nil
}