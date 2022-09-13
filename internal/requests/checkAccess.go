package requests

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func (clientAPI ClientAPI) CheckAccess(firstName string, lastName string) (err error) {
	agent := fiber.AcquireAgent()
	request := agent.Request()
	request.Header.SetMethod(fiber.MethodPost)

	url, err := JoinURL(clientAPI.APIAddr, "/api/v1/demo/access")
	if err != nil {
		log.Fatalln(err)
	}

	request.SetRequestURI(url)

	agent.JSON(fiber.Map{
		"first_name":  firstName,
		"last_name":   lastName,
		"object_name": "main",
	})
	err = agent.Parse()
	if err != nil {
		return
	}

	code, _, errs := agent.Bytes()
	if code != fiber.StatusOK {
		return ErrNoAccess
	}

	if len(errs) > 0 {
		log.Println(errs)
		err = errs[0]
		return
	}

	return
}
