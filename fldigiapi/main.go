package fldigiapi

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/http"
)

// Client represents an XML-RPC client.
type Client struct {
	endpoint string
	client   *http.Client
}

// NewClient creates a new XML-RPC client with the specified endpoint.
func NewClient(endpoint string) *Client {
	return &Client{
		endpoint: endpoint,
		client:   &http.Client{},
	}
}

// Call executes an XML-RPC method call with the given method name and parameters.
func (c *Client) Call(method string, args interface{}) (interface{}, error) {
	// Prepare XML-RPC request body
	requestBody, err := prepareRequestBody(method, args)
	if err != nil {
		return nil, err
	}

	// Make HTTP POST request to the endpoint
	resp, err := c.client.Post(c.endpoint, "text/xml", bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response body
	var buf bytes.Buffer
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}

	// Debug: Print XML response
	fmt.Println("XML Response:", buf.String())

	// Extract and return the value from the XML response
	return extractValueFromXML(buf.String())
}

// Function to extract the value from the XML response
func extractValueFromXML(xmlStr string) (string, error) {
	// Define a struct to hold the XML structure
	type MethodResponse struct {
		Params struct {
			Param struct {
				Value string `xml:"value"`
			} `xml:"param"`
		} `xml:"params"`
	}

	// Unmarshal the XML response into the struct
	var methodResp MethodResponse
	if err := xml.Unmarshal([]byte(xmlStr), &methodResp); err != nil {
		return "", err
	}

	// Extract and return the value
	return methodResp.Params.Param.Value, nil
}

// param type definition moved outside of prepareRequestBody function
type param struct {
	Value interface{} `xml:"value"`
}

// prepareRequestBody prepares the XML-RPC request body.
func prepareRequestBody(method string, args interface{}) ([]byte, error) {
	// XML-RPC request structure
	type methodCall struct {
		XMLName xml.Name
		Method  string  `xml:"methodName"`
		Params  []param `xml:"params>param"`
	}

	call := methodCall{
		XMLName: xml.Name{Local: "methodCall"},
		Method:  method,
	}

	// Prepare parameters
	if args != nil {
		switch a := args.(type) {
		case []interface{}:
			for _, arg := range a {
				call.Params = append(call.Params, param{Value: arg})
			}
		default:
			call.Params = append(call.Params, param{Value: args})
		}
	}

	// Encode XML-RPC request body
	buf, err := xml.MarshalIndent(call, "", "  ")
	if err != nil {
		return nil, err
	}
	return buf, nil
}
