# Contributing to jsReveal

First off, thank you for considering contributing to jsReveal! It's people like you that make jsReveal such a great tool.

## Where do I go from here?

If you've noticed a bug or have a feature request, [make an issue post](https://github.com/SupremeERG/jsReveal/issues/new)! It's generally best if you get confirmation of your bug or approval for your feature request this way before starting to code.

If you have a general question, feel free to reach out to us on [Twitter](https://twitter.com/egx08).

## Fork & create a branch

If this is something you think you can fix, then [fork jsReveal](https://github.com/SupremeERG/jsReveal/fork) and create a branch with a descriptive name.

A good branch name would be (where issue #325 is the ticket you're working on):

```sh
git checkout -b 325-add-japanese-translations
```

## Get the style right

Your patch should follow the same conventions & pass the same code quality checks as the rest of the project.

## Make a Pull Request

At this point, you should switch back to your master branch and make sure it's up to date with jsReveal's master branch:

```sh
git remote add upstream git@github.com:SupremeERG/jsReveal.git
git checkout master
git pull upstream master
```

Then update your feature branch from your local copy of master, and push it!

```sh
git checkout 325-add-japanese-translations
git rebase master
git push --force-with-lease origin 325-add-japanese-translations
```

Finally, go to GitHub and [make a Pull Request](https://github.com/SupremeERG/jsReveal/compare)!

## Keeping your Pull Request updated

If a maintainer asks you to "rebase" your PR, they're saying that a lot of code has changed, and that you need to update your branch so it's easier to merge.

To learn more about rebasing and merging, check out this guide on [merging vs. rebasing](https://www.atlassian.com/git/tutorials/merging-vs-rebasing).

## Merging a PR (for maintainers)

A PR can only be merged by a maintainer if it has at least one approval from a maintainer and all of the checks have passed. A maintainer should ask for reviews from other maintainers and community members.

## Releasing a new version (for maintainers)

To release a new version of jsReveal, you'll need to create a new release on GitHub. The release should include a summary of the changes and a link to the full changelog. The release should also be signed with the maintainer's GPG key.

## Shipping a new version (for maintainers)

To ship a new version of jsReveal, you'll need to create a new tag on GitHub. The tag should be named after the version number (e.g. `v1.0.0`). The tag should also be signed with the maintainer's GPG key.
