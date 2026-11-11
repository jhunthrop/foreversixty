// web/src/lib/guild/raid-comp.ts
// The raid comp (design/specs/2026-11-05-guild-raid-comp.md): the one class order, each row's
// spec and class as the character panel says it, the spec picker's (spec, role) pairs, the late
// clause and the class row. The roster group is the API's (`group`), never rebuilt here.
import { SPECS, type Spec } from '../sim/specs';
import type {
  GuildRaidCounts,
  GuildRaidGroup,
  GuildRaidRole,
  GuildRaidRowComp,
  GuildRaidSignup,
  GuildRaidView,
} from './api';
import type { SelectOption } from '../ui/listbox-model';
import { SEP, compCopy, copyList, kindCountsRoles, kindOf, raidCopy } from './raid-copy';
import { specName } from './spec-name';
import { CLASS_SLUGS } from '../ui/primitives';

/** The game's class order (Warrior to Druid): the class row, every roster group, the copied list
 *  and the Discord post (`raidspec.ClassOrder`). */
export const CLASS_ORDER: readonly string[] = CLASS_SLUGS;

/** The roster's four In groups, in their order. */
export const RAID_GROUPS: readonly GuildRaidGroup[] = ['tank', 'healer', 'melee', 'ranged'];

/** The DPS specs that also sign up as tank (spec 3.2.1): Fury and Arms warriors off-tank. */
const OFF_TANK_SPECS: ReadonlySet<string> = new Set(['warrior-arms', 'warrior-fury']);

/** The class's place in the game's order; an unknown class sorts last. */
export function classRank(classSlug: string): number {
  const index = CLASS_ORDER.indexOf(classSlug.toLowerCase());
  return index < 0 ? CLASS_ORDER.length : index;
}

/** `Warrior` from `warrior`. */
export function classWord(classSlug: string): string {
  return specName(classSlug.toLowerCase());
}

/** The catalogue spec of the class with this key, else `null`. */
export function catalogueSpec(classSlug: string, specKey: string | null | undefined): Spec | null {
  if (specKey == null) return null;
  return SPECS.find((spec) => spec.spec === specKey && spec.class_slug === classSlug.toLowerCase()) ?? null;
}

/** `Holy Paladin`, `Beast Mastery Hunter`; the class alone when the spec does not resolve. */
export function specClassLine(classSlug: string, specKey: string | null | undefined): string {
  const spec = catalogueSpec(classSlug, specKey);
  return spec === null ? classWord(classSlug) : spec.full_name;
}

/** Whether a row tanks on a spec whose own role is not tank: the Fury warrior off-tank. */
function offTanks(spec: Spec | null, role: GuildRaidRole): boolean {
  return spec !== null && role === 'tank' && spec.role !== 'tank';
}

/** The strip's Who line: `Holy Paladin`, and `Fury Warrior · Tank` for an off-tank pick. Below
 *  768 the spec select's face drops the role, so the line always carries it (`withRole`):
 *  `Restoration Druid · Healer`, `Fire Mage · DPS`. */
export function whoLine(
  classSlug: string,
  specKey: string | null | undefined,
  role: GuildRaidRole,
  withRole = false,
): string {
  const line = specClassLine(classSlug, specKey);
  const spec = catalogueSpec(classSlug, specKey);
  const showsRole = spec !== null && (withRole || offTanks(spec, role));
  return showsRole ? `${line}${SEP}${raidCopy.roleLabels[role]}` : line;
}

/** The row's roster group: the API's; a row from an API that predates it reads its role's
 *  group, DPS as ranged (no class or spec rule is rebuilt in the browser). */
export function rowGroup(row: GuildRaidRowComp & { role: GuildRaidRole }): GuildRaidGroup {
  if (row.group !== undefined) return row.group;
  return row.role === 'dps' ? 'ranged' : row.role;
}

/** One option of the spec picker: a (spec, role) pair. */
export interface SpecPick {
  /** The option's value: the spec key, `warrior-fury:tank` for an off-tank pair. */
  value: string;
  spec: Spec;
  role: GuildRaidRole;
  /** `Holy · Healer`, `Fury · Tank`. */
  label: string;
}

const PAIR_SEP = ':';

/** The value of a (spec, role) pair. */
export function pickValue(spec: Spec, role: GuildRaidRole): string {
  return role === spec.role ? spec.spec : `${spec.spec}${PAIR_SEP}${role}`;
}

function picksOf(spec: Spec): SpecPick[] {
  const roles: GuildRaidRole[] = OFF_TANK_SPECS.has(spec.spec) ? [spec.role, 'tank'] : [spec.role];
  return roles.map((role) => ({
    value: pickValue(spec, role),
    spec,
    role,
    label: compCopy.specOption(spec.name, raidCopy.roleLabels[role]),
  }));
}

/** The class's picks in catalogue order, an off-tank pair after its DPS one (spec 3.2). */
export function specPicks(classSlug: string): SpecPick[] {
  return SPECS.filter((spec) => spec.class_slug === classSlug.toLowerCase()).flatMap(picksOf);
}

/** The spec select's options: each pick with its icon, labelled `Holy · Healer` in the list. On a
 *  phone (`compact`) the closed field shows the spec name alone, so `Restoration` and `Beast
 *  Mastery` fit its 176 px; the Who line carries the role there instead (raid comp spec 3.2). */
export function specOptions(
  picks: readonly SpecPick[],
  compact: boolean,
  iconOf: (spec: Spec) => string,
): SelectOption[] {
  return picks.map((pick) => ({
    value: pick.value,
    label: pick.label,
    ...(compact ? { faceLabel: pick.spec.name } : {}),
    icon: { kind: 'spec' as const, src: iconOf(pick.spec) },
  }));
}

/** The pick with this value, else `null` (a forged value is refused, never sent). */
export function pickOf(classSlug: string, value: string): SpecPick | null {
  return specPicks(classSlug).find((pick) => pick.value === value) ?? null;
}

/** The pick a spec key and the sign-up's role make: the off-tank pair when the role is tank and
 *  the spec has one, else the spec's own; `null` when the key is not the class's. */
export function pickFor(
  classSlug: string,
  specKey: string | null | undefined,
  role: GuildRaidRole,
): SpecPick | null {
  const spec = catalogueSpec(classSlug, specKey);
  if (spec === null) return null;
  const pairRole = offTanks(spec, role) && OFF_TANK_SPECS.has(spec.spec) ? 'tank' : spec.role;
  return pickOf(classSlug, pickValue(spec, pairRole));
}

/** The late rows' tanks and healers, for the late clause. */
export function lateRoles(signups: readonly GuildRaidSignup[]): { tanks: number; healers: number } {
  const late = signups.filter((row) => row.state === 'late');
  return {
    tanks: late.filter((row) => rowGroup(row) === 'tank').length,
    healers: late.filter((row) => rowGroup(row) === 'healer').length,
  };
}

/** The count line's clauses: `3 tanks`, `7 healers`, `7 melee`, `5 ranged`, then the late clause
 *  when someone is late. */
export function compBreakdownParts(
  counts: Pick<GuildRaidCounts, 'tanks' | 'healers' | 'dps' | 'melee' | 'ranged' | 'late'>,
  late: { tanks: number; healers: number },
): string[] {
  const melee = counts.melee ?? 0;
  const ranged = counts.ranged ?? counts.dps - melee;
  const parts = [
    compCopy.tanks(counts.tanks),
    compCopy.healers(counts.healers),
    compCopy.melee(melee),
    compCopy.ranged(ranged),
  ];
  return counts.late > 0 ? [...parts, compCopy.late(counts.late, late.tanks, late.healers)] : parts;
}

/** The part of a view the comp reads. */
type CompView = Pick<GuildRaidView, 'counts' | 'signups' | 'zone_kind'>;

/** A raid view's breakdown clauses. A guild event says no role words: only `2 late` when someone
 *  is late (guild round 3). */
export function raidBreakdownParts(raid: CompView): string[] {
  if (!kindCountsRoles(kindOf(raid))) {
    return raid.counts.late > 0 ? [compCopy.late(raid.counts.late, 0, 0)] : [];
  }
  return compBreakdownParts(raid.counts, lateRoles(raid.signups));
}

const NO_BREAK_SPACE = ' ';

/** The breakdown as the page shows it: each clause glued with no-break spaces, so a line wraps
 *  only between clauses and `1 late (healer)` never splits (the copied list keeps plain spaces). */
export function raidBreakdownLine(raid: CompView): string {
  return raidBreakdownParts(raid)
    .map((clause) => clause.replaceAll(' ', NO_BREAK_SPACE))
    .join(SEP);
}

/** One class of the class row. */
export interface ClassRowItem {
  classSlug: string;
  /** `Priest`. */
  name: string;
  count: number;
  /** `Holy 3, Discipline 1, Shadow 1`: the item's `title`. */
  specs: string;
  /** `Priest 5: Holy 3, Discipline 1, Shadow 1`. */
  accessibleName: string;
}

function specOrder(classSlug: string, key: string): number {
  const index = SPECS.findIndex(
    (spec) => spec.spec === key.split(PAIR_SEP)[0] && spec.class_slug === classSlug,
  );
  return index < 0 ? SPECS.length : index;
}

/** The class's In and Late specs, most first, then catalogue order; an unresolved spec counts
 *  as Unknown, an off-tank as `Fury (tank)`. */
function specParts(classSlug: string, rows: readonly GuildRaidSignup[]): string {
  const tally = new Map<string, { label: string; n: number }>();
  for (const row of rows) {
    const spec = catalogueSpec(classSlug, row.spec_key);
    const key = spec === null ? '' : pickValue(spec, offTanks(spec, row.role) ? 'tank' : spec.role);
    const label =
      spec === null
        ? compCopy.unknownSpec
        : offTanks(spec, row.role)
          ? compCopy.offTank(spec.name)
          : spec.name;
    tally.set(key, { label, n: (tally.get(key)?.n ?? 0) + 1 });
  }
  return [...tally.entries()]
    .sort(([a, x], [b, y]) => y.n - x.n || specOrder(classSlug, a) - specOrder(classSlug, b))
    .map(([, part]) => compCopy.specCount(part.label, part.n))
    .join(', ');
}

/** The class counts: the API's, else tallied from the rows (an API that predates them). */
function classCounts(raid: Pick<GuildRaidView, 'counts' | 'signups' | 'not_signed'>): Record<string, number> {
  if (raid.counts.classes !== undefined) return { ...raid.counts.classes };
  const out: Record<string, number> = {};
  for (const row of [...raid.not_signed, ...raid.signups]) out[row.class.toLowerCase()] ??= 0;
  for (const row of raid.signups) {
    if (row.state === 'in' || row.state === 'late') out[row.class.toLowerCase()] += 1;
  }
  return out;
}

/** The band's class row (spec 3.1 line 6): the classes with someone In or Late, in the game's
 *  order, with that count and their specs. A class with nobody counted is left out, on raids,
 *  runs and events alike (guild round 3: a benched paladin read as a greyed `Paladin 0`);
 *  empty when nobody is In or Late, and always for a guild event, which counts no classes. */
export function classRow(raid: CompView & Pick<GuildRaidView, 'not_signed'>): ClassRowItem[] {
  if (!kindCountsRoles(kindOf(raid))) return [];
  const counts = classCounts(raid);
  return Object.keys(counts)
    .filter((slug) => slug !== '' && (counts[slug] ?? 0) > 0)
    .sort((a, b) => classRank(a) - classRank(b))
    .map((slug) => {
      const rows = raid.signups.filter(
        (row) => row.class.toLowerCase() === slug && (row.state === 'in' || row.state === 'late'),
      );
      const name = classWord(slug);
      const count = counts[slug] ?? 0;
      const specs = specParts(slug, rows);
      return { classSlug: slug, name, count, specs, accessibleName: compCopy.classItem(name, count, specs) };
    });
}

/** Class in the game's order, then name: the order within every roster group. */
export function classThenName(
  a: { class: string; name: string },
  b: { class: string; name: string },
): number {
  return classRank(a.class) - classRank(b.class) || a.name.localeCompare(b.name);
}

/** The officer's `Copy the sign-up list` text (raid comp spec 4.2), with the comp rules filled in. */
export function signupListText(view: GuildRaidView, tz: string, locale?: string): string {
  return copyList(
    view,
    tz,
    {
      breakdown: raidBreakdownParts(view),
      classes: classRow(view),
      groupOf: rowGroup,
      specClass: (row) => specClassLine(row.class, row.spec_key),
      order: classThenName,
      roles: kindCountsRoles(kindOf(view)),
    },
    locale,
  );
}
