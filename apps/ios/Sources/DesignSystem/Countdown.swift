import SwiftUI

/// One part of a countdown's clock: its figure and its unit (`02`, `hours`).
public struct FSClockPart: Equatable, Sendable {
    public let value: String
    public let unit: String

    public init(_ value: String, _ unit: String) {
        self.value = value
        self.unit = unit
    }
}

/// `ui/Countdown` (`.cdn.guild`): the raised card with the guild's gold glow, an eyebrow, the clock
/// in parts (`02 hours`, `48 min`, read aloud as one timer) and an aside under it. A live raid has
/// no clock: a red dot leads the eyebrow and the aside follows.
public struct FSCountdown<Aside: View>: View {
    private let eyebrow: String
    private let clock: [FSClockPart]
    private let live: Bool
    private let aside: Aside
    @ScaledMetric(relativeTo: .largeTitle) private var clockLine: CGFloat = 42

    static var padX: CGFloat { FSSpace.space6 }
    static var padY: CGFloat { FSSpace.space5 }
    /// `.cdn .rw { margin-top: 10px }` and its row gap.
    static var clockTop: CGFloat { 10 }
    static var asideTop: CGFloat { FSSpace.space4 }
    static var partGap: CGFloat { 14 }
    static var liveDot: CGFloat { 8 }
    /// `.cdn` radius.
    static var radius: CGFloat { 8 }

    public init(eyebrow: String, clock: [FSClockPart] = [], live: Bool = false, @ViewBuilder aside: () -> Aside) {
        self.eyebrow = eyebrow
        self.clock = clock
        self.live = live
        self.aside = aside()
    }

    public var body: some View {
        VStack(alignment: .leading, spacing: 0) {
            HStack(spacing: FSSpace.space2) {
                if live {
                    Circle().fill(FSColor.kill).frame(width: Self.liveDot, height: Self.liveDot).accessibilityHidden(true)
                }
                Text(eyebrow)
                    .type(TypeRole(face: FSFont.bodyBold, size: 12, lineHeight: 16, scalesWith: .caption))
                    .foregroundStyle(FSColor.muted)
            }
            if !clock.isEmpty {
                HStack(alignment: .top, spacing: Self.partGap) {
                    ForEach(Array(clock.enumerated()), id: \.offset) { _, part in
                        VStack(spacing: 0) {
                            Text(part.value)
                                .type(TypeRole(face: FSFont.bodyBold, size: 40, lineHeight: 42, scalesWith: .largeTitle, tabular: true))
                                .foregroundStyle(FSColor.strong)
                                // Barlow's natural line at 40 is taller than the 42 pt line box.
                                .frame(height: clockLine)
                            Text(part.unit)
                                .type(TypeRole(face: FSFont.body, size: 12, lineHeight: 16, scalesWith: .caption))
                                .foregroundStyle(FSColor.muted)
                        }
                    }
                }
                .padding(.top, Self.clockTop)
                .accessibilityElement(children: .ignore)
                .accessibilityLabel(clock.map { "\(Int($0.value) ?? 0) \($0.unit)" }.joined(separator: ", "))
                .accessibilityAddTraits(.updatesFrequently)
            }
            aside
                .frame(maxWidth: .infinity, alignment: .leading)
                .padding(.top, clock.isEmpty ? Self.clockTop : Self.asideTop)
        }
        .padding(.horizontal, Self.padX)
        .padding(.vertical, Self.padY)
        .frame(maxWidth: .infinity, alignment: .leading)
        .background {
            RoundedRectangle(cornerRadius: Self.radius)
                .fill(FSColor.raised)
                .overlay(
                    RadialGradient(
                        colors: [FSColor.gold.opacity(0.1), .clear], center: .topTrailing, startRadius: 0, endRadius: 280
                    )
                    .clipShape(RoundedRectangle(cornerRadius: Self.radius))
                )
        }
        .overlay(RoundedRectangle(cornerRadius: Self.radius).strokeBorder(FSColor.line, lineWidth: FSStroke.hairline))
    }
}

/// `ui/Tally`: counts against a target, one labelled meter each, in equal columns. The figure is
/// green at the target, gold within a quarter of it and red below.
public struct FSTally: View {
    public struct Item: Equatable, Sendable, Identifiable {
        public let label: String
        public let have: Int
        public let want: Int

        public init(label: String, have: Int, want: Int) {
            self.label = label
            self.have = have
            self.want = want
        }

        public var id: String { label }

        var met: Bool { have >= want }
        var near: Bool { Double(have) >= Double(want) * 0.75 }
        var fraction: Double { want > 0 ? min(1, max(0, Double(have) / Double(want))) : 1 }
    }

    private let items: [Item]

    public init(_ items: [Item]) {
        self.items = items
    }

    static var columnGap: CGFloat { FSSpace.space4 }
    static var meterHeight: CGFloat { 8 }

    public var body: some View {
        HStack(alignment: .top, spacing: Self.columnGap) {
            ForEach(items) { item in
                VStack(alignment: .leading, spacing: 6) {
                    HStack(spacing: FSSpace.space2) {
                        Text(item.label).type(FSType.item).foregroundStyle(FSColor.strong).lineLimit(1)
                        Spacer(minLength: 0)
                        Text("\(item.have)/\(item.want)")
                            .type(TypeRole(face: FSFont.bodyBold, size: 15, lineHeight: 20, scalesWith: .body, tabular: true))
                            .foregroundStyle(item.met ? FSColor.kill : (item.near ? FSColor.gold : FSColor.wipe))
                    }
                    FSMeter(fraction: item.fraction, tone: item.met ? .kill : .gold)
                }
                .frame(maxWidth: .infinity, alignment: .leading)
                .accessibilityElement(children: .ignore)
                .accessibilityLabel("\(item.label): \(item.have) of \(item.want)")
            }
        }
    }
}

/// `ui/Meter`: an 8 pt track in `lineSoft` and its fill.
public struct FSMeter: View {
    public enum Tone: Sendable { case gold, kill, muted }

    private let fraction: Double
    private let tone: Tone

    public init(fraction: Double, tone: Tone = .gold) {
        self.fraction = fraction
        self.tone = tone
    }

    public var body: some View {
        GeometryReader { geo in
            ZStack(alignment: .leading) {
                RoundedRectangle(cornerRadius: FSRadius.pill).fill(FSColor.lineSoft)
                RoundedRectangle(cornerRadius: FSRadius.pill).fill(color).frame(width: geo.size.width * fraction)
            }
        }
        .frame(height: FSTally.meterHeight)
        .accessibilityHidden(true)
    }

    private var color: Color {
        switch tone {
        case .gold: FSColor.gold
        case .kill: FSColor.kill
        case .muted: FSColor.muted
        }
    }
}

/// `ui/Skeleton`: a block at its final size while its content loads, so nothing shifts.
public struct FSSkeleton: View {
    private let height: CGFloat

    public init(height: CGFloat) {
        self.height = height
    }

    public var body: some View {
        RoundedRectangle(cornerRadius: FSRadius.field)
            .fill(FSColor.lineSoft)
            .frame(maxWidth: .infinity, minHeight: height, maxHeight: height)
            .accessibilityHidden(true)
    }
}
