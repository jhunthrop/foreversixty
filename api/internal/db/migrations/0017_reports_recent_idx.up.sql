-- GET /v1/reports/recent lists the newest complete public reports,
-- keyset-paged on (created_at, id). Neither existing reports index
-- covers this: reports_owner_idx and reports_guild_idx are keyed on
-- owner_id and guild_id, which this query does not filter by. Without
-- an index the query is a sequential scan of the whole table plus a
-- sort on every call to an unauthenticated, high-traffic route.
--
-- A partial index scoped to the rows this route ever reads, already
-- ordered the way the query reads them, both answers the filter and
-- serves the ORDER BY with no extra sort step - the same pattern
-- fights_encounter_idx uses above for a comparably shaped read.
create index if not exists reports_recent_idx on reports (created_at desc, id desc)
  where visibility = 'public' and status = 'complete';
