# Contributing

Thanks for helping improve FreshRSS-Recommender. This project is still early, so the most useful contributions are small, well-scoped changes with clear context and verification notes.

## Before You Start

Use the issue templates to make the work easy to pick up and review:

- Use **Bug** when observed behavior is wrong and needs reproduction steps.
- Use **Design** when the direction is not settled yet.
- Use **Task** for concrete implementation work.
- Use **Refactor** for behavior-preserving structure changes.

If an issue already exists, leave a short comment before starting so reviewers can connect the pull request to the right context.

## Pull Requests

Keep pull requests focused. A pull request should have one clear purpose and explicit non-goals when the boundary could be misunderstood.

Before opening a pull request:

1. Link the issue or design discussion that explains the work.
2. Keep unrelated cleanup, formatting, and refactors out of the change.
3. Add or update tests only when the requirement or bug has a stable acceptance scenario.
4. Run the checks that apply to the files you changed, and document what you ran in the pull request.
5. Call out any manual verification, known risk, or intentional non-goal.

If there is no documented command for a check yet, describe the manual validation you performed instead of inventing a project workflow.

## Review Expectations

Review focuses on correctness, scope discipline, maintainability, and whether the change can be verified. Expect reviewers to ask for a narrower diff when a pull request mixes unrelated goals.

Good contributions usually include:

- Clear context for why the change exists.
- A small scope with stated non-goals.
- Observable acceptance criteria.
- Verification notes that another maintainer can repeat.
- Links to related issues, discussions, or prior decisions.

## Project Style

Follow the repository guidance in `AGENTS.md` when working in this codebase. In particular, prefer the smallest complete change, preserve existing terminology, fail fast instead of hiding invalid state, and avoid speculative abstractions or future-facing configuration.
