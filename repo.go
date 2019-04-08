package main

func RepoIndexUser() Users {
	return Users{}
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
