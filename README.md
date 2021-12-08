# sap-api-integrations-purchasing-info-record-reads
sap-api-integrations-purchasing-info-record-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で購買情報レコードを取得するマイクロサービスです。    
sap-api-integrations-purchasing-info-record-reads には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-purchasing-info-record-reads は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。   
https://api.sap.com/api/OP_API_INFORECORD_PROCESS_SRV_0001/overview   

## 動作環境  
sap-api-integrations-purchasing-info-record-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）    

## クラウド環境での利用
sap-api-integrations-purchasing-info-record-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-purchasing-info-record-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_INFORECORD_PROCESS_SRV_0001/overview    
* APIサービス名(=baseURL): API_INFORECORD_PROCESS_SRV

## 本レポジトリ に 含まれる API名
sap-api-integrations-purchasing-info-record-reads には、次の API をコールするためのリソースが含まれています。  

* A_PurchasingInfoRecord（購買情報 - 一般）
* A_PurgInfoRecdOrgPlantData（購買情報 - 購買組織プラント）

## API への 値入力条件 の 初期値
sap-api-integrations-purchasing-info-record-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.PurchasingInfoRecord.PurchasingInfoRecord（購買情報）
* inoutSDC.PurchasingInfoRecord.PurchasingOrganizationPlant.PurchasingInfoRecordCategory（購買情報カテゴリ）
* inoutSDC.PurchasingInfoRecord.PurchasingOrganizationPlant.Supplier（仕入先）
* inoutSDC.PurchasingInfoRecord.PurchasingOrganizationPlant.Material（品目）
* inoutSDC.PurchasingInfoRecord.PurchasingOrganizationPlant.PurchasingOrganization（購買組織）
* inoutSDC.PurchasingInfoRecord.PurchasingOrganizationPlant.Plant（プラント）
* inoutSDC.PurchasingInfoRecord.PurchasingOrganizationPlant.MaterialGroup（品目グループ）
