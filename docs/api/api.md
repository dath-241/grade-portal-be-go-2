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
    "name": {
        "LastName": string,
        "MFName": string
    },
    "email": string
}
```
- __Response__ :
```bash
{
    "data": {
        "_id": string,
        "name": {
            "LastName": string,
            "MFName": string
        },
        "email": string, // @hcmut.edu.vn
        "LastLogin": Date,
        "role": "Admin",
        "updated_at": Date,
        "created_at": Date,
        "admin_info": {
            "admin_id": string
        }
    } | null,
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