package main

import (
	"encoding/xml"
	"strconv"
)

// Api APIサービス用のJSONを格納する構造体
type Api struct {
	NameKanji           string `json:"nameKanji"`
	NameKana            string `json:"nameKana"`
	Year1               string `json:"year1"`
	Year2               string `json:"year2"`
	Month               string `json:"month"`
	Day                 string `json:"day"`
	Gender              string `json:"gender"`
	PostCode            string `json:"postCode"`
	AddressPref         string `json:"addressPref"`
	AddressCity         string `json:"addressCity"`
	AddressNum          string `json:"addressNum"`
	RenrakuTele1        string `json:"RenrakuTele1"`
	RenrakuTele2        string `json:"RenrakuTele2"`
	RenrakuTele3        string `json:"RenrakuTele3"`
	ApplicantTelephone1 string `json:"applicantTelephone1"`
	ApplicantTelephone2 string `json:"applicantTelephone2"`
	ApplicantTelephone3 string `json:"applicantTelephone3"`
	ApplicantFax1       string `json:"applicantFax1"`
	ApplicantFax2       string `json:"applicantFax2"`
	ApplicantFax3       string `json:"applicantFax3"`
	Mailaddress         string `json:"mailaddress"`
	MailaddressConf     string `json:"mailaddressConf"`
	SeikyuY             string `json:"seikyuY"`
	SeikyuM             string `json:"seikyuM"`
	SeikyuD             string `json:"seikyuD"`
	IdouY               string `json:"idouY"`
	IdouM               string `json:"idouM"`
	IdouD               string `json:"idouD"`
	SinseiSimei         string `json:"sinseiSimei"`
	SinseiRenraku       string `json:"sinseiRenraku"`
	SinJushoKana        string `json:"sinJushoKana"`
	SinJusho            string `json:"sinJusho"`
	ImaJusho            string `json:"imaJusho"`
	SetaiNushi          string `json:"setaiNushi"`
	IdouSeiKana1        string `json:"idouSeiKana1"`
	IdouMeiKana1        string `json:"idouMeiKana1"`
	IdouSei1            string `json:"idouSei1"`
	IdouMei1            string `json:"idouMei1"`
	SeinenGengou1       string `json:"seinenGengou1"`
	SeinenY1            string `json:"seinenY1"`
	SeinenM1            string `json:"seinenM1"`
	SeinenD1            string `json:"seinenD1"`
	Seibetsu1           string `json:"seibetsu1"`
	Tsuzukigara1        string `json:"tsuzukigara1"`
	IdouSeiKana2        string `json:"idouSeiKana2"`
	IdouMeiKana2        string `json:"idouMeiKana2"`
	IdouSei2            string `json:"idouSei2"`
	IdouMei2            string `json:"idouMei2"`
	SeinenGengou2       string `json:"seinenGengou2"`
	SeinenY2            string `json:"seinenY2"`
	SeinenM2            string `json:"seinenM2"`
	SeinenD2            string `json:"seinenD2"`
	Seibetsu2           string `json:"seibetsu2"`
	Tsuzukigara2        string `json:"tsuzukigara2"`
	IdouSeiKana3        string `json:"idouSeiKana3"`
	IdouMeiKana3        string `json:"idouMeiKana3"`
	IdouSei3            string `json:"idouSei3"`
	IdouMei3            string `json:"idouMei3"`
	SeinenGengou3       string `json:"seinenGengou3"`
	SeinenY3            string `json:"seinenY3"`
	SeinenM3            string `json:"seinenM3"`
	SeinenD3            string `json:"seinenD3"`
	Seibetsu3           string `json:"seibetsu3"`
	Tsuzukigara3        string `json:"tsuzukigara3"`
	IdouSeiKana4        string `json:"idouSeiKana4"`
	IdouMeiKana4        string `json:"idouMeiKana4"`
	IdouSei4            string `json:"idouSei4"`
	IdouMei4            string `json:"idouMei4"`
	SeinenGengou4       string `json:"seinenGengou4"`
	SeinenY4            string `json:"seinenY4"`
	SeinenM4            string `json:"seinenM4"`
	SeinenD4            string `json:"seinenD4"`
	Seibetsu4           string `json:"seibetsu4"`
	Tsuzukigara4        string `json:"tsuzukigara4"`
}

// SinseiData 申請者情報XML
type SinseiData struct {
	XMLName     xml.Name   `xml:"Data"`
	FormSubject string     `xml:"form-subject,attr"`
	FormItems   []FormItem `xml:"FormItems>FormItem"`
}

// FormItem 申請者情報XML
type FormItem struct {
	ID        string     `xml:"id,attr"`
	Required  string     `xml:"required,attr"`
	Subject   string     `xml:"Subject"`
	Fragments []Fragment `xml:"Fragments>Fragment"`
}

// Fragment 申請者情報XML
type Fragment struct {
	ID          string   `xml:"id,attr"`
	Placeholder string   `xml:"placeholder,attr"`
	Required    string   `xml:"required,attr"`
	Type        string   `xml:"type,attr"`
	Label       string   `xml:"Label"`
	Options     *Options `xml:"Options"`
	Value       string   `xml:"value"`
}

// Options 申請者情報XML
type Options struct {
	Autogenerate string   `xml:"auto-generate,attr,omitempty"`
	Option       []string `xml:"Option,omitempty"`
}

// ApplicationData 申請情報XML
type SinseiInfoData struct {
	XMLName     xml.Name    `xml:"FocusFormSpecification"`
	FormSubject string      `xml:"form-subject,attr"`
	ID          string      `xml:"id,attr"`
	FormItems   []FormItem2 `xml:"FormItems>FormItem"`
}

// FormItem2 申請情報XML
type FormItem2 struct {
	Category    string      `xml:"category,attr"`
	ID          string      `xml:"id,attr"`
	Required    string      `xml:"required,attr"`
	StdCode     string      `xml:"std-code,attr"`
	Subject     Subject2    `xml:"Subject"`
	Description string      `xml:"Description"`
	Fragments   []Fragment2 `xml:"Fragments>Fragment"`
}

// Subject2 申請情報XML
type Subject2 struct {
	Subject           string `xml:",chardata"`
	DisplaySubject    string `xml:"display-subject,attr"`
	StdDisplaySubject string `xml:"std-display-subject,attr"`
	StdSubject        string `xml:"std-subject,attr"`
	SystemSubject     string `xml:"system-subject,attr"`
}

// Fragment2 申請情報XML
type Fragment2 struct {
	AreaID      string    `xml:"area-id,attr,omitempty"`
	FrameID     string    `xml:"frame-id,attr"`
	ID          string    `xml:"id,attr"`
	Placeholder string    `xml:"placeholder,attr"`
	Required    string    `xml:"required,attr"`
	Type        string    `xml:"type,attr"`
	UnitLabel   string    `xml:"unit-label,attr"`
	Label       Label2    `xml:"Label"`
	Comment     Comment2  `xml:"Comment"`
	Options     *Options2 `xml:"Options"`
	Value       string    `xml:"value"`
}

// Label2 申請情報XML
type Label2 struct {
	Label        string `xml:",chardata"`
	DisplayLabel string `xml:"display-label,attr"`
	SystemLabel  string `xml:"system-label,attr"`
}

// Comment2 申請情報XML
type Comment2 struct {
	CData string `xml:",cdata"`
}

// Options2 申請情報XML
type Options2 struct {
	Autogenerate string    `xml:"auto-generate,attr,omitempty"`
	Option       []Option2 `xml:"Option,omitempty"`
}

// Option2 申請情報XML
type Option2 struct {
	Option     string `xml:",chardata"`
	CheckBoxID string `xml:"checkbox-id,attr"`
	ID         string `xml:"id,attr"`
}

// ResBody APIのレスポンス(JSON形式のルート)
type ResBody struct {
	MetaData MetaData `json:"metadata"`
	Links    Links    `json:"_links"`
	Result   Result   `json:"result"`
	Errors   []Error  `json:"errors"`
}

// MetaData APIのレスポンスの一部
type MetaData struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

// Links APIのレスポンスの一部
type Links struct {
	Self      Self      `json:"self"`
	Reference Reference `json:"reference"`
}

// Self APIのレスポンスの一部
type Self struct {
	Href string `json:"href"`
}

// Reference APIのレスポンスの一部
type Reference struct {
	Href string `json:"href"`
}

// Result APIのレスポンスの一部
type Result struct {
	AccessKey       string         `json:"access_key"`
	TempNum         string         `json:"temporary_reference_number"`
	RefNum          string         `json:"reference_number"`
	Finished        string         `json:"finished"`
	StatusCode      string         `json:"status_code"`
	FileSignature   string         `json:"file_for_signature"`
	Copy            string         `json:"customers_copy"`
	Notifications   []Notification `json:"notification"`
	CityServiceCode string         `json:"city_service_code"`
	ReceptionDate   string         `json:"reception_date"`
}

// Error APIのレスポンスの一部
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Notification APIのレスポンスの一部
type Notification struct {
	Message string `json:"message"`
}

func newSinseiDataStruct() *SinseiData {
	data := &SinseiData{
		FormSubject: "申請者情報入力",
		FormItems: []FormItem{
			{
				ID:       "Sub_nameKanji",
				Required: "yes",
				Subject:  "氏名（漢字又はアルファベット）",
				Fragments: []Fragment{
					{
						ID:          "nameKanji",
						Placeholder: "（例）山田 花子、ＪＯＨＮ　ＳＭＩＴＨ",
						Required:    "yes",
						Type:        "text",
						Label:       "氏名（漢字又はアルファベット）",
					},
				},
			},
			{
				ID:       "Sub_nameKana",
				Required: "yes",
				Subject:  "氏名（フリガナ）",
				Fragments: []Fragment{
					{
						ID:          "nameKana",
						Placeholder: "（例）ヤマダ ハナコ",
						Required:    "yes",
						Type:        "text",
						Label:       "氏名（フリガナ）",
					},
				},
			},
			{
				ID:       "Sub_BirthDay",
				Required: "yes",
				Subject:  "生年月日",
				Fragments: []Fragment{
					{
						ID:       "year1",
						Required: "no",
						Type:     "select",
						Label:    "和暦",
						Options: &Options{
							Autogenerate: "year_range:0:-125",
						},
					},
					{
						ID:       "year2",
						Required: "yes",
						Type:     "select",
						Label:    "西暦",
						Options: &Options{
							Autogenerate: "year_range:0:-125",
						},
					},
					{
						ID:       "month",
						Required: "yes",
						Type:     "select",
						Label:    "月",
						Options: &Options{
							Autogenerate: "month",
						},
					},
					{
						ID:       "day",
						Required: "yes",
						Type:     "select",
						Label:    "日",
						Options: &Options{
							Autogenerate: "day",
						},
					},
				},
			},
			{
				ID:       "Sub_female",
				Required: "yes",
				Subject:  "性別",
				Fragments: []Fragment{
					{
						ID:       "gender",
						Required: "yes",
						Type:     "radio",
						Label:    "性別",
						Options: &Options{
							Option: []string{"男性", "女性", "非選択"},
						},
					},
				},
			},
			{
				ID:       "Sub_postCode",
				Required: "yes",
				Subject:  "郵便番号",
				Fragments: []Fragment{
					{
						ID:       "postCode",
						Required: "yes",
						Type:     "text",
						Label:    "郵便番号",
					},
				},
			},
			{
				ID:       "Sub_Address",
				Required: "yes",
				Subject:  "現住所",
				Fragments: []Fragment{
					{
						ID:          "addressPref",
						Placeholder: "（例）東京都",
						Required:    "yes",
						Type:        "text",
						Label:       "都道府県",
					},
					{
						ID:          "addressCity",
						Placeholder: "（例）新宿区",
						Required:    "yes",
						Type:        "text",
						Label:       "市町村",
					},
					{
						ID:          "addressNum",
						Placeholder: "（例）市谷ｘｘｘ－ｘｘｘ",
						Required:    "yes",
						Type:        "text",
						Label:       "番地以下（建物も含む）",
					},
				},
			},
			{
				ID:       "Sub_RenrakuTele",
				Required: "no",
				Subject:  "電話番号（連絡先）",
				Fragments: []Fragment{
					{
						ID:       "RenrakuTele1",
						Required: "no",
						Type:     "tel",
						Label:    "市外局番",
					},
					{
						ID:       "RenrakuTele2",
						Required: "no",
						Type:     "tel",
						Label:    "市内局番",
					},
					{
						ID:       "RenrakuTele3",
						Required: "no",
						Type:     "tel",
						Label:    "加入者番号",
					},
				},
			},
			{
				ID:       "Sub_applicantTelephone",
				Required: "no",
				Subject:  "電話番号",
				Fragments: []Fragment{
					{
						ID:       "applicantTelephone1",
						Required: "no",
						Type:     "tel",
						Label:    "市外局番",
					},
					{
						ID:       "applicantTelephone2",
						Required: "no",
						Type:     "tel",
						Label:    "市内局番",
					},
					{
						ID:       "applicantTelephone3",
						Required: "no",
						Type:     "tel",
						Label:    "加入者番号",
					},
				},
			},
			{
				ID:       "Sub_Fax",
				Required: "no",
				Subject:  "FAX番号",
				Fragments: []Fragment{
					{
						ID:       "applicantFax1",
						Required: "no",
						Type:     "tel",
						Label:    "Fax市外局番",
					},
					{
						ID:       "applicantFax2",
						Required: "no",
						Type:     "tel",
						Label:    "Fax市内局番",
					},
					{
						ID:       "applicantFax3",
						Required: "no",
						Type:     "tel",
						Label:    "Fax加入者番号",
					},
				},
			},
			{
				ID:       "Sub_mailaddress",
				Required: "no",
				Subject:  "メールアドレス",
				Fragments: []Fragment{
					{
						ID:       "mailaddress",
						Required: "no",
						Type:     "text",
						Label:    "メールアドレス",
					},
					{
						ID:       "mailaddressConf",
						Required: "no",
						Type:     "text",
						Label:    "メールアドレス(確認)",
					},
				},
			},
		},
	}

	return data
}

func newSinseiInfoDataStruct() *SinseiInfoData {
	data := &SinseiInfoData{
		FormSubject: "住民票の写し等の交付請求",
		ID:          "",
		FormItems: []FormItem2{
			{
				ID:       "",
				Required: "no",
				Subject: Subject2{
					Subject:        "請求年月日",
					DisplaySubject: "請求年月日",
				},
				Fragments: []Fragment2{
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "select",
						Label: Label2{
							Label: "年",
						},
						Comment: Comment2{
							CData: "年",
						},
						Options: &Options2{
							Autogenerate: "range:1:100",
						},
					},
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "select",
						Label: Label2{
							Label: "月",
						},
						Comment: Comment2{
							CData: "月",
						},
						Options: &Options2{
							Autogenerate: "month",
						},
					},
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "select",
						Label: Label2{
							Label: "日",
						},
						Comment: Comment2{
							CData: "日",
						},
						Options: &Options2{
							Autogenerate: "day",
						},
					},
				},
			},
			{
				Category: "※新しい住所に住み始めた日または予定の日を記入してください",
				ID:       "",
				Required: "no",
				Subject: Subject2{
					Subject:        "異動年月日",
					DisplaySubject: "異動年月日",
				},
				Fragments: []Fragment2{
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "select",
						Label: Label2{
							Label: "年",
						},
						Comment: Comment2{
							CData: "年",
						},
						Options: &Options2{
							Autogenerate: "range:1:100",
						},
					},
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "select",
						Label: Label2{
							Label: "月",
						},
						Comment: Comment2{
							CData: "月",
						},
						Options: &Options2{
							Autogenerate: "month",
						},
					},
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "select",
						Label: Label2{
							Label: "日",
						},
						Comment: Comment2{
							CData: "日",
						},
						Options: &Options2{
							Autogenerate: "day",
						},
					},
				},
			},
			{
				Category: "※異動する本人を記入してください",
				ID:       "",
				Required: "no",
				Subject: Subject2{
					Subject:        "申請者氏名",
					DisplaySubject: "申請者氏名",
				},
				Fragments: []Fragment2{
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "text",
						Label: Label2{
							Label: "申請者氏名",
						},
						Comment: Comment2{
							CData: "申請者氏名",
						},
					},
				},
			},
			{
				Category: "※日中に連絡が取れる電話番号をハイフン付で入力してください",
				ID:       "",
				Required: "no",
				Subject: Subject2{
					Subject:        "申請者連絡先",
					DisplaySubject: "申請者連絡先",
				},
				Fragments: []Fragment2{
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "text",
						Label: Label2{
							Label: "申請者連絡先",
						},
						Comment: Comment2{
							CData: "申請者連絡先",
						},
					},
				},
			},
			{
				ID:       "",
				Required: "no",
				Subject: Subject2{
					Subject:        "新しい住所",
					DisplaySubject: "新しい住所",
				},
				Fragments: []Fragment2{
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "text",
						Label: Label2{
							Label: "フリガナ",
						},
						Comment: Comment2{
							CData: "フリガナ",
						},
					},
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "text",
						Label: Label2{
							Label: "住所",
						},
						Comment: Comment2{
							CData: "住所",
						},
					},
				},
			},
			{
				ID:       "",
				Required: "no",
				Subject: Subject2{
					Subject:        "いままでの住所",
					DisplaySubject: "いままでの住所",
				},
				Fragments: []Fragment2{
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "text",
						Label: Label2{
							Label: "今までの住所",
						},
						Comment: Comment2{
							CData: "今までの住所",
						},
					},
				},
			},
			{
				ID:       "",
				Required: "no",
				Subject: Subject2{
					Subject:        "世帯主",
					DisplaySubject: "世帯主",
				},
				Fragments: []Fragment2{
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "text",
						Label: Label2{
							Label: "世帯主",
						},
						Comment: Comment2{
							CData: "世帯主",
						},
					},
				},
			},
			{
				Category: "※引越しする方全員を記入してください",
				ID:       "a",
				Required: "no",
				Subject: Subject2{
					Subject:        "異動者　1",
					DisplaySubject: "異動者　1",
				},
				Fragments: []Fragment2{
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "text",
						Label: Label2{
							Label: "氏（フリガナ）",
						},
						Comment: Comment2{
							CData: "氏（フリガナ）",
						},
					},
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "text",
						Label: Label2{
							Label: "名（フリガナ）",
						},
						Comment: Comment2{
							CData: "名（フリガナ）",
						},
					},
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "text",
						Label: Label2{
							Label: "氏",
						},
						Comment: Comment2{
							CData: "氏",
						},
					},
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "text",
						Label: Label2{
							Label: "名",
						},
						Comment: Comment2{
							CData: "名",
						},
					},
					{
						ID:       "",
						Required: "no",
						Type:     "radio",
						Label: Label2{
							Label: "元号（生年月日）",
						},
						Comment: Comment2{
							CData: "元号（生年月日）",
						},
						Options: &Options2{
							Option: []Option2{
								{
									Option:     "明",
									CheckBoxID: "",
									ID:         "",
								},
								{
									Option:     "大",
									CheckBoxID: "",
									ID:         "",
								},
								{
									Option:     "昭",
									CheckBoxID: "",
									ID:         "",
								},
								{
									Option:     "平",
									CheckBoxID: "",
									ID:         "",
								},
								{
									Option:     "令",
									CheckBoxID: "",
									ID:         "",
								},
							},
						},
					},
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "select",
						Label: Label2{
							Label: "年（生年月日）",
						},
						Comment: Comment2{
							CData: "年（生年月日）",
						},
						Options: &Options2{
							Autogenerate: "range:1:100",
						},
					},
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "select",
						Label: Label2{
							Label: "月（生年月日）",
						},
						Comment: Comment2{
							CData: "月（生年月日）",
						},
						Options: &Options2{
							Autogenerate: "month",
						},
					},
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "select",
						Label: Label2{
							Label: "日（生年月日）",
						},
						Comment: Comment2{
							CData: "日（生年月日）",
						},
						Options: &Options2{
							Autogenerate: "day",
						},
					},
					{
						ID:       "",
						Required: "no",
						Type:     "radio",
						Label: Label2{
							Label: "性別",
						},
						Comment: Comment2{
							CData: "性別",
						},
						Options: &Options2{
							Option: []Option2{
								{
									Option:     "男",
									CheckBoxID: "",
									ID:         "",
								},
								{
									Option:     "女",
									CheckBoxID: "",
									ID:         "",
								},
							},
						},
					},
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "text",
						Label: Label2{
							Label: "続柄",
						},
						Comment: Comment2{
							CData: "続柄",
						},
					},
				},
			},
			{
				ID:       "",
				Required: "no",
				Subject: Subject2{
					Subject:        "異動者　2",
					DisplaySubject: "異動者　2",
				},
				Fragments: []Fragment2{
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "text",
						Label: Label2{
							Label: "氏（フリガナ）",
						},
						Comment: Comment2{
							CData: "氏（フリガナ）",
						},
					},
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "text",
						Label: Label2{
							Label: "名（フリガナ）",
						},
						Comment: Comment2{
							CData: "名（フリガナ）",
						},
					},
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "text",
						Label: Label2{
							Label: "氏",
						},
						Comment: Comment2{
							CData: "氏",
						},
					},
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "text",
						Label: Label2{
							Label: "名",
						},
						Comment: Comment2{
							CData: "名",
						},
					},
					{
						ID:       "",
						Required: "no",
						Type:     "radio",
						Label: Label2{
							Label: "元号（生年月日）",
						},
						Comment: Comment2{
							CData: "元号（生年月日）",
						},
						Options: &Options2{
							Option: []Option2{
								{
									Option:     "明",
									CheckBoxID: "",
									ID:         "",
								},
								{
									Option:     "大",
									CheckBoxID: "",
									ID:         "",
								},
								{
									Option:     "昭",
									CheckBoxID: "",
									ID:         "",
								},
								{
									Option:     "平",
									CheckBoxID: "",
									ID:         "",
								},
								{
									Option:     "令",
									CheckBoxID: "",
									ID:         "",
								},
							},
						},
					},
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "select",
						Label: Label2{
							Label: "年（生年月日）",
						},
						Comment: Comment2{
							CData: "年（生年月日）",
						},
						Options: &Options2{
							Autogenerate: "range:1:100",
						},
					},
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "select",
						Label: Label2{
							Label: "月（生年月日）",
						},
						Comment: Comment2{
							CData: "月（生年月日）",
						},
						Options: &Options2{
							Autogenerate: "month",
						},
					},
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "select",
						Label: Label2{
							Label: "日（生年月日）",
						},
						Comment: Comment2{
							CData: "日（生年月日）",
						},
						Options: &Options2{
							Autogenerate: "day",
						},
					},
					{
						ID:       "",
						Required: "no",
						Type:     "radio",
						Label: Label2{
							Label: "性別",
						},
						Comment: Comment2{
							CData: "性別",
						},
						Options: &Options2{
							Option: []Option2{
								{
									Option:     "男",
									CheckBoxID: "",
									ID:         "",
								},
								{
									Option:     "女",
									CheckBoxID: "",
									ID:         "",
								},
							},
						},
					},
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "text",
						Label: Label2{
							Label: "続柄",
						},
						Comment: Comment2{
							CData: "続柄",
						},
					},
				},
			},
			{
				ID:       "",
				Required: "no",
				Subject: Subject2{
					Subject:        "異動者　3",
					DisplaySubject: "異動者　3",
				},
				Fragments: []Fragment2{
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "text",
						Label: Label2{
							Label: "氏（フリガナ）",
						},
						Comment: Comment2{
							CData: "氏（フリガナ）",
						},
					},
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "text",
						Label: Label2{
							Label: "名（フリガナ）",
						},
						Comment: Comment2{
							CData: "名（フリガナ）",
						},
					},
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "text",
						Label: Label2{
							Label: "氏",
						},
						Comment: Comment2{
							CData: "氏",
						},
					},
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "text",
						Label: Label2{
							Label: "名",
						},
						Comment: Comment2{
							CData: "名",
						},
					},
					{
						ID:       "",
						Required: "no",
						Type:     "radio",
						Label: Label2{
							Label: "元号（生年月日）",
						},
						Comment: Comment2{
							CData: "元号（生年月日）",
						},
						Options: &Options2{
							Option: []Option2{
								{
									Option:     "明",
									CheckBoxID: "",
									ID:         "",
								},
								{
									Option:     "大",
									CheckBoxID: "",
									ID:         "",
								},
								{
									Option:     "昭",
									CheckBoxID: "",
									ID:         "",
								},
								{
									Option:     "平",
									CheckBoxID: "",
									ID:         "",
								},
								{
									Option:     "令",
									CheckBoxID: "",
									ID:         "",
								},
							},
						},
					},
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "select",
						Label: Label2{
							Label: "年（生年月日）",
						},
						Comment: Comment2{
							CData: "年（生年月日）",
						},
						Options: &Options2{
							Autogenerate: "range:1:100",
						},
					},
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "select",
						Label: Label2{
							Label: "月（生年月日）",
						},
						Comment: Comment2{
							CData: "月（生年月日）",
						},
						Options: &Options2{
							Autogenerate: "month",
						},
					},
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "select",
						Label: Label2{
							Label: "日（生年月日）",
						},
						Comment: Comment2{
							CData: "日（生年月日）",
						},
						Options: &Options2{
							Autogenerate: "day",
						},
					},
					{
						ID:       "",
						Required: "no",
						Type:     "radio",
						Label: Label2{
							Label: "性別",
						},
						Comment: Comment2{
							CData: "性別",
						},
						Options: &Options2{
							Option: []Option2{
								{
									Option:     "男",
									CheckBoxID: "",
									ID:         "",
								},
								{
									Option:     "女",
									CheckBoxID: "",
									ID:         "",
								},
							},
						},
					},
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "text",
						Label: Label2{
							Label: "続柄",
						},
						Comment: Comment2{
							CData: "続柄",
						},
					},
				},
			},
			{
				ID:       "",
				Required: "no",
				Subject: Subject2{
					Subject:        "異動者　4",
					DisplaySubject: "異動者　4",
				},
				Fragments: []Fragment2{
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "text",
						Label: Label2{
							Label: "氏（フリガナ）",
						},
						Comment: Comment2{
							CData: "氏（フリガナ）",
						},
					},
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "text",
						Label: Label2{
							Label: "名（フリガナ）",
						},
						Comment: Comment2{
							CData: "名（フリガナ）",
						},
					},
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "text",
						Label: Label2{
							Label: "氏",
						},
						Comment: Comment2{
							CData: "氏",
						},
					},
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "text",
						Label: Label2{
							Label: "名",
						},
						Comment: Comment2{
							CData: "名",
						},
					},
					{
						ID:       "",
						Required: "no",
						Type:     "radio",
						Label: Label2{
							Label: "元号（生年月日）",
						},
						Comment: Comment2{
							CData: "元号（生年月日）",
						},
						Options: &Options2{
							Option: []Option2{
								{
									Option:     "明",
									CheckBoxID: "",
									ID:         "",
								},
								{
									Option:     "大",
									CheckBoxID: "",
									ID:         "",
								},
								{
									Option:     "昭",
									CheckBoxID: "",
									ID:         "",
								},
								{
									Option:     "平",
									CheckBoxID: "",
									ID:         "",
								},
								{
									Option:     "令",
									CheckBoxID: "",
									ID:         "",
								},
							},
						},
					},
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "select",
						Label: Label2{
							Label: "年（生年月日）",
						},
						Comment: Comment2{
							CData: "年（生年月日）",
						},
						Options: &Options2{
							Autogenerate: "range:1:100",
						},
					},
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "select",
						Label: Label2{
							Label: "月（生年月日）",
						},
						Comment: Comment2{
							CData: "月（生年月日）",
						},
						Options: &Options2{
							Autogenerate: "month",
						},
					},
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "select",
						Label: Label2{
							Label: "日（生年月日）",
						},
						Comment: Comment2{
							CData: "日（生年月日）",
						},
						Options: &Options2{
							Autogenerate: "day",
						},
					},
					{
						ID:       "",
						Required: "no",
						Type:     "radio",
						Label: Label2{
							Label: "性別",
						},
						Comment: Comment2{
							CData: "性別",
						},
						Options: &Options2{
							Option: []Option2{
								{
									Option:     "男",
									CheckBoxID: "",
									ID:         "",
								},
								{
									Option:     "女",
									CheckBoxID: "",
									ID:         "",
								},
							},
						},
					},
					{
						AreaID:   "",
						ID:       "",
						Required: "no",
						Type:     "text",
						Label: Label2{
							Label: "続柄",
						},
						Comment: Comment2{
							CData: "続柄",
						},
					},
				},
			},
		},
	}

	return data
}

func toWareki(years string) (wareki string) {

	var yeari, _ = strconv.Atoi(years)
	var wyear int

	if len(years) < 4 {
		return years
	}
	if yeari < 1868 {
		wyear = 1
	} else if yeari <= 1911 {
		wyear = yeari - 1867
	} else if yeari <= 1925 {
		wyear = yeari - 1911
	} else if yeari <= 1988 {
		wyear = yeari - 1925
	} else if yeari <= 2018 {
		wyear = yeari - 1988
	} else {
		wyear = yeari - 2018
	}

	return strconv.Itoa(wyear)

}

func zeroSup(arg string) (ret string) {
	if len(arg) == 0 {
		return arg
	}
	i, _ := strconv.Atoi(arg)
	return strconv.Itoa(i)
}

func cnvGender(arg string) (ret string) {

	var retStr string
	if len(arg) == 0 {
		return arg
	}

	if arg == "男性" {
		retStr = "男"
	} else if arg == "女性" {
		retStr = "女"
	}
	return retStr
}

func cnvGengo(year string) (gengou string) {

	var yeari, _ = strconv.Atoi(year)
	var gen string

	if yeari < 1868 {
		gen = ""
	} else if yeari <= 1911 {
		gen = "明"
	} else if yeari <= 1925 {
		gen = "大"
	} else if yeari <= 1988 {
		gen = "昭"
	} else if yeari <= 2018 {
		gen = "平"
	} else {
		gen = "令"
	}

	return gen
}
