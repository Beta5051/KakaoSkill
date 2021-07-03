package kakaoskill

import "encoding/json"

type Payload struct {
	Intent struct {
		Id string `json:"id"`
		Name string `json:"name"`
		Extra struct {
			Knowledge struct {
				MatchedKnowledges []struct {
					Answer string `json:"answer"`
					Question string `json:"question"`
					Categories string `json:"categories"`
					LandingUrl string `json:"landingUrl"`
					imageUrl string `json:"imageUrl"`
				} `json:"matched_knowledges"`
			} `json:"knowledge"`
		} `json:"extra"`
	}
	UserRequest struct {
		Timezone string `json:"timezone"`
		Block struct {
			Id string `json:"id"`
			Name string `json:"name"`
		}
		Utterance string `json:"utterance"`
		Lang string `json:"lang"`
		User struct {
			Id string `json:"id"`
			Type string `json:"type"`
			Properties struct {
				PlusfriendUserKey string `json:"plusfriendUserKey"`
				AppUserId string `json:"appUserId"`
				IsFriend bool `json:"isFriend"`
			} `json:"properties"`
		}
	} `json:"userRequest"`
	Bot struct {
		Id string `json:"id"`
		Name string `json:"name"`
	}
	Action struct {
		Id string `json:"id"`
		Name string `json:"name"`
		Params map[string]interface{} `json:"params"`
		DetailParams map[string]struct {
			Origin interface{} `json:"origin"`
			Value interface{} `json:"value"`
			GroupName string `json:"groupName"`
		} `json:"detailParams"`
		ClientExtra map[string]interface{} `json:"clientExtra"`
	} `json:"action"`
}

type Response struct {
	Version string `json:"version"`
	Template *Template `json:"template,omitempty"`
	Context *ContextControl `json:"context,omitempty"`
	Data map[string]interface{} `json:"data,omitempty"`
}

func (response Response) Encode() ([]byte, error) {
	return json.Marshal(response)
}

func (response Response) EncodeToString() (string, error) {
	bs, err := response.Encode()
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

type Template struct {
	Outputs []Component `json:"outputs"`
	QuickReplies []QuickReplies `json:"quickReplies,omitempty"`
}

type ContextControl struct {
	Values []ContextValue `json:"values"`
}

type ContextValue struct {
	Name string `json:"name"`
	LifeSpan int `json:"lifeSpan"`
	Params map[string]string `json:"params,omitempty"`
}

type Component struct {
	SimpleText *SimpleText `json:"simpleText,omitempty"`
	SimpleImage *SimpleImage `json:"simpleImage,omitempty"`
	BasicCard *BasicCard `json:"basicCard,omitempty"`
	CommerceCard *CommerceCard `json:"commerceCard,omitempty"`
	ListCard *ListCard `json:"listCard,omitempty"`
	Carousel *Carousel `json:"carousel,omitempty"`
}

type SimpleText struct {
	Text string `json:"text"`
	Forwardable bool `json:"forwardable,omitempty"`
}

type SimpleImage struct {
	ImageUrl string `json:"imageUrl"`
	AltText string `json:"altText"`
	Forwardable bool `json:"forwardable,omitempty"`
}

type BasicCard struct {
	Title string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Thumbnail Thumbnail `json:"thumbnail"`
	Buttons []Button `json:"buttons,omitempty"`
	Forwardable bool `json:"forwardable,omitempty"`
}

type CommerceCard struct {
	Description string `json:"description"`
	Price int `json:"price"`
	Currency string `json:"currency"`
	Discount int `json:"discount,omitempty"`
	DiscountRate int `json:"discountRate,omitempty"`
	DiscountPrice int `json:"discountPrice,omitempty"`
	Thumbnails []Thumbnail `json:"thumbnails"`
	Profile *Profile `json:"profile,omitempty"`
	Buttons []Button `json:"buttons"`
}

type ListCard struct {
	Header ListItem `json:"header"`
	Items []ListItem `json:"items"`
	Buttons []Button `json:"buttons,omitempty"`
	Forwardable bool `json:"forwardable,omitempty"`
}

type ListItem struct {
	Title string `json:"title"`
	Description string `json:"description,omitempty"`
	ImageUrl string `json:"imageUrl,omitempty"`
	Link *Link `json:"link,omitempty"`
}

type Carousel struct {
	Type string `json:"type"`
	Items []interface{} `json:"items"`
	Header *CarouselHeader `json:"header,omitempty"`
}

type Thumbnail struct {
	ImageUrl string `json:"imageUrl"`
	Link *Link `json:"link,omitempty"`
	FixedRatio bool `json:"fixed_ratio,omitempty"`
	Width int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
}

type Button struct {
	Label string `json:"label"`
	Action string `json:"action"`
	WebLinkUrl string `json:"webLinkUrl,omitempty"`
	MessageText string `json:"messageText,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	BlockId string `json:"blockId,omitempty"`
	Extra map[string]interface{} `json:"extra,omitempty"`
}

type Link struct {
	PC string `json:"pc,omitempty"`
	Mobile string `json:"mobile,omitempty"`
	Web string `json:"web,omitempty"`
}

type CarouselHeader struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Thumbnail Thumbnail `json:"thumbnail"`
}

type Profile struct {
	Nickname string `json:"nickname"`
	ImageUrl string `json:"image_url,omitempty"`
}

type QuickReplies struct {
	Label string `json:"label"`
	Action string `json:"action"`
	MessageText string `json:"messageText"`
	BlockId string `json:"blockId"`
	Extra interface{} `json:"extra"`
}