# setup.sh
# Create the LaunchAgent plist file that will trigger the script on each DNS update (network connection, lease renew, ...)
cat << 'EOF' > ~/Library/LaunchAgents/com.user.networkmonitor.plist
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
        <string>/etc/resolv.conf</string>
    </array>
    <key>RunAtLoad</key>
    <true/>
</dict>
</plist>
EOF

# Make the script executable
mv ./scripts/netcount.sh ~/.local/share/deskday/netcount.sh
chmod +x ~/.local/share/deskday/netcount.sh

# Load the LaunchAgent
launchctl unload -w  ~/Library/LaunchAgents/com.user.networkmonitor.plist  2>/dev/null || true
launchctl load -w   ~/Library/LaunchAgents/com.user.networkmonitor.plist
