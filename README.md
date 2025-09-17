# dagger-withfiles-bug

This attempts to demonstrate a bug in `WithFiles` introduced with Dagger 0.18.18.
Under 0.18.18, using `WithFiles` with files from a subdirectory of a `Directory` will reproduce the path to the subdirectory in the destination path of files created by `WithFiles`.

A short test script is included in the justfile that can be run with `just test`.
This runs the `CheckFile` function with the expected and unexpected paths to illustrate the change in behavior.

When running the test, you can override the Dagger binary by setting `DAGGER`, this can be used to test against both 0.18.18 and 0.18.17.
Under 0.18.17, the test will pass.
