# 🔐 AuthSphere Microservice

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge)](LICENSE)
[![Build Status](https://img.shields.io/badge/Build-Passing-brightgreen.svg?style=for-the-badge)]()
[![Docker Ready](https://img.shields.io/badge/Docker-Ready-2496ED?style=for-the-badge&logo=docker)](https://www.docker.com/)

A high-performance, production-ready **Authentication and Authorization Service** built with Go, PostgreSQL. AuthSphere provides robust JWT-based authentication featuring **stateless short-lived Access Tokens**, **secure HttpOnly cookie-based Refresh Tokens**, blacklisting/revocation support, and role-based protected routes.

---

## 🌟 Key Features

- **User Onboarding & Management**: User registration with hashed password storage (Bcrypt), login, and logout.
- **Dual-Token Architecture**:
  - **Access Token**: Short-lived JWT (e.g., 15 mins) passed via `Authorization: Bearer <token>` header.
  - **Refresh Token**: Long-lived token (e.g., 7 days) stored in a secure `HttpOnly`, `SameSite=Strict`, `Secure` cookie.
- **Token Invalidation & Logout**: Revokes refresh tokens and blacklists active sessions using Redis.
- **Protected Routes**: Middleware enforcement for authenticated users and RBAC (Role-Based Access Control).
- **Security Best Practices**:
  - OWASP compliant password hashing.
  - Rate limiting on public auth endpoints.
  - CORS and CSRF protections.
- **Developer-Friendly**: Hot-reloading setup via [Air](https://github.com/air-verse/air), clean architecture, and full Docker Compose support.

---

## 🏗️ Architecture Overview
