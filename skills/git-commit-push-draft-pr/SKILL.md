---
name: git-commit-push-draft-pr
description: git の変更内容を安全に確認し、意図の明確な commit を作成し、正しいブランチを push して GitHub の Draft PR まで作成するための skill。`commitしてpushしてdraft PRまで作って`、`変更を公開して下書きPRを作って`、`pushしたあとドラフトPRを作成して` のように、ローカル変更の公開と PR 作成まで一気に進めたいときに使う。無関係な差分の巻き込みや危険な git 操作を避けながら進める。
---

# Git Commit Push Draft PR

## 概要

この skill は、ローカルの変更を確認して commit を作り、branch を push し、GitHub に Draft PR を作成するところまでを安全に進めるためのものです。何を commit したか、どの branch を push したか、どの PR を作成したかが短く分かる状態で終えることを目指します。

## 前提確認

最初に次を確認します。

- リポジトリが git 管理されていること
- 現在の branch が想定どおりであること
- working tree に無関係な差分が混ざっていないこと
- GitHub へ push できる remote が設定されていること
- `gh` CLI が利用可能で、GitHub 認証が通っていること

`gh` が使えない場合は、Draft PR の自動作成は進めません。その場合は commit と push まで実施するか、そこで止めるかをユーザーに確認します。

## ワークフロー

1. リポジトリと差分の状態を確認する。
2. commit 対象ファイルを明確にする。
3. 必要な確認を行って commit を作成する。
4. 現在の branch を remote に push する。
5. base branch を確認して Draft PR を作成する。
6. commit hash、push 先、PR URL を共有する。

## 状態確認

まず次のような非対話的なコマンドで状況を確認します。

```powershell
git status --short --branch
git diff --stat
git diff
git diff --cached
git remote -v
git branch --show-current
```

差分が 1 つの commit にまとまるか、分けるべきかを判断できるだけの文脈を読み取ります。ユーザーが対象範囲を指定している場合は、その範囲を優先します。

## 対象範囲の決定

広くまとめて stage するより、明示的に対象を選んで stage することを優先します。

```powershell
git add -- path/to/file
git add -- path/to/another-file
```

無関係な差分がある可能性があるときは、`git add .` や `git commit -a` を避けます。working tree に別件の変更が見える場合は、commit 前に対象範囲を確認します。

## Commit 前の確認

commit 前に次を確認します。

- 含めるべきなのに unstaged のままのファイル
- 誤って stage された生成物、秘密情報、ローカル設定
- 短時間で確認できて関連性の高いテストや linter の失敗
- conflict、rebase 中、detached HEAD のような状態
- push 先や upstream 設定の不足

軽く実行できて関連性が明確な検証は、commit 前に実施します。未実施の検証がある場合は最後に明示します。

## Commit の作成

commit message は作業手順ではなく、ユーザー影響やコード上の変更内容を表すものにします。短い命令形の subject line を優先します。

```powershell
git commit -m "your message"
```

何も stage されていない、hook が失敗する、git identity が未設定といった理由で commit できない場合は、原因を確認して慎重に対応します。ユーザーが明示的に依頼しない限り、`--no-verify` は使いません。

## Push

現在の branch と upstream 設定を確認したうえで push します。

```powershell
git push
```

upstream が未設定なら次を使います。

```powershell
git push -u origin <branch-name>
```

ユーザーが明示的に依頼し、影響を理解している場合を除いて force push はしません。

## Draft PR の作成

Draft PR を自動作成する前に、次を確認します。

- `gh` CLI が利用可能か
- GitHub 認証が通っているか
- PR の base branch が適切か
- title と body が変更内容に合っているか

base branch が不明な場合は、既定 branch や運用ルールを確認します。判断材料がなければユーザーに確認します。

`gh` が使える場合は、たとえば次のように Draft PR を作成します。

```powershell
gh pr create --draft --base <base-branch> --head <head-branch> --title "<title>" --body "<body>"
```

PR title は commit message や変更内容をもとに簡潔にまとめます。PR body には、少なくとも次を含めます。

- 何を変更したか
- なぜ変更したか
- 確認したこと
- 未確認のことや残るリスク

`gh` が使えない、認証されていない、または repository 権限が不足している場合は、Draft PR の自動作成は行わず、その理由を明確に伝えます。

## 結果共有

結果として次を共有します。

- commit hash と subject line
- push した branch
- push 先の remote
- 作成した Draft PR の URL
- 実行した確認やテスト
- 残っているリスクや未実施の検証

## 安全ルール

`git reset --hard`、`git checkout --`、force push のような破壊的な git 操作を、「とりあえず通すため」に使ってはいけません。その操作自体をユーザーが明示的に依頼した場合だけ検討します。

リポジトリが dirty なときに、変更された全ファイルを commit 対象だと決めつけません。

履歴の書き換えを黙って行いません。公開済みの履歴に影響する操作が必要なら、いったん止まって確認します。

非対話的なコマンドを優先します。ユーザーが望んでいない限り、interactive staging や commit editor は避けます。

`gh` CLI が未導入または未認証のときは、Draft PR 作成ができる前提で進めたふりをしません。どこまで実行できたかを正確に共有します。

## 想定トリガー

たとえば次のような依頼で使います。

- "Commit, push, and open a draft PR"
- "変更を commit して push して draft PR まで作って"
- "このブランチを公開してドラフト PR を作成して"
- "対象ファイルだけ stage して commit、push、Draft PR まで進めて"
