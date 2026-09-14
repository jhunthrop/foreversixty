// logs/engine/parquet/file.go
package parquet

import (
	"bytes"
	"fmt"
	"io"
	"sort"

	pq "github.com/parquet-go/parquet-go"
	"github.com/parquet-go/parquet-go/compress/zstd"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
)

// Write encodes a fight's events. Rows are sorted by time then line, as the
// spec requires, and the writer's metadata is pinned, so the same events
// always produce the same bytes.
func Write(w io.Writer, events []event.Event) error {
	rows := make([]Row, len(events))
	for i, e := range events {
		rows[i] = RowOf(e)
	}
	sort.SliceStable(rows, func(i, j int) bool {
		if rows[i].TimeUnixNano != rows[j].TimeUnixNano {
			return rows[i].TimeUnixNano < rows[j].TimeUnixNano
		}
		return rows[i].Line < rows[j].Line
	})
	pw := pq.NewGenericWriter[Row](w,
		pq.Compression(&zstd.Codec{Level: zstd.SpeedDefault}),
		pq.CreatedBy(CreatedBy, CreatedVersion, CreatedBuild),
	)
	if len(rows) > 0 {
		if _, err := pw.Write(rows); err != nil {
			return fmt.Errorf("parquet: write rows: %w", err)
		}
	}
	if err := pw.Close(); err != nil {
		return fmt.Errorf("parquet: close: %w", err)
	}
	return nil
}

// Marshal is Write into a byte slice, which is what the store writes to R2.
func Marshal(events []event.Event) ([]byte, error) {
	var buf bytes.Buffer
	if err := Write(&buf, events); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Read decodes a fight's events.
func Read(r io.ReaderAt, size int64) ([]event.Event, error) {
	rows, err := pq.Read[Row](r, size)
	if err != nil {
		return nil, fmt.Errorf("parquet: read: %w", err)
	}
	out := make([]event.Event, len(rows))
	for i, row := range rows {
		out[i] = EventOf(row)
	}
	return out, nil
}

// Unmarshal is Read from a byte slice.
func Unmarshal(b []byte) ([]event.Event, error) {
	return Read(bytes.NewReader(b), int64(len(b)))
}
