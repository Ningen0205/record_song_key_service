# Record song key service

## 概要

このサービスは以下の目的を実現するために作成されました。

1. 自分がよく歌うカラオケのキーを記録しておくため。
1. よく歌う曲のyoutubeなどへのlinkの保存
1. 自分の最高音、最低音を記録しておくため。
1. 記録した最高音、最低音を元にその曲のおすすめのキーを表示する。


## 機能の要件


1, 2, 3 を実現するためには以下の機能が必要。


- 楽曲に関する機能
  - 曲の登録機能
    - （原キーでの）最低音、最高音、曲のURL（オリジナル）
  - 曲の取得機能
  - 曲の更新機能
  - 曲の削除機能（オプション）
  - 曲のカバーの登録
    - カバー動画のURL
    - 原曲からのキーの変動
    - 曲の最低音、最高音（オプション）
- キーの記録に関する機能
  - キーのCURD
  - ユーザのCURD（認証周りは持たない）
    - 保持するべき情報
      - userId
      - ユーザの最低音
      - ユーザの最高音
  - 曲（カバー含む）とユーザ、キーの紐づけ
  - （オプション: 別のmsの責務かもしれない）楽曲の動画から、最高音、最低音を出す

## DB構造

### song(songs) table

| column_name     | type      | constraints | description      |
| --------------- | --------- | ----------- | ---------------- |
| id              | uuid      | primary key | primary_key      |
| name            | text      | not null    | 曲名             |
| name_kana       | text      | not null    | 曲名（ひらがな） |
| lowest_note_id  |           | foreign key | 最低音           |
| highest_note_id |           | foreign key | 最高音           |
| singer_id       |           | foreign key |                  |
| is_cover        | bool      | not null    |                  |
| video_link      | text      |             | mvへのリンク     |
| created_at      | timestamp |             | 作成日時         |
| updated_at      | timestamp |             | 更新日時         |
| deleted_at      | timestamp |             | 削除日時         |

### CoverSongKeys(cover_songs) table

| column_name | type      | constraints | description |
| ----------- | --------- | ----------- | ----------- |
| id          | uuid      | primary key | primary_key |
| song_id     |           | foreign key |             |
| key_id      |           | foreign key | keyの移動幅 |
| created_at  | timestamp |             | 作成日時    |
| updated_at  | timestamp |             | 更新日時    |
| deleted_at  | timestamp |             | 削除日時    |
### singer(singers) table

| column_name | type      | constraints | description |
| ----------- | --------- | ----------- | ----------- |
| id          | uuid      | primary key | primary_key |
| name        | text      | not null    | 歌手名      |
| created_at  | timestamp |             | 作成日時    |
| updated_at  | timestamp |             | 更新日時    |
| deleted_at  | timestamp |             | 削除日時    |


### note(notes) table

| column_name | type      | constraints | description |
| ----------- | --------- | ----------- | ----------- |
| id          | int       | primary key | primary_key |
| description |           | not null    | 音の高さ    |
| created_at  | timestamp |             | 作成日時    |
| updated_at  | timestamp |             | 更新日時    |
| deleted_at  | timestamp |             | 削除日時    |



### user(users) table

| column_name     | type      | constraints | description           |
| --------------- | --------- | ----------- | --------------------- |
| id              | uuid      | primary key | user_serviceのuser_id |
| lowest_note_id  |           | foreign key | 最低音                |
| highest_note_id |           | foreign key | 最高音                |
| created_at      | timestamp |             | 作成日時              |
| updated_at      | timestamp |             | 更新日時              |
| deleted_at      | timestamp |             | 削除日時              |



### user_song(user_songs) table

| column_name | type      | constraints                      | description  |
| ----------- | --------- | -------------------------------- | ------------ |
| user_id     |           | primary key foreign key not null |              |
| song_id     |           | primary key foreign key          |              |
| key_id      |           | not null                         | キーの移動幅 |
| created_at  | timestamp |                                  | 作成日時     |
| updated_at  | timestamp |                                  | 更新日時     |
| deleted_at  | timestamp |                                  | 削除日時     |


### key(keys) table

| column_name | type      | constraints     | description |
| ----------- | --------- | --------------- | ----------- |
| id          | int       | primary key     | primary_key |
| Value       | int       | not null unique | キー(-7~+7) |
| created_at  | timestamp |                 | 作成日時    |
| updated_at  | timestamp |                 | 更新日時    |
| deleted_at  | timestamp |                 | 削除日時    |

