package requests

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type ResponseAdminListTG struct {
	Status string  `json:"status"`
	Result []int64 `json:"result"`
}

func (clientAPI ClientAPI) AdminListTG() (adminListTG []int64, err error) {
	agent := fiber.AcquireAgent()
	request := agent.Request()
	request.Header.SetMethod(fiber.MethodGet)

	url, err := JoinURL(clientAPI.APIAddr, "/api/v1/tg/user/admins")
	if err != nil {
		log.Fatalln(err)
	}

	request.SetRequestURI(url)
	err = agent.Parse()
	if err != nil {
		return
	}

	code, body, errs := agent.Bytes()
	if code != fiber.StatusOK {
		err = errors.New(fmt.Sprintf("Bad status code: %v", code))
		return
	}

	if len(errs) > 0 {
		log.Println(errs)
		err = errs[0]
		return
	}

	response := ResponseAdminListTG{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}

	adminListTG = response.Result
	return
}
