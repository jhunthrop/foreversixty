-- The addon's inbox v2 (docs/contracts/addon-v2.md contract 3; api/internal/addon/messages.go).
--
-- rating_scores: the per-fight facts the last_raid message shows, worked out by the rater
--   from the card it scored and the fight summary (api/internal/rating/last_raid.go), so the
--   inbox reads rows and never a summary. NULL on a row rated before this migration.
--   deaths             how many times the player died
--   biggest_loss       the plain phrase for the component that cost the most points
--   loss_seconds       downtime seconds when that is the loss, else seconds dead, else 0
--   consumables_missed the plain category names ("flask", "food") the player pulled without
alter table rating_scores
  add column if not exists deaths integer,
  add column if not exists biggest_loss text,
  add column if not exists loss_seconds integer,
  add column if not exists consumables_missed text[];

-- addon_sim_requests: what became of each in-game sim request, for the inbox ack's
-- sim_request_state. 0034 stored only requests that started a run; now a refused one is
-- recorded too, with sim_id NULL: 'queued' (started a run), 'needs_premium', 'duplicate'
-- (another request started a run within the hour) or 'invalid'. Every row 0034 stored
-- started a run.
alter table addon_sim_requests add column if not exists state text;
update addon_sim_requests set state = 'queued' where state is null;
