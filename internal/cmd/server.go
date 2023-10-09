package cmd

import (
	"net/http"
	"notes_server/internal/database"
	"notes_server/internal/domain/note"
	"notes_server/internal/router"
)

// func loop(router *gin.Engine, lifeCycle fx.Lifecycle) {
// 	lifeCycle.Append(fx.Hook{
// 		OnStart: func(ctx context.Context) error {
// 			ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
// 			defer stop()
// 			select {
// 			case <-ctx.Done():
// 				// address := viper.GetString("server.address")
// 				port := viper.GetString("server.port")
// 				if err := router.Run(":" + port); err != nil {
// 					return err
// 				}

// 			}

// 			return nil
// 		},
// 		OnStop: func(context.Context) error {
// 			return nil
// 		},
// 	})
// }

func StartServer() error {
	// server := fx.New(
	// 	fx.Provide(logger.LoggerProvider),
	// 	fx.Provide(database.DatabaseProvider),
	// 	fx.Provide(note.NoteProvider),
	// 	fx.Provide(router.RouterProvider),
	// 	fx.Invoke(loop),
	// )

	// server.Run()

	// if err := server.Err(); err != nil {
	// 	return err
	// }
	db, err := database.DatabaseProvider()
	if err != nil {
		return err
	}
	note := note.NoteProvider(db)
	router := router.RouterProvider(note)
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()
	return nil
}
