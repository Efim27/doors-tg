package requests

type ClientAPI struct {
	APIAddr string
}

func NewClientAPI(APIAddr string) (clientAPI *ClientAPI) {
	clientAPI = new(ClientAPI)
	clientAPI.APIAddr = APIAddr

	return
}