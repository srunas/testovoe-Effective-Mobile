-- name: CreateSubscription :exec
INSERT INTO subscriptions (id, service_name, price, user_id,
                           start_date, end_date)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: GetSubscriptionByID :one
SELECT id, service_name, price, user_id, start_date, end_date,
       created_at, updated_at
FROM subscriptions
WHERE id = $1;

-- name: UpdateSubscription :exec
UPDATE subscriptions
SET service_name = $2,
    price        = $3,
    start_date   = $4,
    end_date     = $5,
    updated_at   = NOW()
WHERE id = $1;

-- name: DeleteSubscription :exec
DELETE FROM subscriptions
WHERE id = $1;

-- name: ListSubscriptions :many
SELECT id, service_name, price, user_id, start_date, end_date,
       created_at, updated_at
FROM subscriptions
WHERE ($1::text = '' OR user_id = $1)
ORDER BY created_at DESC;

-- name: CalculateTotal :one
SELECT COALESCE(SUM(price), 0)::bigint AS total
FROM subscriptions
WHERE ($1::text = '' OR user_id = $1)
  AND ($2::text = '' OR service_name = $2)
  AND start_date >= $3
  AND start_date <= $4;