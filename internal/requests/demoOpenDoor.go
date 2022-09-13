package requests

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
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
		return errors.New(fmt.Sprintf("Bad status code: %v", code))
	}

	if len(errs) > 0 {
		log.Println(errs)
		err = errs[0]
		return
	}

	return
}
