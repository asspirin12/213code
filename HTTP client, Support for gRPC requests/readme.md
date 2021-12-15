Я склонировал себе этот репозиторий https://github.com/grpc/grpc-go.git 

Чтобы сделать вторую гифку как тут https://www.jetbrains.com/go/whatsnew/#http-client, 
открыл grpc-go/examples/route_guide/routeguide/route_guide.proto и нажал на первую иконку слева,
там где строчка rpc GetFeature(Point) returns (Feature) {}. 
IDE откроет файл generated-requests.http с готовым запросом (надо только добавить порт 10000). 

Чтобы сделать первую гифку, можно удалить этот запрос и ввести такой же руками. 

Для третьей гифки надо запустить сервер grpc-go/examples/route_guide/server/server.go.
И запустить из generated-requests.http такой запрос: 

GRPC localhost:10000/routeguide.RouteGuide/ListFeatures

{
  "lo": {
    "latitude": 400000000,
    "longitude": -750000000
  },
  "hi": {
    "latitude": 430000000,
    "longitude": -720000000
  }
}




