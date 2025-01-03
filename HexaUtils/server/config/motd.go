package config

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type MOTD struct {
	IconURL     string
	Message     string
	SampleTexts []string
	IconBase64  string
}

var httpClient = &http.Client{}

func NewMOTD() *MOTD {
	motd := &MOTD{
		IconURL:     "https://i.ibb.co/BTdWQMY/Hexa-Server-Logo.png",
		Message:     "HexaServer\nA minecraft server written in Go",
		SampleTexts: []string{"HexaServer", "A minecraft server", "written in Go"},
	}
	motd.LoadIconBase64()
	return motd
}

func (m *MOTD) SetIconURL(iconURL string) {
	m.IconURL = iconURL
}

func (m *MOTD) LoadIconBase64() string {
	resp, err := httpClient.Get(m.IconURL)
	if err != nil {
		return fmt.Sprintf("Error al obtener la URL: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		retryAfter := resp.Header.Get("Retry-After")
		if retryAfter != "" {
			waitTime, err := time.ParseDuration(retryAfter + "s")
			if err != nil {
				return fmt.Sprintf("Error al analizar Retry-After: %v", err)
			}
			time.Sleep(waitTime)
			return m.LoadIconBase64()
		}
		return fmt.Sprintf("Error en la respuesta del servidor: 429 - Demasiadas solicitudes")
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Sprintf("Error en la respuesta del servidor: %d", resp.StatusCode)
	}

	var buf bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	_, err = io.Copy(encoder, resp.Body)
	if err != nil {
		return fmt.Sprintf("Error al leer el cuerpo de la respuesta: %v", err)
	}
	encoder.Close()

	if buf.Len() == 0 {
		return "Error: No se descargó la imagen (data vacía)"
	}

	m.IconBase64 = buf.String()
	return m.IconBase64
}

func (m *MOTD) GetIconURL() string {
	return m.IconURL
}

func (m *MOTD) GetMessage() string {
	return m.Message
}

func (m *MOTD) GetSampleTexts() []string {
	return m.SampleTexts
}

func (m *MOTD) SetMessage(message string) {
	m.Message = message
}

func (m *MOTD) SetSampleTexts(sampleTexts []string) {
	m.SampleTexts = sampleTexts
}

func (m *MOTD) AddSampleText(sampleText string) {
	m.SampleTexts = append(m.SampleTexts, sampleText)
}

func (m *MOTD) RemoveSampleText(sampleText string) {
	for i, text := range m.SampleTexts {
		if text == sampleText {
			m.SampleTexts = append(m.SampleTexts[:i], m.SampleTexts[i+1:]...)
			break
		}
	}
}

func (m *MOTD) ClearSampleTexts() {
	m.SampleTexts = []string{}
}

func (m *MOTD) GetSampleTextsJSON() []map[string]string {
	sampleTexts := make([]map[string]string, len(m.SampleTexts))
	for i, text := range m.SampleTexts {
		sampleTexts[i] = map[string]string{
			"name": text,
			"id":   "4566e69f-c907-48ee-8d71-d7ba5aa00d20",
		}
	}
	return sampleTexts
}

func (m *MOTD) GetJSON() string {
	image := m.LoadIconBase64()
	loaded := "data:image/png;base64," + image
	jsonMap := map[string]interface{}{
		"version": map[string]interface{}{
			"name":     "HexaServer",
			"protocol": 767,
		},
		"players": map[string]interface{}{
			"max":    20,
			"online": 0,
			"sample": m.GetSampleTextsJSON(),
		},
		"description":        m.Message,
		"favicon":            loaded,
		"enforcesSecureChat": false,
	}
	return JSONToString(jsonMap)
}

func JSONToString(jsonMap map[string]interface{}) string {
	jsonBytes, err := json.Marshal(jsonMap)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}
