package tempconv

import (
	"flag"
	"fmt"
	"gobook/ch2/tempconv"
)

// *celsiusFlag は clag.Value インターフェースを満足します。
type celsiusFlag struct{ tempconv.Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // エラー検査は必要ない
	switch unit {
	case "C", "°C":
		f.Celsius = tempconv.Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = tempconv.FToC(tempconv.Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag は、指定された名前、デフォルト値、使い方を持つ Celsius フラグ
// を定義しており、そのフラグ変数のアドレスを返します。
// フラグ引数は度数と単位です。たとえば "100C" です。
func CelsiusFlag(name string, value tempconv.Celsius, usage string) *tempconv.Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
