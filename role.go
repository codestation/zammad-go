package zammad

import (
	"fmt"
	"net/http"
	"time"
)

type Role struct {
	ID              int       `json:"id,omitempty"`
	Name            string    `json:"name"`
	Preferences     string    `json:"preferences,omitempty"`
	DefaultAtSignup bool      `json:"default_at_signup"`
	Active          bool      `json:"active"`
	Note            string    `json:"note,omitempty"`
	UpdatedByID     int       `json:"updated_by_id"`
	CreatedByID     int       `json:"created_by_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (c *Client) RoleListResult(opts ...Option) *Result[Role] {
	return &Result[Role]{
		res:     nil,
		resFunc: c.RoleListWithOptions,
		opts:    NewRequestOptions(opts...),
	}
}

func (c *Client) RoleList() ([]Role, error) {
	return c.RoleListResult().FetchAll()
}

func (c *Client) RoleListWithOptions(ro RequestOptions) ([]Role, error) {
	var roles []Role

	req, err := c.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.Url, "/api/v1/roles"), nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = ro.URLParams()

	if err = c.sendWithAuth(req, &roles); err != nil {
		return nil, err
	}

	return roles, nil
}

func (c *Client) RoleShow(roleID int) (Role, error) {
	var role Role

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/roles/%d", roleID)), nil)
	if err != nil {
		return role, err
	}

	if err = c.sendWithAuth(req, &role); err != nil {
		return role, err
	}

	return role, nil
}

func (c *Client) RoleCreate(r Role) (Role, error) {
	var role Role

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.Url, "/api/v1/roles"), r)
	if err != nil {
		return role, err
	}

	if err = c.sendWithAuth(req, &role); err != nil {
		return role, err
	}

	return role, nil
}

func (c *Client) RoleUpdate(RoleID int, r Role) (Role, error) {
	var role Role

	req, err := c.NewRequest("PUT", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/roles/%d", RoleID)), r)
	if err != nil {
		return role, err
	}

	if err = c.sendWithAuth(req, role); err != nil {
		return role, err
	}

	return role, nil
}

func (c *Client) RoleDelete(roleID int) error {
	req, err := c.NewRequest("DELETE", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/roles/%d", roleID)), nil)
	if err != nil {
		return err
	}

	if err = c.sendWithAuth(req, nil); err != nil {
		return err
	}

	return nil
}
