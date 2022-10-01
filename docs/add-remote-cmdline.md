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

   ~/dev/proj/go-intro-01$ git branch --set-upstream-to=go-intro-01/main main

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
