package sql

const (
	Q_INSERT_USER  = `INSERT INTO users(id) VALUES ($1)`
	Q_INSERT_BADGE = `INSERT INTO badges(id, name, url, user_id) VALUES ($1, $2, $3, $4)`

	Q_SELECT_ALL_USERS  = `SELECT * from users`
	Q_SELECT_ALL_BADGES = `SELECT * from badges`

	Q_SELECT_BADGES_BY_USER_ID = `SELECT * from badges WHERE user_id=$1`

	// badges_db=# INSERT INTO badges(id,user_id,name,url) VALUES (1,1,'go', 'somewhere');

)
