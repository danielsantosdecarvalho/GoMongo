package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Meu teste Struct
type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func main() {
	// Configuração do cliente MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping no servidor MongoDB
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Selecionando o banco de dados e a coleção
	collection := client.Database("mydatabase").Collection("posts")

	// Criando um post de exemplo
	post := Post{
		ID:    3,
		Title: "Exemplo de post 3",
		Body:  "Conteúdo do post 3",
	}

	// Inserindo o post no banco de dados
	_, err = collection.InsertOne(context.Background(), post)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Post inserido com sucesso!")
}
