package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	gh "github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var accessToken string
var tokenOK bool

const githubEnvVar = "GITHUB_TOKEN"

func init() {
	accessToken, tokenOK = os.LookupEnv(githubEnvVar)
	if !tokenOK {
		log.Fatalf("%s not set\n", githubEnvVar)
	}
	// or comment out code above and enter your token here
	// from https://github.com/settings/tokens
	// accessToken = "secret code"
}

const owner = "bbyars"
const repo = "mountebank"
const filepath = "README.md"
const branch = "master"

var proxyPort = flag.Int("proxy", 0, "Proxy Port")

func main() {

	flag.Parse()

	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)

	tc := oauth2.NewClient(ctx, ts)

	if t, ok := tc.Transport.(*oauth2.Transport); ok {
		dt := http.DefaultTransport.(*http.Transport)
		if *proxyPort != 0 {
			dt.Proxy = http.ProxyURL(proxyURL())
		}
		if dt.TLSClientConfig != nil {
			dt.TLSClientConfig.InsecureSkipVerify = true
		} else {
			dt.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		}
		t.Base = dt
	}

	c := gh.NewClient(tc)

	rc, _, _, err := c.Repositories.GetContents(
		ctx,
		owner,
		repo,
		filepath,
		&gh.RepositoryContentGetOptions{
			Ref: branch,
		},
	)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var content string
	content, err = rc.GetContent()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	println(content)

}

func proxyURL() *url.URL {
	u, err := url.Parse(fmt.Sprintf("https://localhost:%d", *proxyPort))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return u
}
