# Sephera's Biển - YAGPDB's Custom Command
Một số lệnh tùy chỉnh Kiru tạo ra bằng ngôn ngữ Go / Golang để phục vụ con dân Server :3 
> Công cụ: YAGPDB - Custom Command

> Note: YAGPDB đã có sẵn Dashboard hỗ trợ việc nhập liệu (Input), trong đây sẽ chỉ có phần Code về (lưu trữ Data riêng trong Server / xử lý dữ liệu / xuất dữ liệu)

I. Bộ Randomizer tùy chỉnh (Chủ đề: Liên Quân Mobile)
  1. Random đơn lẻ (`random_hero_lqm.go`)
> Random 1 tướng - 1 vị trí - 1 bổ trợ
> 
> Quy định: Vị trí Rừng => Bổ trợ Trừng Trị

![image](https://github.com/user-attachments/assets/71b88a97-6b9e-445a-88bf-bea8c884eef7)

  2. Random kép - đôi (+ gán) (`random_solo_lqm.go`)
> Random 1 chất tướng (Sát thủ / Đấu sĩ / Pháp sư / Trợ thủ / Xạ thủ/ Đỡ đòn)
>
> Thực hiện 2x (Random tướng thuộc chất tướng đã chọn - gán cho người dùng tương ứng)

![image](https://github.com/user-attachments/assets/d8957c33-b0fb-4107-b8d0-9e857c54b965)


II. Bộ "Gacha" (Advanced Randomizer) (Chủ đề: Sephera - Liên Quân Mobile) `[trích]`

> `gacha_uprate_pity_x10.go` 

- Giới thiệu

![image](https://github.com/user-attachments/assets/0807bdc7-ea94-46c7-8fba-567192567fb3)

- Cơ chế
> Randomizer tùy biến sâu
> 
> Lặp 10 lần (cố định)

> - Bộ randomizer chính:
>
> ```Banner S (0.5% / 1% / 2% / 96.5%)
> SSR (0.5%) = List 10 roles
> SR (1%) = Role1
> UC (2%) = Role2
> C (96.5%) = Fail (Nothing)```

> - Bộ randomizer phụ 1 (Fail)
> *Chỉ random các câu "xịt" kèm emoji khác nhau, không có gì đặc biệt*

> - Bộ randomizer phụ 2 (SSR = List 10 roles)
>
> Chọn 1 Role trong List trên
>
> **Cơ chế UPRATE**: 1 Role chọn trước gắn mác "UPRATE" sẽ chiếm tỉ lệ nhiều hon so với các Role còn lại
>
> *(Bộ randomizer phụ 3 quyết định chọn Role UPRATE hoặc không)*

> **Cơ chế PITY**: Dựa vào dữ liệu lưu trong mục DATABASE của YAGPDB, ảnh hưởng kết quả cuối cùng
>
> -> *"Đạt PITY": Người dùng được đảm bảo trúng **SSR** (mặc dù bộ randomizer chính "roll toàn xịt")*

- Phụ
> Cách hiển thị kết quả:
>
> - Chào mừng người dùng
>
> - Kiểm tra PITY, thông báo nếu đạt
>
> - Bắt đầu cơ chế Randomizer chính (lặp 10x, nghỉ 2s giữa các lần) / Randomizer phụ / cơ chế Uprate / ... => Thực hiện **Edit** tin nhắn ban đầu sau mỗi vòng lặp, không gửi tin nhắn mới  
>
> - Thông báo kết quả cuối cùng

![image](https://github.com/user-attachments/assets/4fb3aeb5-1730-48f1-9539-c119de07d4da)

![image](https://github.com/user-attachments/assets/f4643dfc-8cd5-4981-806c-cd949787d73c)

![image](https://github.com/user-attachments/assets/518a1e84-bf22-4847-a426-b9370daba3ab)


