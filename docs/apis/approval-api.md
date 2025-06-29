# 🧾 Approval Service API Documentation

> Dokumentasi API untuk service approval internal. Mendukung versioning per endpoint (`v1`, `v2`, dst).

---

## 🌐 Overview

- **Service Name**: Approval Service
- **Owner**: @taufik
- **Base URLs**:
  - 🧪 Local: `http://localhost:4317`
  - 🚀 Staging: `https://staging.api.company.com/approval`
  - ✅ Production: `https://api.company.com/approval`

---

## 📁 API Versions

---

## 📄 v1

### 📑 Endpoints

| Method | Endpoint                       | Description                  | Auth Required |
| ------ | ------------------------------ | ---------------------------- | ------------- |
| GET    | `/v1/approval-templates`     | List approval templates      | ✅            |
| POST   | `/v1/approval-templates`     | Create approval template     | ✅            |
| GET    | `/v1/approval-templates/:id` | Get approval template detail | ✅            |
| PUT    | `/v1/approval-templates/:id` | Update approval template     | ✅            |

---

### 🧾 Request/Response Examples

**POST /v1/approval-templates**

**Request:**

```json
{
  "slug": "leave-approval",
  "display_name": "Leave Request",
  "resource_type": "leave"
}
```
