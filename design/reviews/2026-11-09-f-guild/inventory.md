# Guild data inventory (batch F, lane guild, 2026-11-09)

What the guild API and DB already hold, and what the guild pages show of it. Migrations: `api/internal/db/migrations/`; handlers: `api/internal/guilds/` unless named. Web types are hand-written (`web/src/lib/guild/*.ts`); `api/openapi.yaml` is the declared spec (it generates only the mobile clients).

Headline: the roster, readiness, raid view, report and loot reads already return most of what an officer needs. The pages render about a fifth of it, and the parity fixtures sent nulls for most of the rest (rating, best_parse, professions, addon, attendance_30), so nobody saw it.

## 1. Returned by the API but not shown (or barely) on the guild pages

| Read | Field(s) | Shown today |
|---|---|---|
| `GET /guilds/{id}/home` roster[] | `rating{overall,output,survival,mechanics,utility,preparation,activity}`, `best_parse{metric,value,percentile,encounter,report_id,fight_index}`, `attendance_30`, `professions[]`, `addon{level,spec,gear_gap,build,seen_at}`, `logged_at`, `last_report_at`, `is_main`/`account_key`, `sim_export`, `role`, `level` | item level and raids only (the rest is coded but hidden behind null fixtures) |
| home `summary` | `below_rating_floor`, `tier`, `named_encounters_down`, `pulls_this_tier`, `updated_at` | no |
| home `standing` | `rank_by_item_level`, `same_spec_count`, `item_level`, `needs_before_next_raid` | partly (checklist) |
| home `reports[]` | `fight_count`, `kill_count`, `duration_ms` | one line |
| home `addon_snapshot` | `received_at`, `members` | no |
| `GET .../readiness` rows[] | `gear_gap{upgrades,gain_dps,unit}`, `enchants.missing_slots`, `consumables.state`, `talent_points_unspent`, `item_level_delta` vs `median_item_level`, `logged_at`, `failing`, officer: `attuned`, `resistances{fire,nature,frost,shadow,arcane}`, `flask_food` | only the failing rows' reasons; the ready raiders as a name list |
| `GET .../attendance/recent` | `raids[]` (last nights), rows[] `cells[]` per night, `slipping` | the four-night grid, slipping rows only |
| `GET .../raids/{raid}` (planned) | signups[] `note`, `late_minutes`, `also[]`, `spec_key`; `not_signed_raiders[]{role,last_week,discord_linked}`; `counts.classes`, `late_groups`; `cover[]` (buff and debuff providers in / late / bench); `comp_target` | counts, rows without figures; cover only on the overview |
| `GET .../raids/{raid}` (done) report | `parses[]` every raider per kill with `value` and `rating` (officers); `kill_order[]{kill_ms,wipes_before,deaths}`; `deaths[]` by cause; `pulls[].death_causes`; `loot[]`; `moves{previous,rows}` (the viewer's changes vs the night before); `night{signed,came,late,not_in_log,joined,excused,kills,wipes,deaths,duration_ms}`; attendance rows `pulls_present/pulls_total`, `first_seen_offset_s` | top three parses, deaths behind a pull picker |
| `GET .../raids` rows[] | `top_dps`, `top_hps`, `raiders`, `deaths`, `wipe_count`, `duration_ms`, `present[]`, `fights[]{deaths,players,duration_ms}` | kills and wipes |
| `GET .../progression` encounters[] | `pulls_by_night[]`, `best_kill_ms`, `best_by_role`, `deaths_per_pull`, `first_kill_at`; `unnamed[]` | one bar per boss |
| `GET .../loot/nights` | every night's `drops[]` and `decisions[]{name,class,ms_os,reason,votes_yes,votes_no,votes_officers,awarded_by,at}` (the whole loot history) | the selected night only |
| `GET .../loot` items[] candidates[] | `ilvl_delta`, `attendance`, `items_won`, `already_equivalent`, `fills_empty_slot`, `replaces_item` (icon), `tier`; items `copies`, `eligible[]`; encounters `kill_nights`, `on_farm` | gain and a text line |
| `GET .../loot/drops/{item}/candidates` | `bis_for_them`, `ms_os`, `votes{yes,no,voters,mine}` | yes |
| `GET .../members/{account}/profile` | `attendance.nights[]`, `loot.items[]`, readiness row, notes | partly |
| `GET .../applications` sheet | `score`, `sim{value,unit}`, `attunement[]`, `enchants`, `bis_gap{pieces,list}`, `history[]` | partly |
| `GET .../raid-badge` | `sims_rank{position,of}`, `last_raid_loot` | no web use |
| `GET /v1/me/raid-nights/{raid}` | `loot[]{gain,unit,replaces}`, `bests[]{dps,percentile,kill_ms}`, `new_bests` | app only |

## 2. In the DB, returned by no endpoint

| Table | Columns |
|---|---|
| `guild_members` | **joined_at** (tenure; used only to blank attendance cells), verified_at |
| `guild_characters` | rank_index (in-game rank number), source, verified_by/at |
| `guild_raid_signups` | **updated_at** (when they answered), signed_by / acted_by (an officer moved them) |
| `guild_raid_comp_targets` | updated_by/at |
| `guild_raid_pings` | pinged_by, mentioned count, per-user mention log |
| `raid_assignments` | set_by/at |
| `loot_votes` | voted_at; per-voter split outside the candidates view |
| `loot_interest` | set_at, set_by |
| `raid_drops` | **loot_method, dropped_at, reported_by/character**, corpse_key |
| `guild_trials` | started_at (billing only) |
| `character_syncs` | every sync attempt and its outcome |
| `rating_scores` | **deaths, biggest_loss, loss_seconds, consumables_missed** (per fight) |
| `fight_metrics` | **damage_taken, active_ms**, buff_count, trinkets, talent_split |
| `percentile_digests` | a per-parse percentile for the report's parses (only the roster's best parse carries one) |
| `character_sims` | the real sim DPS per character (no roster or readiness row carries it; the readiness gain is the BiS-file gain) |
| `rating_scores` (Forever Score) | score and tier per character (applicant sheet only) |
| `activity_events.guild_id` | a guild activity feed (no guild-scoped read) |

Computed but sent only to Discord: guild-wide lowest attendance (`discord_reads_lowest.go`), a character's attendance lines, latest logged night.

## 3. Per-member data a guild view joins

| Data | Source | In a guild read? |
|---|---|---|
| BiS-file gear gain, upgrade count | readiness `gear_gap`, loot candidates, checklist | yes |
| Real sim DPS | `character_sims`, home `guild_sims` ladder | no (applicant sheet only) |
| BiS progress per slot | rankings Character `gear[]` | no (applicant `bis_gap`) |
| Forever rating per fight | `rating_scores.overall` | roster `rating`, report `parses[].rating` |
| Best parse and percentile | `fight_metrics` + digests | roster `best_parse` |
| Attunement, resistances, flask/food, professions | the addon export | readiness (officer), roster professions |
| Last seen | addon `seen_at`, `logged_at`, `last_report_at` | roster |

## 4. What this lane uses (see design/specs/2026-11-09-guild-dense.md)
Everything in section 1 that answers an officer's or member's weekly job. Section 2 items appear only in the spec's "needs data" list unless the lane adds them to the API.
