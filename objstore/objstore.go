package objstore

import (
	"context"
	"io"
	"iter"
	"time"
)

// Bucket represent the Bucket interface contract.
type Bucket interface {
	Put(context.Context, Key, io.Reader, PutOptions) (ObjectInfo, error)
	Get(context.Context, Key) (io.ReadCloser, ObjectInfo, error)
	Stat(context.Context, Key) (ObjectInfo, error)
	Delete(context.Context, Key) error
	List(context.Context, string) iter.Seq2[ObjectInfo, error]
	PresignGet(context.Context, Key, time.Duration) (string, error)
}
