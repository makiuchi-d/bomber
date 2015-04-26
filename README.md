bomber - Unicode BOM Utility
============================

```
Usage:
  bomber [OPTION]... [FILE]

Options:
  -r, --remove  BOMを削除 (default: false)
  -t, --target  対象BOMフォーマット
                8: UTF-8 BOM (0xfe, 0xbb, 0xbf) (デフォルト)
                16le: UTF-16 LE (0xff, 0xfe)
                16be: UTF-16 BE (0xfe, 0xff)
                32le: UTF-16 LE (0xff, 0xfe, 0x00, 0x00)
                32be: UTF-16 BE (0x00, 0x00, 0xfe, 0xff)
  -h, --help    ヘルプの表示
```

テキストファイルにBOMを追加または削除するコマンドです.
指定されたファイルまたは標準入力を読み込み、BOMを追加または削除した結果を標準出力に出力します.

ファイル自体のエンコーディングはチェックせず、オプションで指定した対象BOMのみ追加または削除します.
対象BOM以外であってもBOMが既に存在する場合は追加しません.
