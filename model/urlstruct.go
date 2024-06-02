package model

type UrlGet struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		List []struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			ParentID int    `json:"parent_id"`
			Depth    int    `json:"depth"`
			ChExt    string `json:"ch_ext"`
			Children []struct {
				ID       int           `json:"id"`
				Name     string        `json:"name"`
				ParentID int           `json:"parent_id"`
				Depth    int           `json:"depth"`
				ChExt    string        `json:"ch_ext"`
				Children []interface{} `json:"children"`
				List     []struct {
					ContentID       int    `json:"content_id"`
					Title           string `json:"title"`
					Ext             string `json:"ext"`
					Icon            string `json:"icon"`
					BbsURL          string `json:"bbs_url"`
					ArticleUserName string `json:"article_user_name"`
					ArticleTime     string `json:"article_time"`
					AvatarURL       string `json:"avatar_url"`
					Summary         string `json:"summary"`
				} `json:"list"`
				Layout     string `json:"layout"`
				EntryLimit int    `json:"entry_limit"`
				Hidden     bool   `json:"hidden"`
			} `json:"children"`
			List       []interface{} `json:"list"`
			Layout     string        `json:"layout"`
			EntryLimit int           `json:"entry_limit"`
			Hidden     bool          `json:"hidden"`
		} `json:"list"`
	} `json:"data"`
}

type ChGet struct {
	ContentID       int    `json:"content_id"`
	Title           string `json:"title"`
	Ext             string `json:"ext"`
	Icon            string `json:"icon"`
	BbsURL          string `json:"bbs_url"`
	ArticleUserName string `json:"article_user_name"`
	ArticleTime     string `json:"article_time"`
	AvatarURL       string `json:"avatar_url"`
	Summary         string `json:"summary"`
}

type ExtGet struct {
	C18 struct {
		Filter struct {
			Text string `json:"text"`
		} `json:"filter"`
		Picture struct {
			List []string `json:"list"`
		} `json:"picture"`
	} `json:"c_18"`
}

type CharGet struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		Content struct {
			ID              int    `json:"id"`
			Title           string `json:"title"`
			Content         string `json:"content"`
			Ext             string `json:"ext"`
			AuthorName      string `json:"author_name"`
			EditorName      string `json:"editor_name"`
			Ctime           string `json:"ctime"`
			Mtime           string `json:"mtime"`
			Version         int    `json:"version"`
			Icon            string `json:"icon"`
			Summary         string `json:"summary"`
			URL             string `json:"url"`
			Type            int    `json:"type"`
			BbsURL          string `json:"bbs_url"`
			ArticleUserName string `json:"article_user_name"`
			ArticleTime     string `json:"article_time"`
			AvatarURL       string `json:"avatar_url"`
			Contents        []struct {
				Name string `json:"name"`
				Text string `json:"text"`
			} `json:"contents"`
			ForbidCorrectError bool `json:"forbid_correct_error"`
		} `json:"content"`
		ChannelList []struct {
			Slice []struct {
				ChannelID int    `json:"channel_id"`
				Name      string `json:"name"`
			} `json:"slice"`
		} `json:"channel_list"`
	} `json:"data"`
}
