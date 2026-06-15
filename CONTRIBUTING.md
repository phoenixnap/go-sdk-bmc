# Contributing

## Taking a Release

Follow this process when preparing and publishing a new release.

### 1. Validate the current state

Optionally run `generate-all` and `validate-all` manually, then confirm the latest `validate-all` run completed successfully before continuing.

If `validate-all` failed:

* Review the failed jobs.
* Determine whether failures stem from compilation issues or SDK breaking changes.
* Record any modules affected by SDK breaking changes.

### 2. Prepare the release PRs

Run the `prepare-release-prs` workflow. Version handling is fully automated — no manual version changes are expected.

* Provide SDK breaking-change modules as a comma-separated input; these receive a **major** bump.
* Unlisted modules are bumped **minor** or **patch** based on detected `spec` changes.
* New modules start at `1.0.0`.

### 3. Review and update the draft release PR

Open the generated draft release PR. Its human-readable notes describe the detected spec changes — use them to identify which endpoints need review.

In your IDE:

* Update the tests package with matching payloads and endpoints from the new spec changelog where applicable.

### 4. Merge the release PR

Merge the `release/vX` branch PR into `master`. A workflow then opens an automatic back-merge PR from `release/vX` to `develop` — review and merge that as well. Version handling remains automated throughout.

### 5. Publish GitHub tags

After merging into `master`, manually dispatch the `publish-git-tags` workflow in this order:

1. Run it as a **dry run** and confirm the output is correct.
2. Run it normally.

This publishes Git tags for each sub-module and the root repository module.

### 6. Take a release

Create a release based on the previously generated repo module release tag.

### 7. Confirm release

Confirm the Go SDK release is available on [pkg.go.dev](https://pkg.go.dev/). The release is complete once the expected root module and sub-module versions appear.
