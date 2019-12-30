## Contributing to this project

Contributing to Notification Service? There are a few things you need to know.

### Semantic Versioning

This project follows [semantic versioning](http://semver.org/). We release patch versions for bugfixes, minor versions for new features, and major versions for any breaking changes.

We tag every pull request with a label marking whether the change should go in the next patch, minor, or a major version.

Every significant change is documented in the changelog file.

Because this is more of an example library, semver will be based on the api landscape of the downstream library versions

### Branch organization

The `master` branch should always have a green continuous integration(CI) build, with all linting and tests passing successfully.  It is recommended that you are constantly pulling any changes that exist on the `master` branch to minimize conflicts.

If you send a pull request, please do it against the `master` branch.  Exceptionally, we may at some point process pull requests against feature branches, coordinate accordingly.

Please reference the issue in your commit and your branch to help association.

Name your branches based on the Atlassian's naming convention ([Atlassian Documentation](https://confluence.atlassian.com/bitbucketserver/using-branches-in-bitbucket-server-776639968.html)):
- `{category}/{story#}-{storyDescription}`
    - category:
        - feature
        - bugfix
        - release
    - storyDescription:
        - Typically this can be pulled straight from the JIRA story description.  An example would be, _"Wire up initial bridge data to frontend application"_
 - Full example would look like: `feature/issue-3-resolving-pg-connection-issue`

### Commits
All commits should be atomic and follow these standards:

-   Separate subject from body with a blank line
-   Limit the subject line to 50 characters
-   Capitalize the subject line
-   Do not end the subject line with a period
-   Use the imperative mood in the subject line
-   Wrap the body at 72 characters
-   Use the body to explain what and why vs. how

### Continuous Integration

Pull requests will automatically trigger a Continuous Integration build.  This process is critical to ensuring code quality.

Pull requests without a green build will not be merged into the `master` branch.

### Rollback

It is possible that a rollback may be required due to outside factors that could not be mitigated through the Continuous Integration pipeline.  All rollbacks will be processed as `git revert` operations.  These occurances will be broadcasted through the changelog as well as email and interpersonal communication.
