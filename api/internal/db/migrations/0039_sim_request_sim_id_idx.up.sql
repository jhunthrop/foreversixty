-- Look a sims row's in-game request up by its run (docs/contracts/addon-v2.md contract 3;
-- api/internal/sims/request_job.go).
--
-- addon_sim_requests.sim_id is read by run: sims.Store.Mine's FromAddon (exists ... where
--   sim_id = sims.id, once per row of a page) and the inbox's joins from a request to its run.
--   0034 indexed the table by (character_key, created_at) only.
create index if not exists addon_sim_requests_sim_id_idx on addon_sim_requests (sim_id);
