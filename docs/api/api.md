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

### CHỨC NĂNG: TẠO `ADMIN` -> `Phát`
- __API__: `admin/admin/create`
- __Mô tả__: Tạo Admin mới
- __Body__ :
```bash
{
    "email":    string,    
    "password": string, 
    "name":     string,     
    "faculty":  string,  
    "role":     string     
}
```
- __Response__ :
```bash
{
    "data": null | object,
    "message": string,
    "success": boolean
}
```
---
### `GET`, `PUT`, `DELETE`
---
### CHỨC NĂNG : TẠO `TEACHER` -> `HÂN`
...

---
admin/log
https://localost:8080/admin/login