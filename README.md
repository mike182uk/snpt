# <img src="docs/icon.png" height="28"> snpt

[![Version](https://img.shields.io/github/release/mike182uk/snpt.svg?style=flat-square)](https://github.com/mike182uk/snpt)
[![Build Status](https://img.shields.io/github/actions/workflow/status/mike182uk/snpt/ci.yml?branch=main&style=flat-square)](https://github.com/mike182uk/snpt/actions/workflows/ci.yml?query=workflow%3ACI)
[![Coveralls](https://img.shields.io/coveralls/mike182uk/snpt/main.svg?style=flat-square)](https://coveralls.io/r/mike182uk/snpt)
[![Go Report Card](https://goreportcard.com/badge/github.com/mike182uk/snpt)](https://goreportcard.com/report/github.com/mike182uk/snpt)
[![Downloads](https://img.shields.io/github/downloads/mike182uk/snpt/total.svg?style=flat-square)](https://github.com/mike182uk/snpt)
[![License](https://img.shields.io/github/license/mike182uk/snpt.svg?style=flat-square)](https://github.com/mike182uk/snpt)

A [gist](https://gist.github.com/) powered CLI snippet retriever.

Save a snippet as a gist in GitHub, retrieve the snippet on the command line.

![](docs/example.gif)

## Index

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
  - [Syncing your snippets](#syncing-your-snippets)
  - [Listing available snippets](#listing-available-snippets)
  - [Copying a snippet to the clipboard](#copying-a-snippet-to-the-clipboard)
  - [Creating a file from a snippet](#creating-a-file-from-a-snippet)
  - [Printing a snippet to the screen](#printing-a-snippet-to-the-screen)
  - [Setting a new GitHub access token](#setting-a-new-github-access-token)
  - [Viewing help for a command](#viewing-help-for-a-command)
  - [Improve your workflow with fuzzy search](#improve-your-workflow-with-fuzzy-search)
  - [Alfred Workflow](#alfred-workflow)
- [Uninstalling snpt](#uninstalling-snpt)

## <a id="prerequisites"></a>Prerequisites

- GitHub account (duh!)
- GitHub [access token](https://github.com/blog/1509-personal-api-tokens) with the `gist` scope enabled

## <a id="installation"></a>Installation

Download the binary compatible with your system from  [here](https://github.com/mike182uk/snpt/releases).

If you are using macOS, you can also use `Homebrew` to install snpt:

```bash
brew tap mike182uk/tap
brew install mike182uk/tap/snpt
```

## <a id="usage"></a>Usage

### <a id="syncing"></a>Syncing your snippets

Before you can use snpt you will need to sync your gists:

```bash
snpt sync
```

If this is the first time you have synced your gists you will be prompted to input a GitHub [access token](https://github.com/blog/1509-personal-api-tokens) (you will need create this in your GitHub account). This token should be be created with the `gist` scope enabled.

The sync command will download all of your public and private gists and store them locally for fast retrieval by snpt.

You can prevent specific gists from being synced by placing `[snpt:ignore]` anywhere in the description of the gist.

### <a id="list"></a>Listing available snippets

```
snpt ls
```

This can be useful for searching for a specific snippet: 

```bash
snpt ls | grep <query>
```

### <a id="cp"></a>Copying a snippet to the clipboard

```
snpt cp [snippetID|snippetName]
```

If a `snippetID` or `snippetName` is not supplied a prompt will be displayed allowing you to choose a snippet to copy to the clipboard.

If using `snippetName` to search for a snippet, and there are multiple snippets with the same name, the first snippet matching the name will be used. If you have multiple snippets with the same name it is best to search using `snippetId`.

### <a id="write"></a>Creating a file from a snippet

```
snpt write [snippetID|snippetName]
```

If a `snippetID` or `snippetName` is not supplied a prompt will be displayed allowing you to choose a snippet to create a file from. The created file will be named after the name of the gist file.

If using `snippetName` to search for a snippet, and there are multiple snippets with the same name, the first snippet matching the name will be used. If you have multiple snippets with the same name it is best to search using `snippetId`.

### <a id="print"></a>Printing a snippet to the screen

```
snpt print [snippetID|snippetName]
```

If a `snippetID` or `snippetName` is not supplied a prompt will be displayed allowing you to choose a snippet to print to the screen.

If using `snippetName` to search for a snippet, and there are multiple snippets with the same name, the first snippet matching the name will be used. If you have multiple snippets with the same name it is best to search using `snippetId`.

### <a id="token"></a>Setting a new GitHub access token

```
snpt token
```

This command will prompt you to input a new GitHub [access token](https://github.com/blog/1509-personal-api-tokens).

### <a id="help"></a>Viewing help for a command

You can view help for a command by passing the `-h` flag when running a command:

```bash
snpt sync -h
```

### <a id="fuzzy-search"></a>Improve your workflow with fuzzy search

snpt ❤️ [fzf](https://github.com/junegunn/fzf)


![](docs/fzf-example.gif)

```bash
snpt ls | fzf | snpt cp
```

Speed this up by creating aliases for common usages:

```
alias cs="snpt ls | fzf | snpt cp"    # cs for copy snippet
alias ws="snpt ls | fzf | snpt write" # ws for write snippet
```

`snpt cp` and `snpt write` both accept `stdin` as an input. If `stdin` is detected snpt will try and extract a snippet ID from it. This is how the above `fzf` usage works.

### <a id="alfred-workflow"></a>Alfred Workflow

[alfred-snpt](https://github.com/mike182uk/alfred-snpt) provides quick access to your snippets from [Alfred](https://www.alfredapp.com/).

## <a id="uninstall"></a>Uninstalling snpt

To uninstall snpt from your system you will need to manually delete the snpt binary.

snpt's configuration and gist cache is located at `~/.snpt`. You can safely remove this directory and its contents once you have removed the snpt binary.
