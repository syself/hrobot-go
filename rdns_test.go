package client_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"

	client "github.com/syself/hrobot-go"
	"github.com/syself/hrobot-go/models"
	. "gopkg.in/check.v1"
)

func (s *ClientSuite) TestRDnsGetListSuccess(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		pwd, pwdErr := os.Getwd()
		c.Assert(pwdErr, IsNil)

		data, readErr := ioutil.ReadFile(fmt.Sprintf("%s/test/response/rdns_list.json", pwd))
		c.Assert(readErr, IsNil)

		_, err := w.Write(data)
		c.Assert(err, IsNil)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	rdnsList, err := robotClient.RDnsGetList()
	c.Assert(err, IsNil)
	c.Assert(len(rdnsList), Equals, 2)
	c.Assert(rdnsList[0].IP, Equals, testIP)
	c.Assert(rdnsList[1].IP, Equals, testIP2)
}

func (s *ClientSuite) TestRDnsGetListInvalidResponse(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, err := w.Write([]byte("invalid JSON"))
		c.Assert(err, IsNil)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	_, err := robotClient.RDnsGetList()
	c.Assert(err, Not(IsNil))
}

func (s *ClientSuite) TestRDnsGetListServerError(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	_, err := robotClient.RDnsGetList()
	c.Assert(err, Not(IsNil))
}

func (s *ClientSuite) TestRDnsGetSuccess(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		pwd, pwdErr := os.Getwd()
		c.Assert(pwdErr, IsNil)

		data, readErr := ioutil.ReadFile(fmt.Sprintf("%s/test/response/rdns_get.json", pwd))
		c.Assert(readErr, IsNil)

		_, err := w.Write(data)
		c.Assert(err, IsNil)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	rdns, err := robotClient.RDnsGet(testIP)
	c.Assert(err, IsNil)
	c.Assert(rdns.IP, Equals, testIP)
}

func (s *ClientSuite) TestRDnsGetInvalidResponse(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, err := w.Write([]byte("invalid JSON"))
		c.Assert(err, IsNil)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	_, err := robotClient.RDnsGet(testIP)
	c.Assert(err, Not(IsNil))
}

func (s *ClientSuite) TestRDnsGetServerError(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	_, err := robotClient.RDnsGet(testIP)
	c.Assert(err, Not(IsNil))
}

func (s *ClientSuite) TestRDnsSetSuccess(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqContentType := r.Header.Get("Content-Type")
		c.Assert(reqContentType, Equals, "application/x-www-form-urlencoded")

		body, bodyErr := ioutil.ReadAll(r.Body)
		c.Assert(bodyErr, IsNil)
		c.Assert(string(body), Equals, "ptr=testen.de")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		pwd, pwdErr := os.Getwd()
		c.Assert(pwdErr, IsNil)

		data, readErr := ioutil.ReadFile(fmt.Sprintf("%s/test/response/rdns_get.json", pwd))
		c.Assert(readErr, IsNil)

		_, err := w.Write(data)
		c.Assert(err, IsNil)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	input := &models.RdnsSetInput{
		Ptr: "testen.de",
	}

	rdns, err := robotClient.RDnsSet(testIP, input)
	c.Assert(err, IsNil)
	c.Assert(rdns.IP, Equals, testIP)
}

func (s *ClientSuite) TestRDnsSetInvalidResponse(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqContentType := r.Header.Get("Content-Type")
		c.Assert(reqContentType, Equals, "application/x-www-form-urlencoded")

		body, bodyErr := ioutil.ReadAll(r.Body)
		c.Assert(bodyErr, IsNil)
		c.Assert(string(body), Equals, "ptr=testen.de")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, err := w.Write([]byte("invalid JSON"))
		c.Assert(err, IsNil)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	input := &models.RdnsSetInput{
		Ptr: "testen.de",
	}

	_, err := robotClient.RDnsSet(testIP, input)
	c.Assert(err, Not(IsNil))
}

func (s *ClientSuite) TestRDnsSetServerError(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	input := &models.RdnsSetInput{
		Ptr: "testen.de",
	}

	_, err := robotClient.RDnsSet(testIP, input)
	c.Assert(err, Not(IsNil))
}
