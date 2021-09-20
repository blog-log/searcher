package markdown

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"bloglog.com/search/v1/internal/model"
)

func Get(page *model.Page) (string, error) {
	url, err := GenerateUrl(page)
	if err != nil {
		return "", err
	}

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	if resp.StatusCode >= 300 {
		return "", fmt.Errorf("invalid status code %v", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	//Convert the body to type string
	content := string(body)
	return content, nil
}

func GenerateUrl(page *model.Page) (string, error) {
	const (
		githubRawUrl = "https://raw.githubusercontent.com"
	)

	repoSplit := strings.Split(page.Repo, "/")

	if len(repoSplit) < 3 {
		return "", errors.New("invalid repo url")
	}

	owner := repoSplit[len(repoSplit)-2]
	repoName := repoSplit[len(repoSplit)-1]

	requestUrl := fmt.Sprintf("%s/%s/%s/%s/%s", githubRawUrl, owner, repoName, page.Branch, page.Path)

	return requestUrl, nil
}
