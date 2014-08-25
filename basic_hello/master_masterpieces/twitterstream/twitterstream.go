// twitterstream.go file, the streaming go file for learning purpose.

// ATTN: THIS FILE IS NOT COMPLETE

// ORIGINAL CONTENT FROM: https://github.com/hoisie/twitterstream/blob/master/twitterstream.go

package twitterstream

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"crypto/tls"
	"encoding/base64"
	"http"
	"io"
	"json"
	"net"
	"os"
	"strconv"
	"time"
	"url"
)

var followUrl, _ = url.Parse("https://stream.twitter.com/1/statuses/filter.json")
var trackUrl, _ = url.Parse("http://stream.twitter.com/1/statuses/filter.json")
var sampleUrl, _ = url.Parse("http://stream.twitter.com/1/statuses/sample.json")
var userUrl, _ = url.Parse("http://userstream.twitter.com/2/user.json")
var siteStreamUrl, _ = url.Parse("https://betastream.twitter.com/2b/site.json")

var retryTimeout int64 = 5e9

type streamConn struct {
	clientConn   *http.ClientConn
	url          *url.URL
	stream       chan *Tweet
	eventStream  chan *Event
	friendStream chan *FriendList
	authData     string
	postData     string
	stale        bool
}

func (conn *streamConn) Close() {
	// Just mark the connection as stale, and let the connet() handle close after a read
	conn.stale = true
}

func (conn *streamConn) connect() (*http.Response, os.Error) {
	if conn.stale {
		return nil, os.NewError("stale connection")
	}
	var tcpConn net.Conn
	var err os.Error
	if proxy := os.Getenv("HTTP_PROXY"); len(proxy) > 0 {
		proxy_url, _ := url.Parse(proxy)
		tcpConn, err = net.Dial("tcp", proxy_url.Host)
	}
}
