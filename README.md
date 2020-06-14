# GO IP Address History

This application creates a log file with your current IP Address. You can add it to your startup applications and generate an IP history log.

## How to use?

It's simple. Just add the binary file to your starting applications of your operative system. Or anything else, like crontab task, etc.

## What this application does?

On execution, the application will call the `ipinfo.io` API recovering your public IP Address. Then, it will store your address on the `history.txt` file of the binary folder.

> File creation is managed by the application.

## Development

This application was written in the Go language as a practice project. It contains two "modules".
- **AddressRecover** who manages the responsibility of getting your IP Address. 
- **Logger** who manages console outputs and file logging.