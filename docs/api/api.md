# DANH SÁCH API
## ADMIN
### `POST`
### CHỨC NĂNG : `Login` , `Logout` -> `Minh Tâm`
- __LOGIN__
- __API__ : `admin/login`
    - __Yêu Cầu__ : lưu lại cookie và session trên máy người dùng.
- __JSON__ :
```bash
{
    "idToken" : idToken
}
```
- __LOGOUT__ 
- __API__ : `admin/logout`
    - __Yêu cầu__ : xóa cookie và session trên máy người dùng.
### CHỨC NĂNG : Tạo `Student`
- __API__ : `admin/student/create`
- __Mô tả__ : Tạo mới một sinh viên.
- __Body__ :
```bash
{
    "lastName": string,
    "mfName": string,
    "email": string, // abc@hcmut.edu.vn
    "studentID": string,
    "classID": []
}
```
- __Response__ :
```bash
{
    "data": object,
    "message": string,
    "success": boolean
}
```
---
### `GET`, `PUT`, `DELETE`
---
### CHỨC NĂNG : TẠO `TEACHER` -> `HÂN`
...

admin/log
https://localost:8080/admin/login