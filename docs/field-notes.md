# Field Notes

`parity-mob-forms-vault` is easiest to review by starting with the fixture, not the prose.

The domain cases cover `form pressure`, `sync drift`, `local state`, and `conflict cost`. They sit beside the smaller starter fixture so the project has both a compact scoring check and a domain-flavored review check.

The widest spread is between `local state` and `form pressure`, so those are the first two cases I would preserve during a refactor.

The extra check gives the repository a behavior path that can fail for a domain reason, not only a syntax reason.
