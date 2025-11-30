# 🚀 Rancangan Proyek Final: Quick Typer

Proyek ini mengintegrasikan **Pemrograman Mobile (Android Java)** sebagai klien, **Game Programming** sebagai logika kecepatan dan akurasi, dan **Cloud Computing** sebagai arsitektur *Backend* yang mengelola konten (*Stages*), *Leaderboard*, dan Autentikasi.

## I. 🎯 Mekanisme Game & Scoring

### A. Alur Permainan

1.  **Login/Register:** Pemain login menggunakan kredensial (Backend menggunakan Token API berbasis tabel).
2.  **Stage Selection:** Pemain mengambil daftar *Stage* dari Cloud (`GET /api/stages`) dan memilih salah satunya.
3.  **Game Start:** Aplikasi mengambil detail *Stage* dan semua Frasa yang terurut (`GET /api/stage/{id}`).
4.  **Game Loop:** Pemain mengetik semua **1-5 Frasa** secara berurutan. **Timer Kumulatif** berjalan dari ketikan pertama hingga terakhir.
5.  **Data Recording:** Aplikasi Java mencatat **Total Waktu (ms)** dan **Total Kesalahan (Error)** yang terakumulasi.
6.  **Score Submission:** Setelah selesai, aplikasi mengirim data mentah ke Cloud (`POST /api/score/submit`).

### B. Kalkulasi Skor (Server-Side)

Skor dihitung berdasarkan **Skor per Detik** dikurangi **Penalti Kesalahan**. Logika ini memastikan kecepatan dan akurasi sama pentingnya.

$$\mathbf{Final Score} = \left( \frac{\sum (\text{Panjang Frasa}_i \times \text{Multiplier}_i)}{\text{Total Time Taken in Seconds}} \right) - (\text{Total Error Penalty})$$

| Komponen Skor | Sumber Data | Keterangan |
| :--- | :--- | :--- |
| $\text{Panjang Frasa}_i$ | Tabel `Phrases` (`text` length) | Jumlah karakter dalam frasa $i$. |
| $\text{Multiplier}_i$ | Tabel `Phrases` (`base_multiplier`) | Koefisien skor per karakter (ditetapkan Admin). |
| $\text{Total Time}$ | Data kiriman klien | $\text{Total Time Taken (ms)} / 1000$. |
| $\text{Error Penalty}$ | Data kiriman klien | $\text{Total Errors} \times \text{Constant Value}$ (misalnya, 50). |

---

## II. ☁️ Database Schema Lengkap

Skema ini menggunakan **Kunci Unik Komposit** pada tabel `Scores` untuk menerapkan *Best Score Only* (UPSERT), dan tabel terpisah untuk manajemen Token API.

### 1. 👥 Users

| Field Name | Tipe Data | Keterangan |
| :--- | :--- | :--- |
| `user_id` | PK, Unique ID | ID Utama Pengguna. |
| `username` | String (Unique) | Nama pengguna. |
| `password_hash` | String | Hash kata sandi. |
| `role` | String | **'user' atau 'admin'**. |

### 2. 🔑 PersonalAccessTokens (Untuk Login)

| Field Name | Tipe Data | Keterangan |
| :--- | :--- | :--- |
| `id` | PK, Unique ID | |
| `user_id` | FK | Pemilik token. |
| **`token`** | String (Indexed) | **Hash Token** yang dikirimkan klien. |
| **`expires_at`** | Timestamp | Kapan token kedaluwarsa. |
| **`revoked_at`** | Timestamp | Status pencabutan token (digunakan untuk logout). |

### 3. 🗺️ Stages (Konten Induk)

| Field Name | Tipe Data | Keterangan |
| :--- | :--- | :--- |
| `stage_id` | PK, Unique ID | ID unik Stage. |
| `name` | String | Nama Stage (e.g., "Intro to Java"). |
| `theme` | String | Tema Stage. |
| `difficulty` | String | Kesulitan (e.g., "Easy", "Hard"). |
| `is_active` | Boolean | |

### 4. 📝 Phrases (Konten Detail)

| Field Name | Tipe Data | Keterangan |
| :--- | :--- | :--- |
| `phrase_id` | PK, Unique ID | |
| **`stage_id`** | FK | **Mengikat Frasa ke Stage.** |
| `text` | String | Teks frasa yang harus diketik. |
| **`sequence_number`**| Integer | Urutan frasa dalam Stage (1, 2, 3...). |
| `base_multiplier` | Float | Koefisien skor dasar (dibuat Admin). |

### 5. 🏆 Scores (Best Score Only - UPSERT)

| Field Name | Tipe Data | Keterangan |
| :--- | :--- | :--- |
| **`user_id`** | FK | **Bagian dari Kunci Unik Komposit.** |
| **`stage_id`** | FK | **Bagian dari Kunci Unik Komposit.** |
| `final_score` | Integer/Float | Skor terbaik yang dicapai. |
| `total_time_ms` | Integer | Waktu terbaik yang tercatat. |
| `total_errors` | Integer | Kesalahan terbaik yang tercatat. |
| `completed_at` | Timestamp | Waktu pencatatan skor terbaik. |
| **Constraint** | **Unique Key** | **(`user_id`, `stage_id`)** |

---

## III. 🔑 Kontrak API Lengkap

### 1. Autentikasi (Token API)

| Endpoint | Metode | Input (JSON) | Respons (JSON) |
| :--- | :--- | :--- | :--- |
| `/api/auth/login` | `POST` | `{username, password}` | `{user_id, access_token, token_expires_at}` |
| `/api/auth/register` | `POST` | `{username, password, email}` | `{user_id, access_token, token_expires_at}` |
| `/api/auth/profile` | `GET` | Header `Authorization bearer_{access_token}` | `{user_id, role, token_expires_at}` |


### 2. Content Management System (CMS)

*Semua endpoint CMS memerlukan **Token Admin** dan pengecekan masa berlaku Token di middleware.*

| Endpoint | Metode | Input (JSON) | Keterangan |
| :--- | :--- | :--- | :--- |
| `/admin/stage` | `POST` | `{name, theme, difficulty, is_active}` | Membuat record baru di tabel `Stages`. |
| `/admin/phrase` | `POST` | `{stage_id, text, sequence_number, base_multiplier}` | Menambahkan frasa baru (harus terikat Stage). |

### 3. Game & Leaderboard

*Semua endpoint ini memerlukan Token User dan pengecekan masa berlaku Token.*

| Endpoint | Metode | Input (Query Params) | Respons (JSON) |
| :--- | :--- | :--- | :--- |
| `/api/stages` | `GET` | N/A | `[{stage_id, name, difficulty}, ...]` |
| `/api/stage/{id}` | `GET` | N/A | `{stage_id, name, phrases: [{phrase_id, text, sequence_number, multiplier}, ...]}` |
| `/api/score/submit` | `POST` | `{stage_id, total_time_ms, total_errors}` | `{status: "UPSERTED" / "IGNORED", final_score}` |
| `/api/leaderboard` | `GET` | `?stage_id={id}` | `[{username, final_score, total_time_ms}, ...]` (Dibatasi 10 atau 20) |