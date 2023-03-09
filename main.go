package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type UserInfoChanged struct {
	Meta struct {
		Type      string `json:"type"`
		EventId   string `json:"event_id"`
		CreatedAt int64  `json:"created_at"`
		TraceId   string `json:"trace_id"`
		ServiceId string `json:"service_id"`
	} `json:"meta"`
	Payload struct {
		Id         int      `json:"id"`
		Username   string   `json:"username"`
		Followers  []string `json:"followers"`
		Repos      []string `json:"repos"`
		Email      string   `json:"email"`
		FirstName  string   `json:"first_name"`
		LastName   string   `json:"last_name"`
		TimeZoneId string   `json:"time_zone_id"`
	} `json:"payload"`
}

type UpdateUserData struct {
	Email      string `json:"email"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	TimeZoneId string `json:"time_zone_id"`
}

type GitUser struct {
	Login             string    `json:"login"`
	Id                int       `json:"id"`
	NodeId            string    `json:"node_id"`
	AvatarUrl         string    `json:"avatar_url"`
	GravatarId        string    `json:"gravatar_id"`
	Url               string    `json:"url"`
	HtmlUrl           string    `json:"html_url"`
	FollowersUrl      string    `json:"followers_url"`
	FollowingUrl      string    `json:"following_url"`
	GistsUrl          string    `json:"gists_url"`
	StarredUrl        string    `json:"starred_url"`
	SubscriptionsUrl  string    `json:"subscriptions_url"`
	OrganizationsUrl  string    `json:"organizations_url"`
	ReposUrl          string    `json:"repos_url"`
	EventsUrl         string    `json:"events_url"`
	ReceivedEventsUrl string    `json:"received_events_url"`
	Type              string    `json:"type"`
	SiteAdmin         bool      `json:"site_admin"`
	Name              string    `json:"name"`
	Company           string    `json:"company"`
	Blog              string    `json:"blog"`
	Location          string    `json:"location"`
	Email             string    `json:"email"`
	Hireable          bool      `json:"hireable"`
	Bio               string    `json:"bio"`
	TwitterUsername   string    `json:"twitter_username"`
	PublicRepos       int       `json:"public_repos"`
	PublicGists       int       `json:"public_gists"`
	Followers         int       `json:"followers"`
	Following         int       `json:"following"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type GitRepo struct {
	Id       int    `json:"id"`
	NodeId   string `json:"node_id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Private  bool   `json:"private"`
	Owner    struct {
		Login             string `json:"login"`
		Id                int    `json:"id"`
		NodeId            string `json:"node_id"`
		AvatarUrl         string `json:"avatar_url"`
		GravatarId        string `json:"gravatar_id"`
		Url               string `json:"url"`
		HtmlUrl           string `json:"html_url"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReposUrl          string `json:"repos_url"`
		EventsUrl         string `json:"events_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"owner"`
	HtmlUrl          string    `json:"html_url"`
	Description      *string   `json:"description"`
	Fork             bool      `json:"fork"`
	Url              string    `json:"url"`
	ForksUrl         string    `json:"forks_url"`
	KeysUrl          string    `json:"keys_url"`
	CollaboratorsUrl string    `json:"collaborators_url"`
	TeamsUrl         string    `json:"teams_url"`
	HooksUrl         string    `json:"hooks_url"`
	IssueEventsUrl   string    `json:"issue_events_url"`
	EventsUrl        string    `json:"events_url"`
	AssigneesUrl     string    `json:"assignees_url"`
	BranchesUrl      string    `json:"branches_url"`
	TagsUrl          string    `json:"tags_url"`
	BlobsUrl         string    `json:"blobs_url"`
	GitTagsUrl       string    `json:"git_tags_url"`
	GitRefsUrl       string    `json:"git_refs_url"`
	TreesUrl         string    `json:"trees_url"`
	StatusesUrl      string    `json:"statuses_url"`
	LanguagesUrl     string    `json:"languages_url"`
	StargazersUrl    string    `json:"stargazers_url"`
	ContributorsUrl  string    `json:"contributors_url"`
	SubscribersUrl   string    `json:"subscribers_url"`
	SubscriptionUrl  string    `json:"subscription_url"`
	CommitsUrl       string    `json:"commits_url"`
	GitCommitsUrl    string    `json:"git_commits_url"`
	CommentsUrl      string    `json:"comments_url"`
	IssueCommentUrl  string    `json:"issue_comment_url"`
	ContentsUrl      string    `json:"contents_url"`
	CompareUrl       string    `json:"compare_url"`
	MergesUrl        string    `json:"merges_url"`
	ArchiveUrl       string    `json:"archive_url"`
	DownloadsUrl     string    `json:"downloads_url"`
	IssuesUrl        string    `json:"issues_url"`
	PullsUrl         string    `json:"pulls_url"`
	MilestonesUrl    string    `json:"milestones_url"`
	NotificationsUrl string    `json:"notifications_url"`
	LabelsUrl        string    `json:"labels_url"`
	ReleasesUrl      string    `json:"releases_url"`
	DeploymentsUrl   string    `json:"deployments_url"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	PushedAt         time.Time `json:"pushed_at"`
	GitUrl           string    `json:"git_url"`
	SshUrl           string    `json:"ssh_url"`
	CloneUrl         string    `json:"clone_url"`
	SvnUrl           string    `json:"svn_url"`
	Homepage         *string   `json:"homepage"`
	Size             int       `json:"size"`
	StargazersCount  int       `json:"stargazers_count"`
	WatchersCount    int       `json:"watchers_count"`
	Language         *string   `json:"language"`
	HasIssues        bool      `json:"has_issues"`
	HasProjects      bool      `json:"has_projects"`
	HasDownloads     bool      `json:"has_downloads"`
	HasWiki          bool      `json:"has_wiki"`
	HasPages         bool      `json:"has_pages"`
	HasDiscussions   bool      `json:"has_discussions"`
	ForksCount       int       `json:"forks_count"`
	MirrorUrl        string    `json:"mirror_url"`
	Archived         bool      `json:"archived"`
	Disabled         bool      `json:"disabled"`
	OpenIssuesCount  int       `json:"open_issues_count"`
	License          *struct {
		Key    string `json:"key"`
		Name   string `json:"name"`
		SpdxId string `json:"spdx_id"`
		Url    string `json:"url"`
		NodeId string `json:"node_id"`
	} `json:"license"`
	AllowForking             bool     `json:"allow_forking"`
	IsTemplate               bool     `json:"is_template"`
	WebCommitSignoffRequired bool     `json:"web_commit_signoff_required"`
	Topics                   []string `json:"topics"`
	Visibility               string   `json:"visibility"`
	Forks                    int      `json:"forks"`
	OpenIssues               int      `json:"open_issues"`
	Watchers                 int      `json:"watchers"`
	DefaultBranch            string   `json:"default_branch"`
}

var selectedGitUser GitUser
var userFollowers []GitUser
var userRepos []GitRepo
var userInfoChanged UserInfoChanged
var usersUrl string = "https://api.github.com/users"

func main() {
	handleRequests()
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/produce/{user_id}", produceUser).Methods("POST")
	myRouter.HandleFunc("/users/{user_id}", updateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func createUserInfoChanged(w http.ResponseWriter, gitUser GitUser, fromUpdate bool) {

	// get followers list
	resp, err := http.Get(usersUrl + "/" + gitUser.Login + "/followers")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(body, &userFollowers)

	if err != nil {
		log.Fatalln(err)
	}

	// get repos list
	resp2, err2 := http.Get(usersUrl + "/" + gitUser.Login + "/repos")
	if err2 != nil {
		log.Fatalln(err2)
	}

	body2, err2 := ioutil.ReadAll(resp2.Body)
	if err2 != nil {
		log.Fatalln(err2)
	}
	err2 = json.Unmarshal(body2, &userRepos)

	if err2 != nil {
		log.Fatalln(err2)
	}

	// assigning values to userInfoChanged
	userInfoChanged.Meta.Type = "UserInfoChanged"
	userInfoChanged.Meta.ServiceId = "user-service"
	userInfoChanged.Meta.EventId = uuid.NewV4().String()
	userInfoChanged.Meta.TraceId = uuid.NewV4().String()
	userInfoChanged.Meta.CreatedAt = time.Now().UnixNano()
	userInfoChanged.Payload.Id = gitUser.Id
	userInfoChanged.Payload.Username = gitUser.Login

	if !fromUpdate {
		userInfoChanged.Payload.Email = gitUser.Email

		names := strings.Split(gitUser.Name, " ")
		if len(names) > 1 {
			userInfoChanged.Payload.FirstName = names[0]
			userInfoChanged.Payload.LastName = names[1]
		}
		if len(names) == 1 {
			userInfoChanged.Payload.FirstName = gitUser.Name
			userInfoChanged.Payload.LastName = ""
		}
	}

	for _, follower := range userFollowers {
		userInfoChanged.Payload.Followers = append(userInfoChanged.Payload.Followers, follower.Login)
	}

	for _, repo := range userRepos {
		userInfoChanged.Payload.Repos = append(userInfoChanged.Payload.Followers, repo.Name)
	}

	json.NewEncoder(w).Encode(userInfoChanged)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: updateUser")

	userInfoChanged = UserInfoChanged{}
	var updateUserData UpdateUserData
	json.NewDecoder(r.Body).Decode(&updateUserData)
	// json.NewEncoder(w).Encode(updateUserData)

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	resp, err := http.Get(usersUrl + "/" + params["user_id"])
	if err != nil {
		log.Fatalln(err)
	}

	// assigning user that need to update to selectedGitUser
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(body, &selectedGitUser)

	if err != nil {
		log.Fatalln(err)
	}

	userInfoChanged.Payload.Email = updateUserData.Email
	userInfoChanged.Payload.FirstName = updateUserData.FirstName
	userInfoChanged.Payload.LastName = updateUserData.LastName
	userInfoChanged.Payload.TimeZoneId = updateUserData.TimeZoneId

	createUserInfoChanged(w, selectedGitUser, true)
}

func produceUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: produceUser")

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	resp, err := http.Get(usersUrl + "/" + params["user_id"])
	if err != nil {
		log.Fatalln(err)
	}

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(body, &selectedGitUser)

	if err != nil {
		log.Fatalln(err)
	}
	// json.NewEncoder(w).Encode(selectedGitUser)
	userInfoChanged = UserInfoChanged{}
	createUserInfoChanged(w, selectedGitUser, false)
}
