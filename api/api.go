package api

import (
	"errors"
	"fmt"

	"github.com/aashuprogrammer/document_uploader.git/db/pgdb"
	"github.com/aashuprogrammer/document_uploader.git/token"
	"github.com/aashuprogrammer/document_uploader.git/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/etag"
	"github.com/gofiber/fiber/v3/middleware/favicon"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

type Server struct {
	app    *fiber.App
	store  pgdb.Store
	valid  *validator.Validate
	config utils.Config
	token  token.TokenMaker
}

func NewServer(config utils.Config, store pgdb.Store, tokenMaker token.TokenMaker) (*Server, error) {

	if store == nil {
		return nil, errors.New("store cannot be nil")
	}
	if tokenMaker == nil {
		return nil, errors.New("tokenMaker cannot be nil")
	}
	server := &Server{
		valid:  validator.New(),
		config: config,
		store:  store,
		token:  tokenMaker,
	}
	server.setupApi()
	return server, nil
}

func (server *Server) Start(port int16) error {
	return server.app.Listen(fmt.Sprintf(":%d", port))
}

type msgResponse struct {
	Msg string `json:"msg"`
}

func (server *Server) setupApi() {
	app := fiber.New(fiber.Config{
		ServerHeader:  "Inflection-Fiber",
		ErrorHandler:  errorHandler,
		BodyLimit:     2 * 1024 * 1024,
		CaseSensitive: true,
	})

	app.Use(logger.New(logger.ConfigDefault))

	app.Use(cors.New())

	app.Use(compress.New())

	// app.Use(csrf.New())

	app.Use(etag.New())

	app.Use(favicon.New())

	// app.Use(limiter.New())

	app.Get("/hello", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "hello"})
	})
	app.Post("/login", server.login)
	server.app = app
}
