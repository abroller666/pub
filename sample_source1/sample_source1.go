package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/zenwerk/jptel"
)

// Cnv2Standard 独自フォーマットから標準版フォーマットへの変換を行います
func Cnv2Standard(v interface{}) Standard {
	var standard Standard
	portal := v.(*Portal)

	// 共通
	standard.D0001 = portal.K0001                               // 受付番号ID
	standard.D0002 = portal.K0002                               // 申込日
	standard.D0003 = portal.K0003                               // 申込時刻
	standard.D0004 = portal.K0004                               // 申込者の氏名（姓）
	standard.D0005 = portal.K0005                               // 申込者の氏名（名）
	standard.D0006 = portal.K0006                               // 申込者の氏名（姓）のヨミガナ
	standard.D0007 = portal.K0007                               // 申込者の氏名（名）のヨミガナ
	standard.D0008 = portal.K0008                               // 申込者の性別
	standard.D0009 = portal.K0009                               // 申込者と契約（予定）者との関係
	standard.D0010 = portal.K0010                               // 申込者の会社名
	standard.D0011 = portal.K0011                               // 契約（予定）者の同意
	standard.D0012 = portal.K0012                               // 申込者の電話番号
	standard.D0013 = portal.K0013                               // 申込者の電話番号区分
	standard.D0014 = portal.K0014                               // 申込者のメールアドレス
	standard.D0015 = portal.K0015                               // 契約（予定）者の氏名（姓）
	standard.D0016 = portal.K0016                               // 契約（予定）者の氏名（名）
	standard.D0017 = portal.K0017                               // 契約（予定）者の氏名（姓）のヨミガナ
	standard.D0018 = portal.K0018                               // 契約（予定）者の氏名（名）のヨミガナ
	standard.D0019 = portal.K0019                               // 契約（予定）者の性別
	standard.D0020 = strings.ReplaceAll(portal.K0020, "/", "-") // 契約（予定）者の生年月日
	standard.D0021 = portal.K0021                               // 契約（予定）者の電話番号
	standard.D0022 = portal.K0022                               // 契約（予定）者の電話番号区分
	standard.D0023 = portal.K0023                               // 契約（予定）者のメールアドレス
	standard.D0024 = portal.K0024                               // 契約（予定）者の旧住所（郵便番号）
	standard.D0025 = portal.K0025                               // 契約（予定）者の旧住所（都道府県・市区町村・町名）
	standard.D0026 = portal.K0026                               // 契約（予定）者の旧住所（丁目番地号）
	standard.D0027 = portal.K0027                               // 契約（予定）者の旧住所（建物名）
	standard.D0028 = portal.K0028                               // 契約（予定）者の旧住所（棟）
	standard.D0029 = portal.K0029                               // 契約（予定）者の旧住所（部屋番号）
	standard.D0030 = portal.K0030                               // 契約（予定）者の旧住所の建物形態
	standard.D0031 = portal.K0031                               // 契約（予定）者の旧住所の入口のオートロック有無
	standard.D0032 = portal.K0032                               // 契約（予定）者の新住所（郵便番号）
	standard.D0033 = portal.K0033                               // 契約（予定）者の新住所（都道府県・市区町村・町名）
	standard.D0034 = portal.K0034                               // 契約（予定）者の新住所（丁目番地号）
	standard.D0035 = portal.K0035                               // 契約（予定）者の新住所（建物名）
	standard.D0036 = portal.K0036                               // 契約（予定）者の新住所（棟）
	standard.D0037 = portal.K0037                               // 契約（予定）者の新住所（部屋番号）
	standard.D0038 = portal.K0038                               // 契約（予定）者の新住所の建物形態
	standard.D0039 = portal.K0039                               // 契約（予定）者の新住所の建物区分
	standard.D0040 = portal.K0040                               // 契約（予定）者の新住所の入口のオートロック有無
	standard.D0041 = strings.ReplaceAll(portal.K0041, "/", "-") // 旧住所でのサービス使用停止日
	standard.D0042 = strings.ReplaceAll(portal.K0042, "/", "-") // 新住所でのサービス使用開始日
	standard.D0043 = portal.K0043                               // 請求書又は領収書等の送付先住所（郵便番号）
	standard.D0044 = portal.K0044                               // 請求書又は領収書等の送付先住所（都道府県・市区町村・町名）
	standard.D0045 = portal.K0045                               // 請求書又は領収書等の送付先住所（丁目番地号）
	standard.D0046 = portal.K0046                               // 請求書又は領収書等の送付先住所（建物名）
	standard.D0047 = portal.K0047                               // 請求書又は領収書等の送付先住所（棟）
	standard.D0048 = portal.K0048                               // 請求書又は領収書等の送付先住所（部屋番号）
	standard.D0049 = portal.K0049                               // 請求書又は領収書等の送付先の宛名（姓）
	standard.D0050 = portal.K0050                               // 請求書又は領収書等の送付先の宛名（名）
	standard.D0051 = portal.K0051                               // 請求書又は領収書等の送付先の宛名（姓）のヨミガナ
	standard.D0052 = portal.K0052                               // 請求書又は領収書等の送付先の宛名（名）のヨミガナ
	standard.D0053 = portal.K0053                               // 請求書又は領収書等の送付先の電話番号

	// 電力
	standard.D0054 = portal.K0054                               // （使用停止）法人番号
	standard.D0055 = portal.K0055                               // （使用停止）供給地点特定番号
	standard.D0056 = portal.K0056                               // （使用停止）使用停止場所の建物解体予定の有無
	standard.D0057 = strings.ReplaceAll(portal.K0057, "/", "-") // （使用停止）使用停止場所の建物解体予定年月日
	standard.D0058 = portal.K0058                               // （使用停止）使用停止場所の建物解体予定の有無
	standard.D0059 = portal.K0059                               // （使用停止）使用停止希望時間帯
	standard.D0060 = portal.K0060                               // （使用停止）最終使用分料金の精算方法
	standard.D0061 = portal.K0061                               // （使用停止）訪問希望時間帯
	standard.D0062 = portal.K0062                               // （使用停止）立会者の氏名（姓）
	standard.D0063 = portal.K0063                               // （使用停止）立会者の氏名（名）
	standard.D0064 = portal.K0064                               // （使用停止）立会者の氏名（姓）のヨミガナ
	standard.D0065 = portal.K0065                               // （使用停止）立会者の氏名（名）のヨミガナ
	standard.D0066 = portal.K0066                               // （使用停止）契約（予定）者と立会者との関係
	standard.D0067 = portal.K0067                               // （使用停止）立会者の電話番号
	standard.D0068 = portal.K0068                               // （使用停止）立会者の電話番号区分
	standard.D0069 = portal.K0069                               // （使用停止）付帯契約の有無
	standard.D0070 = portal.K0070                               // （使用停止）付帯契約の同時停止申込の有無
	standard.D0071 = portal.K0071                               // （使用開始）法人番号
	standard.D0072 = portal.K0072                               // （使用開始）供給地点特定番号
	standard.D0073 = portal.K0073                               // （使用開始）使用開始希望時間帯
	standard.D0074 = portal.K0074                               // （使用開始）利用明細の発行希望の有無
	standard.D0075 = portal.K0075                               // （使用開始）電力供給エリア
	standard.D0078 = portal.K0076                               // （使用開始）支払方法
	standard.D0082 = portal.K0077                               // （使用開始）停電周知先の住所（郵便番号）
	standard.D0083 = portal.K0078                               // （使用開始）停電周知先の住所（郵便番号）
	standard.D0084 = portal.K0079                               // （使用開始）停電周知先の住所（丁目番地号）
	standard.D0085 = portal.K0080                               // （使用開始）停電周知先の住所（建物名）
	standard.D0086 = portal.K0081                               // （使用開始）停電周知先の住所（棟）
	standard.D0087 = portal.K0082                               // （使用開始）停電周知先の住所（部屋番号）
	standard.D0088 = portal.K0083                               // （使用開始）停電周知先の宛名（姓）
	standard.D0089 = portal.K0084                               // （使用開始）停電周知先の宛名（名）
	standard.D0090 = portal.K0085                               // （使用開始）停電周知先の宛名（姓）のヨミガナ
	standard.D0091 = portal.K0086                               // （使用開始）停電周知先の宛名（名）のヨミガナ
	standard.D0092 = portal.K0087                               // （使用開始）停電周知先の電話番号

	return standard
}

// Cnv2Api 独自フォーマットからAPIサービス連携フォーマットへの変換を行います
func Cnv2Api(v interface{}) Api {

	var api Api
	portal := v.(*Portal)

	// 申請者情報
	api.NameKanji = portal.K0004 + "　" + portal.K0005
	api.NameKana = portal.K0006 + "　" + portal.K0007
	api.Year2 = portal.K0020[0:4]
	api.Month = portal.K0020[5:7]
	api.Day = portal.K0020[8:10]
	var gendar string
	if "1" == portal.K0008 {
		gendar = "男性"
	} else if "2" == portal.K0008 {
		gendar = "女性"
	} else {
		gendar = "非選択"
	}
	api.Gender = gendar
	api.PostCode = portal.K0024
	api.AddressPref = "東京都"
	api.AddressCity = "●●市"
	api.AddressNum = portal.K0025 + portal.K0026 + portal.K0027 + portal.K0028 + portal.K0029
	tel, _ := jptel.Split(portal.K0012)
	api.RenrakuTele1 = tel.AreaCode
	api.RenrakuTele2 = tel.CityCode
	api.RenrakuTele3 = tel.SubscriberCode
	api.ApplicantTelephone1 = tel.AreaCode
	api.ApplicantTelephone2 = tel.CityCode
	api.ApplicantTelephone3 = tel.SubscriberCode
	// api.ApplicantFax1 =
	// api.ApplicantFax2 =
	// api.ApplicantFax3 =
	api.Mailaddress = portal.K0014
	api.MailaddressConf = portal.K0014

	// 申請情報
	api.SeikyuY = portal.K0002[0:4]
	api.SeikyuM = portal.K0002[5:7]
	api.SeikyuD = portal.K0002[8:10]
	api.SinseiSimei = portal.K0004 + "　" + portal.K0005
	api.SinseiRenraku = portal.K0012
	api.SinJusho = portal.K0033 + portal.K0034 + portal.K0035 + portal.K0036 + portal.K0037
	api.ImaJusho = portal.K0025 + portal.K0026 + portal.K0027 + portal.K0028 + portal.K0029
	api.IdouSeiKana1 = portal.K0017
	api.IdouMeiKana1 = portal.K0018
	api.IdouSei1 = portal.K0015
	api.IdouMei1 = portal.K0016
	api.Tsuzukigara1 = "世帯主"

	// シーケンス番号
	api.SeqNo = portal.K0090

	return api
}

// HandleRequest AWSのLamdaハンドラ（mainから呼ばれる実処理）
func HandleRequest(ctx context.Context, S3Event events.S3Event) (string, error) {

	// S3バケットと出力ファイルのキー環境変数
	var BUCKET = os.Getenv("BUCKET")
	var OUTPUTBETADIR = os.Getenv("OUTPUT_BETA_DIR")
	var OUTPUTDIR = os.Getenv("OUTPUT_DIR")

	// 入力ファイル読み込み
	svc := s3.New(session.New(), &aws.Config{
		Region: aws.String(endpoints.ApNortheast1RegionID),
	})

	filename := S3Event.Records[0].S3.Object.Key
	input := &s3.GetObjectInput{
		Bucket: aws.String(S3Event.Records[0].S3.Bucket.Name),
		Key:    aws.String(filename),
	}

	goo, err := svc.GetObject(input)
	if err != nil {
		log.Print("error occured")
		return "", err
	}

	rc := goo.Body
	defer rc.Close()
	brb := new(bytes.Buffer)
	brb.ReadFrom(rc)

	jsonData, err := ioutil.ReadAll(brb)
	if err != nil {
		log.Print("error occured")
		return "", err
	}

	var portal Portal
	err = json.Unmarshal(jsonData, &portal)
	if err != nil {
		log.Print("error occured")
		return "", err
	}

	// フォーマット変換
	api := Cnv2Api(&portal)

	// サービスAPI向けファイル作成
	outputapi, err := json.MarshalIndent(api, "", "\t")
	if err != nil {
		log.Print("error occured")
		return "", err
	}
	_, err = svc.PutObject(&s3.PutObjectInput{
		Body:                 bytes.NewReader(outputapi),
		Bucket:               aws.String(BUCKET),
		Key:                  aws.String(OUTPUTDIR + "FMT_" + filepath.Base(filename)),
		ACL:                  aws.String("private"),
		ServerSideEncryption: aws.String("AES256"),
	})
	if err != nil {
		log.Print("error occured")
		return "", err
	}

	// 標準版ファイル作成
	standard := Cnv2Standard(&portal)

	outputBeta, err := json.MarshalIndent(standard, "", "\t")
	if err != nil {
		log.Print("error occured")
		return "", err
	}
	_, err = svc.PutObject(&s3.PutObjectInput{
		Body:                 bytes.NewReader(outputBeta),
		Bucket:               aws.String(BUCKET),
		Key:                  aws.String(OUTPUTBETADIR + "FMT_std_" + filepath.Base(filename)),
		ACL:                  aws.String("private"),
		ServerSideEncryption: aws.String("AES256"),
	})
	if err != nil {
		log.Print("error occured")
		return "", err
	}

	defer log.Print("normal end")
	return "normal end", nil

}

func main() {
	lambda.Start(HandleRequest)
}
