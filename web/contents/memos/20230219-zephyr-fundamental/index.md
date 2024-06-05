---
title: Zephyr Fundamental - Get Started
slug: 20230219-zephyr-fundamental
excerpt:
featured_image: 0000.jpeg
published_at: 2023-02-19
last_updated_at: 2024-05-20
published: false
tags:
  - embedded
  - zephyr
---

# Zephyr Fundamental - Get Started

หนึ่งในโลกของการพัฒนา firmware และ microcontrollers โดยเฉพาะกับชิป nRF52840 และ RTOS

## Zephyr คืออะไร?

Zephyr (อ่านว่า เซฟเฟอร์) เป็น open-source real-time operating system (RTOS) ที่ออกแบบมาเพื่อการพัฒนา firmware สำหรับอุปกรณ์ embedded systems โดยเฉพาะ จุดเด่นก็คือขนาดที่เล็กทำให้เหมาะสมกับอุปกรณ์ที่มีทรัพยากรจำกัด

Zephyr รองรับชิปเซ็ตจากผู้ผลิตหลากหลายราย เช่น Intel, NXP, STMicroelectronics และ Nordic Semiconductor ดังนั้นไม่ว่าจะใช้ชิปอะไร Zephyr ก็น่าจะรองรับได้แน่นอน โดยมาพร้อมกับ SDK และเครื่องมือพัฒนาที่ครบครัน ไม่ว่าจะเป็นการสร้าง project การ compile และการ debug

ด้าน Security เองก็มีการพัฒนาระบบที่เข้มงวด รวมถึงการทดสอบอย่างละเอียดเพื่อให้มั่นใจว่า firmware มีความเสถียรและปลอดภัยมากขึ้น

## การเริ่มต้นใช้งาน Zephyr

ในบทความนี้จะติดตั้งลง MacOS Apple Silicon หากเป็นระบบปฏิบัติการอื่นๆ ก็สามารถดูวิธีการลงอย่างละเอียดได้จาก [Link](https://docs.zephyrproject.org/latest/develop/getting_started/index.html)

### 0. เตรียมเครื่องมือก่อนติดตั้ง

ในการติดตั้ง Zephyr SDK เราจะต้องเตรียมเครื่องมือและซอฟต์แวร์ที่จำเป็นไว้ก่อนนั่นคือ Python, CMake, Ninja และ GCC ARM Embedded Toolchain

สำหรับวิธีการติดตั้งนั้นก็ขึ้นอยู่กับของแต่ละระบบปฏิบัติการ สามารถใช้คำสั่งได้ตามนี้

MacOS ผ่าน Homebrew

```shell
brew install cmake ninja gperf python3 ccache qemu dtc libmagic wget openocd
(echo; echo 'export PATH="'$(brew --prefix)'/opt/python/libexec/bin:$PATH"') >> ~/.zprofile
source ~/.zprofile
```

หลังจากนั้นเราจะทำการติดตั้ง west ผ่าน virtual environment กัน

### 1. สร้าง Virtual Environment

เริ่มจากการสร้าง venv ใน zephyrproject และ activate มันขึ้นมา

```shell
python3 -m venv ~/zephyrproject/.venv
source ~/zephyrproject/.venv/bin/activate
```

ถ้า activate แล้วจะมี (.env) อยู่ข้างหน้า shell

ถัดมาให้ install west ผ่าน pip และดึง source code ทั้งหมดลงมาที่ zephyrproject (นานพอสมควร)

```shell
pip install west
west init ~/zephyrproject
cd ~/zephyrproject
west update
west zephyr-export
pip install -r ~/zephyrproject/zephyr/scripts/requirements.txt
```

### 2. ติดตั้ง Zephyr SDK

Zephyr นั้นเป็น open-source ที่มี repository อยู่ที่ GitHub โดยเราสามารถดู release version ได้จาก [Link](https://github.com/zephyrproject-rtos/sdk-ng/releases)

ให้เลือกดาวโหลดตาม OS ที่ใช้ลงมา หรือใช้ wget หรือ curl แบบนี้ โดยในบทความนี้จะใช้ version 0.16.6

```shell
cd ~
curl -L -O https://github.com/zephyrproject-rtos/sdk-ng/releases/download/v0.16.6/zephyr-sdk-0.16.6_macos-aarch64.tar.xz
curl -L https://github.com/zephyrproject-rtos/sdk-ng/releases/download/v0.16.6/sha256.sum | shasum --check --ignore-missing
```

แตกไฟล์ด้วยคำสั่ง tar แล้วรอซักครู่ แล้วเข้าไปรันไฟล์ setup.sh ตอบ y ไปให้หมด

```shell
tar xvf zephyr-sdk-0.16.6_macos-aarch64.tar.xz
cd zephyr-sdk-0.16.6
./setup.sh
```

ณ ตอนนี้เราได้ zephyr-sdk ลงเครื่องเรียบร้อยแล้ว ถัดมาเราจะต้องตั้งค่า environment variables ที่จำเป็นสำหรับการทำงานของ Zephyr SDK ดังนี้

```shell
(echo; echo 'export ZEPHYR_TOOLCHAIN_VARIANT="zephyr"';echo 'export ZEPHYR_SDK_INSTALL_DIR="~/zephyrproject/zephyr-sdk-0.16.6"') >> ~/.zprofile
source ~/.zprofile
```

### 3. ทดลองการ build และ flash

ก่อนจะเริ่มทำอะไร เราจำเป็นต้องทดสอบดูก่อนว่าที่เราลงไปนั้นถูกต้องหรือไม่ ใช้งานได้รึเปล่า ด้วยการ build และ flash sample firmwware ลงไปก่อน โดยสามารถที่จะใช้ samples/basic/blinky หรือ samples/hello_world มาลองก็ได้ แต่สิ่งที่สำคัญเลยก็คือ เราจะต้องรูปว่า board ที่เราใช้เป็น board อะไร และใช้ ID อะไร โดยสามารถดูได้จาก [Link](https://docs.zephyrproject.org/latest/boards/index.html#boards)

```shell
cd ~/zephyrproject/zephyr
west build -p always -b nrf52840dk_nrf52840 samples/basic/blinky
west flash
```

ถ้าไฟ LED1 กระพริบก็ถือว่าจบภารกิจแล้ว :D

## Reference

[Getting Started Guide](https://docs.zephyrproject.org/latest/develop/getting_started/index.html)
