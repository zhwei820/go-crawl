package api

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/result"
	"github.com/llitfkitfk/cirkol/pkg/util"
)

func SearchWXProfile(c *gin.Context, ch chan <- result.Profile) {
	userId := c.Param("userId")
	url := "http://weixin.sogou.com/weixin?type=1&query=" + userId + "&ie=utf8&_sug_=n&_sug_type_="
	body := GetProfileApi(url, ch)

	var profile result.Profile
	profile.UserId = userId
	profile.RawData = body
	profile.Website = util.DecodeString(MatchStrNoCh(0, REGEXP_WEIXIN_URL, body))
	profile.Avatar = MatchStrNoCh(0, REGEXP_WEIXIN_LOGO, body)
	profile.About = MatchStrNoCh(1, REGEXP_WEIXIN_FEATURE, body)

	profile.Status = true
	ch <- profile
}

func SearchWXPosts(c *gin.Context, ch chan <- result.Posts) {
	userId := c.Param("userId")
	url := "http://weixin.sogou.com/weixin?type=1&query=" + userId + "&ie=utf8&_sug_=n&_sug_type_="
	body := GetPostsApi(url, ch)

	urlStr := MatchStrPostCh(0, REGEXP_WEIXIN_URL, body, ch)
	postBody := GetPostsApi(urlStr, ch)

	postsStr := MatchStrPostCh(0, REGEXP_WEIXIN_POSTS, postBody, ch)
	jsonStr := util.DecodeString(postsStr)

	var posts result.Posts
	var data result.WXRawPosts
	ParsePostsJson(jsonStr, &data, ch)
	posts.MergeWXRawPosts(data)
	ch <- posts
}

func SearchWXUID(c *gin.Context, ch chan <-result.UID) {
	rawurl := c.PostForm("url")
	var uid result.UID
	body, err := ReqApi(rawurl)
	if err != nil {
		uid.ErrCode = ERROR_CODE_API_TIMEOUT
		uid.ErrMessage = err.Error()
	} else {
		matcher := util.Matcher(REGEXP_WEIXIN_PROFILE_ID, body)

		uid.Url = rawurl
		uid.Media = "wx"
		if len(matcher) > 0 {
			uid.Status = true
			uid.UserId = matcher[1]
		} else {
			uid.ErrCode = ERROR_CODE_REGEX_MISS_MATCHED
			uid.ErrMessage = ERROR_MSG_REGEX_MISS_MATCHED
		}
	}
	ch <- uid
}