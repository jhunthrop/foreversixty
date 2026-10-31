package gg.foreversixty.app.reports.ui

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.WindowInsets
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.filled.ArrowBack
import androidx.compose.material.icons.filled.Refresh
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.Text
import androidx.compose.material3.TopAppBar
import androidx.compose.material3.TopAppBarDefaults
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.painter.Painter
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.LiveRegionMode
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.liveRegion
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.style.TextOverflow
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.BrandButton
import gg.foreversixty.app.designsystem.BrandButtonKind
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.GameColors
import gg.foreversixty.app.reports.logic.ParseTier
import gg.foreversixty.app.reports.logic.Visibility
import java.util.Locale

/** A pushed screen's bar: back with its label, the uppercase title, and one trailing action. */
@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun ReportsBar(title: String, backLabel: String, back: () -> Unit, trailing: BarAction? = null) {
    TopAppBar(
        title = {
            Text(
                title.uppercase(Locale.ROOT), style = FsType.h2, color = FsColor.strong, maxLines = 2, overflow = TextOverflow.Clip,
                modifier = Modifier.semantics { heading() },
            )
        },
        navigationIcon = {
            IconButton(onClick = back) { Icon(Icons.AutoMirrored.Filled.ArrowBack, contentDescription = backLabel, tint = FsColor.strong) }
        },
        actions = {
            trailing?.let { action ->
                IconButton(onClick = action.run) { Icon(action.icon(), contentDescription = action.label, tint = FsColor.strong) }
            }
        },
        colors = TopAppBarDefaults.topAppBarColors(containerColor = FsColor.bg),
        // The app scaffold already pads for the status bar.
        windowInsets = WindowInsets(0, 0, 0, 0),
    )
}

/** A bar's trailing icon button. */
class BarAction(val label: String, val icon: @Composable () -> Painter, val run: () -> Unit) {
    companion object {
        @Composable
        fun refresh(run: () -> Unit) = BarAction(stringResource(R.string.reports_refresh), { androidx.compose.ui.graphics.vector.rememberVectorPainter(Icons.Filled.Refresh) }, run)

        @Composable
        fun openWeb(run: () -> Unit) = BarAction(stringResource(R.string.report_openWeb), { painterResource(R.drawable.ic_open_external) }, run)
    }
}

/** The failing region's sentence in wipe red with Try again (v1 §5.0). */
@Composable
fun ErrorRegion(text: String, retry: () -> Unit, inset: Boolean = true) {
    Column(Modifier.fillMaxWidth().padding(if (inset) 16.dp else 0.dp),verticalArrangement = Arrangement.spacedBy(12.dp)) {
        Text(text, style = FsType.lead, color = FsColor.wipe, modifier = Modifier.semantics { liveRegion = LiveRegionMode.Polite })
        BrandButton(stringResource(R.string.common_tryAgain), BrandButtonKind.Secondary, retry, height = 48.dp, fill = false, leadingIcon = Icons.Filled.Refresh)
    }
}

@Composable
fun SkeletonBar(width: Dp, height: Dp, modifier: Modifier = Modifier) {
    Box(modifier.width(width).height(height).background(FsColor.cardTop, RoundedCornerShape(4.dp)))
}

@Composable
fun SkeletonDot(size: Dp) {
    Box(Modifier.size(size).background(FsColor.cardTop, CircleShape))
}

/** The outcome dot: kill green, wipe red, trash muted. */
@Composable
fun OutcomeDot(color: Color) {
    Box(Modifier.size(10.dp).background(color, CircleShape))
}

/** The muted 16 dp visibility glyph, with its word for TalkBack. */
@Composable
fun VisibilityGlyph(visibility: Visibility) {
    val word = stringResource(visibility.word)
    val icon = when (visibility) {
        Visibility.Public -> R.drawable.ic_visibility_public
        Visibility.Unlisted -> R.drawable.ic_visibility_unlisted
        Visibility.Guild -> R.drawable.ic_visibility_guild
        Visibility.Private -> R.drawable.ic_visibility_private
    }
    Icon(painterResource(icon), contentDescription = null, tint = FsColor.muted, modifier = Modifier.size(16.dp).semantics { contentDescription = word })
}

/** The parse number's colour: the text steps of `--color-parse-*` (spec W3 §4). */
fun parseColor(tier: ParseTier): Color = when (tier) {
    ParseTier.Grey -> FsColor.parseGreyText
    ParseTier.Green -> FsColor.parseGreen
    ParseTier.Blue -> FsColor.rarityRareText
    ParseTier.Purple -> FsColor.rarityEpicText
    ParseTier.Orange -> FsColor.parseOrange
    ParseTier.Pink -> GameColors.parsePinkText
}
