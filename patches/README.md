# Shadowfork patch

Generate new patch with script.

```sh
git diff main..<target_commit> -- . ':!.github' ':!.golangci.yml' ':!*.log' ':!*.out' ':!*.DS_Store' ':!*.test' > ./patches/shadow_fork.patch
```

Current patch for commit `d805aecf1c690864625392d99cc451be9bf79d99`