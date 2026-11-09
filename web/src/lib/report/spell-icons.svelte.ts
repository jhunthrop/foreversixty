// web/src/lib/report/spell-icons.svelte.ts
// The report tabs' spell icons (logs spec 2026-11-07 section 9, best-addon tenet): the build's
// one icon table (`spellicons.json`, lib/guild/report-spell-icons.ts), read once per page and
// shared by every tab. Until it arrives (and in a server render) every row draws the dark socket
// at its final size, so nothing moves when the icons land.
import activeBuild from '../../data/active-build.json';
import { NO_SPELL_ICONS, loadSpellIcons, type SpellIconOf } from '../guild/report-spell-icons';

let table: Promise<SpellIconOf> | null = null;

/** The active build's icon table, fetched on the first call. */
export function reportSpellIcons(): Promise<SpellIconOf> {
  table ??= loadSpellIcons(activeBuild.build, { chain: true }).catch((error: unknown) => {
    table = null;
    throw error;
  });
  return table;
}

/** A component's view of the table: `of` is the socket until the table loads. Construct it during
 *  a component's init (it starts the load in an effect). */
export class SpellIcons {
  of = $state<SpellIconOf>(NO_SPELL_ICONS);

  constructor(load: () => Promise<SpellIconOf> = reportSpellIcons) {
    $effect(() => {
      let live = true;
      load()
        .then((of) => {
          if (live) this.of = of;
        })
        .catch((error: unknown) => {
          // A missing table is a page of sockets, not a broken tab.
          console.warn('report: the spell icon table did not load', error);
        });
      return () => {
        live = false;
      };
    });
  }
}
