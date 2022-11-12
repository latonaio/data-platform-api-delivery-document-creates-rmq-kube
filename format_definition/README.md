# format_definition
format_definition はデータ連携基盤における入出荷伝票の仕様について、API仕様書のフォーマットに従って簡潔にまとめたものです。

## APIサービス名:【Delivery Document > DPFM_API_DELIVERY_DOCUMENT_SRV】

### API名:【Delivery Document Header > A_Header】
### Accepter名:【Header】

### <項目>
以下の表には、data-platform-api-delivery-document-creates-rmq-kube のデータの項目とその定義が記載されています。  
※ EC の列には、そのデータの存在性確認を行うと定めているかどうか（✔がついていれば、行うと定めている）が記載されています。

| Property                      | Description                                                                                                                      | EC  | 
| ----------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | --- | 
| DeliveryDocument              | 入出荷伝票番号。自動登録される。                                                                                                 |     | 
| Buyer                         | 買い手。オーダーから参照される。変更不可。                                                                                       |     | 
| Seller                        | 売り手。オーダーから参照される。変更不可。                                                                                       |     | 
| ContractType                  | 契約タイプ。オーダーから参照される。変更不可。 <br> MNTH:月額契約。<br>ANNL:年額契約。               |     |
| OrderVaridityStartDate        | オーダー契約開始日付。オーダーから参照される。変更不可。                                                                         |     | 
| OrderValidityEndDate          | オーダー契約終了日付。オーダーから参照される。変更不可。                                                                         |     | 
| InvoiceScheduleStartDate      | 請求計画開始日付。オーダーから参照される。変更不可。                                                                             |     | 
| InvoiceScheduleEndDate        | 請求計画終了日付。オーダーから参照される。変更不可。                                                                             |     | 
| IssuingPlantTimeZone          | 出荷プラントの時間帯。オーダー明細から参照される。変更不可。                                                                     |     | 
| ReceivingPlantTimeZone        | 入荷プラントの時間帯。オーダー明細から参照される。変更不可。                                                                     |     | 
| DocumentDate                  | 入出荷伝票日付。入出荷伝票を入出荷予定として登録した日付。                                                                       |     | 
| PlannedGoodsIssueDate         | 計画在庫出庫日付。                                                                                                               |     | 
| PlannedGoodsIssueTime         | 計画在庫出庫時刻。                                                                                                               |     | 
| PlannedGoodsReceiptDate       | 計画在庫入庫日付。                                                                                                               |     | 
| PlannedGoodsReceiptTime       | 計画在庫入庫時刻。                                                                                                               |     | 
| BillingDocumentDate           | 請求書日付。計画在庫出庫日付および支払条件をもとに参照・計算して請求日付が設定される。                                           |     | 
| CompleteDeliveryIsDefined     | 完納フラグ。入出荷伝票に対して完全に入出荷が完了しているときにtrueが設定される。変更不可。                                       |     | 
| OverallDeliveryStatus         | 全体入出荷ステータス。次のステータスが設定される。変更不可。 <br> NP:未入出荷 <br> PP:部分入出荷完了<br> CL:入出荷完了                                       |       |
| CreationDate                  | 作成日付。自動生成される。変更不可。                                                                                             |     | 
| CreationTime                  | 作成時刻。自動生成される。変更不可。                                                                                             |     | 
| IssuingBlockReason            | 出荷ブロック理由。出荷をブロックする場合、trueを入力する。                                                                       |     | 
| ReceivingBlockReason          | 入荷ブロック理由。入荷をブロックする場合、trueを入力する。                                                                       |     | 
| GoodsIssueOrReceiptSlipNumber | 納品書または受領書番号。必要に応じて入力する。                                                                                   |     | 
| HeaderBillingStatus           | ヘッダ請求ステータス。<br> 完全請求完了:CL <br> 部分請求完了:PP <br> 未請求:NP <br> が設定される。変更不可。      |      |                                                                                                                            | 
| HeaderBillingConfStatus       | ヘッダ請求照合ステータス。<br>完全請求照合完了:CL <br> 部分請求照合完了:PP <br> 未請求照合:NP  <br>が設定される。変更不可。      |      |                                                                                                                            | 
| HeaderBillingBlockReason      | ヘッダ請求ブロック理由。ヘッダ請求ブロック理由。請求をブロックする場合、trueを入力する。                                         |     | 
| HeaderGrossWeight             | ヘッダ総重量。入出荷伝票明細の総重量の合計。重量単位は、HeaderWeightUnitの仕様の通りとして必要な再計算が為される。変更不可。     |     | 
| HeaderNetWeight               | ヘッダ正味重量。入出荷伝票明細の正味重量の合計。重量単位は、HeaderWeightUnitの仕様の通りとして必要な再計算が為される。変更不可。 |     | 
| HeaderWeightUnit              | ヘッダ重量単位。明細の重量単位から最も軽い単位が設定される。変更不可。                                                           |     | 
| Incoterms                     | インコタームズ分類。オーダーから参照される。必要に応じて、インコタームズマスタから選択して指定する。                             | ✔  | 
| IsExportImportDelivery        | 輸出入フラグ。輸出入対象の入出荷である場合にtrueが設定される。オーダーから参照される。変更不可。                                 |     | 
| LastChangeDate                | 最終更新日付。自動生成される。変更不可。                                                                                         |     | 
| IssuingPlantBusinessPartner   | 出荷プラントビジネスパートナ。オーダー明細から参照される。必要に応じてプラントマスタから更新する。                               | ✔  | 
| IssuingPlant                  | 出荷プラント。オーダー明細から参照される。必要に応じてプラントマスタから更新する。                                               | ✔  | 
| ReceivingPlantBusinessPartner | 入荷プラントビジネスパートナ。オーダー明細から参照される。必要に応じてプラントマスタから更新する。                               | ✔  | 
| ReceivingPlant                | 入荷プラント。オーダー明細から参照される。必要に応じてプラントマスタから更新する。                                               | ✔  | 
| DeliverToParty                | 出荷先または入荷先のビジネスパートナコード。オーダーの取引先機能から参照される。必要に応じて変更する。                           | ✔  | 
| DeliverFromParty              | 出荷元または入荷元のビジネスパートナコード。オーダーの取引先機能から参照される。必要に応じて変更する。                           | ✔  | 
| TransactionCurrency           | 取引通貨。オーダーから参照される。変更不可。                                                                                     |     | 
| OverallDelivReltdBillgStatus  | 全体入出荷関連請求ステータス。次のステータスが設定される。変更不可。<br> 完全請求完了:CL <br> 部分請求完了:PP <br> 未請求:PP                     |    |                                                                                                               | 

### API名:【Delivery Document Header Partner> A_HeaderPartner】
### Accepter名:【HeaderPartner】

### <項目>
以下の表には、data-platform-api-delivery-document-creates-rmq-kube のデータの項目とその定義が記載されています。  
※ EC の列には、そのデータの存在性確認を行うと定めているかどうか（✔がついていれば、行うと定めている）が記載されています。

| Property                                    | Description                                                                                                                                                  | EC  | 
| ------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | --- | 
| PartnerFunction                             | 取引先機能。オーダーから参照される。変更不可。なお、入出荷伝票ヘッダにおいて、それぞれの取引先機能レコードは入出荷伝票に対して下記の()内の数だけ設定される。 <br> 参照される取引先機能: <br> BUYER:買い手(一つ) <br> SELLER:売り手(一つ) <br> CUSTOMER:受注先(一つ) <br> SUPPLIER:仕入先(一つ) <br> MANUFACTURER:製造者(複数) <br> DELIVERFROM:入出荷元(複数) <br> DELIVERTO:入出荷先(複数) <br> LOGI:物流業者(複数) <br> BILLTO:請求先(一つ) <br> BILLFROM:請求元(一つ) <br> PAYEE:支払人(一つ) <br> RECEIVER:受取人(一つ)  <br> PSPROVIDER:支払決済サービスプロバイダ(一つ) | ✔   |                                                                                                                                                      | 
| BusinessPartner                             | 取引先機能に対応するビジネスパートナコード。オーダーから設定される。必要に応じて変更する。                                                                   |     | 
| BusinessPartnerFullName                     | ビジネスパートナフルネーム。オーダーから設定される。必要に応じて変更する。                                                                                   |     | 
| BusinessPartnerName                         | ビジネスパートナ名称。オーダーから設定される。必要に応じて変更する。                                                                                         |     | 
| Organization                                | 組織コード。オーダーから設定される。必要に応じて変更する。                                                                                                   |     | 
| Language                                    | 言語コード。オーダーから設定される。必要に応じて変更する。                                                                                                   |     | 
| Currency                                    | 通貨コード。オーダーから設定される。変更不可。                                                                                                               |     | 
| ExternalDocumentID                          | 外部文書ID。オーダーから設定される。必要に応じて変更する。                                                                                                   |     | 
| AddressID                                   | 住所ID。オーダーの住所IDが設定される。マニュアルで住所を更新する場合、住所IDが新たに設定される。変更不可。                                                   |     | 

### API名:【Delivery Document Header Partner Contact > A_HeaderPartnerContact】
### Accepter名:【HeaderPartnerContact】

### <項目>
以下の表には、data-platform-api-delivery-document-creates-rmq-kube のデータの項目とその定義が記載されています。  
※ EC の列には、そのデータの存在性確認を行うと定めているかどうか（✔がついていれば、行うと定めている）が記載されています。

| Property          | Description                                                                                                                                                                | EC  | 
| ----------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --- | 
| PartnerFunction   | 取引先機能。ヘッダパートナから設定される。変更不可。                                                                                                                       |     | 
| ContactID         | コンタクトID。ビジネスパートナ得意先取引先機能コンタクトデータまたは仕入先取引先機能コンタクトデータから設定される。必要に応じて変更する。参照の場合は変更できない。       |     | 
| BusinessPartner   | ビジネスパートナコード。ヘッダパートナから設定される。変更不可。                                                                                                           |     | 
| ContactPersonName | コンタクト担当者名。ビジネスパートナ得意先取引先機能コンタクトデータまたは仕入先取引先機能コンタクトデータから設定される。必要に応じて変更する。参照の場合は変更できない。 |     | 
| EmailAddress      | Eメールアドレス。ビジネスパートナ得意先取引先機能コンタクトデータまたは仕入先取引先機能コンタクトデータから設定される。必要に応じて変更する。参照の場合は変更できない。    |     | 
| PhoneNumber       | 電話番号。ビジネスパートナ得意先取引先機能コンタクトデータまたは仕入先取引先機能コンタクトデータから設定される。必要に応じて変更する。参照の場合は変更できない。           |     | 
| MobilePhoneNumber | モバイル電話番号。ビジネスパートナ得意先取引先機能コンタクトデータまたは仕入先取引先機能コンタクトデータから設定される。必要に応じて変更する。参照の場合は変更できない。   |     | 
| FaxNumber         | ファクス番号。ビジネスパートナ得意先取引先機能コンタクトデータまたは仕入先取引先機能コンタクトデータから設定される。必要に応じて変更する。参照の場合は変更できない。       |     | 
| ContactTag1       | コンタクトタグ1。ビジネスパートナ得意先取引先機能コンタクトデータまたは仕入先取引先機能コンタクトデータから設定される。必要に応じて変更する。参照の場合は変更できない。    |     | 
| ContactTag2       | コンタクトタグ2。ビジネスパートナ得意先取引先機能コンタクトデータまたは仕入先取引先機能コンタクトデータから設定される。必要に応じて変更する。参照の場合は変更できない。    |     | 
| ContactTag3       | コンタクトタグ3。ビジネスパートナ得意先取引先機能コンタクトデータまたは仕入先取引先機能コンタクトデータから設定される。必要に応じて変更する。参照の場合は変更できない。    |     | 
| ContactTag4       | コンタクトタグ4。ビジネスパートナ得意先取引先機能コンタクトデータまたは仕入先取引先機能コンタクトデータから設定される。必要に応じて変更する。参照の場合は変更できない。    |     | 

### API名:【Delivery Document Header PDF > A_HeaderPDF】
### Accepter名:【HeaderPDF】

### <項目>
以下の表には、data-platform-api-delivery-document-creates-rmq-kube のデータの項目とその定義が記載されています。  
※ EC の列には、そのデータの存在性確認を行うと定めているかどうか（✔がついていれば、行うと定めている）が記載されています。

| Property                 | Description                                                                    | EC  | 
| ------------------------ | ------------------------------------------------------------------------------ | --- | 
| DocType                  | 文書タイプ。入出荷伝票の場合、次から選択する。 <br> DELIVERYSLIP:納品書 <br> SHIPPINGINST:出荷指示書  |      |                                                                          | 
| DocVersionID             | 文書のバージョンID。                                                           |     | 
| DocID                    | 文書のID。ハッシュ値で登録される。変更不可。                                   |     | 
| DocIssuerBusinessPartner | 文書発行者ビジネスパートナコード。ビジネスパートナマスタから選択して入力する。 | ✔  | 
| FileName                 | ファイル名称。                                                                 |     | 

### API名:【Delivery Document Address> A_Address】
### Accepter名:【Address】

### <項目>
以下の表には、data-platform-api-delivery-document-creates-rmq-kube のデータの項目とその定義が記載されています。  
※ EC の列には、そのデータの存在性確認を行うと定めているかどうか（✔がついていれば、行うと定めている）が記載されています。

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

### API名:【Delivery Document Header Partner Plant > A_HeaderPartnerPlant】
### Accepter名:【HeaderPartnerPlant】

### <項目>
以下の表には、data-platform-api-delivery-document-creates-rmq-kube のデータの項目とその定義が記載されています。  
※ EC の列には、そのデータの存在性確認を行うと定めているかどうか（✔がついていれば、行うと定めている）が記載されています。

| Property                   | Description                                                                                                                                                  | EC  | 
| -------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | --- | 
| PartnerFunction            | 取引先機能。オーダーから参照される。変更不可。なお、入出荷伝票ヘッダにおいて、それぞれの取引先機能レコードは入出荷伝票に対して下記の()内の数だけ設定される。 <br> 参照される取引先機能: <br> MANUFACTURER:製造者(複数) <br> DELIVERFROM:入出荷元(一つ) <br> DELIVERTO:入出荷先(一つ)   |     |                                                                                                                                                         | 
| BusinessPartner            | 取引先機能に対応するビジネスパートナコード。オーダーから設定される。必要に応じて変更する。                                                                   |     | 
| Plant                      | プラント。取引先機能に対応するプラント。オーダーから設定される。必要に応じて変更する。                                                                       | ✔  | 

### API名:【Delivery Document Item > A_Item】
### Accepter名:【Item】

### <項目>
以下の表には、data-platform-api-delivery-document-creates-rmq-kube のデータの項目とその定義が記載されています。  
※ EC の列には、そのデータの存在性確認を行うと定めているかどうか（✔がついていれば、行うと定めている）が記載されています。

| Property                                 | Description                                                                                                                                                                                                                                              | EC  | 
| ---------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --- | 
| DeliveryDocumentItem                     | 入出荷伝票明細番号。オーダー明細参照で自動登録される。変更不可。                                                                                                                                                                                         |     | 
| DeliveryDocumentItemCategory             | 入出荷伝票明細カテゴリ。オーダー明細から参照される。変更不可。 <br> TAN:標準明細 <br> TAS:仕入先直送明細 <br> TAD:サービス明細                         |       |                                                                                                                                                                                                                                                   | 
| DeliveryDocumentItemText                 | 入出荷伝票明細テキスト。オーダー明細参照で設定される。変更不可。                                                                                                                                                                                         |     | 
| Product                                  | 品目コード。オーダー明細から参照される。変更不可。                                                                                                                                                                                                       |     | 
| ProductGroup                             | 品目グループ。オーダー明細から参照される。変更不可。                                                                                                                                                                                                     |     | 
| ProductStandardID                        | 国際商品識別コード。受注伝票明細から参照される。変更不可。                                                                                                                                                                                               |     | 
| BaseUnit                                 | 基本単位。オーダー明細参照で設定される。変更不可。                                                                                                                                                                                                       |     | 
| OriginalDeliveryQuantity                 | 元の入出荷数量。オーダー明細から参照された時の元の入出荷数量が維持される。変更不可。                                                                                                                                                                     |     | 
| DeliveryQuantityUnit                     | 入出荷数量単位。オーダー明細参照で設定される。変更不可。                                                                                                                                                                                                 | ✔  | 
| ActualGoodsIssueDate                     | 実際在庫出庫日付。実際に在庫を出庫した日付を入力する。                                                                                                                                                                                                   |     | 
| ActualGoodsIssueTime                     | 実際在庫出庫時刻。実際に在庫を出庫した時刻を入力する。                                                                                                                                                                                                   |     | 
| ActualGoodsReceiptDate                   | 実際在庫入庫日付。実際に在庫を入庫した日付を入力する。                                                                                                                                                                                                   |     | 
| ActualGoodsReceiptTime                   | 実際在庫入庫時刻。実際に在庫を入庫した時刻を入力する。                                                                                                                                                                                                   |     | 
| ActualGoodsIssueQtyInBaseUnit            | 実際出荷数量（基本単位）。オーダー明細参照で設定される。必要に応じて変更する。                                                                                                                                                                           |     | 
| ActualGoodsIssueQuantity                 | 実際出荷数量（入出荷数量単位）。オーダー明細参照で設定される。必要に応じて変更する。                                                                                                                                                                     |     | 
| ActualGoodsReceiptQtyInBaseUnit          | 実際入荷数量（基本単位）。オーダー明細参照で設定される。必要に応じて変更する。                                                                                                                                                                           |     | 
| ActualGoodsReceiptQuantity               | 実際入荷数量（入出荷数量単位）。オーダー明細参照で設定される。必要に応じて変更する。                                                                                                                                                                     |     | 
| CompleteItemDeliveryIsDefined            | 完納フラグ。入出荷伝票明細に対して完全に入出荷が完了しているときにtrueが設定される。変更不可。                                                                                                                                                           |     | 
| StockConfirmationPartnerFunction         | 在庫確認を行う取引先機能。オーダー明細参照で設定される。必要に応じて変更する。（PRODUCT_STOCK_AVAIRABILITY_CHECK_SRVを利用、機能はオプション）                                                                                                           | ✔  | 
| StockConfirmationBusinessPartner         | 在庫確認を行うビジネスパートナ。オーダー明細参照で設定される。必要に応じて変更する。（PRODUCT_STOCK_AVAIRABILITY_CHECK_SRVを利用、機能はオプション）                                                                                                     | ✔  | 
| StockConfirmationPlant                   | 在庫確認プラント。オーダー明細参照で設定される。必要に応じて変更する。（PRODUCT_STOCK_AVAIRABILITY_CHECK_SRVを利用、機能はオプション）                                                                                                                   | ✔  | 
| StockConfirmationPolicy                  | 在庫確認方針。オーダー明細参照で設定される。必要に応じて変更する。（PRODUCT_STOCK_AVAIRABILITY_CHECK_SRVを利用、機能はオプション）                                                                                                                       | ✔  | 
| StockConfirmationStatus                  | 在庫確認ステータス。次の在庫確認ステータスが常時更新される。変更不可。 <br> CL:全部完了済  PP:部分完了済 <br> NP:未確認                                |       |                                                                                                                                                                                                                                                   | 
| ProductionPlantPartnerFunction           | 製造プラント取引先機能。オーダー明細から参照される。変更不可。                                                                                                                                                                                           |     | 
| ProductionPlantBusinessPartner           | 製造プラントビジネスパートナ。オーダー明細から参照される。変更不可。                                                                                                                                                                                     |     | 
| ProductionPlant                          | 製造プラント。オーダー明細から参照される。変更不可。                                                                                                                                                                                                     |     | 
| ProductionPlantStorageLocation           | 製造プラントの保管場所。オーダー明細から参照される。必要に応じて保管場所マスタから更新する。                                                                                                                                                             | ✔  | 
| ProductionPlantPartnerFunction           | 出荷プラント取引先機能。入出荷伝票ヘッダから参照される。変更不可。                                                                                                                                                                                       |     | 
| IssuingPlantBusinessPartner              | 出荷プラントビジネスパートナ。入出荷伝票ヘッダから参照される。変更不可。                                                                                                                                                                                 |     | 
| IssuingPlant                             | 出荷プラント。入出荷伝票ヘッダから参照される。変更不可。                                                                                                                                                                                                 |     | 
| IssuingPlantStorageLocation              | 出荷プラントの保管場所。オーダー明細から参照される。必要に応じて保管場所マスタから更新する。                                                                                                                                                             | ✔  | 
| ReceivingPlantPartnerFunction            | 入荷プラント取引先機能。入出荷伝票ヘッダから参照される。変更不可。                                                                                                                                                                                       |     | 
| ReceivingPlantBusinessPartner            | 入荷プラントビジネスパートナ。入出荷伝票ヘッダから参照される。変更不可。                                                                                                                                                                                 |     | 
| ReceivingPlant                           | 入荷プラント。入出荷伝票ヘッダから参照される。変更不可。                                                                                                                                                                                                 |     | 
| ReceivingPlantStorageLocation            | 入荷プラントの保管場所。オーダー明細から参照される。必要に応じて保管場所マスタから更新する。                                                                                                                                                             | ✔  | 
| ProductIsBatchManagedInProductionPlant   | ロット管理フラグ。品目マスタBPプラントデータから参照される。変更不可。                                                                                                                                                                                   |     | 
| ProductIsBatchManagedInIssuingPlant      | ロット管理フラグ。品目マスタBPプラントデータから参照される。変更不可。                                                                                                                                                                                   |     | 
| ProductIsBatchManagedInReceivingPlant    | ロット管理フラグ。品目マスタBPプラントデータから参照される。変更不可。                                                                                                                                                                                   |     | 
| BatchMgmtPolicyInProductionPlant         | ロット管理方針。品目マスタBPプラントデータから参照される。変更不可。                                                                                                                                                                                     |     | 
| BatchMgmtPolicyInIssuingPlant            | ロット管理方針。品目マスタBPプラントデータから参照される。変更不可。                                                                                                                                                                                     |     | 
| BatchMgmtPolicyInReceivingPlant          | ロット管理方針。品目マスタBPプラントデータから参照される。変更不可。                                                                                                                                                                                     |     | 
| ProductionPlantBatch                     | 製造プラントロット番号。ロット管理フラグとロット管理方針に基づき、オーダー明細から参照されるか、必要な場合、オーダー時点で製造プラントのロット番号を入力する、または、ロットが自動決定される。（※ロットマスタに無い場合新規ロットとして自動登録される） | ✔  | 
| IssuingPlantBatch                        | 出荷プラントロット番号。ロット管理フラグとロット管理方針に基づき、オーダー明細から参照されるか、必要な場合、オーダー時点で出荷プラントのロット番号を入力する、または、ロットが自動決定される。（※ロットマスタに無い場合新規ロットとして自動登録される） | ✔  | 
| ReceivingPlantBatch                      | 入荷プラントロット番号。ロット管理フラグとロット管理方針に基づき、オーダー明細から参照されるか、必要な場合、オーダー時点で入荷プラントのロット番号を入力する、または、ロットが自動決定される。（※ロットマスタに無い場合新規ロットとして自動登録される） | ✔  | 
| ProductionPlantBatchValidityStartDate    | 製造プラントロット有効開始日付。ロット管理フラグとロット管理方針に基づき、オーダー明細から参照されるか、必要な場合、製造プラントのロットの有効開始日付を入力する、または、有効開始日付が自動決定される。                                                 |     | 
| ProductionPlantBatchValidityEndDate      | 製造プラントロット有効終了日付。ロット管理フラグとロット管理方針に基づき、オーダー明細から参照されるか、必要な場合、製造プラントのロットの有効終了日付を入力する、または、有効終了日付が自動決定される。                                                 |     | 
| IssuingPlantBatchValidityStartDate       | 出荷プラントロット有効開始日付。ロット管理フラグとロット管理方針に基づき、オーダー明細から参照されるか、必要な場合、出荷プラントのロットの有効開始日付を入力する、または、有効開始日付が自動決定される。                                                 |     | 
| IssuingPlantBatchValidityEndDate         | 出荷プラントロット有効終了日付。ロット管理フラグとロット管理方針に基づき、オーダー明細から参照されるか、必要な場合、出荷プラントのロットの有効終了日付を入力する、または、有効終了日付が自動決定される。                                                 |     | 
| ReceivingPlantBatchValidityStartDate     | 入荷プラントロット有効開始日付。ロット管理フラグとロット管理方針に基づき、オーダー明細から参照されるか、必要な場合、入荷プラントのロットの有効開始日付を入力する、または、有効開始日付が自動決定される。                                                 |     | 
| ReceivingPlantBatchValidityEndDate       | 入荷プラントロット有効終了日付。ロット管理フラグとロット管理方針に基づき、オーダー明細から参照されるか、必要な場合、入荷プラントのロットの有効終了日付を入力する、または、有効終了日付が自動決定される。                                                 |     | 
| CreationDate                             | 登録日。システムにより自動生成される。変更不可。                                                                                                                                                                                                         |     | 
| CreationTime                             | 登録時刻。システムにより自動生成される。変更不可。                                                                                                                                                                                                       |     | 
| ItemBillingStatus                        | 明細請求ステータス。<br> 完全請求完了:CL <br> 部分請求完了:PP   <br> 未請求:NP <br> が設定される。変更不可。                 |        |                                                                                                                                                                                                                                          |        | 
| ItemBillingConfStatus                    | 明細請求照合ステータス。<br> 完全請求照合完了:CL  <br> 部分請求照合完了:PP  <br> 未請求照合:NP<br> が設定される。変更不可。                 |                                                                                                                                                                                                                                                          | 
| SalesCostGLAccount                       | 売上原価総勘定元張勘定。自動設定される。変更不可。                                                                                                                                                                                                       |     | 
| ReceivingGLAccount                       | 受入総勘定元帳勘定。品目在庫もしくは仕入原価の勘定コードが設定される。変更不可。                                                                                                                                                                         |     | 
| IssuingGoodsMovementType                 | 出荷在庫移動タイプ。次の在庫移動タイプが設定される。変更不可。<br> 601:売上出庫(オーダー受注参照出荷)  <br> 111:仕入返品(発注返品リクエスト参照出荷) |          |                                                                                                                                                                                                                                                | 
| ReceivingGoodsMovementType               | 入荷在庫移動タイプ。次の在庫移動タイプが設定される。変更不可。 <br> 101:発注参照入庫(オーダー発注参照入荷) <br> 611:売上返品(受注返品リクエスト参照入荷) |          |                                                                                                                                                                                                                                                | 
| ItemBillingBlockReason                   | 明細請求ブロック理由。請求をブロックする場合、trueを入力する。                                                                                                                                                                                           |     | 
| ItemDeliveryIncompletionStatus           | 明細入出荷不完全ステータス。システムにより設定される。変更不可。<br> 在庫未確認：NC <br> 未入出庫：NP                             |       |                                                                                                                                                                                                                                                   | 
| ItemGrossWeight                          | 明細総重量。オーダー明細から参照される。必要に応じて入力する。                                                                                                                                                                                           |     | 
| ItemNetWeight                            | 明細正味重量。オーダー明細から参照される。必要に応じて入力する。                                                                                                                                                                                         |     | 
| ItemWeightUnit                           | 明細重量単位。オーダー明細から参照される。必要に応じて入力する。                                                                                                                                                                                         | ✔  | 
| ItemVolume                               | 明細数量。項目としては存在するが不使用。                                                                                                                                                                                                                 |     | 
| ItemVolumeUnit                           | 明細数量単位。項目としては存在するが不使用。                                                                                                                                                                                                             |     | 
| ItemIsBillingRelevant                    | 請求関連明細。オーダー明細で” BillingDocumentDate”に値があるものについてtrueが設定される。                                                                                                                                                             |     | 
| LastChangeDate                           | 最終更新日付。システムにより設定される。変更不可。                                                                                                                                                                                                       |     | 
| OrderID                                  | オーダー番号。入出荷伝票登録のために参照されるオーダー番号。                                                                                                                                                                                             | ✔  | 
| OrderItem                                | オーダー明細番号。入出荷伝票登録のために参照されるオーダー明細番号。                                                                                                                                                                                     | ✔  | 
| OrderType                                | オーダータイプ。オーダーから参照される。変更不可。<br> OD:オーダー <br> CT:契約<br> PR:返品リクエスト <br> DR:デビットメモリクエスト <br> CR:クレジットメモリクエスト              |        |                                                                                                                                                                                                                                                  | 
| ContractType                             | 契約タイプ。ヘッダからコピーされる。変更不可。<br> MNTH:月額契約。 <br> ANNL:年額契約。                          |        |                                                                                                                                                                                                                                                  | 
| OrderVaridityStartDate                   | オーダー契約開始日付。ヘッダからコピーされる。変更不可。                                                                                                                                                                                                 |     | 
| OrderValidityEndDate                     | オーダー契約終了日付。ヘッダからコピーされる。変更不可。                                                                                                                                                                                                 |     | 
| InvoiceScheduleStartDate                 | 請求計画開始日付。ヘッダからコピーされる。変更不可。                                                                                                                                                                                                     |     | 
| InvoiceScheduleEndDate                   | 請求計画終了日付。ヘッダからコピーされる。変更不可。                                                                                                                                                                                                     |     | 
| ProductAvailabilityDate                  | 品目利用可能日付。利用可能在庫確認により確認された品目の利用可能日付。変更不可。（PRODUCT_STOCK_AVAIRABILITY_CHECK_SRVを利用、機能はオプション）                                                                                                         |     | 
| Project                                  | プロジェクトコード。オーダー明細から参照される。変更不可。                                                                                                                                                                                               |     | 
| ReferenceDocument                        | 参照された伝票番号。変更不可。                                                                                                                                                                                                                           |     | 
| ReferenceDocumentItem                    | 参照された伝票明細番号。変更不可。                                                                                                                                                                                                                       |     | 
| BPTaxClassification                      | ビジネスパートナ税分類。オーダー明細から参照される。変更不可。                                                                                                                                                                                           |     | 
| ProductTaxClassification                 | 品目税分類。オーダー明細から参照される。変更不可。                                                                                                                                                                                                       |     | 
| BPAccountAssignmentGroup                 | ビジネスパートナ勘定設定グループ。オーダー明細から参照される。変更不可。                                                                                                                                                                                 |     | 
| ProductAccountAssignmentGroup            | 品目勘定設定グループ。オーダー明細から参照される。変更不可。                                                                                                                                                                                             |     | 
| TaxCode                                  | 消費税コード。オーダー明細から参照される。変更不可。                                                                                                                                                                                                     |     | 
| TaxRate                                  | 消費税率。オーダー明細から参照される。変更不可。                                                                                                                                                                                                         |     | 
| CountryOfOrigin                          | 原産国。オーダー明細から参照される。変更不可。                                                                                                                                                                                                           |     | 

### API名:【Delivery Document Item Partner > A_ItemPartner】
### Accepter名:【ItemPartner】

### <項目>
以下の表には、data-platform-api-delivery-document-creates-rmq-kube のデータの項目とその定義が記載されています。  
※ EC の列には、そのデータの存在性確認を行うと定めているかどうか（✔がついていれば、行うと定めている）が記載されています。

| Property                   | Description                                                                                                                                                        | EC  | 
| -------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | --- | 
| PartnerFunction            | 取引先機能。オーダー明細から参照される。必要に応じて変更する。入出荷伝票明細において、それぞれの取引先機能レコードはオーダーに対して下記の()内の数だけ設定される。 <br> 提案される取引先機能: <br> MANUFACTURER:製造者(一つ)  <br> DELIVERFROM:入出荷元(一つ) <br> DELIVERTO:入出荷先(一つ)  <br> LOGI:物流業者(一つ)        | ✔  |                                                                                                                                                               | 
| BusinessPartner            | ビジネスパートナコード。オーダー明細の取引先機能に対応するビジネスパートナコードが設定される。必要に応じて更新する。                                               | ✔  | 

### API名:【Delivery Document Item Partner Plant > A_ItemPartnerPlant】
### Accepter名:【ItemPartnerPlant】

### <項目>
以下の表には、data-platform-api-delivery-document-creates-rmq-kube のデータの項目とその定義が記載されています。  
※ EC の列には、そのデータの存在性確認を行うと定めているかどうか（✔がついていれば、行うと定めている）が記載されています。

| Property             | Description                                                                                                                                          | EC  | 
| -------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | --- | 
| PartnerFunction      | 取引先機能。ヘッダパートナ取引先プラントの取引先機能が設定される。変更不可。<br> MANUFACTURER:製造者 <br> DELIVERFROM:入出荷元 <br> DELIVERTO:入出荷先   |       |                                                                                                                                               | 
| BusinessPartner      | ビジネスパートナコード。ヘッダパートナプラントの上記取引先機能に対してヘッダパートナプラントのビジネスパートナが設定される。変更不可。               |     | 
| Plant                | プラント。ヘッダパートナのヘッダパートナプラントの上記取引先機能に対してヘッダパートナプラントのビジネスパートナが設定される。必要に応じて変更する。 | ✔  | 

