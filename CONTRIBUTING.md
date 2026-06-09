# Contributing: How to Release

This document explains how to prepare and release a new version.

---

## Optional -  OpenAPI generator version bump

If you intend to bump the OpenAPI generator version as part of this release, do it as a prerequisite step.

1. Update the version in `openapitools.json`.
2. Make sure to update any custom mustache templates with newer version mustache files, keeping anything custom as it is.
3. Commit and push to `develop`.

---

## Workflows

| Workflow | Purpose |
|---|---|
| `generate-all` | Regenerates all SDK modules by pulling the latest OpenAPI specs |
| `validate-all` | Validates the generated output compiles and passes contract tests |
| `prepare-release-prs` | Analysis openapi diffs, computes version bumps, openss draft PRs |

---

## Phase 1 - Generate and Validate All

1. Go to **Actions → generate-all → Run workflow**
2. Select branch: `develop`
3. Click **Run workflow**

This regenerates every SDK module pulling its latest spec and automatically triggers `validate-all` on completion.

**If `validate-all` passes:** all modules compiled and contract tests passed. Proceed to Phase 2.

**If `validate-all` fails:** investigate the failing module before proceeding.
- If the failure is a genuine SDK breaking change (renamed method, changed signature - typically caused by a generator version upgrade), take note the affected modules. You will pass them as `force_major_modules` in Phase 2.

---

## Phase 2 - Prepare Release PRs

1. Go to **Actions → prepare-release-prs → Run workflow**
2. Select branch: `develop`
3. (Optional) Fill in the input field:

### Input: `force_major_modules`

Leave blank in normal operation - this is an escape hatch for genuine SDK breaks identified during Phase 1. List the affected modules as a comma-separated string:

```
payments,invoicing
```

Modules listed here will have their bump forced to `major` regardless of what the spec diff shows. The release PR will include a prominent warning for reviewers to inspect those modules carefully.

4. Click **Run workflow**

The workflow will:
- Diff every module's OpenAPI spec (`develop` vs `master`) using [oasdiff](https://github.com/tufin/oasdiff)
- Compute the correct semver bump for each module:

  | Condition | Bump |
  |---|---|
  | Listed in `force_major_modules` | `major` |
  | Spec-Breaking or Additive spec change | `minor` |
  | No or Cosmetic spec changes | `patch` |
  | Brand new module | starts at `1.0.0` |

Note: spec-breaking changes (labelled `[SPEC-BREAKING]` in the changelog) default to `minor`, not `major`.
These typically reflect amendments to fields that were already required in backend services.

- Automatically generates two draft PRs:
  - **`release → master`** - clean release versions, no `-SNAPSHOT`
  - **`release → develop`** - same, plus a `patch+1 -SNAPSHOT` commit on top to prepare `develop` for the next cycle

If the PR changelog looks wrong, delete the release branch and retrigger this workflow.

---

## Phase 3 - Update, Review and Merge

Once the workflow completes:

1. Review the changelog in the release PR body
2. Update any hand-written tests on the release branch to reflect API changes
3. If any modules were passed via `force_major_modules`, diff their generated code carefully before approving
4. Merge release to `master`
5. Create Tags from `master` based on the PR sugggested version bumps
6. Back-Merge release to develop
