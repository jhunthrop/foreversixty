package gg.foreversixty.app.raids.flows

import gg.foreversixty.app.raids.logic.AssignmentEncounterLine
import gg.foreversixty.app.raids.logic.RaidJobLine
import androidx.compose.foundation.layout.Column
import androidx.compose.ui.platform.testTag
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.PanelFrame
import gg.foreversixty.app.designsystem.ParityRegion
import gg.foreversixty.app.flows.FlowBandPage
import gg.foreversixty.app.flows.FlowBlock
import gg.foreversixty.app.flows.FlowButton
import gg.foreversixty.app.flows.FlowLead
import gg.foreversixty.app.flows.FlowLink
import gg.foreversixty.app.flows.FlowNote
import gg.foreversixty.app.flows.FlowPageState
import androidx.compose.foundation.rememberScrollState
import gg.foreversixty.app.flows.FlowPanel
import gg.foreversixty.app.flows.FlowRow
import gg.foreversixty.app.flows.FlowSection
import gg.foreversixty.app.flows.FlowSubLink
import gg.foreversixty.app.flows.TallyItem
import gg.foreversixty.app.flows.SUB
import gg.foreversixty.app.raids.logic.NotesText
import gg.foreversixty.app.raids.ui.LocalRaidIconBuild
import androidx.compose.material3.Text
import androidx.compose.ui.text.style.TextOverflow
import gg.foreversixty.app.home.flows.HomeActions
import gg.foreversixty.app.home.flows.OfficerItems
import gg.foreversixty.app.home.flows.RaidCardModel
import gg.foreversixty.app.home.flows.RaidLiveModel

/** One of the guild's raid pages under the Raid tab (Roster, Calendar, Loot history, Guild sims). */
data class GuildPageRow(val page: GuildPage, val title: String, val sub: String?)

enum class GuildPage { Roster, Calendar, Loot, Sims }

/** A guild that recruits the player's spec, with Apply. */
data class RecruitingGuild(val id: Long, val name: String, val line: String, val applied: String? = null)

/**
 * One item of Before you pull: ticked, missing or a warning, and where to get it; [icon] the item the
 * check is about (the attunement item, the boots) when the data has its icon.
 */
data class ReadyItem(val tone: ReadyTone, val title: String, val sub: String, val where: String?, val link: String? = null, val icon: String? = null)

enum class ReadyTone { Ok, Miss, Warn }

/** One role of Who is coming: its tally (S6), the names shown, how many more, and Damage's split. */
data class RoleLine(val tally: TallyItem, val names: List<RoleName>, val more: String? = null, val split: String? = null) {
    val role: String get() = tally.label
    val have: Int get() = tally.have
    val want: Int? get() = tally.want

    /** The names line as plain text (TalkBack, tests): `Mendara, Gravemend, Rainspeaker and 3 more`. */
    val namesText: String get() = names.joinToString(", ") { it.text } + more.orEmpty()

    /** The line without its target, `Healers 5`: under 60 the Guild tab shows who is coming, not what the raid still needs. */
    fun countOnly(): RoleLine = copy(tally = tally.copy(want = null))
}

/**
 * A line closing Who is coming: the counts, the Covered line ([cover], its names in wipe red beside their
 * icons, [text] read aloud), an officer's readiness (a tap opens the roster on [onlyKeys], the raiders
 * it names).
 */
data class RoleFoot(val text: String, val cover: CoverLine? = null, val onlyKeys: Set<String> = emptySet())

/** A name on Who is coming in its class colour: `Brightwen (Late · 15 min)`. */
data class RoleName(val text: String, val className: String)

/** What the Raid tab shows (spec 2026-11-08 4.2). */
sealed interface RaidTabView {
    data object SignedOut : RaidTabView

    /** No guild: the guilds recruiting the spec (Applied once you did), then Find your guild; under 60 one quiet line. */
    data class NoGuild(val spec: String?, val guilds: List<RecruitingGuild>, val raidsOpen: String? = null) : RaidTabView

    data class Guild(
        val guild: String,
        val next: RaidCardModel?,
        /** An officer's open items and their count, the badge's number. */
        val needs: OfficerItems?,
        /** The newest pending applicant the applicant row opens; null opens the guild page. */
        val applicantId: Long? = null,
        val needsCount: Int,
        val live: RaidLiveModel?,
        /** Who is coming (or, live, who is in): each role with its names. */
        val roles: List<RoleLine>,
        /** The lines under it: the other answers, what the comp misses, an officer's readiness. */
        val foot: List<RoleFoot> = emptyList(),
        /** The raid's officer notes (HTML), two lines under the head; the rest on the raid page. */
        val notes: String? = null,
        /** An officer's Needs you sits directly under the head (G3). */
        val needsFirst: Boolean = false,
        val ready: List<ReadyItem>,
        val readyStamp: String?,
        val pages: List<GuildPageRow>,
        /** The Guild tab's own sections (spec 2026-11-10-guild-tab.md); null until the guild home loads. */
        val tab: GuildTabSections? = null,
        /** Under 60: `Raids open at 60 · you're 36` under the raid head. */
        val raidsOpen: String? = null,
        /** Your job in the raid (spec 2.3a); empty when you hold no slot. */
        val jobs: List<RaidJobLine> = emptyList(),
        /** An officer's read-only assignments; null for everyone else. */
        val assignments: List<AssignmentEncounterLine>? = null,
        /** The pages section's title once the guild head shows the name (`Pages`); null keeps the guild's name. */
        val pagesTitle: String? = null,
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
    /** The not-synced line's tap (row 5): the roster filtered to the characters it names. */
    val onRosterOf: (Set<String>) -> Unit = {},
    /** The Covered line's `and N more` (round 5 row 4): the raid page at its cover list. */
    val onCover: (Long) -> Unit = {},
    /** The Guild tab's own sections (spec 2026-11-10-guild-tab.md). */
    val guild: GuildTabActions = GuildTabActions(),
)

/**
 * The Raid tab root in the round-5 look (spec 4.2, mocks m-android-raid and m-android-signup): the
 * band behind [topBar], then on the ground the one-line raid head with your answer, Who is coming,
 * Before you pull, what needs an officer and the guild's pages. [notice] sits under the band.
 */
@Composable
fun RaidTabScreen(
    view: RaidTabView,
    actions: RaidTabActions,
    topBar: @Composable () -> Unit = {},
    page: FlowPageState = FlowPageState(rememberScrollState()),
    notice: @Composable () -> Unit = {},
) {
    FlowBandPage(topBar, notice = notice, page = page) {
        when (view) {
            RaidTabView.SignedOut -> Lone(stringResource(R.string.raid_tab_withGuild), emptyList()) {
                RaidWideButton(stringResource(R.string.signin_button), primary = true, onClick = actions.onSignIn)
            }
            is RaidTabView.NoGuild -> NoGuild(view, actions)
            is RaidTabView.Guild -> Guild(view, actions)
        }
    }
}

/** A state with one thing to do: its lines on a panel under the title, then the button. */
@Composable
private fun Lone(title: String, lines: List<String>, button: @Composable () -> Unit) {
    FlowSection(title, top = 0.dp, panel = PanelFrame.None) {
        lines.forEachIndexed { index, line ->
            FlowNote(line, Modifier.padding(top = FsSpace.space2), color = if (index == 0 && lines.size > 1) FsColor.strong else FsColor.text)
        }
        button()
    }
}

@Composable
private fun NoGuild(view: RaidTabView.NoGuild, actions: RaidTabActions) {
    view.raidsOpen?.let { RaidsOpenLine(it) }
    if (view.guilds.isNotEmpty() && view.spec != null) {
        FlowSection(stringResource(R.string.raid_tab_recruiting, view.spec), top = 0.dp, panel = PanelFrame.Plain) {
            view.guilds.forEachIndexed { index, guild ->
                FlowRow(guild.name, guild.line, last = index == view.guilds.lastIndex, trail = {
                    val applied = guild.applied
                    if (applied != null) {
                        FlowNote(applied, color = FsColor.muted)
                    } else {
                        FlowButton(stringResource(R.string.raid_tab_apply), { actions.onApply(guild) }, RAID_TARGET)
                    }
                })
            }
        }
    }
    // Find your guild (spec 2.4): the site's guild finder, then Claim.
    FlowSection(stringResource(R.string.guild_tab_find), top = 0.dp, panel = PanelFrame.None) {
        FlowRow(stringResource(R.string.guild_tab_search), onClick = actions.onFindGuild, chevron = true)
        FlowRow(stringResource(R.string.raid_tab_claim), last = true, onClick = actions.onClaim, chevron = true)
    }
}

/** Under 60: one quiet line on the ground, never a block. */
@Composable
private fun RaidsOpenLine(text: String) {
    FlowNote(text, Modifier.fillMaxWidth().testTag("guild-raids-open"), color = FsColor.text)
}

@Composable
private fun Guild(view: RaidTabView.Guild, actions: RaidTabActions) {
    val live = view.live
    val next = view.next
    val tab = view.tab
    // Under 60 the guild leads (spec 2.1); at 60 the raid does (2.2).
    if (tab != null && tab.leveler) GuildHeadBlock(tab.head, tab.spread)
    when {
        live != null -> RaidNightBlock(live, view.roles)
        next != null -> {
            FlowBlock(ParityRegion.section(NEXT_RAID_REGION), top = 0.dp) {
                RaidHead(next, actions.card, view.notes)
            }
            view.raidsOpen?.let { RaidsOpenLine(it) }
            // G3: an officer's open items come straight after the head, the officer's rows with them; then your job.
            if (view.needsFirst) NeedsSection(view, actions)
            OfficerParts(view, actions)
            YourJobSection(view.jobs)
            if (view.roles.isNotEmpty() || view.foot.isNotEmpty()) {
                FlowSection(
                    stringResource(R.string.raid_tab_coming),
                    // With the Guild tab's Roster section on the page, its link is the one way to the roster.
                    action = if (tab == null) stringResource(R.string.raid_tab_rosterLink) else null,
                    onAction = { actions.onPage(GuildPage.Roster) },
                    top = 0.dp,
                    panel = PanelFrame.None,
                ) { RoleRows(view.roles, view.foot, actions.onRosterOf) { actions.onCover(next.raidId) } }
            }
        }
        else -> {
            FlowBlock(ParityRegion.section(NEXT_RAID_REGION), top = 0.dp) {
                FlowPanel { FlowNote(stringResource(R.string.raid_tab_noRaid), Modifier.padding(vertical = FsSpace.space2), color = FsColor.text) }
            }
            view.raidsOpen?.let { RaidsOpenLine(it) }
        }
    }
    if (view.ready.isNotEmpty()) {
        // Once pulled the list is the Ready check, not what comes before the pull.
        val title = stringResource(if (live != null) R.string.raid_tab_readyCheck else R.string.raid_tab_beforePull)
        FlowSection(title, top = 0.dp, panel = PanelFrame.None) {
            ReadyList(view.ready, view.readyStamp, actions.onWhere, inlineWhere = true)
        }
    }
    if (!view.needsFirst || next == null) NeedsSection(view, actions)
    if (next == null) OfficerParts(view, actions)
    if (tab != null) {
        if (!tab.leveler) GuildHeadBlock(tab.head, tab.spread)
        GuildRosterSection(tab, actions.guild.onRoster)
        GuildLootSection(tab.lootNight, tab.loot, actions.guild.onLoot)
        GuildLogsSection(tab.logs, actions.guild)
        GuildProgressSection(tab.progression, actions.guild.onProgression)
    }
    if (view.pages.isNotEmpty()) {
        FlowSection(view.pagesTitle ?: view.guild, top = 0.dp, panel = PanelFrame.None) {
            view.pages.forEachIndexed { index, page ->
                FlowRow(page.title, page.sub, last = index == view.pages.lastIndex, onClick = { actions.onPage(page.page) }, chevron = true)
            }
        }
    }
}

/** An officer's shortcuts and the raid's assignments, read only (spec 2.3, 2.3a). */
@Composable
private fun OfficerParts(view: RaidTabView.Guild, actions: RaidTabActions) {
    view.tab?.let { OfficerSection(it.officer, actions.guild.onOfficer) }
    view.assignments?.let { AssignmentsSection(it, actions.guild.onPlanner) }
}

/** Needs you: an officer's open items with their count. */
@Composable
private fun NeedsSection(view: RaidTabView.Guild, actions: RaidTabActions) {
    val needs = view.needs ?: return
    FlowSection(stringResource(R.string.raid_tab_needsYou), meta = stringResource(R.string.raid_tab_open, view.needsCount), top = 0.dp, panel = PanelFrame.None) {
        Needs(view.next?.raidId ?: 0L, needs, actions.card) { view.applicantId?.let(actions.onApplicant) ?: actions.card.onApplicants() }
    }
}

/** The raid's notes under the head, two lines; the raid page has the rest. */
@Composable
internal fun RaidNotes(html: String) {
    Text(
        NotesText.annotated(html, FsColor.gold),
        style = SUB,
        color = FsColor.text,
        maxLines = NOTES_LINES,
        overflow = TextOverflow.Ellipsis,
        modifier = Modifier.padding(top = FsSpace.space2),
    )
}

@Composable
private fun Needs(raidId: Long, needs: OfficerItems, actions: HomeActions, openApplicant: () -> Unit) {
    needs.unanswered?.let { line ->
        FlowRow(line, needs.short, last = needs.applicants == null, trail = {
            FlowButton(stringResource(R.string.home_officer_preview), { actions.onPreviewMention(raidId) }, RAID_TARGET)
        })
    }
    needs.applicants?.let { FlowRow(it, needs.applicantLine, last = true, onClick = openApplicant) }
}

/**
 * The Before you pull rows, each miss with where to get it, then the export's age (spec F2). With
 * [inlineWhere] (round 5) the fix ends the row's line; otherwise it sits under the row.
 */
@Composable
fun ReadyList(items: List<ReadyItem>, stamp: String?, onWhere: (ReadyItem) -> Unit, inlineWhere: Boolean = false) {
    items.forEachIndexed { index, item ->
        val where = item.where
        FlowRow(
            item.title,
            item.sub,
            FlowLead.Check(item.tone == ReadyTone.Ok, item.tone == ReadyTone.Miss, item.icon, LocalRaidIconBuild.current),
            last = index == items.lastIndex,
            trail = where?.takeIf { !inlineWhere }?.let { { FlowLink(it) { onWhere(item) } } },
            subLink = where?.takeIf { inlineWhere }?.let { FlowSubLink(it) { onWhere(item) } },
        )
    }
    stamp?.let { Column(Modifier.padding(top = FsSpace.space2, bottom = FsSpace.space2)) { FlowNote(it, color = FsColor.muted) } }
}

/** The Raid tab's region for its head: the old Next raid section's name, so every state keeps one. */
private const val NEXT_RAID_REGION = "Next raid"

/** The notes' lines on the tab (round 2); the raid page shows them whole. */
private const val NOTES_LINES = 2
