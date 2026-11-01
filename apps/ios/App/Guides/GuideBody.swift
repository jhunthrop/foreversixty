import DesignSystem
import Models
import SwiftUI

/// The guide under the strip: the at-a-glance card, then each section the guide has, then Sources.
/// Each section reports its top so the strip can follow the scroll, and carries an anchor one
/// strip-height above it so a tap lands the section's label under the strip, not behind it.
struct GuideBody: View {
    let content: GuideContent
    let space: String
    var actions = GuideActions()
    var onTop: (GuideSection, CGFloat) -> Void = { _, _ in }

    private var guide: GuideFile { content.guide }

    var body: some View {
        VStack(alignment: .leading, spacing: 0) {
            GuideGlance(guide: guide).padding(.top, 14)
            ForEach(GuideSection.of(guide)) { section in
                self.section(section)
                    .background(alignment: .top) {
                        Color.clear.frame(height: 1).offset(y: -GuidesLayout.stripHeight).id(section)
                    }
                    .onGeometryChange(for: CGFloat.self) { $0.frame(in: .named(space)).minY } action: { onTop(section, $0) }
            }
            GuideSources(sources: guide.sources).padding(.top, GuidesLayout.sectionGap)
        }
        .padding(.horizontal, GuidesLayout.gutter)
    }

    @ViewBuilder
    private func section(_ section: GuideSection) -> some View {
        VStack(alignment: .leading, spacing: 0) {
            GuideSectionLabel(text: section.title)
            switch section {
            case .overview: GuideOverviewCard(guide: guide)
            case .talents: GuideTalentsSection(talents: guide.talents, build: content.build)
            case .rotation: GuideRotationSection(rotation: guide.rotation, build: content.build)
            case .stats: GuideStatsCard(guide: guide)
            case .gear: GuideGearRow(guide: guide, open: actions.openGear)
            case .races: GuideRacesCard(races: guide.races)
            }
        }
    }
}

/// The at-a-glance card (the one `--surface-card`): Role and its line, the numbered stat priority.
struct GuideGlance: View {
    let guide: GuideFile

    var body: some View {
        VStack(alignment: .leading, spacing: 8) {
            MicroLabel(Copy.guidesRole)
            Text(guide.roleLine ?? GuideRole.label(guide.role))
                .font(ScaledType.lead).foregroundStyle(FSColor.strong)
                .fixedSize(horizontal: false, vertical: true)
            MicroLabel(Copy.guidesStatPriority).padding(.top, 6)
            StatChips(chips: GuideCopy.statChips(guide))
        }
        .padding(GuidesLayout.cardPadding)
        .frame(maxWidth: .infinity, alignment: .leading)
        .primaryCard()
        .accessibilityElement(children: .contain)
        .accessibilityLabel(Copy.guidesGlance)
    }
}

struct GuideOverviewCard: View {
    let guide: GuideFile

    var body: some View {
        VStack(alignment: .leading, spacing: 14) {
            Text(guide.summary).font(ScaledType.body).foregroundStyle(FSColor.text)
                .fixedSize(horizontal: false, vertical: true)
            GuideLines(title: Copy.guidesStrengths, lines: guide.overview.strengths)
            GuideLines(title: Copy.guidesWeaknesses, lines: guide.overview.weaknesses)
            GuideLines(title: Copy.guidesInRaid, lines: guide.overview.raid)
        }
        .padding(GuidesLayout.cardPadding)
        .frame(maxWidth: .infinity, alignment: .leading)
        .quietCard()
    }
}

/// The key talents and the build rows (F9.3). The planner link is lane F13's.
struct GuideTalentsSection: View {
    let talents: GuideTalents
    let build: String

    var body: some View {
        VStack(alignment: .leading, spacing: 12) {
            GuideCardRows(items: talents.key.map(IdentifiedTalent.init)) { GuideTalentRow(talent: $0.talent, build: build) }
            GuideCardRows(items: GuideBuildRows.rows(talents)) { GuideBuildRowView(row: $0) }
        }
    }
}

private struct IdentifiedTalent: Identifiable {
    let talent: GuideKeyTalent
    var id: String { talent.name }
}

struct GuideTalentRow: View {
    let talent: GuideKeyTalent
    let build: String

    var body: some View {
        HStack(alignment: .top, spacing: 12) {
            SpecIcon(icon: talent.icon, build: build, size: GuidesLayout.rowIcon)
            VStack(alignment: .leading, spacing: 4) {
                ViewThatFits(in: .horizontal) {
                    HStack(alignment: .firstTextBaseline, spacing: 8) { name; ranks }
                    VStack(alignment: .leading, spacing: 2) { name; ranks }
                }
                if let text = talent.text {
                    Text(text).font(ScaledType.meta).foregroundStyle(FSColor.nav).fixedSize(horizontal: false, vertical: true)
                }
                if let why = talent.why {
                    Text(why).font(ScaledType.meta).foregroundStyle(FSColor.text).fixedSize(horizontal: false, vertical: true)
                }
            }
            .frame(maxWidth: .infinity, alignment: .leading)
        }
        .padding(GuidesLayout.cardPadding)
        .accessibilityElement(children: .combine)
    }

    private var name: some View {
        Text(talent.name).font(ScaledType.name).foregroundStyle(FSColor.strong).fixedSize(horizontal: false, vertical: true)
    }

    @ViewBuilder
    private var ranks: some View {
        if let ranks = GuideCopy.ranks(talent) {
            Text(ranks).font(ScaledType.monoMeta).foregroundStyle(FSColor.muted)
        }
    }
}

struct GuideBuildRowView: View {
    let row: GuideBuildRow

    var body: some View {
        HStack(alignment: .firstTextBaseline, spacing: 12) {
            VStack(alignment: .leading, spacing: 2) {
                Text(row.title).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                Text(Copy.guidesBuildLevel).font(ScaledType.meta).foregroundStyle(FSColor.muted)
            }
            .frame(maxWidth: .infinity, alignment: .leading)
            Text(row.totals).font(ScaledType.mono).foregroundStyle(FSColor.text)
        }
        .padding(GuidesLayout.cardPadding)
        .frame(minHeight: 56)
        .accessibilityElement(children: .combine)
    }
}

/// Opener, Priority, Cooldowns: numbered rows of the guide's own lists, then the footer.
struct GuideRotationSection: View {
    let rotation: GuideRotation?
    let build: String

    var body: some View {
        VStack(alignment: .leading, spacing: 12) {
            ForEach(RotationRows.groups(rotation)) { group in
                VStack(alignment: .leading, spacing: 8) {
                    MicroLabel(group.title).accessibilityAddTraits(.isHeader)
                    GuideCardRows(items: group.rows) { GuideRotationRow(row: $0, build: build) }
                }
            }
            Text(Copy.guidesRotationFooter).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                .fixedSize(horizontal: false, vertical: true)
        }
    }
}

struct GuideRotationRow: View {
    let row: RotationRow
    let build: String

    var body: some View {
        HStack(alignment: .top, spacing: 12) {
            Text("\(row.number)").font(ScaledType.monoMeta).foregroundStyle(FSColor.muted)
                .frame(minWidth: 16, alignment: .leading)
                .padding(.top, 11)
            if row.hasTile { SpecIcon(icon: row.spell.icon, build: build, size: GuidesLayout.rowIcon) }
            VStack(alignment: .leading, spacing: 3) {
                ViewThatFits(in: .horizontal) {
                    HStack(alignment: .firstTextBaseline, spacing: 8) { name; rank }
                    VStack(alignment: .leading, spacing: 2) { name; rank }
                }
                Text(row.spell.clause).font(ScaledType.meta).foregroundStyle(FSColor.nav).fixedSize(horizontal: false, vertical: true)
            }
            .frame(maxWidth: .infinity, alignment: .leading)
        }
        .padding(GuidesLayout.cardPadding)
        .accessibilityElement(children: .combine)
    }

    private var name: some View {
        Text(row.spell.name).font(ScaledType.name).foregroundStyle(FSColor.strong).fixedSize(horizontal: false, vertical: true)
    }

    @ViewBuilder
    private var rank: some View {
        if let rank = row.rankText { Text(rank).font(ScaledType.monoMeta).foregroundStyle(FSColor.muted) }
    }
}

/// The stat priority again, with the hit cap when the guide has one. No weights (sims own numbers).
struct GuideStatsCard: View {
    let guide: GuideFile

    var body: some View {
        VStack(alignment: .leading, spacing: 12) {
            StatChips(chips: GuideCopy.statChips(guide))
            if let hitCap = guide.hitCap {
                Text(hitCap).font(ScaledType.bodyStrong).foregroundStyle(FSColor.text)
            }
        }
        .padding(GuidesLayout.cardPadding)
        .frame(maxWidth: .infinity, alignment: .leading)
        .quietCard()
    }
}

/// `Best in slot for {spec} · 60 · {preset}`: switches to the BiS tab on that list.
struct GuideGearRow: View {
    let guide: GuideFile
    let open: (GuideGear) -> Void

    var body: some View {
        if let gear = guide.gear, let text = GuideCopy.gearRow(guide) {
            GuideLinkRow(text: text) { open(gear) }.quietCard()
        }
    }
}

struct GuideRacesCard: View {
    let races: [GuideRaceRow]

    var body: some View {
        GuideCardRows(items: races.map(IdentifiedRace.init)) { GuideRaceRowView(race: $0.race) }
    }
}

private struct IdentifiedRace: Identifiable {
    let race: GuideRaceRow
    var id: String { race.slug }
}

struct GuideRaceRowView: View {
    let race: GuideRaceRow

    var body: some View {
        HStack(alignment: .top, spacing: 12) {
            FactionMark(faction: race.faction, size: GuidesLayout.factionMark)
            VStack(alignment: .leading, spacing: 6) {
                HStack(spacing: 8) {
                    Text(race.name).font(ScaledType.name).foregroundStyle(FSColor.strong)
                    if race.recommended { TagView(text: Copy.guidesRecommended, color: FSColor.kill) }
                }
                ForEach(race.racials, id: \.name) { racial in
                    Text("\(Text(racial.name).font(ScaledType.bodyStrong).foregroundStyle(FSColor.text)): \(racial.effect)")
                        .font(ScaledType.meta).foregroundStyle(FSColor.nav)
                        .fixedSize(horizontal: false, vertical: true)
                }
            }
            .frame(maxWidth: .infinity, alignment: .leading)
        }
        .padding(GuidesLayout.cardPadding)
        .accessibilityElement(children: .combine)
    }
}
