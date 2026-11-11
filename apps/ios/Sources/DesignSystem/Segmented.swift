import SwiftUI

/// `ui/Segmented` (brand B1, round 5 `.r5raid .seg`): one connected group, the chosen segment in
/// teal with its edge, hairline dividers between, 44 pt; at the accessibility sizes the segments
/// stack two to a row (2 x 2 for the four answers, as Android at 1.5x) at the control radius, so no
/// word breaks (round 4, row 4). VoiceOver reads it as one segmented control
/// (a radio group: `1 of 4`, the chosen one selected). The press's haptic (R14) is the caller's,
/// through its injected haptics, so a test records it.
public struct FSSegmented<Value: Hashable>: View {
    public struct Option: Identifiable {
        public let value: Value
        public let label: String
        public var id: Value { value }

        public init(_ value: Value, _ label: String) {
            self.value = value
            self.label = label
        }
    }

    private let label: String
    private let options: [Option]
    private let selection: Value?
    private let select: (Value) -> Void
    private let height: CGFloat
    @Environment(\.isEnabled) private var isEnabled
    @Environment(\.dynamicTypeSize) private var typeSize

    static var labelSize: CGFloat { 13 }

    /// `label` names the group to VoiceOver; `selection` nil draws none chosen.
    /// `compact`: the 36 pt face of a row's chips (the raid line), still in a 44 pt-wide group.
    public init(_ label: String, options: [Option], selection: Value?, compact: Bool = false, select: @escaping (Value) -> Void) {
        height = compact ? FSButtonStyle.smallHeight : FSTarget.min
        self.label = label
        self.options = options
        self.selection = selection
        self.select = select
    }

    /// Segments to a row in the stacked form.
    static var stackedColumns: Int { 2 }

    private var radius: CGFloat { typeSize.isAccessibilitySize ? FSRadius.control : FSRadius.field }

    public var body: some View {
        group
            .background(FSColor.raised, in: .rect(cornerRadius: radius))
            .clipShape(.rect(cornerRadius: radius))
            .overlay(RoundedRectangle(cornerRadius: radius).strokeBorder(FSColor.rule, lineWidth: FSStroke.hairline))
            .opacity(isEnabled ? 1 : 0.5)
            .accessibilityRepresentation {
                Picker(label, selection: Binding(get: { selection }, set: { $0.map(select) })) {
                    ForEach(options) { Text($0.label).tag(Optional($0.value)) }
                }
                .pickerStyle(.segmented)
            }
    }

    @ViewBuilder
    private var group: some View {
        if typeSize.isAccessibilitySize {
            let rows = stride(from: 0, to: options.count, by: Self.stackedColumns).map {
                Array(options[$0 ..< min($0 + Self.stackedColumns, options.count)])
            }
            VStack(spacing: 0) {
                ForEach(Array(rows.enumerated()), id: \.offset) { index, row in
                    HStack(spacing: 0) {
                        ForEach(Array(row.enumerated()), id: \.element.id) { column, option in
                            segment(option)
                            if column < Self.stackedColumns - 1 { Rectangle().fill(FSColor.rule).frame(width: FSStroke.hairline) }
                        }
                        // A short last row keeps its cells the width of the rows above.
                        if row.count < Self.stackedColumns { Color.clear.frame(maxWidth: .infinity) }
                    }
                    .fixedSize(horizontal: false, vertical: true)
                    if index < rows.count - 1 { Rectangle().fill(FSColor.rule).frame(height: FSStroke.hairline) }
                }
            }
        } else {
            HStack(spacing: 0) {
                ForEach(Array(options.enumerated()), id: \.element.id) { index, option in
                    segment(option)
                    if index < options.count - 1 { Rectangle().fill(FSColor.rule).frame(width: FSStroke.hairline) }
                }
            }
            .fixedSize(horizontal: false, vertical: true)
        }
    }

    private func segment(_ option: Option) -> some View {
        let chosen = option.value == selection
        return Button { select(option.value) } label: {
            Text(option.label)
                .font(.custom(chosen ? FSFont.bodyBold : FSFont.bodySemibold, size: Self.labelSize, relativeTo: .subheadline))
                .foregroundStyle(chosen ? FSColor.strong : FSColor.text)
                .lineLimit(1)
                .padding(.horizontal, FSSpace.space1)
                .frame(maxWidth: .infinity, minHeight: height)
                .background(chosen ? FSColor.teal : .clear)
                .overlay { if chosen { Rectangle().strokeBorder(FSColor.tealEdge, lineWidth: FSStroke.hairline) } }
                .contentShape(.rect)
        }
        .buttonStyle(SegmentPress())
        .parityItem(.button)
    }
}

/// The press on a teal control (the Segmented, Activity's chips): the edge's teal as a veil.
public enum TealPress {
    public static var veil: Color { FSColor.tealEdge.opacity(0.25) }
}

private struct SegmentPress: ButtonStyle {
    func makeBody(configuration: Configuration) -> some View {
        configuration.label.overlay(configuration.isPressed ? TealPress.veil : .clear)
    }
}
