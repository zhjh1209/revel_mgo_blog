package controllers

import (
	"strings"
	"time"

	"github.com/zhjh1209/revel_mgo_blog/app/models"

	"github.com/revel/revel"
)

type WBlog struct {
	App
}

func (c WBlog) Putup(blog *models.Blog) revel.Result {
	blog.Title = strings.TrimSpace(blog.Title)
	blog.Email = strings.TrimSpace(blog.Email)
	blog.Subject = strings.TrimSpace(blog.Subject)
	blog.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.WBlog)
	}
	dao, err := models.NewDao()
	if err != nil {
		c.Response.Status = 500
		return c.RenderError(err)
	}
	defer dao.Close()
	err = dao.CreateBlog(blog)
	if err != nil {
		c.Response.Status = 500
		return c.RenderError(err)
	}
	if !dao.IsEmailExists(blog.Email) {
		dao.InsertEmail(&models.Email{blog.Email, time.Now()})
	}
	return c.Redirect(App.Index)
}
