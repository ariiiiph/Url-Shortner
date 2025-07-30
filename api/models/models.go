package models

import "time"

type Request struct {
	URL         string        `jason:"url"`
	CustomShort string        `jason:"short"`
	Expiry      time.Duration `jason:"expiry"`
}

type Response struct {
	URL             string        `jason:"url"`
	CustomShort     string        `jason:"short"`
	Expiry          time.Duration `jason:"expiry"`
	XRateRemaining  int           `jason:"rate_limit"`
	XRateLimitReset time.Duration `jason:"rate_limit_reset"`
}

type TagRequest struct {
	ShortID string `json:"shortID"`
	Tag     string `json:"tag"`
}
