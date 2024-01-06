package server

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sanjevscet/go-microservices/internal/database"
	"github.com/sanjevscet/go-microservices/internal/models"
)

type Server interface {
	Start() error
	Readiness(ctx echo.Context) error
	Liveness(ctx echo.Context) error

	GetAllCustomers(ctx echo.Context) error
	GetAllVendors(ctx echo.Context) error
	GetAllProducts(ctx echo.Context) error
	GetAllServices(ctx echo.Context) error
}

type EchoServer struct {
	echo *echo.Echo
	DB   database.DatabaseClient
}

func NewEchoServer(db database.DatabaseClient) Server {
	server := &EchoServer{
		echo: echo.New(),
		DB:   db,
	}
	server.registerRoutes()

	return server
}

// Start implements Server.
func (s *EchoServer) Start() error {
	if err := s.echo.Start(":1414"); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server shutdown error occurred %s", err)
	}

	return nil
}

func (s *EchoServer) Readiness(ctx echo.Context) error {
	ready := s.DB.Ready()
	if ready {
		return ctx.JSON(http.StatusOK, models.Health{Status: "OK"})
	}

	return ctx.JSON(http.StatusInternalServerError, models.Health{Status: "Failure"})
}

func (s *EchoServer) Liveness(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.Health{Status: "OK"})
}

func (s *EchoServer) registerRoutes() {
	s.echo.GET("/readiness", s.Readiness)
	s.echo.GET("/liveness", s.Liveness)

	// Customer Resources
	cg := s.echo.Group("/customers")
	cg.GET("", s.GetAllCustomers)

	// Vendor Resources
	vg := s.echo.Group("/vendors")
	vg.GET("", s.GetAllVendors)

	// Customer Products
	pg := s.echo.Group("/products")
	pg.GET("", s.GetAllProducts)

	// Service Resources
	sg := s.echo.Group("/services")
	sg.GET("", s.GetAllServices)
}
