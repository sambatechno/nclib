package nclib

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
)

type AddActivityPayload struct {
	ActivityName   string         `json:"activity_name"`
	AssetId        string         `json:"asset_id"`
	Timestamp      string         `json:"timestamp"`
	Identity       string         `json:"identity"`
	ActivitySource string         `json:"activity_source"`
	ActivityParams map[string]any `json:"activity_params"`
}

var addActivityEndpoint = map[IDC]string{
	US: "https://api2.netcoresmartech.com/v1/activity/upload",
	EU: "https://apieu2.netcoresmartech.com/v1/activity/upload",
	IN: "https://apiin2.netcoresmartech.com/v1/activity/upload",
}

func (c *client) AddActivity(ctx context.Context, payload ...AddActivityPayload) error {
	res, err := c.client.R().
		SetContext(ctx).
		SetBody(payload).
		Post(addActivityEndpoint[c.idc])
	if err != nil {
		return err
	}

	defer res.Body.Close()
	resJson, _ := io.ReadAll(res.Body)

	var resData map[string]string
	json.Unmarshal(resJson, &resData)

	if res.StatusCode() != 200 {
		return fmt.Errorf("failed to add activity status %d, response: %s", res.StatusCode(), resData["error"])
	}

	return nil
}
