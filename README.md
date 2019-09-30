# GPhotos2Flickr (google-photos-to-flickr)
A demo Go application that imports all of your Google photos to Flickr.

## Disclaimer
This is very much a "hello world" app. I wanted to build something while getting to grips with go, and this is it! Because of that, the code may not always meet Go style guidelines - or even make sense. Feel free to raise an issue to leave feedback, or point me in the direction of anything that you feel should be changed :)
## Setup
1. Make sure you have Go installed by following the [official installation guide](https://golang.org/doc/install).
2. Clone this repo into your Go path, **ensuring the root directory is named gPhotosToFlickr** (the directory name is necessary right now for relative paths to be resolved properly)
`git clone https://github.com/AaronJY/google-photos-to-flickr.git gPhotosToFlickr`
3. Install [GoDep](https://github.com/tools/godep) (a tool for dependancy management) by running `go get github.com/tools/godep`
4. Restore packages using GoDep's `go restore` command within your gPhotosToFlickr folder
5. Take a copy of `config.template.yml` and rename it as `config.yml`. Fill out the details within respectively.
