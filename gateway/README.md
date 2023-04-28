gqlgenのinit
```shell
go run github.com/99designs/gqlgen init
```

gqlgenの構成を更新したい場合, schemaを更新したい場合
（-mod=modを入れないとgo.sumにモジュールが足りていないというエラーが出る https://github.com/99designs/gqlgen/issues/1483）
```shell
go run -mod=mod github.com/99designs/gqlgen generate
```
