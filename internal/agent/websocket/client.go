// websocket/client.go
package websocket

import (
    "fmt"
    "github.com/gorilla/websocket"
)

type Client struct {
    baseURL string
    agentID string
    conn    *websocket.Conn
}

func NewClient(baseURL, agentID string) (*Client, error) {
    if baseURL == "" || agentID == "" {
        return nil, fmt.Errorf("baseURL and agentID cannot be empty")
    }

    return &Client{
        baseURL: baseURL,
        agentID: agentID,
    }, nil
}

func (c *Client) Connect() error {
    // Create WebSocket URL with agent ID
    wsURL := fmt.Sprintf("%s/ws/%s", c.baseURL, c.agentID)

    // Connect to WebSocket server
    conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
    if err != nil {
        return fmt.Errorf("failed to connect to WebSocket server: %v", err)
    }

    c.conn = conn
    return nil
}

// ReadMessage reads a message from the WebSocket connection
func (c *Client) ReadMessage() (int, []byte, error) {
    if c.conn == nil {
        return 0, nil, fmt.Errorf("connection is not established")
    }
    return c.conn.ReadMessage()
}

// WriteMessage sends a message through the WebSocket connection
func (c *Client) WriteMessage(messageType int, data []byte) error {
    if c.conn == nil {
        return fmt.Errorf("connection is not established")
    }
    return c.conn.WriteMessage(messageType, data)
}

// Close closes the WebSocket connection
func (c *Client) Close() error {
    if c.conn == nil {
        return nil
    }
    return c.conn.Close()
}