package gg.foreversixty.app.bis

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.FlowRow
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.navigationBarsPadding
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.selection.selectable
import androidx.compose.foundation.selection.selectableGroup
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.common.ClassCrest
import gg.foreversixty.app.common.FactionMark
import gg.foreversixty.app.common.MicroStyle
import gg.foreversixty.app.common.SecondaryButton
import gg.foreversixty.app.common.pillSurface
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.GamePalette
import gg.foreversixty.app.designsystem.HeadingFamily
import gg.foreversixty.app.items.SpecIcon

private const val CLASS_COLUMNS = 3

/**
 * `Choose a list` (spec F3.2): class, spec, level range and faction. Every choice applies at once
 * and the list redraws under the sheet; the sheet stays until dismissed.
 */
@Composable
fun BisPickerContent(context: BisContext, catalog: SpecCatalog, onSelect: (BisSelection) -> Unit) {
    val res = LocalResources.current
    val selection = context.selection
    Column(
        Modifier.background(GamePalette.sheet).verticalScroll(rememberScrollState()).navigationBarsPadding().padding(16.dp),
        verticalArrangement = Arrangement.spacedBy(22.dp),
    ) {
        Text(
            stringResource(R.string.bis_picker_title).uppercase(),
            style = MaterialTheme.typography.titleMedium.copy(fontFamily = HeadingFamily, fontWeight = FontWeight.ExtraBold),
            color = FsColor.strong,
            modifier = Modifier.semantics { heading() },
        )
        Section(stringResource(R.string.bis_picker_class)) { ClassGrid(context, catalog, onSelect) }
        Section(stringResource(R.string.bis_picker_spec)) {
            FlowRow(Modifier.selectableGroup(), horizontalArrangement = Arrangement.spacedBy(8.dp), verticalArrangement = Arrangement.spacedBy(8.dp)) {
                catalog.specsOf(selection.classSlug).forEach { spec ->
                    val current = spec.spec == selection.spec
                    val shape = RoundedCornerShape(FsRadius.control)
                    Row(
                        Modifier
                            .heightIn(min = 48.dp)
                            .pillSurface(fill(current), line(current), shape)
                            .selectable(selected = current, role = Role.RadioButton) { onSelect(selection.copy(spec = spec.spec)) }
                            .padding(horizontal = 10.dp, vertical = 6.dp),
                        horizontalArrangement = Arrangement.spacedBy(10.dp),
                        verticalAlignment = Alignment.CenterVertically,
                    ) {
                        SpecIcon(spec.icon, size = 32.dp)
                        Column {
                            Text(spec.name, style = MaterialTheme.typography.bodyLarge.copy(fontWeight = FontWeight.SemiBold), color = FsColor.strong)
                            Text(BisLabels.role(res, spec.role), style = MaterialTheme.typography.labelMedium.copy(fontWeight = FontWeight.Normal), color = FsColor.muted)
                        }
                    }
                }
            }
        }
        Section(stringResource(R.string.bis_picker_range)) {
            Pills(BisRange.all, { BisLabels.range(res, it) }, selection.range, mark = null) { onSelect(selection.copy(range = it)) }
        }
        Section(stringResource(R.string.bis_picker_faction)) {
            Pills(BisSelection.factions, { BisLabels.faction(res, it) }, selection.faction, mark = { it }) { onSelect(selection.copy(faction = it)) }
        }
        context.backTarget?.let { (name, own) ->
            SecondaryButton(stringResource(R.string.bis_picker_back, name), { onSelect(own) })
        }
    }
}

@Composable
private fun ClassGrid(context: BisContext, catalog: SpecCatalog, onSelect: (BisSelection) -> Unit) {
    val selection = context.selection
    val slugs = context.classes.map { it.slug }.ifEmpty { catalog.classSlugs }
    Column(Modifier.selectableGroup(), verticalArrangement = Arrangement.spacedBy(8.dp)) {
        slugs.chunked(CLASS_COLUMNS).forEach { row ->
            Row(horizontalArrangement = Arrangement.spacedBy(8.dp)) {
                row.forEach { slug ->
                    val current = slug == selection.classSlug
                    val shape = RoundedCornerShape(FsRadius.control)
                    Column(
                        Modifier
                            .weight(1f)
                            .heightIn(min = 88.dp)
                            .background(if (current) GamePalette.pressedWash else Color.Transparent, shape)
                            .border(1.dp, if (current) FsColor.gold else Color.Transparent, shape)
                            .selectable(selected = current, role = Role.RadioButton) { onSelect(selection.withClass(slug, catalog)) }
                            .padding(8.dp),
                        horizontalAlignment = Alignment.CenterHorizontally,
                        verticalArrangement = Arrangement.spacedBy(6.dp),
                    ) {
                        ClassCrest(slug, size = 44.dp)
                        Text(
                            context.className(slug),
                            style = MaterialTheme.typography.bodyLarge.copy(fontWeight = FontWeight.SemiBold),
                            color = GamePalette.classColor(slug),
                        )
                    }
                }
                repeat(CLASS_COLUMNS - row.size) { Column(Modifier.weight(1f)) {} }
            }
        }
    }
}

@Composable
private fun <T> Pills(options: List<T>, label: (T) -> String, current: T, mark: ((T) -> String)?, onPick: (T) -> Unit) {
    FlowRow(Modifier.selectableGroup(), horizontalArrangement = Arrangement.spacedBy(8.dp), verticalArrangement = Arrangement.spacedBy(8.dp)) {
        options.forEach { option ->
            val isCurrent = option == current
            Row(
                Modifier
                    .heightIn(min = 48.dp)
                    .pillSurface(fill(isCurrent), line(isCurrent))
                    .selectable(selected = isCurrent, role = Role.RadioButton) { onPick(option) }
                    .padding(horizontal = 16.dp, vertical = 8.dp),
                horizontalArrangement = Arrangement.spacedBy(8.dp),
                verticalAlignment = Alignment.CenterVertically,
            ) {
                mark?.let { FactionMark(it(option), size = 22.dp) }
                Text(
                    label(option),
                    style = MaterialTheme.typography.bodyLarge.copy(fontWeight = if (isCurrent) FontWeight.SemiBold else FontWeight.Normal),
                    color = if (isCurrent) FsColor.goldHover else FsColor.text,
                )
            }
        }
    }
}

private fun fill(current: Boolean) = if (current) GamePalette.pressedWash else GamePalette.controlFill

private fun line(current: Boolean) = if (current) FsColor.gold else FsColor.lineCard

@Composable
private fun Section(title: String, body: @Composable () -> Unit) {
    Column(verticalArrangement = Arrangement.spacedBy(10.dp)) {
        Text(title.uppercase(), style = MicroStyle, color = FsColor.muted, modifier = Modifier.semantics { heading() })
        body()
    }
}
