package randomfixtures

import (
	"context"
	"crypto/rand"
	"testing"

	"github.com/ibrt/golang-lib/inject"
	"github.com/ibrt/golang-lib/modules/random"
)

// Fixtures provides test fixtures for Random using a real PRNG.
type Fixtures struct {
	// intentionally empty
}

// BeforeSuite implements fixtures.BeforeSuite.
func (f *Fixtures) BeforeSuite(ctx context.Context) context.Context {
	return inject.MustInject(ctx, random.SingletonInjectorFactory(random.NewRandom(rand.Reader)))
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
