package gitbot

import (
	"fmt"
	"net/http"
	"os"
	"testing"
)

func Setup(plName string, t *testing.T) fmt.Stringer {
	return SetupB(plName, plName, t)
}

func SetupB(plName string, plFile string, t *testing.T) fmt.Stringer {
	module := gitHandler{}
	filename := fmt.Sprintf("test_data/%s.json", plFile)
	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Could not open test data file: %s", err.Error())
	}
	r, _ := http.NewRequest("POST", "/", file)
	r.Header.Set("X-GITHUB-EVENT", plName)
	ret := module.Handle(r)
	if ret == nil {
		t.Fatalf("Handling %s failed to generate a response.", plName)
	}
	return ret
}

func TestCommitComment(t *testing.T) {
	ret := Setup("commit_comment", t)
	expected := "[baxterthehacker/public-repo] New CommitComment by baxterthehacker https://github.com/baxterthehacker/public-repo/commit/9049f1265b7d61be4a8904a9a27120d2064dab3b#commitcomment-11056394"
	if str := ret.String(); str != expected {
		t.Error("String was wrong:", str)
		t.Error(" Expected string:", expected)
	}
	expected = "[baxterthehacker/public-repo] New CommitComment on lolfile.c:15 by baxterthehacker https://github.com/baxterthehacker/public-repo/commit/9049f1265b7d61be4a8904a9a27120d2064dab3b#commitcomment-11056394"
	ret = SetupB("commit_comment", "commit_comment_on_line", t)
	if str := ret.String(); str != expected {
		t.Error("String was wrong:", str)
		t.Error(" Expected string:", expected)
	}
}

func TestCreate(t *testing.T) {
	ret := Setup("create", t)
	expected := "[baxterthehacker/public-repo] baxterthehacker has created tag 0.0.1"
	if str := ret.String(); str != expected {
		t.Error("String was wrong:", str)
		t.Error(" Expected string:", expected)
	}
	expected = "[aarondl/ultimateq] aarondl has created branch reaper"
	ret = SetupB("create", "create_branch", t)
	if str := ret.String(); str != expected {
		t.Error("String was wrong:", str)
		t.Error(" Expected string:", expected)
	}
}

func TestDelete(t *testing.T) {
	ret := Setup("delete", t)
	expected := "[baxterthehacker/public-repo] baxterthehacker has deleted tag simple-tag"
	if str := ret.String(); str != expected {
		t.Error("String was wrong:", str)
		t.Error(" Expected string:", expected)
	}
}

func TestIssues(t *testing.T) {
	ret := Setup("issues", t)
	if str := ret.String(); str != "[baxterthehacker/public-repo] baxterthehacker opened issue #2 (Spelling error in the README file)" {
		t.Error("String was wrong:", str)
	}
}
