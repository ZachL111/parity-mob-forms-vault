# parity-mob-forms-vault

`parity-mob-forms-vault` packages a practical mobile workflows exercise in Go. The emphasis is on deterministic behavior, a small public API, and examples that explain the tradeoffs.

## How I Read Parity Mob Forms Vault

The useful thing to inspect here is how the same score rule is represented in code, metadata, and examples. If those three pieces disagree, the audit script should make the drift visible.

## Problem Shape

This is not a wrapper around a service. It is a self-contained project that shows how the model behaves when demand, capacity, latency, risk, and weight move in different directions.

## Main Behaviors

- Models local state with deterministic scoring and explicit review decisions.
- Uses fixture data to keep sync pressure changes visible in code review.
- Includes extended examples for form constraints, including `recovery` and `degraded`.
- Documents offline paths tradeoffs in `docs/operations.md`.
- Runs locally with a single verification command and no external credentials.

## Internal Model

The project is organized around a compact model rather than a large framework. Inputs are scored, classified, and checked against golden fixtures. The constants live in code and are mirrored in metadata so documentation drift is easy to catch. The Go layout uses small packages and table-oriented tests so the behavior stays easy to follow.

## Repository Map

- `policy`: Go package with the core model
- `cmd`: small command entry point
- `fixtures`: compact golden scenarios
- `examples`: expanded scenario set
- `metadata`: project constants and verification metadata
- `docs`: operations and extension notes
- `scripts`: local verification and audit commands
- `go.mod`: Go module metadata

## Run It Locally

Clone the repository, enter the directory, and run the verifier. No database server, cloud account, or token is required.

## How To Run It

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/verify.ps1
```

This runs the language-level build or test path against the compact fixture set.

## Validation

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/audit.ps1
```

The audit command checks repository structure and README constraints before it delegates to the verifier.

## Scenario Walkthrough

`baseline` is the first example I would inspect because it lands on the `accept` path with a score of 167. The broader file also keeps `degraded` at 6 and `recovery` at 223, which gives the model a useful low-to-high spread.

## Known Edges

The fixture set is deliberately small. That keeps the review surface clear, but it also means the model should not be treated as a complete domain simulator.

## Follow-Up Work

- Add a loader for `examples/extended_cases.csv` and promote selected cases into the language test suite.
- Add a short report command that prints the score breakdown for a single scenario.
- Add malformed input fixtures so the failure path is as visible as the happy path.
- Add one more mobile workflows fixture that focuses on a malformed or borderline input.
