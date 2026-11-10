package gg.foreversixty.app.activity

import androidx.annotation.StringRes
import gg.foreversixty.app.R
import gg.foreversixty.app.bis.SpecCatalog
import gg.foreversixty.app.menu.MenuDestination
import gg.foreversixty.app.model.ActivityRow
import gg.foreversixty.app.navigation.DeepLink
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.push.PushGroup
import gg.foreversixty.app.push.PushLanding
import gg.foreversixty.app.push.PushMessage
import java.time.Instant
import java.time.LocalDate
import java.time.ZoneId

/** Activity's filters (spec 2026-11-08 4.4): All, then the notification groups; account rows show under All only. */
enum class ActivityFilter(val group: PushGroup?, @param:StringRes val label: Int) {
    All(null, R.string.activity_all),
    Raids(PushGroup.Raids, R.string.push_group_raids),
    Gear(PushGroup.Gear, R.string.push_group_gear),
    Logs(PushGroup.Logs, R.string.push_group_logs),
    Ladder(PushGroup.Ladder, R.string.push_group_ladder),
    Tiers(PushGroup.Tiers, R.string.push_group_tiers),
}

/** The day groups rows fall in, newest first. */
enum class ActivityDay(@param:StringRes val label: Int) {
    Today(R.string.activity_today),
    Yesterday(R.string.activity_yesterday),
    ThisWeek(R.string.activity_thisWeek),
    Earlier(R.string.activity_earlier),
}

/** One row as drawn: its day, its group (the icon) and when it happened. */
data class ActivityItem(val row: ActivityRow, val at: Instant, val group: PushGroup)

data class ActivitySection(val day: ActivityDay, val items: List<ActivityItem>)

/** Where a tap on a row lands (spec F4): an app screen, else the tab its group belongs under. */
sealed interface ActivityLanding {
    data class Link(val link: DeepLink) : ActivityLanding
    data class Tab(val tab: ForeverTab) : ActivityLanding
    data class Menu(val destination: MenuDestination) : ActivityLanding
}

object Activity {
    private const val WEEK_DAYS = 7L

    /** The kinds whose title opens on a spec's full name (`Fury Warrior moved to S`); a BiS change keeps its icon, it is about an item. */
    private val NAMED_SPEC_KINDS = setOf("tier.moved")

    fun items(rows: List<ActivityRow>): List<ActivityItem> = rows.mapNotNull { row ->
        val at = runCatching { Instant.parse(row.occurredAt) }.getOrNull() ?: return@mapNotNull null
        ActivityItem(row, at, PushGroup.of(row.group, row.kind))
    }

    /** Today, Yesterday, This week (the six days before), Earlier, in [zone]; empty days left out. */
    fun sections(items: List<ActivityItem>, now: Instant, zone: ZoneId): List<ActivitySection> {
        val today = now.atZone(zone).toLocalDate()
        return items.groupBy { dayOf(it.at.atZone(zone).toLocalDate(), today) }
            .toSortedMap()
            .map { (day, rows) -> ActivitySection(day, rows) }
    }

    fun dayOf(date: LocalDate, today: LocalDate): ActivityDay = when {
        !date.isBefore(today) -> ActivityDay.Today
        date == today.minusDays(1) -> ActivityDay.Yesterday
        date.isAfter(today.minusDays(WEEK_DAYS)) -> ActivityDay.ThisWeek
        else -> ActivityDay.Earlier
    }

    /**
     * The class a row is about, drawn as its crest (one crest language, never a square spec icon): the
     * row's own `crest` when the server names the other player's class, else the
     * spec icon the server sends for a guildmate (guildsim.passed), else the spec a tier move names
     * at its start (`Fury Warrior moved to S`). Null for every other row: it keeps its game icon.
     */
    fun crestClass(row: ActivityRow, specs: SpecCatalog): String? =
        row.crest?.takeIf { it.isNotBlank() }
            ?: specs.specs.firstOrNull { it.icon == row.icon }?.classSlug
            ?: row.takeIf { it.kind in NAMED_SPEC_KINDS }?.let { named ->
                specs.specs.firstOrNull { named.title.startsWith("${it.name} ${it.classSlug} ", ignoreCase = true) }?.classSlug
            }

    /** The row's answer: its app link (or the push's raid, sim or report landing), else its group's tab. */
    fun landing(row: ActivityRow): ActivityLanding? {
        val message = PushMessage(
            kind = row.kind,
            group = PushGroup.of(row.group, row.kind),
            title = row.title,
            body = row.body,
            link = row.link.ifBlank { null },
            category = null,
            collapse = null,
            data = emptyMap(),
        )
        PushLanding.of(message)?.let { return ActivityLanding.Link(it) }
        return when (message.group) {
            PushGroup.Raids, PushGroup.Officer, PushGroup.Ladder -> ActivityLanding.Tab(ForeverTab.Raid)
            PushGroup.Gear -> ActivityLanding.Tab(ForeverTab.Gear)
            // tier.moved: Tiers, the spec marked (spec F4).
            PushGroup.Tiers -> ActivityLanding.Menu(MenuDestination.Tiers)
            PushGroup.Logs -> ActivityLanding.Menu(MenuDestination.Logs)
            PushGroup.Account -> null
        }
    }
}
