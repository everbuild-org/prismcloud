# PrismCloud - the next generation of cloud servers

PrismCloud is a cloud server orchestrator that allows you to run multiple servers on a single machine. PrismCloud is
designed to be easy to use and easy to manage. PrismCloud is built on top of Docker and Kubernetes, which means that it
is highly scalable and highly available.


## Usage (Dev)

- Have a Kubernetes cluster running and accessible from your local machine via `kubectl`.
- Run `docker compose up -d` to start all services required for development.
- Run `make` to regenerate protobuf files.
- Start the apiserver with `go run apiserver --ooc` (ooc is for "out of cluster").
- Run `source ./.alias.sh` to set up aliases for the CLI.
- Run `prism` to see the CLI help.


## Deploy something

If you have port 8085 open on your local machine, you can deploy a simple website with the following command in the [example-deployments/simple_website](example-deployments/simple_website) directory:

```
prism apply .
```

This will create a namespace called `tutorial-website`, a pod called `website-example`, and an LB Ingress called `website-ingress`.

You can then access the website at `http://localhost:8085`.


## Remove the website again

To delete the website, run:

```
prism delete .
```

This will delete the namespace, pod, and LB Ingress.

## License

PrismCloud is licensed under the MIT license. See [LICENSE](LICENSE) for more information.

## Contributing

We welcome contributions to PrismCloud. If you have an idea for a new feature or bug fix, please open an issue or submit a pull request.

---
Prismcloud is a project by [everbuild.org](https://everbuild.org).
