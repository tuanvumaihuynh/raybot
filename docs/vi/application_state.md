# Trạng Thái Ứng Dụng

Chi tiết trạng thái robot mà ứng dụng quản lý.

## BatteryState (Trạng Thái Pin)
| Trường dữ liệu | Kiểu  | Mô tả                                  |
|---------------|-------|----------------------------------------|
| Current       | int   | Dòng điện của pin tính bằng mA         |
| Temp          | int   | Nhiệt độ của pin tính bằng °C          |
| Voltage       | int   | Điện áp của pin tính bằng mV           |
| CellVoltages  | []int | Điện áp của từng cell tính bằng mV     |
| Percent       | uint8 | Phần trăm pin                          |
| Fault         | uint8 | Trạng thái lỗi pin                     |
| Health        | uint8 | Trạng thái sức khỏe pin                |
| Status        | int   | Trạng thái chung của pin               |

## ChargeState (Trạng Thái Sạc)
| Trường dữ liệu | Kiểu  | Mô tả                                  |
|---------------|-------|----------------------------------------|
| CurrentLimit  | int   | Giới hạn dòng điện sạc tính bằng mA    |
| Enabled       | bool  | Cho phép sạc hay không                 |

## DischargeState (Trạng Thái Xả Pin)
| Trường dữ liệu | Kiểu  | Mô tả                                  |
|---------------|-------|----------------------------------------|
| CurrentLimit  | int   | Giới hạn dòng điện xả tính bằng mA     |
| Enabled       | bool  | Cho phép xả pin hay không              |

## DistanceSensorState (Trạng Thái Cảm Biến Khoảng Cách)
| Trường dữ liệu | Kiểu  | Mô tả                                  |
|---------------|-------|----------------------------------------|
| FrontDistance | int   | Khoảng cách đến phía trước robot (cm)  |
| BackDistance  | int   | Khoảng cách đến phía sau robot (cm)    |
| DownDistance  | int   | Khoảng cách đến mặt đất (cm)           |

### LiftMotorDirection (Hướng Động Cơ Nâng)
| Giá trị                  | Mô tả                  |
|--------------------------|------------------------|
| LiftMotorDirectionUp     | Nâng lên               |
| LiftMotorDirectionDown   | Hạ xuống               |

## LiftMotorState (Trạng Thái Động Cơ Nâng)
| Trường dữ liệu | Kiểu                | Mô tả                                       |
|---------------|---------------------|---------------------------------------------|
| Direction     | LiftMotorDirection  | Hướng của động cơ nâng (Lên/Xuống)          |
| Speed         | uint8               | Tốc độ của động cơ nâng (%)                 |

### DriveMotorDirection (Hướng Động Cơ Di Chuyển)
| Giá trị                     | Mô tả                  |
|-----------------------------|------------------------|
| DriveMotorDirectionForward  | Di chuyển tiến         |
| DriveMotorDirectionBackward | Di chuyển lùi          |

## DriveMotorState (Trạng Thái Động Cơ Di Chuyển)
| Trường dữ liệu | Kiểu                | Mô tả                                       |
|---------------|---------------------|---------------------------------------------|
| Direction     | DriveMotorDirection | Hướng của động cơ di chuyển (Tiến/Lùi)      |
| Speed         | uint8               | Tốc độ của động cơ di chuyển (%)            |
