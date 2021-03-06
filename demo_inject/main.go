package main

import (
	"fmt"
	"os"
	"github.com/facebookgo/inject"
	"app/container"
)

func main() {

	var g inject.Graph

	// We provide our graph two "seed" objects, one our empty
	// HomePlanetRenderApp instance which we're hoping to get filled out,
	// and second our DefaultTransport to satisfy our HTTPTransport
	// dependency. We have to provide the DefaultTransport because the
	// dependency is defined in terms of the http.RoundTripper interface,
	// and since it is an interface the library cannot create an instance
	// for it. Instead it will use the given DefaultTransport to satisfy
	// the dependency since it implements the interface:
	var a container.Container
	err := g.Provide(
		&inject.Object{Value: &a},
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Here the Populate call is creating instances of NameAPI &
	// PlanetAPI, and setting the HTTPTransport on both to the
	// http.DefaultTransport provided above:
	if err := g.Populate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// There is a shorthand API for the simple case which combines the
	// three calls above is available as inject.Populate:
	//
	//   inject.Populate(&a, http.DefaultTransport)
	//
	// The above API shows the underlying API which also allows the use of
	// named instances for more complex scenarios.

	a.Render()
	a.Render()
}