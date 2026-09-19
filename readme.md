# Mini Lab
This site (once setup with your canvas api token) allows you to navigate your courses, assignments, and even submit a link.

# Setup Instructions
## 1. Clone the Repository
git clone https://github.com/SethMcDaniel547/CS408-04.01.git

## 2. Install Go Dependencies
go mod tidy

## 3. Install Node.js Dependencies
npm install

## 4. Set Up Environment Variables
update the .env file in the root directory by removing the .example
portion of .env.example and add your canvas api token here
CANVAS_API_TOKEN=

## 5.  Generate templ files
templ generate

## 6. Run the Application
### Run in development mode
go run cmd/main.go

# API Endpoints Used
/api/v1/courses

This is where I get all the courses you are in	

/api/v1/courses/:id/assignments

I get all your assignments for the selected course from here

/api/v1/.../submissions

This is what I used to submit the assignment. You can just put the link in the form it makes and click submit.

# Lessons and what I would do better
Overall this wasn't so difficult.. Except for getting it working initially took some time. My first attempt was so unsuccessful I had restarted to get it working since I messed it up so much.

I used a handler system since my stack would not let my real backend go code talk with my front end at all, so I had to make this middle step which took me a long time to figure out how to get working.


The way this currently works, it only takes a link (and it doesnt look very nice). So if I could I would make it look better and either accept more than just links, or on assignments that need more than a link give a message saying something like "unsupported submission type"