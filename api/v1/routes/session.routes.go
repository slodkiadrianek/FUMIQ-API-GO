package routes

import "FUMIQ_API/middleware"

type SessionRoutes struct {
	SessionController
	AuthMiddleware *middleware.AuthMiddleware
}
