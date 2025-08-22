# gocud

Universal Color Design set for Go

[![Go Report Card](https://goreportcard.com/badge/github.com/KatzMatz/gocud)](https://goreportcard.com/report/github.com/KatzMatz/gocud)
[![GoDoc](https://godoc.org/github.com/KatzMatz/gocud?status.svg)](https://godoc.org/github.com/KatzMatz/gocud)

## About

カラーユニバーサルデザイン推奨配色セットをGoで簡単に利用できるライブラリです。色覚の多様性に配慮したアクセシブルな色選択を提供し、誰もが見やすいグラフやチャートの作成を支援します。

## Features

- 🎨 **標準カラーパレット**: JPMAコード付きの厳選された9色
- 🔄 **複数フォーマット対応**: RGB、RGBA、CMYK色変換をサポート
- 📊 **事前定義パターン**: 様々な可視化ニーズに対応した色パターン
- ♿ **アクセシビリティ**: 色覚特性の違いに関係なく区別可能な色設計
- 🛡️ **型安全性**: 不変な色パターンで意図しない変更を防止

## Installation

```bash
go get github.com/KatzMatz/gocud
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/KatzMatz/gocud"
)

func main() {
    // 定義済み色の使用
    red := gocud.RED
    blue := gocud.BLUE
    
    // RGBA値の取得
    rgba := red.RGB(255) // 完全不透明
    fmt.Printf("赤のRGBA: %+v\n", rgba)
    
    // CMYK値の取得
    cmyk := blue.CMYK()
    fmt.Printf("青のCMYK: %+v\n", cmyk)
    
    // 色パターンの使用
    palette := gocud.Accent4Pattern()
    for i, color := range palette {
        fmt.Printf("色 %d: JPMA=%s\n", i+1, color.JpmaCode)
    }
}
```

## Available Colors

| 色名 | JPMAコード | RGB | 説明 |
|------|-----------|-----|------|
| RED | J08-50V | (255, 75, 0) | 鮮やかな赤 |
| YELLOW | J27-85V | (255, 241, 0) | 明るい黄色 |
| GREEN | J40-60V | (3, 175, 122) | 森の緑 |
| BLUE | J72-40T | (0, 90, 255) | 深い青 |
| SKY_BLUE | J69-70T | (77, 196, 255) | 空色 |
| PINK | J02-70T | (255, 128, 130) | やわらかなピンク |
| ORANGE | J15-65X | (246, 170, 0) | 明るいオレンジ |
| PURPLE | J89-40T | (153, 0, 153) | 濃い紫 |
| BROWN | J09-30H | (128, 64, 0) | 茶色 |

## Color Patterns

### AccentSet()
包括的な可視化に適した全9色を返します。

### Accent6Pattern()
チャートやグラフに最適なバランスの取れた6色パターン:
`[ORANGE, YELLOW, GREEN, BLUE, SKY_BLUE, BROWN]`

### Accent5Pattern()
中程度の複雑さのデータに適した5色パターン:
`[RED, YELLOW, GREEN, BLUE, SKY_BLUE]`

### Accent4Pattern()
シンプルな可視化に適した4色パターン:
`[RED, YELLOW, GREEN, SKY_BLUE]`


## Color Universal Design について

このライブラリはカラーユニバーサルデザイン推奨配色セットを実装しており、以下を保証します：

- **高コントラスト**: 全ての色が十分なコントラスト比を維持
- **識別可能性**: 1型色覚、2型色覚、3型色覚の方でも色の区別が可能
- **プロフェッショナル品質**: ビジネスプレゼンテーション、学術論文、公的資料に適用可能


## 出典

- [Color Universal Design Organization](https://jfly.uni-koeln.de/colorset/)
- 『カラーユニバーサルデザイン推奨配色セット ガイドブック』第２版
- 発行年：2018年
- 発行者：カラーユニバーサルデザイン推奨配色セット制作委員会

## 謝辞

- カラーユニバーサルデザイン推奨配色セット制作委員会
- 色彩標準化における日本塗料工業会（JPMA）
