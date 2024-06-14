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

ช่วงนี้กำลังสนใจภาษา Go อยู่พอดี ก็ไปเจอกับ repository นึงที่น่าสนใจมากๆ นั่นก็คือ [100 go mistakes](https://github.com/teivah/100-go-mistakes) หรือ [100go](https://100go.co/) ซึ่งเป็นเว็บที่รวบรวมข้อผิดพลาดต่างๆ ที่ชาว Gopher รุ่นใหม่ๆ มักจะทำผิดพลาดกันในรูปแบบต่างๆ หรือจะเป็นการเขียนโค้ดท่าต่างๆ ที่เราควรจะหลีกเลี่ยงมัน ซึ่งทางเว็บเองก็จะเป็นการสรุปมาจากหนังสือ 100 Go Mistakes and How to Avoid Them อีกทีหนึ่ง ซึ่งคนเขียนก็เป็น senior software engineer จาก Google บ้านเกิดของภาษา Go นั่นเอง ใครอยากอ่านตั้งแต่ต้นทางก็ไปหามาอ่านกันได้เลย

พออ่านไปซักพักก็เข้าตัวเยอะเหมือนกัน 🤣 รวมไปถึงบางคำศัพท์ก็ไม่ค่อยเข้าใจซักเท่าไร จะไปขอ contribute เพื่อทำแปลเป็นภาษาไทยเผื่อคนอื่นๆ ไปด้วยก็ไม่น่าไหว  ก็เลยพยายามเขียนสรุปตามความเข้าใจในภาษาของตัวเองดีกว่า จะได้เป็นการทำความเข้าใจและย่นเวลาการกลับมาอ่านทบทวนไปด้วย

แต่เนื่องจากมีถึง 100 ข้อ มันคงจะทำให้บทความนี้ยาวววววมาก ดังนั้นจะขอตัดแบ่งออกเป็นทีละหัวข้อหลัก ซึ่งจะมีด้วยกัน 11 หัวข้อดังนี้

1. [Code and Project Organization](https://lazts.com/memos/talents/20170520-100-go-mistakes-code-and-project-organization)
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

เลือกเข้าไปอ่านกันได้เลย 😁

## Reference
[100go](https://100go.co/)
