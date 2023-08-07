package model

type VoiceCallData struct {
	Country             string
	Bandwidth           string
	ResponseTime        string
	Provider            string
	ConnectionStability float32
	TTFB                int
	VoicePurity         int
	MedianOfCallsTime   int
}
type ChVoice struct {
	Data []VoiceCallData
	Err  error
}
