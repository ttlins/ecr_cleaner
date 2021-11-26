# ecr_cleaner
## Description
Simple helper script to list and delete ecr images.

## Instructions
- copy `.env.example` to `.env`
- replace / add required values

### Configuration
- We initialize the aws client with default setup
- For that, it's expected that you follow the standard aws configuration setup. E.g
    - using the standard environment variables;
    - having appropriate profiles in a configuration file in, for example, `~/.aws/credentials`
        - in that case, replacing the variables in the provided `.env` file should suffice
    - etc..
    - for more info about that, check out [their docs](https://docs.aws.amazon.com/en_us/sdk-for-go/v1/developer-guide/configuring-sdk.html)

## Usage
- build / install
- help
```bash
./ecr_cleaner --help
```
Obs.: since this project is built with [cobra](https://github.com/spf13/cobra), all commands have the `--help` flag

### required flags
- `ecr_cleaner` only uses two flags, both required. They are:
    - `-r` or `--repository` -> for passing the ecr repository name
    - `-p` or `--pattern` -> for passing the pattern with which you wish to search the repository with

- list images
```bash
./ecr_cleaner list -r repository_name -p tag_pattern
```

- delete images
```bash
./ecr_cleaner delete -r repository_name -p tag_pattern
```

### Patterns
- user provided patterns are compiled as-is (so no escaping)
    - this means regex metacharacters will be interpreted as such
    - if you wish to match literal characters instead, you have two options
        - escape them yourself (might require double backslack `\\`)
        - use the optional `-e` flag
            - this uses [QuoteMeta](https://pkg.go.dev/regexp#QuoteMeta) to escape all metacharacters in the provided string

- As example, some possible patterns could be:
    - jira ticket references, if you tag your test images with that: `REF-1234`
    - the normal SemVer version itself: `1.9.0`
        - since we ignore normal SemVer version, this will only return the pre-release versions associated with it. e.g.:
            - `1.9.0-testing-new-endpoint`
            - `1.9.0-alpha.1`
            - `1.9.0-rc1`
            - `1.9.0.1-REF-1234`
            - etc..
            - BUT NOT: `1.9.0` itself

### NOTE
Tags matching normal SemVer versions (`X.Y.Z`) are never returned by any commands on purpose, to make sure we don't screw up with important images.
