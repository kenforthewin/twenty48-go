package main

func RepoIndexUser() Users {
	var users Users
	rows, err := DB.Query("SELECT * FROM users ORDER BY top_score DESC LIMIT 50")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		user := User{}
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.CreatedAt,
			&user.Password,
			&user.TopScore)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return users
}

func RepoFindUser(id int) User {
	user := User{}
	err := DB.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(
		&user.ID,
		&user.Name,
		&user.CreatedAt,
		&user.Password,
		&user.TopScore)
	if err != nil {
		panic(err)
	}
	return user
}

func RepoInsertUser(user User) User {
	err := DB.QueryRow("INSERT INTO users(name, created_on) VALUES($1, now()) RETURNING id, created_on", user.Name).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		panic(err)
	}
	return user
}
