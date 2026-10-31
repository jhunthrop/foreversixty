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
        .quietCard()
    }

    private var classes: some View {
        let columns = Array(repeating: GridItem(.flexible(), spacing: 8, alignment: .leading), count: typeSize.isAccessibilitySize ? 1 : 3)
        return LazyVGrid(columns: columns, alignment: .leading, spacing: 10) {
            ForEach(read.classes) { line in
                HStack(spacing: 8) {
                    CrestView(className: line.className, size: 24)
                    Text(line.name).font(Theme.Typography.metaText).lineLimit(1).minimumScaleFactor(0.7)
                        .foregroundStyle(line.count == 0 ? Theme.Palette.muted : Theme.Palette.classColor(line.className))
                    Text("\(line.count)").font(Theme.Typography.monoMeta)
                        .foregroundStyle(line.count == 0 ? Theme.Palette.muted : Theme.Palette.strong)
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
            Text(role).font(Theme.Typography.bodyStrong).foregroundStyle(Theme.Palette.strong)
        }
    }

    private var buffs: some View {
        VStack(alignment: .leading, spacing: 8) {
            SectionLabel(text: Copy.raidsCompBuffs)
            LazyVGrid(columns: [GridItem(.adaptive(minimum: 40), spacing: 8)], alignment: .leading, spacing: 8) {
                ForEach(read.buffs) { buff in
                    GameIcon(stem: buff.icon, size: 36, border: buff.present ? Theme.Palette.kill.opacity(0.6) : Theme.Palette.line, fallback: buff.label)
                        .opacity(buff.present ? 1 : 0.35)
                        .saturation(buff.present ? 1 : 0)
                        .accessibilityHidden(false)
                        .accessibilityElement()
                        .accessibilityLabel(buff.accessibilityLabel)
                }
            }
            Text(read.buffSummary).font(Theme.Typography.metaText)
                .foregroundStyle(read.buffs.allSatisfy(\.present) ? Theme.Palette.kill : Theme.Palette.muted)
        }
    }
}
