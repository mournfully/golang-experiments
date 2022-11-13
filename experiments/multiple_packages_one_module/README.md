You can find the relevant example at `../../exercise-2-urlshort/`

---

The setting has changed, now you need to use this setting in VSCode's settings.json. Then save and restart VSCode to apply changes.
"gopls": {
    "build.experimentalWorkspaceModule": true,
}

From Go 1.18 onwards there is native support for multi-module workspaces. This is done by having a go.work file present in your parent directory.

For a directory structure such as:
```shell
$ tree /my/parent/dir
/my/parent/dir
├── project-one
│   ├── go.mod
│   ├── project-one.go
│   └── project-one_test.go
└── project-two
    ├── go.mod
    ├── project-two.go
    └── project-two_test.go
```

Create and populate the file by executing go work:
```shell
cd /my/parent/dir
go work init
go work use project-one
go work use project-two
```

This will add a go.work file in your parent directory that contains a list of directories you marked for usage:
```
go 1.18

use (
    ./project-one
    ./project-two
)
```

visual studio code - VScode shows an error when having multiple Go Projects in a directory - Stack Overflow
https://stackoverflow.com/questions/65748509/vscode-shows-an-error-when-having-multiple-go-projects-in-a-directory

---

You cannot have two packages per directory, hence the error. So the solution as @Larry Battle said to move your myproject.go to a new directory.

go - Does it make sense to have two packages in the same directory? - Stack Overflow
https://stackoverflow.com/questions/20427890/does-it-make-sense-to-have-two-packages-in-the-same-directory

---

> Can I have multiple packages inside a single go module?
Yes, of course.

> How?
You have to do nothing, it just works.

(Your problem is: You try to import your api package by a wrong name. Import paths inside a module are of the form <modulename>/<relative-filepath-from-module-root>.)



You can have as many package as you want in a single module, have a look here: https://github.com/alessiosavi/GoGPUtils/

The only constraint is that you cane have a single package for every folder.



Can I have multiple packages inside a single go module? How? - Stack Overflow
https://stackoverflow.com/questions/60320844/can-i-have-multiple-packages-inside-a-single-go-module-how

---

The next step is to create a go.mod file within the mymodule directory to define the Go module itself. To do this, you’ll use the go tool’s mod init command and provide it with the module’s name, which in this case is mymodule. Now create the module by running go mod init from the mymodule directory and provide it with the module’s name, mymodule:
```shell
go mod init mymodule
```

With the module created, your directory structure will now look like this:
```
└── projects
    └── mymodule
        └── go.mod
```

How to Use Go Modules | DigitalOcean
https://www.digitalocean.com/community/tutorials/how-to-use-go-modules