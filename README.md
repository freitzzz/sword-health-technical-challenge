# sword-health-technical-challenge

The present repository contains all source used to perform SWORD Health technical challenge. I've chosen `Golang` as the primary programming language for three reasons: its the job position intended language, its rapid for prototyping and prevents from nasty bugs such as null pointers and exceptions. To produce a more valuable, production result, it was decided to follow some strict conventions as documented in the section bellow. These conventions are both suggestions from the official Golang team and past members [2], [3], as well as the community in general [1].

## Challenge

The technical challenge involves producing a system capable of allowing users to manage tasks. These users are distinguished in two roles: `Manager` and `Technician`. The manager supervises in some form technicians, as his role allows him to view their tasks and delete them. Managers also receive notifications when a technician performs a task.

A task is composed of two key elements:

- Summary, supporting a max. of 2500 characters. Summaries can also contain personal information;
- Date, of when it was performed.

These elements raise some concerns: personal information and control management. Data privacy is an hot topic these days, as data protection laws are very strict. This requires the application of a strict encryption algorithm in the task summary. Also, tasks can be updated. Should a task have a creation and update date?

It is also mentioned that tasks are performed during working days. Should it be prevented that employees do not update their tasks during non-working days?

------

For features, it has been asked to:

- Create an API endpoint to save a new task (Create/Update/Delete);
- Create an API endpoint to list tasks;
- Notify manager of each task performed by the tech, without blocking any http request.

For technical requirements:

- Use any language to develop the system HTTP API;
- Use MySQL database to persist data from the application;
- Create a local development environment using docker containing the services and a MySQL database;
- Features should have unit tests to ensure they are working properly.

As a bonus, notification logic should be decoupled from the main application flow, using a message broker. Also, system deployment in Kubernetes is a plus.

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