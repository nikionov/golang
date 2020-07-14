package routes

import (
	"fmt"

	"github.com/nikionov/golang/db/documents"
	"github.com/nikionov/golang/models"
	"github.com/nikionov/golang/session"

	"github.com/martini-contrib/render"

	"gopkg.in/mgo.v2"
)

func IndexHandler(rnd render.Render, s *session.Session, db *mgo.Database) {
	fmt.Println(s.Username)

	postDocuments := []documents.PostDocument{}
	postsCollection := db.C("posts")
	postsCollection.Find(nil).All(&postDocuments)

	posts := []models.Post{}
	for _, doc := range postDocuments {
		post := models.Post{doc.Id, doc.Title, doc.ContentHtml, doc.ContentMarkdown}
		posts = append(posts, post)
	}

	model := models.PostListModel{}
	model.IsAuthorized = s.IsAuthorized
	model.Posts = posts

	rnd.HTML(200, "index", model)
}
