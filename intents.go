package dialogflow

// https://dialogflow.com/docs/reference/agent/intents

import (
	"encoding/json"
	"fmt"
)

// Get all intents.
func (c *Client) AllIntents() (result []Intent, err error) {
	var bytes []byte
	if bytes, err = c.httpGet("intents", nil, nil); err == nil {
		if err = json.Unmarshal(bytes, &result); err == nil {
			return result, nil
		}
	}

	return []Intent{}, err
}

// Get an intent.
func (c *Client) Intent(iid string) (result IntentObject, err error) {
	var bytes []byte
	if bytes, err = c.httpGet(fmt.Sprintf("intents/%s", iid), nil, nil); err == nil {
		if err = json.Unmarshal(bytes, &result); err == nil {
			return result, nil
		}
	}

	return IntentObject{}, err
}

// Create a new intent.
//
// (do not fill Id in IntentObject)
func (c *Client) CreateIntent(intent IntentObject) (result ApiResponse, err error) {
	var bytes []byte
	if bytes, err = c.httpPost("intents", nil, nil, intent); err == nil {
		if err = json.Unmarshal(bytes, &result); err == nil {
			return result, nil
		}
	}

	return ApiResponse{}, err
}

// Update an intent.
func (c *Client) UpdateIntent(iid string, intent IntentObject) (result ApiResponse, err error) {
	var bytes []byte
	if bytes, err = c.httpPut(fmt.Sprintf("intents/%s", iid), nil, intent); err == nil {
		if err = json.Unmarshal(bytes, &result); err == nil {
			return result, nil
		}
	}

	return ApiResponse{}, err
}

// Delete an intent.
func (c *Client) DeleteIntent(iid string) (result ApiResponse, err error) {
	var bytes []byte
	if bytes, err = c.httpDelete(fmt.Sprintf("intents/%s", iid), nil, nil, nil); err == nil {
		if err = json.Unmarshal(bytes, &result); err == nil {
			return result, nil
		}
	}

	return ApiResponse{}, err
}
