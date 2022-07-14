# sdfxshape

A simple wrapper around the https://github.com/deadsy/sdfx Go-based solid modeling package.  Re-enables fluent-style code when it's okay to panic.  Needed because we removed panics in core code in deadsy/sdfx#24, deadsy/sdfx#26, deadsy/sdfx#27.  Background and rationale discussed in more detail in deadsy/sdfx#50 and deadsy/sdfx#51.
