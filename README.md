# Network Traffic Analyzer (Go)

## Design Pattern

- Following "Abstract Factory" Design Pattern. (https://refactoring.guru/design-patterns/abstract-factory/go/example#example-0)

## Requirements

- Tooling: Linter, Formatter
- Feature: Command Line Interface
    - https://github.com/spf13/cobra
- Feature: Packet Capturing
    - Real-Time Packet Capture
    - Packet Filtering
        - IP Addresses
        - Port Numbers
        - Protocols (TCP, UDP, ICMP)
- Feature: Traffic Analysis
    - Protocol Analysis: HTTP, HTTPS, FTP, TCP, UDP
    - Bandwidth Usage
    - Connection Tracking
- Feature: Statistics and Metrics
    - Dashboard: Real-Time Statistics
    - Data Visualization: Graphs and Charts
    - Alerts and Notifications
- Feature: Network Security Features
    - Anomaly Detection
    - Packet Logging and Inspection
- Feature: Data Storage
    - Local Storage of Traffic Data.
    - Export and Import Capabilities

---
