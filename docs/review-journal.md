# Review Journal

The cases below are the review handles I would use before changing the implementation.

The local checks classify each case as `ship`, `watch`, or `hold`. That gives the project a small review vocabulary that matches its mobile workflows focus without claiming live deployment or external usage.

## Cases

- `baseline`: `form pressure`, score 139, lane `watch`
- `stress`: `sync drift`, score 172, lane `ship`
- `edge`: `local state`, score 254, lane `ship`
- `recovery`: `conflict cost`, score 233, lane `ship`
- `stale`: `form pressure`, score 173, lane `ship`

## Note

This file is intentionally plain so the fixture remains the source of truth.
