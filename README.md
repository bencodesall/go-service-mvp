# ardanlabs-service-2.0

## Projects Setup

### Master Policies

1. Use precision-based semantics
2. Make things easy to **understand**, not easy to **do**

        log.Println("main: error: ", err)
        os.Exit(1)
   vs

        log.Faltalln(v ...interface{})

### Packaging and Package-oriented Design

Rules:

- Where possible, packages should be as independent as possible.
- Reduce horizontal dependencies within a package
- Dependencies may flow down but not up (alphabetically) the layers

### Layered Approach to Architecture

Bill Kennedy espouses a 5 layer approach to application architecture that define macro-level engineering decisions around packaging structure. These layers are represented in the directory structure as follows:

- **app** - All code representing "application level". Binary and presentation concerns. Requests, response, cli tooling, UI/UX related to the functioning of the application
- **business** - Solving all business rules. Databases, external systems, business layer data modeling
- **foundation** - Non-business oriented code. Foundational code that could be reusable. Code that could potentially be moved out into a team/org custom "standard library" (sometimes referred to as a "kit" repo)

### Layer Policy (Engineering Constraint)

### Middleware "Onion"

#### Onion "Layers"

1. Readiness
2. Auth
3. Panic
4. Metrics
5. Errors
6. Logging

**Supporting "wrap" func**

This function takes a set of "midlewares" (functions) and returns them layered from right to left.

```go
func wrapMiddleware(mw []Middleware, handler Handler) Handler {
	for i := len(mw) - 1; i >= 0; i-- {
		h := mw[i]
		if h != nil {
			handler = h(handler)
		}
	}
	return handler
}
```

#### Logging

Logging is not an insurance policy. Logging is a engineering decision that should be precise.

- Avoid "logging levels" and focus on correct signal to noise ratios
- Consider the allocation, GC, and other costs of logging
- Write logs to standard out, and do not use the singleton pattern. Pass an instance down from main.

#### Configuration

- Where possible, use sane defaults allowing overrides where necessary.
- Only the `main` package is allowed to talk to configuration system. Use Config structs and/or parameters.
- All config options need to have flags available at the command line.
- Used config options need to be shown on startup to validate overrides.
- Command-line flags must be implemented for overrides. 
- Command-line should include a `--help` flag to show the options.