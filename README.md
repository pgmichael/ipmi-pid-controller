# IPMI PID Controller
## Overview
This command-line interface (CLI) tool automatically adjusts the fan speed of a Dell Precision Rack 7910 to maintain a reasonable noise level. It uses IPMI for direct hardware control, applying a proportional-integral-derivative (PID) logic based on temperature readings.

> ⚠️ Use this tool at your own risk and make sure your machine has fail-safes in place to prevent overheating. Adjusting fan speeds can lead to hardware damage if not done properly.

## Features
- Automatic fan speed adjustments based on temperature.
- Customizable temperature targets and fan speed limits.
- IPMI/iDRAC integration for direct server management.

## Requirements
- `ipmitool` installed on your system.
- Network access to your server's IPMI/iDRAC interface.

## Usage
First, ensure you have `ipmitool` installed and accessible in your path. Then, you can run the tool with the following flags:

```shell
./ipmi-pid-controller -address <IPMI_ADDRESS> -user <IPMI_USER> -password <IPMI_PASSWORD> -interval <CHECK_INTERVAL> -t0 <TEMP_TARGET> -k <CURVE_STEEPNESS> -L <MAX_FAN_SPEED> -minFanSpeed <MIN_FAN_SPEED>
```

For example, to run the tool with default settings:

```shell
./ipmi-pid-controller
```

### Flags
- `address`: IPMI/iDRAC address. Default is `192.168.2.200`.
- `user`: IPMI user. Default is `root`.
- `password`: IPMI password. Default is `calvin`.
- `interval`: Interval in seconds for updating fan speed. Default is `10`.
- `t0`: Midpoint temperature in Celsius for fan speed adjustment. Default is `92.0`.
- `k`: Steepness of the temperature-to-fan speed curve. Default is `0.5`.
- `L`: Maximum fan speed as a percentage. Default is `100`.
- `minFanSpeed`: Minimum fan speed as a percentage. Default is `5`.

## License
This project is open-source and available under the MIT License.