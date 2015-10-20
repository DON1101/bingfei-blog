package auth

func SecretAuth(username, password string) bool {
    return username == "guojing" && password == "abc"
}
