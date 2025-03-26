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
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

// Boolean returns random boolean.
func Boolean() bool {
	b := gofakeit.Number(0, 1)
	return b == 1
}

// TinyInt returns random tinyint.
func TinyInt() int8 {
	return gofakeit.Int8()
}

// UnsignedTinyInt returns random unsigned tinyint.
func UnsignedTinyInt() uint8 {
	return gofakeit.Uint8()
}

// SmallInt returns random smallint.
func SmallInt() int16 {
	return gofakeit.Int16()
}

// UnsignedSmallInt returns random unsigned smallint.
func UnsignedSmallInt() uint16 {
	return gofakeit.Uint16()
}

// MediumInt returns random mediumint.
func MediumInt() int32 {
	return gofakeit.Int32()
}

// UnsignedMediumInt returns random unsigned mediumint.
func UnsignedMediumInt() uint32 {
	return gofakeit.Uint32()
}

// Int returns random int.
func Int() int32 {
	return gofakeit.Int32()
}

// UnsignedInt returns random unsigned int.
func UnsignedInt() uint32 {
	return gofakeit.Uint32()
}

// BigInt returns random bigint.
func BigInt() int64 {
	return gofakeit.Int64()
}

// UnsignedBigInt returns random unsigned bigint.
func UnsignedBigInt() uint64 {
	return gofakeit.Uint64()
}

// Decimal returns random decimal within the given range.
func Decimal(order, precision int) string {
	double := Double(order, precision)
	if double == 0 {
		return "0"
	}

	return fmt.Sprint(double)
}

// UnsignedDecimal returns random decimal within the given range.
func UnsignedDecimal(order, precision int) string {
	double := UnsignedDouble(order, precision)
	if double == 0 {
		return "0"
	}

	return fmt.Sprint(double)
}

// Float returns random float within the given range.
func Float(order, precision int) float32 {
	unsigned := Boolean()

	var (
		minF = float32(math.SmallestNonzeroFloat32)
		maxF float32
	)

	if unsigned {
		maxF = float32(math.Pow(10, float64(order-precision)) - 1) //nolint:mnd
	} else {
		maxF = float32(math.Pow(10, float64(order-precision-1)) - 1) //nolint:mnd
	}

	output := gofakeit.Float32Range(minF, maxF)
	if unsigned {
		return output
	}

	return output * -1
}

// UnsignedFloat returns random unsigned float within the given range.
func UnsignedFloat(order, precision int) float32 {
	var (
		minF float32 = math.SmallestNonzeroFloat32
		maxF         = float32(math.Pow(10, float64(order-precision)) - 1) //nolint:mnd
	)

	return gofakeit.Float32Range(minF, maxF)
}

// Double returns random double within the given range.
func Double(order, precision int) float64 {
	unsigned := Boolean()

	var (
		minF = math.SmallestNonzeroFloat64
		maxF float64
	)

	if unsigned {
		maxF = math.Pow(10, float64(order-precision)) - 1 //nolint:mnd
	} else {
		maxF = math.Pow(10, float64(order-precision-1)) - 1 //nolint:mnd
	}

	output := gofakeit.Float64Range(minF, maxF)
	if unsigned {
		return output
	}

	return output * -1
}

// UnsignedDouble returns random unsigned double within the given range.
func UnsignedDouble(order, precision int) float64 {
	var (
		minF = math.SmallestNonzeroFloat64
		maxF = math.Pow(10, float64(order-precision)) - 1 //nolint:mnd
	)

	return gofakeit.Float64Range(minF, maxF)
}

// Real returns random double within the given range.
func Real(order, precision int) float64 {
	return Double(order, precision)
}

// UnsignedReal returns random unsigned double within the given range.
func UnsignedReal(order, precision int) float64 {
	return UnsignedDouble(order, precision)
}

// Bit returns random bit.
func Bit(order int) string {
	var sb strings.Builder

	sb.WriteString("b'")

	for i := 0; i < order; i++ {
		bit := "0"
		if Boolean() {
			bit = "1"
		}
		sb.WriteString(bit)
	}

	sb.WriteString("'")

	return sb.String()
}

// Date returns random date.
func Date() string {
	// https://dev.mysql.com/doc/refman/8.0/ja/datetime.html
	minT := time.Date(1000, 1, 1, 0, 0, 0, 0, time.UTC)
	maxT := time.Date(9999, 12, 31, 0, 0, 0, 0, time.UTC)
	randTime := gofakeit.DateRange(minT, maxT)

	y, m, d := randTime.Date()
	return fmt.Sprintf("%d-%02d-%02d", y, int(m), d)
}

// DateTime returns random datetime.
func DateTime() string {
	// https://dev.mysql.com/doc/refman/8.0/ja/datetime.html
	minT := time.Date(1000, 1, 1, 0, 0, 0, 0, time.UTC)
	maxT := time.Date(9999, 12, 31, 0, 0, 0, 0, time.UTC)
	randTime := gofakeit.DateRange(minT, maxT)

	y, m, d := randTime.Date()
	hour := randTime.Hour()
	minu := randTime.Minute()
	sec := randTime.Second()
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", y, m, d, hour, minu, sec)
}

// Timestamp returns random timestamp.
func Timestamp() string {
	// https://dev.mysql.com/doc/refman/8.0/ja/datetime.html
	minT := time.Date(1970, 1, 1, 0, 0, 1, 0, time.UTC)
	maxT := time.Date(2038, 1, 19, 3, 14, 7, 0, time.UTC)
	randTime := gofakeit.DateRange(minT, maxT)

	y, m, d := randTime.Date()
	h := randTime.Hour()
	mi := randTime.Minute()
	s := randTime.Second()
	l := randTime.Location()
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d %v", y, m, d, h, mi, s, l)
}

// Time returns random time.
func Time() string {
	// https://dev.mysql.com/doc/refman/8.0/ja/datetime.html
	minT := time.Date(1970, 1, 1, 0, 0, 1, 0, time.UTC)
	maxT := time.Date(2038, 1, 19, 3, 14, 7, 0, time.UTC)
	randTime := gofakeit.DateRange(minT, maxT)

	h := randTime.Hour()
	mi := randTime.Minute()
	s := randTime.Second()
	return fmt.Sprintf("%02d:%02d:%02d", h, mi, s)
}

// Year4 returns random year(4).
func Year4() string {
	// https://dev.mysql.com/doc/refman/8.0/ja/year.html
	minY := 1901
	maxY := 2155
	return fmt.Sprint(gofakeit.Number(minY, maxY))
}

// Year2 returns random year(2).
func Year2() string {
	// https://dev.mysql.com/doc/refman/8.0/ja/year.html
	minY := 0
	maxY := 99
	return fmt.Sprint(gofakeit.Number(minY, maxY))
}

// Char returns random char with the given length.
func Char(length int) string {
	if length < 0 {
		return ""
	}

	return gofakeit.LetterN(uint(length))
}

// VarChar returns random varchar with the given length.
func VarChar(length int) string {
	if length < 0 {
		return ""
	}

	return gofakeit.LetterN(uint(length))
}

// Binary returns random binary with the given length.
func Binary(length int) string {
	if length < 0 {
		return ""
	}

	return gofakeit.LetterN(uint(length))
}

// VarBinary returns random varbinary with the given length.
func VarBinary(length int) string {
	if length < 0 {
		return ""
	}

	return gofakeit.LetterN(uint(length))
}

// TinyBlob returns random tiny blob with the given length.
func TinyBlob(length int) string {
	if length < 0 {
		return ""
	}

	return gofakeit.LetterN(uint(length))
}

// TinyText returns random tiny text with the given length.
func TinyText(length int) string {
	if length < 0 {
		return ""
	}

	return gofakeit.LetterN(uint(length))
}

// Blob returns random blob with the given length.
func Blob(length int) string {
	if length < 0 {
		return ""
	}

	return gofakeit.LetterN(uint(length))
}

// Text returns random text with the given length.
func Text(length int) string {
	if length < 0 {
		return ""
	}

	return gofakeit.LetterN(uint(length))
}

// MediumBlob returns random medium blob with the given length.
func MediumBlob(length int) string {
	if length < 0 {
		return ""
	}

	return gofakeit.LetterN(uint(length))
}

// MediumText returns random medium text with the given length.
func MediumText(length int) string {
	if length < 0 {
		return ""
	}

	return gofakeit.LetterN(uint(length))
}

// LongBlob returns random long blob with the given length.
func LongBlob(length int) string {
	if length < 0 {
		return ""
	}

	return gofakeit.LetterN(uint(length))
}

// LongText returns random long text with the given length.
func LongText(length int) string {
	if length < 0 {
		return ""
	}

	return gofakeit.LetterN(uint(length))
}
