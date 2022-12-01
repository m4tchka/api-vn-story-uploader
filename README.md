# vn-story-uploader

This is an API for creating, fetching and deleting scenes for my other project, _react-visual-novel_, with scenes created using another repo, _cli-vn-dialogue-maker_.

This API is written in Go using the gorilla/mux package, and connects to a MongoDB database.

The intent of this API is to learn how to create a basic API in Go, as well as how to use MongoDB.

| Path         | Method | Params  | Response                                                                       |
| ------------ | ------ | ------- | ------------------------------------------------------------------------------ |
| /            | GET    | None    | Homepage                                                                       |
| /articles    | GET    | None    | Several test articles                                                          |
| /scenes      | GET    | None    | All Scenes in the database collection                                          |
| /scenes/{id} | GET    | sceneId | Specific scene with corresponding id from the database collection if it exists |
| /scenes      | POST   | None    | Newly posted scene from the body of the request                                |
| /scenes/{id} | DELETE | sceneId | Specific scene with corresponding id that was deleted/ confirmation message    |
