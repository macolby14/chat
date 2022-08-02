# Chat Server
- Build a chat server



## Goals
- Users can open a browser and be assigned a unique identity
- Users can enter text in a window and all other users will see it

## Flow Diagram and Class Diagram
- https://app.diagrams.net/


## Steps
- [ ] BE-  Build a go server that can respond to an endpoint
- [ ] BE/FE - Server sends HTML and JS to FE
- [ ] FE - Page shows a user list, chat log, and input window
- [ ] FE/BE - Websocket connection established on page load
- [ ] BE - User created and lobby joined on websocket connection
- [ ] BE - Message log updaetd when user joins and sent to all users
- [ ] FE - FE updates on new user list


## Frontend
- Allow the frontend to send html and javascript in a templating language
- Allow the user to enter a  message and hit a submit button

## Connectivity
- Create a web socket upon opening the browser that will create the user and join the lobby


## Backend
- Implement the actions the server will take to create the user, join the lobby, and notify all users of the updated list

## Frontend
- show the updated user list


