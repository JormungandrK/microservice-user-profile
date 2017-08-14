// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "user-profile": userProfile Resource Client
//
// Command:
// $ goagen
// --design=github.com/JormungandrK/microservice-user-profile/design
// --out=$(GOPATH)/src/github.com/JormungandrK/microservice-user-profile
// --version=v1.2.0-dirty

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// GetUserProfileUserProfilePath computes a request path to the GetUserProfile action of userProfile.
func GetUserProfileUserProfilePath(userID string) string {
	param0 := userID

	return fmt.Sprintf("/user-profile/%s", param0)
}

// Get a UserProfile by UserID
func (c *Client) GetUserProfileUserProfile(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetUserProfileUserProfileRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetUserProfileUserProfileRequest create the request corresponding to the GetUserProfile action endpoint of the userProfile resource.
func (c *Client) NewGetUserProfileUserProfileRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}