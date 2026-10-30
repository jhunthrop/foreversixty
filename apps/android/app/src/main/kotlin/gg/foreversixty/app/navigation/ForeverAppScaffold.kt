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
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import gg.foreversixty.app.R
import gg.foreversixty.app.home.HomeRoute
import gg.foreversixty.app.network.AccountApi
import gg.foreversixty.app.raids.RaidsFeature
import gg.foreversixty.app.raids.RaidsRoute
import androidx.compose.runtime.LaunchedEffect

/** The app shell: a Material 3 navigation bar over one placeholder screen per tab. */
@Composable
fun ForeverAppScaffold(accountApi: AccountApi, raids: RaidsFeature? = null) {
    var selected by rememberSaveable { mutableStateOf(ForeverTab.Character) }
    // Lane F4: a raid link or notification selects the Raids tab.
    val wantsRaids = raids?.navigator?.wantsTab == true
    LaunchedEffect(wantsRaids) {
        if (wantsRaids) {
            selected = ForeverTab.Raids
            raids.navigator.wantsTab = false
        }
    }
    ForeverTabScaffold(selected = selected, onSelect = { selected = it }) {
        when (selected) {
            ForeverTab.Character -> HomeRoute(accountApi)
            ForeverTab.Raids -> if (raids != null && raids.enabled) RaidsRoute(raids) else ComingSoon(selected)
            ForeverTab.Bis, ForeverTab.Sims, ForeverTab.Account -> ComingSoon(selected)
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
                        label = { Text(stringResource(tab.label)) },
                    )
                }
            }
        },
    ) { padding ->
        Box(Modifier.fillMaxSize().padding(padding)) { content() }
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
