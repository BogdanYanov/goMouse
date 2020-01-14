# goMouse

goMouse it is CLI tool that simulates computer mouse.

## About

When you start programm you will see next text:

```
goMouse is CLI tool which simulates mouse control by commands.

Usage:
  goMouse [command]

Available Commands:
  btn-down    The command performs a button push
  btn-up      The command performs a button release
  help        Help about any command
  info        This command shows the settings values of mouse
  move        Moving mouse cursor to X Y coordinates relative to the screen
  reset       This command reset mouse setting to default
  scroll      This command scroll the mouse wheel up or down
  sens        Set mouse sensitivity

Flags:
      --config string   config file (default is $HOME/.goMouse.yaml)
  -h, --help            help for goMouse
  -t, --toggle          Help message for toggle

Use "goMouse [command] --help" for more information about a command.
subcommand is required
```

All settings and states that you will be enter by commands saves in file *"mouse.json"*.

Test coverage for this programm 100%:

```
PASS
coverage: 100.0% of statements
ok      github.com/BogdanYanov/goMouse/mouse    0.018s
```

## Programm example

```
scriptkiller@scriptkiller-X550MJ:~/go/bin$ ./goMouse move 1 1
scriptkiller@scriptkiller-X550MJ:~/go/bin$ ./goMouse sens 3
scriptkiller@scriptkiller-X550MJ:~/go/bin$ ./goMouse scroll -u
scriptkiller@scriptkiller-X550MJ:~/go/bin$ ./goMouse scroll -u
scriptkiller@scriptkiller-X550MJ:~/go/bin$ ./goMouse btn-down -r
scriptkiller@scriptkiller-X550MJ:~/go/bin$ ./goMouse info
Mouse information:
X position - 1
Y position - 1
Sensitivity - 3
Is left button pressed? - false
Is right button pressed? - true
Scroll value - 7
scriptkiller@scriptkiller-X550MJ:~/go/bin$ ./goMouse reset
scriptkiller@scriptkiller-X550MJ:~/go/bin$ ./goMouse info
Mouse information:
X position - 512
Y position - 384
Sensitivity - 1
Is left button pressed? - false
Is right button pressed? - false
Scroll value - 5
```
