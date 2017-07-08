package main

//OutputJSON - The struct that convert the output to Json useable by Slack
type OutputJSON struct {
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
}
