# go-todo-app

```
cmd/
├── go-todo-app
│   └── main // entry point
pkg/
├── adapter // 外部からのリクエスト処理
│   └── controller
├── domain // ドメイン関連の処理
│   ├── model
│   ├── repository // infrastructureと依存性逆転のためのinterface
│   │              // repostitoryの中でquery用とcommand用でinterfaceを分ける
│   └── service  // 複数のmodelをまたぐ業務的なロジックを処理
├── infrastructure // DBアクセス処理
│   ├── dto
│   ├── db // DBの共通処理
│   ├── repository
└── usecase // usecase(application service) repository, domainを扱いアプリケーションの処理を行う
```

start up docker-compose  
```
$ docker-compose up
```
