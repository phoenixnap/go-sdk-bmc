# Contributing: Release Process

This document explains how to prepare and trigger a release.

---

## Workflows

| Workflow | Purpose |
|---|---|
| `generate-all` | Regenerates all SDK modules from their OpenAPI specs |
| `validate-all` | Validates the generated output compiles and passes contract tests |
| `prepare-release-prs` | Computes version bumps, opens draft PRs |

---

## Phase 1 - Generate and Validate

If you intend to bump the OpenAPI generator version for this release, do it before running the workflow. Update the version in `openapitools.json`, commit and push to `develop` first.

1. Go to **Actions → generate-all → Run workflow**
2. Select branch: `develop`
3. Click **Run workflow**

This regenerates every SDK module from its spec and automatically triggers `validate-all` on completion.

**If `validate-all` passes:** all modules compiled and contract tests passed. Proceed to Phase 2 with no extra inputs - no forced major bumps are needed.

**If `validate-all` fails:** investigate the failing module before proceeding.
- If the failure is a test that simply needs updating for a new required field, fix the test and re-run. No forced major bump needed.
- If the failure is a genuine SDK break (renamed method, changed signature, different model shape - typically caused by a generator version upgrade), note the affected modules. You will pass them as `force_major_modules` in Phase 2.

Do not proceed to Phase 2 until `validate-all` is green.

---

## Phase 2 - Prepare Release PRs

1. Go to **Actions → prepare-release-prs → Run workflow**
2. Select branch: `develop`
3. (Optional) Fill in the input field:

### Input: `force_major_modules`

Leave blank in normal operation - this is an escape hatch for genuine SDK breaks identified during Phase 1.

If `validate-all` failed due to a real SDK break, list the affected modules as a comma-separated string:

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
  | No spec changes | `patch` (strips `-SNAPSHOT`, no increment) |
  | Brand new module | starts at `1.0.0` |

  Note: spec-breaking changes (labelled `[SPEC-BREAKING]` in the changelog) default to `minor`, not `major`. These typically reflect corrections for fields that were already required in backend services rather than true consumer-visible SDK breaks. Only use `force_major_modules` when `validate-all` confirms the generated SDK surface actually changed.

- Automatically generates two draft PRs:
  - **`release → master`** - clean release versions, no `-SNAPSHOT`
  - **`release → develop`** - same, plus a `patch+1 -SNAPSHOT` commit on top to prepare `develop` for the next cycle

If the PR changelog looks wrong, delete the release branch and retrigger this workflow.

---

## Phase 3 - Review and Merge

Once the workflow completes:

1. Review the changelog in the release PR body
2. Update any hand-written tests on the release branch to reflect API changes
3. If any modules were passed via `force_major_modules`, diff their generated code carefully before approving
4. Mark **both** draft PRs as ready for review
5. Merge in order:
   - **`release → master` first**
   - **`release → develop` second**

The develop PR contains a snapshot version commit on top of the release commit. Merging master first ensures that snapshot commit never lands on master.
