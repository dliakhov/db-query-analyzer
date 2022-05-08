INSERT INTO query_analyzer.database_query_info (created_at, updated_at, deleted_at, query, execution_time_ms)
SELECT (now() at time zone 'utc'), (now() at time zone 'utc'), null, 'SELECT * FROM users', round(random()*10000)
FROM generate_series(1,20);

INSERT INTO query_analyzer.database_query_info (created_at, updated_at, deleted_at, query, execution_time_ms)
SELECT (now() at time zone 'utc'), (now() at time zone 'utc'), null, 'UPDATE users SET age = ?', round(random()*10000)
FROM generate_series(1,20);

INSERT INTO query_analyzer.database_query_info (created_at, updated_at, deleted_at, query, execution_time_ms)
SELECT (now() at time zone 'utc'), (now() at time zone 'utc'), null, 'DELETE FROM users WHERE id = ?', round(random()*10000)
FROM generate_series(1,20);

INSERT INTO query_analyzer.database_query_info (created_at, updated_at, deleted_at, query, execution_time_ms)
SELECT (now() at time zone 'utc'), (now() at time zone 'utc'), null, 'INSERT INTO users (id, name, age) VALUES (?, ?, ?)', round(random()*10000)
FROM generate_series(1,20);
