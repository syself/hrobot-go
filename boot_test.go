package client_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"

	. "gopkg.in/check.v1"

	client "github.com/syself/hrobot-go"
	"github.com/syself/hrobot-go/models"
)

func (s *ClientSuite) TestBootRescueGetInactiveSuccess(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		pwd, pwdErr := os.Getwd()
		c.Assert(pwdErr, IsNil)

		data, readErr := ioutil.ReadFile(fmt.Sprintf("%s/test/response/boot_rescue_get_inactive.json", pwd))
		c.Assert(readErr, IsNil)

		_, err := w.Write(data)
		c.Assert(err, IsNil)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	rescue, err := robotClient.BootRescueGet(testServerID)
	c.Assert(err, IsNil)
	c.Assert(rescue.ServerNumber, Equals, testServerID)
}

func (s *ClientSuite) TestBootRescueGetGetInvalidResponse(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, err := w.Write([]byte("invalid JSON"))
		c.Assert(err, IsNil)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	_, err := robotClient.BootRescueGet(testServerID)
	c.Assert(err, Not(IsNil))
}

func (s *ClientSuite) TestBootRescueGetServerError(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	_, err := robotClient.BootRescueGet(testServerID)
	c.Assert(err, Not(IsNil))
}

func (s *ClientSuite) TestBootRescueGetActiveSuccess(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		pwd, pwdErr := os.Getwd()
		c.Assert(pwdErr, IsNil)

		data, readErr := ioutil.ReadFile(fmt.Sprintf("%s/test/response/boot_rescue_get_active.json", pwd))
		c.Assert(readErr, IsNil)

		_, err := w.Write(data)
		c.Assert(err, IsNil)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	rescue, err := robotClient.BootRescueGet(testServerID)
	c.Assert(err, IsNil)
	c.Assert(rescue.ServerNumber, Equals, testServerID)
}

func (s *ClientSuite) TestBootRescueDeleteSuccess(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		pwd, pwdErr := os.Getwd()
		c.Assert(pwdErr, IsNil)

		data, readErr := ioutil.ReadFile(fmt.Sprintf("%s/test/response/boot_rescue_delete.json", pwd))
		c.Assert(readErr, IsNil)

		_, err := w.Write(data)
		c.Assert(err, IsNil)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	rescue, err := robotClient.BootRescueDelete(testServerID)
	c.Assert(err, IsNil)
	c.Assert(rescue.ServerNumber, Equals, testServerID)
}

func (s *ClientSuite) TestBootRescueSetSuccess(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqContentType := r.Header.Get("Content-Type")
		c.Assert(reqContentType, Equals, "application/x-www-form-urlencoded")

		body, bodyErr := ioutil.ReadAll(r.Body)
		c.Assert(bodyErr, IsNil)
		c.Assert(string(body), Equals, "arch=64&os=linux")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		pwd, pwdErr := os.Getwd()
		c.Assert(pwdErr, IsNil)

		data, readErr := ioutil.ReadFile(fmt.Sprintf("%s/test/response/boot_rescue_set.json", pwd))
		c.Assert(readErr, IsNil)

		_, err := w.Write(data)
		c.Assert(err, IsNil)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	input := &models.RescueSetInput{
		OS:   "linux",
		Arch: 64,
	}

	rescue, err := robotClient.BootRescueSet(testServerID, input)
	c.Assert(err, IsNil)
	c.Assert(rescue.ServerNumber, Equals, testServerID)
	c.Assert(len(rescue.AuthorizedKey), Equals, 0)
}

func (s *ClientSuite) TestBootRescueSetWithKeySuccess(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqContentType := r.Header.Get("Content-Type")
		c.Assert(reqContentType, Equals, "application/x-www-form-urlencoded")

		body, bodyErr := ioutil.ReadAll(r.Body)
		c.Assert(bodyErr, IsNil)
		c.Assert(string(body), Equals, "arch=64&authorized_key=fi%3Ang%3Aer%3Apr%3Ain%3At0%3A00%3A00%3A00%3A00%3A00%3A00%3A00%3A00%3A00%3A00&os=linux")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		pwd, pwdErr := os.Getwd()
		c.Assert(pwdErr, IsNil)

		data, readErr := ioutil.ReadFile(fmt.Sprintf("%s/test/response/boot_rescue_set_with_key.json", pwd))
		c.Assert(readErr, IsNil)

		_, err := w.Write(data)
		c.Assert(err, IsNil)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	input := &models.RescueSetInput{
		OS:            "linux",
		Arch:          64,
		AuthorizedKey: "fi:ng:er:pr:in:t0:00:00:00:00:00:00:00:00:00:00",
	}

	rescue, err := robotClient.BootRescueSet(testServerID, input)
	c.Assert(err, IsNil)
	c.Assert(len(rescue.AuthorizedKey), Equals, 1)
	c.Assert(rescue.AuthorizedKey[0].Key.Fingerprint, Equals, "fi:ng:er:pr:in:t0:00:00:00:00:00:00:00:00:00:00")
}

func (s *ClientSuite) TestBootRescueSetInvalidResponse(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqContentType := r.Header.Get("Content-Type")
		c.Assert(reqContentType, Equals, "application/x-www-form-urlencoded")

		body, bodyErr := ioutil.ReadAll(r.Body)
		c.Assert(bodyErr, IsNil)
		c.Assert(string(body), Equals, "arch=64&os=linux")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, err := w.Write([]byte("invalid JSON"))
		c.Assert(err, IsNil)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	input := &models.RescueSetInput{
		OS:   "linux",
		Arch: 64,
	}

	_, err := robotClient.BootRescueSet(testServerID, input)
	c.Assert(err, Not(IsNil))
}

func (s *ClientSuite) TestBootRescueSetServerError(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	input := &models.RescueSetInput{
		OS:   "linux",
		Arch: 64,
	}

	_, err := robotClient.BootRescueSet(testServerID, input)
	c.Assert(err, Not(IsNil))
}

func (s *ClientSuite) TestBootLinuxGetInactiveSuccess(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		pwd, pwdErr := os.Getwd()
		c.Assert(pwdErr, IsNil)

		data, readErr := ioutil.ReadFile(fmt.Sprintf("%s/test/response/boot_linux_get_inactive.json", pwd))
		c.Assert(readErr, IsNil)

		_, err := w.Write(data)
		c.Assert(err, IsNil)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	linux, err := robotClient.BootLinuxGet(testServerID)
	c.Assert(err, IsNil)
	c.Assert(linux.ServerNumber, Equals, testServerID)
}

func (s *ClientSuite) TestBootLinuxGetGetInvalidResponse(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, err := w.Write([]byte("invalid JSON"))
		c.Assert(err, IsNil)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	_, err := robotClient.BootLinuxGet(testServerID)
	c.Assert(err, Not(IsNil))
}

func (s *ClientSuite) TestBootLinuxGetServerError(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	_, err := robotClient.BootLinuxGet(testServerID)
	c.Assert(err, Not(IsNil))
}

func (s *ClientSuite) TestBootLinuxDeleteSuccess(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		pwd, pwdErr := os.Getwd()
		c.Assert(pwdErr, IsNil)

		data, readErr := ioutil.ReadFile(fmt.Sprintf("%s/test/response/boot_linux_delete.json", pwd))
		c.Assert(readErr, IsNil)

		_, err := w.Write(data)
		c.Assert(err, IsNil)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	linux, err := robotClient.BootLinuxDelete(testServerID)
	c.Assert(err, IsNil)
	c.Assert(linux.ServerNumber, Equals, testServerID)
}

func (s *ClientSuite) TestBootLinuxSetSuccess(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqContentType := r.Header.Get("Content-Type")
		c.Assert(reqContentType, Equals, "application/x-www-form-urlencoded")

		body, bodyErr := ioutil.ReadAll(r.Body)
		c.Assert(bodyErr, IsNil)
		c.Assert(string(body), Equals, "arch=32&dist=CentOS+5.5+minimal&lang=en")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		pwd, pwdErr := os.Getwd()
		c.Assert(pwdErr, IsNil)

		data, readErr := ioutil.ReadFile(fmt.Sprintf("%s/test/response/boot_linux_set.json", pwd))
		c.Assert(readErr, IsNil)

		_, err := w.Write(data)
		c.Assert(err, IsNil)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	input := &models.LinuxSetInput{
		Dist: "CentOS 5.5 minimal",
		Arch: 32,
		Lang: "en",
	}

	linux, err := robotClient.BootLinuxSet(testServerID, input)
	c.Assert(err, IsNil)
	c.Assert(linux.ServerNumber, Equals, testServerID)
	c.Assert(len(linux.AuthorizedKey), Equals, 0)
}

func (s *ClientSuite) TestBootLinuxSetWithKeySuccess(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqContentType := r.Header.Get("Content-Type")
		c.Assert(reqContentType, Equals, "application/x-www-form-urlencoded")

		body, bodyErr := ioutil.ReadAll(r.Body)
		c.Assert(bodyErr, IsNil)
		c.Assert(string(body), Equals, "arch=32&authorized_key=fi%3Ang%3Aer%3Apr%3Ain%3At0%3A00%3A00%3A00%3A00%3A00%3A00%3A00%3A00%3A00%3A00&dist=CentOS+5.5+minimal&lang=en")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		pwd, pwdErr := os.Getwd()
		c.Assert(pwdErr, IsNil)

		data, readErr := ioutil.ReadFile(fmt.Sprintf("%s/test/response/boot_linux_set_with_key.json", pwd))
		c.Assert(readErr, IsNil)

		_, err := w.Write(data)
		c.Assert(err, IsNil)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	input := &models.LinuxSetInput{
		Dist:          "CentOS 5.5 minimal",
		Arch:          32,
		Lang:          "en",
		AuthorizedKey: "fi:ng:er:pr:in:t0:00:00:00:00:00:00:00:00:00:00",
	}

	linux, err := robotClient.BootLinuxSet(testServerID, input)
	c.Assert(err, IsNil)
	c.Assert(len(linux.AuthorizedKey), Equals, 1)
	c.Assert(linux.AuthorizedKey[0].Key.Fingerprint, Equals, "fi:ng:er:pr:in:t0:00:00:00:00:00:00:00:00:00:00")
}

func (s *ClientSuite) TestBootLinuxSetInvalidResponse(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqContentType := r.Header.Get("Content-Type")
		c.Assert(reqContentType, Equals, "application/x-www-form-urlencoded")

		body, bodyErr := ioutil.ReadAll(r.Body)
		c.Assert(bodyErr, IsNil)
		c.Assert(string(body), Equals, "arch=32&dist=CentOS+5.5+minimal&lang=en")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, err := w.Write([]byte("invalid JSON"))
		c.Assert(err, IsNil)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	input := &models.LinuxSetInput{
		Dist: "CentOS 5.5 minimal",
		Arch: 32,
		Lang: "en",
	}
	_, err := robotClient.BootLinuxSet(testServerID, input)
	c.Assert(err, Not(IsNil))
}

func (s *ClientSuite) TestBootLinuxSetServerError(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	input := &models.LinuxSetInput{
		Dist: "CentOS 5.5 minimal",
		Arch: 32,
		Lang: "en",
	}

	_, err := robotClient.BootLinuxSet(testServerID, input)
	c.Assert(err, Not(IsNil))
}
