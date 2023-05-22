package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	. "main/models"
	"strings"
	"time"
)

type Context struct {
	Db *sql.DB
}

func InitContext() Context {
	log.Println("Initiating the database connection")
	// dsn := fmt.Sprintf("go_root:%s@tcp(db4free.net:3306)/go_microservice?charset=utf8&parseTime=True", "Password!123")
	// log.Println(dsn)
	Db, err := sql.Open("mysql", "go_root:Password!123@tcp(db4free.net:3306)/go_microservice?parseTime=True")
	if err != nil {
		log.Print(err)
		log.Fatalln("Unable to connect to the database. Closing service...")
	}

	// err = Db.Ping()
	// if err != nil {
	// 	log.Fatalln("Unable to connect to the database. Closing service...")
	// }
	return Context{Db}
}

func (ctx Context) GetAllBlogs() ([]Blog, error) {
	var BlogList []Blog
	query := "SELECT * FROM tbl_blogs"

	res, err := ctx.Db.Query(query)

	if err != nil {
		return BlogList, err
	}
	defer res.Close()
	for res.Next() {
		var b Blog
		err := res.Scan(&b.Id, &b.Title, &b.Description, &b.Body, &b.Created_At, &b.Updated_At)
		if err != nil {
			return BlogList, err
		}
		BlogList = append(BlogList, b)
	}

	return BlogList, nil
}

func (ctx Context) GetBlogByID(id int64) (Blog, error) {
	query := "SELECT * FROM tbl_blogs WHERE id = ?"

	res := ctx.Db.QueryRow(query, id)
	//if err != nil {
	//	log.Print(err)
	//	return Blog{}, err
	//}
	//defer res.Close()
	var b Blog

	err := res.Scan(&b.Id, &b.Title, &b.Description, &b.Body, &b.Created_At, &b.Updated_At)

	if err != nil {
		log.Println(err)
		return Blog{}, err
	}

	return b, nil

}

func (ctx Context) CreateBlog(b Blog) error {
	query := "INSERT INTO tbl_blogs (title, description, body, created_at, updated_at) VALUES (?,?,?,?,?)"

	//if b.Title =
	res, err := ctx.Db.Exec(query, b.Title, b.Description, b.Body, b.Created_At, b.Updated_At)
	if err != nil {
		return err
	}
	log.Println(res.LastInsertId())
	return nil
}

func (ctx Context) UpdateBlog(p BlogPatch, id int64) error {

	query := "UPDATE tbl_blogs SET "

	query_params := []string{}
	if p.Title != nil {
		query_params = append(query_params, fmt.Sprintf("title='%s'", *p.Title))
	}

	if p.Description != nil {
		query_params = append(query_params, fmt.Sprintf("description='%s'", *p.Description))
	}

	if p.Body != nil {
		query_params = append(query_params, fmt.Sprintf("body='%s'", *p.Body))
	}

	if len(query_params) > 0 {
		query_params = append(query_params, fmt.Sprintf("updated_at='%v'", time.Now().Format("2006-01-02 15:04:05")))
	}

	query += fmt.Sprintf("%s WHERE id = %d", strings.Join(query_params, ", "), id)
	log.Println(query)

	res, err := ctx.Db.Exec(query)

	if err != nil {
		return err
	} else {
		r, _ := res.RowsAffected()
		if r == 0 {
			return errors.New("404 not Found")
		}
		log.Println(res.RowsAffected())
		return nil
	}
}

func (ctx Context) DeleteBlogByID(id int64) error {
	query := "DELETE FROM tbl_blogs WHERE id = ?"

	res, err := ctx.Db.Exec(query, id)
	if err != nil {
		log.Print(err)
		return err
	}
	r, _ := res.RowsAffected()

	if r == 0 {
		return errors.New("404 not found")
	}
	//defer res.Close()
	return nil
}
