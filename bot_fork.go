package telego

import (
	"context"
	"fmt"
	"reflect"

	"github.com/chococola/telego/internal/json"
	ta "github.com/chococola/telego/telegoapi"
)

// Split original performRequest to two methods, useful if you prepare message in one period of time, and make send delayed
// - PrepareRawRequest - ta.RequestData.Buffer - can be saved to persistent storage
// - PerformRawRequest - all you need to send message

// PrepareRawRequest - part of original method performRequest and constructAndCallRequest.
//
//	Prepare request doesn't matter json or multipart
func (b *Bot) PrepareRawRequest(parameters any) (*ta.RequestData, error) {
	filesParams, hasFiles := filesParameters(parameters)
	var data *ta.RequestData

	if hasFiles {
		parsedParameters, err := parseParameters(parameters)
		if err != nil {
			return nil, fmt.Errorf("parsing parameters: %w", err)
		}

		data, err = b.constructor.MultipartRequest(parsedParameters, filesParams)
		if err != nil {
			return nil, fmt.Errorf("multipart request: %w", err)
		}

	} else {
		var err error
		data, err = b.constructor.JSONRequest(parameters)
		if err != nil {
			return nil, fmt.Errorf("json request: %w", err)
		}
	}

	return data, nil
}

// PerformRawRequest - part of performRequest method
func (b *Bot) PerformRawRequest(ctx context.Context, methodName string, data *ta.RequestData, vs ...any) error {
	resp, err := b.callRawRequest(ctx, methodName, data)
	if err != nil {
		b.log.Errorf("Execution error %s: %s", methodName, err)
		return fmt.Errorf("internal execution: %w", err)
	}
	b.log.Debugf("API response %s: %s", methodName, resp.String())

	if !resp.Ok {
		return fmt.Errorf("api: %w", resp.Error)
	}

	if resp.Result != nil {
		var unmarshalErr error
		for i := range vs {
			unmarshalErr = json.Unmarshal(resp.Result, &vs[i])
			if unmarshalErr == nil {
				break
			}
		}

		if unmarshalErr != nil {
			return fmt.Errorf("unmarshal to %s: %w", reflect.TypeOf(vs[len(vs)-1]), unmarshalErr)
		}
	}

	if b.reportWarningAsErrors && resp.Error != nil {
		return resp.Error
	}

	return nil
}

// part of constructAndCallRequest
func (b *Bot) callRawRequest(ctx context.Context, methodName string, data *ta.RequestData) (*ta.Response, error) {
	var url string
	if b.useTestServerPath {
		url = b.apiURL + botPathPrefix + b.token + "/test/" + methodName
	} else {
		url = b.apiURL + botPathPrefix + b.token + "/" + methodName
	}

	resp, err := b.api.Call(ctx, url, data)
	if err != nil {
		return nil, fmt.Errorf("request call: %w", err)
	}

	return resp, nil
}
