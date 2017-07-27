package haruka

// [person]user_id
type Person struct {
	Name	string `redis:"name"`
	Familiarity int64 `redis:"familiarity"`
	Status	string `redis:"status"`
}

// [memo]chat_id
// [memo-arch]chat_id
type Memo []string

// [memo-incr]
type MemoIncr int

// [memo-detail]memo-incr
type MemoDetail struct {
	Content	string	`redis:"content"`
	Tags 	string	`redis:"tags"`
}

// [tag]chat_id
// [tag-arch]chat_id
type Tag []string