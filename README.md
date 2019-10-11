# GPhotos2Flickr (google-photos-to-flickr)
A Go application that imports all of your Google photos to Flickr.

[![Build Status](https://travis-ci.org/AaronJY/google-photos-to-flickr.svg?branch=master)](https://travis-ci.org/AaronJY/google-photos-to-flickr)

## Disclaimer
This is very much a "hello world" app. I wanted to build something while getting to grips with go, and this is it! Because of that, the code may not always meet Go style guidelines - or even make sense. Feel free to raise an issue to leave feedback, or point me in the direction of anything that you feel should be changed :)

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
5. In `app/`, run the following:

```
npm i
npm run build
```

(or alternatively, you can run `npm run build:dev` for a development build)

5. Run `go run main.go` to run the app!
