# go-todo-app

```
pkg/
├── adapter // 外部からのリクエスト処理
│   └── controller
├── domain // ドメイン関連の処理
│   ├── model
│   ├── repository // infrastructureと依存性逆転のためのinterface
│   │              // repostitoryの中でquery用とcommand用でinterfaceを分ける
│   └── service  // 複数のmodelをまたぐ業務的なロジックを処理
├── errors // アプリケーション全体のエラー処理
│   └── error.go
├── infrastructure // DBアクセス処理
│   ├── dao
│   ├── db.go
│   ├── query
│   ├── repository
│   └── seeds
└── usecase // usecase(application service) repository, domainを扱いアプリケーションの処理を行う
```
