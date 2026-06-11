# Contributing

## Taking a Release

Follow this process when preparing and publishing a new release.

### 1. Validate the current state

Optionally run the following workflows manually:

* `generate-all`
* `validate-all`

Before continuing, confirm that the latest `validate-all` workflow completed successfully.

If `validate-all` failed:

* Review the failed jobs.
* Identify whether the failures are caused by compilation issues or SDK breaking changes.
* Record any modules affected by SDK breaking changes.

### 2. Prepare the release PRs

Run the `prepare-release-prs` workflow.

When starting the workflow:

* Provide any SDK breaking-change modules as a comma-separated input.
* Modules listed as SDK breaking changes will receive a major version bump.
* Any module not listed will be automatically bumped as either `minor` or `patch`, depending on the changes detected in the `spec` files.
* Any new module will automatically start at version `1.0.0`.

Version handling is managed by the workflow. No manual version changes are expected.

### 3. Review and update the draft release PR

Open the generated draft release PR.

The PR includes human-readable notes describing the detected spec changes. Use these notes to identify which endpoints require review.

Using your preferred IDE:

* Update the tests package with matching payloads and endpoints from the new spec changelog where applicable.
* Add or update tests manually as needed.

### 4. Merge the release PR

Merge the `release/vX` branch PR into `master`.

After the merge, an automatic workflow will create a back-merge PR from `release/vX` to `develop`.

Review and merge the back-merge PR as well.

Version handling remains automated throughout this step. No manual version changes are expected.

### 6. Publish GitHub tags

After merging the release branch into master, manually dispatch the `publish-git-tags` workflow.

Run it in the following order:

Run `publish-git-tags` as a dry run and confirm the output is correct.
Run `publish-git-tags` normally.

This workflow publishes Git tags for each sub-module and for the root repository module.

### 7. Take a release

Create a release based on the previously generated repo module release tag.

### 8. Confirm release

Confirm that the Go SDK release is available on pkg.go.dev:

https://pkg.go.dev/github.com

The release is complete once the expected root module and sub-module versions are available.
