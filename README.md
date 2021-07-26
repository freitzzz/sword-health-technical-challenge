# sword-health-technical-challenge

The present repository contains all source used to perform SWORD Health technical challenge. I've chosen `Golang` as the primary programming language for three reasons: its the job position intended language, its rapid for prototyping and prevents from nasty bugs such as null pointers and exceptions. To produce a more valuable, production result, it was decided to follow some strict conventions as documented in the section below. These conventions are both suggestions from the official Golang team and past members [2], [3], as well as the community in general [1].

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

## Design

Once analysed the challenge context, a preliminary design can be thought of. The product dwells on several domains:

- User, authentication and authorization management;
- Tasks;
- Permissions;
- Notifications;
- Security.

To make things easier, I suggest that users can be mocked in some form, as a complete signup/login process is not required by the challenge. To control API access, authentication keys will still be required, but these can be generated by passing the user id.

Permissions are decided based on the role of the user. If the user is a technician, he can: list and view his tasks, create a task, update and delete it. If the user is a manager, he can: list and view technician tasks, as well as deleting them. Also, as a bonus I'd suggest that updates on tasks are prohibited on non-working days (i.e., saturday, sunday), to comply with the challenge description.

Notifications are published under the hood when a task is performed, to be later seen by the managers. This raises a question: are notifications established in a many-to-many relationship? Do all managers receive at the same time a notification from a single update on the task by a technician? Maybe it would be better to allow to opt-in for specific updates on the task (all updates, only create, update or delete) or for a specific technician. Also, what happens when managers are not online and a task is updated? If they want to see the notifications, these need to be stored on a separate queue/mail box, that can be later accessed remotely by the manager.

Security is present not only on the autorization layer, but also on tasks and notifications. Task summary may contain personal identification, so it is required to be applied encryption.

To better visualize these domain concepts, the following domain diagram is proposed:

![domain_diagram](docs/assets/sword_health_technical_challenge_domain_diagram.png)

<center><i>Figure 1 - Domain Diagram illustrating the domain concepts relationships, with UML (Tool: draw.io)</i></center>

Laid down the domain concepts, more architecture decision can be made. The monolith vs micro-service decision can be settle down based on what is known so far. Following a monolith approach would be a quick way to build and ship the system, but it suffers from several issues such as:

- Domain scaling: if more and more concepts are introduced, the system becomes a giant ball-of-mud;
- Too much responsibilities: Having all the code and runtime for the users, authentication, authorization, tasks, permissions and notifications on a single system is hard to both develop and maintain;
- Cross-cutting concerns: performance, single point of failure, etc.

Having a giant system is not acceptable, so there is the need to cut it down in pieces. To follow a micro-service approach, it is needed to apply a strategy to divide the system in smaller services. There are already plenty famous patterns that are adopted and adapted by engineers to achieve this (e.g., business capabilities, bounded-contexts, etc), so no need to reinvent the wheel on this one.

Bounded-Contexts is a design pattern from DDD (Domain Driven Design), and intends to segregate the monolith domain in different smaller domains (i.e., sub-domain). Typically, these smaller domains are identified by aggregate roots and ultimately there is an identification of a microservice per aggregate root (similar to database per service). This segregation is convenient to apply when the domain is quite big and there is time to perform such seperation. Unfortunately, these conditions do not met for this challenge, so applying the bounded-context pattern is rejected. Division by business capabilities on the other hand is rather simpler, not requiring so much time. There are different ways to identify such capabilities, such as through system functionalities/use cases.

The diagram below represents the challenge functionalities connected to each capability:

![business_capabilities_diagram](docs/assets/sword_health_technical_challenge_business_capabilities.png)

<center><i>Figure 2 - Diagram illustrating the business capabilities decomposed by the product functionalities (Tool: draw.io)</i></center>

As seen in the diagram, a total of three business capabilities have been identified: tasks, notifications and authorization management. Each of these capabilities identify a microservice.

Having each microservice identified, it is now possible to apply more design decision in order to strength the system. Typically, microservices each have their own databases, to avoid single point of failures in the data layer (Database per Service). CQRS (Command-Query Responsibility Segregation) could also be applied to reduce latency in the read/write operations, but that's a little bit overkill given the system dimension, as well as the time to develop. API Gateway is a bonus for a more production-ready system, as it serves as a firewall, threshold and load-balancer for the microservices.

Now, before starting development, the only thing missing is designing each of the microservices APIs schemas and decide how deployment occur.

In the tasks services, thinking as a RESTless API, there is the `tasks` collection, allowing for:

- Retrieving all tasks (`GET /tasks`);
- Create a task (`POST /tasks`);
- View a task (`GET /tasks/:id`);
- Update a task (`PUT /tasks/:id`);
- And delete a task (`DELETE /tasks/:id`).

Heading over to the notifications service, there are two APIs:

- Public API for retrieving and viewing non-read notifications, as well as to opt-in for specific notifications (`/notifications` collection);
- Private API for publishing and consuming a task perform event.

Finally, in the authorization service:

- Authenticate (`POST /authenticate`);
- Check if user is authorized to perform certain actions (`GET /permissions`).

-----

For developing the system, I've selected the following libraries and tools:

- [Echo](https://github.com/labstack/echo) as the web framework (serves its purpose, easy to use and had prior experience with it);
- [gorm](https://github.com/go-gorm/gorm) as the "ORM" (also prior experience with);
- RabbitMQ;
- API Gateway... TBD;
- redoc for REST API documentation rendering;
- Swagger Editor for developing REST API documentation.

Golang version: `1.16`

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