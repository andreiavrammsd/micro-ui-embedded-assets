package main

import (
	"net/url"
	"strconv"
	"strings"
)

func getLimit(limit string) int {
	l := getIntValue(limit)

	if l <= 0 || l > maxLimit {
		l = defaultLimit
	}

	return l
}

func getPaginationUrl(offset, limit, dir string) string {
	o := getIntValue(offset)
	l := getIntValue(limit)

	var off int
	if dir == "next" {
		off = l + o
	} else {
		off = o - l
		if off < 0 {
			return ""
		}
	}

	u, err := url.Parse("")
	checkError(err)

	q := u.Query()
	q.Add("offset", strconv.Itoa(off))
	q.Add("limit", limit)
	u.RawQuery = q.Encode()

	return u.String()
}

func getIntValue(text string) int {
	text = strings.Trim(text, " ")
	if len(text) == 0 {
		return 0
	}

	intValue, err := strconv.Atoi(text)
	if err != nil {
		return 0
	}

	return intValue
}

func getAssetsContent(assets []string) ([]byte, error) {
	content := []byte{}
	for _, file := range assets {
		c, err := Asset(file)
		if err != nil {
			return nil, err
		}
		content = append(content, c...)
	}

	return content, nil
}
