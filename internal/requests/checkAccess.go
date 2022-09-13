package requests

import (
	"encoding/json"
	"fmt"
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

	err = agent.Parse()
	if err != nil {
		return
	}

	code, body, errs := agent.Bytes()
	if (code != fiber.StatusOK) && (code != fiber.StatusUnauthorized) {
		err = errors.New(fmt.Sprintf("Bad status code: %v", code))
		return
	}
	if code == fiber.StatusUnauthorized {
		err = ErrNoAccess
		return
	}

	if len(errs) > 0 {
		log.Println(errs)
		err = errs[0]
		return
	}

	ownerUserByTGResponse := OwnerUserByTGResponse{}
	err = json.Unmarshal(body, &ownerUserByTGResponse)
	if err != nil {
		return
	}

	owner = ownerUserByTGResponse.Result
	return
}
