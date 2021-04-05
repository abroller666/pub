package main

import (
	"archive/zip"
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// CnvJSON2Xml 画面から連携されてきたJSON⇒APIサービス申請のフォーマット(XML)に変換
func CnvJSON2Xml(v interface{}, v2 interface{}, v3 interface{}) {
	data := v.(*SinseiData)
	data2 := v2.(*SinseiInfoData)
	api := v3.(*Api)

	// 申請者情報
	data.FormItems[0].Fragments[0].Value = api.NameKanji           // 氏名漢字
	data.FormItems[1].Fragments[0].Value = api.NameKana            // 氏名カナ
	data.FormItems[2].Fragments[1].Value = api.Year2               // 生年月日(西暦)
	data.FormItems[2].Fragments[2].Value = api.Month               // 生年月日(月)
	data.FormItems[2].Fragments[3].Value = api.Day                 // 生年月日(日)
	data.FormItems[3].Fragments[0].Value = api.Gender              // 性別
	data.FormItems[4].Fragments[0].Value = api.PostCode            // 郵便番号
	data.FormItems[5].Fragments[0].Value = api.AddressPref         // 現住所(都道府県)
	data.FormItems[5].Fragments[1].Value = api.AddressCity         // 現住所(市区町村)
	data.FormItems[5].Fragments[2].Value = api.AddressNum          // 現住所(市区町村)
	data.FormItems[6].Fragments[0].Value = api.RenrakuTele1        // 電話番号(連絡先) 市街局番
	data.FormItems[6].Fragments[1].Value = api.RenrakuTele2        // 電話番号(連絡先) 市内局番
	data.FormItems[6].Fragments[2].Value = api.RenrakuTele3        // 電話番号(連絡先) 加入者番号
	data.FormItems[7].Fragments[0].Value = api.ApplicantTelephone1 // 電話番号 市街局番
	data.FormItems[7].Fragments[1].Value = api.ApplicantTelephone2 // 電話番号 市内局番
	data.FormItems[7].Fragments[2].Value = api.ApplicantTelephone3 // 電話番号 加入者番号
	data.FormItems[8].Fragments[0].Value = api.ApplicantFax1       // FAX番号 Fax市街局番
	data.FormItems[8].Fragments[1].Value = api.ApplicantFax2       // FAX番号 Fax市内局番
	data.FormItems[8].Fragments[2].Value = api.ApplicantFax3       // FAX番号 Fax加入者番号
	data.FormItems[9].Fragments[0].Value = api.Mailaddress         // メールアドレス
	data.FormItems[9].Fragments[1].Value = api.MailaddressConf     // メールアドレス（確認用）

	// 申請情報
	data2.FormItems[0].Fragments[0].Value = toWareki(api.SeikyuY)     // 請求年月日（年）
	data2.FormItems[0].Fragments[1].Value = zeroSup(api.SeikyuM)      // 請求年月日（月）
	data2.FormItems[0].Fragments[2].Value = zeroSup(api.SeikyuD)      // 請求年月日（日）
	data2.FormItems[1].Fragments[0].Value = toWareki(api.IdouY)       // 異動年月日（年）
	data2.FormItems[1].Fragments[1].Value = zeroSup(api.IdouM)        // 異動年月日（月）
	data2.FormItems[1].Fragments[2].Value = zeroSup(api.IdouD)        // 異動年月日（日）
	data2.FormItems[2].Fragments[0].Value = api.SinseiSimei           // 申請者氏名
	data2.FormItems[3].Fragments[0].Value = api.SinseiRenraku         // 申請者連絡先
	data2.FormItems[4].Fragments[0].Value = api.SinJushoKana          // 新しい住所（フリガナ）
	data2.FormItems[4].Fragments[1].Value = api.SinJusho              // 新しい住所
	data2.FormItems[5].Fragments[0].Value = api.ImaJusho              // 今までの住所
	data2.FormItems[6].Fragments[0].Value = api.SetaiNushi            // 世帯主
	data2.FormItems[7].Fragments[0].Value = api.IdouSeiKana1          // 異動者１　氏（フリガナ）
	data2.FormItems[7].Fragments[1].Value = api.IdouMeiKana1          // 異動者１　名（フリガナ）
	data2.FormItems[7].Fragments[2].Value = api.IdouSei1              // 異動者１　氏
	data2.FormItems[7].Fragments[3].Value = api.IdouMei1              // 異動者１　名
	data2.FormItems[7].Fragments[4].Value = cnvGengo(api.SeinenY1)    // 異動者１　元号（生年月日）
	data2.FormItems[7].Fragments[5].Value = toWareki(api.SeinenY1)    // 異動者１　年（生年月日）
	data2.FormItems[7].Fragments[6].Value = zeroSup(api.SeinenM1)     // 異動者１　月（生年月日）
	data2.FormItems[7].Fragments[7].Value = zeroSup(api.SeinenD1)     // 異動者１　日（生年月日）
	data2.FormItems[7].Fragments[8].Value = cnvGender(api.Seibetsu1)  // 異動者１　性別
	data2.FormItems[7].Fragments[9].Value = api.Tsuzukigara1          // 異動者１　続柄
	data2.FormItems[8].Fragments[0].Value = api.IdouSeiKana2          // 異動者２　氏（フリガナ）
	data2.FormItems[8].Fragments[1].Value = api.IdouMeiKana2          // 異動者２　名（フリガナ）
	data2.FormItems[8].Fragments[2].Value = api.IdouSei2              // 異動者２　氏
	data2.FormItems[8].Fragments[3].Value = api.IdouMei2              // 異動者２　名
	data2.FormItems[8].Fragments[4].Value = cnvGengo(api.SeinenY2)    // 異動者２　元号（生年月日）
	data2.FormItems[8].Fragments[5].Value = toWareki(api.SeinenY2)    // 異動者２　年（生年月日）
	data2.FormItems[8].Fragments[6].Value = zeroSup(api.SeinenM2)     // 異動者２　月（生年月日）
	data2.FormItems[8].Fragments[7].Value = zeroSup(api.SeinenD2)     // 異動者２　日（生年月日）
	data2.FormItems[8].Fragments[8].Value = cnvGender(api.Seibetsu2)  // 異動者２　性別
	data2.FormItems[8].Fragments[9].Value = api.Tsuzukigara2          // 異動者２　続柄
	data2.FormItems[9].Fragments[0].Value = api.IdouSeiKana3          // 異動者３　氏（フリガナ）
	data2.FormItems[9].Fragments[1].Value = api.IdouMeiKana3          // 異動者３　名（フリガナ）
	data2.FormItems[9].Fragments[2].Value = api.IdouSei3              // 異動者３　氏
	data2.FormItems[9].Fragments[3].Value = api.IdouMei3              // 異動者３　名
	data2.FormItems[9].Fragments[4].Value = cnvGengo(api.SeinenY3)    // 異動者３　元号（生年月日）
	data2.FormItems[9].Fragments[5].Value = toWareki(api.SeinenY3)    // 異動者３　年（生年月日）
	data2.FormItems[9].Fragments[6].Value = zeroSup(api.SeinenM3)     // 異動者３　月（生年月日）
	data2.FormItems[9].Fragments[7].Value = zeroSup(api.SeinenD3)     // 異動者３　日（生年月日）
	data2.FormItems[9].Fragments[8].Value = cnvGender(api.Seibetsu3)  // 異動者３　性別
	data2.FormItems[9].Fragments[9].Value = api.Tsuzukigara3          // 異動者３　続柄
	data2.FormItems[10].Fragments[0].Value = api.IdouSeiKana4         // 異動者４　氏（フリガナ）
	data2.FormItems[10].Fragments[1].Value = api.IdouMeiKana4         // 異動者４　名（フリガナ）
	data2.FormItems[10].Fragments[2].Value = api.IdouSei4             // 異動者４　氏
	data2.FormItems[10].Fragments[3].Value = api.IdouMei4             // 異動者４　名
	data2.FormItems[10].Fragments[4].Value = cnvGengo(api.SeinenY4)   // 異動者４　元号（生年月日）
	data2.FormItems[10].Fragments[5].Value = toWareki(api.SeinenY4)   // 異動者４　年（生年月日）
	data2.FormItems[10].Fragments[6].Value = zeroSup(api.SeinenM4)    // 異動者４　月（生年月日）
	data2.FormItems[10].Fragments[7].Value = zeroSup(api.SeinenD4)    // 異動者４　日（生年月日）
	data2.FormItems[10].Fragments[8].Value = cnvGender(api.Seibetsu4) // 異動者４　性別
	data2.FormItems[10].Fragments[9].Value = api.Tsuzukigara4         // 異動者４　続柄
}

// compress zip圧縮されたバッファを作成する
func compress(applicant []byte, application []byte) (*bytes.Buffer, error) {

	// ZIP作成
	b := new(bytes.Buffer)
	w := zip.NewWriter(b)

	// 申請者情報XML
	hdr := &zip.FileHeader{}
	hdr.Name = "sinsei.xml"
	hdr.Method = zip.Deflate
	hdr.SetModTime(time.Now())
	f, err := w.CreateHeader(hdr)
	if err != nil {
		return b, err
	}
	f.Write([]byte(xml.Header + string(applicant) + "\n"))

	// 申請情報XML
	hdr = &zip.FileHeader{}
	hdr.Name = "sinsei_info.xml"
	hdr.Method = zip.Deflate
	hdr.SetModTime(time.Now())
	f, err = w.CreateHeader(hdr)
	if err != nil {
		return b, err
	}
	f.Write([]byte(xml.Header + string(application) + "\n"))

	// 添付ファイルフォルダ
	hdr = &zip.FileHeader{}
	hdr.Name = "Attachment/"
	hdr.Method = zip.Deflate
	hdr.SetModTime(time.Now())
	f, err = w.CreateHeader(hdr)
	if err != nil {
		return b, err
	}

	w.Close()

	return b, nil
}

// callAuthenticate 認証API呼び出し
func callAuthenticate(sURL string, id string, password string) (string, error) {

	values := url.Values{}
	values.Add("id", id)
	values.Add("password", password)

	req, err := http.NewRequest("POST", sURL, strings.NewReader(values.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer rsp.Body.Close()

	if 200 != rsp.StatusCode {
		log.Print("API Call Error:" + rsp.Status)
		return "", nil
	}

	var resBody ResBody
	err = json.NewDecoder(rsp.Body).Decode(&resBody)
	if err != nil {
		return "", err
	}

	accessKey := resBody.Result.AccessKey

	return accessKey, nil
}

// callApplSet データ送信API呼び出し
func callApplSet(sURL string, accessKey string, providerID string, zipData *bytes.Buffer) (string, error) {
	values := map[string]io.Reader{
		"city_service_code": strings.NewReader("xxxxxxxxxx"),
		"application_zip":   bytes.NewReader(zipData.Bytes()),
		"sign_flg":          strings.NewReader("1"),
	}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fw, err := writer.CreateFormFile("sinsei_zip", "sinsei_zip.zip")
	io.Copy(fw, values["sinsei_zip"])

	fw, err = writer.CreateFormField("service_code")
	io.Copy(fw, values["service_code"])
	fw, err = writer.CreateFormField("sign_flg")
	io.Copy(fw, values["sign_flg"])

	writer.Close()

	req, err := http.NewRequest("POST", sURL, body)

	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("X-Access-Key", accessKey)
	req.Header.Set("X-Provider-Id", providerID)

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer rsp.Body.Close()

	var resBody ResBody

	// log.Print("ボディの中")
	b, err := ioutil.ReadAll(rsp.Body)
	// fmt.Println(string(b))

	if 200 != rsp.StatusCode {

		if 400 == rsp.StatusCode {
			_ = json.NewDecoder(rsp.Body).Decode(&resBody)
			for _, e := range resBody.Errors {
				log.Print("Error Response: Code[" + e.Code + "] Message[" + e.Message + "]")
			}
		}
		return "", errors.New("API Call Error:" + rsp.Status)
	}

	jsonStr := strings.ReplaceAll(string(b), "\n", "")
	err = json.NewDecoder(strings.NewReader(jsonStr)).Decode(&resBody)
	if err != nil {
		return "", err
	}

	tempNum := resBody.Result.TempNum

	return tempNum, nil
}

// callReference 処理状況照会API呼び出し
func callReference(sURL string, accessKey string, providerID string, tempNum string) (string, string, error) {

	values := url.Values{}
	values.Add("temporary_reference_number", tempNum)

	req, err := http.NewRequest("POST", sURL, strings.NewReader(values.Encode()))
	if err != nil {
		return "", "", err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("X-Access-Key", accessKey)
	req.Header.Add("X-Provider-Id", providerID)

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", "", err
	}
	defer rsp.Body.Close()

	var resBody ResBody
	if 200 != rsp.StatusCode {

		if 400 == rsp.StatusCode {
			_ = json.NewDecoder(rsp.Body).Decode(&resBody)
			for _, e := range resBody.Errors {
				log.Print("Error Response: Code[" + e.Code + "] Message[" + e.Message)
			}
		}

		return "", "", errors.New("API Call Error:" + rsp.Status)
	}

	b, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Fatal(err)
	}

	jsonStr := strings.ReplaceAll(string(b), "\n", "")
	err = json.NewDecoder(strings.NewReader(jsonStr)).Decode(&resBody)
	if err != nil {
		return "", "", err
	}

	statusCode := resBody.Result.StatusCode
	retFile := ""

	if statusCode == "003" {
		retFile = resBody.Result.FileSignature
	} else if statusCode == "999" {
		// 入力値チェックエラーの場合はbodyの中身をログ出力
		log.Print("ボディの中")
		log.Print(string(b))
	}

	return statusCode, retFile, nil
}

// HandleRequest AWSのLamdaハンドラ（mainから呼ばれる実処理）
func HandleRequest(ctx context.Context, S3Event events.S3Event) (string, error) {

	// 環境変数
	var providerID = os.Getenv("PROVIDER_ID")                   // 民間サービス事業者ID
	var apiPasswd = os.Getenv("API_PASS")                       // パスワード
	var urlAuth = os.Getenv("URL_AUTH")                         // URL(認証API)
	var urlAppli = os.Getenv("URL_APPLI")                       // URL(データ送信API)
	var urlRef = os.Getenv("URL_REF")                           // URL(処理状況照会API)
	var BUCKET = os.Getenv("BUCKET")                            // S3のバケット
	var OUTPUTDIR = os.Getenv("OUTPUT_DIR")                     // 署名対象データを保存するS3のパス
	var MAXRETCNT, _ = strconv.Atoi(os.Getenv("MAX_RETRY_CNT")) // 処理状況照会APIのリトライ可能最大数
	var SLEEPTIME, _ = strconv.Atoi(os.Getenv("SLEEP_TIME"))    // 処理状況照会APIのリトライ間隔

	// 起動トリガーになったAPIサービス用インプットファイルを取得
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

	sourcename := strings.ReplaceAll(filepath.Base(filename), ".json", "")

	// 申請情報XMLファイル作成
	var api Api
	err = json.Unmarshal(jsonData, &api)
	if err != nil {
		log.Print("error occured")
		return "", err
	}

	var xmlDataSinsei = newSinseiDataStruct()
	var xmlDataSinseiInfo = newSinseiInfoDataStruct()

	// フォーマット変換
	CnvJSON2Xml(xmlDataSinsei, xmlDataSinseiInfo, &api)

	// XMLのバイト配列を生成 (申請者情報XML)
	outputApplicant, err := xml.MarshalIndent(xmlDataSinsei, "", "\t")
	if err != nil {
		log.Print("error occured")
		return "", err
	}

	// XMLのバイト配列を生成 (申請情報XML)
	outputApplication, err := xml.MarshalIndent(xmlDataSinseiInfo, "", "    ")
	if err != nil {
		log.Print("error occured")
		return "", err
	}

	zipBuf, err := compress(outputApplicant, outputApplication)
	if err != nil {
		log.Print("error occured")
		return "", err
	}

	// デバッグ用(APIサービスへ送るXMLの内容を確認したい場合はコメントを外す)
	// _, err = svc.PutObject(&s3.PutObjectInput{
	// 	Body:                 bytes.NewReader(zipBuf.Bytes()),
	// 	Bucket:               aws.String("moving-oss"),
	// 	Key:                  aws.String("test_data/application_zip.zip"),
	// 	ACL:                  aws.String("private"),
	// 	ServerSideEncryption: aws.String("AES256"),
	// })

	// 認証API呼び出し
	passwdHash := sha256.Sum256([]byte(apiPasswd))
	accessKey, err := callAuthenticate(urlAuth, providerID, hex.EncodeToString(passwdHash[:]))
	if err != nil {
		log.Print("error occured")
		return "", err
	}
	log.Print("アクセスキー:" + accessKey)

	// データ送信API呼び出し
	tempNum, err := callApplSet(urlAppli, accessKey, providerID, zipBuf)
	if err != nil {
		log.Print("error occured")
		return "", err
	}
	log.Print("仮受付番号:" + tempNum)

	var stsCode = ""
	var fileForSignStr = ""
	var retCnt = 0

	for {
		retCnt++
		log.Printf("%v回目", retCnt)
		if retCnt > MAXRETCNT {
			log.Print("error occured")
			return "", errors.New("retry count over")
		}
		// 処理状況照会API呼び出し
		stsCode, fileForSignStr, err = callReference(urlRef, accessKey, providerID, tempNum)
		if err != nil {
			log.Print("error occured")
			return "", err
		}

		if stsCode == "001" || stsCode == "002" {
			time.Sleep(time.Second * time.Duration(SLEEPTIME))
			log.Print("リトライします。:" + stsCode)
			continue
		} else if stsCode == "003" {
			log.Print("電子署名待ち[" + sourcename + "]")
			break
		} else if stsCode == "999" {
			log.Print("申請データチェックエラー[" + sourcename + "]")
			fileForSignStr = tempNum
			break
		} else {
			log.Print("想定外のステータス:" + stsCode + " [" + sourcename + "]")
			break
		}

	}

	// 署名対象データS3格納 ※仮受付番号はファイル名の一部として保持する
	fileForSign, _ := base64.StdEncoding.DecodeString(fileForSignStr)
	_, err = svc.PutObject(&s3.PutObjectInput{
		Body:                 bytes.NewReader(fileForSign),
		Bucket:               aws.String(BUCKET),
		Key:                  aws.String(OUTPUTDIR + "SIGN_TGT_" + sourcename + "_" + tempNum + ".dat"),
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
