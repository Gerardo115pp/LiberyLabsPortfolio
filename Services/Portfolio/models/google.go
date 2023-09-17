package models

type RecaptchaVerifyRequest struct {
	Secret   string `json:"secret"`
	Response string `json:"response"`
	Remoteip string `json:"remoteip"`
}

type RecaptchaVerifyResponse struct {
	Success     bool     `json:"success"`
	ChallengeTs string   `json:"challenge_ts"`
	Hostname    string   `json:"hostname"`
	ErrorCodes  []string `json:"error-codes"`
}
