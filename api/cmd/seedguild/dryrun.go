// api/cmd/seedguild/dryrun.go
//
// --dry-run prints exactly what --apply would write, and nothing else: row counts per
// table, then the first two rows of each. It touches no table this tool would otherwise
// write to - the guild and owner lookups are the same read-only queries --apply itself
// opens with.
package main

import "fmt"

// accountRows is the plan's characters reduced to one row per distinct account, in the
// order each account first appears - what the users/guild_members tables actually get
// one row per, now that two roster entries can share an account (roster.go's Account).
func accountRows(plan seedPlan) []plannedCharacter {
	seen := map[string]bool{}
	out := make([]plannedCharacter, 0, plan.accountCount())
	for _, pc := range plan.Characters {
		if !seen[pc.AccountKey] {
			seen[pc.AccountKey] = true
			out = append(out, pc)
		}
	}
	return out
}

// PrintDryRun prints the whole plan's row counts and first-two-rows preview.
func PrintDryRun(plan seedPlan) {
	fmt.Printf("seedguild: dry run for guild %d (%s), tag %q\n", plan.Guild.ID, plan.Guild.Name, plan.Tag)
	fmt.Println()

	accounts := accountRows(plan)
	printTable("users", len(accounts), 2, func(i int) string {
		return fmt.Sprintf("battletag=%s role=user", accounts[i].Battletag)
	})
	printTable("characters", len(plan.Characters), 2, func(i int) string {
		pc := plan.Characters[i]
		return fmt.Sprintf("key=%s name=%s class=%s account=%s", pc.Key, pc.Mock.Name, pc.Mock.Class, pc.AccountKey)
	})
	printTable("addon_exports", len(plan.Characters), 2, func(i int) string {
		pc := plan.Characters[i]
		return fmt.Sprintf("character_key=%s updated_at=%s", pc.Key, pc.ExportedAt.Format("2006-01-02T15:04Z"))
	})
	printTable("guild_characters", len(plan.Characters), 2, func(i int) string {
		pc := plan.Characters[i]
		return fmt.Sprintf("character_key=%s rank=%s verified=%v", pc.Key, pc.Rank, pc.Verified)
	})
	printTable("guild_members", len(accounts), 2, func(i int) string {
		return fmt.Sprintf("user=%s consent=%s", accounts[i].Battletag, accounts[i].Mock.Consent)
	})

	printTable("reports", len(plan.Reports), 2, func(i int) string {
		r := plan.Reports[i]
		return fmt.Sprintf("id=%s title=%q zone=%s created_at=%s fights=%d attendees=%d",
			r.ID, r.Title, r.Plan.Zone, r.CreatedAt.Format("2006-01-02T15:04Z"), len(r.Plan.Fights),
			len(attendees(plan.Characters, r.Plan.Tag)))
	})

	var fightRows, metricRows []string
	fightCount, metricCount := 0, 0
	for _, r := range plan.Reports {
		present := attendees(plan.Characters, r.Plan.Tag)
		roster := mockRosterFrom(present)
		keys := make([]string, len(present))
		for i, pc := range present {
			keys[i] = pc.Key
		}
		for idx, f := range r.Plan.Fights {
			fightCount++
			if len(fightRows) < 2 {
				fightRows = append(fightRows, fmt.Sprintf("report_id=%s fight_index=%d name=%q kill=%v duration_ms=%d players=%d",
					r.ID, idx, f.Name, f.Kill, f.DurationMS, len(present)))
			}
			metricCount += len(present)
			if len(metricRows) < 2 {
				rows := buildFightMetrics(roster, keys, f, seededRand(r.Plan.Tag, idx))
				metricRows = append(metricRows, fmt.Sprintf(
					"report_id=%s fight_index=%d player_key=%s role=%s dps=%.0f hps=%v deaths=%d",
					r.ID, idx, rows[0].PlayerKey, rows[0].Role, rows[0].MetricDPS, rows[0].MetricHPS, rows[0].Deaths))
			}
		}
	}
	printPrecomputedTable("fights", fightCount, fightRows)
	printPrecomputedTable("fight_metrics", metricCount, metricRows)

	fmt.Println()
	fmt.Println("mutations (not row counts - existing rows this would change):")
	fmt.Printf("  guilds id=%d: claimed_by -> owner, officer_max_rank_index -> %d\n", plan.Guild.ID, seedOfficerMaxRankIndex)
	fmt.Println("  guild_characters (owner's existing row(s)): rank -> leader, verified_at -> now")
	fmt.Println()
	fmt.Println("ratings: not written - see ApplySeed's own printed explanation (R2 summary dependency); " +
		"README.md's Readiness/Ratings section has the full reasoning.")
}

func printTable(name string, count, previewN int, row func(i int) string) {
	fmt.Printf("%s: %d row(s)\n", name, count)
	for i := 0; i < previewN && i < count; i++ {
		fmt.Printf("  %d: %s\n", i+1, row(i))
	}
}

func printPrecomputedTable(name string, count int, rows []string) {
	fmt.Printf("%s: %d row(s)\n", name, count)
	for i, r := range rows {
		fmt.Printf("  %d: %s\n", i+1, r)
	}
}
