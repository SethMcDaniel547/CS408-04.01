It submitted itself!

# 1. Clone the Repository
git clone https://github.com/SethMcDaniel547/CS408-04.01.git

# 2. Install Go Dependencies
go mod tidy

# 3. Install Node.js Dependencies
npm install

# 4. Set Up Environment Variables
update the .env file in the root directory by removing the .example
portion of .env.example and add your canvas api token here
CANVAS_API_TOKEN=

# 5.  Generate templ files
templ generate

# 6. Run the Application
## Run in development mode
go run cmd/main.go

## Or run the built binary
./main