# ANGS Pilot

Public, non-production integration fixture for the ANGS trusted-delivery pilot.

This repository contains no credentials, production data, customer code, or private policy. Its purpose is to exercise GitHub pull requests, merge queues, reusable workflows, exact-SHA checks, and failure-closed evidence validation before real-team rollout.

The trusted runner, workflow repository, and GitHub App remain separate trust boundaries. A successful check in this repository is technical pilot evidence only; it is not a claim that M4 trusted team governance is complete.

## First integration scenario

The first pull request verifies the ordinary PR gate. After repository rules enable the merge queue, the queued candidate will exercise the separately authenticated `merge_group` evidence path.

The reusable trusted workflow is consumed from its immutable public commit SHA so the pilot never follows a moving branch or tag.

Trusted evidence runs once for the pull request's GitHub-generated merge candidate and again for the merge queue's final combined candidate. The GitHub App-authored `ANGS Trusted Evidence` check is intended to be required for both stages.

Each stage is bound to the exact SHA delivered by its authenticated GitHub event.

For pull requests, the Runner also verifies the current GitHub test-merge SHA against the authenticated head and base pair before accepting evidence.

That resolution uses a short-lived installation token scoped to this repository with pull-request read and checks write permissions only.
