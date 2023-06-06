package users

type User struct {
	Id            int    `json:"id"`
	Nome          string `json:"nome"`
	Sobrenome     string `json:"sobrenome"`
	Email         string `json:"email"`
	Idade         int    `json:"idade"`
	Altura        int    `json:"altura"`
	Ativo         bool   `json:"ativo"`
	DataDeCriacao string `json:"data_de_criacao"`
}

var users []User
var lastID int

type Repository interface {
	GetAll() ([]User, error)
	Create(id int, nome, sobrenome, email string, idade, altura int, ativo bool, DataDeCriacao string) (User, error)
	LastID() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]User, error) {
	return users, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Create(id int, nome, sobrenome, email string, idade, altura int, ativo bool, DataDeCriacao string) (User, error) {
	u := User{id, nome, sobrenome, email, idade, altura, ativo, DataDeCriacao}
	users = append(users, u)
	lastID = u.Id
	return u, nil
}
