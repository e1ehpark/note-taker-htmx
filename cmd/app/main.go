package main

import (
	"strconv"
	"time"

	"github.com/a-h/templ"
	"github.com/e1ehpark/note-taker-htmx/componets"
	"github.com/e1ehpark/note-taker-htmx/notes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return render(ctx, componets.Index(componets.NotesForm(len(notes.GetAll()))))
	})
	app.Get("/notes", func(c *fiber.Ctx) error {

		return render(ctx, componets.Notes(componets.Notes(len(notes.GetAll()))))
	})
	app.Post("/notes", func(c *fiber.Ctx) error {
		notes.Add(notes.CreateNote{
			Title: ctx.FormValue("title"),
			Body:  ctx.FormValue("body"),
		})
		var notes templ.Component = components.NotesForm(len(notes.GetAll()))
		time.Sleep(1 * time.Second)
		return render(ctx, notes)
	})
	app.Delete("/note/:id", func(ctx *fiber.Ctx) error {
		idStr := ctx.Params("id")

		id, err := strconv.Atoi(idStr)

		if err != nil {
			ctx.Status(404).WriteString(err.Error())
		}

		err = notes.Delete(id)

		if err != nil {
			ctx.Status(404)
		}
		var notes templ.Component = componts.Notes(notes.GetAll())
		return render(ctx, notes)
	})
	app.Listen(":8080")
}
