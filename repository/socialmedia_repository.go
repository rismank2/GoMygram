package repository

import (
	"MyGram/models"
	"database/sql"
	"fmt"
)

func SocmedGetRepo(db *sql.DB) []*models.ResponseGetSocialMedia {
	sqlQuery := `
		select distinct on (sm.id) sm.id, sm.name, sm.social_media_url, sm.userid,
   		u.created_at, u.updated_at, u.id, u.username, p.photo_url 
   		from public.socialmedia sm left join public.users u on sm.userid = u.id
   		left join public.photo p on u.id = p.user_id  `
	rows, err := db.Query(sqlQuery)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	socialmedias := []*models.ResponseGetSocialMedia{}
	for rows.Next() {
		var socialmedia models.ResponseGetSocialMedia
		if serr := rows.Scan(&socialmedia.Id, &socialmedia.Name, &socialmedia.Social_Media_Url, &socialmedia.User_id,
			&socialmedia.CreatedAt, &socialmedia.UpdatedAt, &socialmedia.User.Id, &socialmedia.User.Username,
			&socialmedia.User.Url); serr != nil {
			fmt.Println("Scan error", serr)
		}
		socialmedias = append(socialmedias, &socialmedia)
	}
	return socialmedias
}

func SocmedPostRepo(db *sql.DB, newSocialMedia models.SocialMedia, user_id int) models.ResponsePostSocialMedia {
	sqlQuery := `insert into public.socialmedia (name,social_media_url,userid) values ($1,$2,$3) Returning id`
	// intId, err := strconv.Atoi(id)
	err = db.QueryRow(sqlQuery, newSocialMedia.Name, newSocialMedia.Social_Media_Url, user_id).Scan(&newSocialMedia.Id)
	if err != nil {
		fmt.Println(err)
	}
	response := models.ResponsePostSocialMedia{}
	sqlQuery1 := `select s.id,s.name,s.social_media_url,s.userid,u.created_at 
			from public.socialmedia s left join public.users u on s.userid = u.id where s.id = $1`
	err = db.QueryRow(sqlQuery1, newSocialMedia.Id).Scan(&response.Id, &response.Name,
		&response.Social_Media_Url, &response.User_id, &response.CreatedAt)
	if err != nil {
		panic(err)
	}
	return response
}

func SocmedPutRepo(db *sql.DB, SocialMediap models.SocialMedia, id string) models.ResponsePutSocialMedia {
	sqlQuery := `update public.socialmedia set name = $1, social_media_url = $2 where id = $3`
	//query.scan
	_, err = db.Exec(sqlQuery,
		SocialMediap.Name,
		SocialMediap.Social_Media_Url,
		id,
	)
	if err != nil {
		fmt.Println("error update")
		panic(err)
	}
	response := models.ResponsePutSocialMedia{}
	sqlQuery1 := `select s.id,s.name, s.social_media_url, s.userid, u.updated_at 
	from public.socialmedia s left join public.users u on s.userid = u.id where s.id = $1`
	err = db.QueryRow(sqlQuery1, id).
		Scan(&response.Id, &response.Name, &response.Social_Media_Url, &response.User_id, &response.UpdatedAt)
	// count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	return response
}

func SocmedDelRepo(db *sql.DB, Id string) models.Message {
	sqlQuery := `DELETE from public.socialmedia where id = $1`
	_, err := db.Exec(sqlQuery, Id)

	if err != nil {
		panic(err)
	}
	response := models.Message{
		Message: "Your socialmedia has been successfully deleted",
	}
	return response
}
