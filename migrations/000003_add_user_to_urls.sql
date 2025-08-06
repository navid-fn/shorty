-- +goose Up
-- +goose StatementBegin
ALTER TABLE urls
ADD COLUMN user_id BIGINT REFERENCES users(id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP COLUMN user_id FROM urls;
-- +goose StatementEnd
