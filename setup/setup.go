package setup

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path"
	"strings"
	"syscall"

	log "github.com/Sirupsen/logrus"
	"github.com/flexiant/concerto/utils"

	"github.com/asaskevich/govalidator"
	"github.com/codegangsta/cli"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/net/html"
	"golang.org/x/net/publicsuffix"
)

type WebClient struct {
	client   *http.Client
	endpoint string
	csrf     string
}

func NewWebClient(endpoint string) (*WebClient, error) {
	transport := &http.Transport{}
	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, err := cookiejar.New(&options)
	if err != nil {
		return nil, err
	}
	client := &http.Client{Transport: transport, Jar: jar}

	return &WebClient{client, endpoint, ""}, nil
}

func (w *WebClient) obtainCsrf(b io.Reader) {
	z := html.NewTokenizer(b)
	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			return
		case tt == html.SelfClosingTagToken:
			t := z.Token()
			isMeta := t.Data == "meta"
			if isMeta && len(t.Attr) > 0 {
				if (t.Attr[1].Key == "name") && (t.Attr[1].Val == "csrf-token") {
					w.csrf = t.Attr[0].Val
					log.Debugf("Csrf Token: %s", w.csrf)
				} else if (t.Attr[0].Key == "name") && (t.Attr[0].Val == "csrf-token") {
					w.csrf = t.Attr[1].Val
					log.Debugf("Csrf Token: %s", w.csrf)
				}

			}
		}
	}
}

func (w *WebClient) checkErrorMessage(b io.Reader) error {
	var errorMessage error
	z := html.NewTokenizer(b)
	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			return errorMessage
		case tt == html.StartTagToken:
			t := z.Token()
			if (t.Data == "div") && len(t.Attr) > 0 && (t.Attr[0].Key == "id") && (t.Attr[0].Val == "flash_alert") {
				z.Next()
				errorMessage = errors.New(z.Token().Data)
			}
		}
	}
	return errorMessage
}

func (w *WebClient) login(email string, password string) error {
	response, err := w.client.Get(fmt.Sprintf("%s/accounts/login", w.endpoint))
	if err != nil {
		log.Fatalf("%#v", err)
	}

	b := response.Body
	defer b.Close()

	w.obtainCsrf(b)

	if w.csrf == "" {
		log.Debugf("Can not log into %s as %s", w.endpoint, email)
		return errors.New(fmt.Sprintf("Can not log into %s as %s", w.endpoint, email))
	}

	account := url.Values{}
	account.Set("authenticity_token", w.csrf)
	account.Set("account[email]", email)
	account.Set("account[password]", password)

	response, err = w.client.PostForm(fmt.Sprintf("%s/accounts/login", w.endpoint), account)
	b = response.Body
	defer b.Close()
	if err != nil {
		return err
	}

	err = w.checkErrorMessage(b)

	if err == nil {
		log.Debugf("Logged in %s as %s", w.endpoint, email)
	}

	return err
}

func (w *WebClient) getApiKeys() error {

	response, err := w.client.Get(fmt.Sprintf("%s/settings/api_key.zip", w.endpoint))
	defer response.Body.Close()

	if err != nil {
		return err
	}

	if response.StatusCode < 300 && response.Header.Get("Content-Type") == "application/zip" {
		concertoFolder, server := utils.GetConcertoDir()
		concertoFolderSSL := path.Join(concertoFolder, "ssl")
		if !server {
			os.MkdirAll(path.Join(concertoFolderSSL, "private"), 0755)
			file, err := ioutil.TempFile(os.TempDir(), "api-key.zip")
			if err != nil {
				return err
			}
			defer file.Close()
			io.Copy(file, response.Body)

			err = utils.Unzip(file.Name(), concertoFolderSSL)
			defer os.Remove(file.Name())
			if err != nil {
				return err
			} else {
				log.Infof("Unziped Api Keys in %s. Please enjoy of concerto cli.\n", concertoFolderSSL)
				return nil
			}
		} else {
			return errors.New("You are trying to overwrite server configuration. Please contact your administrator")
		}
	} else {
		return errors.New(fmt.Sprintf("We are not able to download you api keys. Please try by loging to %s/settings/api_key.zip in your web navigator ", w.endpoint))
	}
	return nil
}

func cmdSetupApiKeys(c *cli.Context) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("We are going to log into Concerto %s \nEmail: ", utils.GetConcertoUrl())
	emailUnClean, _ := reader.ReadString('\n')

	fmt.Printf("Password: ")
	passwordUnClean, _ := terminal.ReadPassword(int(syscall.Stdin))

	email := strings.TrimSpace(string(emailUnClean))
	password := strings.TrimSpace(string(passwordUnClean))
	fmt.Printf("\n")

	if govalidator.IsEmail(email) {
		client, err := NewWebClient(utils.GetConcertoUrl())
		if err != nil {
			log.Fatal(err)
		}

		err = client.login(email, password)
		if err != nil {
			log.Fatal(err)
		}
		err = client.getApiKeys()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatalf("Email address %s is not a valid email", email)
	}
}

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "api_keys",
			Usage:  "Install Concerto Api Keys",
			Action: cmdSetupApiKeys,
		},
	}
}
