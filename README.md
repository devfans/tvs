# TVS - IPTV Source Aggregator

A simple Go tool to convert text-based IPTV source lists into valid M3U playlists.

## Features

- Converts simple text files to M3U format
- Supports comments in input files
- Handles basic channel name and URL parsing
- Generates standard-compliant M3U playlists

## Installation

```bash
go install github.com/stefan/tvs@latest
```

Or clone and build manually:

```bash
git clone https://github.com/stefan/tvs.git
cd tvs
go build
```

## Usage

1. Create a text file with your IPTV sources in the following format:
   ```
   channel_name,stream_url
   ```

2. Run the tool:
   ```bash
   ./tvs input.txt output.m3u
   ```

## Input File Format

The input file should contain one channel per line in the format:
```
channel_name,stream_url
```

Example (sample.txt):
```
BBC One,http://example.com/bbc1/index.m3u8
BBC Two,http://example.com/bbc2/index.m3u8
```

Lines starting with # are treated as comments and will be ignored.

## License

MIT