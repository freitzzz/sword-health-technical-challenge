# sword-health-technical-challenge

The present repository contains all source used to perform SWORD Health technical challenge. I've chosen `Golang` as the primary programming language for three reasons: its the job position intended language, its rapid for prototyping and prevents from nasty bugs such as null pointers and exceptions. To produce a more valuable, production result, it was decided to follow some strict conventions as documented in the section bellow. These conventions are both suggestions from the official Golang team and past members [2], [3], as well as the community in general [1].

## Conventions

For naming [2]:

- Keep it short, keep it simple: Reduce common variables to one,two or three letters (e.g., i - index, k - key, ip - ip address, ctx - context);
- Leave **err** for errors;
- Use **camelCase** for local variables, parameters (e.g., sortedList);
- Use **CAPITALCASE** for acronyms (e.g., HTTP, ID);
- Use **PascalCase** for functions, types, interfaces (e.g., PrimeSearch);
- Use **lowercase** only for package names and try to short for one noun (e.g., user, auth);
- Use Error as a **suffix** when defining error types (e.g., UsernameError);
- Use Err as a **prefix** when declaring error variables (e.g., ErrUsername);
- Don't use **snake_case**.

For packages [2], [3]:

- Organize by responsibility, not collection: Avoid models, utils, etc;
- Separate files by responsibility (e.g., http.go, headers.go, cookies.go);
- Top-level package documentation should be written in the `doc.go` file.
- `/cmd/` directory describes the main executable code for the Go progam;
- `/internal/` directory describes internal, non-shared code (i.e., private application code);
- `/pkg/` directory describes shared code externally that other applications can use;
- `/configs/` directory describes configuration files templates;
- `/deployments/` directory describes deployment configuration files;
- `/scripts/` directory describes internal scripts used to automate the application (e.g., build, test, analyse);
- `/third_party/` directory describes external helper tools binaries, files and forked code;
- Directories following /cmd/ should match the name of the executable;
- Directories inside /internal/ can be named as /pkg and /app to separate application and libraries code.

### Sources:

[1] [Community Standard Project Layout for Golang - 2021](https://github.com/golang-standards/project-layout)

[2] [Naming Conventions Proposed By Andrew Gerrand (Google) - 2014](https://talks.golang.org/2014/names.slide)

[3] [Package Conventions Proposed By Rakyll (Google) - 2017](https://rakyll.org/style-packages/)