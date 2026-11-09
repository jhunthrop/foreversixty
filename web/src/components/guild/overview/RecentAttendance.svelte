<!-- web/src/components/guild/overview/RecentAttendance.svelte -->
<!-- Attendance, last 4 raids (web layout spec 2026-11-05 §7.1, officers): the slipping raiders and
     the bench, one column per raid night. Nothing to show reads one quiet line, never an empty grid. -->
<script lang="ts">
  import { DEFAULT_LOCALE } from '../../../lib/guild/raid-time';
  import type { GuildRecentAttendance } from '../../../lib/guild/overview-api';
  import { attendanceGridRows, attendanceTitle } from '../../../lib/guild/overview-role-view';
  import Grid from '../../ui/Grid.svelte';
  import Section from '../../ui/Section.svelte';

  let { page, rosterHref, timezone }: { page: GuildRecentAttendance; rosterHref: string; timezone: string } =
    $props();

  const day = (iso: string) =>
    new Intl.DateTimeFormat(DEFAULT_LOCALE, { weekday: 'short', timeZone: timezone }).format(new Date(iso));
  const rows = $derived(attendanceGridRows(page));
</script>

<Section
  title={attendanceTitle(page.raids.length)}
  panel
  flush
  action={{ label: 'Roster', href: rosterHref }}
  testid="guild-overview-attendance"
>
  {#if rows.length === 0}
    <p class="text-muted text-small px-5 pb-3">Everyone has made their raids.</p>
  {:else}
    <Grid
      caption="Attendance over the last raids"
      columns={page.raids.map((raid) => day(raid.starts_at))}
      {rows}
      edge
    />
  {/if}
</Section>
