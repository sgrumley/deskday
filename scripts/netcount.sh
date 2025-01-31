#!/bin/bash
# netcount.sh

# Function to create the SQLite database if it doesn't exist
init_database() {
    sqlite3 ~/.local/share/deskday/network_connections.db "CREATE TABLE IF NOT EXISTS connections (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
        ssid TEXT,
        manual BOOL,
        type TEXT
    );"
}

# Function to get current SSID (ideally this would use wdutil but sudo makes it hard)
get_ssid() {
    system_profiler SPAirPortDataType | awk '/Current Network/ {getline;$1=$1;print $0 | "tr -d ':'";exit}'
}

# Check if OFFICE_SSID is set
if [ -z "${OFFICE_SSID}" ]; then
    echo "Error: OFFICE_SSID environment variable not set"
    exit 1
fi

# Initialize the database
init_database

# Get current SSID
current_ssid=$(get_ssid)

if [ "$current_ssid" = "${OFFICE_SSID}"]; then
    sqlite3 ~/repo/inoff/network_connections.db "INSERT INTO connections (ssid) VALUES ('${current_ssid}');"
fi
