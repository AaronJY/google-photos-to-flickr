# Google Photos to Flickr
A Go application that imports all of your Google photos to Flickr.

[![Build Status](https://travis-ci.org/AaronJY/google-photos-to-flickr.svg?branch=master)](https://travis-ci.org/AaronJY/google-photos-to-flickr)

## Disclaimer
This is very much a "hello world" app; I wanted to build something while getting to grips with go. Because of that, the code may not be idiomatic. Feel free to raise an issue to leave feedback or report bugs!

## Dependencies
- Go 1.13
- nodejs (built on 12.11.1)
- npm (built on 6.11.3)

## Setup
1. Make sure you have Go 1.13 installed by following the [official installation guide](https://golang.org/doc/install).
2. Clone this repo!
`git clone https://github.com/AaronJY/google-photos-to-flickr.git`
3. Restore packages using `go mod download`
4. Take a copy of `config.template.yml` and rename it as `config.yml`. Fill out the details with your desired config.
5. In the `app/` directory run the following:

    ```
    npm i
    npm run build
    ```

    (or alternatively, you can run `npm run build:dev` for a development build)
6. Create a new Google developer project and enable use of the Google Photos API (see [this guide](https://developers.google.com/photos/library/guides/get-started) for detailed instructions)
7. Create new Web Application credentials with these details:
   - Add http://localhost:1337 as an Authorized JavaScript Origin (changing the port to whatever you've set the `server.port` as in config.yml)
   - Add http://localhost:1337/api/google/authcallback as an Authorized Redirect URI (again, changing the port as necessary)
8. Copy the Client ID and Client Secret from the newly created credential into their respective properties within config.yml
8. Run `go run main.go`
