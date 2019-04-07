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

func RepoDestroyUser(id int) {
	return
}

func RepoInsertUser(user User) User {
	return User{}
}
