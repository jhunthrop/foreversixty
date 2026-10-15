-- The export an in-game sim request was made on (docs/contracts/addon-v2.md contract 3;
-- api/internal/sims/request_job.go).
--
-- addon_sim_requests.export_seen_at: the character's addon_exports.updated_at when the request
--   was recorded, the exact value the inbox ack carries for that export, so the weights
--   message of the request's run can say which copy it ran on. NULL on every row recorded
--   before this migration, and for a character with no stored export.
alter table addon_sim_requests add column if not exists export_seen_at timestamptz;
