# sheetdb-modeler

作成したモデル（Goの構造体）を元に、CRUDのファンクションなどを自動生成するツール。

## Installation

```bash
go get github.com/takuoki/sheetdb/tools/sheetdb-modeler
```

## How to write models

### Type

* 下記の型を構造体フィールドの型として使用可能です。
  * string
  * bool, *bool
  * int, int8, int16, int32, int64, *int, *int8, *int16, *int32, *int64
  * uint, uint8, uint16, uint32, uint64, *uint, *uint8, *uint16, *uint32, *uint64
  * float32, float64, *float32, *float64
  * sheetdb.Date, sheetdb.Datetime, *sheetdb.Date, *sheetdb.Datetime
* Null許容のフィールドはポインタ指定してください。
* ユーザ定義型を使用する場合は、対応するNewXXX関数を作成してください。
  * Sex -> func NewSex(sex string) (Sex, error)
* 現状、外部パッケージのユーザ定義型は使用できません。

### Primary key

* AnnotationTagでPrimaryKeyがどれかを指定する
* 子テーブルの場合は親テーブルのキーを含め複数指定する（順番通りに指定すること）
* generate時、親テーブルの主キーが子テーブルの主キーに含まれていなければエラー
* generate時、子テーブルの主キーは、親テーブルの主キーの数プラス1でなければエラー
* generate時、primaryがNull許可（ポインタ型）の場合はエラー

### Auto numbering of ID

* 主キーの内、フィールド名がIDで終わり、かつ型がintの場合、自動採番とする（AddXXXの引数に含めない）
  * 最上位のテーブルの場合、自動採番されるIDは行番号＋採番初期値-1となる
  * 最上位以外のテーブルの場合、自動採番されるIDは同じ親テーブル内のIDの最大値+1となる（削除されたデータとの重複があり得る）（同じ親の中で最初の場合は採番初期値）

### Constraints

* non-null constraint
  * 型がstringの場合、空文字を許容するかどうかを指定する（primaryとの同時指定不可）
  * 型がstring以外の場合は、nullを許容する場合はポインタ型で定義する
* unique constraint
  * unique制約（複合は不可）が必要な場合は指定する（ひとまずstringのみ）

## How to generate models

* 親テーブル、子テーブルの関係は、generateコマンドの`-parent`, `-children`オプションで指定する
