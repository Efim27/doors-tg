package requests

type ClientAPI struct {
	APIAddr     string
	YaAuthToken string
}

func NewClientAPI(APIAddr, YaAuthToken string) (clientAPI *ClientAPI) {
	clientAPI = new(ClientAPI)
	clientAPI.APIAddr = APIAddr
	clientAPI.YaAuthToken = YaAuthToken

	return
}
