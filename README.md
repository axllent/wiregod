# Wiregod - a [Wireguard](https://www.wireguard.com/) client helper for Linux

```
Wiregod is a helper utility for handling one or more Wireguard client configurations.
Issues & documentation: https://github.com/axllent/wiregod

Usage:
  wiregod
  wiregod [command]

Available Commands:
  down        Take a Wireguard interface offline
  ip          Show your public IP
  status      View current wireguard connection status
  up          Bring a Wireguard interface online
  version     Display the app version & update information

Use "wiregod [command] --help" for more information about a command.
```

## Usage:

- `wiregod up` - Bring a Wireguard interface online. If you have more than one Wireguard config then you are presented a choice.
- `wiregod down` - Take a Wireguard interface offline. If you have more than one Wireguard config then you are presented a choice.


## Requirements

- `sudo`, `wireguard`, `wireguard-tools` installed
- At least one working wireguard configuration