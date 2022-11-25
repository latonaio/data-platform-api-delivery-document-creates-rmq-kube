# subfunc_specs
## 機能サービス名:【Delivery Document > DPFM_API_DELIVERY_DOCUMENT_SRV】

## 機能名：【data_platform_function_delivery_document_creates_subfunc】

## ＜機能概要＞
入出荷伝票登録の補助機能。

## ＜仕様詳細＞
I. 一括処理の場合
I-0. 入力パラメータ  
入出荷伝票登録は、次の通り、入力パラメータを指定して実行する。  

| Property                      | Description                  | Attribute       | Range, Plurality| 
| ----------------------------- | ---------------------------- | --------------- | ------ | 
| IssuingPlantPartnerFunction   | 出荷プラント取引先機能       | “DELIVERFROM” |  (固定値)        | 
| IssuingPlantBusinessPartner   | 出荷プラントビジネスパートナ | Mandatory       | ✔     | 
| IssuingPlant                  | 出荷プラント                 | Mandatory       | ✔     | 
| ReceivingPlantPartnerFunction | 入荷プラント取引先機能       | “DELIVERTO”   | (固定値)   | 
| ReceivingPlantBusinessPartner | 入荷プラントビジネスパートナ | Mandatory       | ✔     | 
| ReceivingPlant                | 入荷プラント                 | Mandatory       | ✔     | 
| ConfirmedDeliveryDate         | 在庫確認済納入日付           | Mandatory       | ✔     | 

I-0-1. [OrderID]が未入出荷であり、かつ、[OrderID]に入出荷伝票未登録残がある、明細の取得  

I-0-1-1. [OrderID]が未入出荷であり、かつ、[OrderID]に入出荷伝票未登録残がある、明細の取得  
オーダーヘッダから、HeaderCompleteDeliveryIsDefined=falseであり、かつ、OverallDeliveryStatus <>”CL”である、OrderIDを検索して保持する。  
処理結果がゼロ件である場合、エラーメッセージを出力して終了する。  

| Target / Processing Type       | Key(s)                                | 
| ------------------------------ | ------------------------------------- | 
| OrderID <br> / Get and Hold    | HeaderCompleteDeliveryIsDefined=false <br> OverallDeliveryStatus <>”CL” | 
| <b>Table Searched</b>                 | <b>Name of the Table</b>                     | 
| Orders Header                  | data_platform_orders_header_data      | 
| <b>Field Searched</b>                 | <b>Data Type / Number of Digits</b>          | 
| OrderID                        | int / 16                              | 
| <b>Single Record or Array</b>         | <b>Memo</b>                                  | 
| Array                          | 

I-0-2. ヘッダパートナプラントのデータ取得  
I-0-2-1. I-0-1-1で保持した[OrderID]をキーとして、オーダーヘッダから、ヘッダパートナプラントデータ[PartnerFunction, BusinessPartner, Plant]を取得する。  
処理結果がゼロ件である場合、エラーメッセージを出力して終了する。  

| Target / Processing Type    | Key(s)                                          | 
| --------------------------- | ----------------------------------------------- | 
| 下記Field Searchedと同一 <br>/ Get and Hold              | OrderID(I-0-1-1で保持)                          | 
| <b>Table Searched</b>              | <b>Name of the Table</b>                               | 
| Orders Header Partner Plant | data_platform_orders_header_partner_plant__data | 
| <b>Field Searched</b>              | <b>Data Type / Number of Digits</b>                    | 
| PartnerFunction <br> BusinessPartner <br> Plant  | string(varchar) / 40 <br> int / 12 <br> string(varchar) / 4         | 
| <b>Single Record or Array</b>      | <b>Memo</b>                                            | 
| Array                       | 	                                            |

I-0-2-2. ヘッダパートナプラントの絞り込み(入力パラメータ)  
I-0-2-1で保持したデータを、入力パラメータの絞り込み条件で絞り込み、[絞り込まれたデータ]を保持する。  
処理結果がゼロ件である場合、エラーメッセージを出力して終了する。  

| Property                                                    | 絞り込み条件         | 
| ----------------------------------------------------------- | -------------------- | 
| IssuingPlantBusinessPartner(出荷プラントビジネスパートナ)   | 入力パラメータの条件 | 
| IssuingPlant(出荷プラント)                                  | 入力パラメータの条件 | 
| ReceivingPlantBusinessPartner(入荷プラントビジネスパートナ) | 入力パラメータの条件 | 
| ReceivingPlant(入荷プラント)                                | 入力パラメータの条件 | 

I-0-3. オーダー明細納入日程行の取得  
I-0-2-2で[絞り込まれたデータ]の[オーダーID]をキーとして、オーダー明細納入日程行から、[ConfirmedDeliveryDate]が入力パラメータの条件を満たす[OrderID, OrderItem, OrderItemScheduleLine]を検索して保持する。  
処理結果がゼロ件である場合、エラーメッセージを出力して終了する。  

| Target / Processing Type | Key(s)                                       | 
| ------------------------ | -------------------------------------------- | 
| OrderID <br> OrderItem <br> OrderItemScheduleLine <br> / Get and Hold   | [ConfirmedDeliveryDate]=入力パラメータの条件 | 
| <b>Table Searched</b>           | <b>Name of the Table</b>                            | 
| Orders Header            | data_platform_orders_item_schedule_line_data | 
| <b>Field Searched</b>           | <b>Data Type / Number of Digits</b>                 | 
| OrderID <br> OrderItem <br> OrderItemScheduleLine    | int / 16 <br> int / 6 <br> int / 3                  | 
| <b>Single Record or Array</b>   | <b>Memo</b>                                         | 
| Array                    | 	

II 個別処理の場合  
II-0-1-1. OrderIDが未入出荷であり、かつ、OrderIDに入出荷伝票未登録残がある、明細の取得  
OrderIDをキーとして、オーダーヘッダから、HeaderCompleteDeliveryIsDefined=falseであり、かつ、OverallDeliveryStatus <>”CL”である、OrderIDを検索して保持する。  
処理結果がゼロ件である場合、エラーメッセージを出力して終了する。  

| Target / Processing Type       | Key(s)                                                          | 
| ------------------------------ | --------------------------------------------------------------- | 
| OrderID <br> / Get and Hold     | HeaderCompleteDeliveryIsDefined=false <br> OverallDeliveryStatus <>”CL” | 
| <b>Table Searched</b>                 | <b>Name of the Table</b>                                               | 
| Orders Header                  | data_platform_orders_header_data                                | 
| <b>Field Searched</b>                 | <b>Data Type / Number of Digits</b>                                    | 
| OrderID                        | int / 16                                                        | 
| <b>Single Record or Array</b>         | <b>Memo</b>                                                            | 
| Array                          | 実際はSingle Recordだが後続をIと共通処理とするためArrayで持つ。 |   

※以下、I(一括処理), II(個別処理) からの後続として共通処理  

1. Delivery Document Header 
次の補助機能の開発を行う。

1. オーダー参照レコード・値の取得  

1-1. オーダー参照レコード・値の取得（オーダーヘッダ）  
IまたはIIの処理結果の[OrderID]をキーとして、オーダーヘッダから、[入出荷伝票登録のために参照するレコードと値]を取得する。  

| Target / Processing Type | Key(s)                           | 
| ------------------------ | -------------------------------- | 
| 下記Field Searchedと同一 <br> / Get and Hold           | [OrderID]                        | 
| <b>Table Searched</b>           | <b>Name of the Table</b>                | 
| Orders Header            | data_platform_orders_header_data | 
| <b>Field Searched</b>           | <b>Data Type / Number of Digits</b>     | 
| OrderType <br> Buyer <br> Seller <br> ContractType <br> OrderVaridityStartDate <br> OrderVaridityEndDate <br> InvoiceScheduleStartDate <br> InvoiceScheduleEndDate <br> TransactionCurrency <br> Incoterms <br> PaymentTerms <br> IsExportImportDelivery   | string(varchar) / 3  <br> int / 12  <br> int / 12  <br> string(varchar) / 4  <br> date      <br> date      <br> date      <br> date      <br> string(varchar) / 5  <br> string(varchar) / 4 <br> string(varchar) / 4 <br> bool(tinyint) / 1        | 
| <b>Single Record or Array</b>   | <b>Memo</b>                             | 
| Array                    |                                  |

1-2. オーダー参照レコード・値の取得（オーダーヘッダパートナ）  
[OrderID]をキーとして、オーダーヘッダパートナから、[入出荷伝票登録のために参照するレコードと値]を取得する。  

| Target / Processing Type | Key(s)                                   | 
| ------------------------ | ---------------------------------------- | 
| 下記Field Searchedと同一 <br> / Get and Hold           | [OrderID]                                | 
| <b>Table Searched</b>           | <b>Name of the Table</b>                        | 
| Orders Header Partner    | data_platform_orders_header_partner_data | 
| <b>Field Searched</b>           | <b>Data Type / Number of Digits</b>             | 
| PartnerFunction <br> BusinessPartner <br> BusinessPartnerFullName <br> BusinessPartnerName <br> Organization <br> Language <br> Currency <br> ExternalDocumentID <br> AddressID  | string(varchar) / 40 <br> int / 12 <br> string(varchar) / 200 <br> string(varchar) / 100 <br> string(varchar) / 6 <br> string(varchar) / 2 <br> string(varchar) / 5 <br> string(varchar) / 40 <br> int / 12                 | 
| <b>Single Record or Array</b>   | <b>Memo</b>                                     | 
| Array                    |                                          |

1-3. オーダー参照レコード・値の取得（オーダーヘッダパートナコンタクト）  
[OrderID]をキーとして、オーダーヘッダパートナコンタクトから、[入出荷伝票登録のために参照するレコードと値]を取得する。  

| Target / Processing Type      | Key(s)                                           | 
| ----------------------------- | ------------------------------------------------ | 
| 下記Field Searchedと同一 <br> / Get and Hold                | [OrderID]                                        | 
| <b>Table Searched</b>                | <b>Name of the Table</b>                                | 
| Orders Header Partner Contact | data_platform_orders_header_partner_contact_data | 
| <b>Field Searched</b>                | <b>Data Type / Number of Digits</b>                     | 
| PartnerFunction <br> BusinessPartner <br> ContactID <br> ContactPersonName <br> EmailAddress <br> PhoneNumber <br> MobilePhoneNumber <br> FaxNumber <br> ContactTag1 <br> ContactTag2 <br> ContactTag3  <br> ContactTag4                   | string(varchar) / 40 <br> int / 12 <br>string(varchar) / 4   <br> string(varchar) / 100 <br> string(varchar) / 200 <br> string(varchar) / 100 <br> string(varchar) / 100 <br> string(varchar) / 100 <br> string(varchar) / 40  <br> string(varchar) / 40  <br> string(varchar) / 40  <br> string(varchar) / 40  |
| <b>Single Record or Array</b>        | <b>Memo</b>                                             | 
| Array                         |                                                  |

1-4. オーダー参照レコード・値の取得（オーダーヘッダパートナプラント）  
[OrderID]をキーとして、オーダーヘッダパートナプラントから、[入出荷伝票登録のために参照するレコードと値]を取得する。  

| Target / Processing Type    | Key(s)                                         | 
| --------------------------- | ---------------------------------------------- | 
| 下記Field Searchedと同一 <br> / Get and Hold              | [OrderID]                                      | 
| <b>Table Searched</b>              | <b>Name of the Table</b>                              | 
| Orders Header Partner Plant | data_platform_orders_header_partner_plant_data | 
| <b>Field Searched</b>              | <b>Data Type / Number of Digits</b>                   | 
| PartnerFunction <br> BusinessPartner <br> Plant  | string(varchar) / 40 <br> int / 12 <br> string(varchar) / 4         | 
| <b>Single Record or Array</b>      | <b>Memo</b>                                           | 
| Array                       |                                                |

1-5. オーダー参照レコード・値の取得（オーダー明細）  
[OrderID、ItemCompleteDeliveryIsDefined=false、StockConfirmationStatus<>”NP”]をキーとして、オーダー明細から、[入出荷伝票登録のために参照するレコードと値]を取得する。  

| Target / Processing Type                              | Key(s)                         | 
| ----------------------------------------------------- | ------------------------------ | 
| 下記Field Searchedと同一 <br> / Get and Hold          | OrderID<br> ItemCompleteDeliveryIsDefined=false <br> StockConfirmationStatus<>”NP”                       | 
| <b>Table Searched</b>                                        | <b>Name of the Table</b>              | 
| Orders Item                                           | data_platform_orders_item_data | 
| <b>Field Searched</b>                                        | <b>Data Type / Number of Digits</b>   | 
| OrderID <br> OrderItem <br>OrderItemCategory <br>OrderItemText <br>Product <br>ProductStandardID <br>BaseUnit <br>OrderQuantity <br>OrderQuantityUnit <br>StockConfirmationPartnerFunction <br>StockConfirmationBusinessPartner <br>StockConfirmationPlant <br>StockConfirmationPolicy <br>StockConfirmationStatus <br>ConfdOrderQtyByPDTAvailCheck <br>ItemGrossWeight <br>ItemNetWeight <br>ItemWeightUnit <br>ProductGroup <br>ProductionPlantPartnerFunction <br>ProductionPlantBusinessPartner <br>ProductionPlant <br>ProductionPlantTimeZone <br>ProductionPlantStorageLocation <br>IssuingPlantPartnerFunction <br>IssuingPlantBusinessPartner <br>IssuingPlant <br>IssuingPlantTimeZone <br>IssuingPlantStorageLocation <br>ReceivingPlantPartnerFunction <br>ReceivingPlantBusinessPartner <br>ReceivingPlant <br>ReceivingPlantTimeZone <br>ReceivingPlantStorageLocation <br>ProductIsBatchManagedInProductionPlant <br>ProductIsBatchManagedInIssuingPlant <br>ProductIsBatchManagedInReceivingPlant <br>ProductionPlantBatch <br>IssuingPlantBatch <br>ReceivingPlantBatch <br>ProductionPlantBatchValidityStartDate <br>ProductionPlantBatchValidityEndDate <br>IssuingPlantBatchValidityStartDate <br>IssuingPlantBatchValidityEndDate <br>ReceivingPlantBatchValidityStartDate <br>ReceivingPlantBatchValidityEndDate <br>Incoterms <br>BPTaxClassification <br>ProductTaxClassification <br>BPAccountAssignmentGroup <br>ProductAccountAssignmentGroup <br>PaymentTerms <br>PaymentMethod <br>Project <br>TaxCode <br>TaxRate <br> CountryOfOrigin <br>| int / 16 <br> int / 6 <br> string(varchar) / 4  <br>string(varchar) / 100 <br> string(varchar) / 40 <br> string(varchar) / 40 <br>string(varchar) / 3  <br> float / 15 <br> string(varchar) / 3  <br> string(varchar) / 40 <br>int / 12 <br>string(varchar) / 4 <br>string(varchar) / 4 <br>string(varchar) / 2 <br> float / 15 <br> float / 15 <br>float / 15 <br> string(varchar) / 3 <br>string(varchar) / 9 <br>string(varchar) / 40  <br>int 12 <br>string(varchar) / 4 <br> string(varchar) / 3 <br> string(varchar) / 4  <br> string(varchar) / 40 <br> int 12 <br> string(varchar) / 4 <br> string(varchar) / 3 <br> string(varchar) / 4  <br> string(varchar) / 40 <br> int 12 <br> string(varchar) / 4 <br>string(varchar) / 3 <br> string(varchar) / 4  <br> bool(tinyint) / 1  <br> bool(tinyint) / 1 <br> bool(tinyint) / 1  <br>string(varchar) / 10  <br>string(varchar) / 10  <br> string(varchar) / 10 <br> date <br> date <br> date <br> date <br> date <br> date <br> string(varchar) / 4 <br> string(varchar) / 1 <br> string(varchar) / 1  <br> string(varchar) / 2 <br>string(varchar) / 2 <br> string(varchar) / 4 <br>string(varchar) / 1 <br>string(varchar) / 24 <br> string(varchar) / 2 <br> float / 6 <br> string(varchar) / 3                                   | 
| <b>Single Record or Array</b>                                | <b>Memo</b>                           | 
| Array                                                 |                                | 

1-6. オーダー参照レコード・値の取得（オーダー明細納入日程行）  
1-5で取得した[OrderID、OrderItem]、[DelivBlockReasonForSchedLine=false]をキーとして、オーダー明細納入日程行から、[入出荷伝票登録のために参照するレコードと値]を取得する。  

| Target / Processing Type             | Key(s)                                        | 
| ------------------------------------ | --------------------------------------------- | 
| 下記Field Searchedと同一 <br> / Get and Hold   | [OrderID(1-5で取得) <br> OrderItem(1-5で取得)] <br> [DelivBlockReasonForSchedLine=false] | 
| <b>Table Searched</b>                       | <b>Name of the Table</b>                             | 
| Orders Item Schedule Line            | data_platform_orders_item_schedule_line_data | 
| <b>Field Searched</b>                       | <b>Data Type / Number of Digits</b>                  | 
| OrderID <br> OrderItem <br> ScheduleLine <br> ConfirmedDeliveryDate <br> ScheduleLineOrderQuantity <br> OpenConfdDelivQtyInOrdQtyUnit        | int / 16 <br> int / 6 <br> int / 3 <br> date <br> float / 15 <br> float / 15                           | 
| <b>Single Record or Array</b>               | <b>Memo</b>                                          | 
| Array                                |                                               | 


1-7. DeliveryDocument　　

| Property | Description |
| -------- | ----------- |
| DeliveryDocument | 入出荷伝票番号。自動採番される。 |

1-7-1. 入力ファイルのservice_label、当処理のProperty(=”DeliveryDocument”)をキーとして、対象のNumber Range Latest NumberのLatestNumberを検索し保持する。 

| Target / Processing Type                       | Key(s)                                        | 
| ---------------------------------------------- | --------------------------------------------- | 
| ServiceLabel FieldNameWithNumberRange <br> LatestNumber <br> / Get and Hold  | service_label <br> FieldNameWithNumberRange =“DeliveryDocument” | 
| <b>Table Searched</b>                                 | <b>Name of the Table</b>                             | 
| Number Range SQL Latest Number Data            | data_platform_number_range_latest_number_data | 
| <b>Field Searched</b>                                 | <b>Data Type / Number of Digits</b>                  | 
| ServiceLabel <br> FieldNameWithNumberRange <br> LatestNumber                                   | string(varchar) / 50 <br> string(varchar) / 100 <br> int / 10                                       | 
| <b>Single Record or Array</b>                         | <b>Memo</b>                                          | 
| Single Record                                  |                                               | 

1-7-2. 保持されたLatestNumberに1を足したものをDeliveryDocumentにセットする。  

1-8. DocumentDate  

| Property | Description |
| -------- | ----------- |
| DocumentDate | 入出荷伝票日付。データ連携時のシステム日付が設定される。 |

1-8-1. システム日付を設定する。

1-9. BillingDocumentDate  

| Property | Description |
| -------- | ----------- |
| BillingDocumentDate | 請求書日付。計画在庫出庫日付および支払条件をもとに参照・計算して請求日付が設定される。 |

1-9-1. BillingDocumentDate[請求書日付（＝請求書の締め日）]を計算するために、入力ファイルのPlannedGoodsIssueDateを保持する。  

| Target / Processing Type | Key(s)                       | 
| ------------------------ | ---------------------------- | 
| PlannedGoodsIssueDate <br> / Hold  | None               | 
| <b>Table Searched</b>           | <b>Name of the Table</b>            | 
| None                     | None                         | 
| <b>Field Searched</b>           | <b>Data Type / Number of Digits</b> | 
| None                     |                              | 
| <b>Single Record or Array</b>   | <b>Memo</b>                         | 
| Single Record            |                              | 

1-9-2. BillingDocumentDate[請求書日付（＝請求書の締め日）]を計算するために、1-1で取得したPaymentTermsをキーとして、対象の支払条件のDueDate、BaseDateCalcFixedDate、BaseDateCalcAddMonth、を検索して保持する。  

| Target / Processing Type             | Key(s)                                         | 
| ------------------------------------ | ---------------------------------------------- | 
| DueDate <br> BaseDateCalcFixedDate <br> BaseDateCalcAddMonth <br> / Get and Hold                       | PaymentTerms                                   | 
| <b>Table Searched</b>                       | <b>Name of the Table</b>                              | 
| Payment Terms SQL Payment Terms Data | data_platform_payment_terms_payment_terms_data | 
| <b>Field Searched</b>                       | <b>Data Type / Number of Digits</b>                   | 
| DueDate <br> BaseDateCalcFixedDate <br> BaseDateCalcAddMonth  | int / 2 <br> int / 2 <br> int / 2                              | 
| <b>Single Record or Array</b>               | <b>Memo</b>                                           | 
| Single Record                        |                                                | 

1-9-3. PlannedGoodsIssueDate、DueDate、BaseDateCalcFixedDate、BaseDateCalcAddMonthをもとに、請求書日付（＝請求書の締め日）を計算し、BillingDocumentDateにセットする。  
e.g.) RequestedDeliveryDateが2022-03-08、DueDateが31、BaseDateCalcFixedDateが31、BaseDateCalcAddMonthが0だった場合、請求書日付は2022-03-31と計算される。  

1-10. CreationDate  

| Property | Description | EC |
| -------- | ----------- | --- |
| CreationDate | 作成日付。自動生成される。変更不可。 |    |	

1-10-1. CreationDateに、システム日付をセットする。  

1-11. CreationTime  

| Property | Description | EC |
| -------- | ----------- | --- |
| CreationTime | 作成時刻。自動生成される。変更不可。 |     |

1-11-1. CreationDateに、システム時刻をセットする。  

1-12. HeaderBillingStatus  

| Property | Description | EC |
| -------- | ----------- | --- |
| HeaderBillingStatus | ヘッダ請求ステータス。 <br> 完全請求完了:CL <br> 部分請求完了:PP <br> 未請求:NP <br>が設定される。変更不可。 |     |

1-12-1.入力ファイルの[DeliveryDocument, DeliveryDocumentItem]をキーとして、対象の全ての請求伝票明細[InvoiceDocument, InvoiceDocumentItem]を検索して保持する。  

| Target / Processing Type  | Key(s)                                   | 
| ------------------------- | ---------------------------------------- | 
| InvoiceDocument <br> InvoiceDocumentItem <br> / Get and Hold  | DeliveryDocument <br> DeliveryDocumentItem      | 
| <b>Table Searched</b>            | <b>Name of the Table</b>                        | 
| Invoice Document SQL Item | data_platform_invoice_document_item_data | 
| <b>Field Searched</b>            | <b>Data Type / Number of Digits</b>             | 
| InvoiceDocument <br> InvoiceDocumentItem | int / 16 <br> int / 6     | 
| <b>Single Record or Array</b>    | <b>Memo</b>                                     | 
| Array                     | 

1-12-2. 検索結果が0件であった場合、HeaderBillingStatusに未請求”NP”をセットする。検索結果が1件以上であり、かつ、全ての検索結果値がfalseの場合は未請求”NP”をセットする。検索結果が1件以上であり、かつ、 全ての検索結果値がtrueの場合は完全請求完了”CL”をセットする。それ以外の場合は部分請求完了”PP”をセットする。  
<ロジックまとめ>  

| 検索結果件数 | 検索結果値 | HeaderBillingStatusにセットする値 | 
| ------------ | ---------- | --------------------------------- | 
| 0件          | -          | “NP”                            | 
| 1件以上      | -          | “NP”                            | 
| 1件以上      | -          | “CL”                            | 
| 1件以上      | -          | “PP”                            | 


＜項目＞  

| Property                      | Description                                                                                          | EC  | 
| ----------------------------- | ---------------------------------------------------------------------------------------------------- | --- | 
| CompleteDeliveryIsDefined     | 完納フラグ。入出荷伝票に対して完全に入出荷が完了しているときにtrueが設定される。変更不可。           |     | 
| OverallDeliveryStatus         | 全体入出荷ステータス。次のステータスが設定される。変更不可。 <br> NP:未入出荷 <br> PP:部分入出荷完了 <br> CL:入出荷完了                 |                                                                                                      | 
| HeaderBillgIncompletionStatus | ヘッダ請求不完全ステータス。システムにより設定される。変更不可。 <br> 請求未確認：NC  <br> 未請求：NP <br> 請求エラー：ER                |                                                                                                      | 
| HeaderGrossWeight             | ヘッダ総重量。オーダーから参照される。必要に応じて入力する。                                         |     | 
| HeaderNetWeight               | ヘッダ正味重量。オーダーから参照される。必要に応じて入力する。                                       |     | 
| HeaderVolume                  | ヘッダ数量。オーダーから参照される。必要に応じて入力する。                                           |     | 
| HeaderVolumeUnit              | ヘッダ数量単位。オーダーから参照される。必要に応じて入力する。                                       | ✔  | 
| HeaderWeightUnit              | ヘッダ重量単位。オーダーから参照される。必要に応じて入力する。                                       | ✔  | 
| Incoterms                     | インコタームズ分類。オーダーから参照される。必要に応じて、インコタームズマスタから選択して指定する。 | ✔  | 
| IsExportImportDelivery        | 輸出入フラグ。輸出入対象の入出荷である場合にtrueを設定する。                                         |     | 
| LastChangeDate                | 最終更新日付。自動生成される。変更不可。                                                             |     | 
| IssuingPlant                  | 出荷プラント。オーダーから参照される。必要に応じてプラントマスタから更新する。                       | ✔  | 
| ReceivingPlant                | 入荷プラント。オーダーから参照される。必要に応じてプラントマスタから更新する。                       | ✔  | 
| DeliverToParty                | 出荷先または入荷先のビジネスパートナコード。オーダーの取引先機能から参照される。変更不可。           |     | 
| DeliverFromParty              | 出荷元または入荷元のビジネスパートナコード。オーダーの取引先機能から参照される。変更不可。           |     | 
| DeliverToPartyLanguage        | 出荷先/入荷先言語。出荷先または入先のビジネスパートナマスタから設定される。必要に応じて変更する。    | ✔  | 
| DeliverFromPartyLanguage      | 出荷先/入荷先言語。出荷元または入荷元のビジネスパートナマスタから設定される。必要に応じて変更する。  | ✔  | 
| TransactionCurrency           | 取引通貨。オーダーから参照される。変更不可。                                                         |     | 
| OverallDelivReltdBillgStatus  | 全体入出荷関連請求ステータス。次のステータスが設定される。変更不可。 <br> 完全請求完了:CL <br> 部分請求完了:PP <br> 未請求:PP                     |                                                                                                      | 



| Property                      | Description                                                                                                            | 
| ----------------------------- | ---------------------------------------------------------------------------------------------------------------------- | 
| DeliveryDocument              | 入出荷伝票番号。自動登録される。                                                                                       | 
| OrderID                       | オーダー番号。入出荷伝票登録のために参照されるオーダー番号を入力する。                                                 | 
| OrderType                     | オーダータイプ。参照されたオーダーの契約タイプが設定される。変更不可。 <br> OD:オーダー <br> CT:契約 <br> PR:返品リクエスト <br> DR:デビットメモリクエスト <br> CR:クレジットメモリクエスト   | 
| ContractType                  | 契約タイプ。参照されたオーダーの契約タイプが設定される。変更不可。 <br> MNTH:月額契約。<br>ANNL:年額契約。               | 
| OrderVaridityStartDate        | オーダー契約開始日付。参照されたオーダーの契約開始日付が設定される。変更不可。                                         | 
| OrderValidityEndDate          | オーダー契約終了日付。参照されたオーダーの契約終了日付が設定される。変更不可。                                         | 
| InvoiceScheduleStartDate      | 請求計画開始日付。参照されたオーダーの請求計画開始日付が設定される。変更不可。                                         | 
| InvoiceScheduleEndDate        | 請求計画終了日付。参照されたオーダーの請求計画終了日付が設定される。変更不可。                                         | 
| IssuingLocationTimeZone       | 出荷プラントの時間帯。オーダー参照で設定された出荷プラントのタイムゾーンが設定される。変更不可。                       | 
| ReceivingLocationTimeZone     | 入荷プラントの時間帯。オーダー参照で設定された入荷プラントのタイムゾーンが設定される。変更不可。                       | 
| DocumentDate                  | 入出荷伝票日付。データ連携時のシステム日付が設定される。                                                               | 
| PlannedGoodsIssueDate         | 計画在庫出庫日付。実際に在庫を出庫した日付を入力する。                                                                 | 
| PlannedGoodsIssueTime         | 計画在庫出庫時刻。実際に在庫を出庫した時刻を入力する。                                                                 | 
| PlannedGoodsReceiptDate       | 計画在庫入庫日付。実際に在庫を入庫した日付を入力する。                                                                 | 
| PlannedGoodsReceiptTime       | 計画在庫入庫時刻。実際に在庫を入庫した時刻を入力する。                                                                 | 
| BillingDocumentDate           | 請求書日付。オーダーの請求書日付または支払条件をもとに参照・計算して請求日付が設定される。                             | 
| CompleteDeliveryIsDefined     | 完納フラグ。入出荷伝票に対して完全に入出荷が完了しているときにtrueが設定される。変更不可。                             | 
| OverallDeliveryStatus         | 全体入出荷ステータス。次のステータスが設定される。変更不可。 <br> NP:未入出荷 <br> PP:部分入出荷完了 <br> CL:入出荷完了                 | 
| CreationDate                  | 作成日付。自動生成される。変更不可。                                                                                   | 
| CreationTime                  | 作成時刻。自動生成される。変更不可。                                                                                   | 
| IssuingBlockReason            | 出荷ブロック理由。出荷をブロックする場合、trueを入力する。                                                             | 
| ReceivingBlockReason          | 入荷ブロック理由。入荷をブロックする場合、trueを入力する。                                                             | 
| GoodsIssueOrReceiptSlipNumber | 納品書または受領書番号。必要に応じて入力する。                                                                         | 
| HeaderBillgIncompletionStatus | ヘッダ請求不完全ステータス。システムにより設定される。変更不可。 <br> 請求未確認：NC <br> 未請求：NP <br>請求エラー：ER                | 
| HeaderBillingBlockReason      | ヘッダ請求ブロック理由。ヘッダ請求ブロック理由。請求をブロックする場合、trueを入力する。                               | 
| HeaderGrossWeight             | ヘッダ総重量。オーダーから参照される。必要に応じて入力する。                                                           | 
| HeaderNetWeight               | ヘッダ正味重量。オーダーから参照される。必要に応じて入力する。                                                         | 
| HeaderVolume                  | ヘッダ数量。オーダーから参照される。必要に応じて入力する。                                                             | 
| HeaderVolumeUnit              | ヘッダ数量単位。オーダーから参照される。必要に応じて入力する。                                                         | 
| HeaderWeightUnit              | ヘッダ重量単位。オーダーから参照される。必要に応じて入力する。                                                         | 
| Incoterms                     | インコタームズ分類。オーダーから参照される。必要に応じて、インコタームズマスタから選択して指定する。                   | 
| IsExportImportDelivery        | 輸出入フラグ。輸出入対象の入出荷である場合にtrueを設定する。                                                           | 
| LastChangeDate                | 最終更新日付。自動生成される。変更不可。                                                                               | 
| IssuingPlant                  | 出荷プラント。オーダーから参照される。必要に応じてプラントマスタから更新する。                                         | 
| ReceivingPlant                | 入荷プラント。オーダーから参照される。必要に応じてプラントマスタから更新する。                                         | 
| DeliverToParty                | 出荷先または入荷先のビジネスパートナコード。オーダーの取引先機能から参照される。変更不可。                             | 
| DeliverFromParty              | 出荷元または入荷元のビジネスパートナコード。オーダーの取引先機能から参照される。変更不可。                             | 
| DeliverToPartyLanguage        | 出荷先/入荷先言語。出荷先または入先のビジネスパートナマスタから設定される。必要に応じて変更する。                      | 
| DeliverFromPartyLanguage      | 出荷先/入荷先言語。出荷元または入荷元のビジネスパートナマスタから設定される。必要に応じて変更する。                    | 
| TransactionCurrency           | 取引通貨。オーダーから参照される。変更不可。                                                                           | 
| OverallDelivReltdBillgStatus  | 全体入出荷関連請求ステータス。次のステータスが設定される。変更不可。 <br> 完全請求完了:CL <br> 部分請求完了:PP <br> 未請求:PP                     | 
| BusinessPartner               | オーダーから参照されたビジネスパートナコード。項目としては存在するが、不使用。                                         | 
| OrderType                     | 参照されたオーダーのオーダータイプが設定される。項目としては存在するが、不使用。                                       | 
| DeliveryDocumentType          | 入出荷伝票タイプ。項目としては存在するが不使用。                                                                       | 
| DeliveryDate                  | 入出荷日付。項目としては存在するが不使用。                                                                             | 
| DeliveryTime                  | 入出荷時刻。項目としては存在するが不使用。                                                                             | 
| ActualDeliveryRoute           | 実際入出荷ルート。項目としては存在するが不使用。                                                                       | 
| ConfirmationTime              | 在庫確認時刻。項目としては存在するが不使用。                                                                           | 
| CustomerGroup                 | 得意先グループ。項目としては存在するが不使用。                                                                         | 
| DeliveryDocumentBySupplier    | 仕入先の入出荷伝票番号。項目としては存在するが、不使用。                                                               | 
| DeliveryDocumentType          | 入出荷伝票タイプ。項目としては存在するが不使用。                                                                       | 
| DeliveryIsInPlant             | プラント内入出荷。項目としては存在するが不使用。                                                                       | 
| DeliveryPriority              | 入出荷優先順位。項目としては存在するが不使用。                                                                         | 
| HeaderDelivIncompletionStatus | 明細入出荷の不完全ステータス。項目としては存在するが不使用。                                                           | 
| HeaderPackingIncompletionSts  | ヘッダパッキング不完全ステータス。項目としては存在するが不使用。                                                       | 
| HeaderPickgIncompletionStatus | ヘッダピッキング不完全ステータス。項目としては存在するが不使用。                                                       | 
| LoadingDate                   | 積載日付。項目としては存在するが不使用。                                                                               | 
| LoadingPoint                  | 積載ポイント。項目としては存在するが不使用。                                                                           | 
| LoadingTime                   | 積載時刻。項目としては存在するが不使用。                                                                               | 
| MeansOfTransport              | 輸送手段。項目としては存在するが不使用。                                                                               | 
| OrderCombinationIsAllowed     | 組合せ入出荷可能。項目としては存在するが不使用。（※複数受発注を組合せての入出荷登録には対応しない）                   | 
| PickedItemsLocation           | ピッキング場所。項目としては存在するが不使用。                                                                         | 
| PickingDate                   | ピッキング日付。項目としては存在するが不使用。                                                                         | 
| PickingTime                   | ピッキング時間。項目としては存在するが不使用。                                                                         | 
| PlannedGoodsIssueDate         | 計画入出庫日付。周辺業務システム内で登録された計画入出庫日付を入力する。計画入出庫日付。項目としては存在するが不使用。 | 
| ProposedDeliveryRoute         | 提案入出荷ルート。項目としては存在するが不使用。                                                                       | 
| RouteSchedule                 | ルートスケジュール。項目としては存在するが不使用。                                                                     | 
| SalesOffice                   | 営業所。項目としては存在するが不使用。                                                                                 | 
| SalesOrganization             | 販売組織。項目としては存在するが、不使用。                                                                             | 
| PurchaseOrganization          | 購買組織。項目としては存在するが、不使用。                                                                             | 
| SDDocumentCategory            | 販売伝票カテゴリ。項目としては存在するが不使用。                                                                       | 
| ShipmentBlockReason           | 入出荷ブロック理由。項目としては存在するが不使用。                                                                     | 
| ShippingCondition             | 出荷条件。項目としては存在するが不使用。                                                                               | 
| ShippingPoint                 | 出荷ポイント。項目としては存在するが不使用。                                                                           | 
| ShippingType                  | 出荷タイプ。項目としては存在するが不使用。                                                                             | 
| Customer                      | 受注先のビジネスパートナコード。項目としては存在するが、不使用。                                                       | 
| Supplier                      | 仕入先のビジネスパートナコード。項目としては存在するが、不使用。                                                       | 
| TotalBlockStatus              | トータルブロックステータス。項目としては存在するが不使用。                                                             | 
| TotalCreditCheckStatus        | トータル与信確認ステータス。項目としては存在するが不使用。                                                             | 
| TotalNumberOfPackage          | パッケージ総数。項目としては存在するが不使用。                                                                         | 
| TransportationGroup           | 輸送グループ。項目としては存在するが不使用。                                                                           | 
| TransportationPlanningDate    | 輸送計画日付。項目としては存在するが不使用。                                                                           | 
| TransportationPlanningStatus  | 輸送計画ステータス。項目としては存在するが不使用。                                                                     | 
| TransportationPlanningTime    | 輸送計画時間。項目としては存在するが不使用。                                                                           | 
| UnloadingPointName            | 積載ポイント名称。項目としては存在するが不使用。                                                                       | 

1-1. Language  

| Property | Description | EC |
| -------- | ----------- | --- |
| Language | 言語。ビジネスパートナマスタから設定される。必要に応じて変更する。または参照されて設定される。参照の場合は変更できない。 | ✔  |

1-1-1. 入力ファイルのBusinessPartnerをキーとして、対象のビジネスパートナのLanguageを検索して値をセットする。

<処理サマリ>  

| Target / Processing Type          | Key(s)                                      | 
| --------------------------------- | ------------------------------------------- | 
| Language / Get and Set            | BusinessPartner                             | 
| <b>Table Searched</b>                    | <b>Name of the Table</b>                           | 
| Business Partner SQL General Data | data_platform_business_partner_general_data | 
| <b>Field Searched</b>                    | <b>Data Type / Number of Digits</b>                | 
| Language                          | string(varchar) / 2                         | 
| <b>Single Record or Array</b>            | <b>Memo</b>                                        | 
| Single Record                     |                                             | 

【Address】
1. Delivery Document Address  

1-1. 住所マスタからの住所データの取得([PostalCode, LocalRegion, Country, District, StreetName, CityName, Building, Floor, Room])  

| Property    | Description                                                                                                        | EC  | 
| ----------- | ------------------------------------------------------------------------------------------------------------------ | --- | 
| AddressID   | 住所ID。オーダーから設定される。入出荷伝票上でマニュアルで住所を更新する場合、住所IDが新たに設定される。変更不可。 |     | 
| PostalCode  | 郵便番号。オーダーから設定される。必要に応じて変更する。                                                           | ✔  | 
| LocalRegion | ローカル地域。オーダーから設定される。必要に応じて変更する。                                                       | ✔  | 
| Country     | 国。オーダーから設定される。必要に応じて変更する。                                                                 | ✔  | 
| District    | ディストリクト。オーダーから設定される。必要に応じて変更する。                                                     | ✔  | 
| StreetName  | 地名・番地。オーダーから設定される。必要に応じて変更する。                                                         |     | 
| CityName    | 市区町村。オーダーから設定される。必要に応じて変更する。                                                           |     | 
| Building    | 建物名。オーダーから設定される。必要に応じて変更する。                                                             |     | 
| Floor       | 階数。オーダーから設定される。必要に応じて変更する。                                                               |     | 
| Room        | 部屋番号。オーダーから設定される。必要に応じて変更する。                                                           |     | 






1-1-1. 2-2-2で取得した[AddressID]をコピーしてセットする。

1-1-2. [1-1-1でセットした AddressID]と、[ValidityEndDate≧システム日付]をキーとして、対象の住所データの[PostalCode, LocalRegion, Country, District, StreetName, CityName, Building, Floor, Room]を検索しセットする。  

| Target / Processing Type        | Key(s)                             | 
| ------------------------------- | ---------------------------------- | 
| 下記Field Searchedと同一        | 
| / Get and Set                   | [AddressID] <b> [ValidityEndDate≧システム日付] | 
| <b>Table Searched</b>                  | <b>Name of the Table</b>                  | 
| Address SQL Address Data        | data_platform_address_address_data | 
| <b>Field Searched</b>                  | <b>Data Type / Number of Digits</b>       | 
| PostalCode <br> LocalRegion<br> Country <br> District <br> StreetName <br> CityName <br> Building <br> Floor  <br> Room | string(varchar) / 10 <br> string(varchar) / 3 <br> string(varchar) / 3 <br> string(varchar) / 6 <br> string(varchar) / 200 <br> string(varchar) / 200 <br>string(varchar) / 100 <br> int / 4 <br> int / 8                         | 
| <b>Single Record or Array</b>          | <b>Memo</b>                               | 
| Array                           |                                    | 

1-2. AddressIDの登録(ユーザーが任意の住所を入力ファイルで指定した場合)  

| Property | Description | EC |
| -------- | ----------- | ---- |
|AddressID | 住所ID。ビジネスパートナマスタまたは見積/引合から設定される。オーダーでマニュアルで住所を更新する場合、住所IDが新たに設定される。変更不可。 |    |	

1-2-1. 入力ファイルの[ValidityEndDate, ValidityStartDate, PostalCode, LocalRegion, Country, GlobalRegion, TimeZone, StreetName, CityName]が、ブランクでなく、かつnullでない場合、住所IDの登録を行う。  

1-2-2. AddressIDについては、ServiceLabel=”ADDRESS_ID”, FieldNameWithNumberRange=当処理のProperty(=”AddressID”)をキーとして、対象のNumber Range Latest NumberのLatestNumberを検索し保持する。  

| Target / Processing Type                                  | Key(s)                                        | 
| --------------------------------------------------------- | --------------------------------------------- | 
| ServiceLabel <br> FieldNameWithNumberRange <br> LatestNumber <br> / Get and Hold | ServiceLabel=”ADDRESS_ID” <br> FieldNameWithNumberRange =”AddressID”(当処理のProperty) | 
| <b>Table Searched</b>                                            | <b>Name of the Table</b>                             | 
| Number Range SQL Latest Number Data                       | data_platform_number_range_latest_number_data | 
| <b>Field Searched</b>                                            | <b>Data Type / Number of Digits</b>                  | 
| ServiceLabel <br> FieldNameWithNumberRange <br> LatestNumber  | string(varchar) / 50 <br> string(varchar) / 100 <br> int / 10                                                  | 
| <b>Single Record or Array</b>                                    | <b>Memo</b>                                          | 
| Single Record                                             |                                               | 

1-2-3. 保持されたLatestNumberに1を足したものをAddressIDにセットする。  

1-2-4. アドレスマスタレコードの生成  
次の処理の通り、アドレスマスタレコードを生成し、アドレスマスタに登録する。  

| 生成する項目      | 生成する値等のロジック            | 
| ----------------- | --------------------------------- | 
| AddressID         | 8-2-3でセットしたAddressID。      | 
| ValidityEndDate   | 入力ファイルのValidityEndDate。   | 
| ValidityStartDate | 入力ファイルのValidityStartDate。 | 
| PostalCode        | 入力ファイルのPostalCode。        | 
| LocalRegion       | 入力ファイルのLocalRegion。       | 
| Country           | 入力ファイルのCountry。           | 
| GlobalRegion      | 入力ファイルのGlobalRegion。      | 
| TimeZone          | 入力ファイルのTimeZone。          | 
| District          | 入力ファイルのDistrict。          | 
| StreetName        | 入力ファイルのStreetName。        | 
| CityName          | 入力ファイルのCityName。          | 
| Builiding         | 入力ファイルのBuiliding。         | 
| Floor             | 入力ファイルのFloor。             | 
| Room              | 入力ファイルのRoom。              | 

| Target / Processing Type                                                               | Key(s)                                | 
| -------------------------------------------------------------------------------------- | ------------------------------------- | 
| アドレスマスタの各項目(上記表の通り) / Creates                                         | AddressID(7-2-3でセットしたAddressID) | 
| <b>Table Searched</b>                                                                         | <b>Name of the Table</b>                     | 
| None                                                                                   | None                                  | 
| <b>Field Searched</b>                                                                         | <b>Data Type / Number of Digits</b>          | 
| アドレスマスタの各項目(上記表の通り)                                                   | 上記表の通り                          | 
| <b>Single Record or Array</b>                                                                 | <b>Memo</b>                                  | 
| Single Record                                                                          | data_platform_api_address_creates <br> ※本処理を行ってからOrders Createsの最終処理であるDBへのレコード登録を行う必要がある。 | 

### Accepter名:【Item】  

1-0. StockConfirmationPartnerFunction / StockConfirmationBusinessPartner /StockConfirmationPlant  

| Property                         | Description                                                                                                                                          | 
| -------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | 
| StockConfirmationPartnerFunction | 在庫確認を行う取引先機能。オーダー明細参照で設定される。必要に応じて変更する。（PRODUCT_STOCK_AVAIRABILITY_CHECK_SRVを利用、機能はオプション）       | 
| StockConfirmationBusinessPartner | 在庫確認を行うビジネスパートナ。オーダー明細参照で設定される。必要に応じて変更する。（PRODUCT_STOCK_AVAIRABILITY_CHECK_SRVを利用、機能はオプション） | 
| StockConfirmationPlant           | 在庫確認プラント。オーダー明細参照で設定される。必要に応じて変更する。（PRODUCT_STOCK_AVAIRABILITY_CHECK_SRVを利用、機能はオプション）               | 
| StockConfirmationStatus          | 在庫確認ステータス。次の在庫確認ステータスが常時更新される。変更不可。<br> CL:全部完了済 <br> PP:部分完了済 <br> NP:未確認                        | 

1-1. 在庫確認  
Delivery_Document_Creates_Subfuncに含まれる入力ファイルの[Delivery Document ItemCategory=”INVP”の明細]に対して、本処理を実行する。  

1-1-1.在庫確認是非の判断  

| Property                     | Description                                                    | EC  | 
| ---------------------------- | -------------------------------------------------------------- | --- | 
| DeliveryDocumentItemCategory | 入出荷伝票明細カテゴリ。オーダー明細から参照される。変更不可。<br> INVP:在庫明細  <br> SRVP:サービス/非在庫明細 <br> SDTP:仕入先直送明細 <br> BTOP:受注生産明細 <br> PJPP:プロジェクト生産明細    |                                                                | 

1-1-1-1. [DeliveryDocumentItemCategory =”INVP”]の時、本在庫確認の処理を実行する。  

1-1-2. 在庫確認①(ロット単位での在庫確認)  

1-1-2-1. [data_platform_function_orders_creates_subfunc 内の5-5-1の品目在庫確認プラントにおけるIsBatchManagementRequired]が[true]だった場合、[対象のDeliveryDocument, DeliveryDocumentItem, Product]についてロット単位での在庫確認を行う。  

1-1-2-2. [DeliveryDocument, DeliveryDocumentItem, IssuingPlantPartnerFunction, IssuingPlantBusinessPartner, IssuingPlant, DeliveryQuantityUnit]に基づき、[入力ファイルのProduct, IssuingPlantBusinessPartner, IssuingPlant, 入力ファイルのIssuingPlantBatch, 入力ファイルのIssuingPlantBatchValidityStartDate, 入力ファイルのIssuingPlantBatchValidityEndDate, 入力ファイルのPlannedGoodsIssueDate]をinputとして、data_platform_function_product_availability_check(利用可能在庫確認)を実行し、品目利用可能在庫データの[Product, BusinessPartner, Plant, Batch, BatchValidityEndDate, ProductStockAvailabilityDate, AvailableProductStock]、を取得して保持する。プラントやロット、日付等のキー違いなどでエラーが発生した場合、エラーメッセージを出力して終了する。  

| Target / Processing Type                          | Key(s)                                            | 
| ------------------------------------------------- | ------------------------------------------------- | 
| Product<br> BusinessPartner <br> Plant <br> Batch <br> BatchValidityEndDate <br> ProductStockAvailabilityDate <br> AvailableProductStock  <br> / Get and Hold | Product(入力ファイル) <br> IssuingPlantBusinessPartner <br> IssuingPlant  <br> PlannedGoodsIssueDate (入力ファイル) <br> IssuingPlantBatch (入力ファイル)  <br> IssuingPlantBatchValidityStartDate (入力ファイル) <br> IssuingPlantBatchValidityEndDate (入力ファイル)   | 
| <b>Table Searched</b>                                    | <b>Name of the Function</b>                              | 
| None                                              | data_platform_function_product_availability_check | 
| <b>Field Searched</b>                                    | <b>Data Type / Number of Digits</b>                      | 
Product<br> BusinessPartner<br> Plant<br> Batch<br> BatchValidityEndDate<br> ProductStockAvailabilityDate<br> AvailableProductStock | string(varchar) / 40 <br> int / 12 <br> string(varchar) / 4 <br> string(varchar) / 10 <br> date <br> date <br> float / 15                                        | 
| <b>Single Record or Array</b>                            | <b>Memo</b>                                              | 
| Array                                             |                                                   |

1-1-2-3. 入出荷伝票在庫確認データの生成(ロット在庫)  
[1-1-2-2で在庫確認が行われたDeliveryDocument, DeliveryDocumentItemの明細]について、同じ明細数だけ、次の処理の通り、[入出荷伝票在庫確認データ]を生成する。
※在庫確認の結果の在庫確認数量がゼロである場合も入出荷伝票在庫確認データが生成される。  

| 生成する項目                                                                                     | 生成する値等のロジック                                                                                         | 
| ------------------------------------------------------------------------------------------------ | -------------------------------------------------------------------------------------------------------------- | 
| DeliveryDocument (※Key)                                                                         | 本処理(1-1-2)にて保持済のDeliveryDocument。                                                                    | 
| DeliveryDocumentItem (※Key)                                                                     | 本処理(1-1-2)にて保持済のDeliveryDocumentItem。                                                                | 
| Product                                                                                          | 本処理(1-1-2)にて保持済みのProduct。                                                                           | 
| IssuingPlantPartnerFunction                                                                      | 本処理(1-1-2)にて保持済みのIssuingPlantPartnerFunction。                                                       | 
| IssuingPlantBusinessPartner                                                                      | 本処理(1-1-2)にて保持済みのIssuingPlantBusinessPartner。                                                       | 
| IssuingPlant                                                                                     | 本処理(1-1-2)にて保持済みのIssuingPlant。                                                                      | 
| IssuingPlantBatch                                                                                | 本処理(1-1-2)にて保持済みのIssuingPlantBatch。                                                                 | 
| IssuingPlantBatchValidityStartDate                                                               | 本処理(1-1-2)にて保持済みのIssuingPlantBatchValidityStartDate                                                  | 
| IssuingPlantBatchValidityEndDate                                                                 | 本処理(1-1-2)にて保持済みのIssuingPlantBatchValidityEndDate                                                    | 
| PlannedGoodsIssueDate                                                                            | 本処理(1-1-2)にて保持済みのPlannedGoodsIssueDate。                                                             | 
| ConfirmedDeliveryDate                                                                            | 1-1-2-2で取得したProductStockAvailabilityDate <br> ※確認された在庫が無い場合、API処理の結果により、PlannedGoodsIssueDateと同じ日付がセットされる。 | 
| DeliveryQuantityUnit                                                                             | 本処理(1-1-2)にて保持済みのDeliveryQuantityUnit                                                                | 
| ConfdOrderQtyByPDTAvailCheck                                                                     | 1-1-2-2で取得したAvailableProductStock <br> ※確認された在庫が無い場合、API処理の結果により、ゼロがセットされる。                            | 
| DeliveredQtyInOrderQtyUnit                                                                       | nullをセット。                                                                                                 | 
| OpenConfdDelivQtyInOrdQtyUnit                                                                    | OrderQuantityInBaseUnitからConfdOrderQtyByPDTAvailCheckを減算した値を計算してセット。                          | 
| StockIsFullyConfirmed                                                                            | OrderQuantityInBaseUnit= ConfdOrderQtyByPDTAvailCheckである場合、trueをセット。それ以外の場合、falseをセット。 | 
| DelivBlockReasonForSchedLine                                                                     | Falseをセット。                                                                                                | 
| PlusMinusFlag                                                                                    | “-“(マイナス)をセット。                                                                                      | 

1-1-2-4. 入出荷伝票明細への在庫確認済フラグ/在庫確認ステータスのセット  

| Property | Description | EC |
| -------- | ----------- | --- |
| StockConfirmationStatus | 在庫確認ステータス。次の在庫確認ステータスが常時更新される。変更不可。 <br> CL:全部完了済 <br> PP:部分完了済 <br> NP:確認済数量無し	|     |

1-1-2-3において、[StockIsFullyConfirmed]がtrueの場合、[StockConfirmationStatus]に”CL”をセットする。[StockIsFullyConfirmed]がfalseでかつ[ConfdOrderQtyByPDTAvailCheck]がゼロである場合、”NP”をセットする。それ以外の場合、”PP”をセットする。  

1-1-3. 在庫確認②(通常の在庫確認)  
1-1-3-1. data_platform_function_orders_creates_subfunc 5-5-1の品目在庫確認プラントにおけるIsBatchManagementRequiredがfalseだった場合、次の処理を行う。  

1-1-3-2. [DeliveryDocument, DeliveryDocumentItem, IssuingPlantPartnerFunction, IssuingPlantBusinessPartner, IssuingPlant, DeliveryQuantityUnit]に基づき、[入力ファイルのProduct, 1-0で取得したIssuingPlantBusinessPartner, IssuingPlant, 入力ファイルのPlannedGoodsIssueDate]をinputとして、data_platform_function_product_availability_check(利用可能在庫確認)を実行し、品目利用可能在庫データの[Product, BusinessPartner, Plant, ProductStockAvailabilityDate, AvailableProductStock]、を取得して保持する。プラント、日付等のキー違いなどでエラーが発生した場合、エラーメッセージを出力して終了する。  

| Target / Processing Type                | Key(s)                                            | 
| --------------------------------------- | ------------------------------------------------- | 
| Product <br> BusinessPartner <br> Plant <br> ProductStockAvailabilityDate <br> AvailableProductStock <br> / Get and Hold  | Product(入力ファイル)  <br> IssuingPlantBusinessPartner (1-0で取得) <br> IssuingPlant (1-0で取得) <br> PlannedGoodsIssueDate (入力ファイル)    | 
| <b>Table Searched</b>                          | <b>Name of the Function</b>                              | 
| None                                    | data_platform_function_product_availability_check | 
| <b>Field Searched</b>                          | <b>Data Type / Number of Digits</b>                      | 
| Product <br> BusinessPartner<br> Plant<br> ProductStockAvailabilityDate<br> AvailableProductStock | string(varchar) / 40 <br> int / 12 <br>string(varchar) / 4 <br> date <br> float / 15                              | 
| <b>Single Record or Array</b>                  | <b>Memo</b>                                              | 
| Array                                   | 	                                               |

1-1-3-3. 入出荷伝票在庫確認データの生成(通常の在庫確認)  
[1-1-3-2で在庫確認が行われたDeliveryDocument, DeliveryDocumentItem,の明細]について、同じ明細数だけ、次の処理の通り、[入出荷伝票在庫確認データ]を生成する。
※在庫確認の結果の在庫確認数量がゼロである場合も入出荷伝票在庫確認データが生成される。  

| 生成する項目                                                                                     | 生成する値等のロジック                                                                                        | 
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------- | 
| DeliveryDocument(※Key)                                                                          | 本処理(1-1-3)にて保持済のDeliveryDocument。                                                                   | 
| DeliveryDocumentItem(※Key)                                                                      | 本処理(1-1-3)にて保持済のDeliveryDocumentItem。                                                               | 
| Product                                                                                          | 本処理(1-1-3)にて保持済みのProduct。                                                                          | 
| IssuingPlantPartnerFunction                                                                      | 本処理(1-1-3)にて保持済みのIssuingPlantPartnerFunction。                                                      | 
| IssuingPlantBusinessPartner                                                                      | 本処理(1-1-3)にて保持済みのIssuingPlantBusinessPartner。                                                      | 
| IssuingPlant                                                                                     | 本処理(1-1-3)にて保持済みのIssuingPlant。                                                                     | 
| IssuingPlantBatch                                                                                | “”(ブランク)をセット。                                                                                      | 
| IssuingPlantBatchValidityStartDate                                                               | nullをセット。                                                                                                | 
| IssuingPlantBatchValidityEndDate                                                                 | nullをセット。                                                                                                | 
| PlannedGoodsIssueDate                                                                            | 本処理(1-1-3)にて保持済みのPlannedGoodsIssueDate。                                                            | 
| ConfirmedDeliveryDate                                                                            | 1-1-3-2で取得したProductStockAvailabilityDate。 <br> ※確認された在庫が無い場合、API処理の結果により、PlannedGoodsIssueDateと同じ日付がセットされる。 | 
| DeliveryQuantityUnit                                                                             | 本処理(1-1-3)にて保持済みのDeliveryQuantityUnit。                                                             | 
| ConfdOrderQtyByPDTAvailCheck                                                                     | 1-1-3-2で取得したAvailableProductStock <br> ※確認された在庫が無い場合、API処理の結果により、ゼロがセットされる。                            | 
| DeliveredQtyInOrderQtyUnit                                                                       | nullをセット。                                                                                                | 
| OpenConfdDelivQtyInOrdQtyUnit                                                                    | PlannedGoodsIssueDateからConfdOrderQtyByPDTAvailCheckを減算した値を計算してセット。                           | 
| StockIsFullyConfirmed                                                                            | PlannedGoodsIssueDate＝ ConfdOrderQtyByPDTAvailCheckである場合、trueをセット。それ以外の場合、falseをセット。 | 
| DelivBlockReasonForSchedLine                                                                     | Falseをセット。                                                                                               | 
| PlusMinusFlag                                                                                    | “-“(マイナス)をセット。                                                                                     | 

1-1-3-4. 入出荷伝票明細への在庫確認済フラグ/在庫確認ステータスのセット


| Property | Description | EC |
| -------- | ----------- | --- |
| StockConfirmationStatus | 在庫確認ステータス。次の在庫確認ステータスが常時更新される。変更不可。<br> CL:全部完了済 <br> PP:部分完了済 <br> NP:確認済数量無し	|     |

1-1-3-3において、[StockIsFullyConfirmed]がtrueの場合、[StockConfirmationStatus]に”CL”をセットする。[StockIsFullyConfirmed]がfalseでかつ[ConfdOrderQtyByPDTAvailCheck]がゼロである場合、”NP”をセットする。それ以外の場合、”PP”をセットする。

