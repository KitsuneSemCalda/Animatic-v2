# Animatic-v2

## Description

Animatic-v2 is a Go-based application designed to search and download anime episodes from the web. It provides a command-line interface for users to input the name of the anime they wish to download.

## Features

- Checks for internet connection at startup.
- Searches for the requested anime on AnimeFire.
- Downloads all episodes of the selected anime.
- Provides error messages for various failure cases such as lack of internet connection, failure to locate the anime, and inability to access episode URLs.

## Installation

>[!WARNING]
> Do you need go 1.21.5 installed in your system

To install this project we need follow this steps:

- Compile the project in your enviroment:

```bash
go build
```

- Setting Path to new local:

```bash
export PATH=$PATH:~/.local/bin
```

- Move the project from some directory in your path

```bash
mv animatic-v2 ~/.local/bin/animatic
```

- Re-open your shell

## Usage

>[!NOTE]
> This script only download episodes to Plex/Jellyfin and playing the anime is your responsability.

>[!NOTE]
> The downloaded episodes is only pt-br subtitles and dub

To use this code is simple:

- Run in your shell:

```bash
animatic
```

The code open the prompt-ui label to you write the anime name to be downloaded.

After write the name press enter and the code list animes with related names, use the arrow from select the anime and press enter.

Before this is await the download ending

## Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the [MIT License](https://github.com/KitsuneSemCalda/Animatic-v2/tree/master/LICENSE).
