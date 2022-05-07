CREATE SCHEMA query_analyzer;
CREATE TABLE query_analyzer.database_query_info
(
    id                SERIAL PRIMARY KEY NOT NULL,
    created_at        TIMESTAMP       NOT NULL,
    updated_at        TIMESTAMP,
    delete_at         TIMESTAMP,
    query             TEXT,
    execution_time_ms INTEGER
);

CREATE INDEX idx_database_query_info_query ON query_analyzer.database_query_info (query);
CREATE INDEX idx_database_query_info_delete ON query_analyzer.database_query_info (delete_at);

-- insert test data
INSERT INTO query_analyzer.database_query_info (created_at, updated_at, delete_at, query, execution_time_ms)
SELECT (now() at time zone 'utc'), (now() at time zone 'utc'), null, 'SELECT * FROM users', round(random()*10000)
FROM generate_series(1,20);

INSERT INTO query_analyzer.database_query_info (created_at, updated_at, delete_at, query, execution_time_ms)
SELECT (now() at time zone 'utc'), (now() at time zone 'utc'), null, 'UPDATE users SET age = ?', round(random()*10000)
FROM generate_series(1,20);

INSERT INTO query_analyzer.database_query_info (created_at, updated_at, delete_at, query, execution_time_ms)
SELECT (now() at time zone 'utc'), (now() at time zone 'utc'), null, 'DELETE FROM users WHERE id = ?', round(random()*10000)
FROM generate_series(1,20);

INSERT INTO query_analyzer.database_query_info (created_at, updated_at, delete_at, query, execution_time_ms)
SELECT (now() at time zone 'utc'), (now() at time zone 'utc'), null, 'INSERT INTO users (id, name, age) VALUES (?, ?, ?)', round(random()*10000)
FROM generate_series(1,20);
