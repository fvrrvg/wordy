# Wordy
## cli game like wordle but not limited to 5-letter words

Wordy is a fun word-guessing game inspired by Wordle. Unlike Wordle, Wordy is not limited to 5-letter words, making it an exciting challenge for word enthusiasts. Test your vocabulary and word-guessing skills while enjoying a relaxing gaming experience.

## Installation

### Build from source
Make sure you have [Go](https://golang.org/) installed. Then
clone the repo
```sh
git clone https://github.com/fvrrvg/wordy
```
Build it.

```sh
cd wordy
go build -o wordy
```

Run

```sh
./wordy
```
### Download pre-built binaries
Download the latest release from [here](https://github.com/fvrrvg/wordy/releases/) and run it.
### Homebrew
```sh
brew tap fvrrvg/tap/wordy
brew install wordy
```

## Usage
* --common (-c): Guess a common word
* --rare (-r): Guess a rare word
* --help (-h): Display the help message
* --rules (-R): Display the rules of the game

## Contributing
If you'd like to contribute to Wordy, feel free to open an issue or submit a pull request.
