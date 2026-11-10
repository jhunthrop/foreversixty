package gg.foreversixty.app.designsystem

import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.BoxScope
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.RowScope
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.WindowInsets
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.filled.ArrowBack
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.Text
import androidx.compose.material3.TopAppBar
import androidx.compose.material3.TopAppBarDefaults
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.style.TextOverflow

/**
 * The 12 sp label (`FsType.label`): 700, muted, sentence case (spec ruling N1). A field's name above
 * its value, never a section's title ([SectionTitle]) and never a button. iOS: `Label12`.
 */
@Composable
fun Label12(text: String, modifier: Modifier = Modifier, color: Color = FsColor.muted) {
    Text(text, modifier, style = FsType.label, color = color)
}

/**
 * A section's title over the ground in the round-5 look (the mocks' `.sh h2`): the display face at
 * 21, cream, read as a heading. The tab roots' `FlowSection` uses the same role. iOS: `SectionTitle`.
 */
@Composable
fun SectionTitle(text: String, modifier: Modifier = Modifier, color: Color = FsColor.cream) {
    Text(text, modifier.semantics { heading() }, style = ScaledType.sectionTitle, color = color)
}

/**
 * A pushed screen's top app bar in the round-5 look (spec Mobile, Android): the painted band behind
 * it with its black bottom edge, Back at the start, [title] in the display face at 22 in sentence
 * or title case as written, [actions] at the end. Material 3's bar carries the title, so the page
 * never repeats it. iOS: the navigation bar over `BandArt`.
 */
@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun FsTopBar(
    title: String?,
    backLabel: String,
    back: () -> Unit,
    titleColor: Color = FsColor.strong,
    /** False when the page's head continues the band under the bar (the guide's crest and name, as on the web). */
    bottomLine: Boolean = true,
    actions: @Composable RowScope.() -> Unit = {},
) {
    TopAppBar(
        title = {
            if (title != null) {
                Text(
                    title, style = ScaledType.barTitle, color = titleColor, maxLines = 2, overflow = TextOverflow.Ellipsis,
                    modifier = Modifier.semantics { heading() },
                )
            }
        },
        navigationIcon = {
            IconButton(onClick = back) { Icon(Icons.AutoMirrored.Filled.ArrowBack, contentDescription = backLabel, tint = FsColor.strong) }
        },
        actions = actions,
        colors = TopAppBarDefaults.topAppBarColors(containerColor = Color.Transparent, scrolledContainerColor = Color.Transparent),
        // The app scaffold already pads for the status bar.
        windowInsets = WindowInsets(0, 0, 0, 0),
        modifier = Modifier.surfaceBand(bottomLine),
    )
}

/**
 * A pushed screen in the round-5 look: the painted ground, [FsTopBar] with [title], then [content]
 * filling the rest (it scrolls itself). Routes and their screenshot tests draw the same chrome.
 */
@Composable
fun PushedScreen(
    title: String?,
    backLabel: String,
    back: () -> Unit,
    titleColor: Color = FsColor.strong,
    actions: @Composable RowScope.() -> Unit = {},
    content: @Composable BoxScope.() -> Unit,
) {
    Column(Modifier.fillMaxSize().screenGround()) {
        FsTopBar(title, backLabel, back, titleColor, actions = actions)
        Box(Modifier.fillMaxWidth().weight(1f), content = content)
    }
}
