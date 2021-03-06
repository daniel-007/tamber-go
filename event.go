package tamber

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type EventParams struct {
	User     string   `json:"user"`     // required
	Item     string   `json:"item"`     // required
	Behavior string   `json:"behavior"` // required
	Value    *float64 `json:"value,omitempty"`
	Hit      *bool    `json:"hit,omitempty"`
	Context  []string `json:"context,omitempty"`
	Created  *int64   `json:"created,omitempty"`
	// GetRecs can only be set when making event track requests.
	GetRecs *DiscoverParams `json:"get_recs,omitempty"`
}

type EventRetrieveParams struct {
	User, Item, Behavior        *string
	CreatedSince, CreatedBefore *int64
	Number                      *int // default:200 | max:500
}

type Event struct {
	User     string  `json:"user"`
	Item     string  `json:"item"`
	Behavior string  `json:"behavior"`
	Value    float64 `json:"value"`
	Created  int64   `json:"created"`
	Object   string  `json:"object"`
}

type EventResult struct {
	Events []Event      `json:"events"`
	Recs   *[]Discovery `json:"recommended,omitempty"`
}

type EventResponse struct {
	Succ   bool        `json:"success"`
	Result EventResult `json:"result"`
	Error  string      `json:"error"`
	Time   float64     `json:"time"`
	ResponseInfo
}

func (r *EventResponse) SetInfo(info ResponseInfo) {
	r.ResponseInfo = info
}

func (params *EventParams) AppendToBody(v *url.Values) {
	if len(params.User) > 0 {
		v.Add("user", params.User)
	}
	if len(params.Item) > 0 {
		v.Add("item", params.Item)
	}
	if len(params.Behavior) > 0 {
		v.Add("behavior", params.Behavior)
	}

	if params.Value != nil {
		v.Add("value", strconv.FormatFloat(*(params.Value), 'f', -1, 64))
	}
	if params.Created != nil {
		v.Add("created", strconv.FormatInt(*(params.Created), 10))
	}

	getRecs, _ := json.Marshal(params.GetRecs)
	if getRecs != nil {
		v.Add("get_recs", string(getRecs))
	}
}

func (params *EventRetrieveParams) AppendToBody(v *url.Values) {
	if params.User != nil {
		v.Add("user", *params.User)
	}
	if params.Item != nil {
		v.Add("item", *params.Item)
	}
	if params.Behavior != nil {
		v.Add("behavior", *params.Behavior)
	}
	if params.CreatedSince != nil {
		v.Add("created_since", strconv.FormatInt(*params.CreatedSince, 10))
	}
	if params.CreatedBefore != nil {
		v.Add("created_before", strconv.FormatInt(*params.CreatedBefore, 10))
	}
	if params.Number != nil {
		v.Add("number", strconv.Itoa(*params.Number))
	}
}

//Batch
type EventBatchParams struct {
	Events []EventParams `json:"events"`
}

type BatchResult struct {
	Object         string `json:"object"`
	NumBatchEvents int    `json:"num_batch_events"` //events in the batch that have been tracked
	NumBatchUsers  int    `json:"num_batch_users"`  //total user count in the batch
	NumBatchItems  int    `json:"num_batch_items"`  //total item count in the batch
	NumUsersAdded  int    `json:"num_users_added"`  //users added from the batch
	NumItemsAdded  int    `json:"num_items_added"`  //items added from the batch
}

type BatchResponse struct {
	Succ   bool        `json:"success"`
	Result BatchResult `json:"result"`
	Error  string      `json:"error"`
	Time   float64     `json:"time"`
	ResponseInfo
}

func (r *BatchResponse) SetInfo(info ResponseInfo) {
	r.ResponseInfo = info
}

func (params *EventBatchParams) AppendToBody(v *url.Values) {
	events, _ := json.Marshal(params.Events)
	v.Add("events", string(events))
}
