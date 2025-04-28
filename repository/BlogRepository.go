package repository

// "blog/model"
// "database/sql"

// func AddBlog(db *sql.DB, blog model.Blog) (int64, error) {
// 	query := `INSERT INTO blogs (title, content, user_id, category_id, status) VALUES (?, ?, ?, ?, ?)`
// 	result, err := db.Exec(query, blog.Title, blog.Content, blog.UserId, blog.CategoryId, "Pending")
// 	if err != nil {
// 		return 0, err
// 	}
// 	return result.LastInsertId()
// }

// func GetAllBlogs(db *sql.DB) ([]model.Blog, error) {
// 	query := `SELECT blog_id, title, content, user_id, category_id, status, create_at, update_at FROM blogs`
// 	result, err := db.Query(query)
// 	if err != nil {
// 		return nil, err}
// 	slicesBlogs := make([]model.Blog, 0)
// 	for result.Next() {
// 		var blog model.Blog
// 		err = result.Scan(&blog.BlogId, &blog.Title, &blog.Content, &blog.UserId, &blog.CategoryId, &blog.Status, &blog.CreateAt, &blog.UpdateAt)
// 		if err != nil {
// 			return nil, err
// 		}
// 		slicesBlogs = append(slicesBlogs, blog)
// 	}
// 	return slicesBlogs, nil
// }
