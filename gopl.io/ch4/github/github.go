// 包gitgub提供了GitHub issue跟踪接口和Go API
// 详细查看https://developer.github.com/v3/search/#search-issues.
package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const issuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues函数查询GitHub的issue跟踪接口
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(issuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// 我们必须在所有可能的分支上面关闭resp.Bogy
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
