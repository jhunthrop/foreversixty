package gg.foreversixty.app.navigation

import androidx.compose.runtime.getValue
import androidx.compose.runtime.setValue
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Icon
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.NavigationBar
import androidx.compose.material3.NavigationBarItem
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.unit.Density
import androidx.compose.ui.res.stringResource
import androidx.lifecycle.compose.collectAsStateWithLifecycle
import gg.foreversixty.app.R
import gg.foreversixty.app.account.AccountDeps
import gg.foreversixty.app.account.CharacterSelectorBar
import gg.foreversixty.app.account.HomeAccountGate
import gg.foreversixty.app.account.YouRoute
import gg.foreversixty.app.auth.session.SessionPhase
import gg.foreversixty.app.home.HomeDependencies
import gg.foreversixty.app.bis.BisModel
import gg.foreversixty.app.bis.BisOwner
import gg.foreversixty.app.bis.BisRoute
import gg.foreversixty.app.home.HomeRoute
import gg.foreversixty.app.raids.RaidsFeature
import gg.foreversixty.app.raids.RaidsRoute

/**
 * The app shell: a Material 3 navigation bar over one screen per tab. [account] is lane F1's
 * session and sign-in, null while its feature switch is off.
 */
@Composable
fun ForeverAppScaffold(
    home: HomeDependencies,
    bis: BisModel? = null,
    account: AccountDeps? = null,
    raids: RaidsFeature? = null,
) {
    var selected by rememberSaveable { mutableStateOf(ForeverTab.Home) }
    val openBis = { selected = ForeverTab.Bis }
    // Lane F4: a raid link or notification selects the Raids tab.
    val wantsRaids = raids?.navigator?.wantsTab == true
    LaunchedEffect(wantsRaids) {
        if (wantsRaids) {
            selected = ForeverTab.Raids
            raids.navigator.wantsTab = false
        }
    }
    // Lane F5: a sim link selects Home, where Sims and the result are pushed.
    val wantsHome = home.sims?.navigator?.wantsHome == true
    LaunchedEffect(wantsHome) {
        if (wantsHome) {
            selected = ForeverTab.Home
            home.sims?.navigator?.wantsHome = false
        }
    }
    ForeverTabScaffold(selected = selected, onSelect = { selected = it }) {
        // The tab slots. Each lane replaces its placeholder branch: F2 Home, F4 Raids, F3 BiS, F1 You.
        when (selected) {
            ForeverTab.Home ->
                if (account == null) {
                    HomeRoute(home, openBis = openBis) { ComingSoon(selected) }
                } else {
                    HomeAccountGate(account, openBis) {
                        HomeRoute(home, account.session, openBis) { ComingSoon(selected) }
                    }
                }
            ForeverTab.Raids -> if (raids == null) ComingSoon(selected) else RaidsTab(raids, account)
            ForeverTab.Bis ->
                if (bis == null) {
                    ComingSoon(selected)
                } else {
                    BisTab(bis, home, account)
                }
            ForeverTab.You -> if (account == null) ComingSoon(selected) else YouRoute(account, openBis)
        }
    }
}

/**
 * The BiS tab under the character selector (spec §3.4): the list follows the session's current
 * character, and waits while a signed-in session's characters are still on their way.
 */
@Composable
private fun BisTab(bis: BisModel, home: HomeDependencies, account: AccountDeps?) {
    val session = account?.session
    val phase = session?.phase?.collectAsStateWithLifecycle()?.value
    val picked = session?.pickedKey?.collectAsStateWithLifecycle()?.value
    val characterKey = phase?.let { session.currentCharacter(it, picked)?.key }
    val waiting = phase == SessionPhase.Importing || (phase is SessionPhase.SignedIn && phase.me == null)
    BisRoute(bis, home.catalog, BisOwner(characterKey, waiting)) { if (session != null) CharacterSelectorBar(session) }
}

/**
 * Raids (lane F4) for the session's current character's guild, under the same selector as Home
 * and BiS; signed out or with sign-in off, it follows `/v1/me`'s main character.
 */
@Composable
private fun RaidsTab(raids: RaidsFeature, account: AccountDeps?) {
    val session = account?.session
    val phase = session?.phase?.collectAsStateWithLifecycle()?.value
    val picked = session?.pickedKey?.collectAsStateWithLifecycle()?.value
    val characterKey = phase?.let { session.currentCharacter(it, picked)?.key }
    RaidsRoute(raids, characterKey) { if (session != null) CharacterSelectorBar(session) }
}

/** The scaffold with a navigation bar, stateless so screenshot tests can drive it. */
@Composable
fun ForeverTabScaffold(
    selected: ForeverTab,
    onSelect: (ForeverTab) -> Unit,
    content: @Composable () -> Unit,
) {
    Scaffold(
        containerColor = MaterialTheme.colorScheme.background,
        bottomBar = {
            NavigationBar(containerColor = MaterialTheme.colorScheme.surface) {
                ForeverTab.entries.forEach { tab ->
                    NavigationBarItem(
                        selected = tab == selected,
                        onClick = { onSelect(tab) },
                        icon = { Icon(tab.icon, contentDescription = null) },
                        label = { TabLabel(stringResource(tab.label)) },
                    )
                }
            }
        },
    ) { padding ->
        Box(Modifier.fillMaxSize().padding(padding)) { content() }
    }
}

/** The largest font scale a tab label grows to; past it the label stays one line (spec §10). */
private const val TAB_LABEL_MAX_FONT_SCALE = 1.3f

/** A tab label on one line whatever the font scale: the scale is capped, the icon carries the rest. */
@Composable
private fun TabLabel(text: String) {
    val density = LocalDensity.current
    CompositionLocalProvider(
        LocalDensity provides Density(density.density, density.fontScale.coerceAtMost(TAB_LABEL_MAX_FONT_SCALE)),
    ) {
        Text(text, maxLines = 1, softWrap = false)
    }
}

@Composable
private fun ComingSoon(tab: ForeverTab) {
    Box(Modifier.fillMaxSize(), contentAlignment = Alignment.Center) {
        Text(
            text = stringResource(R.string.shell_arrivesNext, stringResource(tab.label)),
            style = MaterialTheme.typography.bodyLarge,
            color = MaterialTheme.colorScheme.onSurfaceVariant,
        )
    }
}
