Adding Remote GitHub from command line
--------------------------------

    git status
    On branch main
    Changes not staged for commit:
    (use "git add/rm <file>..." to update what will be committed)
    (use "git restore <file>..." to discard changes in working directory)
        deleted:    dup2

    Untracked files:
    (use "git add <file>..." to include in what will be committed)
        ../../.gitignore
        ../../README.md

    no changes added to commit (use "git add" and/or "git commit -a")
    cd ../..
    git status
    On branch main
    Changes not staged for commit:
    (use "git add/rm <file>..." to update what will be committed)
    (use "git restore <file>..." to discard changes in working directory)
        deleted:    ch1/dup/dup2

    Untracked files:
    (use "git add <file>..." to include in what will be committed)
        .gitignore
        README.md

    no changes added to commit (use "git add" and/or "git commit -a")
    git add .
    git status
    On branch main
    Changes to be committed:
    (use "git restore --staged <file>..." to unstage)
        new file:   .gitignore
        new file:   README.md
        deleted:    ch1/dup/dup2

    git commit -m "Intro go 2"
    [main a19a82b] Intro go 2
    3 files changed, 360 insertions(+)
    create mode 100644 .gitignore
    create mode 100644 README.md
    delete mode 100755 ch1/dup/dup2
    git config --global user.name "peteranton1"
    git push
    fatal: No configured push destination.
    Either specify the URL from the command-line or configure a remote repository using

        git remote add <name> <url>

    and then push using the remote name

        git push <name>

    cd ..
    ~/dev/proj$ mv go-intro go-intro-01
    ~/dev/proj$ cd go-intro-01/
    ~/dev/proj/go-intro-01$ git remote add go-intro-01 git@github.com:peteranton1/go-intro-01.git
    ~/dev/proj/go-intro-01$ git push
    fatal: No configured push destination.
    Either specify the URL from the command-line or configure a remote repository using

        git remote add <name> <url>

    and then push using the remote name

        git push <name>

    ~/dev/proj/go-intro-01$ git push go-intro-01
    To github.com:peteranton1/go-intro-01.git
    ! [rejected]        main -> main (fetch first)
    error: failed to push some refs to 'github.com:peteranton1/go-intro-01.git'
    hint: Updates were rejected because the remote contains work that you do
    hint: not have locally. This is usually caused by another repository pushing
    hint: to the same ref. You may want to first integrate the remote changes
    hint: (e.g., 'git pull ...') before pushing again.
    hint: See the 'Note about fast-forwards' in 'git push --help' for details.
    ~/dev/proj/go-intro-01$ git pull
    There is no tracking information for the current branch.
    Please specify which branch you want to merge with.
    See git-pull(1) for details.

        git pull <remote> <branch>

    If you wish to set tracking information for this branch you can do so with:

        git branch --set-upstream-to=go-intro-01/<branch> main

    ~/dev/proj/go-intro-01$ git pull go-intro-01 main
    remote: Enumerating objects: 3, done.
    remote: Counting objects: 100% (3/3), done.
    remote: Compressing objects: 100% (2/2), done.
    remote: Total 3 (delta 0), reused 0 (delta 0), pack-reused 0
    Unpacking objects: 100% (3/3), 1.18 KiB | 603.00 KiB/s, done.
    From github.com:peteranton1/go-intro-01
    * branch            main       -> FETCH_HEAD
    * [new branch]      main       -> go-intro-01/main
    hint: You have divergent branches and need to specify how to reconcile them.
    hint: You can do so by running one of the following commands sometime before
    hint: your next pull:
    hint: 
    hint:   git config pull.rebase false  # merge (the default strategy)
    hint:   git config pull.rebase true   # rebase
    hint:   git config pull.ff only       # fast-forward only
    hint: 
    hint: You can replace "git config" with "git config --global" to set a default
    hint: preference for all repositories. You can also pass --rebase, --no-rebase,
    hint: or --ff-only on the command line to override the configured default per
    hint: invocation.
    fatal: Need to specify how to reconcile divergent branches.
    ~/dev/proj/go-intro-01$ git config pull.rebase false
    ~/dev/proj/go-intro-01$ git config --global pull.rebase false
    ~/dev/proj/go-intro-01$ git pull go-intro-01 main
    From github.com:peteranton1/go-intro-01
    * branch            main       -> FETCH_HEAD
    fatal: refusing to merge unrelated histories
    ~/dev/proj/go-intro-01$ git config --global pull.ff only
    ~/dev/proj/go-intro-01$ git pull go-intro-01 main
    From github.com:peteranton1/go-intro-01
    * branch            main       -> FETCH_HEAD
    fatal: Not possible to fast-forward, aborting.
    ~/dev/proj/go-intro-01$ git config --global pull.rebase true
    ~/dev/proj/go-intro-01$ git pull go-intro-01 main
    From github.com:peteranton1/go-intro-01
    * branch            main       -> FETCH_HEAD
    fatal: Not possible to fast-forward, aborting.
    ~/dev/proj/go-intro-01$ git pull
    There is no tracking information for the current branch.
    Please specify which branch you want to merge with.
    See git-pull(1) for details.

        git pull <remote> <branch>

    If you wish to set tracking information for this branch you can do so with:

        git branch --set-upstream-to=go-intro-01/<branch> main

    ~/dev/proj/go-intro-01$ git branch --set-upstream-to=go-intro-01/main main
    Branch 'main' set up to track remote branch 'main' from 'go-intro-01'.
    ~/dev/proj/go-intro-01$ git push
    To github.com:peteranton1/go-intro-01.git
    ! [rejected]        main -> main (non-fast-forward)
    error: failed to push some refs to 'github.com:peteranton1/go-intro-01.git'
    hint: Updates were rejected because the tip of your current branch is behind
    hint: its remote counterpart. Integrate the remote changes (e.g.
    hint: 'git pull ...') before pushing again.
    hint: See the 'Note about fast-forwards' in 'git push --help' for details.
    ~/dev/proj/go-intro-01$ git pull -f
    fatal: Not possible to fast-forward, aborting.
    ~/dev/proj/go-intro-01$ git pull 
    fatal: Not possible to fast-forward, aborting.
    ~/dev/proj/go-intro-01$ git pull --rebase
    Successfully rebased and updated refs/heads/main.
    ~/dev/proj/go-intro-01$ git status
    On branch main
    Your branch is ahead of 'go-intro-01/main' by 2 commits.
    (use "git push" to publish your local commits)

    nothing to commit, working tree clean
    ~/dev/proj/go-intro-01$ git push
    Enumerating objects: 22, done.
    Counting objects: 100% (22/22), done.
    Delta compression using up to 12 threads
    Compressing objects: 100% (19/19), done.
    Writing objects: 100% (21/21), 1.02 MiB | 8.04 MiB/s, done.
    Total 21 (delta 1), reused 0 (delta 0), pack-reused 0
    remote: Resolving deltas: 100% (1/1), done.
    To github.com:peteranton1/go-intro-01.git
    84fb555..76ed93b  main -> main

    ~/dev/proj/go-intro-01$ 
