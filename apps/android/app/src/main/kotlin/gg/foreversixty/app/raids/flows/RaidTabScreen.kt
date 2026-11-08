package gg.foreversixty.app.raids.flows

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.padding
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.ParityRegion
import gg.foreversixty.app.flows.FlowBlock
import gg.foreversixty.app.flows.FlowButton
import gg.foreversixty.app.flows.FlowCountdown
import gg.foreversixty.app.flows.FlowHead
import gg.foreversixty.app.flows.FlowLink
import gg.foreversixty.app.flows.FlowLead
import gg.foreversixty.app.flows.FlowNote
import gg.foreversixty.app.flows.FlowPage
import gg.foreversixty.app.flows.FlowRow
import gg.foreversixty.app.flows.FlowSection
import gg.foreversixty.app.flows.FlowWideButton
import gg.foreversixty.app.home.flows.HomeActions
import gg.foreversixty.app.home.flows.OfficerItems
import gg.foreversixty.app.home.flows.RaidCard
import gg.foreversixty.app.home.flows.RaidCardModel
import gg.foreversixty.app.home.flows.RaidLiveModel

/** One of the guild's raid pages under the Raid tab (Roster, Calendar, Loot history, Guild sims). */
data class GuildPageRow(val page: GuildPage, val title: String, val sub: String?)

enum class GuildPage { Roster, Calendar, Loot, Sims }

/** A guild that recruits the player's spec, with Apply. */
data class RecruitingGuild(val id: Long, val name: String, val line: String)

/** One item of Ready check: ticked, missing or a warning, and where to get it. */
data class ReadyItem(val tone: ReadyTone, val title: String, val sub: String, val where: String?, val link: String? = null)

enum class ReadyTone { Ok, Miss, Warn }

/** What the Raid tab shows (spec 2026-11-08 4.2). */
sealed interface RaidTabView {
    data object SignedOut : RaidTabView

    /** Under 60: when raids start, and Find a guild. */
    data class Leveler(val levelsLeft: Int) : RaidTabView

    /** No guild: the guilds recruiting the spec, then Claim your guild. */
    data class NoGuild(val spec: String?, val guilds: List<RecruitingGuild>) : RaidTabView

    data class Guild(
        val guild: String,
        val next: RaidCardModel?,
        /** An officer's open items and their count, the badge's number. */
        val needs: OfficerItems?,
        /** The newest pending applicant the applicant row opens; null opens the guild page. */
        val applicantId: Long? = null,
        val needsCount: Int,
        val live: RaidLiveModel?,
        /** Raid night's roles with names (`Tanks 2 of 2`, `Grunvald, Brakka`). */
        val roles: List<Pair<String, String>>,
        val ready: List<ReadyItem>,
        val readyStamp: String?,
        val pages: List<GuildPageRow>,
    ) : RaidTabView
}

/** What the Raid tab's taps do; the route binds them. */
data class RaidTabActions(
    val card: HomeActions = HomeActions(),
    val onPage: (GuildPage) -> Unit = {},
    val onApply: (RecruitingGuild) -> Unit = {},
    val onClaim: () -> Unit = {},
    val onFindGuild: () -> Unit = {},
    val onSignIn: () -> Unit = {},
    val onWhere: (ReadyItem) -> Unit = {},
    val onApplicant: (Long) -> Unit = {},
)

/** The Raid tab root: the next raid first with your answer, then what needs an officer, then the guild's pages. */
@Composable
fun RaidTabScreen(view: RaidTabView, actions: RaidTabActions) {
    FlowPage {
        when (view) {
            RaidTabView.SignedOut -> {
                FlowHead(stringResource(R.string.raid_tab_withGuild), listOf(stringResource(R.string.raid_tab_outLine)))
                FlowWideButton(stringResource(R.string.signin_button), primary = true, onClick = actions.onSignIn)
            }
            is RaidTabView.Leveler -> {
                FlowHead(stringResource(R.string.raid_tab_leveler), listOf(stringResource(R.string.raid_tab_levelsToGo, view.levelsLeft)))
                FlowNote(stringResource(R.string.raid_tab_levelerLine), Modifier.padding(top = FsSpace.space2))
                FlowWideButton(stringResource(R.string.home_findGuild), primary = false, onClick = actions.onFindGuild)
            }
            is RaidTabView.NoGuild -> NoGuild(view, actions)
            is RaidTabView.Guild -> Guild(view, actions)
        }
    }
}

@Composable
private fun NoGuild(view: RaidTabView.NoGuild, actions: RaidTabActions) {
    FlowHead(stringResource(R.string.raid_tab_withGuild), listOf(stringResource(R.string.raid_tab_withGuildLine)))
    if (view.guilds.isNotEmpty() && view.spec != null) {
        FlowSection(stringResource(R.string.raid_tab_recruiting, view.spec)) {
            view.guilds.forEachIndexed { index, guild ->
                FlowRow(guild.name, guild.line, last = index == view.guilds.lastIndex, trail = {
                    FlowButton(stringResource(R.string.raid_tab_apply), { actions.onApply(guild) })
                })
            }
        }
    }
    FlowSection(stringResource(R.string.raid_tab_notHere)) {
        FlowNote(stringResource(R.string.raid_tab_claimLine), Modifier.padding(top = FsSpace.space2))
        FlowWideButton(stringResource(R.string.raid_tab_claim), primary = false, onClick = actions.onClaim)
    }
}

@Composable
private fun Guild(view: RaidTabView.Guild, actions: RaidTabActions) {
    val live = view.live
    if (live != null) {
        FlowSection(stringResource(R.string.raid_tab_raidNight)) {
            FlowBlock(ParityRegion.COUNTDOWNS) { FlowCountdown(live.eyebrow, emptyList(), "", live = true) { FlowNote(live.line) } }
            view.roles.forEachIndexed { index, (title, names) -> FlowRow(title, names, slim = true, last = index == view.roles.lastIndex) }
        }
        if (view.ready.isNotEmpty()) ReadyCheck(view, actions)
    } else {
        FlowSection(stringResource(R.string.raid_tab_nextRaid)) {
            val next = view.next
            if (next != null) {
                FlowBlock(ParityRegion.COUNTDOWNS) { RaidCard(next.copy(officer = null), actions.card) }
            } else {
                FlowNote(stringResource(R.string.raid_tab_noRaid), Modifier.padding(top = FsSpace.space2))
            }
        }
        view.needs?.let { needs ->
            FlowSection(stringResource(R.string.raid_tab_needsYou), meta = stringResource(R.string.raid_tab_open, view.needsCount)) {
                Needs(view.next?.raidId ?: 0L, needs, actions.card) { view.applicantId?.let(actions.onApplicant) ?: actions.card.onApplicants() }
            }
        }
    }
    FlowSection(view.guild) {
        view.pages.forEachIndexed { index, page ->
            FlowRow(page.title, page.sub, last = index == view.pages.lastIndex, onClick = { actions.onPage(page.page) })
        }
    }
}

@Composable
private fun Needs(raidId: Long, needs: OfficerItems, actions: HomeActions, openApplicant: () -> Unit) {
    needs.unanswered?.let { line ->
        FlowRow(line, needs.short, last = needs.applicants == null, trail = {
            FlowButton(stringResource(R.string.home_officer_preview), { actions.onPreviewMention(raidId) })
        })
    }
    needs.applicants?.let { FlowRow(it, needs.applicantLine, last = true, onClick = openApplicant) }
}

@Composable
private fun ReadyCheck(view: RaidTabView.Guild, actions: RaidTabActions) {
    FlowSection(stringResource(R.string.raid_tab_readyCheck)) {
        view.ready.forEachIndexed { index, item ->
            FlowRow(
                item.title,
                item.sub,
                FlowLead.Check(item.tone == ReadyTone.Ok, item.tone == ReadyTone.Miss),
                last = index == view.ready.lastIndex,
                trail = item.where?.let { where -> { FlowLink(where) { actions.onWhere(item) } } },
            )
        }
        view.readyStamp?.let { Column(Modifier.padding(top = FsSpace.space2)) { FlowNote(it, color = FsColor.muted) } }
    }
}
