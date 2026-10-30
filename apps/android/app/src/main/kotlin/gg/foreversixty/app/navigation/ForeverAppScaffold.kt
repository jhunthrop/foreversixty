package gg.foreversixty.app.navigation

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
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.unit.Density
import androidx.compose.ui.res.stringResource
import gg.foreversixty.app.R
import gg.foreversixty.app.account.AccountDeps
import gg.foreversixty.app.account.HomeAccountGate
import gg.foreversixty.app.account.YouRoute
import gg.foreversixty.app.home.HomeDependencies
import gg.foreversixty.app.home.HomeRoute

/**
 * The app shell: a Material 3 navigation bar over one screen per tab. [account] is lane F1's
 * session and sign-in, null while its feature switch is off.
 */
@Composable
fun ForeverAppScaffold(home: HomeDependencies, account: AccountDeps? = null) {
    var selected by rememberSaveable { mutableStateOf(ForeverTab.Home) }
    val openBis = { selected = ForeverTab.Bis }
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
            ForeverTab.Raids -> ComingSoon(selected)
            ForeverTab.Bis -> ComingSoon(selected)
            ForeverTab.You -> if (account == null) ComingSoon(selected) else YouRoute(account, openBis)
        }
    }
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
