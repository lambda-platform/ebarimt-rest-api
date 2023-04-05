package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lambda-platform/ebarimt-rest-api/ebarimt"
)

// @Summary        Retrieve API information
// @Summary_mn        API-ийн мэдээллийг авах
// @Description    Get information about the API, such as its version and name.
// @Description_mn    API-ийн талаарх мэдээллийг, жишээ нь хувилбар ба нэрээ авах.
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

// @Summary        Check API connectivity and status
// @Summary_mn        API-ийн холболтыг шалгах ба статусыг авах
// @Description    Check the API connection and retrieve the status of the database, configuration, and network.
// @Description_mn    API-ийн холболтыг шалгах мөн тохиргооны мэдээлл, сүлжээний статусыг аваh, өгөгдөлийн сантай харьцахад саад буй эсэхийг шалгаж болно
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

// @Summary        Invoke a specified function
// @Summary_mn        Тодорхой функцыг дуудах
// @Description    Call a function with the specified name and parameters.
// @Description_mn    Тодорхой нэртэй болон параметртэй функцыг дуудах.
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

// @Summary        Submit a bill to the system
// @Summary_mn        Системд нэхэмжлэхийг оруулах
// @Description    Add a bill to the system with the provided data.
// @Description_mn    Өгөгдсөн мэдээллийг ашиглан системд нэхэмжлэхийг нэмэх.
// @Tags              ebarimt
// @Accept            json
// @Produce           json
// @Param             data body posapi.PutInput true "Bill data"
// @Success           200 {object} posapi.PutOutput
// @Router            /put [post]
func Put(c *fiber.Ctx) error {

	data := c.Body()
	result, err := ebarimt.PosAPI.Put(string(data))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to call Put function"})
	}

	return c.JSON(fiber.Map{"result": result})
}

// @Summary        Process a bill return
// @Summary_mn        Нэхэмжлэхийг буцаах
// @Description Return a bill with the provided data.
// @Description_mn Өгөгдсөн мэдээллийг ашиглан нэхэмжлэхийг буцаах.
// @Tags ebarimt
// @Accept json
// @Produce json
// @Param data body posapi.BillInput true "Return bill data"
// @Success 200 {object} posapi.BillOutput
// @Router /return [post]
func Return(c *fiber.Ctx) error {

	data := c.Body()
	result, err := ebarimt.PosAPI.ReturnBill(string(data))
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
