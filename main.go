// A generated module for DWithfiles functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"dagger/withfiles/internal/dagger"
)

type Withfiles struct{}

func (m *Withfiles) Container() *dagger.Container {
	dir := dag.Directory().
		WithFile("/sub/file", dag.File("file", "contents", dagger.FileOpts{Permissions: 0644})).
		Directory("/sub")
	return dag.Container().
		From("alpine:3.22.1").
		WithFiles("/opt", []*dagger.File{
			dir.File("file"),
		})
}
