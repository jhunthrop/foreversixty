package gg.foreversixty.app.sims.logic

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.bis.SpecCatalog
import gg.foreversixty.app.items.ItemVocabulary
import gg.foreversixty.app.model.SimResult
import gg.foreversixty.app.model.SimRow
import gg.foreversixty.app.raids.logic.GameClass
import java.util.Locale

/** One icon of `The gear it ran`, in the character frame's slot order. */
data class SimGearCell(
    val slot: String,
    val itemId: Int,
    /** A slot the Top Gear result's best set changes (the gold corner). */
    val changed: Boolean,
)

/** An item of the gear grid, as the build's items file names it. */
data class SimGearLook(val name: String, val icon: String, val quality: Int)

/**
 * What F5.3 draws for one result (spec §5.5 F5.3). Reads the result and, when it has one, the
 * list row for its headline; nothing here re-derives a number the engine wrote.
 */
data class SimResultModel(
    val title: String,
    val kind: String,
    val specName: String,
    val className: String,
    val classSlug: String,
    val level: Int?,
    val answer: Answer,
    val errorLine: String,
    val gear: List<SimGearCell>,
    val breakdownUrl: String,
) {
    sealed interface Answer {
        /** A plain run: the mean to one decimal, `DPS` beside it. */
        data class Figure(val value: String) : Answer

        /** Every other kind: the row's headline sentence. */
        data class Sentence(val text: String) : Answer
    }

    /** `sims.changed` under the grid, when a slot carries the corner. */
    val showsChangedLegend: Boolean get() = gear.any { it.changed }

    companion object {
        private const val SITE = "https://foreversixty.gg"

        fun of(id: String, result: SimResult, row: SimRow?, specs: SpecCatalog, res: Resources): SimResultModel {
            val specKey = result.request?.spec ?: row?.spec.orEmpty()
            val spec = specs.spec(specKey)
            val character = result.request?.character
            val slug = character?.className ?: spec?.classSlug.orEmpty()
            val specName = spec?.name ?: specKey
            val named = result.title ?: row?.title.orEmpty()
            val kindKey = row?.kind ?: kindOf(result)
            val headline = row?.headline.orEmpty()
            val sentence = kindKey != "run" || result.healing != null || result.tank != null
            return SimResultModel(
                title = named.ifEmpty { specName },
                kind = SimsCopy(res).kindLabel(kindKey),
                specName = specName,
                className = GameClass.of(slug)?.let { res.getString(it.label) } ?: slug.replaceFirstChar { it.titlecase(Locale.US) },
                classSlug = slug,
                level = character?.level,
                answer = if (sentence && headline.isNotEmpty()) Answer.Sentence(headline) else Answer.Figure(figure(result.dps.mean)),
                errorLine = res.getString(
                    R.string.sims_error, figure(result.dps.error), String.format(Locale.US, "%,d", result.iterationsRun), result.engineVersion,
                ),
                gear = gear(result, kindKey),
                breakdownUrl = "$SITE/sim/$id",
            )
        }

        fun figure(value: Double?): String = String.format(Locale.US, "%,.1f", value ?: 0.0)

        /** A result names its kind by its blocks, as the site's `kindOf`. */
        fun kindOf(result: SimResult): String = when {
            result.weights != null -> "weights"
            else -> result.request?.bulk?.mode ?: "run"
        }

        /** The gear the run wore, in frame order; on Top Gear the slots its best set changes are marked. */
        fun gear(result: SimResult, kind: String): List<SimGearCell> {
            val base = result.request?.character?.gear.orEmpty().associate { it.slot to it.itemId }
            val changed = if (kind == "gear") result.combos?.firstOrNull()?.substitutions.orEmpty().mapNotNull { it.slot }.toSet() else emptySet()
            return ItemVocabulary.frameSlots.mapNotNull { slot -> base[slot]?.let { SimGearCell(slot, it, slot in changed) } }
        }
    }

    fun meta(res: Resources): String = listOfNotNull("$specName $className", level?.let { res.getString(R.string.sims_level, it) }).joinToString(" · ")
}
