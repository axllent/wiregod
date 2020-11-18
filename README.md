# Wiregod - a [WireGuard](https://www.wireguard.com/) client helper for Linux

```
Wiregod is a helper utility for handling one or more WireGuard client configurations.
Issues & documentation: https://github.com/axllent/wiregod

Usage:
  wiregod
  wiregod [command]

Available Commands:
  down        Take a WireGuard interface offline
  ip          Show your public IP
  status      View current wireguard connection status
  up          Bring a WireGuard interface online
  version     Display the app version & update information

Use "wiregod [command] --help" for more information about a command.
```

## Usage:

- `wiregod up` - Bring a WireGuard interface online. If you have more than one WireGuard config then you are presented a choice.
- `wiregod down` - Take a WireGuard interface offline. If you have more than one WireGuard config then you are presented a choice.


## Requirements

- `sudo`, `wireguard`, `wireguard-tools` installed
- At least one working wireguard configuration