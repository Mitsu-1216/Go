---
name: git-commit-push-existing-pr
description: Safely review local git changes, make a focused commit, push to the current branch, and update the existing GitHub PR tied to that branch. Use when the user wants to continue an in-flight PR instead of creating a new one, especially for follow-up fixes, review feedback, or small incremental updates on the same branch.
---

# Git Commit Push Existing PR

## Overview

Use this skill when the user wants new commits added to an already-open PR on the current branch. The main rule is to confirm that the branch is actually tied to an existing PR before pushing, rather than assuming the update target.

## Workflow

1. Review the repository state and current branch.
2. Check whether the current branch has an existing open PR.
3. If it does, stage only the intended files and create one focused commit.
4. Push the current branch to its existing upstream.
5. Report that the existing PR has been updated, along with the PR URL and commit hash.

## Preflight Checks

Confirm the following before acting:

- The repository is a git repository.
- The working tree contains the changes the user intends to publish.
- The current branch is not detached.
- A writable remote exists for push.
- `gh` CLI is installed and authenticated if PR verification is required.

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

If `gh` is unavailable or unauthenticated, explain the limitation clearly. Do not claim the branch is tied to an existing PR unless you verified it another way.

## Existing PR Rule

Treat PR verification as a required step for this workflow.

- If the current branch has an open PR, keep using that branch.
- If the current branch does not have an open PR, stop and tell the user this skill does not apply.
- If the user actually wants a separate PR, switch to the fresh-draft skill instead.

This skill is for continuing the same review thread, not splitting work into a new branch.

## Commit Scope Rules

Stage only the files intended for this PR update. Avoid broad staging commands when unrelated edits may be present.

```powershell
git add -- path/to/file
git add -- path/to/another-file
git commit -m "your concise commit message"
```

Before committing, verify there are no unrelated staged or unstaged changes that would make the PR update noisy or misleading. If the intended subset is unclear, pause and ask rather than guessing.

## Push Rules

Push to the branch that already backs the existing PR.

```powershell
git push
```

If upstream is missing but the branch is clearly meant to update an existing PR, set upstream to the matching remote branch and then push.

```powershell
git push -u origin <current-branch>
```

Do not create a new branch in this workflow unless the user changes direction.

## PR Verification

Before pushing, confirm the existing PR details:

- the PR is open
- the head branch matches the current branch
- the repository and base branch look correct

Useful checks:

```powershell
gh pr status
gh pr view --json url,number,title,headRefName,baseRefName,state
```

If the branch points to a closed or merged PR, stop and tell the user instead of reviving the branch by accident.

## Safety Rules

Never:

- create a fresh branch when the task is specifically to update the existing PR
- push to a different branch than the one verified as the PR head
- force-push unless the user explicitly asks and understands the risk
- use `git reset --hard` or `git checkout --` as part of this workflow

If PR targeting is ambiguous, pause and explain the ambiguity before pushing.

## Report Back

Always include:

- the commit hash and subject
- the branch that was pushed
- the existing PR URL that was updated
- any blocker that prevented completion

## Example Triggers

Use this skill for prompts like:

- "Commit these changes and update the current PR."
- "Push this review fix to the existing PR."
- "Add one more commit to the PR on this branch."
- "Keep using the same PR and publish this follow-up."
