package gg.foreversixty.app.raids.flows

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.bis.BisRange
import gg.foreversixty.app.home.flows.RaidCardWords
import gg.foreversixty.app.model.RaidBadge
import gg.foreversixty.app.model.RaidView
import gg.foreversixty.app.raids.logic.PlannedRaids
import gg.foreversixty.app.raids.logic.RaidClock
import gg.foreversixty.app.raids.logic.RaidsTabContent
import gg.foreversixty.app.raids.logic.RaidsTabState
import gg.foreversixty.app.raids.logic.SignupController
import gg.foreversixty.app.raids.logic.SignupState
import gg.foreversixty.app.home.flows.HomeFlowState
import java.time.Instant

/**
 * The Raid tab's view from the tab model's state (spec 2026-11-08 4.2). Null for the states the tab
 * still draws with its older views (loading, failed, waiting for approval).
 */
class RaidTabPresenter(private val res: Resources, private val clock: RaidClock, private val now: Instant) {
    fun view(state: RaidsTabState, controller: SignupController?, badge: RaidBadge?, hasExtras: Boolean): RaidTabView? = when (state) {
        RaidsTabState.SignedOut -> RaidTabView.SignedOut
        is RaidsTabState.NoGuild -> {
            val level = state.level
            if (level != null && level < BisRange.LAST_BAND) RaidTabView.Leveler(BisRange.LAST_BAND - level) else RaidTabView.NoGuild(state.spec, emptyList())
        }
        is RaidsTabState.Loaded -> guild(state.content, controller, badge, hasExtras)
        else -> null
    }

    private fun guild(content: RaidsTabContent, controller: SignupController?, badge: RaidBadge?, hasExtras: Boolean): RaidTabView.Guild {
        val words = RaidCardWords(res, clock, now)
        val key = content.character.key
        val next = PlannedRaids(content.planned, now).next?.let { raid ->
            val answer = controller?.takeIf { it.raid.id == raid.id }?.block(key)?.shown ?: words.ownAnswer(raid, key)
            val own = raid.signups.firstOrNull { it.own && it.characterKey == key }
            words.card(raid, content.head.name, key, answer, lateOpen = own?.lateMinutes == null, badge = badge)
        }
        val live = content.planned.firstOrNull { HomeFlowState.isLive(it, now) }
        return RaidTabView.Guild(
            guild = content.head.name,
            next = next,
            needs = next?.officer,
            needsCount = badge?.count ?: 0,
            live = live?.let { words.live(it, content.head.name, key, words.ownAnswer(it, key)) },
            roles = live?.let(::roles).orEmpty(),
            ready = emptyList(),
            readyStamp = null,
            pages = pages(content, hasExtras),
        )
    }

    /** Raid night's tanks and healers with their names (`Brightwen (Late, 15 min)`). */
    private fun roles(raid: RaidView): List<Pair<String, String>> {
        val target = raid.compTarget
        return listOf(TANK to R.string.raid_group_tanks, HEALER to R.string.raid_group_healers).mapNotNull { (role, label) ->
            val coming = raid.signups.filter { it.role == role && (SignupState.of(it.state)?.isComing == true) }
            if (coming.isEmpty()) return@mapNotNull null
            val names = coming.joinToString(", ") { s ->
                val minutes = s.lateMinutes
                if (s.state == SignupState.Late.wire && minutes != null) res.getString(R.string.raid_tab_lateNote, s.name, minutes) else s.name
            }
            val want = if (role == TANK) target?.tanks else target?.healers
            val title = res.getString(label).let { word -> want?.let { res.getString(R.string.raid_tab_roleLine, word, coming.size, it) } ?: word }
            title to names
        }
    }

    /** The guild's raid pages: Roster and Loot history where the app has them, Calendar from the planned raids. */
    private fun pages(content: RaidsTabContent, hasExtras: Boolean): List<GuildPageRow> {
        val calendar = content.planned.take(CALENDAR_ITEMS).joinToString(", ") {
            res.getString(R.string.raid_tab_calendarItem, it.zoneName.substringBefore("'"), clock.shortDate(Instant.parse(it.startsAt)).substringBefore(" "))
        }.ifEmpty { null }
        return listOfNotNull(
            GuildPageRow(GuildPage.Roster, res.getString(R.string.raid_tab_roster), null).takeIf { hasExtras },
            GuildPageRow(GuildPage.Calendar, res.getString(R.string.raid_tab_calendar), calendar).takeIf { content.head.webUrl != null },
            GuildPageRow(GuildPage.Loot, res.getString(R.string.raid_tab_loot), null).takeIf { hasExtras },
            GuildPageRow(GuildPage.Sims, res.getString(R.string.raid_tab_sims), null).takeIf { content.head.webUrl != null },
        )
    }

    private companion object {
        const val TANK = "tank"
        const val HEALER = "healer"
        const val CALENDAR_ITEMS = 2
    }
}
