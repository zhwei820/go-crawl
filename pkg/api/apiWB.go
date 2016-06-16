package api

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/result"
	"encoding/json"
	"github.com/llitfkitfk/cirkol/pkg/util"
)

func SearchWBProfile(c *gin.Context, ch chan <- result.Profile) {
	userId := c.Param("userId")

	url := "http://mapi.weibo.com/2/profile?gsid=_&c=&s=&user_domain=" + userId
	var profile result.Profile
	var data result.WBRawProfile

	body, err := ReqApi(url)
	if err != nil {
		profile.ErrCode = ERROR_CODE_API_TIMEOUT
		profile.ErrMessage = err.Error()
	} else {
		profile.Website = url
		profile.RawData = body
		err = json.Unmarshal([]byte(profile.RawData), &data)
		if err != nil {
			profile.ErrCode = ERROR_CODE_JSON_ERROR
			profile.ErrMessage = err.Error()
		} else {
			profile.MergeWBRawProfile(data)
		}
	}
	ch <- profile
}

func SearchWBUID(c *gin.Context, ch chan <-result.UID) {
	rawurl := c.PostForm("url")
	var uid result.UID
	body, err := ReqApi(rawurl)
	if err != nil {
		uid.ErrCode = ERROR_CODE_API_TIMEOUT
		uid.ErrMessage = err.Error()
	} else {
		matcher := util.Matcher(REGEXP_WEIBO_PROFILE_ID, body)
		uid.Url = rawurl
		uid.Media = "wb"
		if len(matcher) > 1 {
			uid.Status = true
			uid.UserId = matcher[1]
		} else {
			uid.ErrCode = ERROR_CODE_REGEX_MISS_MATCHED
			uid.ErrMessage = ERROR_MSG_REGEX_MISS_MATCHED
		}
	}
	ch <- uid
}