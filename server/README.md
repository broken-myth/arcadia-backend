# Server

The server utilises REST API which uses models, routers and controllers to handle requests. Entity refers to a set of model, router and controller. For example, the `User` entity has a `User` model, `user_router.go` router, and controllers present in `controller/user` folder.

## How to Create a Model:

1. Make sure you have a model file in the `server/models` folder. If you don't, create one. The name of the model file should be `nameOfEntity.go` (e.g. `User.go`).

1. Create the model according to the [GORM documentation](https://gorm.io/docs/models.html).

1. Add the model to the `server/models/migrate_db.go` file to make sure it is migrated when the backend starts.

## How to Create a Router:

1. Make sure you have a router file in the `server/router` folder. If you don't, create one. The name of the router file should be `<entity>_router.go` (e.g. `user_router.go`).

1. Add the path and associated controller function to the `<entity>_router.go` file in the `server/router` folder. Seperate the paths by Private and Public by using the `Auth` middleware.

## How to Create a Controller:

1. Make sure you have a controller file in the `server/controllers/<entity>` folder. If you don't, create one. The name of the controller function should be `nameOfControllerNameOfEntity<METHOD>` (e.g. `SignupUserPOST` or `GetLeaderboardGET`)

1. Add the path and associated controller function to the `<entity>_router.go` file in the `server/router` folder
