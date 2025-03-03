# Robot State
Detail state of robot that application manages.
## BatteryState
| Field        | Type  | Description                                  |
|-------------|-------|----------------------------------------------|
| Current     | int   | Current of the battery in mA                |
| Temp        | int   | Temperature of the battery in °C            |
| Voltage     | int   | Voltage of the battery in mV                |
| CellVoltages| []int | Voltage of each cell in mV                  |
| Percent     | uint8 | Battery percentage                          |
| Fault       | uint8 | Battery fault status                        |
| Health      | uint8 | Battery health status                       |
| Status      | int   | General status of the battery               |

## ChargeState
| Field         | Type  | Description                                  |
|--------------|-------|----------------------------------------------|
| CurrentLimit | int   | Current limit of the charge in mA          |
| Enabled      | bool  | Whether the charge is enabled               |

## DischargeState
| Field         | Type  | Description                                  |
|--------------|-------|----------------------------------------------|
| CurrentLimit | int   | Current limit of the discharge in mA       |
| Enabled      | bool  | Whether the discharge is enabled            |

## DistanceSensorState
| Field         | Type  | Description                                  |
|--------------|-------|----------------------------------------------|
| FrontDistance | int   | Distance to the front of the robot in cm   |
| BackDistance  | int   | Distance to the back of the robot in cm    |
| DownDistance  | int   | Distance to the bottom of the robot in cm  |

### LiftMotorDirection
| Value | Description          |
|-------|----------------------|
| 0     | Lift moving up       |
| 1     | Lift moving down     |

## LiftMotorState
| Field      | Type                | Description                                       |
|-----------|---------------------|---------------------------------------------------|
| Direction | LiftMotorDirection  | Direction of the lift motor (Up/Down)            |
| Speed     | uint8               | Speed of the lift motor in %                     |

### DriveMotorDirection
| Value | Description          |
|-------|----------------------|
| 0     | Moving forward    |
| 1     | Moving backward   |

## DriveMotorState
| Field      | Type                | Description                                       |
|-----------|---------------------|---------------------------------------------------|
| Direction | DriveMotorDirection | Direction of the drive motor (Forward/Backward)  |
| Speed     | uint8               | Speed of the drive motor in %                     |
