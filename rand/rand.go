/*
Package rand ...

Copyright © 2019 hajime-terasawa <terako.studio@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package rand

import (
	crand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"strings"
	"time"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	minFloat = 0.0
)

// nolint:gochecknoinits
func init() {
	seed, err := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		rand.Seed(time.Now().UnixNano())
	} else {
		rand.Seed(seed.Int64())
	}
}

// GenInt generates int64 between given range
func GenInt(min, max int64) int64 {
	// nolint:gosec
	return rand.Int63n(max-min) + min
}

func genString(n int) string {
	b := make([]byte, n)
	for i := range b {
		// nolint:gosec
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func genTime(min, max time.Time) time.Time {
	minI := min.Unix()
	maxI := max.Unix()
	delta := maxI - minI
	// nolint:gosec
	sec := rand.Int63n(delta) + minI
	return time.Unix(sec, 0)
}

// Boolean returns random boolean
func Boolean() bool {
	// nolint:gosec
	return rand.Float32() < 0.5
}

// TinyInt returns random tinyint
func TinyInt() int8 {
	return int8(GenInt(-128, 127))
}

// UnsignedTinyInt returns random unsigned tinyint
func UnsignedTinyInt() uint8 {
	return uint8(GenInt(0, 255))
}

// SmallInt returns random smallint
func SmallInt() int16 {
	return int16(GenInt(-32768, 32767))
}

// UnsignedSmallInt returns random unsigned smallint
func UnsignedSmallInt() uint16 {
	return uint16(GenInt(0, 65535))
}

// MediumInt returns random mediumint
func MediumInt() int32 {
	return int32(GenInt(-8388608, 8388607))
}

// UnsignedMediumInt returns random unsigned mediumint
func UnsignedMediumInt() uint32 {
	return uint32(GenInt(0, 16777215))
}

// Int returns random int
func Int() int32 {
	return int32(GenInt(-2147483648, 2147483647))
}

// UnsignedInt returns random unsigned int
func UnsignedInt() uint32 {
	// nolint:gosec
	return rand.Uint32()
}

// BigInt returns random bigint
func BigInt() int64 {
	if Boolean() {
		return GenInt(0, 9223372036854775807)
	}
	return GenInt(-9223372036854775808, -1)
}

// UnsignedBigInt returns random unsigned bigint
func UnsignedBigInt() uint64 {
	// nolint:gosec
	return uint64(rand.Uint32())<<32 + uint64(rand.Uint32())
}

// Decimal returns random decimal within the given range
func Decimal(order, precision int) string {
	double := Double(order, precision)
	if double == 0 {
		return "0"
	}
	return fmt.Sprint(double)
}

// UnsignedDecimal returns random decimal within the given range
func UnsignedDecimal(order, precision int) string {
	double := UnsignedDouble(order, precision)
	if double == 0 {
		return "0"
	}
	return fmt.Sprint(double)
}

// Float returns random float within the given range
func Float(order, precision int) float32 {
	unsigned := Boolean()
	min := float32(minFloat)
	var max float32

	if unsigned {
		// nolint:gomnd
		max = float32(math.Pow(10, float64(order-precision)) - 1)
	} else {
		// nolint:gomnd
		max = float32(math.Pow(10, float64(order-precision-1)) - 1)
	}

	// nolint:gosec
	float := min + rand.Float32()*(max-min)
	if unsigned {
		return float
	}
	return float * -1
}

// UnsignedFloat returns random unsigned float within the given range
func UnsignedFloat(order, precision int) float32 {
	min := float32(minFloat)
	// nolint:gomnd
	max := float32(math.Pow(10, float64(order-precision)) - 1)
	// nolint:gosec
	return min + rand.Float32()*(max-min)
}

// Double returns random double within the given range
func Double(order, precision int) float64 {
	unsigned := Boolean()

	var max float64

	if unsigned {
		// nolint:gomnd
		max = math.Pow(10, float64(order-precision)) - 1
	} else {
		// nolint:gomnd
		max = math.Pow(10, float64(order-precision-1)) - 1
	}

	// nolint:gosec
	double := minFloat + rand.Float64()*(max-minFloat)
	if unsigned {
		return double
	}
	return double * -1
}

// UnsignedDouble returns random unsigned double within the given range
func UnsignedDouble(order, precision int) float64 {
	// nolint:gomnd
	max := math.Pow(10, float64(order-precision)) - 1
	// nolint:gosec
	return minFloat + rand.Float64()*(max-minFloat)
}

// Real returns random double within the given range
func Real(order, precision int) float64 {
	return Double(order, precision)
}

// UnsignedReal returns random unsigned double within the given range
func UnsignedReal(order, precision int) float64 {
	return UnsignedDouble(order, precision)
}

// Bit returns random bit
func Bit(order int) string {
	var sb strings.Builder
	sb.WriteString("b'")
	for i := 0; i < order; i++ {
		sb.WriteString(fmt.Sprint(GenInt(0, 1)))
	}
	sb.WriteString("'")
	return sb.String()
}

// Date returns random date
func Date() string {
	min := time.Date(1000, 1, 1, 0, 0, 0, 0, time.UTC)
	max := time.Date(9999, 12, 31, 0, 0, 0, 0, time.UTC)
	randTime := genTime(min, max)
	y, m, d := randTime.Date()

	return fmt.Sprintf("%d-%02d-%02d", y, int(m), d)
}

// DateTime returns random datetime
func DateTime() string {
	min := time.Date(1000, 1, 1, 0, 0, 0, 0, time.UTC)
	max := time.Date(9999, 12, 31, 0, 0, 0, 0, time.UTC)
	randTime := genTime(min, max)
	y, m, d := randTime.Date()
	hour := randTime.Hour()
	minu := randTime.Minute()
	sec := randTime.Second()

	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", y, m, d, hour, minu, sec)
}

// Timestamp returns random timestamp
func Timestamp() string {
	min := time.Date(1970, 1, 1, 0, 0, 1, 0, time.UTC)
	max := time.Date(2038, 1, 19, 3, 14, 7, 0, time.UTC)
	randTime := genTime(min, max)
	y, m, d := randTime.Date()
	h := randTime.Hour()
	mi := randTime.Minute()
	s := randTime.Second()
	l := randTime.Location()

	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d %v", y, m, d, h, mi, s, l)
}

// Time returns random time
func Time() string {
	min := time.Date(1970, 1, 1, 0, 0, 1, 0, time.UTC)
	max := time.Date(2038, 1, 19, 3, 14, 7, 0, time.UTC)
	randTime := genTime(min, max)
	h := randTime.Hour()
	mi := randTime.Minute()
	s := randTime.Second()

	return fmt.Sprintf("%02d:%02d:%02d", h, mi, s)
}

// Year4 returns random year(4)
func Year4() string {
	// nolint:gomnd
	return fmt.Sprint(GenInt(1901, 2155))
}

// Year2 returns random year(2)
func Year2() string {
	return fmt.Sprint(GenInt(0, 99))
}

// Char returns random char with the given length
func Char(length int) string {
	return genString(length)
}

// VarChar returns random varchar with the given length
func VarChar(length int) string {
	return genString(length)
}

// Binary returns random binary with the given length
func Binary(length int) string {
	return genString(length)
}

// VarBinary returns random varbinary with the given length
func VarBinary(length int) string {
	return genString(length)
}

// TinyBlob returns random tiny blob with the given length
func TinyBlob(length int) string {
	return genString(length)
}

// TinyText returns random tiny text with the given length
func TinyText(length int) string {
	return genString(length)
}

// Blob returns random blob with the given length
func Blob(length int) string {
	return genString(length)
}

// Text returns random text with the given length
func Text(length int) string {
	return genString(length)
}

// MediumBlob returns random medium blob with the given length
func MediumBlob(length int) string {
	return genString(length)
}

// MediumText returns random medium text with the given length
func MediumText(length int) string {
	return genString(length)
}

// LongBlob returns random long blob with the given length
func LongBlob(length int) string {
	return genString(length)
}

// LongText returns random long text with the given length
func LongText(length int) string {
	return genString(length)
}
