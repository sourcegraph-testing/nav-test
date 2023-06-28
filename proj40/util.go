package main

import (
	"context"

	process "github.com/go-nacelle/process/v2"
)

type initializer struct{}

func (*initializer) Init(ctx context.Context) error { return nil }

var _ process.Initializer = &initializer{}

//
//

type initializerAndRunner struct{}

func (*initializerAndRunner) Init(ctx context.Context) error { return nil }
func (*initializerAndRunner) Run(ctx context.Context) error  { return nil }

var _ process.Initializer = &initializerAndRunner{}
var _ process.Runner = &initializerAndRunner{}

//
//

type initializerRunnerAndStopper struct{}

func (*initializerRunnerAndStopper) Init(ctx context.Context) error { return nil }
func (*initializerRunnerAndStopper) Run(ctx context.Context) error  { return nil }
func (*initializerRunnerAndStopper) Stop(ctx context.Context) error { return nil }

var _ process.Initializer = &initializerRunnerAndStopper{}
var _ process.Runner = &initializerRunnerAndStopper{}
var _ process.Stopper = &initializerRunnerAndStopper{}

//
//

type initializerRunnerStopperAndFinalizer struct{}

func (*initializerRunnerStopperAndFinalizer) Init(ctx context.Context) error     { return nil }
func (*initializerRunnerStopperAndFinalizer) Run(ctx context.Context) error      { return nil }
func (*initializerRunnerStopperAndFinalizer) Stop(ctx context.Context) error     { return nil }
func (*initializerRunnerStopperAndFinalizer) Finalize(ctx context.Context) error { return nil }

var _ process.Initializer = &initializerRunnerStopperAndFinalizer{}
var _ process.Runner = &initializerRunnerStopperAndFinalizer{}
var _ process.Stopper = &initializerRunnerStopperAndFinalizer{}
var _ process.Finalizer = &initializerRunnerStopperAndFinalizer{}
