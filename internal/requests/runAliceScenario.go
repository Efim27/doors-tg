package requests

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

const (
	ScenarioRoomOpenNoAccessOrNoBatch = "265b9916-5b2d-4eaf-8104-a67980747f14"
	ScenarioRoomOpenSuccess           = "13997be5-042c-48b2-943d-05690e308b4a"
	ScenarioRoomOpenCancel            = "79f6f5a9-e57d-4eb6-8d57-45eb6becd6f0"
)

func (clientAPI ClientAPI) RunAliceScenario(scenarioID string) (err error) {
	agent := fiber.AcquireAgent()
	request := agent.Request()
	request.Header.SetMethod(fiber.MethodPost)

	fullURL := fmt.Sprintf("https://api.iot.yandex.net/v1.0/scenarios/%v/actions", scenarioID)
	request.SetRequestURI(fullURL)

	agent.Set("Authorization", fmt.Sprintf("Bearer %v", clientAPI.YaAuthToken))

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
