# プログラムのサンプル
## sample_source1の処理
1. 画面より連携されたJSONファイルをS3より取得
2. 取得したJSONファイルをフォーマット変換
3. 変換したファイルを申込書としてXMLファイルに変換(オンメモリでZIP圧縮)
4. 申込書を外部APIに連携

## sample_source2の処理
1. スマホアプリより署名されたファイルをS3より取得
2. 取得した電子署名ファイルを外部APIに連携
