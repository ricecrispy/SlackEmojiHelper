package main

type OutputJson struct {
	Output string
}

type SlackPostRequest struct {
	Token string `json:"token"`
	Team_id string `json:"team_id"`
	Team_domain string `json:"team_domain"`
	Channel_id string `json:"channel_id"`
	Channel_name string `json:"channel_name"`
	User_id string `json:"user_id"`
	User_name string `json:"user_name"`
	Command string `json:"command"`
	Text string `json:"text"`
	Response_url string `json:"response_url"`
}
