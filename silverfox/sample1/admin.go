package sample1

import (
	"bytes"
	"encoding/csv"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"silverfox/sample1/stock/service"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"google.golang.org/appengine/user"
)

type Html struct {
	Url   string
	Admin bool
	Err   string
}

func adminHandler(w http.ResponseWriter, r *http.Request) {

	ctx := appengine.NewContext(r)
	u := user.Current(ctx)

	html := Html{}
	tmpl := template.New("new")
	if u == nil {
		loginUrl, _ := user.LoginURL(ctx, "/admin")
		html.Url = loginUrl
		tmpl.Parse("<html><body><a href='{{.Url}}'>ログイン</a></body></html>")
	} else {
		logoutUrl, _ := user.LogoutURL(ctx, "/admin")
		html.Url = logoutUrl
		tmpl.Parse("<html><body>{{if .Admin}}<p><a href='/admin/reset'>データリセット</a></p>{{end}}{{if .Err}}<p>{{.Err}}</p>{{end}}<p><a href='{{.Url}}'>ログアウト</a></p></body></html>")
		html.Admin = u.Admin
	}

	if html.Admin && r.URL.Path == "/admin/reset" {
		err := refreshData(w, r)
		if err != nil {
			html.Err = err.Error()
		}
	}

	tmpl.Execute(w, html)
}

func refreshData(w http.ResponseWriter, r *http.Request) error {

	ctx := appengine.NewContext(r)

	// 全件削除
	service.DeleteAll(ctx)

	// 設定ファイルにしようかと思ったけどとりあえずそのまま
	url := "https://silverfox-sample1.appspot.com/data/data.csv"
	client := urlfetch.Client(ctx)
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// CSVファイルがShift-JISなのでUTF-8に変換して読込
	reader := csv.NewReader(transform.NewReader(bytes.NewReader(body), japanese.ShiftJIS.NewDecoder()))

	var header []string
	i := 0
	for {
		row, err := reader.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			// 読み込みエラー発生
			break
		}
		if len(header) == 0 {
			header = row
			continue
		}

		entity, dxoErr := service.CsvToEntity(row)
		if len(dxoErr) == 0 {
			_, err := service.FullUpdate(ctx, entity.Code, entity)
			if err == nil {
				i++
			}
		}
	}
	return nil
}
