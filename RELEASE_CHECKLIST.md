# EduGo v1.0.0 release checklist

EduGo enters feature freeze when this checklist is used. Release work should fix correctness, pedagogy, reproducibility, publishing, or documentation defects rather than add new course scope.

## Content

- [ ] Chapters 0–105 exist in both Norwegian and English.
- [ ] Norwegian and English chapter pairs teach the same learning objective.
- [ ] All code shown as runnable has been checked against Go 1.25.
- [ ] Capstone instructions and acceptance criteria are complete.
- [ ] No secrets, credentials, malware samples, or proprietary course dependencies are present.

## Verification

- [ ] CI format gate passes.
- [ ] `go vet ./...` passes for every module.
- [ ] `go test ./...` passes for every module.
- [ ] `go test -race ./...` passes for every module.
- [ ] Student OCI builds.
- [ ] Student OCI smoke test reports the expected Go toolchain and can compile/run a minimal Go program.

## Publishing

- [ ] Title, author, publisher, copyright, language, and CC-BY-4.0 metadata are present.
- [ ] HTML/web output builds.
- [ ] EPUB output builds.
- [ ] Kindle-compatible output builds.
- [ ] PDF output builds.
- [ ] Internal links and code blocks survive all publication formats.

## Release

- [ ] README and course index match the frozen scope.
- [ ] Fresh-student walkthrough completed from the documented environment.
- [ ] Known limitations are documented.
- [ ] Release notes summarize the v1.0 scope.
- [ ] Tag `v1.0.0` only after all required gates above are complete.
