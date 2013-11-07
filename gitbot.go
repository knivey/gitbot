// gitbot project gitbot.go
package gitbot

import (
	"encoding/json"
	"fmt"
	"github.com/aarondl/gitio"
	"github.com/aarondl/ultimateq/data"
	"log"
	"net/http"
	"strings"
)

var (
	Endpoint *data.DataEndpoint
	Port     string
	Chan     string
)

type Payload struct {
	Before      string
	After       string
	Ref         string
	Compare     string
	Created     bool
	Deleted     bool
	Forced      bool
	Commits     []*Commit
	Head_commit *Commit
	Pusher      *Person //may not have Username, beware
	Repository  *Repository
}

type Commit struct {
	Added     []string
	Modified  []string
	Removed   []string
	Author    *Person
	Committer *Person
	Distinct  bool
	Id        string
	Message   string
	Timestamp string //Might be able to use time here
	Url       string
}

type Person struct {
	Email    string
	Name     string
	Username string
}

type Repository struct {
	Created_at    uint64
	Description   string
	Fork          bool
	Forks         uint
	Has_downloads bool
	Has_issues    bool
	Has_wiki      bool
	Homepage      string
	Id            uint
	Language      string
	Master_branch string
	Name          string
	Open_issues   uint
	Owner         *Person
	Private       bool
	Pushed_at     uint64
	Size          uint
	Stargazers    uint
	Url           string
	Watchers      uint
}

func alertChan(pl Payload) {
	name := pl.Repository.Name
	num := len(pl.Commits)
	who := pl.Pusher.Name
	refs := strings.Split(pl.Ref, "/")
	branch := refs[len(refs)-1]
	comp, err := gitio.Shorten(pl.Compare)
	if err != nil {
		comp = pl.Compare
	}

	msg := fmt.Sprintf("[%s] %s pushed %v commits to %s %s", name, who, num, branch, comp)
	Endpoint.Privmsg(Chan, msg)

	var c *Commit
	for i := 0; i < num; i++ {
		c = pl.Commits[i]
		msg = fmt.Sprintf("[%s/%s] %s: %s: %s", name, branch, c.Id[:7], c.Committer.Username, c.Message)
		Endpoint.Privmsg(Chan, msg)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	payload := r.FormValue("payload")
	var Pl Payload
	err := json.Unmarshal([]byte(payload), &Pl)
	if err != nil {
		log.Println(err)
	}
}

//Start the http listener
func Listen() (err error) {
	http.HandleFunc("/", handler)
	err = http.ListenAndServe(":"+Port, nil)
	if err != nil {
		return err
	}
	return nil
}
