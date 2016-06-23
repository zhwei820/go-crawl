package controllers

import (
	"time"
	"github.com/llitfkitfk/cirkol/pkg/models"
	"encoding/json"
	"github.com/llitfkitfk/cirkol/pkg/common"
)

type post struct {
	ID                 string `json:"id"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
	ShareCount         int    `json:"share_count"`
	LikeCount          int    `json:"like_count"`
	CommentCount       int    `json:"comment_count"`
	ViewCount          int    `json:"view_count"`
	ContentType        string `json:"content_type"`
	ContentCaption     string `json:"content_caption"`
	ContentBody        string `json:"content_body"`
	ContentFullPicture string `json:"content_full_picture"`
	PermalinkUrl       string `json:"permalink_url"`
	HasComment         bool   `json:"has_comment"`
	RawData            string `json:"raw_data"`
	Date               int64  `json:"date"`
}

type Posts struct {
	Items      []post `json:"data"`
	Date       int64  `json:"date"`
	Status     bool   `json:"status"`
	RawData    string `json:"raw_data"`
	ErrCode    int    `json:"error_code"`
	ErrMessage string `json:"error_message"`
}

func (p *Posts) FetchErr(err error) {
	p.ErrCode = ERROR_CODE_API_FETCH
	p.ErrMessage = err.Error()
	p.Date = time.Now().Unix()
	p.Status = false
}
func TimeOutPosts() Posts {
	var p Posts
	p.ErrCode = ERROR_CODE_API_TIMEOUT
	p.ErrMessage = ERROR_MSG_API_TIMEOUT
	p.Date = time.Now().Unix()
	p.Status = false
	return p
}

func (p *Posts) ParseRawPosts(data models.FBRawPosts) {
	for _, item := range data.Data {
		var data post
		data.ID = item.ID
		data.CreatedAt = common.DateFormat(item.CreatedTime)
		data.UpdatedAt = common.DateFormat(item.UpdatedTime)
		data.ShareCount = item.Shares.Count
		data.LikeCount = item.Likes.Summary.TotalCount
		data.CommentCount = item.Comments.Summary.TotalCount

		data.ContentType = item.Type

		data.ContentBody = item.Message
		data.ContentFullPicture = item.FullPicture
		data.PermalinkUrl = item.PermalinkURL
		data.HasComment = item.Comments.Summary.CanComment
		data.RawData = common.JsonToString(json.Marshal(item))
		data.Date = time.Now().Unix()

		p.Items = append(p.Items, data)
	}

	p.Status = true
}
