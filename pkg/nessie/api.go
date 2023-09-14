package nessie

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func get[T any](path string, client *Client) (T, error) {
	url, err := client.createURL(path, nil)
	if err != nil {
		var t T
		return t, err
	}
	return underlyingGet[T](url, client)
}

func getWithQueryParams[T any](path string, params map[string]string, client *Client) (T, error) {
	url, err := client.createURL(path, &params)
	if err != nil {
		var t T
		return t, err
	}
	return underlyingGet[T](url, client)
}

func underlyingGet[T any](url string, client *Client) (decodedBody T, err error) {
	resp, err := client.underlyingClient.Get(url)
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

	url, err := client.createURL(path, nil)
	if err != nil {
		return err
	}

	resp, err := client.underlyingClient.Post(url, "application/json", bytes.NewBuffer(b))
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

	url, err := client.createURL(path, nil)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(b))
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
	url, err := client.createURL(path, nil)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("DELETE", url, nil)
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
