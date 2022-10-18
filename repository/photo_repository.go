package repository

import (
	"MyGram/models"
	"database/sql"
	"fmt"
	"time"
)

func PhotoGetRepo(db *sql.DB) []*models.ResponseGetPhoto {
	sqlQuery := `
	select p.id, p.title,p.caption, p.photo_url, p.user_id, p.created_at,
   	p.updated_at, u.email, u.username 
    from public.photo as p inner join public.users as u on p.user_id = u.id`
	rows, err := db.Query(sqlQuery)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	photos := []*models.ResponseGetPhoto{}
	for rows.Next() {
		var photo models.ResponseGetPhoto
		if serr := rows.Scan(&photo.Id, &photo.Title, &photo.Caption, &photo.Url, &photo.User_id, &photo.CreatedAt, &photo.UpdatedAt, &photo.Users.Email, &photo.Users.Username); serr != nil {
			fmt.Println("Scan error", serr)
		}
		photos = append(photos, &photo)
	}
	return photos
}

func PhotoPostRepo(db *sql.DB, newPhotos models.Photo, user_id int) models.ResponsePostPhoto {
	sqlQuery := `insert into photo
	(title,caption,photo_url,user_id,created_at,updated_at)
	values ($1,$2,$3,$4,$5,$5) Returning id`
	//query.scan
	err = db.QueryRow(sqlQuery,
		newPhotos.Title,
		newPhotos.Caption,
		newPhotos.Url,
		user_id,
		time.Now(),
	).Scan(&newPhotos.Id)
	if err != nil {
		panic(err)
	}
	response := models.ResponsePostPhoto{
		Id:        newPhotos.Id,
		Title:     newPhotos.Title,
		Caption:   newPhotos.Caption,
		Url:       newPhotos.Url,
		User_id:   int(user_id),
		CreatedAt: time.Now(),
	}
	return response
}

func PhotoPutRepo(db *sql.DB, newPhotos models.Photo, id string) models.ResponsePutPhoto {
	sqlQuery := `update public.photo set title = $1, caption = $2 , photo_url = $3, updated_at = $4 where id = $5`
	_, err = db.Exec(sqlQuery,
		newPhotos.Title,
		newPhotos.Caption,
		newPhotos.Url,
		time.Now(),
		id,
	)
	if err != nil {
		fmt.Println("error update")
		panic(err)
	}
	sqlQuery1 := `select p.id, p.title, p.caption, p.photo_url, p.user_id, p.created_at,
				p.updated_at  from public.photo as p where p.id= $1`
	err = db.QueryRow(sqlQuery1, id).
		Scan(&newPhotos.Id, &newPhotos.Title, &newPhotos.Caption, &newPhotos.Url,
			&newPhotos.User_id, &newPhotos.CreatedAt, &newPhotos.UpdatedAt)
	// count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	response := models.ResponsePutPhoto{
		Id:        newPhotos.Id,
		Title:     newPhotos.Title,
		Caption:   newPhotos.Caption,
		Url:       newPhotos.Url,
		User_id:   newPhotos.User_id,
		UpdatedAt: newPhotos.UpdatedAt,
	}
	return response
}

func PhotoDeleteRepo(db *sql.DB, id string) models.Message {
	sqlQuery := `delete from public.photo where id = $1`
	_, err := db.Exec(sqlQuery, id)
	if err != nil {
		fmt.Println("error delete")
		panic(err)
	}
	return models.Message{
		Message: "Your photo has been Successfully deleted",
	}
}
