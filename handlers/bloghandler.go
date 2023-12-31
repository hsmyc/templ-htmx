package handlers

import (
	"context"
	"hsmyc/htmx/db"
	"hsmyc/htmx/models"
	"hsmyc/htmx/views/components/blogpage"
	"hsmyc/htmx/views/layout"
	"time"

	"github.com/a-h/templ"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	blogCollection *mongo.Collection = db.GetCollection(db.DB, "blogs")
	blog           models.Blogmodel
)

func BlogHandler(blogSlug string) templ.Component {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"slug": blogSlug}
	err := blogCollection.FindOne(ctx, filter).Decode(&blog)
	if err != nil {
		return layout.Index(nil, nil, nil)
	}
	return layout.Index(nil, blogpage.Blogpage(blog), nil)
}
