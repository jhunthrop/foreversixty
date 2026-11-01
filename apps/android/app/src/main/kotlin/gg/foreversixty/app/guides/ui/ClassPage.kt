package gg.foreversixty.app.guides.ui

import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.common.OfflineStamp
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.OfflineStrip
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.guides.logic.GuideDates
import gg.foreversixty.app.guides.logic.GuidesCopy
import gg.foreversixty.app.guides.logic.Loadable
import gg.foreversixty.app.items.SpecIcon
import gg.foreversixty.app.model.GuideIndex
import gg.foreversixty.app.model.GuideIndexSpec
import gg.foreversixty.app.reports.ui.Chevron
import gg.foreversixty.app.reports.ui.ReportsBar
import gg.foreversixty.app.reports.ui.isLargeText
import java.time.Instant

/** F9.6: a class's spec rows, each `Saved` or `Not saved`; pushed from a crest. */
@Composable
fun ClassPage(classSlug: String, index: Loadable<GuideIndex>, saved: Map<String, Long>, now: Instant, back: () -> Unit, open: (String) -> Unit) {
    val loaded = index as? Loadable.Loaded
    val cls = loaded?.value?.classes?.firstOrNull { it.classSlug == classSlug }
    Column(Modifier.fillMaxWidth()) {
        ReportsBar(cls?.className.orEmpty(), stringResource(R.string.guides_title), back, titleColor = classColor(classSlug))
        loaded?.offlineSince?.let { OfflineStrip(stringResource(R.string.offline_strip, OfflineStamp.format(it, now))) }
        if (cls != null) {
            Column(Modifier.screenMargins().padding(top = GuideSpace.margin), verticalArrangement = Arrangement.spacedBy(GuideSpace.rowGap)) {
                QuietRows(cls.specs) { spec -> SpecRow(spec, loaded.value.build, spec.spec in saved) { open(spec.spec) } }
                // Offline with a guide not on the phone: say so under the card (board F9.6).
                if (loaded.offlineSince != null && cls.specs.any { it.spec !in saved }) {
                    Text(stringResource(R.string.guides_notCached), style = FsType.meta, color = FsColor.muted)
                }
            }
        }
    }
}

@Composable
private fun SpecRow(spec: GuideIndexSpec, build: String, saved: Boolean, open: () -> Unit) {
    val copy = GuidesCopy(LocalResources.current)
    val words = @Composable { modifier: Modifier ->
        Column(modifier, verticalArrangement = Arrangement.spacedBy(2.dp)) {
            Text(spec.specName, style = FsType.bodyStrong, color = FsColor.strong)
            Text(copy.cardMeta(spec.role, GuideDates.short(spec.updated)), style = FsType.meta, color = FsColor.muted)
        }
    }
    Row(
        Modifier.fillMaxWidth().heightIn(min = ROW_MIN).clickable(role = Role.Button, onClick = open).cardInside(),
        horizontalArrangement = Arrangement.spacedBy(GuideSpace.rowGap),
        verticalAlignment = Alignment.CenterVertically,
    ) {
        SpecIcon(spec.icon, build, ICON)
        if (isLargeText()) {
            // Largest text: the saved mark moves under the name, so neither is cut.
            Column(Modifier.weight(1f), verticalArrangement = Arrangement.spacedBy(6.dp)) {
                words(Modifier)
                SavedMark(saved)
            }
        } else {
            words(Modifier.weight(1f))
            SavedMark(saved)
        }
        Chevron(CHEVRON)
    }
}

private val ROW_MIN = 68.dp
private val ICON = 40.dp
private val CHEVRON = 20.dp
