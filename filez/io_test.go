package filez_test

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-lib/errorz"
	"github.com/ibrt/golang-lib/filez"
	"github.com/ibrt/golang-lib/fixturez"
)

type IOSuite struct {
	// intentionally empty
}

func TestIOSuite(t *testing.T) {
	fixturez.RunSuite(t, &IOSuite{})
}

func (*IOSuite) TestFilesAndDirs(g *WithT) {
	{
		filePath := filez.MustCreateTempFile([]byte("content"))
		defer func() { errorz.MaybeMustWrap(os.Remove(filePath)) }()

		g.Expect(filez.MustReadFile(filePath)).To(Equal([]byte("content")))
		g.Expect(filez.MustReadFileString(filePath)).To(Equal("content"))
	}

	{
		filePath := filez.MustCreateTempFileString("content")
		defer func() { errorz.MaybeMustWrap(os.Remove(filePath)) }()

		g.Expect(filez.MustReadFile(filePath)).To(Equal([]byte("content")))
		g.Expect(filez.MustReadFileString(filePath)).To(Equal("content"))
	}

	{
		dirPath := filez.MustCreateTempDir()
		defer filez.MustRemoveAll(dirPath)

		filePath := filez.MustWriteFile(filepath.Join(dirPath, "first"), 0777, 0666, []byte("content"))
		g.Expect(filez.MustReadFile(filePath)).To(Equal([]byte("content")))
		g.Expect(filez.MustReadFileString(filePath)).To(Equal("content"))

		filePath = filez.MustWriteFileString(filepath.Join(dirPath, "second"), 0777, 0666, "content")
		g.Expect(filez.MustReadFile(filePath)).To(Equal([]byte("content")))
		g.Expect(filez.MustReadFileString(filePath)).To(Equal("content"))

		g.Expect(filez.MustCheckPathExists(dirPath)).To(BeTrue())
		g.Expect(filez.MustCheckPathExists(filepath.Join(dirPath, "third"))).To(BeFalse())
		g.Expect(func() { filez.MustCheckPathExists(string([]byte{0})) }).To(Panic())

		g.Expect(filez.MustCheckFileExists(filePath)).To(BeTrue())
		g.Expect(filez.MustCheckFileExists(filepath.Join(dirPath, "third"))).To(BeFalse())
		g.Expect(filez.MustCheckFileExists(dirPath)).To(BeFalse())
		g.Expect(func() { filez.MustCheckFileExists(string([]byte{0})) }).To(Panic())

		filez.MustPrepareDir(dirPath, 0777)
		g.Expect(filez.MustCheckFileExists(filePath)).To(BeFalse())
	}
}
