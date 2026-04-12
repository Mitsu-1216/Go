---
name: git-commit-push
description: git の変更内容を安全に確認し、意図の明確な commit を作成して正しいブランチを push するための skill。`commitして`、`pushして`、`commit,pushまでやって`、`変更を公開して` のように、ローカルの git 作業を仕上げたいときに使う。無関係な差分の巻き込みや危険な git 操作を避けながら進める。
---

# Git Commit Push

## 概要

この skill は、ローカルの変更がある状態から、整理された commit を作って branch を push するところまでを安全に進めるためのものです。事故を減らしつつ、何を commit してどこへ push したかが短く分かる状態で終えることを目指します。

## ワークフロー

1. 変更前にリポジトリの状態を確認する。
2. どのファイルを commit に含めるかを明確にする。
3. commit 前に分かりやすいリスクを確認する。
4. 焦点の絞られた commit を 1 つ作る。
5. 意図した branch を push し、結果を共有する。

## 状態確認

まず working tree と branch の状態を確認します。高速で非対話的な次のようなコマンドを優先します。

```powershell
git status --short --branch
git diff --stat
git diff
git diff --cached
```

差分が 1 つの commit にまとまるか、分けるべきかを判断できるだけの文脈を読み取ります。ユーザーが対象範囲を指定している場合は、変更ファイルを全部まとめて commit するのではなく、その範囲を優先します。

## 対象範囲の決定

広くまとめて stage するより、明示的に対象を選んで stage することを優先します。

次のような対象限定のコマンドを使います。

```powershell
git add -- path/to/file
git add -- path/to/another-file
```

無関係な差分が混ざっている可能性があるときは、`git add .` や `git commit -a` を避けます。

working tree に無関係そうな変更がある場合は、commit 前にいったん止まって対象範囲を確認します。明示的に依頼されない限り、無関係な差分を revert したり discard したりしません。

## Commit 前の確認

commit 前に次を確認します。

- 含めるべきなのに unstaged のままのファイル
- 誤って stage された生成物、秘密情報、ローカル設定
- 短時間で確認できて関連性の高いテストや linter の失敗
- conflict、rebase 中、detached HEAD のような状態
- upstream branch の未設定や誤設定

軽く実行できて関連性が明確な検証は、commit 前に実施します。重い検証や、実行が難しい検証、必要性が曖昧な検証については、何を確認したか、何をしていないかを明示します。

## Commit の作成

commit message は、作業手順ではなくユーザー影響やコード上の変更内容を表すものにします。短い命令形の subject line を優先します。例:

```text
fix login redirect after session expiry
add CSV export for billing report
refactor cache invalidation for profile updates
```

そのうえで、非対話的なコマンドで commit します。

```powershell
git commit -m "your message"
```

何も stage されていない、hook が失敗する、git identity が未設定といった理由で commit できない場合は、原因を確認して慎重に対応します。ユーザーから明示的な依頼がない限り、`--no-verify` で hook を回避しません。

## Push

現在の branch と、remote branch の tracking 設定があるかを確認します。

基本は次を優先します。

```powershell
git push
```

upstream が未設定なら次を使います。

```powershell
git push -u origin <branch-name>
```

ユーザーが明示的に依頼し、影響を理解している場合を除いて force push はしません。push が拒否されたら、その先の操作に進む前に branch が behind か、remote 側に変更があるか、権限不足かを確認します。

## 結果共有

結果として次を共有します。

- commit hash と subject line
- push した branch
- 分かる場合は push 先の remote
- 実行した確認やテスト
- 残っているリスクや未実施の検証

## 安全ルール

`git reset --hard`、`git checkout --`、force push のような破壊的な git コマンドを、「とりあえず通すため」に使ってはいけません。その操作自体をユーザーが明示的に依頼した場合だけ検討します。

リポジトリが dirty なときに、変更された全ファイルを commit 対象だと決めつけません。

履歴の書き換えを黙って行いません。すでに公開済みの履歴に影響する操作が必要なら、いったん止まって確認します。

非対話的なコマンドを優先します。ユーザーが望んでいない限り、interactive staging や commit editor は避けます。

## 想定トリガー

たとえば次のような依頼で使います。

- "Commit and push my changes"
- "Please make a git commit for this fix"
- "Publish the current branch"
- "Stage only the files for this feature, then push"
