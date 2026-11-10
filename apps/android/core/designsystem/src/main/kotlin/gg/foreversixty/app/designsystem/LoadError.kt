package gg.foreversixty.app.designsystem

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Refresh
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.semantics.LiveRegionMode
import androidx.compose.ui.semantics.liveRegion
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.dp

/**
 * The one load error (spec 3.4, `ui/LoadError`): a Plate with the sentence in wipe, announced once,
 * and a secondary Try again. Wipe ink reads on the panel's dark, never loose on the painted ground.
 * iOS: `LoadError`.
 */
@Composable
fun LoadError(text: String, retryLabel: String, retry: () -> Unit, modifier: Modifier = Modifier) {
    Column(modifier.fillMaxWidth().panel().padding(16.dp), verticalArrangement = Arrangement.spacedBy(12.dp)) {
        Text(text, style = ScaledType.body, color = FsColor.wipe, modifier = Modifier.semantics { liveRegion = LiveRegionMode.Polite })
        BrandButton(retryLabel, BrandButtonKind.Secondary, retry, height = 48.dp, fill = false, leadingIcon = Icons.Filled.Refresh)
    }
}
