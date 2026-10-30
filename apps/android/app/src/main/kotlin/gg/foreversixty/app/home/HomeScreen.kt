package gg.foreversixty.app.home

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.produceState
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.model.mainCharacter
import gg.foreversixty.app.network.AccountApi
import gg.foreversixty.app.network.dataOrNull

/** Loads /v1/me and shows the main character. Placeholder until F2 builds the real home. */
@Composable
fun HomeRoute(accountApi: AccountApi) {
    val character by produceState<AccountCharacter?>(initialValue = null, accountApi) {
        value = accountApi.me().dataOrNull()?.mainCharacter
    }
    HomeScreen(character)
}

/** The placeholder home: the tagline over your main character, stateless for previews and tests. */
@Composable
fun HomeScreen(character: AccountCharacter?) {
    Column(
        modifier = Modifier.fillMaxSize().padding(horizontal = 16.dp, vertical = 24.dp),
        verticalArrangement = Arrangement.spacedBy(8.dp),
    ) {
        Text(
            text = stringResource(R.string.shell_home_tagline),
            style = MaterialTheme.typography.labelMedium,
            color = MaterialTheme.colorScheme.primary,
        )
        if (character == null) {
            Text(
                text = stringResource(R.string.shell_home_noCharacter),
                style = MaterialTheme.typography.bodyLarge,
                color = MaterialTheme.colorScheme.onSurfaceVariant,
            )
        } else {
            CharacterHeader(character)
        }
    }
}

@Composable
private fun CharacterHeader(character: AccountCharacter) {
    Text(
        text = character.name,
        style = MaterialTheme.typography.headlineMedium,
        color = MaterialTheme.colorScheme.onBackground,
    )
    val level = character.level
    val characterClass = character.className
    if (level != null && characterClass != null) {
        Text(
            text = stringResource(R.string.shell_home_levelClass, level, characterClass.replaceFirstChar { it.uppercase() }),
            style = MaterialTheme.typography.bodyLarge,
            color = MaterialTheme.colorScheme.onSurfaceVariant,
        )
    }
}
