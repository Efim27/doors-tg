package requests

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func (clientAPI ClientAPI) DemoOpenDoor() (err error) {
	agent := fiber.AcquireAgent()
	request := agent.Request()
	request.Header.SetMethod(fiber.MethodPost)
	request.SetRequestURI("http://213.189.221.50/api/v1/demo/door/open")

	err = agent.Parse()
	if err != nil {
		return
	}

	code, _, errs := agent.Bytes()
	if code != fiber.StatusOK {
		return err
	}

	if len(errs) > 0 {
		log.Println(errs)
		err = errs[0]
		return
	}

	return
}
