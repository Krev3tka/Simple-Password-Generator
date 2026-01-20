 # Simple Password Generator by Krev3tka


## Desription

Simple CLI-password generator written in GoLang.


### Functions

- Password Generation of any length

- Possibility to add special characters or digits

- Passwords keep locally to .csv-files which makes it secure to keep, it looks like this:


```csv

Service,Password

Google,j4GhI15XnWnO6t  #without special characters (comments won't appear in the file)

YouTube,D>SbClrn$NYB  #without digits

Firefox,X+^Wh>?Z1N^/E@G  #with digits and special characters

GitHub,LETlXKdzUJzvQXcg  #without digits and special characters

```


## Installation

1. Install needed packages

```bash

sudo apt update && sudo apt install git go # for Debian and Ubuntu

```

```bash

sudo dnf install git go # for Fedora

```

```bash

sudo pacman -S git go # for Arch

```

2. Clone the repo

```bash

git clone https://github.com/Krev3tka/Simple-Password-Generator.git ~/spg

```

3. Go to directory

```bash

cd ~/spg

```

4. Compile project and move it to directory

```bash

go build -o spg

mv spg ~/.local/bin/

```

5. Launch it

```bash

spg

``` 
