package gg.foreversixty.app.flows

import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.setValue
import gg.foreversixty.app.activity.ActivityFixtures
import gg.foreversixty.app.gear.RunPick
import gg.foreversixty.app.home.HomeCharacter
import gg.foreversixty.app.home.flows.HomeActions
import gg.foreversixty.app.home.flows.HomeComposer
import gg.foreversixty.app.home.flows.HomeFlowScreen
import gg.foreversixty.app.home.flows.HomeFlowService
import gg.foreversixty.app.home.flows.HomeFlowState
import gg.foreversixty.app.home.flows.HomeInputs
import gg.foreversixty.app.home.flows.HomeView
import gg.foreversixty.app.home.flows.RaidCardWords
import gg.foreversixty.app.menu.CharacterBar
import gg.foreversixty.app.model.CharacterHome
import gg.foreversixty.app.model.HomeSection
import gg.foreversixty.app.model.Me
import gg.foreversixty.app.model.mainCharacter
import gg.foreversixty.app.raids.MemoryAskStore
import gg.foreversixty.app.raids.MemoryRaidCache
import gg.foreversixty.app.raids.RaidFixture
import gg.foreversixty.app.raids.RecordingHaptics
import gg.foreversixty.app.raids.clock
import gg.foreversixty.app.raids.logic.NotificationPermission
import gg.foreversixty.app.raids.logic.RaidsTabModel
import gg.foreversixty.app.raids.logic.SignupController
import gg.foreversixty.app.raids.resources
import java.time.Instant
import kotlinx.coroutines.launch

/**
 * Home's raid card as [gg.foreversixty.app.home.flows.HomeFlowContent] wires it, over a fake clock:
 * HomeFlowState over the Raid tab's model, one SignupController per next raid built exactly as the
 * route builds it (no onRaidChanged), and the route's private `homeView` card rule. HomeFlowContent
 * itself cannot be hosted deterministically: it reads Instant.now() and ZoneId.systemDefault() and
 * needs a RaidsFeature, whose clock and service are fixed.
 */
class HomeFlow(
    val service: EchoRaidService = EchoRaidService(),
    val haptics: RecordingHaptics = RecordingHaptics(),
    val now: () -> Instant = { FlowTime.now },
) {
    val character: HomeCharacter = HomeCharacter.from(requireNotNull(RaidFixture.load(Me.serializer(), "me").mainCharacter))
    private val tab = RaidsTabModel(service, MemoryRaidCache(), resources(), now)
    val state = HomeFlowState(RaidDaySections, tab, null, now)
    var controller: SignupController? by mutableStateOf(null)

    /** The raid-day sections: the card leads. */
    private object RaidDaySections : HomeFlowService {
        override suspend fun sections(characterKey: String) = listOf(HomeSection(type = "next_raid", style = "card"))
        override suspend fun bestRun(home: CharacterHome): RunPick? = null
    }

    /** HomeFlowContent's `homeView`, the raid card only (no upgrades answer in these flows). */
    fun view(): HomeView {
        val res = resources()
        val composer = HomeComposer(res, now(), clock().zone)
        val words = RaidCardWords(res, clock(), now())
        val guild = state.raidContent?.head?.name.orEmpty()
        val card = state.nextRaid?.let { raid ->
            val own = raid.signups.firstOrNull { it.own && it.characterKey == character.key }
            val answer = controller?.takeIf { it.raid.id == raid.id }?.block(character.key)?.shown ?: words.ownAnswer(raid, character.key)
            words.card(raid, guild, character.key, answer, lateOpen = own?.lateMinutes == null, badge = null)
        }
        val inputs = HomeInputs(state.sections, null, null, card, null, emptyList(), 0, firstRaids = null)
        return HomeView.Blocks(composer.head(character), composer.blocks(inputs), ActivityFixtures.BUILD)
    }
}

/** Home with its raid card, inside the tab shell. */
@Composable
fun HomeFlowRoot(flow: HomeFlow) {
    val scope = rememberCoroutineScope()
    LaunchedEffect(flow) { flow.state.load(flow.character) }
    // HomeFlowContent: one controller per raid, the Raid tab's, built in composition.
    val raid = flow.state.nextRaid
    val guild = flow.state.raidContent?.head
    if (raid != null && guild != null && flow.controller?.raid?.id != raid.id) {
        flow.controller = SignupController(
            guild.id,
            raid,
            SignupController.Dependencies(flow.service, resources(), flow.haptics, NotificationPermission { false }, MemoryAskStore(), flow.now, scope, RaidFlow.HOLD),
        )
    }
    val key = flow.character.key
    val actions = remember(flow) {
        HomeActions(
            onAnswer = { _, answer -> scope.launch { flow.controller?.press(key, answer) } },
            onLate = { _, minutes -> scope.launch { flow.controller?.pressLate(key, minutes) } },
        )
    }
    HomeFlowScreen(flow.view(), actions) { CharacterBar(null) }
}
