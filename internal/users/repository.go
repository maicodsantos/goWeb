package users

import "fmt"

var users []User
var lastID int

type Repository interface {
	GetAll() ([]User, error)
	Create(id int, nome, sobrenome, email string, idade, altura int, ativo bool, dataDeCriacao string) (User, error)
	LastID() (int, error)
	Update(id int, nome, sobrenome, email string, idade, altura int, ativo bool, dataDeCriacao string) (User, error)
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

func (r *repository) Create(id int, nome, sobrenome, email string, idade, altura int, ativo bool, dataDeCriacao string) (User, error) {
	u := User{id, nome, sobrenome, email, idade, altura, ativo, dataDeCriacao}
	users = append(users, u)
	lastID = u.Id
	return u, nil
}

func (repository) Update(id int, nome, sobrenome, email string, idade, altura int, ativo bool, dataDeCriacao string) (User, error) {
	u := User{Nome: nome, Sobrenome: sobrenome, Email: email, Idade: idade, Altura: altura, Ativo: ativo, DataDeCriacao: dataDeCriacao} // Instância de "p" para Update
	updated := false                                                                                                                    // Atribuição false para Updated - não foi realizado nenhum update até aqui
	for i := range users {                                                                                                              // Este For percorrerá a lista dos elementos criados no array para buscar o elemento com o Id que já existe
		if users[i].Nome == nome { // Caso encontre esse Id ...
			u.Id = id      // ... o Id do novo produto será o mesmo do já existente (basicamente, o Id que passamos substituirá o já existente, só que são iguais)...
			users[i] = u   // ... e aqui, irá atualizar (neste Id), todos os valores dos elementos que enviarmos no Put...
			updated = true // ... alterando o seu status para "True"
		}
	}
	if !updated { // Caso não tenha havido esse update, ou seja, se continuar como 'false'...
		return User{}, fmt.Errorf("usuario %d não encontrado", id) // ... nos será enviada uma mensagem de erro
	}
	return u, nil // Retorno do novo produto com um erro do tipo 'nil'
}
