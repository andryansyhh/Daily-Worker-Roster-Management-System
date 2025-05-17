CREATE TABLE IF NOT EXISTS workers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS shifts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date TEXT NOT NULL,
    start_time TEXT NOT NULL,
    end_time TEXT NOT NULL,
    role TEXT NOT NULL,
    location TEXT
);

CREATE TABLE IF NOT EXISTS shift_requests (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    shift_id INTEGER NOT NULL,
    worker_id INTEGER NOT NULL,
    status TEXT NOT NULL CHECK (status IN ('pending', 'approved', 'rejected')),
    created_at TEXT NOT NULL,
    FOREIGN KEY (shift_id) REFERENCES shifts(id),
    FOREIGN KEY (worker_id) REFERENCES workers(id)
);

CREATE TABLE IF NOT EXISTS assignments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    shift_id INTEGER NOT NULL,
    worker_id INTEGER NOT NULL,
    FOREIGN KEY (shift_id) REFERENCES shifts(id),
    FOREIGN KEY (worker_id) REFERENCES workers(id)
);