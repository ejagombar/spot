![spotLogo](https://github.com/ejagombar/spot/assets/77460324/29c96587-abd3-4ac1-a0b9-8e2ad6a5ffb9)

### A minimal commmand line tool for spotify, written in Go
Compatible with Linux, Windows, and MacOS

https://github.com/ejagombar/spot/assets/77460324/17f6dbc7-3f31-4717-9aaf-28d5a572b5d4

# Commands
`play`- Plays the current song

`pause` - Pauses the current song

`next` - Skips to the next song

`back` - Skips back to the previous song

`shuffle` - Enable/Disable shuffle

`song` - Play a specified song

`album` - Play a specified album

`artist` - Play a specified artist

`playlist` - Play a specified playlist

`info` - Show info about the current song, or info about your spotify account

`add` - Add the current song to a specified playlist

`login` - Login to spot with your spotify account

`config` - Configure settings

`help` - Show help for any command


# Installation
### Method 1 - Install from binary
- Visit the releases tab on the left side of the github repository page
- Download the appropriarate file for your OS
- Unzip the file.
- Add the binary to the path.

### Method 2 - Build from source
Prerequisites: Go must be properly installed on your system
- Clone the repository: `git clone git@github.com:ejagombar/spot.git`
- In the spot directory, download the dependancies: `go mod download`
- Build and install spot `go install spot`

# Help

**Step 1:** Go to https://developer.spotify.com/ and login with your spotify account.

**Step 2:** Go to your dashboard (https://developer.spotify.com/dashboard) and click on the "create an app" button.

**Step 3:** Enter an app name and description of your choice. (Anything will do)

**Step 4:** Set the Redirect URL to: http://localhost:8080/callback

**Step 5:** Click 'create'

**Step 6:** You will now be greeted with an overview page. At the top right, click "settings"

**Step 7:** You will see a your client ID and below, a button to reveal your client secret.

**Step 8:** Copy these values into the spot config file. This hidden file should be found in your home directory with the name .spot.json
    Open the file and enter the client ID and client secret in the appropriate boxes then save and close the file.

**Step 9:** run the command "spot login". If everything is done correctly, a link will be generated which you can click to login with your spotify account.

