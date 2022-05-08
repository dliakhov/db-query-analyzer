CREATE TABLE query_analyzer.database_query_info
(
    id                SERIAL PRIMARY KEY NOT NULL,
    created_at        TIMESTAMP       NOT NULL,
    updated_at        TIMESTAMP,
    deleted_at         TIMESTAMP,
    query             TEXT,
    execution_time_ms INTEGER
);

CREATE INDEX idx_database_query_info_query ON query_analyzer.database_query_info (query);
CREATE INDEX idx_database_query_info_deleted_at ON query_analyzer.database_query_info (deleted_at);
