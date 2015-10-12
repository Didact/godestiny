package destiny

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type MembershipType int

var db *sql.DB

const (
	All  MembershipType = 254
	PSN  MembershipType = 2
	Xbox MembershipType = 1
)

func (m MembershipType) String() string {
	return fmt.Sprintf("%d", int(m))
}

type c struct {
	apikey string
}

func (c *c) RoundTrip(r *http.Request) (*http.Response, error) {
	fmt.Println(r.URL)
	r.Header.Add("X-API-Key", c.apikey)
	return http.DefaultTransport.RoundTrip(r)
}

func init() {
	http.DefaultClient.Transport = &c{os.Getenv("BNETAPI")}
	content, err := sql.Open("sqlite3", "/home/dakota/go/src/github.com/Didact/destiny/content.db")
	if err != nil {
		log.Fatal(err)
	}
	db = content
}

func GetAccount(mt MembershipType, destinyMembershipId string) {
}

func GetAccountSummary(mt MembershipType, destinyMembershipId string, definitions bool) {
}

func GetActivityHistory(mt MembershipType, destinyMembershipId string, characterId string, params map[string]string) {
}

func GetAllItemsSummary(mt MembershipType, destinyMembershipId string, definitions bool) *ItemsSummary {
	url := fmt.Sprintf("http://www.bungie.net/Platform/Destiny/%s/Account/%s/Items/", mt, destinyMembershipId)
	if definitions {
		url += "?definitions=true"
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil
	}
	is := ItemsSummary{}
	err = json.Unmarshal(body, &is)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &is
}

func GetItemDetail(mt MembershipType, destinyMembershipId, characterId, itemInstanceId string, definitions bool) {
}

func SearchDestinyPlayer(mt MembershipType, displayName string) string {
	type ss struct {
		Response []struct {
			IconPath       string
			MembershipType float64
			MembershipId   string
			DisplayName    string
		}
		ErrorCode       float64
		ThrottleSeconds float64
		ErrorStatus     string
		Message         string
		MessageData     interface{}
	}

	url := fmt.Sprintf("http://www.bungie.net/Platform/Destiny/SearchDestinyPlayer/%s/%s/", mt, displayName)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	s := ss{}
	err = json.Unmarshal(body, &s)
	if err != nil {
		log.Fatal(err)
	}
	if len(s.Response) == 0 {
		return ""
	}
	return s.Response[0].MembershipId
}
