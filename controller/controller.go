package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lambda-platform/ebarimt-rest-api/ebarimt"
	"github.com/lambda-platform/ebarimt/posapi"
)

// @Summary_en        Retrieve API information
// @Summary        API-ийн мэдээллийг авах
// @Description_en    Get information about the API, such as its version and name.
// @Description    API-ийн талаарх мэдээллийг, жишээ нь хувилбар ба нэрээ авах.
// @Tags              ebarimt
// @Produce           json
// @Success           200 {object} posapi.InformationOutput
// @Router            /info [get]
func Info(c *fiber.Ctx) error {

	info, err := ebarimt.PosAPI.GetInformation()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to call GetInformation function"})
	}

	return c.JSON(info)
}

// @Summary_en        Check API connectivity and status
// @Summary        API-ийн холболтыг шалгах ба статусыг авах
// @Description_en    Check the API connection and retrieve the status of the database, configuration, and network.
// @Description    API-ийн холболтыг шалгах мөн тохиргооны мэдээлл, сүлжээний статусыг аваh, өгөгдөлийн сантай харьцахад саад буй эсэхийг шалгаж болно
// @Tags              ebarimt
// @Produce           json
// @Success           200 {object} posapi.CheckOutput
// @Router            /check [get]
func Check(c *fiber.Ctx) error {

	result, err := ebarimt.PosAPI.CheckApi()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to call CheckApi function"})
	}

	return c.JSON(fiber.Map{"result": result})
}

// @Summary_en        Invoke a specified function
// @Summary        Тодорхой функцыг дуудах
// @Description_en    Call a function with the specified name and parameters.
// @Description    Тодорхой нэртэй болон параметртэй функцыг дуудах.
// @Tags              ebarimt
// @Produce           json
// @Param             functionName query string true "Name of the function to call"
// @Param             params query string true "Parameters to pass to the function"
// @Success           200 {string} string
// @Router            /call [get]
func Call(c *fiber.Ctx) error {

	functionName := c.Query("functionName")
	params := c.Query("params")
	result, err := ebarimt.PosAPI.CallFunction(functionName, params)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to call CallFunction function"})
	}

	return c.JSON(result)
}

// @Summary_en        Submit a bill to the system
// @Summary        Системд нэхэмжлэхийг оруулах
// @Description_en    Add a bill to the system with the provided data.
// @Description    Өгөгдсөн мэдээллийг ашиглан системд нэхэмжлэхийг нэмэх.
// @Tags              ebarimt
// @Accept            json
// @Produce           json
// @Param             data body posapi.PutInput true "Bill data"
// @Success           200 {object} posapi.PutOutput
// @Router            /put [post]
func Put(c *fiber.Ctx) error {
	input := posapi.PutInput{}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	result, err := ebarimt.PosAPI.Put(input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to call Put function"})
	}

	return c.JSON(fiber.Map{"result": result})
}

// @Summary_en        Process a bill return
// @Summary        Нэхэмжлэхийг буцаах
// @Description Return a bill with the provided data.
// @Description Өгөгдсөн мэдээллийг ашиглан нэхэмжлэхийг буцаах.
// @Tags ebarimt
// @Accept json
// @Produce json
// @Param data body posapi.BillInput true "Return bill data"
// @Success 200 {object} posapi.BillOutput
// @Router /return [post]
func Return(c *fiber.Ctx) error {

	input := posapi.BillInput{}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	result, err := ebarimt.PosAPI.ReturnBill(input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to call ReturnBill function"})
	}

	return c.JSON(fiber.Map{"result": result})
}

// @Summary      Transmit data
// @Description  Send data via the API.
// @Tags         ebarimt
// @Produce      json
// @Success      200 {object} posapi.DataOutput
// @Router       /send [get]
func Send(c *fiber.Ctx) error {

	result, err := ebarimt.PosAPI.SendData()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to call SendData function"})
	}

	return c.JSON(fiber.Map{"result": result})
}
