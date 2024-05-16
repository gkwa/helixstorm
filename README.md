# helixstorm

Purpose:
Learn ugrep.

I couldn't get the syntax right in order to get files that contain both 'list' and 'unzip'.  It turns out I should be using '--bool --files'.

I struggled with this for a while and in the process, I generated test files.  Hence this go app.

Its a one off really.

tl;dr:

```bash

# find files that contain both unzip and list anywhere in the file
ugrep /tmp/ugrep-gh4X/ --recursive --bool --files --regexp unzip --regexp list 

```

## example usage

```bash


```

## install helixstorm


on macos/linux:
```bash

brew install taylormonacelli/homebrew-tools/helixstorm

```


on windows:

```powershell

TBD

```
