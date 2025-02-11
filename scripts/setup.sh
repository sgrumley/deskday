# setup.sh
init_database() {
    sqlite3 ~/.local/share/deskday/network_connections.db "CREATE TABLE IF NOT EXISTS connections (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
        ssid TEXT,
        manual BOOL,
        type TEXT
    );"
}
# Create the LaunchAgent plist file that will trigger the script on each DNS update (network connection, lease renew, ...)
cat << EOF > ~/Library/LaunchAgents/com.user.networkmonitor.plist
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>com.user.networkmonitor</string>
    <key>ProgramArguments</key>
    <array>
        <string>/Users/${USER}/.local/share/deskday/netcount.sh</string>
    </array>
    <key>WatchPaths</key>
    <array>
        <string>/var/run/resolv.conf</string>
    </array>
    <key>RunAtLoad</key>
    <true/>
</dict>
</plist>
EOF

# Make the script executable
mkdir -p ~/.local/share/deskday
cp -f ./scripts/netcount.sh ~/.local/share/deskday/netcount.sh
chmod +x ~/.local/share/deskday/netcount.sh
init_database

# Load the LaunchAgent
launchctl unload -w  ~/Library/LaunchAgents/com.user.networkmonitor.plist  2>/dev/null || true
launchctl load -w   ~/Library/LaunchAgents/com.user.networkmonitor.plist


