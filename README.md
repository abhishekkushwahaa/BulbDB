# BulbDB

===
<a target="_blank" href="https://discord.gg/abhiimaurya"><img src="https://dcbadge.limes.pink/api/server/abhiimaurya?style=flat" alt="discord community" /></a>
<a href="https://x.com/AbhishekKushwaa">![X](https://img.shields.io/badge/X-000000?style=flat-square&logo=x)</a>
<a href="https://linkedin.com/in/abhishekkushwahaa">![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=flat-square&logo=linkedin)</a>
<a href="https://abhishekkushwaha.tech">![Website](https://img.shields.io/badge/Website-FF4500?style=flat-square)</a>
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

BulbDB is a lightweight, relational database built from scratch in Go, featuring a B+Tree storage engine and SQL-like query execution.

## Features

- **B+Tree-based storage** for efficient indexing and retrieval
- **SQL-like query support** for structured data management
- **Custom storage engine** optimized for performance
- **ACID compliance** for reliability and data integrity
- **Written in Go** with a focus on simplicity and efficiency

## Installation

```sh
git clone https://github.com/abhishekkushwahaa/BulbDB.git
cd BulbDB
make build
```

## Usage

Start the database server:

```sh
./bulbdb
```

Execute queries using the interactive shell:

```sql
CREATE TABLE users (id INT PRIMARY KEY, name TEXT, email TEXT);
INSERT INTO users VALUES (1, 'Test', 'test@gmail.com');
SELECT * FROM users;
```

## Roadmap

- [ ] Implement SQL Parser
- [ ] Enhance Query Optimizer
- [ ] Add Transaction Support
- [ ] Improve Storage Engine

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests.

## License

MIT License. See [LICENSE](LICENSE) for details.
