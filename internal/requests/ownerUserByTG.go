package requests

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

var ErrNoAccess = errors.New("user not found or no access")

type User struct {
	Id         uint32 `json:"id" form:"id"`
	Name       string `json:"name" form:"name"`
	Surname    string `json:"surname" form:"surname"`
	Patronymic string `json:"patronymic" form:"patronymic"`
}

type OwnerUserByTGResponse struct {
	Status string `json:"status"`
	Result User   `json:"result"`
}

func (clientAPI ClientAPI) OwnerUserByTG(TgID int64) (owner User, err error) {
	agent := fiber.AcquireAgent()
	request := agent.Request()
	request.Header.SetMethod(fiber.MethodGet)

	url, err := JoinURL(clientAPI.APIAddr, fmt.Sprintf("/api/v1/tg/user/owner/%v", TgID))
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
