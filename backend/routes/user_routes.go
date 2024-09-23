package routes

import(
	"github.com/gin-gonic/gin"
	"github.com/UprightBiswa/local-marketplace/controllers"
)

//UserRouters defines user routes for regisreation and login
func UserRouters(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/register", controllers.RegisterUser)
	// incomingRoutes.POST("/login", controllers.LoginUser)	
}