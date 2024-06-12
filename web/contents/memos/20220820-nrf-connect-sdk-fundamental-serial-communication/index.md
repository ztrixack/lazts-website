---
title: nRF Connect SDK Fundamental; Serial Communication
slug: 20220820-nrf-connect-sdk-fundamental-serial-communication
excerpt:
featured_image:
published_at: 2022-08-20
last_updated_at: 2024-05-20
published: false
tags:
  - embedded
  - nRF52840
---

## UART

## I2C/TWI

I2C (Inter-Integrated Circuit) เป็นโปรโตคอลการสื่อสารแบบซิงโครนัสสองสายที่ใช้กันอย่างแพร่หลาย มักใช้สำหรับการเชื่อมต่อไมโครคอนโทรลเลอร์กับเซ็นเซอร์และ IC ต่างๆ รวมถึงไมโครคอนโทรลเลอร์อื่นๆ เนื่องจากใช้สายสองเส้น โปรโตคอล I2C จึงมีชื่อเรียกว่า Two-Wire Interface (TWI)

I2C ถูกออกแบบมาเพื่อการเชื่อมต่อในระยะสั้นและภายในแผงวงจรของอุปกรณ์ ส่วนใหญ่ SoC ของ Nordic มีตัวควบคุม I2C อย่างน้อยหนึ่งตัวที่สามารถทำงานได้สองบทบาท คือ เป็น "controller" ที่เริ่มการสื่อสารและควบคุมสัญญาณนาฬิกา หรือเป็น "target" ที่ตอบสนองต่อคำสั่งการสื่อสาร

ตัวควบคุม I2C ส่วนใหญ่ในชิปของ Nordic รองรับหลายความเร็ว: 100 kbps (I2C_BITRATE_STANDARD), 400 kbps (I2C_BITRATE_FAST) และ 1000 kbps (I2C_BITRATE_FAST_PLUS) โดยความเร็วเริ่มต้นคือ 100 kbps

เส้นสองสายของ I2C เรียกว่า serial clock (SCL) และ serial data (SDA) อุปกรณ์ทั้งหมดบนบัสจะเชื่อมต่อกับสายสองเส้นนี้

### การทำงานของ I2C
SCL (Serial Clock Line): สัญญาณนาฬิกาที่ถูกสร้างขึ้นโดย I2C controller เพื่อซิงค์อุปกรณ์ทั้งหมดบนบัสให้เป็นนาฬิกาเดียวกัน
SDA (Serial Data Line): สายส่งข้อมูลที่เป็นสองทิศทาง ข้อมูลสามารถเดินทางได้ทั้งจาก controller ไปยัง target หรือจาก target ไปยัง controller

### ที่อยู่ของอุปกรณ์ I2C
อุปกรณ์ I2C target แต่ละตัวมีที่อยู่ที่ไม่ซ้ำกันเพื่อแยกแยะระหว่างอุปกรณ์ I2C target บนบัสเดียวกัน ที่อยู่มักจะเป็นค่า 7-bit แต่บางอุปกรณ์ I2C target ก็ใช้ค่า 10-bit ด้วย

การตั้งค่า I2C ใน nRF Connect SDK และการใช้งาน API ของ I2C Controller
การเปิดใช้งานไดรเวอร์
เปิดใช้งานไดรเวอร์ I2C โดยเพิ่มบรรทัดต่อไปนี้ในไฟล์การกำหนดค่าแอปพลิเคชัน prj.conf:

```conf
CONFIG_I2C=y
```

รวมไฟล์เฮดเดอร์ของ I2C API ในไฟล์ซอร์สโค้ดของคุณ:

```c
#include <zephyr/drivers/i2c.h>
```

### การตั้งค่าอุปกรณ์
เหมือนกับไดรเวอร์ GPIO ที่เราเรียนรู้ในบทที่ 2 ไดรเวอร์ I2C ทั่วไปใน Zephyr มีโครงสร้างเฉพาะ i2c_dt_spec ซึ่งประกอบด้วยตัวชี้อุปกรณ์สำหรับบัส I2C const struct device *bus และที่อยู่ของ target uint16_t addr.

เพื่อดึงโครงสร้างนี้ เราต้องใช้ฟังก์ชันเฉพาะ API I2C_DT_SPEC_GET()

### การกำหนดค่า I2C Controller
ถ้าเซ็นเซอร์ยังไม่ได้ถูกกำหนดใน devicetree ของบอร์ด คุณต้องเพิ่มเซ็นเซอร์ของคุณเป็นโหนดย่อยของ I2C controller โดยใช้ไฟล์ overlay:

```dts
&i2c0 {
    mysensor: mysensor@4a {
        compatible = "i2c-device";
        reg = <0x4a>;
        label = "MYSENSOR";
    };
};
```

ในตัวอย่างนี้ เซ็นเซอร์เชื่อมต่อกับ I2C controller i2c0 และมีที่อยู่ target เป็น 0x4a

การกำหนดตัวระบุโหนด
ใช้มาโคร DT_NODELABEL() เพื่อดึงสัญลักษณ์ตัวระบุโหนด I2C0_NODE ซึ่งจะแทน I2C hardware controller i2c0:

```c
#define I2C0_NODE DT_NODELABEL(mysensor)
```

I2C0_NODE มีข้อมูลเกี่ยวกับพินที่ใช้สำหรับ SDA และ SCL, แผนที่หน่วยความจำของ I2C controller, ความถี่ I2C เริ่มต้น และที่อยู่ของ target device

การดึงโครงสร้างอุปกรณ์เฉพาะ API
มาโคร I2C_DT_SPEC_GET() คืนค่าโครงสร้าง i2c_dt_spec ซึ่งประกอบด้วยตัวชี้อุปกรณ์สำหรับบัส I2C รวมถึงที่อยู่ target:

```c
static const struct i2c_dt_spec dev_i2c = I2C_DT_SPEC_GET(I2C0_NODE);
```

การตรวจสอบว่าอุปกรณ์พร้อมใช้งาน
ใช้ `device_is_ready()` เพื่อยืนยันว่าอุปกรณ์พร้อมใช้งาน:

```c
if (!device_is_ready(dev_i2c.bus)) {
    printk("I2C bus %s is not ready!\n\r", dev_i2c.bus->name);
    return;
}
```

ตอนนี้เรามีโครงสร้างอุปกรณ์ i2c_dt_spec *dev_i2c ที่สามารถผ่านไปยัง I2C generic API interface เพื่อดำเนินการอ่าน/เขียน

#### การเขียน I2C
ฟังก์ชันที่ง่ายที่สุดในการเขียนไปยังอุปกรณ์ target คือ i2c_write_dt():

```c
uint8_t config[2] = {0x03, 0x8C};
int ret = i2c_write_dt(&dev_i2c, config, sizeof(config));
if (ret != 0) {
    printk("Failed to write to I2C device address %x at reg. %x \n\r", dev_i2c.addr, config[0]);
}
```

#### การอ่าน I2C
ฟังก์ชันที่ง่ายที่สุดในการอ่านจากอุปกรณ์ target คือ i2c_read_dt():

```c
uint8_t data;
int ret = i2c_read_dt(&dev_i2c, &data, sizeof(data));
if (ret != 0) {
    printk("Failed to read from I2C device address %x at reg. %x \n\r", dev_i2c.addr, config[0]);
}
```

#### การอ่านข้อมูล I2C แบบต่อเนื่อง
การอ่านข้อมูลจากหลายๆ รีจิสเตอร์ต่อเนื่องกันสามารถทำได้โดยใช้ฟังก์ชัน `i2c_burst_read_dt()`:

```c
uint8_t rgb_value[6] = {0};
// อ่านต่อเนื่อง 6 bytes เนื่องจากแต่ละช่องสีมีขนาด 2 bytes
ret = i2c_burst_read_dt(&dev_i2c, BH1749_RED_DATA_LSB, rgb_value, sizeof(rgb_value));
```

#### การเขียน/อ่าน I2C
เป็นเรื่องปกติที่จะเขียนที่อยู่ของรีจิสเตอร์ภายในที่จะอ่าน แล้วตามด้วยการอ่านข้อมูลจากรีจิสเตอร์นั้นโดยตรง ฟังก์ชัน `i2c_write_read_dt()` มีรูปแบบการเรียกดังนี้:

```c
uint8_t sensor_regs[2] = {0x02, 0x00};
uint8_t temp_reading[2] = {0};
int ret = i2c_write_read_dt(&dev_i2c, &sensor_regs[0], 1, &temp_reading[0], 1);
if (ret != 0) {
    printk("Failed to write/read I2C device address %x at reg. %x \n\r", dev_i2c.addr, sensor_regs[0]);
}
```

### สรุป
การตั้งค่าและการใช้งาน I2C ใน nRF Connect SDK สามารถทำได้โดยการกำหนดค่า devicetree, เปิดใช้งานไดรเวอร์ใน prj.conf และการใช้งานฟังก์ชัน API เช่น i2c_write_dt(), i2c_read_dt(), i2c_burst_read_dt() และ i2c_write_read_dt() เพื่อสื่อสารกับอุปกรณ์ I2C

การสื่อสารผ่าน I2C ใน nRF Connect SDK สามารถตั้งค่าและใช้งานได้ง่ายโดยการเปิดใช้งานไดรเวอร์ I2C ใน prj.conf และกำหนดค่า devicetree overlay สำหรับพอร์ต I2C การสื่อสารกับเซ็นเซอร์และส่วนประกอบภายนอกสามารถทำได้โดยใช้ฟังก์ชัน i2c_write() และ i2c_read() ในโค้ดแอปพลิเคชันของคุณ

## SPI

## Reference
[Nordic SDK documents](https://docs.nordicsemi.com/bundle/ncs-latest/page/nrf/index.html)
