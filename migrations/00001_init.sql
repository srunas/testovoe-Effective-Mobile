-- +goose Up
CREATE TABLE subscriptions (
                               id           TEXT        PRIMARY KEY,
                               service_name TEXT        NOT NULL,
                               price        BIGINT      NOT NULL,
                               user_id      TEXT        NOT NULL,
                               start_date   DATE        NOT NULL,
                               end_date     DATE,
                               created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                               updated_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_subscriptions_user_id      ON subscriptions
    (user_id);
CREATE INDEX idx_subscriptions_service_name ON subscriptions
    (service_name);

-- +goose Down
DROP TABLE subscriptions;