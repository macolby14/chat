# Chat Server

- Build a chat server

## Goals

- Users can open a browser and be assigned a unique identity
- Users can enter text in a window and all other users will see it

## Flow Diagram and Class Diagram

- https://app.diagrams.net/

## Steps

- [x] BE- Build a go server that can respond to an endpoint
- [x] BE/FE - Server sends HTML and JS to FE
- [x] FE - Page shows a user list, chat log, and input window
- [x] FE/BE - Websocket connection established on page load
- [x] BE - Lobby joined on websocket connection
- [x] BE - Message log updaetd when user joins and sent to all users
- [ ] BE - Create a session cookie to assign to a user when joining
- [ ] BE - Create a random username
- [ ] FE - FE updates on new user list

## Frontend

- Allow the frontend to send html and javascript in a templating language
- Allow the user to enter a message and hit a submit button

## Connectivity

- Create a web socket upon opening the browser that will create the user and join the lobby

## Backend

- Implement the actions the server will take to create the user, join the lobby, and notify all users of the updated list

## Frontend

- show the updated user list

## Proxying

- https://create-react-app.dev/docs/proxying-api-requests-in-development/
- Simple proxy config for websockets does not work https://stackoverflow.com/questions/58088218/websockets-in-create-react-app-with-webpack-proxy/67177366#67177366
