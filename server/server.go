package server

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/nizhunt/urlShortner/model"
	"github.com/nizhunt/urlShortner/utils"
)

func redirect(app *fiber.Ctx) error {
	golyURL := app.Params("redirect")
	goly, err := model.FindByGolyUrl(golyURL)
	if err != nil {
		return app.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"could not find the url in db" + err.Error(),
		})
	}

	goly.Clicked +=1
	err = model.UpdateGoly(goly)
	if err!= nil {
		fmt.Printf("err updating:%v\n", err)
	}
	return app.Redirect(goly.Redirect, fiber.StatusTemporaryRedirect)
}

func getAllGolies(app *fiber.Ctx) error {
	golies, err := model.GetAllGolies()
	if err != nil {
		return app.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting all goly links" + err.Error(),
		})
	}
	return app.Status(fiber.StatusOK).JSON(golies)
}

func getGoly(app *fiber.Ctx) error {
	id, err := strconv.ParseUint( app.Params("id"), 10, 64)
	if err != nil {
		return app.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not parse id" + err.Error(),
		})
	}

	goly, err := model.GetGoly(id)
	if err!= nil {
		return app.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Goly Id not found" + err.Error(),
		})
	}
	return app.Status(fiber.StatusOK).JSON(goly)
}

func createGoly(app *fiber.Ctx) error {
	var goly model.Goly

	app.Accepts("application/json")
	err := app.BodyParser(&goly)
	if err != nil{
		return app.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"error parsing JSON"+ err.Error(),
		})
	}

	if goly.Random {
		goly.Goly = utils.RandomURL(8)
	}

	err = model.CreateGoly(goly)
	if err!= nil{
		return app.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not create goly in db" + err.Error(),
		})
	}

	return app.Status(fiber.StatusOK).JSON(goly)
}

func updateGoly(app *fiber.Ctx) error {
	app.Accepts("application/json")

	var goly model.Goly

	err:= app.BodyParser(&goly)
	if err != nil{
		return app.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"could not parse json" + err.Error(),
		})
	}
	
	err = model.UpdateGoly(goly)
	if err != nil{
	return app.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message":"could not parse json" + err.Error(),
		})
	}

	return app.Status(fiber.StatusOK).JSON(goly)
}

func deleteGoly(c *fiber.Ctx) error {
	id, err := strconv.ParseUint( c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not parse id from url " + err.Error(),
		})
	}

	err = model.DeleteGoly(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not delete from db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map {
		"message": "goly deleted.",
	})
}

func SetupAndListen() {
	 app := fiber.New()

	 app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders:"Origin, Content-Type, Accept",
	 }))

	 app.Get("/goly", getAllGolies)
	 app.Get("/goly/:id", getGoly)
	 app.Post("/goly", createGoly)
	 app.Delete("/goly/:id", deleteGoly)
	 app.Patch("/goly", updateGoly)
	 app.Get("/r/:redirect", redirect)

	 

	 app.Listen(":4545")
}