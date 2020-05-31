package db

import (
	"log"
	"tasks/types"
)

func GetCategories(username string)[]types.CategoryCount{
	userID, err := GetUserId(username)
	if err != nil {
		return nil
	}
	stmt := "select 'UNCATEGORIZED' as name, count(1) from task where cat_id=0 union  select c.name, count(*) from   category c left outer join task t  join status s on  c.id = t.cat_id and t.task_status_id=s.id where s.status!='DELETED' and c.user_id=?   group by name    union     select name, 0  from category c, user u where c.user_id=? and name not in (select distinct name from task t join category c join status s on s.id = t.task_status_id and t.cat_id = c.id and s.status!='DELETED' and c.user_id=?)"
	rows := db.query(stmt, userID, userID, userID)
	var categories []types.CategoryCount
	var category types.CategoryCount

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&category.Name, &category.Count)
		if err != nil {
			log.Println(err)
		}
		categories = append(categories, category)
	}
	return categories
}
