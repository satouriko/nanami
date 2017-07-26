package haruka

type Person struct {
	Name	string `redis:"name"`
	Familiarity int64 `redis:"familiarity"`
	Status	string `redis:"status"`
}