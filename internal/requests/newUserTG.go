package requests

import (
	"errors"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func (clientAPI ClientAPI) NewUserTG(TgID, ChatID int64, Username, FirstName, LastName, LanguageCode string) (err error) {
	agent := fiber.AcquireAgent()
	request := agent.Request()
	request.Header.SetMethod(fiber.MethodPost)

	url, err := JoinURL(clientAPI.APIAddr, "/api/v1/tg/user")
	if err != nil {
		log.Fatalln(err)
	}

	request.SetRequestURI(url)

	agent.JSON(fiber.Map{
		"tg_id":         TgID,
		"chat_id":       ChatID,
		"username":      Username,
		"first_name":    FirstName,
		"last_name":     LastName,
		"language_code": LanguageCode,
	})
	err = agent.Parse()
	if err != nil {
		return
	}

	code, _, errs := agent.Bytes()
	if (code != fiber.StatusOK) && (code != fiber.StatusConflict) {
		err = errors.New(fmt.Sprintf("Bad status code: %v", code))
		return
	}

	if len(errs) > 0 {
		log.Println(errs)
		err = errs[0]
		return
	}

	return
}
