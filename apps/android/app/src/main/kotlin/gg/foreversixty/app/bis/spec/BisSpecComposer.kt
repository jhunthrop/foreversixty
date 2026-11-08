package gg.foreversixty.app.bis.spec

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.bis.BisLabels
import gg.foreversixty.app.bis.BisListBuilder
import gg.foreversixty.app.bis.BisRange
import gg.foreversixty.app.bis.BisSelection
import gg.foreversixty.app.bis.BisUnit
import gg.foreversixty.app.bis.SpecCatalog
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.items.ItemVocabulary
import gg.foreversixty.app.model.BisBand
import gg.foreversixty.app.model.BisFile
import gg.foreversixty.app.model.BisSlot
import gg.foreversixty.app.model.Item
import java.time.Instant
import java.time.ZoneId
import java.time.format.DateTimeFormatter
import java.time.format.DateTimeParseException
import java.util.Locale
import kotlin.math.abs

/** What the composer reads: the published file and band, the build's items, the list asked for and the sources left on. */
data class BisSpecInput(
    val file: BisFile,
    val band: BisBand,
    val items: List<Item>,
    val selection: BisSelection,
    val className: String,
    val sourcesOff: Set<BisSource> = emptySet(),
)

/**
 * Words a published band as the spec page (mocks `bis-spec-*`). Nothing is re-ranked: with every source on
 * each slot shows the file's pick; a source turned off falls to the slot's first runner-up still on.
 */
class BisSpecComposer(private val res: Resources, private val specs: SpecCatalog, private val zone: ZoneId = ZoneId.systemDefault()) {
    fun content(input: BisSpecInput): BisSpecContent {
        val words = BisListBuilder(res, input.items, emptySet(), zone)
        val unit = BisUnit.forRole(input.band.role)
        val items = input.items.associateBy { it.id }
        val picks = input.band.slots.map { pick(it, input.sourcesOff) }
        val cells = picks.map { cell(it, input, items, words, unit) }
        val specName = specs.spec(input.selection.spec)?.name ?: input.selection.spec
        return BisSpecContent(
            title = res.getString(R.string.bis_specChip, specName, input.className),
            classSlug = input.selection.classSlug,
            meta = res.getString(R.string.bisSpec_meta),
            answer = answer(input, unit),
            controls = controls(input, specName),
            summary = summary(input, picks.count { it.itemId != null }),
            notes = notes(input.band, items),
            cells = cells,
            groups = groups(picks.zip(cells)),
            weights = weights(input.band, unit),
            guide = res.getString(R.string.bisSpec_guideLink, res.getString(R.string.bis_specChip, specName, input.className)),
        )
    }

    /** The sources a band's picks and runners-up come from, for the Sources sheet. */
    fun sources(band: BisBand): List<BisSource> =
        band.slots.flatMap { slot -> listOf(slot.sourceKind) + slot.alternatives.orEmpty().map { it.sourceKind } }
            .map(BisSource::of).distinct().sortedBy { it.ordinal }

    private fun answer(input: BisSpecInput, unit: BisUnit): BisSpecAnswer {
        val band = input.band
        val level = if (band.band == BisRange.LAST_BAND) band.band else band.band + BisRange.BAND_STEP - 1
        return BisSpecAnswer(
            label = res.getString(R.string.bisSpec_simmed, simmed(input.file.generatedAt)),
            value = band.setDps?.takeIf { unit != BisUnit.Score }?.let(BisListBuilder::figure),
            unit = res.getString(unit.word),
            facts = emptyList(),
            line = res.getString(R.string.bisSpec_line, level, race(band.race), band.talentPoints, profile(input.file, band.preset)),
            rankingSim = null,
            openSim = band.role == null || band.role == DPS,
        )
    }

    private fun controls(input: BisSpecInput, specName: String): BisSpecControls {
        val selection = input.selection
        val range = selection.range
        val faction = BisLabels.faction(res, selection.faction)
        val bandWords = bandLabel(range.band)
        val atTop = range.band == BisRange.LAST_BAND
        val profile = profileWord(range.preset)
        val bands = (BisRange.FIRST_BAND..BisRange.LAST_BAND step BisRange.BAND_STEP).toList()
        return BisSpecControls(
            bar = if (atTop) res.getString(R.string.bisSpec_bar, specName, bandWords, faction, profile) else res.getString(R.string.bisSpec_barBare, specName, bandWords, faction),
            specs = specs.specsOf(selection.classSlug).map { BisSpecOption(it.name, it.spec == selection.spec, it.icon, selection.copy(spec = it.spec)) },
            levels = bands.map { b ->
                val preset = if (b == BisRange.LAST_BAND) BisRange.RAID else BisRange.BARE
                BisSpecOption(bandLabel(b), b == range.band, pick = selection.copy(range = BisRange(b, preset)))
            },
            factions = BisSelection.factions.map { BisSpecOption(BisLabels.faction(res, it), it == selection.faction, pick = selection.copy(faction = it)) },
            profiles = if (!atTop) emptyList() else listOf(BisRange.RAID, BisRange.BARE).map {
                BisSpecOption(profileWord(it), it == range.preset, pick = selection.copy(range = BisRange(range.band, it)))
            },
            race = race(input.band.race),
            sources = if (input.sourcesOff.isEmpty()) {
                res.getString(R.string.bisSpec_sources_all)
            } else {
                res.getString(R.string.bisSpec_sources_some, sources(input.band).count { it !in input.sourcesOff })
            },
        )
    }

    private fun summary(input: BisSpecInput, picked: Int): BisSpecSummary {
        val range = input.selection.range
        val count = res.getString(R.string.bisSpec_summary, picked, input.band.slots.size)
        if (range.band >= BisRange.LAST_BAND) return BisSpecSummary(count, null, null)
        val next = range.band + BisRange.BAND_STEP
        val nextRange = BisRange(next, if (next == BisRange.LAST_BAND) BisRange.RAID else BisRange.BARE)
        val words = res.getString(R.string.bisSpec_nextBand, BisLabels.band(res, range), BisLabels.band(res, nextRange))
        return BisSpecSummary(count, words, input.selection.copy(range = nextRange))
    }

    private fun notes(band: BisBand, items: Map<Int, Item>): List<String> {
        val mainHand = band.slots.firstOrNull { it.slot == BisListBuilder.MAIN_HAND }
        val twoHanded = mainHand?.itemId?.let { items[it]?.twoHand } == true
        return listOfNotNull(mainHand?.itemName?.takeIf { twoHanded }?.let { res.getString(R.string.bisSpec_twoHander, it) })
    }

    private fun cell(pick: Pick, input: BisSpecInput, items: Map<Int, Item>, words: BisListBuilder, unit: BisUnit): BisSpecCell {
        val slotLabel = ItemVocabulary.slotLabel(res, pick.slot.slot, input.selection.classSlug)
        val id = pick.itemId
        if (id == null) {
            val mainHand = input.band.slots.firstOrNull { it.slot == BisListBuilder.MAIN_HAND }?.itemName
            val reason = if (pick.slot.emptyReason == BisListBuilder.TWO_HAND_EQUIPPED && mainHand != null) {
                res.getString(R.string.bisSpec_twoHanderSlot, mainHand)
            } else {
                res.getString(R.string.bis_noPick, slotLabel)
            }
            return BisSpecCell(slotLabel, null, null, "", reason, null, null, null, null)
        }
        val item = items[id]
        val runnerUp = pick.runnerUp
        val newGain = runnerUp?.takeIf { pick.slot.slot in newSlots(input.band) }?.let { "+" + BisListBuilder.figure(abs(it.delta)) }
        val isNew = newGain != null
        val gap = runnerUp?.let { BisListBuilder.signedDelta(it.delta) }
        return BisSpecCell(
            slot = slotLabel,
            level = item?.requiredLevel?.takeIf { it > 1 }?.let { res.getString(R.string.bisSpec_itemLevel, it) },
            look = ItemLook(id, pick.name, item?.quality, item?.icon, item?.itemLevel),
            stats = stats(item),
            source = pick.sourceLine ?: words.sourceLine(pick.source, pick.kind),
            gain = newGain ?: res.getString(R.string.bisSpec_same),
            unit = if (isNew) res.getString(unit.word) else null,
            versus = when {
                runnerUp == null -> res.getString(R.string.bisSpec_noAlternative)
                isNew || gap == null -> res.getString(R.string.bisSpec_vs)
                else -> res.getString(R.string.bisSpec_nextVs, gap)
            },
            versusItem = runnerUp?.let { ItemLook(it.id, it.name, items[it.id]?.quality, items[it.id]?.icon) },
        )
    }

    private fun groups(cells: List<Pair<Pick, BisSpecCell>>): List<BisSpecGroup> =
        cells.filter { it.first.itemId != null }.groupBy { place(it.first) }.map { (place, members) ->
            BisSpecGroup(place, res.getQuantityString(R.plurals.bisSpec_pieces, members.size, members.size), members.map { it.second })
        }

    /** Where a pick is got: the instance of a drop, the profession, the reputation, or the kind. */
    private fun place(pick: Pick): String {
        val line = pick.sourceLine.orEmpty()
        return when (BisSource.of(pick.kind)) {
            BisSource.Raid, BisSource.Dungeon -> line.split(SEP).getOrNull(1)?.substringBefore(" (")?.takeIf { it.isNotBlank() } ?: pick.source.substringBefore(":")
            BisSource.Crafted -> res.getString(R.string.bisSpec_place_crafted, pick.source)
            BisSource.Reputation -> line.substringBefore(SEP).ifBlank { pick.source }
            BisSource.Quest -> res.getString(R.string.bisSpec_place_quest)
            BisSource.World -> res.getString(R.string.bisSpec_place_world)
            BisSource.Vendor -> res.getString(R.string.bisSpec_place_vendor)
            BisSource.Pvp, BisSource.Other -> pick.source
        }
    }

    private fun weights(band: BisBand, unit: BisUnit): BisSpecWeights? {
        val ranked = band.weights.sortedByDescending { it.scaleFactor }
        val significant = ranked.filter { !it.insignificant }
        val top = significant.firstOrNull() ?: return null
        val names = significant.take(STACK_NAMES).map { statName(it.stat).lowercase(Locale.US) }
        val sentence = when (names.size) {
            1 -> res.getString(R.string.bisSpec_stack_two, names[0], names[0])
            2 -> res.getString(R.string.bisSpec_stack_two, names[0], names[1])
            else -> res.getString(R.string.bisSpec_stack_three, names[0], names[1], names[2])
        }
        val unitWord = res.getString(unit.word)
        return BisSpecWeights(
            sentence = sentence,
            legend = res.getString(R.string.bisSpec_stack_legend, statName(top.stat), unitWord),
            stats = (significant + ranked.filter { it.insignificant }).map {
                if (it.insignificant) {
                    BisSpecWeight(statName(it.stat), null, 0f, null)
                } else {
                    BisSpecWeight(
                        statName(it.stat), String.format(Locale.US, "%.2f", it.scaleFactor), it.scaleFactor.toFloat(),
                        res.getString(R.string.bisSpec_stack_perPoint, String.format(Locale.US, "%.3f", it.dpsPerPoint), unitWord),
                    )
                }
            },
        )
    }

    private fun stats(item: Item?): String {
        item ?: return ""
        return ItemVocabulary.statOrder.mapNotNull { (key, label) -> item.stats[key]?.takeIf { it != 0 }?.let { "+$it ${res.getString(label)}" } }
            .take(STATS_SHOWN).joinToString(SEP)
    }

    private fun statName(stat: String): String =
        ItemVocabulary.statOrder.firstOrNull { it.first == stat }?.let { res.getString(it.second) } ?: stat.replace('_', ' ').replaceFirstChar { it.titlecase(Locale.US) }

    private fun newSlots(band: BisBand): Set<String> = band.newAtBand.orEmpty().map { it.substringBefore(":").trim() }.toSet()

    private fun bandLabel(band: Int): String =
        if (band == BisRange.LAST_BAND) band.toString() else "$band-${band + BisRange.BAND_STEP - 1}"

    private fun profileWord(preset: String): String =
        res.getString(if (preset == BisRange.RAID) R.string.bisSpec_profile_raid else R.string.bisSpec_profile_bare)

    private fun profile(file: BisFile, preset: String): String = file.presets[preset]?.label ?: profileWord(preset)

    private fun race(slug: String): String = slug.replaceFirstChar { it.titlecase(Locale.US) }

    private fun simmed(generatedAt: String): String = try {
        SIMMED.withZone(zone).format(Instant.parse(generatedAt)).lowercase(Locale.US)
    } catch (e: DateTimeParseException) {
        generatedAt
    }

    /** A slot's pick after the sources filter: the file's pick, or its first runner-up from a source still on. */
    private fun pick(slot: BisSlot, off: Set<BisSource>): Pick {
        val candidates = listOfNotNull(
            slot.itemId?.takeIf { it > 0 }?.let { Candidate(it, slot.itemName.orEmpty(), slot.source.orEmpty(), slot.sourceKind, slot.sourceLine, 0.0) },
        ) + slot.alternatives.orEmpty().map { Candidate(it.itemId, it.itemName, it.source, it.sourceKind, it.sourceLine, it.dpsDelta) }
        val usable = candidates.filter { BisSource.of(it.kind) !in off }
        val chosen = usable.firstOrNull()
        val next = usable.drop(1).firstOrNull()
        return Pick(
            slot = slot,
            itemId = chosen?.id,
            name = chosen?.name.orEmpty(),
            source = chosen?.source.orEmpty(),
            kind = chosen?.kind,
            sourceLine = chosen?.line,
            runnerUp = if (chosen != null && next != null) RunnerUp(next.id, next.name, next.delta - chosen.delta) else null,
        )
    }

    private data class Candidate(val id: Int, val name: String, val source: String, val kind: String?, val line: String?, val delta: Double)

    private data class RunnerUp(val id: Int, val name: String, val delta: Double)

    private data class Pick(
        val slot: BisSlot,
        val itemId: Int?,
        val name: String,
        val source: String,
        val kind: String?,
        val sourceLine: String?,
        val runnerUp: RunnerUp?,
    )

    private companion object {
        const val DPS = "dps"
        const val SEP = " · "
        const val STACK_NAMES = 3
        const val STATS_SHOWN = 2
        val SIMMED: DateTimeFormatter = DateTimeFormatter.ofPattern("d MMM", Locale.US)
    }
}

/** The source kinds the Sources sheet turns on and off, in its order. */
enum class BisSource(val label: Int) {
    Raid(R.string.bisSpec_sources_kind_raid),
    Dungeon(R.string.bisSpec_sources_kind_dungeon),
    Crafted(R.string.bisSpec_sources_kind_crafted),
    Reputation(R.string.bisSpec_sources_kind_reputation),
    Quest(R.string.bisSpec_sources_kind_quest),
    Vendor(R.string.bisSpec_sources_kind_vendor),
    World(R.string.bisSpec_sources_kind_world),
    Pvp(R.string.bisSpec_sources_kind_pvp),
    Other(R.string.bisSpec_sources_kind_other);

    companion object {
        fun of(kind: String?): BisSource = when (kind) {
            "raid" -> Raid
            "zone", "dungeon" -> Dungeon
            "crafted" -> Crafted
            "rep", "reputation" -> Reputation
            "quest" -> Quest
            "vendor" -> Vendor
            "world", "world_drop" -> World
            "pvp" -> Pvp
            else -> Other
        }
    }
}
