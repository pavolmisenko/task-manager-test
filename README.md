# FAKE TAsk management eXperIence

This is trainig repository to get familiar with `go + HTMX` for web development.

Styling is done via [TaiwindCSS](https://tailwindui.com/) and [daisyUI](https://daisyui.com/)

# Setup

Make sure you have Node installed. See instalation guide [here](https://nodejs.org/en/download/package-manager)

Make sure you have Golang installed. See instalation guide [here](https://go.dev/doc/install)

Install dependencies

```bash
npm install
```

# Run application

Run DB

```bash
cd db
docker-compose up
```

## Run Development server

Install [air](https://pkg.go.dev/github.com/air-verse/air) for hot-reload on go server

```bash
go install github.com/air-verse/air@latest
```

Run air in root directory

```bash
air
```

After building, application should be available on `localhost:7070`

# Bonus: How to setup TailwindCSS with repository

### NodeJS (with daisyUI)

1. [install tailwind](https://tailwindcss.com/docs/installation)
2. [install daisyUI](https://daisyui.com/docs/install/)

### Alternative: Standalone CLI tailwind without npm dependencies (just tailwind UI library)

Approach made by [this](https://tailwindcss.com/blog/standalone-cli) process to enable TailwindCSS.

This only enables tailwindCSS without out of 3rd party libraries such as daisyUI
