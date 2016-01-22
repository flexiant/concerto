package webservice

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/flexiant/concerto/config"
)

const contentDispositionRegex = "filename=\\\"([^\\\"]*){1}\\\""

type Webservice struct {
	config *config.Config
	client *http.Client
}

func NewWebService() (*Webservice, error) {
	config, err := config.ConcertoServerConfiguration()
	if err != nil {
		return nil, err
	}

	client, err := httpClient(config)
	if err != nil {
		return nil, err
	}

	return &Webservice{config, client}, nil
}

func httpClient(config *config.Config) (*http.Client, error) {

	// Loads Clients Certificates and creates and 509KeyPair
	cert, err := tls.LoadX509KeyPair(config.Certificate.Cert, config.Certificate.Key)
	if err != nil {
		return nil, err
	}

	// Load CA file used to verify server
	ca_pool := x509.NewCertPool()
	ca_cert, err := ioutil.ReadFile(config.Certificate.Ca)
	if err != nil {
		return nil, fmt.Errorf("could not load CA file: %v", err)
	}
	ca_pool.AppendCertsFromPEM(ca_cert)

	// Create a client with specific transport configurations
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      ca_pool,
			Certificates: []tls.Certificate{cert},
		},
	}

	client := &http.Client{Transport: transport}

	return client, nil
}

func (w *Webservice) Post(endpoint string, json []byte) (error, []byte, int) {
	log.Debugf("Connecting: %s%s", w.config.ApiEndpoint, endpoint)
	output := strings.NewReader(string(json))
	response, err := w.client.Post(w.config.ApiEndpoint+endpoint, "application/json", output)

	log.Debugf("Posting: %s", output)
	if err != nil {
		return err, nil, 4000
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	log.Debugf("Response: %s", body)
	log.Debugf("Status code: %s", response.Status)

	return nil, body, response.StatusCode
}

func (w *Webservice) Put(endpoint string, json []byte) (error, []byte, int) {
	log.Debugf("Connecting: %s%s", w.config.ApiEndpoint, endpoint)
	output := strings.NewReader(string(json))

	request, err := http.NewRequest("PUT", w.config.ApiEndpoint+endpoint, output)
	if err != nil {
		return err, nil, -1
	}

	request.Header = map[string][]string{"Content-type": {"application/json"}}
	response, err := w.client.Do(request)

	log.Debugf("Putting: %s", endpoint)
	if err != nil {
		return err, nil, -1
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	log.Debugf("Response: %s", body)
	log.Debugf("Status code: %s", response.Status)
	return nil, body, response.StatusCode
}

func (w *Webservice) Delete(endpoint string) (error, []byte, int) {
	log.Debugf("Connecting: %s%s", w.config.ApiEndpoint, endpoint)

	request, err := http.NewRequest("DELETE", w.config.ApiEndpoint+endpoint, nil)
	response, err := w.client.Do(request)

	log.Debugf("Deleting: %s", endpoint)
	if err != nil {
		return err, nil, -1
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	log.Debugf("Response: %s", body)
	log.Debugf("Status code: %s", response.Status)
	return nil, body, response.StatusCode
}

func (w *Webservice) Get(endpoint string) (error, []byte, int) {

	log.Debugf("Connecting: %s%s", w.config.ApiEndpoint, endpoint)
	response, err := w.client.Get(w.config.ApiEndpoint + endpoint)
	if err != nil {
		return err, nil, -1
	}
	defer response.Body.Close()

	log.Debugf("Status code: %s", response.Status)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err, nil, -1
	}

	log.Debugf("Response: %s", string(body))
	return nil, body, response.StatusCode
}

func (w *Webservice) GetFile(endpoint string, directoryPath string) (string, error) {

	log.Debugf("Connecting: %s%s", w.config.ApiEndpoint, endpoint)
	response, err := w.client.Get(w.config.ApiEndpoint + endpoint)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	log.Debugf("Status code: %s", response.Status)

	r, err := regexp.Compile(contentDispositionRegex)
	if err != nil {
		return "", err
	}

	fileName := r.FindStringSubmatch(response.Header.Get("Content-Disposition"))[1]
	if err != nil {
		return "", err
	}
	realFileName := fmt.Sprintf("%s/%s", directoryPath, fileName)

	output, err := os.Create(realFileName)
	if err != nil {
		return "", err
	}
	defer output.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		return "", err
	}

	log.Debugf("%#v bytes downloaded", n)
	return realFileName, nil
}
