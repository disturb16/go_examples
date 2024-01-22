# Go Project structure (template)

How to build docker container for "advanced" project

```shell
docker build -t myproject-api -f advanced/Dockerfile --build-arg MAIN_FOLDER=api -p 8080:8080 .
```
or
```shell
docker build -t myproject-api -f advanced/Dockerfile --build-arg MAIN_FOLDER=worker .
```

1. Set a tag using the `-t` shortcode
2. Select Dockerfile using `-f`
3. Pass the type of project using a `--build-arg`
    - It has to match the name of the folders `/api` or `/worker`
4. Map a port if needed

Note:
The docker build command should be executed at the `root` of repository.
