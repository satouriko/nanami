package haruka

// [person]user_id
type Person struct {
	Name	string `redis:"name"`
	Familiarity int64 `redis:"familiarity"`
	Status	string `redis:"status"`
}

// [memo]chat_id:memo-incr
// [memo-arch]chat_id:memo-incr
// [tag]tag@chat_id:memo-incr
// [tag-arch]tag@chat_id:memo-incr
type Memo []string

// [memo-incr]
type MemoIncr int

// [memo-detail]memo-incr
type MemoDetail struct {
	Content	string	`redis:"content"`
}

// [tags]memo-incr:tag
type Tags []string