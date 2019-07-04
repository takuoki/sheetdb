package sheetdb

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/takuoki/gsheets"
)

var modelSets = map[string][]model{}

type Client struct {
	gsClient      *gsheets.Client
	spreadsheetID string
	modelSetName  string
	warningFunc   func([]gsheets.UpdateValue, interface{})

	models []model
}

type model struct {
	name     string
	loadFunc func(data *gsheets.Sheet) error
}

type ClientOption func(client *Client) *Client

func ModelSetName(modelSetName string) func(client *Client) *Client {
	return func(client *Client) *Client {
		if client != nil {
			client.modelSetName = modelSetName
		}
		return client
	}
}

func WarningFunc(f func([]gsheets.UpdateValue, interface{})) func(client *Client) *Client {
	return func(client *Client) *Client {
		if client != nil {
			client.warningFunc = f
		}
		return client
	}
}

func defaultWarningFunc(data []gsheets.UpdateValue, e interface{}) {
	log.Println(e)
	log.Printf("Data is not reflected in the sheet (data=%+v)", data)
}

func New(ctx context.Context, credentials, token, spreadsheetID string, opts ...ClientOption) (*Client, error) {
	gsClient, err := gsheets.New(ctx, credentials, token, gsheets.ClientWritable())
	if err != nil {
		return nil, fmt.Errorf("Unable to create gsheets client: %v", err)
	}
	client := &Client{
		gsClient:      gsClient,
		spreadsheetID: spreadsheetID,
		modelSetName:  "default",
		warningFunc:   defaultWarningFunc,
	}
	for _, opt := range opts {
		client = opt(client)
	}
	client.models = modelSets[client.modelSetName]
	return client, nil
}

func SetModel(modelSetName, modelName string, loadFunc func(data *gsheets.Sheet) error) {
	m := model{name: modelName, loadFunc: loadFunc}
	if s, ok := modelSets[modelSetName]; ok {
		s = append(s, m)
	} else {
		s = []model{m}
	}
}

func (c *Client) LoadData(ctx context.Context) error {

	if c.gsClient == nil {
		return errors.New("This package is not initialized")
	}

	for _, m := range c.models {
		data, err := c.gsClient.GetSheet(ctx, c.spreadsheetID, m.name)
		if err != nil {
			return err
		}
		err = m.loadFunc(data)
		if err != nil {
			return fmt.Errorf("Unable to load '%s' data: %v", m.name, err)
		}
	}

	return nil
}

func (c *Client) AsyncUpdate(data []gsheets.UpdateValue) error {

	if c.gsClient == nil {
		return errors.New("This package is not initialized")
	}

	go func() {
		defer func() {
			if e := recover(); e != nil {
				c.warningFunc(data, e)
			}
		}()

		if err := c.gsClient.BatchUpdate(context.Background(), c.spreadsheetID, data...); err != nil {
			panic(fmt.Sprintf("Unable to update spreadsheet: %v", err))
		}
	}()

	return nil
}
