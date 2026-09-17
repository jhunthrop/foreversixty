// web/src/lib/report/death-health.ts
// Which health readings on a death card can be believed.
//
// The client writes the target's advanced block at full health on a swing a shield ate
// whole, and it keeps writing that stale reading on the next swing or two, so a dying
// player's bar jumps back up in the middle of the run-up with nothing healing them. The
// engine already drops the reading on the absorbed swing itself (summary/deaths.go: the
// landed line only fills in health when something landed), but the rows after it carry
// the same fiction, and the recap's "the highest they stood at in that span" quoted it:
// Hanabanana's 1:18 card on the sample log read 73% where the real ceiling was 47%.
//
// Damage only ever takes health off. So a reading on a damage row is believed only when
// it is no higher than the last reading that was believed; a heal is the one thing that
// legitimately lifts the bar, so a heal's own reading is always believed and becomes the
// new ceiling for the rows after it. A reading that cannot be believed is dropped rather
// than replaced with a guess: the card prints no bar and no figure there, exactly as it
// already does for the absorbed swing.

/** A death card's row, as much of it as a health reading needs. */
export interface HealthEvent {
  kind: 'damage' | 'heal';
  hp_after?: number;
  max_hp?: number;
}

/** A reading as a percentage of the full bar, or null when the log carries no reading. */
export function healthPct(reading: { hp_after?: number; max_hp?: number }): number | null {
  if (!reading.max_hp) return null;
  return Math.max(0, Math.min(100, ((reading.hp_after ?? 0) / reading.max_hp) * 100));
}

/**
 * The believable health percentage for each row of a death card, in the order given (the
 * card's own order: hits and heals interleaved by time), null where there is none.
 */
export function believableHealth(events: readonly HealthEvent[]): (number | null)[] {
  let ceiling: number | null = null;
  return events.map((event) => {
    const pct = healthPct(event);
    if (pct === null) return null;
    // Percentages divide integers, so a hit that took nothing off can land a hair above
    // the last reading on the same health; the epsilon keeps that from reading as a rise.
    if (event.kind === 'damage' && ceiling !== null && pct > ceiling + 1e-9) return null;
    ceiling = pct;
    return pct;
  });
}
