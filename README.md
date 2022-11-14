# golang-final-project2-team2

Link Deploy API : https://golang-final-project2-team2-production.up.railway.app/

Repository Team 2 Untuk Final Project 2 Golang (Kampus Merdeka Hacktiv8)

Anggota Kelompok :

- JOVIN LIDAN (GLNG-KS04-023)
  Mengerjakan :
  - Setup Docker dan Init Project
  - Semua API User dan Postman user
  - Semua API Photo dan Postman photo
  - Setup deployment menggunakan railway
  - Helper : GenerateToken, VerifyToken, ValidateRequest, ComparePass, HashPass
  - Helper_test : TestSuccessGenerateToken, TestFailedGenerateToken, TestSuccessComparePass, TestFailedComparePass, TestSuccessHashPass, TestFailedHashPass
- GUSTIO NUSAMBA (GLNG-KS04-025)
  Mengerjakan :
  - Semua API Comment
  - Semua API Social Media

## Cara Install

1. Buka dan jalankan aplikasi docker.
2. Jalankan command `docker compose up --build` untuk menjalankan database postgres di dalam docker container , go dan air auto reload. Tunggu agar docker sudah berjalan dengan baik.
3. Setelah docker container semuanya berjalan dengan baik. Maka port default yang akan dibuka adalah `8080`

_Note : Memerlukan docker terinstall didalam perangkat anda_

_Nama File Postman : `MyGram.postman_collection.json`_

## List Route
### Users
- **`POST`- Users Register `api/users/register`**, Digunakan untuk membuat user baru.
- **`POST`- Users Login `api/users/login`**, Digunakan untuk melakukan login atau autentikasi user.
- **`PUT`- Users Update `api/users/:userId`**, Digunakan untuk mengubah data user berdasarkan idnya.
- **`DELETE`- Users Delete `api/users`**, Digunakan untuk menghapus data user.

### Photos
- **`GET`- Photos Index `api/photos`**, Digunakan untuk mengambil seluruh data photos dari database.
- **`POST`- Photos Store `api/photos`**, Digunakan untuk membuat photo baru.
- **`PUT`- Photos Update `api/photos/:photoId`**, Digunakan untuk mengubah data photo berdasarkan idnya.
- **`DELETE`- Photos Delete `api/photos/:photoId`**, Digunakan untuk menghapus data photo berdasarkan idnya.
