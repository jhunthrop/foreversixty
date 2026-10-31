"""Boards for design/specs/2026-10-31-mobile-wave3.md. Run from the repo root:
python3 design/mocks/2026-10-31-mobile-wave3/gen_wave3.py && node design/mocks/2026-10-31-mobile-wave3/render.mjs
One iOS and one Android board per feature F9 to F16 (iPhone 17 402 x 874 pt, Pixel 9 412 x 915 dp)."""
import sys
from pathlib import Path

HERE = Path(__file__).resolve().parent
sys.path.insert(0, str(HERE))
from w3kit import DEVICES, board, cell, offline_strip  # noqa: E402
import screens_w3a as A  # noqa: E402
import screens_w3b as B  # noqa: E402
import screens_w3c as C  # noqa: E402

RENDERS = HERE / "renders"


def boards(p: str) -> dict[str, tuple[str, list[str]]]:
    i = p == "ios"
    f14 = ([cell("<b>F14.1 Home screen, 7:07 pm on raid day.</b> Small next raid, small next upgrade, medium next raid with In / Late / Out, medium upgrades.", C.home_widgets(p)),
            cell("<b>F14.2 Widget states.</b> Every small state: signed out, no guild, none planned, locked, cancelled, older than 12 h, unanswered, no spec.", C.widget_states(p)),
            cell("<b>F14.3 Lock screen.</b> Inline, rectangular and circular widgets, and the raid Live Activity from 1 hour before invites.", C.lock_widgets()),
            cell("<b>F14.4 Dynamic Island and sim.</b> Compact and expanded raid countdown; the server sim running and finished.", C.island())]
           if i else
           [cell("<b>F14.1 Home screen.</b> Next raid and next upgrade at 2 x 2, next raid with In / Late / Out and upgrades at 4 x 2.", C.home_widgets(p)),
            cell("<b>F14.2 Widget states.</b> The same eight states as iOS.", C.widget_states(p)),
            cell("<b>F14.3 Shade and tile.</b> Ongoing raid countdown, server sim progress and finished, the Next raid quick-settings tile.", C.and_ongoing())])
    f15 = ([cell("<b>F15.1 Siri, sign up.</b> The App Intent asks once, then answers with the count.", C.siri("ios")),
            cell("<b>F15.2 Siri, next upgrade.</b> Item icon, quality colour, source and the measured gain.", C.siri("ios", "up")),
            cell("<b>F15.3 Spotlight.</b> Characters, guides, raids and talents, and the two shortcuts.", C.spotlight()),
            cell("<b>F15.4 Offline data.</b> You, Offline data.", C.offline_data(p)),
            cell("<b>F15.5 Stale and new-data states.</b>", C.banner_states(p))]
           if i else
           [cell("<b>F15.1 Launcher shortcuts.</b> Long-press the icon.", C.launcher_menu()),
            cell("<b>F15.2 Assistant.</b> The Open app feature action.", C.assistant()),
            cell("<b>F15.3 Launcher search.</b> Character shortcuts.", C.launcher_search()),
            cell("<b>F15.4 Offline data.</b> You, Offline data.", C.offline_data(p)),
            cell("<b>F15.5 Stale and new-data states.</b>", C.banner_states(p))])
    return {
        f"f9-guides-{p}": ("F9 guides, tier lists and what changed", [
            cell("<b>F9.1 Guides tab.</b> The fifth tab. Your guide first, then all nine classes.", A.guides_root(p)),
            cell("<b>F9.2 A guide.</b> Role and stat priority first; sections in a strip.", A.guide_page(p)),
            cell("<b>F9.3 Talents.</b> Every talent has its icon, rank and the game's own text.", A.guide_mid(p)),
            cell("<b>F9.3b Rotation.</b> The same page scrolled. From the guide's own loop; Flame Shock is shown as not in it.", A.guide_mid(p, 895)),
            cell("<b>F9.4 Tier lists.</b> Sample player: a Fire Mage. Elemental ships only after its band is verified (spec, F9 accuracy gate).", A.tier_list(p)),
            cell("<b>F9.5 What changed.</b> Key dates, then the changelog.", A.whats_new(p)),
            cell("<b>F9.6 Offline class page.</b> Saved and not-saved guides.", A.class_page(p)),
            cell("<b>F9.7 Loading.</b>", A.guide_loading(p))]),
        f"f10-reports-{p}": ("F10 your reports", [
            cell("<b>F10.1 Your reports.</b> Live, processing, failed and finished.", A.reports_list(p)),
            cell("<b>F10.2 A report.</b> Your parses first; trash as one entry.", A.report_page(p)),
            cell("<b>F10.3 A fight.</b> Meters with your row marked.", A.report_page(p, overlay=A.fight_sheet(p))),
            cell("<b>F10.4 Coach (Premium).</b> The report scrolled down. Free: no card, no prompt.", A.report_page(p, scroll=330)),
            cell("<b>F10.5 Report states.</b>", A.report_states(p)),
            cell("<b>F10.6 No reports yet.</b> The web's sentence.", A.reports_list(p, "empty")),
            cell("<b>F10.7 Loading and error.</b>", A.reports_list(p, "loading")),
            cell("<b>F10.8 Home gains Logs.</b> And Rankings (F11); the search button joins the bar.", A.home_rows(p))]),
        f"f11-rankings-{p}": ("F11 rankings and character pages", [
            cell("<b>F11.1 Rankings.</b> Your rank, one summary line, your row marked in place.", A.rankings(p)),
            cell("<b>F11.2 Filters.</b> A sheet; every choice applies at once.", A.filters_sheet(p)),
            cell("<b>F11.3 Off the page.</b> Rank 204 pins your best kill above the board.", A.rankings(p, "pinned")),
            cell("<b>F11.4 A character.</b> Forever Score, bests, recent fights.", A.char_page(p, 1)),
            cell("<b>F11.5 Gear.</b> The paperdoll, gear gap, builds seen.", A.char_page(p, 2)),
            cell("<b>F11.6 No logs yet.</b>", A.char_page(p, 1, "empty")),
            cell("<b>F11.7 Not found.</b>", A.char_page(p, 1, "notfound"))]),
        f"f12-search-{p}": ("F12 search and the item page", [
            cell("<b>F12.1 Search, empty.</b> For your character, watching, recent.", B.search_empty(p)),
            cell("<b>F12.2 All results.</b> Guides, talents, characters.", B.search_all(p)),
            cell("<b>F12.3 Items.</b> Quality colour, slot, level.", B.search_items(p)),
            cell("<b>F12.4 An item.</b> Your pick, the tooltip, where it comes from. Build 70658.", B.item_page(p)),
            cell("<b>F12.5 Best in slot for.</b> Every spec that lists it.", B.item_page(p, 2)),
            cell("<b>F12.6 No results.</b>", B.search_none(p))]),
        f"f13-planner-{p}": ("F13 the talent planner", [
            cell("<b>F13.1 Plan a build.</b> Tap adds a point, long-press removes one.", B.editor(p, "mid")),
            cell("<b>F13.2 A refusal.</b> The inspector says why, in words.", B.editor(p, "low")),
            cell("<b>F13.3 Share.</b> The web's confirm step.", B.editor(p, "mid", B.confirm_sheet(p))),
            cell("<b>F13.4 System share.</b> The link carries the build card.", B.editor(p, "mid", B.share_sheet(p))),
            cell("<b>F13.5 A shared build.</b> Opened from foreversixty.gg/b/…", B.viewer(p)),
            cell("<b>F13.6 Your builds.</b>", B.my_builds(p))]),
        f"f14-surfaces-{p}": ("F14 home-screen and lock-screen surfaces", f14),
        f"f15-system-{p}": ("F15 system integration", f15),
        f"f16-guild-{p}": ("F16 guild extras", [
            cell("<b>F16.1 Roster.</b> Spec, level, rank, attendance.", C.roster(p)),
            cell("<b>F16.2 No match.</b>", C.roster(p, "empty")),
            cell("<b>F16.3 Raid history.</b> Paged by cursor. Drawn on Thu 17 Dec; F16.4 to F16.7 on Tue 8 Dec.", C.raid_history(p)),
            cell("<b>F16.4 Calendar.</b> A calendar the app keeps up to date, and a link.", C.calendar_settings(p)),
            cell("<b>F16.5 Access off.</b>", C.calendar_settings(p, True)),
            cell("<b>F16.6 The event.</b> " + ("Apple Calendar." if p == "ios" else "Google Calendar."), C.cal_event(p)),
            cell("<b>F16.7 Raids tab.</b> Roster and Raid history rows under the guild head.", A.raids_rows(p))]),
    }


def main() -> None:
    RENDERS.mkdir(exist_ok=True)
    for p in ("ios", "and"):
        for name, (title, cells) in boards(p).items():
            (RENDERS / f"{name}.html").write_text(board(title, DEVICES[p]["label"], cells))
            print(name)


if __name__ == "__main__":
    main()
