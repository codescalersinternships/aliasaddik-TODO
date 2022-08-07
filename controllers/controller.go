package controllers

import (
	"fmt"
	"net/http"

	"github.com/aliasaddik/todo-project/models"
	"github.com/aliasaddik/todo-project/services"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Handle services.Handles
}

func New(handle services.Handles) Controller {
	return Controller{
		Handle: handle,
	}
}

func (controller *Controller) CreateTask(ctx *gin.Context) {

	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := controller.Handle.CreateTask(&task)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	//change the status ok to the specific request
	ctx.JSON(http.StatusAccepted, gin.H{"message": "success"})
}

func (Controller *Controller) GetTask(ctx *gin.Context) {

	tasks, err := Controller.Handle.GetTask()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}
func (Controller *Controller) EditTask(ctx *gin.Context) {

	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := Controller.Handle.EditTask(&task)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{"message": "success"})
}
func (Controller *Controller) DeleteTask(ctx *gin.Context) {

	objID := ctx.Param("id")

	fmt.Print("my objID", objID)
	err := Controller.Handle.DeleteTask(objID)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{"message": "success"})
}

func (controller *Controller) RegisterRoutes(routerGroup *gin.RouterGroup) {

	userroute := routerGroup.Group("/")

	userroute.POST("/", controller.CreateTask)
	userroute.GET("/", controller.GetTask)
	userroute.PUT("/", controller.EditTask)
	userroute.DELETE("/:id", controller.DeleteTask)

}
