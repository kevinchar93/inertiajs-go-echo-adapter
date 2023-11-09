# inertiajs-go-echo-adapter

A simple Pokédex that shows off a partial implementation of the [Inertia.js](https://inertiajs.com) protocol using the Go programming language with the [Echo web framework.](https://echo.labstack.com)

This project was made along with the video ["Inertia.js - How it works"](https://www.youtube.com/watch?v=MmQgm9EGP_U) to illustrate ***in code*** how Inertia.js works.

There is a video series ["Let's Build An Inertia.js Adapter in Go"](https://www.youtube.com/watch?v=2THwHiRf8Vg&list=PLSkOWppXkXWRELTBMtwLenv88_Q-aJVcm), that shows the building of this project with commentary.

## Prerequisites

- You must have an installation of **the Go programming language**, this project used `go version go1.21.1 darwin/arm64` at the time it was created
    - Find out how to [install it here](https://go.dev/doc/install)
- You must have an installation of Node.js, this project used `Node.js v18.17.1` at the time it was created
    - Find out how to [install it here](https://nodejs.org/en/learn/getting-started/how-to-install-nodejs)

## Running the example app

- Navigate to the `views` directory & run `npm install`

If you are on a Mac/Unix/Linux machine run `make dev` in the project root to start the vite dev server & go dev server simultaneously; you can also start them individually by running `make dev-client` & `make dev-server` respectively

If you are on a Windows machine you can start the vite dev server by running `$env:BUILD_ENV='development'; npm --prefix ./views run dev; Remove-Item Env:\BUILD_ENV` in the project root; you can start the go dev server by running `$env:BUILD_ENV='development'; go run main.go; Remove-Item Env:\BUILD_ENV` in the project root.

Note that ***both*** the vite dev server & go dev server need to be running for this example to work.

Navigate to [http:localhost:3000](http:localhost:3000) & you should see the example Inertia.js app running

## Copyright Notice

All images in repository & the word "Pokédex" are copyrighted by the Pokémon Company and its affiliates. This repository uses them exclusively for educational purposes to build an example "Pokédex" while teching about the Inertia.js library.

The **pokedex.json** file is from the repository ["fanzeyi/pokemon.json"](https://github.com/fanzeyi/pokemon.json), which itself 
a compilation of data collected by the editors of [Bulbapedia](https://bulbapedia.bulbagarden.net/wiki/Main_Page).