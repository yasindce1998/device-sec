package websocket

import (
    "github.com/gorilla/websocket"
    "time"
)

type Client struct {
    conn     *websocket.Conn
    serverURL string
    agentID   string
}

func NewClient(serverURL, agentID string) (*Client, error) {
    return &Client{
        serverURL: serverURL,
        agentID:   agentID,
    }, nil
}

func (c *Client) Connect() error {
    dialer := websocket.Dialer{
        HandshakeTimeout: 10 * time.Second,
    }

    conn, _, err := dialer.Dial(c.serverURL + "?agent_id=" + c.agentID, nil)
    if err != nil {
        return err
    }

    c.conn = conn
    return nil
}