# DeskDay

A lightweight macOS solution for tracking office attendance through network connectivity. DeskDay automatically logs when you're connected to your office network, making it easy to monitor your in-office days.

## How it Works

DeskDay uses macOS's native system monitoring capabilities to:
1. Detect network changes via `/etc/resolv.conf` file changes
2. Check if you're connected to the office network (This comes from env var `OFFICE_SSID`)
3. Log connections to a SQLite database with timestamps

The system runs entirely locally and requires no external services.

## Components

- `netcount.sh`: Main script that is run upon connection, logs SSID and timestamp to SQLite
- `setup.sh`: Installation script that sets up the LaunchAgent for automatic monitoring
- `deskday`: Executable compiled from the go app
- SQLite database: Stores connection timestamps (created automatically)
- LaunchAgent: Ensures the monitoring script runs on network changes (TODO: this can be found at ...) the launch agent `com.user.networkmonitor` can be found at `~/Library/LaunchAgents/`

## Prerequisites

- macOS 
- SQLite3 (should be included in macOS by default)
- Use of terminal to see the output

## Installation

1. Clone this repository:
   ```bash
   git clone github
   cd DeskDay
   ```

2. Run the setup script:
   ```bash
   ./setup.sh
   ```

This will:
- Make the monitoring script executable
- Create and configure the LaunchAgent
- Initialize the SQLite database
- Start the monitoring service

## Database Location

The SQLite database is created at:
TODO: which dir is assets usually stored in
```
~/repo/inoff/network_connections.db
```

## Customization

To monitor a different network, modify the SSID check in `netcount.sh`:
TODO: update to env var

```bash
if [ "$current_ssid" = "YOUR_NETWORK_NAME" ]; then
```

## Troubleshooting

If the monitoring isn't working:

1. Check LaunchAgent status:
   ```bash
   launchctl list | grep networkmonitor
   # or 
   launchctl list | rg networkmonitor
   ```

2. Verify database permissions:
   ```bash
   ls -l ~/repo/inoff/network_connections.db
   ```

3. Check system logs:
   ```bash
   log show --predicate 'processImagePath contains "netcount"'
   ```

## Limitations

- Currently only supports macOS
- Once the script is running in will only track from then on, currently no options to manually add or remove days
- Network detection relies on system profiler output. This is because the [airport](https://support.apple.com/en-au/guide/aputility/aprtc6ff2ed9/mac) software included with Mac seems to be deprecated and the replacement [wdutil](https://ss64.com/mac/wdutil.html) requires sudo.

## Contributing

Feel free to open issues or submit pull requests with improvements.
The code is quite simple so feel free to take what you need and make your own to suit your own needs

## License

MIT License
