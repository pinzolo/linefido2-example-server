package main

var store = make(map[string]RegistrationSession)

type RegistrationSession struct {
	Id    string
	Login string
}

func SaveSession(session RegistrationSession) {
	store[session.Id] = session
}

func GetSession(sessionId string) RegistrationSession {
	return store[sessionId]
}

func RemoveSession(sessionId string) {
	delete(store, sessionId)
}
