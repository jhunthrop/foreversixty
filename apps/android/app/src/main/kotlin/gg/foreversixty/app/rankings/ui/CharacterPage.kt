package gg.foreversixty.app.rankings.ui

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.runtime.Composable
import androidx.compose.runtime.remember
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.LiveRegionMode
import androidx.compose.ui.semantics.liveRegion
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.bis.SpecCatalog
import gg.foreversixty.app.common.OfflineStamp
import gg.foreversixty.app.common.SavedAnswerStrip
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.flows.FlowNote
import gg.foreversixty.app.flows.FlowPage
import gg.foreversixty.app.model.CharacterFight
import gg.foreversixty.app.model.CharacterProfile
import gg.foreversixty.app.model.FeedRow
import gg.foreversixty.app.model.Item
import gg.foreversixty.app.paperdoll.Paperdoll
import gg.foreversixty.app.paperdoll.PaperdollSlot
import gg.foreversixty.app.rankings.logic.CharacterCard
import gg.foreversixty.app.rankings.logic.CharacterExtras
import gg.foreversixty.app.rankings.logic.CharacterLink
import gg.foreversixty.app.rankings.logic.CharacterPageComposer
import gg.foreversixty.app.rankings.logic.CharacterPhase
import gg.foreversixty.app.rankings.logic.CharacterRoute
import gg.foreversixty.app.rankings.logic.RankingsViewer
import gg.foreversixty.app.rankings.logic.WatchState
import java.time.Instant
import java.time.ZoneId

/** What a character page draws, apart from its store so the parity tests drive every state. */
data class CharacterScreenState(
    val route: CharacterRoute,
    val phase: CharacterPhase,
    val watch: WatchState?,
    val viewer: RankingsViewer,
    /** The character's class items file, by id, for the gear's icons and quality. */
    val items: Map<Int, Item> = emptyMap(),
    /** The data build the icons come from. */
    val build: String? = null,
    /** The `reports` switch: a best fight opens its report. */
    val reportsOn: Boolean = false,
    /** The band's upgrades, the owner's standing and the activity, as they answer. */
    val extras: CharacterExtras = CharacterExtras(),
)

data class CharacterActions(
    val back: () -> Unit = {},
    val retry: () -> Unit = {},
    val toggleWatch: () -> Unit = {},
    val share: () -> Unit = {},
    /** A link on the page: the planner, the simulator, best in slot, a web page. */
    val open: (CharacterLink) -> Unit = {},
    /** A paperdoll cell: its item's tooltip. */
    val openSocket: (PaperdollSlot) -> Unit = {},
    val openFight: (CharacterFight) -> Unit = {},
    val openActivity: (FeedRow) -> Unit = {},
    /** Not found: look a name up in Search. */
    val lookUp: (String) -> Unit = {},
    val signIn: () -> Unit = {},
    val getAddon: () -> Unit = {},
)

/** F11.4 to F11.7 (design/specs/2026-11-06-character.md): one character's page under its back bar. */
@Composable
fun CharacterContent(state: CharacterScreenState, specs: SpecCatalog, now: Instant, actions: CharacterActions = CharacterActions(), zone: ZoneId = ZoneId.systemDefault()) {
    val loaded = state.phase as? CharacterPhase.Loaded
    Column(Modifier.fillMaxSize()) {
        loaded?.offlineSince?.let { SavedAnswerStrip(OfflineStamp.format(it, now, zone), loaded.stale, actions.retry) }
        FlowPage(top = 0.dp) {
            when (val phase = state.phase) {
                CharacterPhase.Loading -> CharacterLoading()
                CharacterPhase.Failed -> CharacterError(state.route.name, actions.retry)
                CharacterPhase.NotFound -> CharacterNotFound(state.route, state.viewer.signedIn, actions)
                is CharacterPhase.Loaded -> Loaded(phase.profile, state, specs, now, zone, actions)
            }
        }
    }
}

@Composable
private fun Loaded(profile: CharacterProfile, state: CharacterScreenState, specs: SpecCatalog, now: Instant, zone: ZoneId, actions: CharacterActions) {
    val res = LocalResources.current
    val view = remember(profile, state.extras, state.items, state.viewer, now) {
        val slots = Paperdoll.fromProfile(profile.gear, state.items, res, CharacterCard.classSlug(profile))
        CharacterPageComposer(res, specs, zone).compose(profile, slots, state.extras, state.viewer, now)
    }
    CharacterHead(view, actions)
    if (state.watch?.failed == true) {
        FlowNote(
            stringResource(R.string.char_watchFailed, view.name),
            Modifier.semantics { liveRegion = LiveRegionMode.Polite },
            color = FsColor.wipe,
        )
    }
    CharacterSummary(view, actions)
    if (view.cells.isNotEmpty()) CharacterGear(view, state.build.orEmpty(), actions)
    CharacterScore(view, actions, top = if (view.cells.isEmpty()) FsSpace.space5 else SECTION_GAP)
    CharacterFights(view, state.reportsOn, actions)
    CharacterActivity(view.activity, now, actions)
    CharacterOwner(view, actions)
}

/** The 48 between the page's sections (`.section + .section`). */
internal val SECTION_GAP = FsSpace.space12
