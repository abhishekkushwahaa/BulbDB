# BulbDB

<a target="_blank" href="https://discord.gg/m9wE84YGJH"><img src="https://dcbadge.limes.pink/api/server/m9wE84YGJH?style=flat" alt="Join our Discord community" /></a>
<a href="https://linkedin.com/in/abhishekkushwahaa">![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=flat-square&logo=linkedin)</a>
<a href="https://x.com/AbhishekKushwaa">![X](https://img.shields.io/badge/X-000000?style=flat-square&logo=x)</a>
<a href="https://abhishekkushwaha.tech">![Website](https://img.shields.io/badge/Website-FF4500?style=flat-square)</a>
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## Overview

**BulbDB** is a lightweight relational database built from the ground up in Go. It features a **B+Tree storage engine** and an **SQL-like query execution layer**, making it ideal for structured data management with speed and efficiency.

## Key Features

- 🚀 **Optimized B+Tree-based storage** for fast indexing and data retrieval.
- 🛠 **SQL-like query execution** for intuitive and structured interactions.
- 🔒 **ACID-compliant transactions** ensuring data consistency and reliability.
- ⚡ **Custom-built storage engine** tailored for high performance.
- 📌 **Written in Go** with a focus on efficiency, scalability, and maintainability.

## Installation

To get started with BulbDB, clone the repository and build it:

```sh
git clone https://github.com/abhishekkushwahaa/BulbDB.git
cd BulbDB
make build
```

## Usage

Start the BulbDB server:

```sh
./bulbdb
```

Run queries using the interactive shell:

```sql
CREATE TABLE users (id INT PRIMARY KEY, name TEXT, email TEXT);
INSERT INTO users VALUES (1, 'Test', 'test@gmail.com');
SELECT * FROM users;
```

## Roadmap

BulbDB is an evolving project. Here are key milestones we are working on:

- ✅ Initial implementation of B+Tree storage
- 🔜 SQL Parser for advanced queries
- 🔜 Query Optimizer enhancements
- 🔜 ACID Transaction Support
- 🔜 Performance tuning and scalability improvements

## Contributing

We welcome contributions from the community! Feel free to **submit issues, feature requests, or pull requests** to help improve BulbDB.

### How to Contribute:

1. Fork the repository and create a new branch.
2. Implement your changes with clear commit messages.
3. Submit a pull request with a detailed description.

## License

BulbDB is released under the **MIT License**. See [LICENSE](LICENSE) for more details.

---

🚀 **Stay connected!** Join our community on **[Discord](https://discord.gg/m9wE84YGJH)** and follow updates on **[LinkedIn](https://linkedin.com/in/abhishekkushwahaa)** & **[X](https://x.com/AbhishekKushwaa)**.
