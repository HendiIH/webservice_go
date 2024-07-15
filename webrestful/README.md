# Simple crud prototype

Prototype simple untuk pengaplikasian CRUD, bisa dilakukan melalui curl ataupun
menggunakan template dari GO yang untuk kedepan bisa digabungkan htmx
> Data telah dideklarisan pada ```var albums = []album```

---
Untuk melakukan get cukup melakukan curl biasa seperti
```
$ curl localhost:8080/albums
```
Untuk melakukan pengisian data bisa dilakukan perintah POST seperti contoh
di bawah yang sesuai dari dokumentasi go
```
$ curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```
## TODO
- [x] Protype Rest API untuk mendapatkan data album
- [x] Templating agar dapat melakukan create tanpa menggunakan curl
- [ ] Melakukan fitur RUD pada template
- [ ] Implemnetasi error handling, validasi input, dan juga feedback
- [ ] Implementasi htmx pada template
