# make styles
npx tailwindcss -i ./src/web/styles/main.css -o ./src/web/static/css/styles.css

# build go program
go build -o ./tmp/main ./src/main.go

# run go program
./tmp/main
