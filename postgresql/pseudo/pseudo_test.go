package pseudo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPseudoEncrypt(t *testing.T) {
	Convey("pseudo", t, func() {
		So(PseudoEncrypt(123456789), ShouldEqual, -854259877)
	})
}

func TestPseudoEncrypt64(t *testing.T) {
	Convey("pseudo given a limit len to gennerate a number of pseudo ", t, func() {
		So(PseudoEncrypt(12345678912345), ShouldEqual, 218442452601)
	})
}
