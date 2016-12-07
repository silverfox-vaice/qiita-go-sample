package rest

import (
	"errors"
	"strings"
)

type UrlParam struct {
	Kind       string
	Keys       []string
	Params     map[string]string
	Conditions map[string]string
}

func (u *UrlParam) GetParam(index int) string {
	return u.Params[u.Keys[index]]
}

// URLのGET値をベースにParamを作成。
func NewUrlParam(url string, index int) (*UrlParam, error) {
	path := strings.Trim(url, "/")
	s := strings.Split(path, "/")
	if len(s) < index {
		return nil, errors.New(" index out of range.")
	}

	param := &UrlParam{
		Kind:       s[index-1],
		Keys:       make([]string, 0),
		Params:     map[string]string{},
		Conditions: map[string]string{},
	}

	for i := index; i < len(s); i = i + 2 {
		val := ""
		if len(s) > i+1 {
			val = s[i+1]
		}
		param.Keys = append(param.Keys, s[i])
		param.Params[s[i]] = val
	}
	return param, nil
}
