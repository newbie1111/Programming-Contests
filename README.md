# Programming-Contests

## 目次

- [Programming-Contests](#programming-contests)
  - [目次](#目次)
  - [リポジトリについて](#リポジトリについて)
  - [ojによるテスト](#ojによるテスト)
    - [準備](#準備)
    - [入出力テスト](#入出力テスト)
    - [提出](#提出)

## リポジトリについて

競技プログラミングに関係する回答コードを、各サービスごとにディレクトリに分けて保存したリポジトリ

企業の個別テストや規約によって公開できないものは.gitignoreにより保存しないようする。
何かの拍子で公開してしまい、漏洩してしまうことを防ぐため。

## ojによるテスト

`online-judge-tools`及び`oj-prepare`のwrapperとなるシェルスクリプトを作成したため、これを使う。

- `oj-prepares.sh`:`oj-prepare`を複数コンテスト分一括で行う
- `oj-submits.sh`:`oj s`を複数ファイル一括で行う
- `oj-tests.sh`:`oj t`を複数ファイル一括で行う

### 準備

```bash
sh oj-prepares.sh [<contest URL> ...]
```

- コンテスト用のディレクトリを追加
  - コンテストそのもののディレクトリ
  - 各問題用のディレクトリ
- 回答コードの作成
  - `.py`と`.cpp`は入出力のコードが自動で生成される（たまにされない）

### 入出力テスト

```bash
sh oj-tests.sh [<submission file path> ...]
```

- それぞれのファイルパスに基づいて、`oj t`を行う
- 現状対応している言語
  - `.go`
  - `.py`
  - `.cpp`
- 対応していない場合は提出されない

### 提出

```bash
sh oj-submits.sh [<submission file path> ...]
```
