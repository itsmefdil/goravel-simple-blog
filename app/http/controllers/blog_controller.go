package controllers

import (
	"github.com/goravel/framework/contracts/http"
)

type BlogController struct {
	//Dependent services
}

func NewBlogController() *BlogController {
	return &BlogController{
		//Inject services
	}
}

func (r *BlogController) Index(ctx http.Context) http.Response {
	return nil
}	
