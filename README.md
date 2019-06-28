# sheetdb

A golang package for using Google spreadsheets as a database instead of the actual database management system.

**!!! Caution !!!**

Currently we are not measuring performance. It is intended for use in small applications where performance is not an issue.

## Features

* Load sheet data into cache
* Apply cache update information to sheet asynchronously
* Exclusive control when updating cache and inserting a row to a sheet
* Automatic generation of CRUD functions based on model (structure definition)
* Automatic numbering and initial value setting of ID
* Unique and non-null constraints
* Cascade delete child data when deleting parent data
* Notification mechanism on asynchronous update error

The following features are not included.

* SQL query
* Transaction control (commit and rollback)
* Read Lock for Update

## Installation

```bash
go get github.com/takuoki/sheetdb
```

## Requirement

This package uses Google OAuth2.0. So before executing tool, you have to prepare credentials.json.
See [Go Quickstart](https://developers.google.com/sheets/api/quickstart/go), or [blog post (Japanese)](https://medium.com/veltra-engineering/how-to-use-google-sheets-api-with-golang-9e50ee9e0abc) for the details.

## Usage

### How to write a model

```go
//go:generate sheetdb-modeler -type=User -children=Foo,Bar
type User struct {
  UserID   int           `json:"user_id"` primary key
  Name     string        `json:"name"`
  Email    string        `json:"email"` unique
  Sex      Sex           `json:"sex"`
  Birthday *sheetdb.Date `json:"birthday"`
}

//go:generate sheetdb-modeler -type=Foo -parent=User
type Foo struct {
  UserID int     `json:"user_id"` primary key
  FooID  int     `json:"foo_id"` primary
  Value  float32 `json:"value"`
  Note   string  `json:"note"` allowempty
}

//go:generate sheetdb-modeler -type=Bar -parent=User
type Bar struct {
  UserID   int              `json:"user_id"` primary key
  Datetime sheetdb.Datetime `json:"datetime"` primary
  Value    float32          `json:"value"`
  Note     string           `json:"note"` allowempty
}
```

* PrimaryKeyがどれかを指定する（複数指定可能、順番通りに指定すること）
* Null許容のフィールドはポインタ指定する
* 型がstringの場合、空文字を許容するかどうかを指定する（primary以外）
* unique制約（複合は不可）が必要な場合は指定する（ひとまずstringのみ）

* ユーザ定義型を使用する場合は、対応するNewXXX関数を作成すること
  * Sex -> func NewSex(sex string) (Sex, error)

* 親テーブル、子テーブルの関係は、generateコマンドの`-parent`, `-children`オプションで指定する
* generate時、親テーブルの主キーが子テーブルの主キーに含まれていなければエラー
* generate時、子テーブルの主キーは、親テーブルの主キーの数プラス1でなければエラー
* generate時、primaryがNull許可（ポインタ型）の場合はエラー

* 親テーブルの削除ファンクションでは、子テーブルの削除を同時に行う（Cascade Delete）
* 主キーの内、フィールド名がIDで終わり、かつ型がintの場合、自動採番とする（AddXXXの引数に含めない）
  * 最上位のテーブルの場合、自動採番されるIDは行番号＋採番初期値-1となる
  * 最上位以外のテーブルの場合、自動採番されるIDは同じ親テーブル内のIDの最大値+1となる（削除されたデータとの重複があり得る）（同じ親の中で最初の場合は採番初期値）

### How to generate from model

```bash
go generate
```

### How to set up Google spreadsheet

### Package initialization

```go
err := sheetdb.Initialize(
  ctx,
  `{"installed":{"client_id":"..."}`, // Google API credentials
  `{"access_token":"..."`,            // Google API token
  "xxxxx",                            // Google spreadsheet ID
)
```

### Load sheet data

```go
err := sheetdb.LoadData(ctx)
```

### CRUD

The functions in this section are generated automatically.

#### Create (Add/Insert)

```go
user, err := AddUser(name, email, sex, birthday)
foo, err := user.AddFoo(value, note)
bar, err := user.AddBar(datetime, value, note)
```

```go
foo, err := AddFoo(userID, value, note)
bar, err := AddBar(userID, datetime, value, note)
```

#### Read (Get/Select)

```go
user, err := GetUser(userID)
foo, err := user.GetFoo(fooID)
bar, err := user.GetBar(datetime)
```

```go
foo, err := GetFoo(userID, fooID)
bar, err := GetBar(userID, datetime)
```

get list

```go
users, err := GetUsers()
foos, err := user.GetFoos()
bars, err := user.GetBars()
```

```go
foos, err := GetFoos(userID)
bars, err := GetBars(userID)
```

filterable

```go
Foos, err := user.GetFoos(FooFilter(func(foo *Foo) bool {
  return foo.Value > 0
}))
```

sortable

```go
bars, err := user.GetBars(BarSort(func(bars []*Bar) {
  sort.Slice(bars, func(i, j int) bool {
    return bars[i].Datetime.After(bars[j].Datetime)
  })
}))
```

#### Update

```go
user, err := UpdateUser(userID, name, email, sex, birthday)
foo, err := user.UpdateFoo(fooID, value, note)
bar, err := user.UpdateBar(datetime, value, note)
```

```go
foo, err := UpdateFoo(userID, fooID, value, note)
bar, err := UpdateBar(userID, datetime, value, note)
```

#### Delete

```go
err := DeleteUser(userID)
err = user.DeleteFoo(fooID)
err = user.DeleteBar(datetime)
```

```go
err = DeleteFoo(userID, fooID)
err = DeleteBar(userID, datetime)
```

### Set up nortificaltion
