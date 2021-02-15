package bark

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// New returns a Device representing the given registered device.
func New(endpointURL, deviceKey string) *Device {
	if _, err := url.ParseRequestURI(endpointURL); err != nil {
		panic(err)
	}
	if isBlank(deviceKey) {
		panic(errors.New("blank device key"))
	}

	return &Device{
		endpointURL: strings.TrimRight(endpointURL, "/"),
		deviceKey:   deviceKey,
	}
}

// Ping sends a empty message to the device.
func (d *Device) Ping(opts ...Options) error {
	return d.sendData(nil, opts...)
}

// SendShortMessage sends a message contains only body content to the device.
func (d *Device) SendShortMessage(content string, opts ...Options) error {
	return d.sendData(url.Values{
		"body": {content},
	}, opts...)
}

// SendMessage sends a message with title and body content to the device.
func (d *Device) SendMessage(title, body string, opts ...Options) error {
	return d.sendData(url.Values{
		"title": {title},
		"body":  {body},
	}, opts...)
}

var (
	client = http.Client{
		Timeout: 15 * time.Second,
	}
)

// Copied from https://github.com/Finb/bark-server/blob/master/server.go
type serverResponse struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func (d *Device) sendData(data url.Values, opts ...Options) error {
	var payload io.Reader
	if len(data) > 0 {
		payload = strings.NewReader(data.Encode())
	}
	urlStr := fmt.Sprintf("%s/%s/", d.endpointURL, d.deviceKey)

	// build request
	req, err := http.NewRequest("POST", urlStr, payload)
	if err != nil {
		return err
	}

	// set header for form data
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// set url parameters as per options
	query := req.URL.Query()
	for _, o := range opts {
		if rt := string(o.Ringtone); isNotBlank(rt) {
			query.Set("sound", rt)
		}
		if isNotBlank(o.OpenURL) {
			query.Set("url", o.OpenURL)
		}
		if isNotBlank(o.CopyText) {
			query.Set("copy", o.CopyText)
		}
		if o.ForceArchive {
			query.Set("isArchive", "1")
		}
		if o.ForceCopy {
			query.Set("automaticallyCopy", "1")
		}
	}
	req.URL.RawQuery = query.Encode()

	// send request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// verify response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if !(200 <= resp.StatusCode && resp.StatusCode <= 299) {
		return fmt.Errorf("http error status, code: %d, body: %s", resp.StatusCode, body)
	}

	var respData serverResponse
	if err := json.Unmarshal(body, &respData); err != nil {
		return fmt.Errorf("fail to unmarshal server response, error: %w, body: %s", err, body)
	}

	if respData.Code != 200 {
		return fmt.Errorf("server error, code: %d, message: %s", respData.Code, respData.Message)
	}
	return nil
}

func isBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}

func isNotBlank(s string) bool {
	return strings.TrimSpace(s) != ""
}
