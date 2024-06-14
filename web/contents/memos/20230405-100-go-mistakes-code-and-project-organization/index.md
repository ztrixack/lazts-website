---
title: 100 Go Mistakes and How to Avoid Them
slug: 20230428-100-go-mistakes-and-how-to-avoid-them
part: 1
excerpt:
featured_image: https://picsum.photos/1024/768?random=1
published_at: 2023-04-05
last_updated_at: 2024-06-12
published: false
tags:
  - 100-go-mistakes
  - development
  - go
---



1. Code and Project Organization
2. [Data Types](https://lazts.com/memos/talents/20170520-100-go-mistakes-data-types)
3. [Control Structures](https://lazts.com/memos/talents/20170520-100-go-mistakes-control-structures)
4. [Strings](https://lazts.com/memos/talents/20170520-100-go-mistakes-strings)
5. [Functions and Methods](https://lazts.com/memos/talents/20170520-100-go-mistakes-functions-and-methods)
6. [Error Management](https://lazts.com/memos/talents/20170520-100-go-mistakes-error-management)
7. [Concurrency: Foundations](https://lazts.com/memos/talents/20170520-100-go-mistakes-concurrency-foundations)
8. [Concurrency: Practice](https://lazts.com/memos/talents/20170520-100-go-mistakes-concurrency-practice)
9. [Standard Library](https://lazts.com/memos/talents/20170520-100-go-mistakes-standard-library)
10. [Testing](https://lazts.com/memos/talents/20170520-100-go-mistakes-testing)
11. [Optimizations](https://lazts.com/memos/talents/20170520-100-go-mistakes-optimizations)

## \#1 สร้าง Variable Shadowing ขึ้นโดยไม่ตั้งใจ

Variable shadowing คือสถานการณ์ที่ตัวแปรใน scope หนึ่งไปทับตัวแปรที่มีชื่อเดียวกันใน scope ที่ใหญ่กว่า ลองดูตัวอย่างโค้ดนี้

```go
func main() {
    // ตัวอย่างของ variable shadowing ในระดับ function scope
    count := 100
    fmt.Println("Initial count:", count) // พิมพ์ค่า 100 ออกมา

    // ใช้ loop ซึ่งจะมีตัวแปรที่ทับซ้อนกัน
    for i := 0; i < 3; i++ {
        // ตัวแปร count ภายใน loop นี้ทับตัวแปร count ภายนอก
        count := i * 10
        fmt.Println("Count inside loop:", count) // จะพิมพ์ค่า 0, 10, 20 ตามลำดับ
    }

    fmt.Println("Count after loop:", count) // จะพิมพ์ค่า 100 ซึ่งเป็น count ภายนอก

    // การใช้ inner function ที่มีตัวแปรทับกัน
    func() {
        count := 50
        fmt.Println("Count inside inner function:", count) // จะพิมพ์ค่า 50 ซึ่งเป็นค่าของ count ภายใน
    }()

    fmt.Println("Count after inner function:", count) // พิมพ์ค่า 100 ออกมา
}
```

ในตัวอย่าง count ใน main function จะเป็นค่า 100 เสมอ ไม่ถูกเปลี่ยนแปลง เพราะ count ที่อื่นๆ จะเป็น count ภายใน scope ตัวเอง ซึ่งอาจทำให้เกิดข้อผิดพลาดหรือสร้างความสับสนให้กับทีมได้ แต่ถ้าเรารู้ตัวดีว่าเรากำลังทำอะไรอยู่ การใช้ variable shadowing ก็ช่วยให้สะดวกเหมือนกัน เช่น การใช้กับ error ที่กลายเป็นท่ามาตรฐานในรูปแบบนี้


```go

```

ดังนั้น ถ้าเลี่ยงได้ก็เลี่ยง เลี่ยงไม่ได้ก็ใช้อย่างระมัดระวัง

## \#2. มี Scope ซ้อนกันโดยไม่จำเป็น

หลักการเขียนโค้ดที่ดี และอ่านเข้าใจง่ายอย่างหนึ่งคือ flow ของโค้ดที่เป็น happy path ควรจะ align ทางซ้ายมือสุดเสมอ เช่น

```go
// ❌ ไม่ควรเขียนแบบนี้
if foo() {
    // ...
    return true
} else {
    // ...
}

// ✅ ควรเขียนแบบนี้ดีกว่า
if foo() {
    // ...
    return true
}
// ...
```

```go
// ❌ ไม่ควรเขียนแบบนี้
if s != "" {
    // ...
} else {
    return errors.New("empty string")
}

// ✅ ควรเขียนแบบนี้ดีกว่า
if foo() {
    // ...
    return true
}
// ...
```

## \#3. ใช้ init() ในทางที่ผิด

init() function เป็น function ชื่อพิเศษที่ go สงวนไว้สำหรับการ initialize หรือตั้งค่าเริ่มต้นของโปรแกรม แต่มันจะไม่รับ arguments และไม่ return ค่าใดๆ ออกมา ทุกอย่างจะจบภายในนั้น ดังนั้นจึงไม่ควรทำอะไรที่ซับซ้อน หรือการจัดการใดๆ ก็ตามที่มี error เข้ามาทำภายในนี้ เช่น init database ควรใช้ในแง่ของการกำหนดค่าคงที่ของพวก configuration มากกว่า

จริงๆ งานทั่วๆ ไปก็ไม่จำเป็นต้องใช้ init() เลย เพราะเราก็จัดการทุกๆ อย่างได้ภายใน main function อยู่แล้ว

## \#4. ใช้ getter/setter pattern

ภาษา Go ออกแบบมาอย่างดีมากๆ แล้วในการ encapulate ข้อมูล และการรับ/ส่งค่าก็มีท่ามาตรฐานเฉพาะกับภาษา Go จึงไม่มีความจำเป็นใดๆ ต้องสร้าง Getter/Setter ยัดเข้าไปใน struct เข้ามาอีก

```go
// ❌ ไม่ควรเขียนแบบนี้
type Data struct {
  value string
}

func (d *Data) SetValue(value string) {
  d.value = value
}

func (d *Data) GetValue() string {
  return value
}

// ✅ ควรเขียนแบบนี้ดีกว่า
type Data struct {
  Value string
}
```

## \#5. การครอบ interface กับทุกอย่าง

อย่าพยายามสร้างความซับซ้อนขึ้นโดยไม่จำเป็น ให้สร้าง interface เฉพาะสิ่งที่เราต้องการ ไม่ต้องสร้างเผื่อใช้ในอนาคต เพราะมันจะทำให้เข้าใจยาก โค้ดรก และ implement ลำบากไปด้วย

concept ของ interface ใน Go ไม่ใช่

## Reference
[100go](https://100go.co/)
