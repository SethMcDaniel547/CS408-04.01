# 1. Clone the Repository
git clone https://github.com/SethMcDaniel547/CS408-04.01.git

# 2. Update Root Folder
The root folder needs to be named HelloWorld

# 3. Install Go Dependencies
go mod tidy

# 4. Install Node.js Dependencies
npm install

# 5. Set Up Environment Variables
Create a .env file in the root directory and
add your canvas api token with
CANVAS_API_TOKEN=

# 6. Configure Tailwind CSS (if needed)
npm run build

# 7.  Generate templ files
templ generate

# 8. Run the Application
## Run in development mode
go run cmd/main.go

## Or run the built binary
./main