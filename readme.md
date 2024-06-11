# Headstart
* Still on development 
Headstart is a productivity cli tool, file/dir creation, boiler code gen etc..

## Installation
```bash
git clone git@github.com:BarunW/headstart.git 

# Add to .zshrc or .bashrc file or other shell rc file
# /path/to = where the headstart repo is
export PATH="/path/to/headstart/bin:$PATH"

#souce 
source ~/.zshrc or ~/.bashrc (linux)
```

## Usage 
* headstart link <filePath(file or dir)> <name the link> 
* headstart gen <name of the link> <file/dir name> 
* headstart <Name of text Editor> <name of the link> 
* headstart [list all link name]

``` bash
# open your terminal and type headstart and usage will be shown
# project setup
headstart link boilerPlate.go  goBoil
headstart nvim goBoil 
headstart gen goBoil "../newGo" 
headstart  link <binary or bashscript or any executable> 
headstart exec <that link executable>
```
## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
