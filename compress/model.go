package compress

type BaseModel struct {
	//client_id: `${this.customer_name}_client`,
	ClientId   	string `json:"client_id"`
	ApiKey 		string `json:"api_key"`
}

type ResponseServer struct {
	Message  string      `json:"message"`
	Response string      `json:"response"`
	Data     interface{} `json:"data"`
}
