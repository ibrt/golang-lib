package randomfixtures

import (
	"bytes"
	"context"
	"testing"

	"github.com/ibrt/golang-lib/errors"
	"github.com/ibrt/golang-lib/inject"
	"github.com/ibrt/golang-lib/modules/random"
)

// Fixtures provides test fixtures for Random using a real PRNG.
type Fixtures struct {
	releaser inject.Releaser
}

// BeforeSuite implements fixtures.BeforeSuite.
func (f *Fixtures) BeforeSuite(ctx context.Context) context.Context {
	injector, releaser, err := random.Initializer(ctx)
	errors.MaybeMustWrap(err)
	f.releaser = releaser
	return inject.MustInject(ctx, injector)
}

// AfterSuite implements fixtures.AfterSuite.
func (f *Fixtures) AfterSuite(ctx context.Context) {
	inject.SafeRelease(f.releaser)
}

// DeterministicFixtures provides text fixtures for Random using a deterministic sequence.
type DeterministicFixtures struct {
	// intentionally empty
}

// BeforeTest implements the fixtures.BeforeTest interface.
func (f *DeterministicFixtures) BeforeTest(ctx context.Context, _ *testing.T) context.Context {
	return inject.MustInject(ctx, random.SingletonInjectorFactory(random.NewRandom(&deterministicReader{})))
}

type deterministicReader struct {
	counter byte
}

// Read implements the io.Reader interface.
func (r *deterministicReader) Read(buf []byte) (int, error) {
	for i := 0; i < len(buf); i++ {
		buf[i] = r.counter
		r.counter++
	}
	return len(buf), nil
}

// ConfigurableFixtures provides text fixtures for Random using a configurable sequence.
type ConfigurableFixtures struct {
	// intentionally empty
}

// Set injects a new Random that reads from the given buffer.
func (f *ConfigurableFixtures) Set(ctx context.Context, buf []byte) context.Context {
	return inject.MustInject(ctx, random.SingletonInjectorFactory(random.NewRandom(bytes.NewReader(buf))))
}

// MakeSequence returns a buffer containing a sequence of bytes of the given length.
// The first byte is equal to "start" and each consecutive byte is increased by one.
func MakeSequence(start byte, length int) []byte {
	buf := make([]byte, length)
	for i := 0; i < len(buf); i++ {
		buf[i] = start + byte(i)
	}
	return buf
}
