package main

import (
  "fmt"
  "math"
  //"strings"
  "reflect"
)

type El struct {
    X interface{}
}

const intSize = 32 << (^uint(0) >> 63)

func main() {
    var bl bool
    var str string
    var ui8 uint8
    var bt byte
    var ui16 uint16
    var ui32 uint32
    var ui64 uint64
    var i8 int8
    var i16 int16
    var i32 int32
    var rn rune
    var i64 int64
    var f32 float32
    var f64 float64
    var cpx64 complex64
    var cpx128 complex128
    var ui uint
    var i int
    var uip uintptr
    // reflectで動的にmathを呼び出せるかチャレンジしたがダメだったのでコメントアウト ここから
    //mathObj := math // こんなことはできない
    //method := reflect.ValueOf(&mathObj).MethodByName("MinInt8") // こんなことはできない
    //var v = method.Call([]reflect.Value{}) //こんなことはできない
    //var v = math.MinInt8 // こんなことはできない
    // reflectで動的にmathを呼び出せるかチャレンジしたがダメだったのでコメントアウト ここまで
     
    names := []string{"bl","str","ui8","bt","ui16","ui32","ui64","i8","i16","i32","rn","i64","f32","f64","cpx64","cpx128","ui","i","uip"}
    elements := []El{El{bl},El{str},El{ui8},El{bt},El{ui16},El{ui32},El{ui64},El{i8},El{i16},El{i32},El{rn},El{i64},El{f32},El{f64},El{cpx64},El{cpx128},El{ui},El{i},El{uip}}
    typeReader(names, elements)
} 

func typeReader(names []string, elements []El) {
   // map定義
   maxIntMap := map[string]uint64{
      "ui8": math.MaxUint8,
      "ui16": math.MaxUint16,
      "ui32": math.MaxUint32,
      "ui64": math.MaxUint64,
      "i8": math.MaxInt8,
      "i16": math.MaxInt16,
      "i32": math.MaxInt32,
      "i64": math.MaxInt64,
    }
    maxFloatMap := map[string]float64{
      "f32": math.MaxFloat32,
      "f64": math.MaxFloat64,
    }
    maxEtcMap := map[string]string{
      "str": "stringはintの最大が文字列長の最大と考えられる",
      "bt": "byteはuint8のalias",
      "rn": "runeはint32のalias",
      "ui": "動作環境のOSに依存します。32bit OR 64bit",
      "i": "動作環境のOSに依存します。32bit OR 64bit",
      "uip": "ポインタの値を表現するに十分なサイズの符号なし整数",
      "cpx64": "複素数。実数部と虚数部として内部にそれぞれfloat32の値を持つ",
      "cpx128": "複素数。実数部と虚数部として内部にそれぞれfloat64の値を持つ",
    }
    minIntMap := map[string]int64{
      "ui8": 0,
      "ui16": 0,
      "ui32": 0,
      "ui64": 0,
      "i8": math.MinInt8,
      "i16": math.MinInt16,
      "i32": math.MinInt32,
      "i64": math.MinInt64,
    }
    minEtcMap := map[string]string{
      "str": "MINという表現は当てはまらない",
      "bt": "byteはuint8のalias",
      "rn": "runeはint32のalias",
      "ui": "動作環境のOSに依存します。32bit OR 64bit",
      "i": "動作環境のOSに依存します。32bit OR 64bit",
      "uip": "ポインタの値を表現するに十分なサイズの符号なし整数",
    }
    fmt.Printf("この実行環境のOSは[%d]bitです\n", intSize)

    for i := 0; i < len(elements); i++ {
        fmt.Printf("%-6s TYPE -> [%-10s] ", names[i], reflect.TypeOf(elements[i].X))
        fmt.Print("DEFAULT VALUE -> [")
        fmt.Print(elements[i].X)
        fmt.Print("]")
        if val, ok := maxIntMap[names[i]]; ok {
          fmt.Printf(" MAX VALUE -> [%d] MIN VALUE -> [%d]", val, minIntMap[names[i]])
        }
        if val, ok := maxFloatMap[names[i]]; ok {
          fmt.Printf(" MAX VALUE -> [%f]", val)
        }
        if val, ok := maxEtcMap[names[i]]; ok {
          fmt.Printf(" MAX VALUE -> [%s] MIN VALUE -> [%s]", val, minEtcMap[names[i]])
        }
        //文字列前方一致で場合分けする場合のコード
        //if strings.HasPrefix(names[i], "u") {
        //}
        fmt.Println("")
    }
}

