package vend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	baseUrlString  = "vendhq.com/api/"
	libraryVersion = "0.1"
	userAgent      = "go-vend/" + libraryVersion
)

var (
	store    string
	username string
	password string
)

// Vend client which talks to the Vend API
type Client struct {

	// HTTP client
	client    *http.Client
	UserAgent string

	Config    *ConfigService
	Products  *ProductService
	Taxes     *TaxService
	Users     *UserService
	Outlets   *OutletService
	Registers *RegisterService
	Customers *CustomerService
}

func NewClient(st, usrname, passwd string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	store = st
	username = usrname
	password = passwd

	c := &Client{client: httpClient, UserAgent: userAgent}
	c.Config = &ConfigService{client: c}
	c.Products = &ProductService{client: c}
	c.Taxes = &TaxService{client: c}
	c.Users = &UserService{client: c}
	c.Outlets = &OutletService{client: c}
	c.Registers = &RegisterService{client: c}
	c.Customers = &CustomerService{client: c}

	return c
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {

	u := fmt.Sprintf("https://%s.vendhq.com/api/%s", store, urlStr)
	fmt.Println(u)
	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u, buf)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(username, password)

	req.Header.Add("User-Agent", c.UserAgent)

	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := newResponse(resp)

	err = CheckResponse(resp)
	if err != nil {
		return response, err
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}
	return response, err
}

type Response struct {
	*http.Response
}

func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}

/*
{"error":"API auth required"}
*/

type ErrorResponse struct {
	Response *http.Response
	Message  string `json:"error,omitempty"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method,
		r.Response.Request.URL,
		r.Response.StatusCode,
		r.Message)
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}
	return errorResponse
}
