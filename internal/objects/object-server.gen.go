// Package objects provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package objects

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// FindObjects request
	FindObjects(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AddNewObject request
	AddNewObject(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteObjectByID request
	DeleteObjectByID(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindObjectByID request
	FindObjectByID(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) FindObjects(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewFindObjectsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AddNewObject(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAddNewObjectRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteObjectByID(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteObjectByIDRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) FindObjectByID(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewFindObjectByIDRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewFindObjectsRequest generates requests for FindObjects
func NewFindObjectsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/objects")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewAddNewObjectRequest generates requests for AddNewObject
func NewAddNewObjectRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/objects")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewDeleteObjectByIDRequest generates requests for DeleteObjectByID
func NewDeleteObjectByIDRequest(server string, id int64) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/objects/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewFindObjectByIDRequest generates requests for FindObjectByID
func NewFindObjectByIDRequest(server string, id int64) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/objects/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// FindObjects request
	FindObjectsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*FindObjectsResponse, error)

	// AddNewObject request
	AddNewObjectWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*AddNewObjectResponse, error)

	// DeleteObjectByID request
	DeleteObjectByIDWithResponse(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*DeleteObjectByIDResponse, error)

	// FindObjectByID request
	FindObjectByIDWithResponse(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*FindObjectByIDResponse, error)
}

type FindObjectsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Object
}

// Status returns HTTPResponse.Status
func (r FindObjectsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r FindObjectsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AddNewObjectResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *[]Object
}

// Status returns HTTPResponse.Status
func (r AddNewObjectResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AddNewObjectResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteObjectByIDResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r DeleteObjectByIDResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteObjectByIDResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type FindObjectByIDResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Object
}

// Status returns HTTPResponse.Status
func (r FindObjectByIDResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r FindObjectByIDResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// FindObjectsWithResponse request returning *FindObjectsResponse
func (c *ClientWithResponses) FindObjectsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*FindObjectsResponse, error) {
	rsp, err := c.FindObjects(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseFindObjectsResponse(rsp)
}

// AddNewObjectWithResponse request returning *AddNewObjectResponse
func (c *ClientWithResponses) AddNewObjectWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*AddNewObjectResponse, error) {
	rsp, err := c.AddNewObject(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAddNewObjectResponse(rsp)
}

// DeleteObjectByIDWithResponse request returning *DeleteObjectByIDResponse
func (c *ClientWithResponses) DeleteObjectByIDWithResponse(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*DeleteObjectByIDResponse, error) {
	rsp, err := c.DeleteObjectByID(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteObjectByIDResponse(rsp)
}

// FindObjectByIDWithResponse request returning *FindObjectByIDResponse
func (c *ClientWithResponses) FindObjectByIDWithResponse(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*FindObjectByIDResponse, error) {
	rsp, err := c.FindObjectByID(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseFindObjectByIDResponse(rsp)
}

// ParseFindObjectsResponse parses an HTTP response from a FindObjectsWithResponse call
func ParseFindObjectsResponse(rsp *http.Response) (*FindObjectsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &FindObjectsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Object
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseAddNewObjectResponse parses an HTTP response from a AddNewObjectWithResponse call
func ParseAddNewObjectResponse(rsp *http.Response) (*AddNewObjectResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AddNewObjectResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest []Object
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}

// ParseDeleteObjectByIDResponse parses an HTTP response from a DeleteObjectByIDWithResponse call
func ParseDeleteObjectByIDResponse(rsp *http.Response) (*DeleteObjectByIDResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteObjectByIDResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseFindObjectByIDResponse parses an HTTP response from a FindObjectByIDWithResponse call
func ParseFindObjectByIDResponse(rsp *http.Response) (*FindObjectByIDResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &FindObjectByIDResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Object
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Return all objects
	// (GET /objects)
	FindObjects(ctx echo.Context) error
	// Add a new object
	// (POST /objects)
	AddNewObject(ctx echo.Context) error
	// Delete object
	// (DELETE /objects/{id})
	DeleteObjectByID(ctx echo.Context, id int64) error
	// Returns an object by ID
	// (GET /objects/{id})
	FindObjectByID(ctx echo.Context, id int64) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// FindObjects converts echo context to params.
func (w *ServerInterfaceWrapper) FindObjects(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindObjects(ctx)
	return err
}

// AddNewObject converts echo context to params.
func (w *ServerInterfaceWrapper) AddNewObject(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddNewObject(ctx)
	return err
}

// DeleteObjectByID converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteObjectByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteObjectByID(ctx, id)
	return err
}

// FindObjectByID converts echo context to params.
func (w *ServerInterfaceWrapper) FindObjectByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindObjectByID(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/objects", wrapper.FindObjects)
	router.POST(baseURL+"/objects", wrapper.AddNewObject)
	router.DELETE(baseURL+"/objects/:id", wrapper.DeleteObjectByID)
	router.GET(baseURL+"/objects/:id", wrapper.FindObjectByID)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8xUXW/TPBT+K0fnfaWBFNoOJi5yV1RAldCKEFxN0+TFJ80Zie3ZTqtS5b8jO+l3yhg3",
	"cJfY59jPeT68xkxXRitS3mG6RpcVVIn4Obt/oMyHL1GWsxzTmzUaqw1ZzxQrMkvCk7wTsUqSyywbz1ph",
	"2nVfOIhFrBVI4QkTzLWtQgOG/1eeq7DoV4YwRectqzk2Cebf79q1s+eGbZhOds2sPM3Jhm6Wv2isFT/W",
	"XesWDCv/9qr3rLBytxBl/SQYJfpmaRK09FizJYnpTYC2my7Zp/C2uW1CNatcn171tWAH489TEGWplw4s",
	"CclqDkJJWFr24fvj7P2n8fUEdITlAhb2ZQBztAMVZ1Y7sgvOAooFWddeczkYDUZhbG1ICcOY4pvB5WCE",
	"CRrhiyj7cHN+usY59Wj/hbxlWpADXxCwBJ0H2Hu4gouiK6YSU/zASs62e5ac0cq1Fns9GkWnaeVJtWY0",
	"puQsNg8fXLhv49rwxZ6q2Pi/pRxT/G+48/ewM/ewc3az1UpYK1YY2e8T2EEhFgT3RApsN5uMyrq6qoRd",
	"tTPXVh2M2SRotOvhZywlCFC07ErhBcsEdlZIop9envA0lvKalh36E6Iu/z5RUAjX0tQNc0TS8eBxe2On",
	"4Zpl05JVku9J2ySug1Ab2u6FIwlagQDHal52oT4kre1q8b1bxQIjrKjIk3XxUTu8ZToJdu1u8Bpy8lmB",
	"IZaYxgxggjHpaZvlXbi9rSnZo/jJx6W5PRHx6tzYMqhw1Yp8uP9NidoX2vKPTVHPIdMJKO0h17U6FqWj",
	"VW+1Ppfp2ir3XPZ30f73uX/eS/M7uTmXk4vwfrd3/4liPVqsAvlN0zQ/AwAA//8K1xqd0QcAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
