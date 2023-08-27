package nessie

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func get[T any](path string, client *Client) (decodedBody T, err error) {
	resp, err := client.underlyingClient.Get(client.createURL(path))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return decodedBody, fmt.Errorf("unable to get, status: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&decodedBody)
	return
}

func post[T any](path string, input T, client *Client) error {
	b, err := json.Marshal(input)
	if err != nil {
		return err
	}

	resp, err := client.underlyingClient.Post(client.createURL(path), "application/json", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unable to post, status: %d", resp.StatusCode)
	}
	return nil
}

func put[T any](path string, input T, client *Client) error {
	b, err := json.Marshal(input)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", client.createURL(path), bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.underlyingClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("unable to put, status: %d", resp.StatusCode)
	}

	return nil
}

func delete(path string, client *Client) error {
	req, err := http.NewRequest("DELETE", client.createURL(path), nil)
	if err != nil {
		return err
	}
	resp, err := client.underlyingClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unable to delete, status: %d", resp.StatusCode)
	}

	return nil
}
