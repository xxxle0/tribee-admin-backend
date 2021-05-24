package databases

var schema = `
	CREATE TABLE user (
		full_name varchar(255)

	);
`

type User struct {
	FullName string `db:"full_name"`
}
