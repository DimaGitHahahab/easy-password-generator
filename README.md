# easy-password-generator

Simple CLI tool to generate passwords with specific properties.

### Usage:

Run `password-generator [flags]` in terminal.

### Flags:

`-h` or `--help`: Display help information.

`-l <int>` or `--length <int>`: Set the length of the generated password (default is 10 characters).

`-s` or `--special`: Exclude special characters.

`-n` or `--numbers`: Exclude numbers.

`-u` or `--upper`: Exclude uppercase.

`-L` or `--lower`: Exclude lowercase letters.

### Example:

`password-generator -l 12 -L -n`
