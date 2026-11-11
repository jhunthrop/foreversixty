package gg.foreversixty.app.raids

import gg.foreversixty.app.designsystem.screenGround
import gg.foreversixty.app.common.rememberHaptics
import android.Manifest
import android.content.Context
import androidx.activity.compose.BackHandler
import androidx.activity.compose.rememberLauncherForActivityResult
import androidx.activity.result.contract.ActivityResultContracts
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateListOf
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.platform.LocalUriHandler
import androidx.compose.ui.platform.LocalView
import gg.foreversixty.app.common.Connectivity
import gg.foreversixty.app.guild.GuildExtrasFeature
import gg.foreversixty.app.guild.GuildExtrasRows
import gg.foreversixty.app.guild.HistoryRoute
import gg.foreversixty.app.guild.RosterRoute
import gg.foreversixty.app.guild.logic.CalendarInput
import gg.foreversixty.app.guild.logic.Loadable
import gg.foreversixty.app.navigation.DeepLink
import gg.foreversixty.app.navigation.RAID_VIEW_UNANSWERED
import gg.foreversixty.app.model.RaidBadge
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.network.StaticDataClient
import gg.foreversixty.app.network.dataOrNull
import gg.foreversixty.app.raids.logic.FileRaidCache
import gg.foreversixty.app.raids.logic.LiveRaidService
import gg.foreversixty.app.raids.logic.PlannedRaids
import gg.foreversixty.app.raids.logic.PrefsReminderAskStore
import gg.foreversixty.app.raids.logic.RaidCache
import gg.foreversixty.app.raids.logic.RaidClock
import gg.foreversixty.app.raids.logic.RaidNightModel
import gg.foreversixty.app.raids.logic.RaidOpenAction
import gg.foreversixty.app.raids.logic.RaidPageModel
import gg.foreversixty.app.raids.logic.RaidService
import gg.foreversixty.app.raids.logic.RaidsTabModel
import gg.foreversixty.app.raids.logic.RaidsTabState
import gg.foreversixty.app.raids.logic.ReminderAskStore
import gg.foreversixty.app.raids.logic.SignupController
import gg.foreversixty.app.raids.logic.SignupState
import gg.foreversixty.app.raids.logic.SystemNotificationPermission
import gg.foreversixty.app.raids.ui.LocalRaidIconBuild
import gg.foreversixty.app.raids.ui.RaidNightScreen
import gg.foreversixty.app.reports.ReportRoute
import gg.foreversixty.app.reports.ReportRouteScreen
import gg.foreversixty.app.reports.ReportsFeature
import gg.foreversixty.app.reports.logic.ReportsViewer
import gg.foreversixty.app.raids.ui.RaidPageScreen
import gg.foreversixty.app.raids.ui.RaidsTabContentView
import gg.foreversixty.app.raids.flows.GuildPage
import gg.foreversixty.app.raids.flows.RaidTabActions
import gg.foreversixty.app.raids.flows.RaidTabPresenter
import gg.foreversixty.app.raids.flows.RaidTabScreen
import gg.foreversixty.app.raids.flows.NotAnsweredModel
import gg.foreversixty.app.raids.flows.ApplicantModel
import gg.foreversixty.app.raids.flows.ApplicantRoute
import gg.foreversixty.app.raids.flows.LiveRaidFlowService
import gg.foreversixty.app.raids.flows.RaidFlowService
import gg.foreversixty.app.raids.flows.RaidTabExtras
import gg.foreversixty.app.raids.flows.NextRaidSignup
import gg.foreversixty.app.raids.push.RaidPushPayload
import gg.foreversixty.app.flows.FlowPageState
import gg.foreversixty.app.flows.rememberFlowRefresh
import androidx.compose.foundation.rememberScrollState
import gg.foreversixty.app.raids.logic.RaidLinks
import gg.foreversixty.app.home.flows.HomeFlowState
import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.raids.flows.NotAnsweredRoute
import gg.foreversixty.app.raids.flows.RaidNightRoute
import gg.foreversixty.app.raids.logic.RaidsTabContent
import gg.foreversixty.app.home.flows.HomeActions
import kotlinx.coroutines.CoroutineScope
import gg.foreversixty.app.common.SavedAnswerStrip
import java.io.File
import java.time.Instant
import kotlin.time.Duration.Companion.seconds
import kotlinx.coroutines.launch

/** A screen pushed on the Raids stack. */
sealed interface RaidsDestination {
    /** A raid's page; [toCover] scrolls to its composition's cover (the Raid tab's `and N more`). */
    data class Raid(val guild: Long, val raid: Long, val open: RaidOpenAction? = null, val toCover: Boolean = false) : RaidsDestination

    /** An application to the guild, for its officers (spec 2026-11-08 F3). */
    data class Applicant(val guild: Long, val application: Long) : RaidsDestination

    /** The raid page filtered to Not answered (spec 2026-11-08 F3), an officer's. */
    data class NotAnswered(val guild: Long, val raid: Long) : RaidsDestination
    data class Night(val raid: Long) : RaidsDestination

    /** Lane F10: a report pushed from Your raid night. */
    data class Report(val id: String) : RaidsDestination
    /** Lane F16: the guild's roster and its logged raid nights. */
    /** The guild's roster; [only] narrows it to those characters (the Raid tab's not-synced line). */
    data class Roster(val guild: Long, val only: Set<String> = emptySet()) : RaidsDestination
    data class History(val guild: Long) : RaidsDestination
}

/** Lane F10: the reports screens Your raid night pushes, and who reads them; null while `reports` is off. */
class RaidsReports(val feature: ReportsFeature, val viewer: ReportsViewer)

/** The Raids tab's own back stack; deep links and notifications push here (spec 5.0, 6.4). */
class RaidsNavigator {
    val stack = mutableStateListOf<RaidsDestination>()

    /** Set when a link asks for the Raids tab; the scaffold selects it and clears this. */
    var wantsTab by mutableStateOf(false)

    fun push(destination: RaidsDestination) {
        stack.add(destination)
    }

    fun pop() {
        stack.removeLastOrNull()
    }

    /** Opens a raid link on a fresh stack, so back goes to the tab's root. False for other links. */
    fun open(link: DeepLink): Boolean {
        val destinations = when (link) {
            is DeepLink.Raid -> {
                // spec=ask presses the link's state once the spec is picked; without one, In (links.json).
                val press = SignupState.of(link.state) ?: SignupState.In.takeIf { link.spec == RaidPushPayload.SPEC_ASK }
                if (link.view == RAID_VIEW_UNANSWERED) {
                    listOf(RaidsDestination.NotAnswered(link.guild, link.raid))
                } else {
                    listOf(RaidsDestination.Raid(link.guild, link.raid, open = press?.let(::RaidOpenAction)))
                }
            }
            is DeepLink.RaidNight -> listOf(RaidsDestination.Raid(link.guild, link.raid), RaidsDestination.Night(link.raid))
            is DeepLink.Guild -> emptyList()
            is DeepLink.GuildApplicant -> listOf(RaidsDestination.Applicant(link.guild, link.application))
            else -> return false
        }
        stack.clear()
        stack.addAll(destinations)
        wantsTab = true
        return true
    }
}

/** The raids feature's dependencies, built once in the app graph. */
class RaidsFeature(
    context: Context,
    apiClient: ApiClient,
    private val statics: StaticDataClient,
    /** Lane F16: roster, raid history and calendar sync; null while `guild_extras` is off. */
    val extras: GuildExtrasFeature? = null,
    private val isAvailable: () -> Boolean,
) {
    val service: RaidService = LiveRaidService(apiClient)
    val flows: RaidFlowService = LiveRaidFlowService(apiClient)
    val cache: RaidCache = FileRaidCache(File(context.filesDir, "raids"))
    val askStore: ReminderAskStore = PrefsReminderAskStore(context.getSharedPreferences("forever_raids", Context.MODE_PRIVATE))
    val navigator = RaidsNavigator()
    val now: () -> Instant = Instant::now

    /** Lane F15: whether a failed refresh is offline or stale. */
    val connectivity: Connectivity = Connectivity.of(context)

    /** The Raids tab ships dark behind `raids` (AppGraph.isAvailable: on in debug and fixture builds). */
    val enabled: Boolean get() = isAvailable()

    suspend fun iconBuild(): String? = statics.activeBuild().dataOrNull()?.build

    /** The Raid tab's badge (spec 2026-11-08 3): the member's unanswered raid or the officer's open items. */
    var badge: RaidBadge? by mutableStateOf(null)
        private set

    /** Reads the badge for [guild]; the tab bar and the officer's Needs you show its count. */
    suspend fun refreshBadge(guild: Long?) {
        badge = guild?.let { (service.badge(it) as? ApiResult.Success)?.data }
    }

    /** The applicant, recruiting guilds and Ready check the Raid tab shows beside the raids. */
    var tabExtras: RaidTabExtras by mutableStateOf(RaidTabExtras())
        private set

    /**
     * Reads what [state] needs: an officer's newest applicant, the guilds recruiting a guildless spec,
     * and the Ready check of [readyRaid] (the live raid, else the next: spec 4.2 Before you pull).
     */
    suspend fun refreshExtras(state: RaidsTabState, readyRaid: Long?) {
        tabExtras = when (state) {
            is RaidsTabState.NoGuild -> RaidTabExtras(recruiting = recruiting(state.character))
            is RaidsTabState.Loaded -> {
                val guild = state.content.head.id
                val applicant = if ((badge?.applicants ?: 0) > 0 && badge?.officer == true) newestApplicant(guild) else null
                val ready = readyRaid?.let { (flows.ready(state.content.character.key, it) as? ApiResult.Success)?.data }
                RaidTabExtras(applicant = applicant, ready = ready)
            }
            else -> RaidTabExtras()
        }
    }

    private suspend fun newestApplicant(guild: Long) = (flows.applications(guild) as? ApiResult.Success)?.data?.rows
        ?.filter { it.status == ApplicantModel.PENDING }
        ?.maxByOrNull { it.createdAt }

    private suspend fun recruiting(character: AccountCharacter?): List<gg.foreversixty.app.model.RecruitingGuild> {
        val spec = character?.specKey() ?: return emptyList()
        val faction = character.faction ?: return emptyList()
        return (flows.recruiting(spec, faction, character.region, character.ruleset) as? ApiResult.Success)?.data?.rows.orEmpty()
    }
}

/**
 * F4: the Raids tab root and its stack. [characterKey] is the session's current character (lane
 * F1), null to follow `/v1/me`'s main; [selector] is the character selector row over the root.
 */
@Composable
fun RaidsRoute(
    feature: RaidsFeature,
    characterKey: String? = null,
    reports: RaidsReports? = null,
    signIn: () -> Unit = {},
    selector: @Composable () -> Unit = {},
) {
    val context = LocalContext.current
    val view = LocalView.current
    val uri = LocalUriHandler.current
    val scope = rememberCoroutineScope()
    val res = androidx.compose.ui.platform.LocalResources.current
    val clock = remember(res) { RaidClock(res) }
    val permissionLauncher = rememberLauncherForActivityResult(ActivityResultContracts.RequestPermission()) { }
    val requestPermission = {
        // Lane F7 registers the FCM token once permission is granted.
        if (android.os.Build.VERSION.SDK_INT >= android.os.Build.VERSION_CODES.TIRAMISU) {
            permissionLauncher.launch(Manifest.permission.POST_NOTIFICATIONS)
        }
    }
    val haptics = rememberHaptics()
    val deps = remember(feature, haptics) {
        SignupController.Dependencies(
            service = feature.service,
            res = res,
            haptics = haptics,
            permission = SystemNotificationPermission(context),
            askStore = feature.askStore,
            now = feature.now,
            scope = scope,
            hold = HOLD,
        )
    }
    val tab = remember(feature) { RaidsTabModel(feature.service, feature.cache, res, feature.now, connectivity = feature.connectivity) }
    var iconBuild by remember { mutableStateOf<String?>(null) }

    val extras = feature.extras

    /** Lane F16: the calendar follows every refresh and every sign-up answer. */
    fun syncCalendar() {
        val loaded = (tab.state as? RaidsTabState.Loaded)?.content ?: return
        if (loaded.offlineSince != null || extras == null) return
        scope.launch { extras.syncCalendar(CalendarInput(loaded.head.id, loaded.head.name, loaded.planned)) }
    }

    val signup = remember(feature) { NextRaidSignup(tab, feature.now) { syncCalendar() } }
    val next = signup.controller
    // The root keeps its place while a page is pushed over it and popped (Roster, a raid).
    val rootScroll = rememberScrollState()

    suspend fun reload() {
        tab.load()
        signup.sync(deps)
        syncCalendar()
        val loaded = (tab.state as? RaidsTabState.Loaded)?.content
        feature.refreshBadge(loaded?.head?.id)
        val readyRaid = loaded?.planned?.let { planned -> planned.firstOrNull { HomeFlowState.isLive(it, feature.now()) } ?: PlannedRaids(planned, feature.now()).next }
        feature.refreshExtras(tab.state, readyRaid?.id)
        if (extras != null && loaded != null) extras.roster(loaded.head.id).load()
    }

    val refresh = rememberFlowRefresh { reload() }
    LaunchedEffect(feature) { iconBuild = feature.iconBuild() }
    LaunchedEffect(feature, characterKey) {
        tab.currentKey = characterKey
        reload()
    }
    val content = (tab.state as? RaidsTabState.Loaded)?.content
    BackHandler(enabled = feature.navigator.stack.isNotEmpty()) { feature.navigator.pop() }
    CompositionLocalProvider(LocalRaidIconBuild provides iconBuild) {
        when (val top = feature.navigator.stack.lastOrNull()) {
            null -> Column(Modifier.fillMaxSize().screenGround()) {
                val home = content?.let { extras?.roster(it.head.id)?.state as? Loadable.Loaded }?.value
                val view = RaidTabPresenter(res, clock, feature.now()).view(tab.state, next, feature.badge, extras != null, feature.tabExtras.copy(home = home))
                if (view != null) {
                    val actions = raidTabActions(feature, content, next, scope, uri::openUri, signIn)
                    RaidTabScreen(view, actions, topBar = selector, page = FlowPageState(rootScroll, refresh)) {
                        content?.offlineSince?.let { SavedAnswerStrip(clock.stripTime(it, feature.now()), content.stale, { scope.launch { reload() } }) }
                    }
                } else {
                    selector()
                    RaidsTabContentView(
                        state = tab.state,
                        clock = clock,
                        now = feature.now,
                        nextController = next,
                        openRaid = { raid, unanswered -> feature.navigator.push(raidDestination(content?.head?.id ?: 0, raid.id, unanswered)) },
                        openNight = { feature.navigator.push(RaidsDestination.Night(it)) },
                        openWeb = { uri.openUri(it) },
                        retry = { scope.launch { reload() } },
                        requestPermission = requestPermission,
                        guildRows = { loaded ->
                            if (extras != null) {
                                GuildExtrasRows(
                                    onRoster = extras.roster(loaded.head.id).onRoster,
                                    openRoster = { feature.navigator.push(RaidsDestination.Roster(loaded.head.id)) },
                                    openHistory = { feature.navigator.push(RaidsDestination.History(loaded.head.id)) },
                                )
                            }
                        },
                    )
                }
            }
            is RaidsDestination.Raid -> {
                val model = remember(top) { RaidPageModel(top.guild, top.raid, deps, feature.cache, top.open, signup::replaced, feature.connectivity, characterKey ?: content?.character?.key) }
                RaidPageScreen(
                    model = model,
                    clock = clock,
                    now = feature.now,
                    currentKey = content?.character?.key,
                    guildName = content?.head?.name.orEmpty(),
                    guildUrl = content?.head?.webUrl,
                    openNight = { feature.navigator.push(RaidsDestination.Night(it)) },
                    back = { feature.navigator.pop() },
                    requestPermission = requestPermission,
                    toCover = top.toCover,
                )
            }
            is RaidsDestination.NotAnswered -> {
                val model = remember(top) { NotAnsweredModel(top.guild, top.raid, feature.service, feature.flows, res, scope, HOLD, signup::replaced) }
                NotAnsweredRoute(model, clock, content?.head?.name.orEmpty(), content?.head?.webUrl, iconBuild, uri::openUri) { feature.navigator.pop() }
            }
            is RaidsDestination.Applicant -> {
                val model = remember(top) { ApplicantModel(top.guild, top.application, feature.flows, res) }
                ApplicantRoute(model, clock, content?.head?.name.orEmpty(), feature.now) {
                    feature.navigator.pop()
                    scope.launch { feature.refreshBadge(top.guild) }
                }
            }
            is RaidsDestination.Night -> {
                val model = remember(top) { RaidNightModel(top.raid, feature.service, feature.cache, feature.now, connectivity = feature.connectivity) }
                val openReport: ((String) -> Unit)? = reports?.let { { id -> feature.navigator.push(RaidsDestination.Report(id)) } }
                val raid = content?.planned?.firstOrNull { it.id == top.raid }
                RaidNightRoute(
                    model, clock, feature.now, raid?.compTarget, iconBuild, openReport,
                    openRoster = { content?.head?.id?.let { feature.navigator.push(RaidsDestination.Raid(it, top.raid)) } },
                    openWeb = uri::openUri,
                ) { feature.navigator.pop() }
            }
            is RaidsDestination.Report -> reports?.let {
                ReportRouteScreen(it.feature, ReportRoute(top.id), it.viewer) { feature.navigator.pop() }
            }
            is RaidsDestination.Roster -> extras?.let {
                RosterRoute(it.roster(top.guild), clock, feature.now, top.only) { feature.navigator.pop() }
            }
            is RaidsDestination.History -> extras?.let {
                val pager = remember(top) { it.history(top.guild) }
                HistoryRoute(pager, tab.accountKeys, clock, feature.now, { raid -> feature.navigator.push(RaidsDestination.Night(raid)) }) {
                    feature.navigator.pop()
                }
            }
        }
    }
}

/** The Raid tab root's taps: answers through the next raid's controller, pages pushed on the tab's stack. */
private fun raidTabActions(
    feature: RaidsFeature,
    content: RaidsTabContent?,
    next: SignupController?,
    scope: CoroutineScope,
    openWeb: (String) -> Unit,
    signIn: () -> Unit,
): RaidTabActions {
    val guild = content?.head
    val key = content?.character?.key
    val openRaid = { raid: Long, unanswered: Boolean -> guild?.let { feature.navigator.push(raidDestination(it.id, raid, unanswered)) } ?: Unit }
    val guildPage = { guild?.webUrl?.let(openWeb) ?: Unit }
    return RaidTabActions(
        card = HomeActions(
            onAnswer = { _, state -> if (key != null) scope.launch { next?.press(key, state) } },
            onLate = { _, minutes -> if (key != null) scope.launch { next?.pressLate(key, minutes) } },
            onRetryAnswer = { if (key != null) scope.launch { next?.retry(key) } },
            onRaid = { openRaid(it, false) },
            onReady = { openRaid(it, false) },
            onPreviewMention = { openRaid(it, true) },
            onApplicants = guildPage,
        ),
        onPage = { page ->
            when (page) {
                GuildPage.Roster -> guild?.let { feature.navigator.push(RaidsDestination.Roster(it.id)) }
                GuildPage.Loot -> guild?.let { feature.navigator.push(RaidsDestination.History(it.id)) }
                GuildPage.Calendar, GuildPage.Sims -> guildPage()
            }
        },
        onRosterOf = { keys -> guild?.let { feature.navigator.push(RaidsDestination.Roster(it.id, keys)) } },
        onCover = { raid -> guild?.let { feature.navigator.push(RaidsDestination.Raid(it.id, raid, toCover = true)) } },
        onApply = { openWeb("${RaidLinks.SITE}/app/guild/${it.id}") },
        onClaim = { openWeb(GUILDS) },
        onFindGuild = { openWeb(GUILDS) },
        onSignIn = signIn,
        onApplicant = { id -> guild?.let { feature.navigator.push(RaidsDestination.Applicant(it.id, id)) } },
        onWhere = { item -> item.link?.let(openWeb) },
    )
}

/** The raid page, or an officer's Not answered view of it. */
fun raidDestination(guild: Long, raid: Long, unanswered: Boolean): RaidsDestination =
    if (unanswered) RaidsDestination.NotAnswered(guild, raid) else RaidsDestination.Raid(guild, raid)

/** The site's guild finder, where a player applies to or claims a guild. */
private const val GUILDS = "https://foreversixty.gg/guild"

/** `Saved.`, `Copied` and the lock toast stay 4 s (spec F4.3, F4.18, F4.21). */
private val HOLD = 4.seconds

/** The specs.json key (`warrior-fury`) from the character's class slug and spec name; null when either is unknown. */
internal fun AccountCharacter.specKey(): String? {
    val cls = className?.takeIf { it.isNotBlank() } ?: return null
    val name = spec?.takeIf { it.isNotBlank() } ?: return null
    return "${cls.lowercase()}-${name.lowercase().replace(' ', '-')}"
}
