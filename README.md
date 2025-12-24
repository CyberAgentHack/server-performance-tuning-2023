# server-performance-tuning-2023

## 実行方法

### Docker を使用する場合（推奨）
```shell
# MySQL と Redis を起動
$ docker-compose up -d

# アプリケーションを起動
$ make run-local

# 停止する場合
$ docker-compose down
```

### Docker を使用しない場合
```shell
$ make setup
$ make run-local
```