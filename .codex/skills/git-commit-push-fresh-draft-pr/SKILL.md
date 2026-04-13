---
name: git-commit-push-fresh-draft-pr
description: Safely review local git changes, make a focused commit, push to a fresh remote branch, and create a new GitHub Draft PR without reusing an existing open PR. Use when the user wants commit/push/Draft PR help and the work must not be pushed onto a branch that already has an open PR, or when branch isolation matters for follow-up work.
---

# Git Commit Push Fresh Draft PR

## Overview

Use this skill when the user wants to publish changes as a brand-new Draft PR instead of updating an existing open PR. The core rule is simple: never keep using a head branch that already has an open PR unless the user explicitly asks for that PR to be updated.

## Workflow

1. Review the repository state before making changes.
2. Check whether the current branch already has an open PR.
3. If it does, create and switch to a fresh branch from the current HEAD before pushing anything.
4. Stage only the intended files, create one focused commit, and push the fresh branch.
5. Create a new Draft PR from that fresh branch.
6. Report the new branch name, commit hash, and Draft PR URL.

## Preflight Checks

Confirm the following before acting:

- The repository is a git repository.
- The working tree contains the changes the user intends to publish.
- A writable remote exists for push.
- `gh` CLI is installed and authenticated if Draft PR creation is required.

Use non-interactive checks like:

```powershell
git status --short --branch
git diff --stat
git diff
git diff --cached
git remote -v
git branch --show-current
gh pr status
gh auth status
```

If `gh` is unavailable or unauthenticated, stop after explaining the blocker. Do not pretend the Draft PR step succeeded.

## Branch Isolation Rule

Treat an existing open PR on the current branch as a stop sign for direct push.

- If the current branch already has an open PR, do not push new commits to that branch.
- If the current branch is already tracking a remote branch used by an open PR, do not reuse it for this workflow.
- If the user explicitly asks to update the existing PR, this skill is the wrong skill; use the normal draft-PR workflow instead.

Prefer to create a fresh branch even when unsure. A new branch is safer than accidentally mutating an in-flight review.

## Creating the Fresh Branch

When branch isolation is needed, create a new branch from the current HEAD before push. Reuse the current worktree contents; do not discard or rewrite user changes.

Suggested naming patterns:

```text
<current-branch>-followup
<current-branch>-draft
<ticket-or-feature>-v2
<ticket-or-feature>-followup-YYYYMMDD
```

Use a short, descriptive name. If a candidate branch already exists locally or remotely, generate another unique name instead of overwriting anything.

Example:

```powershell
git switch -c <new-branch-name>
```

If already on an isolated branch with no open PR tied to it, keep using it.

## Staging and Commit Rules

Stage only the files intended for this publish step. Avoid broad staging commands when unrelated edits may be present.

```powershell
git add -- path/to/file
git add -- path/to/another-file
git commit -m "your concise commit message"
```

Before committing, verify there are no unrelated unstaged or staged changes that would make the commit misleading. If the tree is mixed and the intended subset is unclear, stop and ask the user rather than guessing.

## Push and Draft PR Creation

Push the fresh branch with upstream configured.

```powershell
git push -u origin <new-branch-name>
```

Create a Draft PR that points from the fresh branch to the intended base branch.

```powershell
gh pr create --draft --base <base-branch> --head <new-branch-name> --title "<title>" --body "<body>"
```

Before creating the PR, confirm:

- `base` is correct.
- `head` is the fresh branch, not the old PR branch.
- The title matches the commit or feature intent.

If `gh pr status` or `gh pr list --head <current-branch>` shows an open PR for the original branch, that is expected. Still create the new Draft PR from the fresh branch.

## Safety Rules

Never:

- push additional commits to a branch with an existing open PR unless the user explicitly asks
- force-push to escape a branch naming or PR reuse problem
- use `git reset --hard` or `git checkout --` as part of this workflow
- use interactive git flows that are likely to hang the session

If the repository is dirty in a way that makes safe isolation unclear, explain the risk and pause rather than guessing.

## Report Back

Always include:

- the commit hash and subject
- the branch that was pushed
- whether a fresh branch was created because of an existing open PR
- the new Draft PR URL
- any blocker that prevented completion

## Example Triggers

Use this skill for prompts like:

- "Commit these changes, but do not add to the existing PR. Open a new draft PR."
- "Publish this as a separate draft PR from the current branch."
- "There is already a PR for this branch. Make a new branch and create another draft PR."
- "Commit and push this follow-up work without reusing the old PR."
