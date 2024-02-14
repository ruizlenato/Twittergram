package twitter

type TweetContent struct {
	Medias  []Medias `json:"medias"`
	Caption string   `json:"caption,omitempty"`
}

type Medias struct {
	Height int    `json:"height,omitempty"`
	Width  int    `json:"width,omitempty"`
	Source string `json:"src"`
	Video  bool   `json:"is_video"`
}

type TwitterAPIData struct {
	Data *struct {
		ThreadedConversationWithInjectionsV2 *struct {
			Instructions []Instruction `json:"instructions"`
		} `json:"threaded_conversation_with_injections_v2"`
		User *struct {
			Result struct {
				Legacy Legacy `json:"legacy"`
			} `json:"result"`
		} `json:"user,omitempty"`
	} `json:"data"`
}

type Instruction struct {
	Entries []struct {
		EntryID string `json:"entryId"`
		Content struct {
			ItemContent struct {
				TweetResults struct {
					Result `json:"result"`
				} `json:"tweet_results"`
			} `json:"itemContent"`
		} `json:"content"`
	} `json:"entries,omitempty"`
}
type Result struct {
	Typename string `json:"__typename"`
	Tweet    struct {
		Legacy Legacy `json:"legacy"`
	} `json:"tweet"`
	Legacy Legacy `json:"legacy"`
}

type Legacy *struct {
	FullText         string `json:"full_text"`
	ExtendedEntities struct {
		Media []Media `json:"media"`
	} `json:"extended_entities"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Entities    struct {
		Description struct {
			Urls []interface{} `json:"urls"`
		} `json:"description"`
	} `json:"entities"`
	Verified       bool   `json:"verified"`
	FollowersCount int    `json:"followers_count"`
	FriendsCount   int    `json:"friends_count"`
	StatusesCount  int    `json:"statuses_count"`
	Location       string `json:"location"`
}
type Media struct {
	DisplayURL           string `json:"display_url"`
	ExpandedURL          string `json:"expanded_url"`
	Indices              []int  `json:"indices"`
	MediaURLHTTPS        string `json:"media_url_https"`
	Type                 string `json:"type"`
	URL                  string `json:"url"`
	ExtMediaAvailability struct {
		Status string `json:"status"`
	} `json:"ext_media_availability"`
	Sizes struct {
		Large Size `json:"large"`
		Thumb Size `json:"thumb"`
	} `json:"sizes"`
	OriginalInfo struct {
		Height     int   `json:"height"`
		Width      int   `json:"width"`
		FocusRects []any `json:"focus_rects"`
	} `json:"original_info"`
	VideoInfo struct {
		AspectRatio    []int `json:"aspect_ratio"`
		DurationMillis int   `json:"duration_millis"`
		Variants       []struct {
			Bitrate     int    `json:"bitrate,omitempty"`
			ContentType string `json:"content_type"`
			URL         string `json:"url"`
		} `json:"variants"`
	} `json:"video_info"`
}
type Size struct {
	H      int    `json:"h"`
	W      int    `json:"w"`
	Resize string `json:"resize"`
}
