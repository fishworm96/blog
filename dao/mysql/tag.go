package mysql

func CreateTag(name string) (err error) {
	sqlStr := `insert into tag(tag_name) values (?)`
	_, err = db.Exec(sqlStr, name)
	return err
}