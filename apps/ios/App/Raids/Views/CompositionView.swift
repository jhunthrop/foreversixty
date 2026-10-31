import DesignSystem
import SwiftUI

/// The composition block at the top of the roster: classes with names and counts, roles against
/// the raid's targets, and the raid buffs present or missing.
struct CompositionView: View {
    let read: CompositionRead
    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        VStack(alignment: .leading, spacing: 14) {
            SectionLabel(text: Copy.raidsCompTitle)
            classes
            roles
            buffs
        }
        .padding(14)
        .raidCardBox(0).quietCard()
    }

    private var classes: some View {
        let columns = Array(repeating: GridItem(.flexible(), spacing: 8, alignment: .leading), count: typeSize.isAccessibilitySize ? 1 : 3)
        return LazyVGrid(columns: columns, alignment: .leading, spacing: 10) {
            ForEach(read.classes) { line in
                HStack(spacing: 8) {
                    ClassCrest(className: line.className, size: 24)
                    Text(line.name).font(ScaledType.meta).lineLimit(1).minimumScaleFactor(0.7)
                        .foregroundStyle(line.count == 0 ? FSColor.muted : ClassColor.of(line.className))
                    Text("\(line.count)").font(ScaledType.monoMeta)
                        .foregroundStyle(line.count == 0 ? FSColor.muted : FSColor.strong)
                }
            }
        }
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(read.classesAccessibilityLabel)
    }

    private var roles: some View {
        ViewThatFits(in: .horizontal) {
            HStack(spacing: 16) { roleTexts }
            VStack(alignment: .leading, spacing: 6) { roleTexts }
        }
    }

    private var roleTexts: some View {
        ForEach(read.roles, id: \.self) { role in
            Text(role).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
        }
    }

    private var buffs: some View {
        VStack(alignment: .leading, spacing: 8) {
            SectionLabel(text: Copy.raidsCompBuffs)
            LazyVGrid(columns: [GridItem(.adaptive(minimum: 40), spacing: 8)], alignment: .leading, spacing: 8) {
                ForEach(read.buffs) { buff in
                    GameIcon(
                        stem: buff.icon, size: 36, border: buff.present ? FSColor.kill.opacity(0.6) : FSColor.line,
                        fallback: buff.label, dimmed: !buff.present
                    )
                    .accessibilityHidden(false)
                    .accessibilityElement()
                    .accessibilityLabel(buff.accessibilityLabel)
                }
            }
            Text(read.buffSummary).font(ScaledType.meta)
                .foregroundStyle(read.buffs.allSatisfy(\.present) ? FSColor.kill : FSColor.muted)
        }
    }
}
