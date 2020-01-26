package main

import (
	"github.com/abyssparanoia/catharsis-gcp/default/handler/api"
	"github.com/abyssparanoia/catharsis-gcp/default/infrastructure/repository"
	"github.com/abyssparanoia/catharsis-gcp/default/usecase"
	"github.com/abyssparanoia/catharsis-gcp/internal/pkg/gluefirebaseauth"
	"github.com/abyssparanoia/catharsis-gcp/internal/pkg/gluemysql"
	"github.com/abyssparanoia/catharsis-gcp/internal/pkg/httpheader"
	"github.com/abyssparanoia/catharsis-gcp/internal/pkg/log"
)

// Dependency ... dependency
type Dependency struct {
	Log              *log.Middleware
	gluefirebaseauth *gluefirebaseauth.Middleware
	DummyHTTPHeader  *httpheader.Middleware
	HTTPHeader       *httpheader.Middleware
	UserHandler      *api.UserHandler
}

// Inject ... indect dependency
func (d *Dependency) Inject(e *Environment) {

	var lCli log.Writer
	var firebaseauth gluefirebaseauth.Firebaseauth

	authCli := gluefirebaseauth.NewClient(e.ProjectID)
	// fCli := gluefirestore.NewClient(e.ProjectID)

	if e.ENV == "LOCAL" {
		lCli = log.NewWriterStdout()
		firebaseauth = gluefirebaseauth.NewDebug(authCli)
	} else {
		lCli = log.NewWriterStackdriver(e.ProjectID)
		firebaseauth = gluefirebaseauth.New(authCli)
	}

	// Config
	dbCfg := gluemysql.NewConfig()

	// pkg
	dbConn := gluemysql.NewClient(dbCfg)

	// Repository
	uRepo := repository.NewUser(dbConn)

	// Service
	dhh := httpheader.NewDummy()
	hh := httpheader.New()
	uSvc := usecase.NewUser(uRepo)

	// Middleware
	d.Log = log.NewMiddleware(lCli, e.MinLogSeverity)
	d.gluefirebaseauth = gluefirebaseauth.NewMiddleware(firebaseauth)
	d.DummyHTTPHeader = httpheader.NewMiddleware(dhh)
	d.HTTPHeader = httpheader.NewMiddleware(hh)

	// Handler
	d.UserHandler = api.NewUserHandler(uSvc)
}
