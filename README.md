# 🔐 **Security & Authorization Domain**

## 📖 Description
The **Security & Authorization** domain is responsible for managing authentication, authorization, and data encryption within the hospital system. Each functionality is implemented as an independent microservice to ensure **security, scalability, and modularity**.

---

## 🔹 Microservices

### 🛡️ **2. Password Encryption and Verification**
- **📌 Description:** Encrypts passwords before storing them in the database and verifies if a provided password matches the stored encrypted password.
- **🔹 Method:** `POST`
- **📥 Inputs:** Plain text password or encrypted password
- **📤 Outputs:** Encrypted password 🔐 or verification result (match ✅ or no match ❌)

---

## 🛠️ **Technologies Used**
- **⚙️ Backend:** Go, bycrypt 💻
- **🗄️ Database:** PostgreSQL 🐘, MySQL 🐬

---

## 🔗 **Integrations**
- **🏥 Patient Management Domain:** Encryption is necessary to keep patient passwords secure.
- **🩺 Doctor Management Domain:** Encryption is necessary to keep doctor passwords secure.
- **🧑 Admin Management Domain:** Encryption is necessary to keep admin passwords secure.
