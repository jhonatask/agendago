package main

import (
    "agenda/internal/entity"
    "agenda/internal/router"
    "agenda/pkg/database"
    "log"
    "os"
)

func main() {
    db, err := database.NewPostgresDB()
    if err != nil {
        log.Fatal(err)
    }

    // Migrar a entidade Agenda
    db.AutoMigrate(&entity.Agenda{})

    r := router.SetupRouter(db)

    // Configurar proxies confiáveis
    r.SetTrustedProxies([]string{"127.0.0.1"})

    // Definir a porta a partir da variável de ambiente ou usar a porta padrão 8080
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    r.Run(":" + port)
}