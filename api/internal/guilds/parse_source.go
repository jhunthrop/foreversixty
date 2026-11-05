// api/internal/guilds/parse_source.go
package guilds

// The one rule every guild surface's output figure reads (guild round 2, WP-13): the Raids tab's
// top figure for a night (fillTopParse), the Progression tab's best per role (bestByRole) and the
// roster's best (bestParses). Round 1 let the Raids tab rank every kill fight, so a healer's whelp
// pack outranked their Onyxia kill and the same healer read three different numbers on three
// tabs. Each reader now picks from kills of a named boss only, ordered by the row's own role
// metric (HPS for a healer, DPS for everyone else), so for one healer and one kill the three
// agree; they differ only by scope (one night, one boss, the whole log history).

// bossKillMetrics is the where-clause fragment over fight_metrics `m`: a kill of a named boss
// (a whelp pack is no one's figure) that was not removed.
const bossKillMetrics = `m.kill and coalesce(m.encounter_id, 0) > 0 and m.state <> 'removed'`

// roleMetricOrder orders fight_metrics `m` by the figure its own role is judged on.
const roleMetricOrder = `(case when m.role = 'healer' then coalesce(m.metric_hps, 0) else coalesce(m.metric_dps, 0) end)`
