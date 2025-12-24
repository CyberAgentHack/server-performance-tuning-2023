# server-performance-tuning-2023

## 実行方法

### Docker を使用する場合（推奨）
```shell
# MySQL と Redis を起動
$ docker-compose up -d

# テストデータを投入（オプション）
$ ./scripts/insert-test-data.sh

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

## テストデータについて

`scripts/insert-test-data.sh` を実行すると、以下のテストデータが投入されます：
- 5 genres（アニメ、ドラマ、バラエティ、映画、格闘技）
- 5 series（ONE PIECE、名探偵コナン、ドラゴンボール、NARUTO、進撃の巨人）
- 3 seasons
- 3 episodes