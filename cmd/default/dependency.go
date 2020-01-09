package main

import (
	"github.com/abyssparanoia/catharsis-gcp/default/handler/api"
	"github.com/abyssparanoia/catharsis-gcp/default/infrastructure/repository"
	"github.com/abyssparanoia/catharsis-gcp/default/service"
	"github.com/abyssparanoia/catharsis-gcp/pkg/firebaseauth"
	"github.com/abyssparanoia/catharsis-gcp/pkg/httpheader"
	"github.com/abyssparanoia/catharsis-gcp/pkg/log"
	"github.com/abyssparanoia/catharsis-gcp/pkg/mysql"
)

// Dependency ... dependency
type Dependency struct {
	Log             *log.Middleware
	FirebaseAuth    *firebaseauth.Middleware
	DummyHTTPHeader *httpheader.Middleware
	HTTPHeader      *httpheader.Middleware
	UserHandler     *api.UserHandler
}

// Inject ... indect dependency
func (d *Dependency) Inject(e *Environment) {

	var lCli log.Writer
	var firebaseAuth firebaseauth.Firebaseauth

	if e.ENV == "LOCAL" {
		lCli = log.NewWriterStdout()
		firebaseAuth = firebaseauth.NewDebug()
	} else {
		lCli = log.NewWriterStackdriver(e.ProjectID)
		firebaseAuth = firebaseauth.New()
	}

	// fCli := cloudfirestore.NewClient(e.ProjectID)

	// Config
	dbCfg := mysql.NewConfig()

	// pkg
	dbConn := mysql.NewClient(dbCfg)

	// Repository
	uRepo := repository.NewUser(dbConn)

	// Service
	dhh := httpheader.NewDummy()
	hh := httpheader.New()
	uSvc := service.NewUser(uRepo)

	// Middleware
	d.Log = log.NewMiddleware(lCli, e.MinLogSeverity)
	d.FirebaseAuth = firebaseauth.NewMiddleware(firebaseAuth)
	d.DummyHTTPHeader = httpheader.NewMiddleware(dhh)
	d.HTTPHeader = httpheader.NewMiddleware(hh)

	// Handler
	d.UserHandler = api.NewUserHandler(uSvc)
}
